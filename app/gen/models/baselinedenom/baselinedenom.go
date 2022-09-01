// Code generated (@generated) by entc, DO NOT EDIT.

package baselinedenom

const (
	// Label holds the string label denoting the baselinedenom type in the database.
	Label = "baseline_denom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineDenomUnits holds the string denoting the baseline_denom_units field in the database.
	FieldBaselineDenomUnits = "baseline_denom_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineDenomCountList holds the string denoting the baseline_denom_count_list edge name in mutations.
	EdgeBaselineDenomCountList = "baseline_denom_count_list"
	// Table holds the table name of the baselinedenom in the database.
	Table = "baseline_denom"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_denom"
	// ParentInverseTable is the table name for the BaselineCharacteristicsModule entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecharacteristicsmodule" package.
	ParentInverseTable = "baseline_characteristics_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_characteristics_module_baseline_denom_list"
	// BaselineDenomCountListTable is the table that holds the baseline_denom_count_list relation/edge.
	BaselineDenomCountListTable = "baseline_denom_count"
	// BaselineDenomCountListInverseTable is the table name for the BaselineDenomCount entity.
	// It exists in this package in order to avoid circular dependency with the "baselinedenomcount" package.
	BaselineDenomCountListInverseTable = "baseline_denom_count"
	// BaselineDenomCountListColumn is the table column denoting the baseline_denom_count_list relation/edge.
	BaselineDenomCountListColumn = "baseline_denom_baseline_denom_count_list"
)

// Columns holds all SQL columns for baselinedenom fields.
var Columns = []string{
	FieldID,
	FieldBaselineDenomUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_denom"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_characteristics_module_baseline_denom_list",
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
