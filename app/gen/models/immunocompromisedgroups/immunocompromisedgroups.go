// Code generated (@generated) by entc, DO NOT EDIT.

package immunocompromisedgroups

const (
	// Label holds the string label denoting the immunocompromisedgroups type in the database.
	Label = "immunocompromised_groups"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupName holds the string denoting the group_name field in the database.
	FieldGroupName = "group_name"
	// Table holds the table name of the immunocompromisedgroups in the database.
	Table = "immunocompromised_groups"
)

// Columns holds all SQL columns for immunocompromisedgroups fields.
var Columns = []string{
	FieldID,
	FieldGroupName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// GroupNameValidator is a validator for the "group_name" field. It is called by the builders before save.
	GroupNameValidator func(string) error
)
