// Code generated (@generated) by entc, DO NOT EDIT.

package moreinfomodule

const (
	// Label holds the string label denoting the moreinfomodule type in the database.
	Label = "more_info_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLimitationsAndCaveatsDescription holds the string denoting the limitations_and_caveats_description field in the database.
	FieldLimitationsAndCaveatsDescription = "limitations_and_caveats_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeCertainAgreement holds the string denoting the certain_agreement edge name in mutations.
	EdgeCertainAgreement = "certain_agreement"
	// EdgePointOfContact holds the string denoting the point_of_contact edge name in mutations.
	EdgePointOfContact = "point_of_contact"
	// Table holds the table name of the moreinfomodule in the database.
	Table = "more_info_module"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "more_info_module"
	// ParentInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ParentInverseTable = "results_definition"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "results_definition_more_info_module"
	// CertainAgreementTable is the table that holds the certain_agreement relation/edge.
	CertainAgreementTable = "more_info_module"
	// CertainAgreementInverseTable is the table name for the CertainAgreement entity.
	// It exists in this package in order to avoid circular dependency with the "certainagreement" package.
	CertainAgreementInverseTable = "certain_agreement"
	// CertainAgreementColumn is the table column denoting the certain_agreement relation/edge.
	CertainAgreementColumn = "more_info_module_certain_agreement"
	// PointOfContactTable is the table that holds the point_of_contact relation/edge.
	PointOfContactTable = "more_info_module"
	// PointOfContactInverseTable is the table name for the PointOfContact entity.
	// It exists in this package in order to avoid circular dependency with the "pointofcontact" package.
	PointOfContactInverseTable = "point_of_contact"
	// PointOfContactColumn is the table column denoting the point_of_contact relation/edge.
	PointOfContactColumn = "more_info_module_point_of_contact"
)

// Columns holds all SQL columns for moreinfomodule fields.
var Columns = []string{
	FieldID,
	FieldLimitationsAndCaveatsDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "more_info_module"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"more_info_module_certain_agreement",
	"more_info_module_point_of_contact",
	"results_definition_more_info_module",
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
