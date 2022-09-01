// Code generated (@generated) by entc, DO NOT EDIT.

package baselinegroup

const (
	// Label holds the string label denoting the baselinegroup type in the database.
	Label = "baseline_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineGroupID holds the string denoting the baseline_group_id field in the database.
	FieldBaselineGroupID = "baseline_group_id"
	// FieldBaselineGroupTitle holds the string denoting the baseline_group_title field in the database.
	FieldBaselineGroupTitle = "baseline_group_title"
	// FieldBaselineGroupDescription holds the string denoting the baseline_group_description field in the database.
	FieldBaselineGroupDescription = "baseline_group_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the baselinegroup in the database.
	Table = "baseline_group"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_group"
	// ParentInverseTable is the table name for the BaselineCharacteristicsModule entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecharacteristicsmodule" package.
	ParentInverseTable = "baseline_characteristics_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_characteristics_module_baseline_group_list"
)

// Columns holds all SQL columns for baselinegroup fields.
var Columns = []string{
	FieldID,
	FieldBaselineGroupID,
	FieldBaselineGroupTitle,
	FieldBaselineGroupDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_group"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_characteristics_module_baseline_group_list",
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
