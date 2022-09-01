package clinicaltrials

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const tagName = "csv"

type SearchOptions struct {
	UseCase                      string   `json:"use_case"`
	IncomeGroups                 []string `json:"income_groups"`
	ParentArraySeparator         string   `json:"parent_array_separator"`
	ChildArraySeparator          string   `json:"child_array_separator"`
	Vaccine                      []string `json:"vaccine"`
	Phase                        []string `json:"phase"`
	STDAgeList                   []string `json:"std_age_list"`
	CountryCodes                 []string `json:"country_codes"`
	ContinentCodes               []string `json:"continent_codes"`
	ResponsibleParty             []string `json:"responsible_party"`
	EthnicityRace                []string `json:"ethnicity_race"`
	ImmunocompromisedPopulations []string `json:"immunocompromised_populations"`
}

type FieldOptions struct {
	LocationFacility                           bool `json:"location_facility"`
	LocationCity                               bool `json:"location_city"`
	LocationCountry                            bool `json:"location_country"`
	LocationCountryCode                        bool `json:"location_country_code"`
	LocationContinentName                      bool `json:"location_continent"`
	LocationContinentCode                      bool `json:"location_continent_code"`
	EligibilityCriteria                        bool `json:"study_eligibility_criteria"`
	HealthyVolunteers                          bool `json:"study_eligibility_healthy_volunteers"`
	Gender                                     bool `json:"study_eligibility_gender"`
	MinimumAge                                 bool `json:"study_eligibility_minimum_age"`
	MaximumAge                                 bool `json:"study_eligibility_maximum_age"`
	StandardAgeList                            bool `json:"study_eligibility_standard_age_list"`
	Ethnicity                                  bool `json:"study_eligibility_ethnicity"`
	ClinicalTrialStudyName                     bool `json:"clinical_trial_study_name"`
	ClinicalTrialStudyID                       bool `json:"clinical_trial_study_id"`
	Sponsor                                    bool `json:"clinical_trial_sponsor"`
	ResponsibleParty                           bool `json:"clinical_trial_responsible_party"`
	StudyType                                  bool `json:"clinical_trial_study_type"`
	Phase                                      bool `json:"clinical_trial_phase"`
	ActualEnrollment                           bool `json:"clinical_trial_actual_enrollment"`
	Allocation                                 bool `json:"clinical_trial_allocation"`
	InterventionModel                          bool `json:"clinical_trial_intervention_model"`
	Masking                                    bool `json:"clinical_trial_masking"`
	PrimaryPurpose                             bool `json:"clinical_trial_primary_purpose"`
	OfficialTitle                              bool `json:"clinical_trial_official_title"`
	ActualStudyStartDate                       bool `json:"clinical_trial_actual_study_start_date"`
	ActualPrimaryCompletionDate                bool `json:"clinical_trial_actual_primary_completion_date"`
	ActualStudyCompletionDate                  bool `json:"clinical_trial_actual_study_completion_date"`
	OutcomeMeasureType                         bool `json:"outcome_measure_type"`
	OutcomeMeasureTitle                        bool `json:"outcome_measure_title"`
	OutcomeMeasureDescription                  bool `json:"outcome_measure_description"`
	OutcomeMeasurePopulationDescription        bool `json:"outcome_measure_population_description"`
	OutcomeMeasureReportingStatus              bool `json:"outcome_measure_reporting_status"`
	OutcomeMeasureAnticipatedPostingDate       bool `json:"outcome_measure_anticipated_posting_date"`
	OutcomeMeasureParamType                    bool `json:"outcome_measure_param_type"`
	OutcomeMeasureDispersionType               bool `json:"outcome_measure_dispersion_type"`
	OutcomeMeasureUnitOfMeasure                bool `json:"outcome_measure_unit_of_measure"`
	OutcomeMeasureCalculatePCT                 bool `json:"outcome_measure_calculate_pct"`
	OutcomeMeasureTimeFrame                    bool `json:"outcome_measure_time_frame"`
	OutcomeMeasureTypeUnitsAnalyzed            bool `json:"outcome_measure_type_units_analyzed"`
	OutcomeMeasureDenomUnitsSelected           bool `json:"outcome_measure_denom_units_selected"`
	OutcomeOverviewTitle                       bool `json:"outcome_overview_title"`
	OutcomeOverviewID                          bool `json:"outcome_overview_id"`
	OutcomeOverviewDescription                 bool `json:"outcome_overview_description"`
	OutcomeOverviewTimeFrame                   bool `json:"outcome_overview_time_frame"`
	OutcomeOverviewAssay                       bool `json:"outcome_overview_assay"`
	OutcomeOverviewDoseNumber                  bool `json:"outcome_overview_dose_number"`
	OutcomeOverviewParticipants                bool `json:"outcome_overview_participants"`
	OutcomeOverviewSerotype                    bool `json:"outcome_overview_serotype"`
	OutcomeOverviewValue                       bool `json:"outcome_overview_value"`
	OutcomeOverviewUpperLimit                  bool `json:"outcome_overview_upper_limit"`
	OutcomeOverviewLowerLimit                  bool `json:"outcome_overview_lower_limit"`
	OutcomeOverviewRatio                       bool `json:"outcome_overview_ratio"`
	OutcomeOverviewVaccine                     bool `json:"outcome_overview_vaccine"`
	OutcomeOverviewImmunocompromisedPopulation bool `json:"outcome_overview_immunocompromised_population"`
	OutcomeOverviewManufacturer                bool `json:"outcome_overview_manufacturer"`
	OutcomeOverviewDoseDescription             bool `json:"outcome_overview_dose_description"`
	OutcomeOverviewSchedule                    bool `json:"outcome_overview_schedule"`
	OutcomeOverviewTimeFrameWeeks              bool `json:"outcome_overview_time_frame_weeks"`
	OutcomeOverviewConfidenceInterval          bool `json:"outcome_overview_confidence_interval"`
	OutcomeOverviewPercentResponders           bool `json:"outcome_overview_percent_responders"`
}

