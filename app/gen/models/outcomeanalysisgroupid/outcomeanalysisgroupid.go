// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeanalysisgroupid

const (
	// Label holds the string label denoting the outcomeanalysisgroupid type in the database.
	Label = "outcome_analysis_group_id"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeAnalysisGroupID holds the string denoting the outcome_analysis_group_id field in the database.
	FieldOutcomeAnalysisGroupID = "outcome_analysis_group_id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomeanalysisgroupid in the database.
	Table = "outcome_analysis_group_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_analysis_group_id"
	// ParentInverseTable is the table name for the OutcomeAnalysis entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeanalysis" package.
	ParentInverseTable = "outcome_analysis"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_analysis_outcome_analysis_group_id_list"
)

// Columns holds all SQL columns for outcomeanalysisgroupid fields.
var Columns = []string{
	FieldID,
	FieldOutcomeAnalysisGroupID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_analysis_group_id"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_analysis_outcome_analysis_group_id_list",
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
