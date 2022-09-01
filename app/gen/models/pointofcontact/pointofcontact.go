// Code generated (@generated) by entc, DO NOT EDIT.

package pointofcontact

const (
	// Label holds the string label denoting the pointofcontact type in the database.
	Label = "point_of_contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPointOfContactTitle holds the string denoting the point_of_contact_title field in the database.
	FieldPointOfContactTitle = "point_of_contact_title"
	// FieldPointOfContactOrganization holds the string denoting the point_of_contact_organization field in the database.
	FieldPointOfContactOrganization = "point_of_contact_organization"
	// FieldPointOfContactEmail holds the string denoting the point_of_contact_email field in the database.
	FieldPointOfContactEmail = "point_of_contact_email"
	// FieldPointOfContactPhone holds the string denoting the point_of_contact_phone field in the database.
	FieldPointOfContactPhone = "point_of_contact_phone"
	// FieldPointOfContactPhoneExt holds the string denoting the point_of_contact_phone_ext field in the database.
	FieldPointOfContactPhoneExt = "point_of_contact_phone_ext"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the pointofcontact in the database.
	Table = "point_of_contact"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "point_of_contact"
	// ParentInverseTable is the table name for the MoreInfoModule entity.
	// It exists in this package in order to avoid circular dependency with the "moreinfomodule" package.
	ParentInverseTable = "more_info_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "point_of_contact_parent"
)

// Columns holds all SQL columns for pointofcontact fields.
var Columns = []string{
	FieldID,
	FieldPointOfContactTitle,
	FieldPointOfContactOrganization,
	FieldPointOfContactEmail,
	FieldPointOfContactPhone,
	FieldPointOfContactPhoneExt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "point_of_contact"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"point_of_contact_parent",
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