type CSVData struct {
	ClinicalTrialStudyName      string    `json:"clinical_trial_study_name,omitempty" csv:"clinical_trial_study_name"`
	ClinicalTrialStudyID        string    `json:"clinical_trial_study_id,omitempty" csv:"clinical_trial_study_id"`
	Sponsor                     string    `json:"clinical_trial_sponsor,omitempty" csv:"clinical_trial_sponsor"`
	ResponsibleParty            string    `json:"clinical_trial_responsible_party,omitempty" csv:"clinical_trial_responsible_party"`
	StudyType                   string    `json:"clinical_trial_study_type,omitempty" csv:"clinical_trial_study_type"`
	Phase                       string    `json:"clinical_trial_phase,omitempty" csv:"clinical_trial_phase"`
	ActualEnrollment            string    `json:"clinical_trial_actual_enrollment,omitempty" csv:"clinical_trial_actual_enrollment"`
	Allocation                  string    `json:"clinical_trial_allocation,omitempty" csv:"clinical_trial_allocation"`
	InterventionModel           string    `json:"clinical_trial_intervention_model,omitempty" csv:"clinical_trial_intervention_model"`
	Masking                     string    `json:"clinical_trial_masking,omitempty" csv:"clinical_trial_masking"`
	PrimaryPurpose              string    `json:"clinical_trial_primary_purpose,omitempty" csv:"clinical_trial_primary_purpose"`
	OfficialTitle               string    `json:"clinical_trial_official_title,omitempty" csv:"clinical_trial_official_title"`
	ActualStudyStartDate        time.Time `json:"clinical_trial_actual_study_start_date,omitempty" csv:"clinical_trial_actual_study_start_date"`
	ActualPrimaryCompletionDate time.Time `json:"clinical_trial_actual_primary_completion_date,omitempty" csv:"clinical_trial_actual_primary_completion_date"`
	ActualStudyCompletionDate   time.Time `json:"clinical_trial_actual_study_completion_date,omitempty" csv:"clinical_trial_actual_study_completion_date"`

	LocationFacility      string `json:"location_facility,omitempty" csv:"location_facility"`
	LocationCity          string `json:"location_city,omitempty" csv:"location_city"`
	LocationCountry       string `json:"location_country,omitempty" csv:"location_country"`
	LocationCountryCode   string `json:"location_country_code,omitempty" csv:"location_country_code"`
	LocationContinentName string `json:"location_continent,omitempty" csv:"location_continent"`
	LocationContinentCode string `json:"location_continent_code,omitempty" csv:"location_continent_code"`

	EligibilityCriteria string `json:"study_eligibility_criteria,omitempty" csv:"study_eligibility_criteria"`
	HealthyVolunteers   string `json:"study_eligibility_healthy_volunteers,omitempty" csv:"study_eligibility_healthy_volunteers"`
	Gender              string `json:"study_eligibility_gender,omitempty" csv:"study_eligibility_gender"`
	MinimumAge          string `json:"study_eligibility_minimum_age,omitempty" csv:"study_eligibility_minimum_age"`
	MaximumAge          string `json:"study_eligibility_maximum_age,omitempty" csv:"study_eligibility_maximum_age"`
	StandardAgeList     string `json:"study_eligibility_standard_age_list,omitempty" csv:"study_eligibility_standard_age_list"`
	Ethnicity           string `json:"study_eligibility_ethnicity,omitempty" csv:"study_eligibility_ethnicity"`

	OutcomeMeasureType                   string `json:"outcome_measure_type,omitempty" csv:"outcome_measure_type"`
	OutcomeMeasureTitle                  string `json:"outcome_measure_title,omitempty" csv:"outcome_measure_title"`
	OutcomeMeasureDescription            string `json:"outcome_measure_description,omitempty" csv:"outcome_measure_description"`
	OutcomeMeasurePopulationDescription  string `json:"outcome_measure_population_description,omitempty" csv:"outcome_measure_population_description"`
	OutcomeMeasureReportingStatus        string `json:"outcome_measure_reporting_status,omitempty" csv:"outcome_measure_reporting_status"`
	OutcomeMeasureAnticipatedPostingDate string `json:"outcome_measure_anticipated_posting_date,omitempty" csv:"outcome_measure_anticipated_posting_date"`
	OutcomeMeasureParamType              string `json:"outcome_measure_param_type,omitempty" csv:"outcome_measure_param_type"`
	OutcomeMeasureDispersionType         string `json:"outcome_measure_dispersion_type,omitempty" csv:"outcome_measure_dispersion_type"`
	OutcomeMeasureUnitOfMeasure          string `json:"outcome_measure_unit_of_measure,omitempty" csv:"outcome_measure_unit_of_measure"`
	OutcomeMeasureCalculatePCT           string `json:"outcome_measure_calculate_pct,omitempty" csv:"outcome_measure_calculate_pct"`
	OutcomeMeasureTimeFrame              string `json:"outcome_measure_time_frame,omitempty" csv:"outcome_measure_time_frame"`
	OutcomeMeasureTypeUnitsAnalyzed      string `json:"outcome_measure_type_units_analyzed,omitempty" csv:"outcome_measure_type_units_analyzed"`
	OutcomeMeasureDenomUnitsSelected     string `json:"outcome_measure_denom_units_selected,omitempty" csv:"outcome_measure_denom_units_selected"`

	OutcomeOverviewTitle                       string  `json:"outcome_overview_title,omitempty" csv:"outcome_overview_title"`
	OutcomeOverviewID                          string  `json:"outcome_overview_id,omitempty" csv:"outcome_overview_id"`
	OutcomeOverviewDescription                 string  `json:"outcome_overview_description,omitempty" csv:"outcome_overview_description"`
	OutcomeOverviewTimeFrame                   string  `json:"outcome_overview_time_frame,omitempty" csv:"outcome_overview_time_frame"`
	OutcomeOverviewAssay                       string  `json:"outcome_overview_assay,omitempty" csv:"outcome_overview_assay"`
	OutcomeOverviewDoseNumber                  int64   `json:"outcome_overview_dose_number,omitempty" csv:"outcome_overview_dose_number"`
	OutcomeOverviewParticipants                string  `json:"outcome_overview_participants,omitempty" csv:"outcome_overview_participants"`
	OutcomeOverviewSerotype                    string  `json:"outcome_overview_serotype,omitempty" csv:"outcome_overview_serotype"`
	OutcomeOverviewValue                       float64 `json:"outcome_overview_value,omitempty" csv:"outcome_overview_value"`
	OutcomeOverviewUpperLimit                  string  `json:"outcome_overview_upper_limit,omitempty" csv:"outcome_overview_upper_limit"`
	OutcomeOverviewLowerLimit                  string  `json:"outcome_overview_lower_limit,omitempty" csv:"outcome_overview_lower_limit"`
	OutcomeOverviewRatio                       string  `json:"outcome_overview_ratio,omitempty" csv:"outcome_overview_ratio"`
	OutcomeOverviewVaccine                     string  `json:"outcome_overview_vaccine,omitempty" csv:"outcome_overview_vaccine"`
	OutcomeOverviewImmunocompromisedPopulation string  `json:"outcome_overview_immunocompromised_population,omitempty" csv:"outcome_overview_immunocompromised_population"`
	OutcomeOverviewManufacturer                string  `json:"outcome_overview_manufacturer,omitempty" csv:"outcome_overview_manufacturer"`
	OutcomeOverviewDoseDescription             string  `json:"outcome_overview_dose_description,omitempty" csv:"outcome_overview_dose_description"`
	OutcomeOverviewSchedule                    string  `json:"outcome_overview_schedule,omitempty" csv:"outcome_overview_schedule"`
	OutcomeOverviewTimeFrameWeeks              int64   `json:"outcome_overview_time_frame_weeks,omitempty" csv:"outcome_overview_time_frame_weeks"`
	OutcomeOverviewConfidenceInterval          string  `json:"outcome_overview_confidence_interval,omitempty" csv:"outcome_overview_confidence_interval"`
	OutcomeOverviewPercentResponders           float64 `json:"outcome_overview_percent_responders,omitempty" csv:"outcome_overview_percent_responders"`
}

