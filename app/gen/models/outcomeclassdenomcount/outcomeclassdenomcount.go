// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclassdenomcount

const (
	// Label holds the string label denoting the outcomeclassdenomcount type in the database.
	Label = "outcome_class_denom_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeClassDenomCountGroupID holds the string denoting the outcome_class_denom_count_group_id field in the database.
	FieldOutcomeClassDenomCountGroupID = "outcome_class_denom_count_group_id"
	// FieldOutcomeClassDenomCountValue holds the string denoting the outcome_class_denom_count_value field in the database.
	FieldOutcomeClassDenomCountValue = "outcome_class_denom_count_value"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomeclassdenomcount in the database.
	Table = "outcome_class_denom_count"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_class_denom_count"
	// ParentInverseTable is the table name for the OutcomeClassDenom entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclassdenom" package.
	ParentInverseTable = "outcome_class_denom"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_class_denom_outcome_class_denom_count_list"
)

// Columns holds all SQL columns for outcomeclassdenomcount fields.
var Columns = []string{
	FieldID,
	FieldOutcomeClassDenomCountGroupID,
	FieldOutcomeClassDenomCountValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_class_denom_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_class_denom_outcome_class_denom_count_list",
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
