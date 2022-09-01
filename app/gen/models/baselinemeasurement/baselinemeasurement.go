// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasurement

const (
	// Label holds the string label denoting the baselinemeasurement type in the database.
	Label = "baseline_measurement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineMeasurementGroupID holds the string denoting the baseline_measurement_group_id field in the database.
	FieldBaselineMeasurementGroupID = "baseline_measurement_group_id"
	// FieldBaselineMeasurementValue holds the string denoting the baseline_measurement_value field in the database.
	FieldBaselineMeasurementValue = "baseline_measurement_value"
	// FieldBaselineMeasurementSpread holds the string denoting the baseline_measurement_spread field in the database.
	FieldBaselineMeasurementSpread = "baseline_measurement_spread"
	// FieldBaselineMeasurementLowerLimit holds the string denoting the baseline_measurement_lower_limit field in the database.
	FieldBaselineMeasurementLowerLimit = "baseline_measurement_lower_limit"
	// FieldBaselineMeasurementUpperLimit holds the string denoting the baseline_measurement_upper_limit field in the database.
	FieldBaselineMeasurementUpperLimit = "baseline_measurement_upper_limit"
	// FieldBaselineMeasurementComment holds the string denoting the baseline_measurement_comment field in the database.
	FieldBaselineMeasurementComment = "baseline_measurement_comment"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the baselinemeasurement in the database.
	Table = "baseline_measurement"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_measurement"
	// ParentInverseTable is the table name for the BaselineCategory entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecategory" package.
	ParentInverseTable = "baseline_category"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_category_baseline_measurement_list"
)

// Columns holds all SQL columns for baselinemeasurement fields.
var Columns = []string{
	FieldID,
	FieldBaselineMeasurementGroupID,
	FieldBaselineMeasurementValue,
	FieldBaselineMeasurementSpread,
	FieldBaselineMeasurementLowerLimit,
	FieldBaselineMeasurementUpperLimit,
	FieldBaselineMeasurementComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_measurement"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_category_baseline_measurement_list",
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
