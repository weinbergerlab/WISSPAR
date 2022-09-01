// Code generated (@generated) by entc, DO NOT EDIT.

package outcomedenom

const (
	// Label holds the string label denoting the outcomedenom type in the database.
	Label = "outcome_denom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeDenomUnits holds the string denoting the outcome_denom_units field in the database.
	FieldOutcomeDenomUnits = "outcome_denom_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeDenomCountList holds the string denoting the outcome_denom_count_list edge name in mutations.
	EdgeOutcomeDenomCountList = "outcome_denom_count_list"
	// Table holds the table name of the outcomedenom in the database.
	Table = "outcome_denom"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_denom"
	// ParentInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	ParentInverseTable = "outcome_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measure_outcome_denom_list"
	// OutcomeDenomCountListTable is the table that holds the outcome_denom_count_list relation/edge.
	OutcomeDenomCountListTable = "outcome_denom_count"
	// OutcomeDenomCountListInverseTable is the table name for the OutcomeDenomCount entity.
	// It exists in this package in order to avoid circular dependency with the "outcomedenomcount" package.
	OutcomeDenomCountListInverseTable = "outcome_denom_count"
	// OutcomeDenomCountListColumn is the table column denoting the outcome_denom_count_list relation/edge.
	OutcomeDenomCountListColumn = "outcome_denom_outcome_denom_count_list"
)

// Columns holds all SQL columns for outcomedenom fields.
var Columns = []string{
	FieldID,
	FieldOutcomeDenomUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_denom"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measure_outcome_denom_list",
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
