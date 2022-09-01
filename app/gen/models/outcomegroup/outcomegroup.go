// Code generated (@generated) by entc, DO NOT EDIT.

package outcomegroup

const (
	// Label holds the string label denoting the outcomegroup type in the database.
	Label = "outcome_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeGroupID holds the string denoting the outcome_group_id field in the database.
	FieldOutcomeGroupID = "outcome_group_id"
	// FieldOutcomeGroupTitle holds the string denoting the outcome_group_title field in the database.
	FieldOutcomeGroupTitle = "outcome_group_title"
	// FieldOutcomeGroupDescription holds the string denoting the outcome_group_description field in the database.
	FieldOutcomeGroupDescription = "outcome_group_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomegroup in the database.
	Table = "outcome_group"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_group"
	// ParentInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	ParentInverseTable = "outcome_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measure_outcome_group_list"
)

// Columns holds all SQL columns for outcomegroup fields.
var Columns = []string{
	FieldID,
	FieldOutcomeGroupID,
	FieldOutcomeGroupTitle,
	FieldOutcomeGroupDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_group"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measure_outcome_group_list",
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
