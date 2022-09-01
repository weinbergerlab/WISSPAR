// Code generated (@generated) by entc, DO NOT EDIT.

package flowperiod

const (
	// Label holds the string label denoting the flowperiod type in the database.
	Label = "flow_period"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowPeriodTitle holds the string denoting the flow_period_title field in the database.
	FieldFlowPeriodTitle = "flow_period_title"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeFlowMilestoneList holds the string denoting the flow_milestone_list edge name in mutations.
	EdgeFlowMilestoneList = "flow_milestone_list"
	// EdgeFlowDropWithdrawList holds the string denoting the flow_drop_withdraw_list edge name in mutations.
	EdgeFlowDropWithdrawList = "flow_drop_withdraw_list"
	// Table holds the table name of the flowperiod in the database.
	Table = "flow_period"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_period"
	// ParentInverseTable is the table name for the ParticipantFlowModule entity.
	// It exists in this package in order to avoid circular dependency with the "participantflowmodule" package.
	ParentInverseTable = "participant_flow_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "participant_flow_module_flow_period_list"
	// FlowMilestoneListTable is the table that holds the flow_milestone_list relation/edge.
	FlowMilestoneListTable = "flow_milestone"
	// FlowMilestoneListInverseTable is the table name for the FlowMilestone entity.
	// It exists in this package in order to avoid circular dependency with the "flowmilestone" package.
	FlowMilestoneListInverseTable = "flow_milestone"
	// FlowMilestoneListColumn is the table column denoting the flow_milestone_list relation/edge.
	FlowMilestoneListColumn = "flow_period_flow_milestone_list"
	// FlowDropWithdrawListTable is the table that holds the flow_drop_withdraw_list relation/edge.
	FlowDropWithdrawListTable = "flow_drop_withdraw"
	// FlowDropWithdrawListInverseTable is the table name for the FlowDropWithdraw entity.
	// It exists in this package in order to avoid circular dependency with the "flowdropwithdraw" package.
	FlowDropWithdrawListInverseTable = "flow_drop_withdraw"
	// FlowDropWithdrawListColumn is the table column denoting the flow_drop_withdraw_list relation/edge.
	FlowDropWithdrawListColumn = "flow_period_flow_drop_withdraw_list"
)

// Columns holds all SQL columns for flowperiod fields.
var Columns = []string{
	FieldID,
	FieldFlowPeriodTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_period"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"participant_flow_module_flow_period_list",
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
