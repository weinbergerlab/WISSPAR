// Code generated (@generated) by entc, DO NOT EDIT.

package flowdropwithdraw

const (
	// Label holds the string label denoting the flowdropwithdraw type in the database.
	Label = "flow_drop_withdraw"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowDropWithdrawType holds the string denoting the flow_drop_withdraw_type field in the database.
	FieldFlowDropWithdrawType = "flow_drop_withdraw_type"
	// FieldFlowDropWithdrawComment holds the string denoting the flow_drop_withdraw_comment field in the database.
	FieldFlowDropWithdrawComment = "flow_drop_withdraw_comment"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeFlowReasonList holds the string denoting the flow_reason_list edge name in mutations.
	EdgeFlowReasonList = "flow_reason_list"
	// Table holds the table name of the flowdropwithdraw in the database.
	Table = "flow_drop_withdraw"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_drop_withdraw"
	// ParentInverseTable is the table name for the FlowPeriod entity.
	// It exists in this package in order to avoid circular dependency with the "flowperiod" package.
	ParentInverseTable = "flow_period"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "flow_period_flow_drop_withdraw_list"
	// FlowReasonListTable is the table that holds the flow_reason_list relation/edge.
	FlowReasonListTable = "flow_reason"
	// FlowReasonListInverseTable is the table name for the FlowReason entity.
	// It exists in this package in order to avoid circular dependency with the "flowreason" package.
	FlowReasonListInverseTable = "flow_reason"
	// FlowReasonListColumn is the table column denoting the flow_reason_list relation/edge.
	FlowReasonListColumn = "flow_drop_withdraw_flow_reason_list"
)

// Columns holds all SQL columns for flowdropwithdraw fields.
var Columns = []string{
	FieldID,
	FieldFlowDropWithdrawType,
	FieldFlowDropWithdrawComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_drop_withdraw"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"flow_period_flow_drop_withdraw_list",
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