func (t *ResultsContext) GetExport(searchFilter SearchOptions, fields FieldOptions) (*bytes.Buffer, error) {
	data, err := t.GetData(searchFilter, fields)
	if err != nil {
		return nil, err
	}

	return dataToCSV(data, fields)
}

func (t *ResultsContext) GetData(searchFilter SearchOptions, fields FieldOptions) ([]CSVData, error) {
	if searchFilter.UseCase == "" {
		return nil, fmt.Errorf("use case value not set")
	}

	includeAnyEligibilityFields := fields.EligibilityCriteria || fields.HealthyVolunteers || fields.Gender || fields.MinimumAge || fields.MaximumAge || fields.StandardAgeList || fields.Ethnicity
	includeAnyLocationFields := fields.LocationFacility || fields.LocationCity || fields.LocationCountry || fields.LocationCountryCode || fields.LocationContinentName || fields.LocationContinentCode

	if searchFilter.ParentArraySeparator == "" {
		searchFilter.ParentArraySeparator = ","
		searchFilter.ChildArraySeparator = " "
	}

	trialWhere := []predicate.ClinicalTrial{}
	if len(searchFilter.IncomeGroups) > 0 {
		trialWhere = append(trialWhere, clinicaltrial.HasStudyLocationsWith(studylocation.LocationCountryCodeIn(incomeGroupToCountryCode(searchFilter.IncomeGroups)...)))
	}

	if len(searchFilter.Phase) > 0 {
		phaseList := []predicate.ClinicalTrial{}
		for _, phase := range searchFilter.Phase {
			phaseList = append(phaseList, clinicaltrial.PhaseContains(phase))
		}
		trialWhere = append(trialWhere, clinicaltrial.Or(phaseList...))
	}

	if len(searchFilter.ResponsibleParty) > 0 {
		trialWhere = append(trialWhere, clinicaltrial.ResponsiblePartyIn(searchFilter.ResponsibleParty...))
	}

	if len(searchFilter.STDAgeList) > 0 {
		studyEligibilityWhere := []predicate.StudyEligibility{}
		for _, age := range searchFilter.STDAgeList {
			studyEligibilityWhere = append(studyEligibilityWhere, studyeligibility.StdAgeListContains(fmt.Sprintf("\"%s\"",age)))
		}
		trialWhere = append(trialWhere, clinicaltrial.HasStudyEligibilityWith(studyeligibility.Or(studyEligibilityWhere...)))
	}

	outcomeOverviewWhere := []predicate.OutcomeOverview{}
	if len(searchFilter.ImmunocompromisedPopulations) > 0 {
		ip := []predicate.OutcomeOverview{}
		for _, icp := range searchFilter.ImmunocompromisedPopulations {
			ip = append(ip, outcomeoverview.OutcomeOverviewImmunocompromisedPopulationContains(icp))
		}
		outcomeOverviewWhere = append(outcomeOverviewWhere, outcomeoverview.Or(ip...))
	}

	outcomeOverviewWhere = append(outcomeOverviewWhere, outcomeoverview.OutcomeOverviewUseCaseCode(searchFilter.UseCase))

	if len(searchFilter.Vaccine) > 0 {
		outcomeOverviewWhere = append(outcomeOverviewWhere, outcomeoverview.OutcomeOverviewVaccineIn(searchFilter.Vaccine...))
	}

	trialWhere = append(trialWhere, clinicaltrial.HasResultsDefinitionWith(resultsdefinition.HasOutcomeMeasuresModuleWith(outcomemeasuresmodule.HasOutcomeMeasureListWith(outcomemeasure.HasOutcomeOverviewListWith(outcomeoverview.And(outcomeOverviewWhere...))))))

	query := t.client.ClinicalTrial.Query().WithStudyLocations().WithStudyEligibility().WithResultsDefinition(
		func(q *models.ResultsDefinitionQuery) {
			q.WithOutcomeMeasuresModule(
				func(q *models.OutcomeMeasuresModuleQuery) {
					q.WithOutcomeMeasureList(
						func(q *models.OutcomeMeasureQuery) {
							q.WithOutcomeOverviewList()
						},
					)
				},
			)
		},
	).Where(clinicaltrial.And(trialWhere...))

	res, err := query.All(t.ctx)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	data := []CSVData{}

	for _, trial := range res {
		eligibilityCriteria := ""
		healthyVolunteers := ""
		gender := ""
		minimumAge := ""
		maximumAge := ""
		standardAgeList := ""
		ethnicity := ""

		if includeAnyEligibilityFields {
			eligibility, err := trial.Edges.StudyEligibilityOrErr()
			if err != nil {
				return nil, err
			}

			if len(eligibility) > 0 {
				eligibilityCriteria = eligibility[0].EligibilityCriteria
				healthyVolunteers = eligibility[0].HealthyVolunteers
				gender = eligibility[0].Gender
				minimumAge = eligibility[0].MinimumAge
				maximumAge = eligibility[0].MaximumAge
				standardAgeList = eligibility[0].StdAgeList
				ethnicity = eligibility[0].Ethnicity
			}
		}

		facilities := []string{}
		cities := []string{}
		countries := []string{}
		countryCodes := []string{}
		continents := []string{}
		continentCodes := []string{}

		if includeAnyLocationFields {
			locations, err := trial.Edges.StudyLocationsOrErr()
			if err != nil {
				return nil, err
			}

			for _, location := range locations {
				facilities = append(facilities, location.LocationFacility)
				cities = append(cities, location.LocationCity)
				countries = append(countries, location.LocationCountry)
				countryCodes = append(countryCodes, location.LocationCountryCode)
				continents = append(continents, strings.Join(strings.Split(location.LocationContinentName, ","), searchFilter.ChildArraySeparator))
				continentCodes = append(continentCodes, strings.Join(strings.Split(location.LocationContinentCode, ","), searchFilter.ChildArraySeparator))
			}
		}

		rd, err := trial.Edges.ResultsDefinitionOrErr()
		if err != nil {
			return nil, err
		}

		omm, err := rd.Edges.OutcomeMeasuresModuleOrErr()
		if err != nil {
			return nil, err
		}

		oml, err := omm.Edges.OutcomeMeasureListOrErr()
		if err != nil {
			return nil, err
		}

		for _, om := range oml {
			ovl, err := om.Edges.OutcomeOverviewListOrErr()
			if err != nil {
				return nil, err
			}

			for _, ov := range ovl {
				if len(searchFilter.ImmunocompromisedPopulations) > 0 {
					foundICP := false
					for _, icp := range searchFilter.ImmunocompromisedPopulations {
						if strings.Contains(ov.OutcomeOverviewImmunocompromisedPopulation, icp) {
							foundICP = true
							break
						}
						if icp == "" && ov.OutcomeOverviewImmunocompromisedPopulation == "" {
							foundICP = true
						}
					}
					if !foundICP {
						continue
					}
				}

				data = append(data, CSVData{
					EligibilityCriteria: eligibilityCriteria,
					HealthyVolunteers:   healthyVolunteers,
					Gender:              gender,
					MinimumAge:          minimumAge,
					MaximumAge:          maximumAge,
					StandardAgeList:     standardAgeList,
					Ethnicity:           ethnicity,

					ClinicalTrialStudyName:      trial.StudyName,
					ClinicalTrialStudyID:      	 trial.StudyID,
					Sponsor:                     trial.Sponsor,
					ResponsibleParty:            trial.ResponsibleParty,
					StudyType:                   trial.StudyType,
					Phase:                       trial.Phase,
					ActualEnrollment:            trial.ActualEnrollment,
					Allocation:                  trial.Allocation,
					InterventionModel:           trial.InterventionModel,
					Masking:                     trial.Masking,
					PrimaryPurpose:              trial.PrimaryPurpose,
					OfficialTitle:               trial.OfficialTitle,
					ActualStudyStartDate:        trial.ActualStudyStartDate,
					ActualPrimaryCompletionDate: trial.ActualPrimaryCompletionDate,
					ActualStudyCompletionDate:   trial.ActualStudyCompletionDate,

					OutcomeMeasureType:                   om.OutcomeMeasureType,
					OutcomeMeasureDescription:            om.OutcomeMeasureDescription,
					OutcomeMeasurePopulationDescription:  om.OutcomeMeasurePopulationDescription,
					OutcomeMeasureReportingStatus:        om.OutcomeMeasureReportingStatus,
					OutcomeMeasureAnticipatedPostingDate: om.OutcomeMeasureAnticipatedPostingDate,
					OutcomeMeasureParamType:              om.OutcomeMeasureParamType,
					OutcomeMeasureDispersionType:         om.OutcomeMeasureDispersionType,
					OutcomeMeasureUnitOfMeasure:          om.OutcomeMeasureUnitOfMeasure,
					OutcomeMeasureCalculatePCT:           om.OutcomeMeasureCalculatePct,
					OutcomeMeasureTimeFrame:              om.OutcomeMeasureTimeFrame,
					OutcomeMeasureTypeUnitsAnalyzed:      om.OutcomeMeasureTypeUnitsAnalyzed,
					OutcomeMeasureDenomUnitsSelected:     om.OutcomeMeasureDenomUnitsSelected,

					LocationFacility:      strings.Join(facilities, searchFilter.ParentArraySeparator),
					LocationCity:          strings.Join(cities, searchFilter.ParentArraySeparator),
					LocationCountry:       strings.Join(countries, searchFilter.ParentArraySeparator),
					LocationCountryCode:   strings.Join(countryCodes, searchFilter.ParentArraySeparator),
					LocationContinentName: strings.Join(continents, searchFilter.ParentArraySeparator),
					LocationContinentCode: strings.Join(continentCodes, searchFilter.ParentArraySeparator),

					OutcomeOverviewTitle:                       ov.OutcomeOverviewTitle,
					OutcomeOverviewID:                          ov.OutcomeOverviewID,
					OutcomeOverviewDescription:                 ov.OutcomeOverviewDescription,
					OutcomeOverviewTimeFrame:                   ov.OutcomeOverviewTimeFrame,
					OutcomeOverviewAssay:                       ov.OutcomeOverviewAssay,
					OutcomeOverviewDoseNumber:                  ov.OutcomeOverviewDoseNumber,
					OutcomeMeasureTitle:                        ov.OutcomeOverviewMeasureTitle,
					OutcomeOverviewParticipants:                ov.OutcomeOverviewParticipants,
					OutcomeOverviewSerotype:                    ov.OutcomeOverviewSerotype,
					OutcomeOverviewValue:                       ov.OutcomeOverviewValue,
					OutcomeOverviewUpperLimit:                  ov.OutcomeOverviewUpperLimit,
					OutcomeOverviewLowerLimit:                  ov.OutcomeOverviewLowerLimit,
					OutcomeOverviewRatio:                       ov.OutcomeOverviewRatio,
					OutcomeOverviewVaccine:                     ov.OutcomeOverviewVaccine,
					OutcomeOverviewImmunocompromisedPopulation: ov.OutcomeOverviewImmunocompromisedPopulation,
					OutcomeOverviewManufacturer:                ov.OutcomeOverviewManufacturer,
					OutcomeOverviewDoseDescription:             ov.OutcomeOverviewDoseDescription,
					OutcomeOverviewSchedule:                	ov.OutcomeOverviewSchedule,
					OutcomeOverviewTimeFrameWeeks:              ov.OutcomeOverviewTimeFrameWeeks,
					OutcomeOverviewConfidenceInterval:          ov.OutcomeOverviewConfidenceInterval,
					OutcomeOverviewPercentResponders:           ov.OutcomeOverviewPercentResponders,
				})
			}
		}
	}

	return data, nil
}

