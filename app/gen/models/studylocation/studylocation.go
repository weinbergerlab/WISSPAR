// Code generated (@generated) by entc, DO NOT EDIT.

package studylocation

const (
	// Label holds the string label denoting the studylocation type in the database.
	Label = "study_location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocationFacility holds the string denoting the locationfacility field in the database.
	FieldLocationFacility = "location_facility"
	// FieldLocationCity holds the string denoting the locationcity field in the database.
	FieldLocationCity = "location_city"
	// FieldLocationCountry holds the string denoting the locationcountry field in the database.
	FieldLocationCountry = "location_country"
	// FieldLocationCountryCode holds the string denoting the locationcountrycode field in the database.
	FieldLocationCountryCode = "location_country_code"
	// FieldLocationContinentName holds the string denoting the locationcontinentname field in the database.
	FieldLocationContinentName = "location_continent_name"
	// FieldLocationContinentCode holds the string denoting the locationcontinentcode field in the database.
	FieldLocationContinentCode = "location_continent_code"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the studylocation in the database.
	Table = "study_location"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "study_location"
	// ParentInverseTable is the table name for the ClinicalTrial entity.
	// It exists in this package in order to avoid circular dependency with the "clinicaltrial" package.
	ParentInverseTable = "clinical_trial"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "clinical_trial_study_locations"
)

// Columns holds all SQL columns for studylocation fields.
var Columns = []string{
	FieldID,
	FieldLocationFacility,
	FieldLocationCity,
	FieldLocationCountry,
	FieldLocationCountryCode,
	FieldLocationContinentName,
	FieldLocationContinentCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "study_location"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"clinical_trial_study_locations",
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
