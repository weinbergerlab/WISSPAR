// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclassdenom

const (
	// Label holds the string label denoting the baselineclassdenom type in the database.
	Label = "baseline_class_denom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineClassDenomUnits holds the string denoting the baseline_class_denom_units field in the database.
	FieldBaselineClassDenomUnits = "baseline_class_denom_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineClassDenomCountList holds the string denoting the baseline_class_denom_count_list edge name in mutations.
	EdgeBaselineClassDenomCountList = "baseline_class_denom_count_list"
	// Table holds the table name of the baselineclassdenom in the database.
	Table = "baseline_class_denom"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_class_denom"
	// ParentInverseTable is the table name for the BaselineClass entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclass" package.
	ParentInverseTable = "baseline_class"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_class_baseline_class_denom_list"
	// BaselineClassDenomCountListTable is the table that holds the baseline_class_denom_count_list relation/edge.
	BaselineClassDenomCountListTable = "baseline_class_denom_count"
	// BaselineClassDenomCountListInverseTable is the table name for the BaselineClassDenomCount entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclassdenomcount" package.
	BaselineClassDenomCountListInverseTable = "baseline_class_denom_count"
	// BaselineClassDenomCountListColumn is the table column denoting the baseline_class_denom_count_list relation/edge.
	BaselineClassDenomCountListColumn = "baseline_class_denom_baseline_class_denom_count_list"
)

// Columns holds all SQL columns for baselineclassdenom fields.
var Columns = []string{
	FieldID,
	FieldBaselineClassDenomUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_class_denom"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_class_baseline_class_denom_list",
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