func dataToCSV(input []CSVData, fieldOptions FieldOptions) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)

	data := structToStrings(input, fieldOptions)
	if data == nil {
		return nil, fmt.Errorf("err: can't make csv with no data")
	}

	w.WriteAll(data)

	return b, w.Error()
}

func structToStrings(input []CSVData, fieldOptions FieldOptions) [][]string {
	if len(input) == 0 {
		return nil
	}

	output := [][]string{}

	output = append(output, getHeaders(input[0], fieldOptions))

	for _, row := range input {
		output = append(output, getValues(row, fieldOptions))
	}

	return output
}

func getHeaders(input interface{}, fieldOptions FieldOptions) []string {
	if input == nil {
		return nil
	}

	output := []string{}

	t := reflect.TypeOf(input)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		if tag == "" {
			fmt.Printf("couldn't find tag for %s", t.Field(i).Name)
			continue
		}

		fieldName := getFieldNameFromTag(tag)
		if fieldName == "" {
			fmt.Printf("couldn't find field name for %s", tag)
			continue
		}

		if !reflect.
			Indirect(reflect.ValueOf(fieldOptions)).
			FieldByName(fieldName).
			Bool() {
			continue
		}

		output = append(output, tag)
	}

	return output
}

func getValues(input interface{}, fieldOptions FieldOptions) []string {
	if input == nil {
		return nil
	}

	values := reflect.ValueOf(input)
	t := reflect.TypeOf(input)
	rowVals := []string{}
	for i := 0; i < values.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		if tag == "" {
			fmt.Printf("couldn't find tag for %s", t.Field(i).Name)
			fmt.Println()
			continue
		}

		fieldName := getFieldNameFromTag(tag)
		if fieldName == "" {
			fmt.Printf("couldn't find field name for %s", tag)
			fmt.Println()
			continue
		}

		if !reflect.
			Indirect(reflect.ValueOf(fieldOptions)).
			FieldByName(fieldName).
			Bool() {
			continue
		}

		result := ""

		switch values.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = strconv.FormatInt(values.Field(i).Int(), 10)
		case reflect.Float32, reflect.Float64:
			result = strconv.FormatFloat(values.Field(i).Float(), 'g', 'g', 64)
		default:
			result = values.Field(i).String()
		}

		rowVals = append(rowVals, result)
	}

	return rowVals
}

