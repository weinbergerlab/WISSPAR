// Code generated (@generated) by entc, DO NOT EDIT.

package manufacturer

const (
	// Label holds the string label denoting the manufacturer type in the database.
	Label = "manufacturer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldManufacturerName holds the string denoting the manufacturer_name field in the database.
	FieldManufacturerName = "manufacturer_name"
	// Table holds the table name of the manufacturer in the database.
	Table = "manufacturer"
)

// Columns holds all SQL columns for manufacturer fields.
var Columns = []string{
	FieldID,
	FieldManufacturerName,
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
	// ManufacturerNameValidator is a validator for the "manufacturer_name" field. It is called by the builders before save.
	ManufacturerNameValidator func(string) error
)
