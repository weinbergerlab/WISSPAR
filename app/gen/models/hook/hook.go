// Code generated (@generated) by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
)

// The AdverseEventsModuleFunc type is an adapter to allow the use of ordinary
// function as AdverseEventsModule mutator.
type AdverseEventsModuleFunc func(context.Context, *models.AdverseEventsModuleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f AdverseEventsModuleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.AdverseEventsModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.AdverseEventsModuleMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineCategoryFunc type is an adapter to allow the use of ordinary
// function as BaselineCategory mutator.
type BaselineCategoryFunc func(context.Context, *models.BaselineCategoryMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineCategoryFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineCategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineCategoryMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineCharacteristicsModuleFunc type is an adapter to allow the use of ordinary
// function as BaselineCharacteristicsModule mutator.
type BaselineCharacteristicsModuleFunc func(context.Context, *models.BaselineCharacteristicsModuleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineCharacteristicsModuleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineCharacteristicsModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineCharacteristicsModuleMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineClassFunc type is an adapter to allow the use of ordinary
// function as BaselineClass mutator.
type BaselineClassFunc func(context.Context, *models.BaselineClassMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineClassFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineClassMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineClassMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineClassDenomFunc type is an adapter to allow the use of ordinary
// function as BaselineClassDenom mutator.
type BaselineClassDenomFunc func(context.Context, *models.BaselineClassDenomMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineClassDenomFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineClassDenomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineClassDenomMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineClassDenomCountFunc type is an adapter to allow the use of ordinary
// function as BaselineClassDenomCount mutator.
type BaselineClassDenomCountFunc func(context.Context, *models.BaselineClassDenomCountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineClassDenomCountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineClassDenomCountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineClassDenomCountMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineDenomFunc type is an adapter to allow the use of ordinary
// function as BaselineDenom mutator.
type BaselineDenomFunc func(context.Context, *models.BaselineDenomMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineDenomFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineDenomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineDenomMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineDenomCountFunc type is an adapter to allow the use of ordinary
// function as BaselineDenomCount mutator.
type BaselineDenomCountFunc func(context.Context, *models.BaselineDenomCountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineDenomCountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineDenomCountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineDenomCountMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineGroupFunc type is an adapter to allow the use of ordinary
// function as BaselineGroup mutator.
type BaselineGroupFunc func(context.Context, *models.BaselineGroupMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineGroupFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineGroupMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineMeasureFunc type is an adapter to allow the use of ordinary
// function as BaselineMeasure mutator.
type BaselineMeasureFunc func(context.Context, *models.BaselineMeasureMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineMeasureFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineMeasureMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineMeasureMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineMeasureDenomFunc type is an adapter to allow the use of ordinary
// function as BaselineMeasureDenom mutator.
type BaselineMeasureDenomFunc func(context.Context, *models.BaselineMeasureDenomMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineMeasureDenomFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineMeasureDenomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineMeasureDenomMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineMeasureDenomCountFunc type is an adapter to allow the use of ordinary
// function as BaselineMeasureDenomCount mutator.
type BaselineMeasureDenomCountFunc func(context.Context, *models.BaselineMeasureDenomCountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineMeasureDenomCountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineMeasureDenomCountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineMeasureDenomCountMutation", m)
	}
	return f(ctx, mv)
}

// The BaselineMeasurementFunc type is an adapter to allow the use of ordinary
// function as BaselineMeasurement mutator.
type BaselineMeasurementFunc func(context.Context, *models.BaselineMeasurementMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f BaselineMeasurementFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.BaselineMeasurementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.BaselineMeasurementMutation", m)
	}
	return f(ctx, mv)
}

// The CertainAgreementFunc type is an adapter to allow the use of ordinary
// function as CertainAgreement mutator.
type CertainAgreementFunc func(context.Context, *models.CertainAgreementMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f CertainAgreementFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.CertainAgreementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.CertainAgreementMutation", m)
	}
	return f(ctx, mv)
}

// The ClinicalTrialFunc type is an adapter to allow the use of ordinary
// function as ClinicalTrial mutator.
type ClinicalTrialFunc func(context.Context, *models.ClinicalTrialMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ClinicalTrialFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ClinicalTrialMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ClinicalTrialMutation", m)
	}
	return f(ctx, mv)
}

// The ClinicalTrialMetadataFunc type is an adapter to allow the use of ordinary
// function as ClinicalTrialMetadata mutator.
type ClinicalTrialMetadataFunc func(context.Context, *models.ClinicalTrialMetadataMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ClinicalTrialMetadataFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ClinicalTrialMetadataMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ClinicalTrialMetadataMutation", m)
	}
	return f(ctx, mv)
}

// The DoseDescriptionFunc type is an adapter to allow the use of ordinary
// function as DoseDescription mutator.
type DoseDescriptionFunc func(context.Context, *models.DoseDescriptionMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f DoseDescriptionFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.DoseDescriptionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.DoseDescriptionMutation", m)
	}
	return f(ctx, mv)
}

// The EventGroupFunc type is an adapter to allow the use of ordinary
// function as EventGroup mutator.
type EventGroupFunc func(context.Context, *models.EventGroupMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f EventGroupFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.EventGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.EventGroupMutation", m)
	}
	return f(ctx, mv)
}

// The FlowAchievementFunc type is an adapter to allow the use of ordinary
// function as FlowAchievement mutator.
type FlowAchievementFunc func(context.Context, *models.FlowAchievementMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowAchievementFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowAchievementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowAchievementMutation", m)
	}
	return f(ctx, mv)
}

// The FlowDropWithdrawFunc type is an adapter to allow the use of ordinary
// function as FlowDropWithdraw mutator.
type FlowDropWithdrawFunc func(context.Context, *models.FlowDropWithdrawMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowDropWithdrawFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowDropWithdrawMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowDropWithdrawMutation", m)
	}
	return f(ctx, mv)
}

// The FlowGroupFunc type is an adapter to allow the use of ordinary
// function as FlowGroup mutator.
type FlowGroupFunc func(context.Context, *models.FlowGroupMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowGroupFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowGroupMutation", m)
	}
	return f(ctx, mv)
}

// The FlowMilestoneFunc type is an adapter to allow the use of ordinary
// function as FlowMilestone mutator.
type FlowMilestoneFunc func(context.Context, *models.FlowMilestoneMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowMilestoneFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowMilestoneMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowMilestoneMutation", m)
	}
	return f(ctx, mv)
}