func getFieldNameFromTag(s string) string {
	output := ""
	switch s {
	case "location_facility":
		output = "LocationFacility"
		break
	case "location_city":
		output = "LocationCity"
		break
	case "location_country":
		output = "LocationCountry"
		break
	case "location_country_code":
		output = "LocationCountryCode"
		break
	case "location_continent":
		output = "LocationContinentName"
		break
	case "location_continent_code":
		output = "LocationContinentCode"
		break
	case "study_eligibility_criteria":
		output = "EligibilityCriteria"
		break
	case "study_eligibility_healthy_volunteers":
		output = "HealthyVolunteers"
		break
	case "study_eligibility_gender":
		output = "Gender"
		break
	case "study_eligibility_minimum_age":
		output = "MinimumAge"
		break
	case "study_eligibility_maximum_age":
		output = "MaximumAge"
		break
	case "study_eligibility_standard_age_list":
		output = "StandardAgeList"
		break
	case "study_eligibility_ethnicity":
		output = "Ethnicity"
		break
	case "clinical_trial_study_name":
		output = "ClinicalTrialStudyName"
		break
	case "clinical_trial_study_id":
		output = "ClinicalTrialStudyID"
		break
	case "clinical_trial_sponsor":
		output = "Sponsor"
		break
	case "clinical_trial_responsible_party":
		output = "ResponsibleParty"
		break
	case "clinical_trial_study_type":
		output = "StudyType"
		break
	case "clinical_trial_phase":
		output = "Phase"
		break
	case "clinical_trial_actual_enrollment":
		output = "ActualEnrollment"
		break
	case "clinical_trial_allocation":
		output = "Allocation"
		break
	case "clinical_trial_intervention_model":
		output = "InterventionModel"
		break
	case "clinical_trial_masking":
		output = "Masking"
		break
	case "clinical_trial_primary_purpose":
		output = "PrimaryPurpose"
		break
	case "clinical_trial_official_title":
		output = "OfficialTitle"
		break
	case "clinical_trial_actual_study_start_date":
		output = "ActualStudyStartDate"
		break
	case "clinical_trial_actual_primary_completion_date":
		output = "ActualPrimaryCompletionDate"
		break
	case "clinical_trial_actual_study_completion_date":
		output = "ActualStudyCompletionDate"
		break
	case "outcome_measure_type":
		output = "OutcomeMeasureType"
		break
	case "outcome_measure_title":
		output = "OutcomeMeasureTitle"
		break
	case "outcome_measure_description":
		output = "OutcomeMeasureDescription"
		break
	case "outcome_measure_population_description":
		output = "OutcomeMeasurePopulationDescription"
		break
	case "outcome_measure_reporting_status":
		output = "OutcomeMeasureReportingStatus"
		break
	case "outcome_measure_anticipated_posting_date":
		output = "OutcomeMeasureAnticipatedPostingDate"
		break
	case "outcome_measure_param_type":
		output = "OutcomeMeasureParamType"
		break
	case "outcome_measure_dispersion_type":
		output = "OutcomeMeasureDispersionType"
		break
	case "outcome_measure_unit_of_measure":
		output = "OutcomeMeasureUnitOfMeasure"
		break
	case "outcome_measure_calculate_pct":
		output = "OutcomeMeasureCalculatePCT"
		break
	case "outcome_measure_time_frame":
		output = "OutcomeMeasureTimeFrame"
		break
	case "outcome_measure_type_units_analyzed":
		output = "OutcomeMeasureTypeUnitsAnalyzed"
		break
	case "outcome_measure_denom_units_selected":
		output = "OutcomeMeasureDenomUnitsSelected"
		break
	case "outcome_overview_title":
		output = "OutcomeOverviewTitle"
		break
	case "outcome_overview_id":
		output = "OutcomeOverviewID"
		break
	case "outcome_overview_description":
		output = "OutcomeOverviewDescription"
		break
	case "outcome_overview_time_frame":
		output = "OutcomeOverviewTimeFrame"
		break
	case "outcome_overview_assay":
		output = "OutcomeOverviewAssay"
		break
	case "outcome_overview_dose_number":
		output = "OutcomeOverviewDoseNumber"
		break
	case "outcome_overview_measure_title":
		output = "OutcomeOverviewMeasureTitle"
		break
	case "outcome_overview_participants":
		output = "OutcomeOverviewParticipants"
		break
	case "outcome_overview_serotype":
		output = "OutcomeOverviewSerotype"
		break
	case "outcome_overview_value":
		output = "OutcomeOverviewValue"
		break
	case "outcome_overview_upper_limit":
		output = "OutcomeOverviewUpperLimit"
		break
	case "outcome_overview_lower_limit":
		output = "OutcomeOverviewLowerLimit"
		break
	case "outcome_overview_ratio":
		output = "OutcomeOverviewRatio"
		break
	case "outcome_overview_vaccine":
		output = "OutcomeOverviewVaccine"
		break
	case "outcome_overview_immunocompromised_population":
		output = "OutcomeOverviewImmunocompromisedPopulation"
		break
	case "outcome_overview_manufacturer":
		output = "OutcomeOverviewManufacturer"
		break
	case "outcome_overview_dose_description":
		output = "OutcomeOverviewDoseDescription"
		break
	case "outcome_overview_schedule":
		output = "OutcomeOverviewSchedule"
		break
	case "outcome_overview_time_frame_weeks":
		output = "OutcomeOverviewTimeFrameWeeks"
		break
	case "outcome_overview_confidence_interval":
		output = "OutcomeOverviewConfidenceInterval"
		break
	case "outcome_overview_percent_responders":
		output = "OutcomeOverviewPercentResponders"
		break
	}
	return output
}

