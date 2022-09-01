// Code generated (@generated) by entc, DO NOT EDIT.

package clinicaltrialmetadata

const (
	// Label holds the string label denoting the clinicaltrialmetadata type in the database.
	Label = "clinical_trial_metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTagName holds the string denoting the tag_name field in the database.
	FieldTagName = "tag_name"
	// FieldTagValue holds the string denoting the tag_value field in the database.
	FieldTagValue = "tag_value"
	// FieldUseCaseCode holds the string denoting the use_case_code field in the database.
	FieldUseCaseCode = "use_case_code"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the clinicaltrialmetadata in the database.
	Table = "clinical_trial_metadata"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "clinical_trial_metadata"
	// ParentInverseTable is the table name for the ClinicalTrial entity.
	// It exists in this package in order to avoid circular dependency with the "clinicaltrial" package.
	ParentInverseTable = "clinical_trial"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "clinical_trial_metadata"
)

// Columns holds all SQL columns for clinicaltrialmetadata fields.
var Columns = []string{
	FieldID,
	FieldTagName,
	FieldTagValue,
	FieldUseCaseCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "clinical_trial_metadata"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"clinical_trial_metadata",
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

var (
	// DefaultUseCaseCode holds the default value on creation for the "use_case_code" field.
	DefaultUseCaseCode string
)
