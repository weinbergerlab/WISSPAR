// Code generated (@generated) by entc, DO NOT EDIT.

package resultsdefinition

const (
	// Label holds the string label denoting the resultsdefinition type in the database.
	Label = "results_definition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeParticipantFlowModule holds the string denoting the participant_flow_module edge name in mutations.
	EdgeParticipantFlowModule = "participant_flow_module"
	// EdgeBaselineCharacteristicsModule holds the string denoting the baseline_characteristics_module edge name in mutations.
	EdgeBaselineCharacteristicsModule = "baseline_characteristics_module"
	// EdgeOutcomeMeasuresModule holds the string denoting the outcome_measures_module edge name in mutations.
	EdgeOutcomeMeasuresModule = "outcome_measures_module"
	// EdgeAdverseEventsModule holds the string denoting the adverse_events_module edge name in mutations.
	EdgeAdverseEventsModule = "adverse_events_module"
	// EdgeMoreInfoModule holds the string denoting the more_info_module edge name in mutations.
	EdgeMoreInfoModule = "more_info_module"
	// Table holds the table name of the resultsdefinition in the database.
	Table = "results_definition"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "results_definition"
	// ParentInverseTable is the table name for the ClinicalTrial entity.
	// It exists in this package in order to avoid circular dependency with the "clinicaltrial" package.
	ParentInverseTable = "clinical_trial"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "clinical_trial_results_definition"
	// ParticipantFlowModuleTable is the table that holds the participant_flow_module relation/edge.
	ParticipantFlowModuleTable = "participant_flow_module"
	// ParticipantFlowModuleInverseTable is the table name for the ParticipantFlowModule entity.
	// It exists in this package in order to avoid circular dependency with the "participantflowmodule" package.
	ParticipantFlowModuleInverseTable = "participant_flow_module"
	// ParticipantFlowModuleColumn is the table column denoting the participant_flow_module relation/edge.
	ParticipantFlowModuleColumn = "results_definition_participant_flow_module"
	// BaselineCharacteristicsModuleTable is the table that holds the baseline_characteristics_module relation/edge.
	BaselineCharacteristicsModuleTable = "baseline_characteristics_module"
	// BaselineCharacteristicsModuleInverseTable is the table name for the BaselineCharacteristicsModule entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecharacteristicsmodule" package.
	BaselineCharacteristicsModuleInverseTable = "baseline_characteristics_module"
	// BaselineCharacteristicsModuleColumn is the table column denoting the baseline_characteristics_module relation/edge.
	BaselineCharacteristicsModuleColumn = "results_definition_baseline_characteristics_module"
	// OutcomeMeasuresModuleTable is the table that holds the outcome_measures_module relation/edge.
	OutcomeMeasuresModuleTable = "outcome_measures_module"
	// OutcomeMeasuresModuleInverseTable is the table name for the OutcomeMeasuresModule entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasuresmodule" package.
	OutcomeMeasuresModuleInverseTable = "outcome_measures_module"
	// OutcomeMeasuresModuleColumn is the table column denoting the outcome_measures_module relation/edge.
	OutcomeMeasuresModuleColumn = "results_definition_outcome_measures_module"
	// AdverseEventsModuleTable is the table that holds the adverse_events_module relation/edge.
	AdverseEventsModuleTable = "adverse_events_module"
	// AdverseEventsModuleInverseTable is the table name for the AdverseEventsModule entity.
	// It exists in this package in order to avoid circular dependency with the "adverseeventsmodule" package.
	AdverseEventsModuleInverseTable = "adverse_events_module"
	// AdverseEventsModuleColumn is the table column denoting the adverse_events_module relation/edge.
	AdverseEventsModuleColumn = "results_definition_adverse_events_module"
	// MoreInfoModuleTable is the table that holds the more_info_module relation/edge.
	MoreInfoModuleTable = "more_info_module"
	// MoreInfoModuleInverseTable is the table name for the MoreInfoModule entity.
	// It exists in this package in order to avoid circular dependency with the "moreinfomodule" package.
	MoreInfoModuleInverseTable = "more_info_module"
	// MoreInfoModuleColumn is the table column denoting the more_info_module relation/edge.
	MoreInfoModuleColumn = "results_definition_more_info_module"
)

// Columns holds all SQL columns for resultsdefinition fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "results_definition"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"clinical_trial_results_definition",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
