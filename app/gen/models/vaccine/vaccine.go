// Code generated (@generated) by entc, DO NOT EDIT.

package vaccine

const (
	// Label holds the string label denoting the vaccine type in the database.
	Label = "vaccine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVaccineName holds the string denoting the vaccine_name field in the database.
	FieldVaccineName = "vaccine_name"
	// Table holds the table name of the vaccine in the database.
	Table = "vaccine"
)

// Columns holds all SQL columns for vaccine fields.
var Columns = []string{
	FieldID,
	FieldVaccineName,
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
	// VaccineNameValidator is a validator for the "vaccine_name" field. It is called by the builders before save.
	VaccineNameValidator func(string) error
)
