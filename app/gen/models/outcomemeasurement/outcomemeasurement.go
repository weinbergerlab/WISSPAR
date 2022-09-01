// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasurement

const (
	// Label holds the string label denoting the outcomemeasurement type in the database.
	Label = "outcome_measurement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeMeasurementGroupID holds the string denoting the outcome_measurement_group_id field in the database.
	FieldOutcomeMeasurementGroupID = "outcome_measurement_group_id"
	// FieldOutcomeMeasurementValue holds the string denoting the outcome_measurement_value field in the database.
	FieldOutcomeMeasurementValue = "outcome_measurement_value"
	// FieldOutcomeMeasurementSpread holds the string denoting the outcome_measurement_spread field in the database.
	FieldOutcomeMeasurementSpread = "outcome_measurement_spread"
	// FieldOutcomeMeasurementLowerLimit holds the string denoting the outcome_measurement_lower_limit field in the database.
	FieldOutcomeMeasurementLowerLimit = "outcome_measurement_lower_limit"
	// FieldOutcomeMeasurementUpperLimit holds the string denoting the outcome_measurement_upper_limit field in the database.
	FieldOutcomeMeasurementUpperLimit = "outcome_measurement_upper_limit"
	// FieldOutcomeMeasurementComment holds the string denoting the outcome_measurement_comment field in the database.
	FieldOutcomeMeasurementComment = "outcome_measurement_comment"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomemeasurement in the database.
	Table = "outcome_measurement"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_measurement"
	// ParentInverseTable is the table name for the OutcomeCategory entity.
	// It exists in this package in order to avoid circular dependency with the "outcomecategory" package.
	ParentInverseTable = "outcome_category"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_category_outcome_measurement_list"
)

// Columns holds all SQL columns for outcomemeasurement fields.
var Columns = []string{
	FieldID,
	FieldOutcomeMeasurementGroupID,
	FieldOutcomeMeasurementValue,
	FieldOutcomeMeasurementSpread,
	FieldOutcomeMeasurementLowerLimit,
	FieldOutcomeMeasurementUpperLimit,
	FieldOutcomeMeasurementComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_measurement"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_category_outcome_measurement_list",
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
