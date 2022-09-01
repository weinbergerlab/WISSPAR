package clinicaltrials

import (
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
)

func (t *ResultsContext) EditExistingTrial(studyID string, outcomes []OutcomeOverview, metadata []Metadata, useCase string, outcomeMeasures []OutcomeMeasureCustom) error {
	trial, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(studyID)).First(t.ctx)
	if err != nil {
		return err
	}

	measuresList, err := t.client.ClinicalTrial.Query().
		Where(clinicaltrial.StudyIDEQ(studyID)).
		QueryResultsDefinition().
		QueryOutcomeMeasuresModule().
		QueryOutcomeMeasureList().
		All(t.ctx)
	if err != nil {
		return err
	}

	measureModule, err := t.client.ClinicalTrial.Query().
		Where(clinicaltrial.StudyIDEQ(studyID)).
		QueryResultsDefinition().
		QueryOutcomeMeasuresModule().
		First(t.ctx)
	if err != nil {
		return err
	}

	for _, com := range outcomeMeasures {
		found := false
		for _, om := range measuresList {
			if com.OutcomeMeasureTitle == om.OutcomeMeasureTitle {
				found = true
				break
			}
		}
		if found {
			continue
		}

		_, err := t.client.OutcomeMeasure.Create().
			SetParent(measureModule).
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
	}

	measuresList, err = t.client.ClinicalTrial.Query().
		Where(clinicaltrial.StudyIDEQ(studyID)).
		QueryResultsDefinition().
		QueryOutcomeMeasuresModule().
		QueryOutcomeMeasureList().
		All(t.ctx)
	if err != nil {
		return err
	}

	measureModule, err = t.client.ClinicalTrial.Query().
		Where(clinicaltrial.StudyIDEQ(studyID)).
		QueryResultsDefinition().
		QueryOutcomeMeasuresModule().
		First(t.ctx)
	if err != nil {
		return err
	}

	if err := t.UpdateOverviewList(outcomes, measuresList, useCase); err != nil {
		return err
	}

	if _, err := t.client.ClinicalTrialMetadata.Delete().Where(clinicaltrialmetadata.HasParentWith(clinicaltrial.IDEQ(trial.ID)), clinicaltrialmetadata.UseCaseCode(useCase)).Exec(t.ctx); err != nil {
		return err
	}

	for _, md := range metadata {
		if _, err := t.client.ClinicalTrialMetadata.Create().
			SetParent(trial).
			SetTagName(md.TagName).
			SetTagValue(md.TagValue).
			SetUseCaseCode(md.TagUseCase).
			Save(t.ctx); err != nil {
				return err
		}
	}

	return nil
}

func (t *ResultsContext) UpdateOverviewList(outcomes []OutcomeOverview, measureList []*models.OutcomeMeasure, useCase string) error {
	for _, measure := range measureList {
		if _, err := t.client.OutcomeOverview.Delete().Where(
				outcomeoverview.HasParentWith(outcomemeasure.IDEQ(measure.ID)),
				outcomeoverview.OutcomeOverviewUseCaseCode(useCase),
			).Exec(t.ctx); err != nil {
			return err
		}
	}

	finalOutcomes := []*models.OutcomeOverview{}

	for _, outcome := range outcomes {
		oo, err := t.client.OutcomeOverview.Create().
			SetOutcomeOverviewMeasureTitle(outcome.OutcomeMeasureTitle).
			SetOutcomeOverviewTitle(outcome.OutcomeOverviewTitle).
			SetOutcomeOverviewParticipants(outcome.OutcomeOverviewParticipants).
			SetOutcomeOverviewAssay(outcome.OutcomeOverviewAssay).
			SetOutcomeOverviewDescription(outcome.OutcomeOverviewDescription).
			SetOutcomeOverviewDoseNumber(outcome.OutcomeOverviewDoseNumber).
			SetOutcomeOverviewGroupID(outcome.OutcomeOverviewGroupID).
			SetOutcomeOverviewRatio(outcome.OutcomeOverviewRatio).
			SetOutcomeOverviewID(outcome.OutcomeOverviewID).
			SetOutcomeOverviewSerotype(outcome.OutcomeOverviewSerotype).
			SetOutcomeOverviewTimeFrame(outcome.OutcomeOverviewTimeFrame).
			SetOutcomeOverviewValue(outcome.OutcomeOverviewValue).
			SetOutcomeOverviewLowerLimit(outcome.OutcomeOverviewLowerLimit).
			SetOutcomeOverviewTimeFrameWeeks(outcome.OutcomeOverviewTimeFrameWeeks).
			SetOutcomeOverviewVaccine(outcome.OutcomeOverviewVaccine).
			SetOutcomeOverviewImmunocompromisedPopulation(outcome.OutcomeOverviewImmunocompromisedPopulation).
			SetOutcomeOverviewManufacturer(outcome.OutcomeOverviewManufacturer).
			SetOutcomeOverviewUpperLimit(outcome.OutcomeOverviewUpperLimit).
			SetOutcomeOverviewConfidenceInterval(outcome.OutcomeOverviewConfidenceInterval).
			SetOutcomeOverviewDoseDescription(outcome.OutcomeOverviewDoseDescription).
			SetOutcomeOverviewSchedule(outcome.OutcomeOverviewSchedule).
			SetOutcomeOverviewPercentResponders(outcome.OutcomeOverviewPercentResponders).
			SetOutcomeOverviewUseCaseCode(outcome.OutcomeOverviewUseCase).
			Save(t.ctx)
		if err != nil {
			return err
		}

		finalOutcomes = append(finalOutcomes, oo)
	}

	for _, measure := range measureList {
		for _, outcomeOverview := range finalOutcomes {
			if outcomeOverview.OutcomeOverviewMeasureTitle == measure.OutcomeMeasureTitle {
				if _, err := outcomeOverview.Update().SetParent(measure).Save(t.ctx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
