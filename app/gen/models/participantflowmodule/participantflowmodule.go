// Code generated (@generated) by entc, DO NOT EDIT.

package participantflowmodule

const (
	// Label holds the string label denoting the participantflowmodule type in the database.
	Label = "participant_flow_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowPreAssignmentDetails holds the string denoting the flow_pre_assignment_details field in the database.
	FieldFlowPreAssignmentDetails = "flow_pre_assignment_details"
	// FieldFlowRecruitmentDetails holds the string denoting the flow_recruitment_details field in the database.
	FieldFlowRecruitmentDetails = "flow_recruitment_details"
	// FieldFlowTypeUnitsAnalyzed holds the string denoting the flow_type_units_analyzed field in the database.
	FieldFlowTypeUnitsAnalyzed = "flow_type_units_analyzed"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeFlowGroupList holds the string denoting the flow_group_list edge name in mutations.
	EdgeFlowGroupList = "flow_group_list"
	// EdgeFlowPeriodList holds the string denoting the flow_period_list edge name in mutations.
	EdgeFlowPeriodList = "flow_period_list"
	// Table holds the table name of the participantflowmodule in the database.
	Table = "participant_flow_module"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "participant_flow_module"
	// ParentInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ParentInverseTable = "results_definition"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "results_definition_participant_flow_module"
	// FlowGroupListTable is the table that holds the flow_group_list relation/edge.
	FlowGroupListTable = "flow_group"
	// FlowGroupListInverseTable is the table name for the FlowGroup entity.
	// It exists in this package in order to avoid circular dependency with the "flowgroup" package.
	FlowGroupListInverseTable = "flow_group"
	// FlowGroupListColumn is the table column denoting the flow_group_list relation/edge.
	FlowGroupListColumn = "participant_flow_module_flow_group_list"
	// FlowPeriodListTable is the table that holds the flow_period_list relation/edge.
	FlowPeriodListTable = "flow_period"
	// FlowPeriodListInverseTable is the table name for the FlowPeriod entity.
	// It exists in this package in order to avoid circular dependency with the "flowperiod" package.
	FlowPeriodListInverseTable = "flow_period"
	// FlowPeriodListColumn is the table column denoting the flow_period_list relation/edge.
	FlowPeriodListColumn = "participant_flow_module_flow_period_list"
)

// Columns holds all SQL columns for participantflowmodule fields.
var Columns = []string{
	FieldID,
	FieldFlowPreAssignmentDetails,
	FieldFlowRecruitmentDetails,
	FieldFlowTypeUnitsAnalyzed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "participant_flow_module"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"results_definition_participant_flow_module",
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
