// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclassdenom

const (
	// Label holds the string label denoting the outcomeclassdenom type in the database.
	Label = "outcome_class_denom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeClassDenomUnits holds the string denoting the outcome_class_denom_units field in the database.
	FieldOutcomeClassDenomUnits = "outcome_class_denom_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeClassDenomCountList holds the string denoting the outcome_class_denom_count_list edge name in mutations.
	EdgeOutcomeClassDenomCountList = "outcome_class_denom_count_list"
	// Table holds the table name of the outcomeclassdenom in the database.
	Table = "outcome_class_denom"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_class_denom"
	// ParentInverseTable is the table name for the OutcomeClass entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclass" package.
	ParentInverseTable = "outcome_class"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_class_outcome_class_denom_list"
	// OutcomeClassDenomCountListTable is the table that holds the outcome_class_denom_count_list relation/edge.
	OutcomeClassDenomCountListTable = "outcome_class_denom_count"
	// OutcomeClassDenomCountListInverseTable is the table name for the OutcomeClassDenomCount entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclassdenomcount" package.
	OutcomeClassDenomCountListInverseTable = "outcome_class_denom_count"
	// OutcomeClassDenomCountListColumn is the table column denoting the outcome_class_denom_count_list relation/edge.
	OutcomeClassDenomCountListColumn = "outcome_class_denom_outcome_class_denom_count_list"
)

// Columns holds all SQL columns for outcomeclassdenom fields.
var Columns = []string{
	FieldID,
	FieldOutcomeClassDenomUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_class_denom"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_class_outcome_class_denom_list",
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
