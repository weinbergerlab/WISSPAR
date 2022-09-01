// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclass

const (
	// Label holds the string label denoting the outcomeclass type in the database.
	Label = "outcome_class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeClassTitle holds the string denoting the outcome_class_title field in the database.
	FieldOutcomeClassTitle = "outcome_class_title"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeClassDenomList holds the string denoting the outcome_class_denom_list edge name in mutations.
	EdgeOutcomeClassDenomList = "outcome_class_denom_list"
	// EdgeOutcomeCategoryList holds the string denoting the outcome_category_list edge name in mutations.
	EdgeOutcomeCategoryList = "outcome_category_list"
	// Table holds the table name of the outcomeclass in the database.
	Table = "outcome_class"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_class"
	// ParentInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	ParentInverseTable = "outcome_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measure_outcome_class_list"
	// OutcomeClassDenomListTable is the table that holds the outcome_class_denom_list relation/edge.
	OutcomeClassDenomListTable = "outcome_class_denom"
	// OutcomeClassDenomListInverseTable is the table name for the OutcomeClassDenom entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclassdenom" package.
	OutcomeClassDenomListInverseTable = "outcome_class_denom"
	// OutcomeClassDenomListColumn is the table column denoting the outcome_class_denom_list relation/edge.
	OutcomeClassDenomListColumn = "outcome_class_outcome_class_denom_list"
	// OutcomeCategoryListTable is the table that holds the outcome_category_list relation/edge.
	OutcomeCategoryListTable = "outcome_category"
	// OutcomeCategoryListInverseTable is the table name for the OutcomeCategory entity.
	// It exists in this package in order to avoid circular dependency with the "outcomecategory" package.
	OutcomeCategoryListInverseTable = "outcome_category"
	// OutcomeCategoryListColumn is the table column denoting the outcome_category_list relation/edge.
	OutcomeCategoryListColumn = "outcome_class_outcome_category_list"
)

// Columns holds all SQL columns for outcomeclass fields.
var Columns = []string{
	FieldID,
	FieldOutcomeClassTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_class"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measure_outcome_class_list",
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
