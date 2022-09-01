// Code generated (@generated) by entc, DO NOT EDIT.

package baselinedenomcount

const (
	// Label holds the string label denoting the baselinedenomcount type in the database.
	Label = "baseline_denom_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineDenomCountGroupID holds the string denoting the baseline_denom_count_group_id field in the database.
	FieldBaselineDenomCountGroupID = "baseline_denom_count_group_id"
	// FieldBaselineDenomCountValue holds the string denoting the baseline_denom_count_value field in the database.
	FieldBaselineDenomCountValue = "baseline_denom_count_value"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the baselinedenomcount in the database.
	Table = "baseline_denom_count"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_denom_count"
	// ParentInverseTable is the table name for the BaselineDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselinedenom" package.
	ParentInverseTable = "baseline_denom"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_denom_baseline_denom_count_list"
)

// Columns holds all SQL columns for baselinedenomcount fields.
var Columns = []string{
	FieldID,
	FieldBaselineDenomCountGroupID,
	FieldBaselineDenomCountValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_denom_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_denom_baseline_denom_count_list",
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