func incomeGroupToCountryCode(incomeGroups []string) []string {
	output := []string{}
	for _, incomeGroup := range incomeGroups {
		switch incomeGroup {
		case "High income":
			output = append(output, []string{"AW", "AD", "AE", "AG", "AU", "AT", "BE", "BH", "BS", "BM", "BB", "BN", "CA", "CH", "JE", "CL", "CW", "KY", "CY", "CZ", "DE", "DK", "ES", "EE", "FI", "FR", "FO", "GB", "GI", "GR", "GL", "GU", "HK", "HR", "HU", "IM", "IE", "IS", "IL", "IT", "JP", "KM", "KR", "KW", "LI", "LT", "LU", "LV", "MO", "MF", "MC", "MT", "MP", "NC", "NL", "NO", "NR", "NZ", "OM", "PW", "PL", "PR", "PT", "PF", "QA", "SA", "SG", "SM", "SK", "SI", "SE", "SX", "SC", "TC", "TT", "UY", "US", "VG", "VI"}...)
		case "Upper middle income":
			output = append(output, []string{"AL", "AR", "AM", "AS", "AZ", "BG", "BA", "BY", "BR", "BW", "CN", "CO", "CR", "CU", "DM", "DO", "EC", "FJ", "GA", "GE", "GQ", "GD", "GT", "GY", "IQ", "JM", "JO", "KZ", "LB", "LY", "LC", "MD", "MV", "MX", "MH", "MK", "ME", "MU", "MY", "NA", "PA", "PE", "PY", "RO", "RU", "RS", "SR", "TH", "TM", "TO", "TR", "TV", "VC", "XK", "ZA"}...)
		case "Lower middle income":
			output = append(output, []string{"AO", "BJ", "BD", "BZ", "BO", "BT", "CI", "CM", "CG", "KM", "CV", "DJ", "DZ", "EG", "FM", "GH", "HN", "HT", "ID", "IN", "IR", "KE", "KG", "KH", "KI", "LA", "LK", "LS", "MA", "MM", "MN", "MR", "NG", "NI", "NP", "PK", "PH", "PG", "PS", "SN", "SB", "SV", "ST", "SZ", "TJ", "TL", "TN", "TZ", "UA", "UZ", "VN", "VU", "WS", "ZM", "ZW"}...)
		case "Low income":
			output = append(output, []string{"AF", "BI", "BF", "CF", "CD", "ER", "ET", "GN", "GM", "GW", "LR", "MG", "ML", "MZ", "MW", "NE", "KP", "RW", "SD", "SL", "SO", "SS", "SY", "TD", "TG", "UG", "YE"}...)
		case "Not classified":
			output = append(output, "VE")
		default:
			continue
		}
	}
	return output
}
