// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasuresmodule

const (
	// Label holds the string label denoting the outcomemeasuresmodule type in the database.
	Label = "outcome_measures_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeMeasureList holds the string denoting the outcome_measure_list edge name in mutations.
	EdgeOutcomeMeasureList = "outcome_measure_list"
	// Table holds the table name of the outcomemeasuresmodule in the database.
	Table = "outcome_measures_module"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_measures_module"
	// ParentInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ParentInverseTable = "results_definition"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "results_definition_outcome_measures_module"
	// OutcomeMeasureListTable is the table that holds the outcome_measure_list relation/edge.
	OutcomeMeasureListTable = "outcome_measure"
	// OutcomeMeasureListInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	OutcomeMeasureListInverseTable = "outcome_measure"
	// OutcomeMeasureListColumn is the table column denoting the outcome_measure_list relation/edge.
	OutcomeMeasureListColumn = "outcome_measures_module_outcome_measure_list"
)

// Columns holds all SQL columns for outcomemeasuresmodule fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_measures_module"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"results_definition_outcome_measures_module",
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
