// Code generated (@generated) by entc, DO NOT EDIT.

package outcomecategory

const (
	// Label holds the string label denoting the outcomecategory type in the database.
	Label = "outcome_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeCategoryTitle holds the string denoting the outcome_category_title field in the database.
	FieldOutcomeCategoryTitle = "outcome_category_title"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeMeasurementList holds the string denoting the outcome_measurement_list edge name in mutations.
	EdgeOutcomeMeasurementList = "outcome_measurement_list"
	// Table holds the table name of the outcomecategory in the database.
	Table = "outcome_category"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_category"
	// ParentInverseTable is the table name for the OutcomeClass entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclass" package.
	ParentInverseTable = "outcome_class"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_class_outcome_category_list"
	// OutcomeMeasurementListTable is the table that holds the outcome_measurement_list relation/edge.
	OutcomeMeasurementListTable = "outcome_measurement"
	// OutcomeMeasurementListInverseTable is the table name for the OutcomeMeasurement entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasurement" package.
	OutcomeMeasurementListInverseTable = "outcome_measurement"
	// OutcomeMeasurementListColumn is the table column denoting the outcome_measurement_list relation/edge.
	OutcomeMeasurementListColumn = "outcome_category_outcome_measurement_list"
)

// Columns holds all SQL columns for outcomecategory fields.
var Columns = []string{
	FieldID,
	FieldOutcomeCategoryTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_category"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_class_outcome_category_list",
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
