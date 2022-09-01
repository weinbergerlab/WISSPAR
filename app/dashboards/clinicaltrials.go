package dashboards

import (
	"context"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
	"strings"
)

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

type ClinicalTrialResult struct {
	NCTID string `json:"nct_id"`
	StudyName string `json:"study_name"`
	Phase string `json:"phase"`
	Vaccine []string `json:"vaccine"`
	Schedule []string `json:"schedule"`
	Manufacturer []string `json:"manufacturer"`
	AgeList string `json:"age_list"`
	Ethnicity string `json:"ethnicity"`
	LocationContinents []string `json:"location_continents"`
	ImmunocompromisedGroups []string `json:"immunocompromised_groups"`
	Gender string `json:"gender"`
	OutcomesCuratedPerUseCase map[string]int `json:"outcomes_curated_per_use_case"`
}

func (t *ResultsContext) GetClinicalTrial(isOr bool, agelist []string, vaccines []string, schedules []string, manufacturers []string, continents []string, immunocompromisedGroups []string) ([]ClinicalTrialResult, error) {
	trialWhere := []predicate.ClinicalTrial{}

	if len(agelist) > 0 {
		studyEligibilityWhere := []predicate.StudyEligibility{}
		for _, age := range agelist {
			studyEligibilityWhere = append(studyEligibilityWhere, studyeligibility.StdAgeListContains(age))
		}
		trialWhere = append(trialWhere, clinicaltrial.HasStudyEligibilityWith(studyeligibility.Or(studyEligibilityWhere...)))
	}

	if len(continents) > 0 {
		studyLocationWhere := []predicate.StudyLocation{}
		for _, continent := range continents {
			studyLocationWhere = append(studyLocationWhere, studylocation.LocationContinentNameContains(continent))
		}
		trialWhere = append(trialWhere, clinicaltrial.HasStudyLocationsWith(studylocation.Or(studyLocationWhere...)))
	}

	if len(vaccines) > 0 || len(manufacturers) > 0  || len(schedules) > 0{
		overviewVaccineWhere := []predicate.OutcomeOverview{}
		overviewManufacturerWhere := []predicate.OutcomeOverview{}
		overviewScheduleWhere := []predicate.OutcomeOverview{}
		overviewImmunocompromisedGroupsWhere := []predicate.OutcomeOverview{}

		for _, vaccine := range vaccines {
			overviewVaccineWhere = append(overviewVaccineWhere, outcomeoverview.OutcomeOverviewVaccine(vaccine))
		}

		for _, manufacturer := range manufacturers {
			overviewManufacturerWhere = append(overviewManufacturerWhere, outcomeoverview.OutcomeOverviewManufacturer(manufacturer))
		}

		for _, schedule := range schedules {
			overviewScheduleWhere = append(overviewScheduleWhere, outcomeoverview.OutcomeOverviewSchedule(schedule))
		}

		for _, ig := range immunocompromisedGroups {
			overviewImmunocompromisedGroupsWhere = append(overviewImmunocompromisedGroupsWhere, outcomeoverview.OutcomeOverviewImmunocompromisedPopulationContains(ig))
		}

		overviewWhere := []predicate.OutcomeOverview{}
		if len(vaccines) > 0 {
			overviewWhere = append(overviewWhere, outcomeoverview.Or(overviewVaccineWhere...))
		}

		if len(manufacturers) > 0 {
			overviewWhere = append(overviewWhere, outcomeoverview.Or(overviewManufacturerWhere...))
		}

		if len(schedules) > 0 {
			overviewWhere = append(overviewWhere, outcomeoverview.Or(overviewScheduleWhere...))
		}

		if len(immunocompromisedGroups) > 0 {
			overviewWhere = append(overviewWhere, outcomeoverview.Or(overviewImmunocompromisedGroupsWhere...))
		}

		trialWhere = append(trialWhere, clinicaltrial.HasResultsDefinitionWith(resultsdefinition.HasOutcomeMeasuresModuleWith(outcomemeasuresmodule.HasOutcomeMeasureListWith(outcomemeasure.HasOutcomeOverviewListWith(outcomeoverview.And(overviewWhere...))))))
	}

	trials, err := t.client.ClinicalTrial.Query().WithStudyLocations().WithStudyEligibility().WithResultsDefinition(
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
		}).Where(trialWhere...).All(t.ctx)
	if err != nil {
		return nil, err
	}

	results := []ClinicalTrialResult{}

	for _, trial := range trials {
		result := ClinicalTrialResult{}
		result.NCTID = trial.StudyID
		result.StudyName = trial.StudyName
		result.Phase = trial.Phase

		if rd, err := trial.Edges.ResultsDefinitionOrErr(); err == nil {
			if omm, err := rd.Edges.OutcomeMeasuresModuleOrErr(); err == nil {
				if oml, err := omm.Edges.OutcomeMeasureListOrErr(); err == nil {

					trialVaccines := []string{}
					trialSchedules := []string{}
					trialManufacturers := []string{}
					trialImmunocompromisedGroups := []string{}

					for _, om := range oml {
						if ovl, err := om.Edges.OutcomeOverviewListOrErr(); err == nil {
							for _, ov := range ovl {
								if result.OutcomesCuratedPerUseCase == nil {
									result.OutcomesCuratedPerUseCase = make(map[string]int)
								}
								result.OutcomesCuratedPerUseCase[ov.OutcomeOverviewUseCaseCode] = result.OutcomesCuratedPerUseCase[ov.OutcomeOverviewUseCaseCode] + 1
								trialVaccines = distinctValueAdd(trialVaccines, ov.OutcomeOverviewVaccine)
								trialSchedules = distinctValueAdd(trialSchedules, ov.OutcomeOverviewSchedule)
								trialManufacturers = distinctValueAdd(trialManufacturers, ov.OutcomeOverviewManufacturer)

								if ov.OutcomeOverviewImmunocompromisedPopulation != "" {
									if strings.Contains(ov.OutcomeOverviewImmunocompromisedPopulation, ",") {
										for _, ig := range strings.Split(ov.OutcomeOverviewImmunocompromisedPopulation, ",") {
											trialImmunocompromisedGroups = distinctValueAdd(trialImmunocompromisedGroups, ig)
										}
									} else {
										trialImmunocompromisedGroups = distinctValueAdd(trialImmunocompromisedGroups, ov.OutcomeOverviewImmunocompromisedPopulation)
									}
								}
							}
						}
					}

					if !isOr {
						vaccineSearchCount := len(vaccines)
						actualVaccineCount := 0
						for _, trialVaccine := range trialVaccines {
							if vaccineSearchCount == actualVaccineCount {
								break
							}
							for _, s := range vaccines {
								if trialVaccine == s {
									actualVaccineCount++
								}
							}
						}

						if vaccineSearchCount != actualVaccineCount {
							continue
						}
					}

					result.Vaccine = trialVaccines
					result.Schedule = trialSchedules
					result.Manufacturer = trialManufacturers
					result.ImmunocompromisedGroups = trialImmunocompromisedGroups
				}
			}
		}

		if se, err := trial.Edges.StudyEligibilityOrErr(); err == nil {
			for _, ses := range se {
				result.AgeList = ses.StdAgeList
				result.Gender = ses.Gender
				result.Ethnicity = ses.Ethnicity
			}
		}

		if sls, err := trial.Edges.StudyLocationsOrErr(); err == nil {
			continents := []string{}
			for _, sl := range sls {
				cn := strings.Replace(sl.LocationContinentName, ",", "/", -1)
				continents = distinctValueAdd(continents, cn)
			}
			result.LocationContinents = continents
		}

		if len(immunocompromisedGroups) > 0 {
			hasGroup := false
			for _, immunocompromisedGroup := range immunocompromisedGroups  {
				for _, selectedGroup := range result.ImmunocompromisedGroups {
					if immunocompromisedGroup == selectedGroup {
						hasGroup = true
						break
					}
				}
				if len(result.ImmunocompromisedGroups) == 0 && immunocompromisedGroup == "" {
					hasGroup = true
				}
				if hasGroup {
					break
				}
			}
			if !hasGroup {
				continue
			}
		}

		results = append(results, result)
	}

	return results, nil
}

func distinctValueAdd(arr []string, str string) []string {
	if str != "" && !needleInHaystack(str, arr) {
		arr = append(arr, str)
	}

	return arr
}

func needleInHaystack(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}