// The FlowPeriodFunc type is an adapter to allow the use of ordinary
// function as FlowPeriod mutator.
type FlowPeriodFunc func(context.Context, *models.FlowPeriodMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowPeriodFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowPeriodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowPeriodMutation", m)
	}
	return f(ctx, mv)
}

// The FlowReasonFunc type is an adapter to allow the use of ordinary
// function as FlowReason mutator.
type FlowReasonFunc func(context.Context, *models.FlowReasonMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f FlowReasonFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.FlowReasonMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.FlowReasonMutation", m)
	}
	return f(ctx, mv)
}

// The ImmunocompromisedGroupsFunc type is an adapter to allow the use of ordinary
// function as ImmunocompromisedGroups mutator.
type ImmunocompromisedGroupsFunc func(context.Context, *models.ImmunocompromisedGroupsMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ImmunocompromisedGroupsFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ImmunocompromisedGroupsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ImmunocompromisedGroupsMutation", m)
	}
	return f(ctx, mv)
}

// The ManufacturerFunc type is an adapter to allow the use of ordinary
// function as Manufacturer mutator.
type ManufacturerFunc func(context.Context, *models.ManufacturerMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ManufacturerFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ManufacturerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ManufacturerMutation", m)
	}
	return f(ctx, mv)
}

// The MoreInfoModuleFunc type is an adapter to allow the use of ordinary
// function as MoreInfoModule mutator.
type MoreInfoModuleFunc func(context.Context, *models.MoreInfoModuleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f MoreInfoModuleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.MoreInfoModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.MoreInfoModuleMutation", m)
	}
	return f(ctx, mv)
}

