package clinicaltrials

import (
	"encoding/json"
	"fmt"
	"github.com/DrGrimshaw/gohtml"
	"github.com/rs/zerolog/log"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const outsideTrialError = "incorrect prefix for study ID"

func (t *ResultsContext) GetTrial(ID string) (Study, error) {
	if !strings.Contains(strings.ToUpper(ID), "NCT") { return Study{}, fmt.Errorf(outsideTrialError) }
	apiFetch := fmt.Sprintf(apiURL, ID)

	resp, err := http.Get(apiFetch)
	if err != nil {
		return Study{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		return Study{}, fmt.Errorf("unable to unmarshal data to response struct, err: %s", err)
	}

	if len(response.FullStudiesResponse.FullStudies) == 0 {
		return Study{}, fmt.Errorf("studies in response for ID: %s", ID)
	}

	return response.FullStudiesResponse.FullStudies[0].Study, nil
}

func (t *ResultsContext) GetOutcomeOverviewList(study Study, useCase string) map[string]map[string][]OutcomeOverview {
	outcomeOverviews, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(study.ProtocolSection.IdentificationModule.NCTID)).QueryResultsDefinition().QueryOutcomeMeasuresModule().QueryOutcomeMeasureList().QueryOutcomeOverviewList().Where(outcomeoverview.OutcomeOverviewUseCaseCodeEQ(useCase)).All(t.ctx)

	sorting := map[string]map[string][]OutcomeOverview{}

	if err == nil && len(outcomeOverviews) > 0 {
		for _, overview := range outcomeOverviews {
			if sorting[overview.OutcomeOverviewMeasureTitle] == nil {
				sorting[overview.OutcomeOverviewMeasureTitle] = make(map[string][]OutcomeOverview)
			}

			if sorting[overview.OutcomeOverviewMeasureTitle][overview.OutcomeOverviewTitle] == nil {
				sorting[overview.OutcomeOverviewMeasureTitle][overview.OutcomeOverviewTitle] = []OutcomeOverview{}
			}

			sorting[overview.OutcomeOverviewMeasureTitle][overview.OutcomeOverviewTitle] = append(sorting[overview.OutcomeOverviewMeasureTitle][overview.OutcomeOverviewTitle], OutcomeOverview{
				OutcomeOverviewTitle:                       overview.OutcomeOverviewTitle,
				OutcomeOverviewID:                          overview.OutcomeOverviewID,
				OutcomeOverviewDescription:                 overview.OutcomeOverviewDescription,
				OutcomeOverviewTimeFrame:                   overview.OutcomeOverviewTimeFrame,
				OutcomeOverviewAssay:                       overview.OutcomeOverviewAssay,
				OutcomeOverviewDoseNumber:                  overview.OutcomeOverviewDoseNumber,
				OutcomeMeasureTitle:                        overview.OutcomeOverviewMeasureTitle,
				OutcomeOverviewParticipants:                overview.OutcomeOverviewParticipants,
				OutcomeOverviewSerotype:                    overview.OutcomeOverviewSerotype,
				OutcomeOverviewValue:                       overview.OutcomeOverviewValue,
				OutcomeOverviewUpperLimit:                  overview.OutcomeOverviewUpperLimit,
				OutcomeOverviewLowerLimit:                  overview.OutcomeOverviewLowerLimit,
				OutcomeOverviewGroupID:                     overview.OutcomeOverviewGroupID,
				OutcomeOverviewRatio:                       overview.OutcomeOverviewRatio,
				OutcomeOverviewVaccine:                     overview.OutcomeOverviewVaccine,
				OutcomeOverviewImmunocompromisedPopulation: overview.OutcomeOverviewImmunocompromisedPopulation,
				OutcomeOverviewManufacturer:                overview.OutcomeOverviewManufacturer,
				OutcomeOverviewTimeFrameWeeks:              overview.OutcomeOverviewTimeFrameWeeks,
				OutcomeOverviewConfidenceInterval:          overview.OutcomeOverviewConfidenceInterval,
				OutcomeOverviewDoseDescription:             overview.OutcomeOverviewDoseDescription,
				OutcomeOverviewSchedule:                    overview.OutcomeOverviewSchedule,
				OutcomeOverviewPercentResponders:           overview.OutcomeOverviewPercentResponders,
				OutcomeOverviewUseCase:                     overview.OutcomeOverviewUseCaseCode,
				Enabled:                                    false,
			})
		}
	}

	for _, outcomeMeasure := range study.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure {
		serotypeCheck := regexp.MustCompile(`(?mi)(\bserotype|\bserotypes|\bserotype-specific)`)
		if !serotypeCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) &&
			!serotypeCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) &&
			!serotypeCheck.Match([]byte(outcomeMeasure.OutcomeMeasurePopulationDescription)) {
			continue
		}

		dose := 0

		flowPeriod := study.ResultsSection.ParticipantFlowModule.FlowPeriodList.FlowPeriod

		assay := "Other"

		if gmcCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) ||
			gmcCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
			gmcCheck.Match([]byte(outcomeMeasure.OutcomeMeasureParamType)) {
			assay = "GMC"
		} else if opaCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) ||
			opaCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
			opaCheck.Match([]byte(outcomeMeasure.OutcomeMeasureParamType)) {
			assay = "OPA"
		} else if gmtCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) ||
			gmtCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
			gmtCheck.Match([]byte(outcomeMeasure.OutcomeMeasureParamType)) {
			assay = "GMT"
		}

		for i, outcomeGroup := range outcomeMeasure.OutcomeGroupList.OutcomeGroup {
			if len(flowPeriod) > 0 {
				flowMilestoneArr := study.ResultsSection.ParticipantFlowModule.FlowPeriodList.FlowPeriod[0].FlowMilestoneList.FlowMilestone
				if len(flowMilestoneArr) > i {
					if firstDoseCheckFlow.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) ||
						firstDoseCheck.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) {
						dose = 1
					} else if secondDoseCheckFlow.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) ||
						secondDoseCheck.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) {
						dose = 2
					} else if thirdDoseCheckFlow.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) ||
						thirdDoseCheck.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) {
						dose = 3
					} else if fourthDoseCheckFlow.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) ||
						fourthDoseCheck.Match([]byte(flowMilestoneArr[i].FlowMilestoneType)) {
						dose = 4
					}
				}
			}

			if dose == 0 {
				if firstDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTimeFrame)) ||
					firstDoseCheck.Match([]byte(outcomeMeasure.OutcomeGroupList.OutcomeGroup[i].OutcomeGroupDescription)) ||
					firstDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasurePopulationDescription)) ||
					firstDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
					firstDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) {
					dose = 1
				} else if secondDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTimeFrame)) ||
					secondDoseCheck.Match([]byte(outcomeMeasure.OutcomeGroupList.OutcomeGroup[i].OutcomeGroupDescription)) ||
					secondDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasurePopulationDescription)) ||
					secondDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
					secondDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) {
					dose = 2
				} else if thirdDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTimeFrame)) ||
					thirdDoseCheck.Match([]byte(outcomeMeasure.OutcomeGroupList.OutcomeGroup[i].OutcomeGroupDescription)) ||
					thirdDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasurePopulationDescription)) ||
					thirdDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
					thirdDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) {
					dose = 3
				} else if fourthDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTimeFrame)) ||
					fourthDoseCheck.Match([]byte(outcomeMeasure.OutcomeGroupList.OutcomeGroup[i].OutcomeGroupDescription)) ||
					fourthDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasurePopulationDescription)) ||
					fourthDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureDescription)) ||
					fourthDoseCheck.Match([]byte(outcomeMeasure.OutcomeMeasureTitle)) {
					dose = 4
				}
			}

			outcomeOverview := OutcomeOverview{
				OutcomeOverviewTitle:        outcomeGroup.OutcomeGroupTitle,
				OutcomeOverviewID:           outcomeGroup.OutcomeGroupID,
				OutcomeOverviewDescription:  outcomeGroup.OutcomeGroupDescription,
				OutcomeOverviewTimeFrame:    outcomeMeasure.OutcomeMeasureTimeFrame,
				OutcomeOverviewAssay:        assay,
				OutcomeOverviewDoseNumber:   int64(dose),
				OutcomeMeasureTitle:         outcomeMeasure.OutcomeMeasureTitle,
				OutcomeOverviewParticipants: "0",
				OutcomeOverviewUseCase:      useCase,
				Enabled:                     true,
			}

			outcomeDenoms := outcomeMeasure.OutcomeDenomList.OutcomeDenom
			if len(outcomeDenoms) > 0 {
				if len(outcomeDenoms[0].OutcomeDenomCountList.OutcomeDenomCount) > i {
					outcomeOverview.OutcomeOverviewParticipants = outcomeDenoms[0].OutcomeDenomCountList.OutcomeDenomCount[i].OutcomeDenomCountValue
				}
			}

			for _, outcomeClass := range outcomeMeasure.OutcomeClassList.OutcomeClass {
				serotype := getSerotype(outcomeClass.OutcomeClassTitle)

				if dose == 0 {
					if firstDoseCheck.Match([]byte(outcomeClass.OutcomeClassTitle)) {
						dose = 1
					} else if secondDoseCheck.Match([]byte(outcomeClass.OutcomeClassTitle)) {
						dose = 2
					} else if thirdDoseCheck.Match([]byte(outcomeClass.OutcomeClassTitle)) {
						dose = 3
					} else if fourthDoseCheck.Match([]byte(outcomeClass.OutcomeClassTitle)) {
						dose = 4
					}
				}

				thisOutcomeOverview := outcomeOverview
				thisOutcomeOverview.OutcomeOverviewSerotype = serotype
				thisOutcomeOverview.OutcomeOverviewValue = 0
				thisOutcomeOverview.OutcomeOverviewUpperLimit = "0"
				thisOutcomeOverview.OutcomeOverviewLowerLimit = "0"

				outcomeCategories := outcomeClass.OutcomeCategoryList.OutcomeCategory
				if len(outcomeCategories) > 0 {
					measurements := outcomeCategories[0].OutcomeMeasurementList.OutcomeMeasurement
					if len(measurements) > i {
						value, err := strconv.ParseFloat(measurements[i].OutcomeMeasurementValue, 64)
						if err != nil {
							log.Err(err)
							value = -1
						}

						thisOutcomeOverview.OutcomeOverviewValue = value
						thisOutcomeOverview.OutcomeOverviewUpperLimit = measurements[i].OutcomeMeasurementUpperLimit
						thisOutcomeOverview.OutcomeOverviewLowerLimit = measurements[i].OutcomeMeasurementLowerLimit
					}
				}
				if len(sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle]) == 0 {
					if sorting[thisOutcomeOverview.OutcomeMeasureTitle] == nil {
						sorting[thisOutcomeOverview.OutcomeMeasureTitle] = make(map[string][]OutcomeOverview)
					}

					if sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle] == nil {
						sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle] = []OutcomeOverview{}
					}
				}

				overviewExists := false

				for _, compareOverview := range sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle] {
					if thisOutcomeOverview.OutcomeOverviewSerotype == compareOverview.OutcomeOverviewSerotype {
						overviewExists = true
						break
					}
				}

				if overviewExists {
					break
				}

				sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle] = append(sorting[thisOutcomeOverview.OutcomeMeasureTitle][thisOutcomeOverview.OutcomeOverviewTitle], thisOutcomeOverview)
			}
		}
	}

	return sorting
}

func GetTrialHTML(study Study) (string, error) {
	return gohtml.Encode(study)
}

func GetProtocolHTML(study Study) (string, error) {
	return gohtml.Encode(study.ProtocolSection)
}

func GetDerivedHTML(study Study) (string, error) {
	return gohtml.Encode(study.DerivedSection)
}

func GetResultsHTML(study Study) (string, error) {
	return gohtml.Encode(study.ResultsSection)
}

func GetDocumentHTML(study Study) (string, error) {
	return gohtml.Encode(study.DocumentSection)
}
