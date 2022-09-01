// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclass

const (
	// Label holds the string label denoting the baselineclass type in the database.
	Label = "baseline_class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineClassTitle holds the string denoting the baseline_class_title field in the database.
	FieldBaselineClassTitle = "baseline_class_title"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineClassDenomList holds the string denoting the baseline_class_denom_list edge name in mutations.
	EdgeBaselineClassDenomList = "baseline_class_denom_list"
	// EdgeBaselineCategoryList holds the string denoting the baseline_category_list edge name in mutations.
	EdgeBaselineCategoryList = "baseline_category_list"
	// Table holds the table name of the baselineclass in the database.
	Table = "baseline_class"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_class"
	// ParentInverseTable is the table name for the BaselineMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasure" package.
	ParentInverseTable = "baseline_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_measure_baseline_class_list"
	// BaselineClassDenomListTable is the table that holds the baseline_class_denom_list relation/edge.
	BaselineClassDenomListTable = "baseline_class_denom"
	// BaselineClassDenomListInverseTable is the table name for the BaselineClassDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclassdenom" package.
	BaselineClassDenomListInverseTable = "baseline_class_denom"
	// BaselineClassDenomListColumn is the table column denoting the baseline_class_denom_list relation/edge.
	BaselineClassDenomListColumn = "baseline_class_baseline_class_denom_list"
	// BaselineCategoryListTable is the table that holds the baseline_category_list relation/edge.
	BaselineCategoryListTable = "baseline_category"
	// BaselineCategoryListInverseTable is the table name for the BaselineCategory entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecategory" package.
	BaselineCategoryListInverseTable = "baseline_category"
	// BaselineCategoryListColumn is the table column denoting the baseline_category_list relation/edge.
	BaselineCategoryListColumn = "baseline_class_baseline_category_list"
)

// Columns holds all SQL columns for baselineclass fields.
var Columns = []string{
	FieldID,
	FieldBaselineClassTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_class"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_measure_baseline_class_list",
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
