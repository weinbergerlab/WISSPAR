// Code generated (@generated) by entc, DO NOT EDIT.

package flowreason

const (
	// Label holds the string label denoting the flowreason type in the database.
	Label = "flow_reason"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowReasonGroupID holds the string denoting the flow_reason_group_id field in the database.
	FieldFlowReasonGroupID = "flow_reason_group_id"
	// FieldFlowReasonComment holds the string denoting the flow_reason_comment field in the database.
	FieldFlowReasonComment = "flow_reason_comment"
	// FieldFlowReasonNumSubjects holds the string denoting the flow_reason_num_subjects field in the database.
	FieldFlowReasonNumSubjects = "flow_reason_num_subjects"
	// FieldFlowReasonNumUnits holds the string denoting the flow_reason_num_units field in the database.
	FieldFlowReasonNumUnits = "flow_reason_num_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the flowreason in the database.
	Table = "flow_reason"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_reason"
	// ParentInverseTable is the table name for the FlowDropWithdraw entity.
	// It exists in this package in order to avoid circular dependency with the "flowdropwithdraw" package.
	ParentInverseTable = "flow_drop_withdraw"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "flow_drop_withdraw_flow_reason_list"
)

// Columns holds all SQL columns for flowreason fields.
var Columns = []string{
	FieldID,
	FieldFlowReasonGroupID,
	FieldFlowReasonComment,
	FieldFlowReasonNumSubjects,
	FieldFlowReasonNumUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_reason"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"flow_drop_withdraw_flow_reason_list",
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
