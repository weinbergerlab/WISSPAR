// Code generated (@generated) by entc, DO NOT EDIT.

package usecase

const (
	// Label holds the string label denoting the usecase type in the database.
	Label = "use_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUseCaseName holds the string denoting the use_case_name field in the database.
	FieldUseCaseName = "use_case_name"
	// FieldUseCaseDescription holds the string denoting the use_case_description field in the database.
	FieldUseCaseDescription = "use_case_description"
	// FieldUseCaseCode holds the string denoting the use_case_code field in the database.
	FieldUseCaseCode = "use_case_code"
	// Table holds the table name of the usecase in the database.
	Table = "use_case"
)

// Columns holds all SQL columns for usecase fields.
var Columns = []string{
	FieldID,
	FieldUseCaseName,
	FieldUseCaseDescription,
	FieldUseCaseCode,
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
