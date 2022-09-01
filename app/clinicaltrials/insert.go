package clinicaltrials

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/run"
	"regexp"
	"strings"
	"time"
)

const timeLayout = "Jan 2, 2006"
const timeLayoutAlt = "January 2, 2006"
const timeLayoutAlt2 = "January 2006"
var serotypes = [...]string{"23F", "23F", "22F", "19F","19A", "18C","14","9V","8","7F","6C","6B","6A","5","4","3","2","1"}
var firstDoseCheck = regexp.MustCompile(`(?mi)(\b1st Dose|\bFirst Dose|\bone dose|\b1 dose)`)
var firstDoseCheckFlow = regexp.MustCompile(`(?mi)(\Wdose 1)`)
var secondDoseCheck = regexp.MustCompile(`(?mi)(\b2nd Dose|\bSecond Dose|\btwo doses|\b2 doses)`)
var secondDoseCheckFlow = regexp.MustCompile(`(?mi)(\Wdose 2)`)
var thirdDoseCheck = regexp.MustCompile(`(?mi)(\b3rd Dose|\bThird Dose|\bthree doses|\b3 doses)`)
var thirdDoseCheckFlow = regexp.MustCompile(`(?mi)(\Wdose 3)`)
var fourthDoseCheck = regexp.MustCompile(`(?mi)(\b4th Dose|\bFourth Dose|\bfour doses|\b4 doses)`)
var fourthDoseCheckFlow = regexp.MustCompile(`(?mi)(\Wdose 4)`)
var gmcCheck = regexp.MustCompile(`(?mi)(\bGeometric Mean Concentration|\bGMC|\bGeometric Mean)`)
var gmtCheck = regexp.MustCompile(`(?mi)(\bGeometric Mean Titer|\bGMT)`)
var opaCheck = regexp.MustCompile(`(?mi)(\bo-phthalaldehyde|\bOPA)`)

type ResultsContext struct {
	client *models.Client
	ctx    context.Context
}

func NewResultsContext(ctx context.Context, client *models.Client) *ResultsContext {
	return &ResultsContext{
		client: client,
		ctx:    ctx,
	}
}

type Response struct {
	FullStudiesResponse struct{
		FullStudies []struct{
			Study Study `json:"Study"`
		} `json:"FullStudies"`
	} `json:"FullStudiesResponse"`
}

const apiURL = "https://clinicaltrials.gov/api/query/full_studies?expr=%s&fmt=JSON"

func (t *ResultsContext) InsertAll(study Study, outcomeMeasures []OutcomeMeasureCustom, overviews []OutcomeOverview, metadata []Metadata, useCase string) error {
	exists, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(study.ProtocolSection.IdentificationModule.NCTID)).Exist(t.ctx)
	if err != nil {
		return err
	}

	if exists {
		return t.EditExistingTrial(study.ProtocolSection.IdentificationModule.NCTID, overviews, metadata, useCase, outcomeMeasures)
	}

	if err := t.InsertOutcomeOverviews(overviews); err != nil {
		return err
	}

	return t.Insert(study, metadata, outcomeMeasures)
}

func (t *ResultsContext) Insert(study Study, metadata []Metadata, outcomeMeasures []OutcomeMeasureCustom) error {
	phase := ""

	for i, p := range study.ProtocolSection.DesignModule.PhaseList.Phase {
		if i > 0 {
			phase += ", "
		}

		phase += p
	}

	startDate, err := getDate(study.ProtocolSection.StatusModule.StartDateStruct.StartDate)
	if err != nil {
		return err
	}

	primaryCompletionDate, err := getDate(study.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.PrimaryCompletionDate)
	if err != nil {
		return err
	}

	studyCompletionDate, err := getDate(study.ProtocolSection.StatusModule.CompletionDateStruct.CompletionDate)
	if err != nil {
		return err
	}

	clinicalTrial, err := t.client.ClinicalTrial.Create().
		SetStudyName(study.ProtocolSection.IdentificationModule.BriefTitle).
		SetSponsor(study.ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorName).
		SetResponsibleParty(study.ProtocolSection.IdentificationModule.Organization.OrgFullName).
		SetStudyType(study.ProtocolSection.DesignModule.StudyType).
		SetPhase(phase).
		SetActualEnrollment(study.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount).
		SetAllocation(study.ProtocolSection.DesignModule.DesignInfo.DesignAllocation).
		SetInterventionModel(study.ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel).
		SetMasking(study.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking).
		SetPrimaryPurpose(study.ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose).
		SetOfficialTitle(study.ProtocolSection.IdentificationModule.OfficialTitle).
		SetActualStudyStartDate(startDate).
		SetActualPrimaryCompletionDate(primaryCompletionDate).
		SetActualStudyCompletionDate(studyCompletionDate).
		SetStudyID(study.ProtocolSection.IdentificationModule.NCTID).
		Save(t.ctx)
	if err != nil {
		return fmt.Errorf("unable to create clinical trial, err: %s", err)
	}

	resultsDefinition, err := t.client.ResultsDefinition.Create().
		SetParent(clinicalTrial).
		Save(t.ctx)
	if err != nil {
		if err := t.client.ClinicalTrial.DeleteOne(clinicalTrial).Exec(t.ctx); err != nil {
			return fmt.Errorf("unable to delete clinical trial, err: %s", err)
		}
		return fmt.Errorf("unable to create results definition, err: %s", err)
	}

	if err := run.Go(t.ctx,
		func() error {return t.insertParticipantFlowModule(resultsDefinition, study)},
		func() error {return t.insertBaselineCharacteristicsModule(resultsDefinition, study)},
		func() error {return t.insertOutcomeMeasureModule(resultsDefinition, study, outcomeMeasures)},
		func() error {return t.insertAdverseEventsModule(resultsDefinition, study)},
		func() error {return t.insertMoreInfoModule(resultsDefinition, study)},
		func() error {return t.insertStudyLocations(clinicalTrial, study)},
		func() error {return t.insertStudyEligibility(clinicalTrial, study)},
		func() error {return t.insertMetadata(clinicalTrial, metadata)},
		); err != nil {
		log.Err(err)
		if err := t.cleanClinicalTrial(clinicalTrial); err != nil {
			return fmt.Errorf("unable to delete delete data, err: %s", err)
		}

		return fmt.Errorf("unable to insert results data, err: %s", err)
	}

	return nil
}

