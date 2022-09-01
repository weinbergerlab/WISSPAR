// Code generated (@generated) by entc, DO NOT EDIT.

package outcomedenomcount

const (
	// Label holds the string label denoting the outcomedenomcount type in the database.
	Label = "outcome_denom_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeDenomCountGroupID holds the string denoting the outcome_denom_count_group_id field in the database.
	FieldOutcomeDenomCountGroupID = "outcome_denom_count_group_id"
	// FieldOutcomeDenomCountValue holds the string denoting the outcome_denom_count_value field in the database.
	FieldOutcomeDenomCountValue = "outcome_denom_count_value"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomedenomcount in the database.
	Table = "outcome_denom_count"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_denom_count"
	// ParentInverseTable is the table name for the OutcomeDenom entity.
	// It exists in this package in order to avoid circular dependency with the "outcomedenom" package.
	ParentInverseTable = "outcome_denom"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_denom_outcome_denom_count_list"
)

// Columns holds all SQL columns for outcomedenomcount fields.
var Columns = []string{
	FieldID,
	FieldOutcomeDenomCountGroupID,
	FieldOutcomeDenomCountValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_denom_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_denom_outcome_denom_count_list",
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
