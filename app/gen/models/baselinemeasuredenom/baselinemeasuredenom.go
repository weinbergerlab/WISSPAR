// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasuredenom

const (
	// Label holds the string label denoting the baselinemeasuredenom type in the database.
	Label = "baseline_measure_denom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineMeasureDenomUnits holds the string denoting the baseline_measure_denom_units field in the database.
	FieldBaselineMeasureDenomUnits = "baseline_measure_denom_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineMeasureDenomCountList holds the string denoting the baseline_measure_denom_count_list edge name in mutations.
	EdgeBaselineMeasureDenomCountList = "baseline_measure_denom_count_list"
	// Table holds the table name of the baselinemeasuredenom in the database.
	Table = "baseline_measure_denom"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_measure_denom"
	// ParentInverseTable is the table name for the BaselineMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasure" package.
	ParentInverseTable = "baseline_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_measure_baseline_measure_denom_list"
	// BaselineMeasureDenomCountListTable is the table that holds the baseline_measure_denom_count_list relation/edge.
	BaselineMeasureDenomCountListTable = "baseline_measure_denom_count"
	// BaselineMeasureDenomCountListInverseTable is the table name for the BaselineMeasureDenomCount entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasuredenomcount" package.
	BaselineMeasureDenomCountListInverseTable = "baseline_measure_denom_count"
	// BaselineMeasureDenomCountListColumn is the table column denoting the baseline_measure_denom_count_list relation/edge.
	BaselineMeasureDenomCountListColumn = "baseline_measure_denom_baseline_measure_denom_count_list"
)

// Columns holds all SQL columns for baselinemeasuredenom fields.
var Columns = []string{
	FieldID,
	FieldBaselineMeasureDenomUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_measure_denom"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_measure_baseline_measure_denom_list",
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