func (t *ResultsContext) insertMetadata(clinicalTrial *models.ClinicalTrial, metadata []Metadata) error {
	if len(metadata) == 0 {
		return nil
	}

	if _, err := t.client.ClinicalTrialMetadata.Delete().Where(
			clinicaltrialmetadata.HasParentWith(clinicaltrial.StudyIDEQ(clinicalTrial.StudyID)),
			clinicaltrialmetadata.UseCaseCode(metadata[0].TagUseCase),
		).Exec(t.ctx); err != nil {
		return err
	}

	for _, datum := range metadata {
		if _, err := t.client.ClinicalTrialMetadata.Create().
			SetTagName(datum.TagName).
			SetTagValue(datum.TagValue).
			SetUseCaseCode(datum.TagUseCase).
			SetParent(clinicalTrial).
			Save(t.ctx); err != nil {
			return err
		}
	}

	return nil
}

func (t *ResultsContext) insertStudyLocations(clinicalTrial *models.ClinicalTrial, study Study) error {
	for _, studyLocations := range study.ProtocolSection.ContactsLocationsModule.LocationList.Location {
		cc := ""

		switch studyLocations.LocationCountry {
			case "Andorra": cc = "AD"; break
			case "United Arab Emirates": cc = "AE"; break
			case "Afghanistan": cc = "AF"; break
			case "Antigua and Barbuda": cc = "AG"; break
			case "Anguilla": cc = "AI"; break
			case "Albania": cc = "AL"; break
			case "Armenia": cc = "AM"; break
			case "Angola": cc = "AO"; break
			case "Antarctica": cc = "AQ"; break
			case "Argentina": cc = "AR"; break
			case "American Samoa": cc = "AS"; break
			case "Austria": cc = "AT"; break
			case "Australia": cc = "AU"; break
			case "Aruba": cc = "AW"; break
			case "Åland Islands": cc = "AX"; break
			case "Azerbaijan": cc = "AZ"; break
			case "Bosnia and Herzegovina": cc = "BA"; break
			case "Barbados": cc = "BB"; break
			case "Bangladesh": cc = "BD"; break
			case "Belgium": cc = "BE"; break
			case "Burkina": cc = "BF"; break
			case "Bulgaria": cc = "BG"; break
			case "Bahrain": cc = "BH"; break
			case "Burundi": cc = "BI"; break
			case "Benin": cc = "BJ"; break
			case "Saint Barthélemy": cc = "BL"; break
			case "Bermuda": cc = "BM"; break
			case "Brunei": cc = "BN"; break
			case "Bolivia": cc = "BO"; break
			case "Bonaire": cc = "BQ"; break
			case "Brazil": cc = "BR"; break
			case "Bahamas": cc = "BS"; break
			case "Bhutan": cc = "BT"; break
			case "Bouvet Island": cc = "BV"; break
			case "Botswana": cc = "BW"; break
			case "Belarus": cc = "BY"; break
			case "Belize": cc = "BZ"; break
			case "Canada": cc = "CA"; break
			case "Cocos (Keeling) Islands": cc = "CC"; break
			case "Democratic Republic of the Congo": cc = "CD"; break
			case "Central African Republic": cc = "CF"; break
			case "Congo": cc = "CG"; break
			case "Switzerland": cc = "CH"; break
			case "Ivory Coast": cc = "CI"; break
			case "Cook Islands": cc = "CK"; break
			case "Chile": cc = "CL"; break
			case "Cameroon": cc = "CM"; break
			case "China": cc = "CN"; break
			case "Colombia": cc = "CO"; break
			case "Costa": cc = "CR"; break
			case "Cuba": cc = "CU"; break
			case "Cabo Verde": cc = "CV"; break
			case "Curaçao": cc = "CW"; break
			case "Christmas Island": cc = "CX"; break
			case "Cyprus": cc = "CY"; break
			case "Czech Republic": cc = "CZ"; break
			case "Germany": cc = "DE"; break
			case "Djibouti": cc = "DJ"; break
			case "Denmark": cc = "DK"; break
			case "Dominica": cc = "DM"; break
			case "Dominican Republic": cc = "DO"; break
			case "Algeria": cc = "DZ"; break
			case "Ecuador": cc = "EC"; break
			case "Estonia": cc = "EE"; break
			case "Egypt": cc = "EG"; break
			case "Western Sahara": cc = "EH"; break
			case "Eritrea": cc = "ER"; break
			case "Spain": cc = "ES"; break
			case "Ethiopia": cc = "ET"; break
			case "Finland": cc = "FI"; break
			case "Fiji": cc = "FJ"; break
			case "Falkland Islands": cc = "FK"; break
			case "Micronesia (Federated States of)": cc = "FM"; break
			case "Faroe": cc = "FO"; break
			case "France": cc = "FR"; break
			case "Gabon": cc= "GA"; break
			case "United Kingdom of Great Britain and Northern Ireland": cc = "GB"; break
			case "Grenada": cc = "GD"; break
			case "Georgia": cc = "GE"; break
			case "French Guiana": cc = "GF"; break
			case "Guernsey": cc = "GG"; break
			case "Ghana": cc = "GH"; break
			case "Gibraltar": cc = "GI"; break
			case "Greenland": cc = "GL"; break
			case "Gambia": cc = "GM"; break
			case "Guinea": cc = "GN"; break
			case "Guadeloupe": cc = "GP"; break
			case "Equatorial Guinea": cc = "GQ"; break
			case "Greece": cc = "GR"; break
			case "South Georgia and the South Sandwich Islands": cc = "GS"; break
			case "Guatemala": cc = "GT"; break
			case "Guam": cc = "GU"; break
			case "Guinea-Bissau": cc = "GW"; break
			case "Guyana": cc = "GY"; break
			case "Hong": cc = "HK"; break
			case "Hong Kong": cc = "HK"; break
			case "Heard Island and McDonald Islands": cc = "HM"; break
			case "Honduras": cc = "HN"; break
			case "Croatia": cc = "HR"; break
			case "Haiti": cc = "HT"; break
			case "Hungary": cc = "HU"; break
			case "Indonesia": cc = "ID"; break
			case "Ireland": cc = "IE"; break
			case "Israel": cc = "IL"; break
			case "Isle of Man": cc = "IM"; break
			case "India": cc = "IN"; break
			case "British Indian Ocean Territory": cc = "IO"; break
			case "Iraq": cc = "IQ"; break
			case "Iran": cc = "IR"; break
			case "Iceland": cc = "IS"; break
			case "Italy": cc = "IT"; break
			case "Jersey": cc = "JE"; break
			case "Jamaica": cc = "JM"; break
			case "Jordan": cc = "JO"; break
			case "Japan": cc = "JP"; break
			case "Kenya": cc = "KE"; break
			case "Kyrgyzstan": cc = "KG"; break
			case "Cambodia": cc = "KH"; break
			case "Kiribati": cc = "KI"; break
			case "Comoros Kitts and Nevis": cc = "KM"; break
			case "Saint": cc = "KN"; break
			case "North Korea": cc = "KP"; break
			case "South Korea": cc = "KR"; break
			case "Kuwait": cc = "KW"; break
			case "Cayman Islands": cc = "KY"; break
			case "Kazakhstan": cc = "KZ"; break
			case "Laos": cc = "LA"; break
			case "Lebanon": cc = "LB"; break
			case "Saint Lucia": cc = "LC"; break
			case "Liechtenstein": cc = "LI"; break
			case "Sri Lanka": cc = "LK"; break
			case "Liberia": cc = "LR"; break
			case "Lesotho": cc = "LS"; break
			case "Lithuania": cc = "LT"; break
			case "Luxembourg": cc = "LU"; break
			case "Latvia": cc = "LV"; break
			case "Libya": cc = "LY"; break
			case "Morocco": cc = "MA"; break
			case "Monaco": cc = "MC"; break
			case "Moldova": cc = "MD"; break
			case "Montenegro": cc = "ME"; break
			case "Saint Martin": cc = "MF"; break
			case "Madagascar": cc = "MG"; break
			case "Marshall Islands": cc = "MH"; break
			case "North Macedonia": cc = "MK"; break
			case "Mali": cc = "ML"; break
			case "Myanmar": cc = "MM"; break
			case "Mongolia": cc = "MN"; break
			case "Macao": cc = "MO"; break
			case "Northern Mariana Islands": cc = "MP"; break
			case "Martinique": cc = "MQ"; break
			case "Mauritania": cc = "MR"; break
			case "Montserrat": cc = "MS"; break
			case "Malta": cc = "MT"; break
			case "Mauritius": cc = "MU"; break
			case "Maldives": cc = "MV"; break
			case "Malawi": cc = "MW"; break
			case "Mexico": cc = "MX"; break
			case "Malaysia": cc = "MY"; break
			case "Mozambique": cc = "MZ"; break
			case "Namibia": cc = "NA"; break
			case "New Caledonia": cc = "NC"; break
			case "Niger": cc = "NE"; break
			case "Norfolk Island": cc = "NF"; break
			case "Nigeria": cc = "NG"; break
			case "Nicaragua": cc = "NI"; break
			case "Netherlands": cc = "NL"; break
			case "Norway": cc = "NO"; break
			case "Nepal": cc = "NP"; break
			case "Nauru": cc = "NR"; break
			case "Niue": cc = "NU"; break
			case "New Zealand": cc = "NZ"; break
			case "Oman": cc = "OM"; break
			case "Panama": cc = "PA"; break
			case "Peru": cc = "PE"; break
			case "French Polynesia": cc = "PF"; break
			case "Papua New Guinea": cc = "PG"; break
			case "Philippines": cc = "PH"; break
			case "Pakistan": cc = "PK"; break
			case "Poland": cc = "PL"; break
			case "Saint Pierre and Miquelon": cc = "PM"; break
			case "Pitcairn": cc = "PN"; break
			case "Puerto Rico": cc = "PR"; break
			case "Palestine": cc = "PS"; break
			case "Portugal": cc = "PT"; break
			case "Palau": cc = "PW"; break
			case "Paraguay": cc = "PY"; break
			case "Qatar": cc = "QA"; break
			case "Réunion": cc = "RE"; break
			case "Romania": cc = "RO"; break
			case "Serbia": cc = "RS"; break
			case "Russia": cc = "RU"; break
			case "Rwanda": cc = "RW"; break
			case "Saudi Arabia": cc = "SA"; break
			case "Solomon Islands": cc = "SB"; break
			case "Seychelles": cc = "SC"; break
			case "Sudan": cc = "SD"; break
			case "Sweden": cc = "SE"; break
			case "Singapore": cc = "SG"; break
			case "Saint Helen": cc = "SH"; break
			case "Slovenia": cc = "SI"; break
			case "Svalbard and Jan Mayen": cc = "SJ"; break
			case "Slovakia": cc = "SK"; break
			case "Sierra Leone": cc = "SL"; break
			case "San Marino": cc = "SM"; break
			case "Senegal": cc = "SN"; break
			case "Somalia": cc = "SO"; break
			case "Suriname": cc = "SR"; break
			case "South": cc = "SS"; break
			case "Sao Tome and Principe": cc = "ST"; break
			case "El": cc = "SV"; break
			case "Sint Maarten": cc = "SX"; break
			case "Syria": cc = "SY"; break
			case "Eswatini": cc = "SZ"; break
			case "Turks and Caicos Islands": cc = "TC"; break
			case "Chad": cc = "TD"; break
			case "French Southern Territories": cc = "TF"; break
			case "Togo": cc = "TG"; break
			case "Thailand": cc = "TH"; break
			case "Tajikistan": cc = "TJ"; break
			case "Tokelau": cc = "TK"; break
			case "Timor-Leste": cc = "TL"; break
			case "Turkmenistan": cc = "TM"; break
			case "Tunisia": cc = "TN"; break
			case "Tonga": cc = "TO"; break
			case "Turkey": cc = "TR"; break
			case "Trinidad and Tobago": cc = "TT"; break
			case "Tuvalu": cc = "TV"; break
			case "Taiwan": cc = "TW"; break
			case "Tanzania": cc = "TZ"; break
			case "Ukraine": cc = "UA"; break
			case "Uganda": cc = "UG"; break
			case "United States Minor Outlying Islands": cc = "UM"; break
			case "United States of America": cc = "US"; break
			case "Uruguay": cc = "UY"; break
			case "Uzbekistan": cc = "UZ"; break
			case "Holy See": cc = "VA"; break
			case "Saint Vincent and the Grenadines": cc = "VC"; break
			case "Venezuela": cc = "VE"; break
			case "Virgin Islands (British)": cc = "VG"; break
			case "Virgin Islands (U.S.)": cc = "VI"; break
			case "Vietnam": cc = "VN"; break
			case "Vanuatu": cc = "VU"; break
			case "Wallis and Futuna": cc = "WF"; break
			case "Samoa": cc = "WS"; break
			case "Yemen": cc = "YE"; break
			case "Mayotte": cc = "YT"; break
			case "Kosovo": cc = "XK"; break
			case "South Africa": cc = "ZA"; break
			case "Zambia": cc = "ZM"; break
			case "Zimbabwe": cc = "ZW"; break
			case "United Kingdom": cc = "UK"; break
			case "United States": cc = "US"; break
			case "Russian Federation": cc = "RU"; break
		default:
			cc = ""
		}

		continents, continentCodes := ccToContinent(cc)

		_, err := t.client.StudyLocation.Create().
			SetLocationFacility(studyLocations.LocationFacility).
			SetLocationCity(studyLocations.LocationCity).
			SetLocationCountry(studyLocations.LocationCountry).
			SetLocationCountryCode(cc).
			SetLocationContinentName(continents).
			SetLocationContinentCode(continentCodes).
			SetParent(clinicalTrial).
			Save(t.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *ResultsContext) insertStudyEligibility(clinicalTrial *models.ClinicalTrial, study Study) error {
	stdAgeJson, err := json.Marshal(study.ProtocolSection.EligibilityModule.StdAgeList.StdAge)
	if err != nil {
		return err
	}

	ethnicity := []string{}

	for _, baselineMeasure := range study.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure {
		if !strings.Contains(strings.ToLower(baselineMeasure.BaselineMeasureTitle),"race") && !strings.Contains(strings.ToLower(baselineMeasure.BaselineMeasureTitle),"ethnicity") {
			continue
		}

		for _, baselineClass := range baselineMeasure.BaselineClassList.BaselineClass {
			if baselineClass.BaselineClassTitle != "" {
				ethnicity = append(ethnicity, baselineClass.BaselineClassTitle)
				continue
			}

			for _, baselineCategory := range baselineClass.BaselineCategoryList.BaselineCategory {
				if baselineCategory.BaselineCategoryTitle != "" {
					ethnicity = append(ethnicity, baselineCategory.BaselineCategoryTitle)
				}
			}
		}
	}

	ethnicityJson, err := json.Marshal(ethnicity)
	if err != nil {
		return err
	}

	_, err = t.client.StudyEligibility.Create().
		SetEligibilityCriteria(study.ProtocolSection.EligibilityModule.EligibilityCriteria).
		SetHealthyVolunteers(study.ProtocolSection.EligibilityModule.HealthyVolunteers).
		SetGender(study.ProtocolSection.EligibilityModule.Gender).
		SetMinimumAge(study.ProtocolSection.EligibilityModule.MinimumAge).
		SetMaximumAge(study.ProtocolSection.EligibilityModule.MaximumAge).
		SetStdAgeList(string(stdAgeJson)).
		SetEthnicity(string(ethnicityJson)).
		SetParent(clinicalTrial).
		Save(t.ctx)

	return err
}

func (t *ResultsContext) insertParticipantFlowModule(resultsDefinition *models.ResultsDefinition, study Study) error {
	participantFlowModule := study.ResultsSection.ParticipantFlowModule

	participantFlowModuleParent, err := t.client.ParticipantFlowModule.Create().
		SetFlowPreAssignmentDetails(participantFlowModule.FlowPreAssignmentDetails).
		SetFlowRecruitmentDetails(participantFlowModule.FlowRecruitmentDetails).
		SetFlowTypeUnitsAnalyzed(participantFlowModule.FlowTypeUnitsAnalyzed).
		SetParent(resultsDefinition).
		Save(t.ctx)
	if err != nil {
		return err
	}

	for _, flowGroup := range study.ResultsSection.ParticipantFlowModule.FlowGroupList.FlowGroup {
		if _, err := t.client.FlowGroup.Create().
			SetFlowGroupID(flowGroup.FlowGroupID).
			SetFlowGroupTitle(flowGroup.FlowGroupTitle).
			SetFlowGroupDescription(flowGroup.FlowGroupDescription).
			SetParent(participantFlowModuleParent).
			Save(t.ctx); err != nil {
				return err
		}
	}

	for _, flowPeriod := range study.ResultsSection.ParticipantFlowModule.FlowPeriodList.FlowPeriod {
		flowPeriodParent, err := t.client.FlowPeriod.Create().
			SetFlowPeriodTitle(flowPeriod.FlowPeriodTitle).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, flowMilestone := range flowPeriod.FlowMilestoneList.FlowMilestone {
			flowMilestoneParent, err := t.client.FlowMilestone.Create().
				SetFlowMilestoneType(flowMilestone.FlowMilestoneType).
				SetFlowMilestoneComment(flowMilestone.FlowMilestoneComment).
				SetParent(flowPeriodParent).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, flowAchievement := range flowMilestone.FlowAchievementList.FlowAchievement {
				_, err = t.client.FlowAchievement.Create().
					SetParent(flowMilestoneParent).
					SetFlowAchievementGroupID(flowAchievement.FlowAchievementGroupID).
					SetFlowAchievementComment(flowAchievement.FlowAchievementNumUnits).
					SetFlowAchievementNumSubjects(flowAchievement.FlowAchievementNumSubjects).
					SetFlowAchievementNumUnits(flowAchievement.FlowAchievementNumUnits).
					Save(t.ctx)
				if err != nil {
					return err
				}
			}
		}

		for _, flowDropWithdraw := range flowPeriod.FlowDropWithdrawList.FlowDropWithdraw {
			flowDropWithdrawParent, err := t.client.FlowDropWithdraw.Create().
				SetParent(flowPeriodParent).
				SetFlowDropWithdrawType(flowDropWithdraw.FlowDropWithdrawType).
				SetFlowDropWithdrawComment(flowDropWithdraw.FlowDropWithdrawComment).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, flowReason := range flowDropWithdraw.FlowReasonList.FlowReason {
				_, err = t.client.FlowReason.Create().
					SetParent(flowDropWithdrawParent).
					SetFlowReasonGroupID(flowReason.FlowReasonGroupID).
					SetFlowReasonNumSubjects(flowReason.FlowReasonNumSubjects).
					SetFlowReasonNumUnits(flowReason.FlowReasonNumUnits).
					SetFlowReasonComment(flowReason.FlowReasonComment).
					Save(t.ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (t *ResultsContext) insertBaselineCharacteristicsModule(resultsDefinition *models.ResultsDefinition, study Study) error {
	baselineCharacteristicsModule := study.ResultsSection.BaselineCharacteristicsModule

	baselineCharacteristicsParent, err := t.client.BaselineCharacteristicsModule.Create().
		SetParent(resultsDefinition).
		SetBaselinePopulationDescription(baselineCharacteristicsModule.BaselinePopulationDescription).
		SetBaselineTypeUnitsAnalyzed(baselineCharacteristicsModule.BaselineTypeUnitsAnalyzed).
		Save(t.ctx)
	if err != nil {
		return err
	}

	for _, baselineGroup := range baselineCharacteristicsModule.BaselineGroupList.BaselineGroup {
		_, err := t.client.BaselineGroup.Create().
			SetParent(baselineCharacteristicsParent).
			SetBaselineGroupID(baselineGroup.BaselineGroupID).
			SetBaselineGroupTitle(baselineGroup.BaselineGroupTitle).
			SetBaselineGroupDescription(baselineGroup.BaselineGroupDescription).
			Save(t.ctx)
		if err != nil {
			return err
		}
	}

	for _, baselineDenom := range baselineCharacteristicsModule.BaselineDenomList.BaselineDenom {
		baselineDenomParent, err := t.client.BaselineDenom.Create().
			SetParent(baselineCharacteristicsParent).
			SetBaselineDenomUnits(baselineDenom.BaselineDenomUnits).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, baselineDenomCount := range baselineDenom.BaselineDenomCountList.BaselineDenomCount {
			_, err := t.client.BaselineDenomCount.Create().
				SetParent(baselineDenomParent).
				SetBaselineDenomCountGroupID(baselineDenomCount.BaselineDenomCountGroupID).
				SetBaselineDenomCountValue(baselineDenomCount.BaselineDenomCountValue).
				Save(t.ctx)
			if err != nil {
				return err
			}
		}
	}

	for _, baselineMeasure := range baselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure {
		baselineMeasureParent, err := t.client.BaselineMeasure.Create().
			SetParent(baselineCharacteristicsParent).
			SetBaselineMeasureTitle(baselineMeasure.BaselineMeasureTitle).
			SetBaselineMeasureDescription(baselineMeasure.BaselineMeasureDescription).
			SetBaselineMeasureDispersionType(baselineMeasure.BaselineMeasureDispersionType).
			SetBaselineMeasureCalculatePct(baselineMeasure.BaselineMeasureCalculatePct).
			SetBaselineMeasureParamType(baselineMeasure.BaselineMeasureParamType).
			SetBaselineMeasureUnitOfMeasure(baselineMeasure.BaselineMeasureUnitOfMeasure).
			SetBaselineMeasurePopulationDescription(baselineMeasure.BaselineMeasurePopulationDescription).
			SetBaselineMeasureDenomUnitsSelected(baselineMeasure.BaselineMeasureDenomUnitsSelected).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, baselineClass := range baselineMeasure.BaselineClassList.BaselineClass {
			baselineClassParent, err := t.client.BaselineClass.Create().
				SetParent(baselineMeasureParent).
				SetBaselineClassTitle(baselineClass.BaselineClassTitle).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, baselineCategory := range baselineClass.BaselineCategoryList.BaselineCategory {
				baselineCategoryParent, err := t.client.BaselineCategory.Create().
					SetBaselineCategoryTitle(baselineCategory.BaselineCategoryTitle).
					SetParent(baselineClassParent).
					Save(t.ctx)
				if err != nil {
					return err
				}

				for _, baselineMeasurement := range baselineCategory.BaselineMeasurementList.BaselineMeasurement {
					_, err := t.client.BaselineMeasurement.Create().
						SetBaselineMeasurementGroupID(baselineMeasurement.BaselineMeasurementGroupID).
						SetBaselineMeasurementValue(baselineMeasurement.BaselineMeasurementValue).
						SetBaselineMeasurementSpread(baselineMeasurement.BaselineMeasurementSpread).
						SetBaselineMeasurementLowerLimit(baselineMeasurement.BaselineMeasurementLowerLimit).
						SetBaselineMeasurementUpperLimit(baselineMeasurement.BaselineMeasurementUpperLimit).
						SetBaselineMeasurementComment(baselineMeasurement.BaselineMeasurementComment).
						SetParent(baselineCategoryParent).
						Save(t.ctx)
					if err != nil {
						return err
					}
				}
			}
		}

		for _, baselineMeasureDenom := range baselineMeasure.BaselineMeasureDenomList.BaselineMeasureDenom {
			baselineMeasureDenomParent, err := t.client.BaselineMeasureDenom.Create().
				SetParent(baselineMeasureParent).
				SetBaselineMeasureDenomUnits(baselineMeasureDenom.BaselineMeasureDenomUnits).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, baselineMeasureDenomCount := range baselineMeasureDenom.BaselineMeasureDenomCountList.BaselineMeasureDenomCount {
				_, err := t.client.BaselineMeasureDenomCount.Create().
					SetParent(baselineMeasureDenomParent).
					SetBaselineMeasureDenomCountGroupID(baselineMeasureDenomCount.BaselineMeasureDenomCountGroupID).
					SetBaselineMeasureDenomCountValue(baselineMeasureDenomCount.BaselineMeasureDenomCountValue).
					Save(t.ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (t *ResultsContext) insertOutcomeMeasureModule(resultsDefinition *models.ResultsDefinition, study Study, customOutcomeMeasures []OutcomeMeasureCustom) error {
	outcomeMeasureModule := study.ResultsSection.OutcomeMeasuresModule

	outcomeMeasureModuleParent, err := t.client.OutcomeMeasuresModule.Create().
		SetParent(resultsDefinition).
		Save(t.ctx)
	if err != nil {
		return err
	}

	for _, com := range customOutcomeMeasures {
		found := false
		for _, outcomeMeasure := range outcomeMeasureModule.OutcomeMeasureList.OutcomeMeasure {
			if outcomeMeasure.OutcomeMeasureTitle == com.OutcomeMeasureTitle {
				found = true
				break
			}
		}
		if found {
			continue
		}

		outcomeMeasureParent, err := t.client.OutcomeMeasure.Create().
			SetParent(outcomeMeasureModuleParent).
			SetOutcomeMeasureType(com.OutcomeMeasureType).
			SetOutcomeMeasureTitle(com.OutcomeMeasureTitle).
			SetOutcomeMeasureDescription(com.OutcomeMeasureDescription).
			SetOutcomeMeasurePopulationDescription(com.OutcomeMeasurePopulationDescription).
			SetOutcomeMeasureReportingStatus(com.OutcomeMeasureReportingStatus).
			SetOutcomeMeasureParamType(com.OutcomeMeasureParamType).
			SetOutcomeMeasureDispersionType(com.OutcomeMeasureDispersionType).
			SetOutcomeMeasureUnitOfMeasure(com.OutcomeMeasureUnitOfMeasure).
			SetOutcomeMeasureTimeFrame(com.OutcomeMeasureTimeFrame).
			SetOutcomeMeasureAnticipatedPostingDate(com.OutcomeMeasureAnticipatedPostingDate).
			SetOutcomeMeasureCalculatePct(com.OutcomeMeasureCalculatePct).
			SetOutcomeMeasureDenomUnitsSelected(com.OutcomeMeasureDenomUnitsSelected).
			SetOutcomeMeasureTypeUnitsAnalyzed(com.OutcomeMeasureTypeUnitsAnalyzed).
			Save(t.ctx)
		if err != nil {
			return err
		}

		outcomeOverviews, err := t.client.OutcomeOverview.Query().Where(outcomeoverview.Not(outcomeoverview.HasParent())).All(t.ctx)
		if err != nil {
			return err
		}

		for _, outcomeOverview := range outcomeOverviews {
			if outcomeOverview.OutcomeOverviewMeasureTitle == outcomeMeasureParent.OutcomeMeasureTitle {
				if _, err := outcomeOverview.Update().SetParent(outcomeMeasureParent).Save(t.ctx); err != nil {
					return err
				}
			}
		}
	}

	for _, outcomeMeasure := range outcomeMeasureModule.OutcomeMeasureList.OutcomeMeasure {
		outcomeMeasureParent, err := t.client.OutcomeMeasure.Create().
			SetParent(outcomeMeasureModuleParent).
			SetOutcomeMeasureType(outcomeMeasure.OutcomeMeasureType).
			SetOutcomeMeasureTitle(outcomeMeasure.OutcomeMeasureTitle).
			SetOutcomeMeasureDescription(outcomeMeasure.OutcomeMeasureDescription).
			SetOutcomeMeasurePopulationDescription(outcomeMeasure.OutcomeMeasurePopulationDescription).
			SetOutcomeMeasureReportingStatus(outcomeMeasure.OutcomeMeasureReportingStatus).
			SetOutcomeMeasureParamType(outcomeMeasure.OutcomeMeasureParamType).
			SetOutcomeMeasureDispersionType(outcomeMeasure.OutcomeMeasureDispersionType).
			SetOutcomeMeasureUnitOfMeasure(outcomeMeasure.OutcomeMeasureUnitOfMeasure).
			SetOutcomeMeasureTimeFrame(outcomeMeasure.OutcomeMeasureTimeFrame).
			SetOutcomeMeasureAnticipatedPostingDate(outcomeMeasure.OutcomeMeasureAnticipatedPostingDate).
			SetOutcomeMeasureCalculatePct(outcomeMeasure.OutcomeMeasureCalculatePct).
			SetOutcomeMeasureDenomUnitsSelected(outcomeMeasure.OutcomeMeasureDenomUnitsSelected).
			SetOutcomeMeasureTypeUnitsAnalyzed(outcomeMeasure.OutcomeMeasureTypeUnitsAnalyzed).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, outcomeAnalysis := range outcomeMeasure.OutcomeAnalysisList.OutcomeAnalysis {
			outcomeAnalysisParent, err := t.client.OutcomeAnalysis.Create().
				SetParent(outcomeMeasureParent).
				SetOutcomeAnalysisCiLowerLimit(outcomeAnalysis.OutcomeAnalysisCILowerLimit).
				SetOutcomeAnalysisCiLowerLimitComment(outcomeAnalysis.OutcomeAnalysisCILowerLimitComment).
				SetOutcomeAnalysisCiNumSides(outcomeAnalysis.OutcomeAnalysisCINumSides).
				SetOutcomeAnalysisCiPctValue(outcomeAnalysis.OutcomeAnalysisCIPctValue).
				SetOutcomeAnalysisCiUpperLimit(outcomeAnalysis.OutcomeAnalysisCIUpperLimit).
				SetOutcomeAnalysisCiUpperLimitComment(outcomeAnalysis.OutcomeAnalysisCIUpperLimitComment).
				SetOutcomeAnalysisDispersionType(outcomeAnalysis.OutcomeAnalysisDispersionType).
				SetOutcomeAnalysisDispersionValue(outcomeAnalysis.OutcomeAnalysisDispersionValue).
				SetOutcomeAnalysisEstimateComment(outcomeAnalysis.OutcomeAnalysisEstimateComment).
				SetOutcomeAnalysisGroupDescription(outcomeAnalysis.OutcomeAnalysisGroupDescription).
				SetOutcomeAnalysisNonInferiorityComment(outcomeAnalysis.OutcomeAnalysisNonInferiorityComment).
				SetOutcomeAnalysisNonInferiorityType(outcomeAnalysis.OutcomeAnalysisNonInferiorityType).
				SetOutcomeAnalysisOtherAnalysisDescription(outcomeAnalysis.OutcomeAnalysisDescription).
				SetOutcomeAnalysisParamType(outcomeAnalysis.OutcomeAnalysisParamType).
				SetOutcomeAnalysisParamValue(outcomeAnalysis.OutcomeAnalysisParamValue).
				SetOutcomeAnalysisPValue(outcomeAnalysis.OutcomeAnalysisPValue).
				SetOutcomeAnalysisPValueComment(outcomeAnalysis.OutcomeAnalysisPValueComment).
				SetOutcomeAnalysisStatisticalComment(outcomeAnalysis.OutcomeAnalysisStatisticalComment).
				SetOutcomeAnalysisStatisticalMethod(outcomeAnalysis.OutcomeAnalysisStatisticalMethod).
				SetOutcomeAnalysisTestedNonInferiority(outcomeAnalysis.OutcomeAnalysisTestedNonInferiority).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, outcomeAnalysisGroupID := range outcomeAnalysis.OutcomeAnalysisGroupIDList.OutcomeAnalysisGroupID {
				_, err := t.client.OutcomeAnalysisGroupID.Create().
					SetParent(outcomeAnalysisParent).
					SetOutcomeAnalysisGroupID(outcomeAnalysisGroupID).
					Save(t.ctx)
				if err != nil {
					return err
				}
			}
		}

		for _, outcomeClass := range outcomeMeasure.OutcomeClassList.OutcomeClass {
			outcomeClassParent, err := t.client.OutcomeClass.Create().
				SetParent(outcomeMeasureParent).
				SetOutcomeClassTitle(outcomeClass.OutcomeClassTitle).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, outcomeClassDenom := range outcomeClass.OutcomeClassDenomList.OutcomeClassDenom {
				outcomeClassDenomParent, err := t.client.OutcomeClassDenom.Create().
					SetParent(outcomeClassParent).
					SetOutcomeClassDenomUnits(outcomeClassDenom.OutcomeClassDenomUnits).
					Save(t.ctx)
				if err != nil {
					return err
				}

				for _, outcomeClassDenomCount := range outcomeClassDenom.OutcomeClassDenomCountList.OutcomeClassDenomCount {
					_, err := t.client.OutcomeClassDenomCount.Create().
						SetParent(outcomeClassDenomParent).
						SetOutcomeClassDenomCountGroupID(outcomeClassDenomCount.OutcomeClassDenomCountGroupID).
						SetOutcomeClassDenomCountValue(outcomeClassDenomCount.OutcomeClassDenomCountValue).
						Save(t.ctx)
					if err != nil {
						return err
					}
				}
			}

			for _, outcomeCategory := range outcomeClass.OutcomeCategoryList.OutcomeCategory {
				outcomeCategoryParent, err := t.client.OutcomeCategory.Create().
					SetParent(outcomeClassParent).
					SetOutcomeCategoryTitle(outcomeCategory.OutcomeCategoryTitle).
					Save(t.ctx)
				if err != nil {
					return err
				}

				for _, outcomeMeasurement := range outcomeCategory.OutcomeMeasurementList.OutcomeMeasurement {
					_, err := t.client.OutcomeMeasurement.Create().
						SetParent(outcomeCategoryParent).
						SetOutcomeMeasurementGroupID(outcomeMeasurement.OutcomeMeasurementGroupID).
						SetOutcomeMeasurementLowerLimit(outcomeMeasurement.OutcomeMeasurementLowerLimit).
						SetOutcomeMeasurementUpperLimit(outcomeMeasurement.OutcomeMeasurementUpperLimit).
						SetOutcomeMeasurementSpread(outcomeMeasurement.OutcomeMeasurementSpread).
						SetOutcomeMeasurementValue(outcomeMeasurement.OutcomeMeasurementValue).
						SetOutcomeMeasurementComment(outcomeMeasurement.OutcomeMeasurementComment).
						Save(t.ctx)
					if err != nil {
						return err
					}
				}
			}
		}

		for _, outcomeDenom := range outcomeMeasure.OutcomeDenomList.OutcomeDenom {
			outcomeDenomParent, err := t.client.OutcomeDenom.Create().
				SetParent(outcomeMeasureParent).
				SetOutcomeDenomUnits(outcomeDenom.OutcomeDenomUnits).
				Save(t.ctx)
			if err != nil {
				return err
			}

			for _, outcomeDenomCount := range outcomeDenom.OutcomeDenomCountList.OutcomeDenomCount {
				_, err := t.client.OutcomeDenomCount.Create().
					SetParent(outcomeDenomParent).
					SetOutcomeDenomCountGroupID(outcomeDenomCount.OutcomeDenomCountGroupID).
					SetOutcomeDenomCountValue(outcomeDenomCount.OutcomeDenomCountValue).
					Save(t.ctx)
				if err != nil {
					return err
				}
			}
		}

		for _, outcomeGroup := range outcomeMeasure.OutcomeGroupList.OutcomeGroup {
			_, err := t.client.OutcomeGroup.Create().
				SetParent(outcomeMeasureParent).
				SetOutcomeGroupID(outcomeGroup.OutcomeGroupID).
				SetOutcomeGroupTitle(outcomeGroup.OutcomeGroupTitle).
				SetOutcomeGroupDescription(outcomeGroup.OutcomeGroupDescription).
				Save(t.ctx)
			if err != nil {
				return err
			}
		}

		outcomeOverviews, err := t.client.OutcomeOverview.Query().Where(outcomeoverview.Not(outcomeoverview.HasParent())).All(t.ctx)
		if err != nil {
			return err
		}

		for _, outcomeOverview := range outcomeOverviews {
			if outcomeOverview.OutcomeOverviewMeasureTitle == outcomeMeasure.OutcomeMeasureTitle {
				if _, err := outcomeOverview.Update().SetParent(outcomeMeasureParent).Save(t.ctx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (t *ResultsContext) insertAdverseEventsModule(resultsDefinition *models.ResultsDefinition, study Study) error {
	adverseEventsModule := study.ResultsSection.AdverseEventsModule

	adverseEventsModuleParent, err := t.client.AdverseEventsModule.Create().
		SetParent(resultsDefinition).
		SetEventsDescription(adverseEventsModule.EventsDescription).
		SetEventsFrequencyThreshold(adverseEventsModule.EventsFrequencyThreshold).
		SetEventsTimeFrame(adverseEventsModule.EventsTimeFrame).
		Save(t.ctx)
	if err != nil {
		return err
	}

	for _, eventGroup := range adverseEventsModule.EventGroupList.EventGroup {
		_, err = t.client.EventGroup.Create().
			SetParent(adverseEventsModuleParent).
			SetEventGroupID(eventGroup.EventGroupID).
			SetEventGroupTitle(eventGroup.EventGroupTitle).
			SetEventGroupDescription(eventGroup.EventGroupDescription).
			SetEventGroupDeathsNumAffected(eventGroup.EventGroupDeathsNumAffected).
			SetEventGroupDeathsNumAtRisk(eventGroup.EventGroupDeathsNumAtRisk).
			SetEventGroupSeriousNumAffected(eventGroup.EventGroupSeriousNumAffected).
			SetEventGroupSeriousNumAtRisk(eventGroup.EventGroupSeriousNumAtRisk).
			SetEventGroupOtherNumAffected(eventGroup.EventGroupOtherNumAffected).
			SetEventGroupOtherNumAtRisk(eventGroup.EventGroupOtherNumAtRisk).
			Save(t.ctx)
		if err != nil {
			return err
		}
	}

	for _, seriousEvent := range adverseEventsModule.SeriousEventList.SeriousEvent {
		seriousEventParent, err := t.client.SeriousEvent.Create().
			SetParent(adverseEventsModuleParent).
			SetSeriousEventTerm(seriousEvent.SeriousEventTerm).
			SetSeriousEventOrganSystem(seriousEvent.SeriousEventOrganSystem).
			SetSeriousEventSourceVocabulary(seriousEvent.SeriousEventSourceVocabulary).
			SetSeriousEventAssessmentType(seriousEvent.SeriousEventAssessmentType).
			SetSeriousEventNotes(seriousEvent.SeriousEventNotes).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, seriousEventStats := range seriousEvent.SeriousEventStatsList.SeriousEventStats {
			_, err = t.client.SeriousEventStats.Create().
				SetSeriousEventStatsGroupID(seriousEventStats.SeriousEventStatsGroupID).
				SetSeriousEventStatsNumEvents(seriousEventStats.SeriousEventStatsNumEvents).
				SetSeriousEventStatsNumAffected(seriousEventStats.SeriousEventStatsNumAffected).
				SetSeriousEventStatsNumAtRisk(seriousEventStats.SeriousEventStatsNumAtRisk).
				SetParent(seriousEventParent).
				Save(t.ctx)
			if err != nil {
				return err
			}
		}
	}

	for _, otherEvent := range adverseEventsModule.OtherEventList.OtherEvent {
		otherEventParent, err := t.client.OtherEvent.Create().
			SetParent(adverseEventsModuleParent).
			SetOtherEventTerm(otherEvent.OtherEventTerm).
			SetOtherEventOrganSystem(otherEvent.OtherEventOrganSystem).
			SetOtherEventSourceVocabulary(otherEvent.OtherEventSourceVocabulary).
			SetOtherEventAssessmentType(otherEvent.OtherEventAssessmentType).
			SetOtherEventNotes(otherEvent.OtherEventNotes).
			Save(t.ctx)
		if err != nil {
			return err
		}

		for _, otherEventStats := range otherEvent.OtherEventStatsList.OtherEventStats {
			_, err = t.client.OtherEventStats.Create().
				SetOtherEventStatsGroupID(otherEventStats.OtherEventStatsGroupID).
				SetOtherEventStatsNumEvents(otherEventStats.OtherEventStatsNumEvents).
				SetOtherEventStatsNumAffected(otherEventStats.OtherEventStatsNumAffected).
				SetOtherEventStatsNumAtRisk(otherEventStats.OtherEventStatsNumAtRisk).
				SetParent(otherEventParent).
				Save(t.ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *ResultsContext) insertMoreInfoModule(resultsDefinition *models.ResultsDefinition, study Study) error {
	moreInfoModule := study.ResultsSection.MoreInfoModule

	moreInfoModuleParent, err := t.client.MoreInfoModule.Create().
		SetParent(resultsDefinition).
		SetLimitationsAndCaveatsDescription(moreInfoModule.LimitationsAndCaveatsDescription).
		Save(t.ctx)
	if err != nil {
		return err
	}

	_, err = t.client.PointOfContact.Create().
		SetParent(moreInfoModuleParent).
		SetPointOfContactTitle(moreInfoModule.PointOfContact.PointOfContactTitle).
		SetPointOfContactOrganization(moreInfoModule.PointOfContact.PointOfContactOrganization).
		SetPointOfContactEmail(moreInfoModule.PointOfContact.PointOfContactEmail).
		SetPointOfContactPhone(moreInfoModule.PointOfContact.PointOfContactPhone).
		SetPointOfContactPhoneExt(moreInfoModule.PointOfContact.PointOfContactPhoneExt).
		Save(t.ctx)
	if err != nil {
		return err
	}

	_, err = t.client.CertainAgreement.Create().
		SetParent(moreInfoModuleParent).
		SetAgreementPiSponsorEmployee(moreInfoModule.CertainAgreement.AgreementPISponsorEmployee).
		SetAgreementRestrictionType(moreInfoModule.CertainAgreement.AgreementRestrictionType).
		SetAgreementRestrictiveAgreement(moreInfoModule.CertainAgreement.AgreementRestrictiveAgreement).
		SetAgreementOtherDetails(moreInfoModule.CertainAgreement.AgreementOtherDetails).
		Save(t.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *ResultsContext) InsertOutcomeOverviews(overviews []OutcomeOverview) error {
	for _, overview := range overviews {
		if _, err := t.client.OutcomeOverview.Create().
		SetOutcomeOverviewMeasureTitle(overview.OutcomeMeasureTitle).
		SetOutcomeOverviewTitle(overview.OutcomeOverviewTitle).
		SetOutcomeOverviewParticipants(overview.OutcomeOverviewParticipants).
		SetOutcomeOverviewAssay(overview.OutcomeOverviewAssay).
		SetOutcomeOverviewDescription(overview.OutcomeOverviewDescription).
		SetOutcomeOverviewDoseNumber(overview.OutcomeOverviewDoseNumber).
		SetOutcomeOverviewGroupID(overview.OutcomeOverviewGroupID).
		SetOutcomeOverviewRatio(overview.OutcomeOverviewRatio).
		SetOutcomeOverviewID(overview.OutcomeOverviewID).
		SetOutcomeOverviewSerotype(overview.OutcomeOverviewSerotype).
		SetOutcomeOverviewTimeFrame(overview.OutcomeOverviewTimeFrame).
		SetOutcomeOverviewValue(overview.OutcomeOverviewValue).
		SetOutcomeOverviewLowerLimit(overview.OutcomeOverviewLowerLimit).
		SetOutcomeOverviewTimeFrameWeeks(overview.OutcomeOverviewTimeFrameWeeks).
		SetOutcomeOverviewVaccine(overview.OutcomeOverviewVaccine).
		SetOutcomeOverviewImmunocompromisedPopulation(overview.OutcomeOverviewImmunocompromisedPopulation).
		SetOutcomeOverviewManufacturer(overview.OutcomeOverviewManufacturer).
		SetOutcomeOverviewUpperLimit(overview.OutcomeOverviewUpperLimit).
		SetOutcomeOverviewConfidenceInterval(overview.OutcomeOverviewConfidenceInterval).
		SetOutcomeOverviewDoseDescription(overview.OutcomeOverviewDoseDescription).
		SetOutcomeOverviewSchedule(overview.OutcomeOverviewSchedule).
		SetOutcomeOverviewPercentResponders(overview.OutcomeOverviewPercentResponders).
		SetOutcomeOverviewUseCaseCode(overview.OutcomeOverviewUseCase).
			Save(t.ctx); err != nil {
			return err
		}
	}
	return nil
}

func getDate(date string) (time.Time, error) {
	retDate, err := time.Parse(timeLayout, date)
	if err != nil {
		retDate, err = time.Parse(timeLayoutAlt, date)
		if err != nil {
			retDate, err = time.Parse(timeLayoutAlt2, date)
			if err != nil {
				retDate = time.Date(1900, 1, 1, 0, 0,0, 0, time.Local)
			}
		}
	}

	return retDate, nil
}

func matchRegex(reg *regexp.Regexp, items...string) bool {
	for _, item := range items {
		if reg.MatchString(item) {
			return true
		}
	}
	return false
}

func getSerotype(haystack string) string {
	for _, needle := range serotypes {
		if findWord(needle, haystack) {
			return needle
		}
	}
	return haystack
}

func findWord(needle string, haystack string) bool {
	return regexp.MustCompile(fmt.Sprintf(`(?mi)(\w%s|\b%s)`,needle,needle)).Match([]byte(haystack))
}