// The OtherEventFunc type is an adapter to allow the use of ordinary
// function as OtherEvent mutator.
type OtherEventFunc func(context.Context, *models.OtherEventMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OtherEventFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OtherEventMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OtherEventMutation", m)
	}
	return f(ctx, mv)
}

// The OtherEventStatsFunc type is an adapter to allow the use of ordinary
// function as OtherEventStats mutator.
type OtherEventStatsFunc func(context.Context, *models.OtherEventStatsMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OtherEventStatsFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OtherEventStatsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OtherEventStatsMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeAnalysisFunc type is an adapter to allow the use of ordinary
// function as OutcomeAnalysis mutator.
type OutcomeAnalysisFunc func(context.Context, *models.OutcomeAnalysisMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeAnalysisFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeAnalysisMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeAnalysisMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeAnalysisGroupIDFunc type is an adapter to allow the use of ordinary
// function as OutcomeAnalysisGroupID mutator.
type OutcomeAnalysisGroupIDFunc func(context.Context, *models.OutcomeAnalysisGroupIDMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeAnalysisGroupIDFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeAnalysisGroupIDMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeAnalysisGroupIDMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeCategoryFunc type is an adapter to allow the use of ordinary
// function as OutcomeCategory mutator.
type OutcomeCategoryFunc func(context.Context, *models.OutcomeCategoryMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeCategoryFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeCategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeCategoryMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeClassFunc type is an adapter to allow the use of ordinary
// function as OutcomeClass mutator.
type OutcomeClassFunc func(context.Context, *models.OutcomeClassMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeClassFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeClassMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeClassMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeClassDenomFunc type is an adapter to allow the use of ordinary
// function as OutcomeClassDenom mutator.
type OutcomeClassDenomFunc func(context.Context, *models.OutcomeClassDenomMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeClassDenomFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeClassDenomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeClassDenomMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeClassDenomCountFunc type is an adapter to allow the use of ordinary
// function as OutcomeClassDenomCount mutator.
type OutcomeClassDenomCountFunc func(context.Context, *models.OutcomeClassDenomCountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeClassDenomCountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeClassDenomCountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeClassDenomCountMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeDenomFunc type is an adapter to allow the use of ordinary
// function as OutcomeDenom mutator.
type OutcomeDenomFunc func(context.Context, *models.OutcomeDenomMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeDenomFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeDenomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeDenomMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeDenomCountFunc type is an adapter to allow the use of ordinary
// function as OutcomeDenomCount mutator.
type OutcomeDenomCountFunc func(context.Context, *models.OutcomeDenomCountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeDenomCountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeDenomCountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeDenomCountMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeGroupFunc type is an adapter to allow the use of ordinary
// function as OutcomeGroup mutator.
type OutcomeGroupFunc func(context.Context, *models.OutcomeGroupMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeGroupFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeGroupMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeMeasureFunc type is an adapter to allow the use of ordinary
// function as OutcomeMeasure mutator.
type OutcomeMeasureFunc func(context.Context, *models.OutcomeMeasureMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeMeasureFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeMeasureMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeMeasureMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeMeasurementFunc type is an adapter to allow the use of ordinary
// function as OutcomeMeasurement mutator.
type OutcomeMeasurementFunc func(context.Context, *models.OutcomeMeasurementMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeMeasurementFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeMeasurementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeMeasurementMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeMeasuresModuleFunc type is an adapter to allow the use of ordinary
// function as OutcomeMeasuresModule mutator.
type OutcomeMeasuresModuleFunc func(context.Context, *models.OutcomeMeasuresModuleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeMeasuresModuleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeMeasuresModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeMeasuresModuleMutation", m)
	}
	return f(ctx, mv)
}

// The OutcomeOverviewFunc type is an adapter to allow the use of ordinary
// function as OutcomeOverview mutator.
type OutcomeOverviewFunc func(context.Context, *models.OutcomeOverviewMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f OutcomeOverviewFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.OutcomeOverviewMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.OutcomeOverviewMutation", m)
	}
	return f(ctx, mv)
}

// The ParticipantFlowModuleFunc type is an adapter to allow the use of ordinary
// function as ParticipantFlowModule mutator.
type ParticipantFlowModuleFunc func(context.Context, *models.ParticipantFlowModuleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ParticipantFlowModuleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ParticipantFlowModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ParticipantFlowModuleMutation", m)
	}
	return f(ctx, mv)
}

// The PointOfContactFunc type is an adapter to allow the use of ordinary
// function as PointOfContact mutator.
type PointOfContactFunc func(context.Context, *models.PointOfContactMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f PointOfContactFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.PointOfContactMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.PointOfContactMutation", m)
	}
	return f(ctx, mv)
}

// The ResultsDefinitionFunc type is an adapter to allow the use of ordinary
// function as ResultsDefinition mutator.
type ResultsDefinitionFunc func(context.Context, *models.ResultsDefinitionMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ResultsDefinitionFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ResultsDefinitionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ResultsDefinitionMutation", m)
	}
	return f(ctx, mv)
}

// The ScheduleFunc type is an adapter to allow the use of ordinary
// function as Schedule mutator.
type ScheduleFunc func(context.Context, *models.ScheduleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f ScheduleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.ScheduleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.ScheduleMutation", m)
	}
	return f(ctx, mv)
}

// The SeriousEventFunc type is an adapter to allow the use of ordinary
// function as SeriousEvent mutator.
type SeriousEventFunc func(context.Context, *models.SeriousEventMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f SeriousEventFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.SeriousEventMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.SeriousEventMutation", m)
	}
	return f(ctx, mv)
}

// The SeriousEventStatsFunc type is an adapter to allow the use of ordinary
// function as SeriousEventStats mutator.
type SeriousEventStatsFunc func(context.Context, *models.SeriousEventStatsMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f SeriousEventStatsFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.SeriousEventStatsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.SeriousEventStatsMutation", m)
	}
	return f(ctx, mv)
}

// The StudyEligibilityFunc type is an adapter to allow the use of ordinary
// function as StudyEligibility mutator.
type StudyEligibilityFunc func(context.Context, *models.StudyEligibilityMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f StudyEligibilityFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.StudyEligibilityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.StudyEligibilityMutation", m)
	}
	return f(ctx, mv)
}

// The StudyLocationFunc type is an adapter to allow the use of ordinary
// function as StudyLocation mutator.
type StudyLocationFunc func(context.Context, *models.StudyLocationMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f StudyLocationFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.StudyLocationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.StudyLocationMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *models.TaskMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.TaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.TaskMutation", m)
	}
	return f(ctx, mv)
}

// The UseCaseFunc type is an adapter to allow the use of ordinary
// function as UseCase mutator.
type UseCaseFunc func(context.Context, *models.UseCaseMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f UseCaseFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.UseCaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.UseCaseMutation", m)
	}
	return f(ctx, mv)
}

// The VaccineFunc type is an adapter to allow the use of ordinary
// function as Vaccine mutator.
type VaccineFunc func(context.Context, *models.VaccineMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f VaccineFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.VaccineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.VaccineMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, models.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op models.Op) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk models.Hook, cond Condition) models.Hook {
	return func(next models.Mutator) models.Mutator {
		return models.MutateFunc(func(ctx context.Context, m models.Mutation) (models.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, models.Delete|models.Create)
//
func On(hk models.Hook, op models.Op) models.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, models.Update|models.UpdateOne)
//
func Unless(hk models.Hook, op models.Op) models.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) models.Hook {
	return func(models.Mutator) models.Mutator {
		return models.MutateFunc(func(context.Context, models.Mutation) (models.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []models.Hook {
//		return []models.Hook{
//			Reject(models.Delete|models.Update),
//		}
//	}
//
func Reject(op models.Op) models.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []models.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...models.Hook) Chain {
	return Chain{append([]models.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() models.Hook {
	return func(mutator models.Mutator) models.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...models.Hook) Chain {
	newHooks := make([]models.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
