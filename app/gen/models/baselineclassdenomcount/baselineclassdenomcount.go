// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclassdenomcount

const (
	// Label holds the string label denoting the baselineclassdenomcount type in the database.
	Label = "baseline_class_denom_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineClassDenomCountGroupID holds the string denoting the baseline_class_denom_count_group_id field in the database.
	FieldBaselineClassDenomCountGroupID = "baseline_class_denom_count_group_id"
	// FieldBaselineClassDenomCountValue holds the string denoting the baseline_class_denom_count_value field in the database.
	FieldBaselineClassDenomCountValue = "baseline_class_denom_count_value"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the baselineclassdenomcount in the database.
	Table = "baseline_class_denom_count"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_class_denom_count"
	// ParentInverseTable is the table name for the BaselineClassDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclassdenom" package.
	ParentInverseTable = "baseline_class_denom"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_class_denom_baseline_class_denom_count_list"
)

// Columns holds all SQL columns for baselineclassdenomcount fields.
var Columns = []string{
	FieldID,
	FieldBaselineClassDenomCountGroupID,
	FieldBaselineClassDenomCountValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_class_denom_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_class_denom_baseline_class_denom_count_list",
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
