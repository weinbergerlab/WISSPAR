// Code generated (@generated) by entc, DO NOT EDIT.

package baselinecategory

const (
	// Label holds the string label denoting the baselinecategory type in the database.
	Label = "baseline_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineCategoryTitle holds the string denoting the baseline_category_title field in the database.
	FieldBaselineCategoryTitle = "baseline_category_title"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineMeasurementList holds the string denoting the baseline_measurement_list edge name in mutations.
	EdgeBaselineMeasurementList = "baseline_measurement_list"
	// Table holds the table name of the baselinecategory in the database.
	Table = "baseline_category"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_category"
	// ParentInverseTable is the table name for the BaselineClass entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclass" package.
	ParentInverseTable = "baseline_class"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_class_baseline_category_list"
	// BaselineMeasurementListTable is the table that holds the baseline_measurement_list relation/edge.
	BaselineMeasurementListTable = "baseline_measurement"
	// BaselineMeasurementListInverseTable is the table name for the BaselineMeasurement entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasurement" package.
	BaselineMeasurementListInverseTable = "baseline_measurement"
	// BaselineMeasurementListColumn is the table column denoting the baseline_measurement_list relation/edge.
	BaselineMeasurementListColumn = "baseline_category_baseline_measurement_list"
)

// Columns holds all SQL columns for baselinecategory fields.
var Columns = []string{
	FieldID,
	FieldBaselineCategoryTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_category"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_class_baseline_category_list",
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
