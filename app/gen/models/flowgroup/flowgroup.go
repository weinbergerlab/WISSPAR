// Code generated (@generated) by entc, DO NOT EDIT.

package flowgroup

const (
	// Label holds the string label denoting the flowgroup type in the database.
	Label = "flow_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowGroupID holds the string denoting the flow_group_id field in the database.
	FieldFlowGroupID = "flow_group_id"
	// FieldFlowGroupTitle holds the string denoting the flow_group_title field in the database.
	FieldFlowGroupTitle = "flow_group_title"
	// FieldFlowGroupDescription holds the string denoting the flow_group_description field in the database.
	FieldFlowGroupDescription = "flow_group_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the flowgroup in the database.
	Table = "flow_group"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_group"
	// ParentInverseTable is the table name for the ParticipantFlowModule entity.
	// It exists in this package in order to avoid circular dependency with the "participantflowmodule" package.
	ParentInverseTable = "participant_flow_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "participant_flow_module_flow_group_list"
)

// Columns holds all SQL columns for flowgroup fields.
var Columns = []string{
	FieldID,
	FieldFlowGroupID,
	FieldFlowGroupTitle,
	FieldFlowGroupDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_group"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"participant_flow_module_flow_group_list",
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
