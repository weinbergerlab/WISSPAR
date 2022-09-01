// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasuredenomcount

const (
	// Label holds the string label denoting the baselinemeasuredenomcount type in the database.
	Label = "baseline_measure_denom_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineMeasureDenomCountGroupID holds the string denoting the baseline_measure_denom_count_group_id field in the database.
	FieldBaselineMeasureDenomCountGroupID = "baseline_measure_denom_count_group_id"
	// FieldBaselineMeasureDenomCountValue holds the string denoting the baseline_measure_denom_count_value field in the database.
	FieldBaselineMeasureDenomCountValue = "baseline_measure_denom_count_value"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the baselinemeasuredenomcount in the database.
	Table = "baseline_measure_denom_count"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_measure_denom_count"
	// ParentInverseTable is the table name for the BaselineMeasureDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasuredenom" package.
	ParentInverseTable = "baseline_measure_denom"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_measure_denom_baseline_measure_denom_count_list"
)

// Columns holds all SQL columns for baselinemeasuredenomcount fields.
var Columns = []string{
	FieldID,
	FieldBaselineMeasureDenomCountGroupID,
	FieldBaselineMeasureDenomCountValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_measure_denom_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_measure_denom_baseline_measure_denom_count_list",
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
