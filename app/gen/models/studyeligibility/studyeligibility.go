// Code generated (@generated) by entc, DO NOT EDIT.

package studyeligibility

const (
	// Label holds the string label denoting the studyeligibility type in the database.
	Label = "study_eligibility"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEligibilityCriteria holds the string denoting the eligibilitycriteria field in the database.
	FieldEligibilityCriteria = "eligibility_criteria"
	// FieldHealthyVolunteers holds the string denoting the healthyvolunteers field in the database.
	FieldHealthyVolunteers = "healthy_volunteers"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldMinimumAge holds the string denoting the minimumage field in the database.
	FieldMinimumAge = "minimum_age"
	// FieldMaximumAge holds the string denoting the maximumage field in the database.
	FieldMaximumAge = "maximum_age"
	// FieldStdAgeList holds the string denoting the stdagelist field in the database.
	FieldStdAgeList = "std_age_list"
	// FieldEthnicity holds the string denoting the ethnicity field in the database.
	FieldEthnicity = "ethnicity"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the studyeligibility in the database.
	Table = "study_eligibility"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "study_eligibility"
	// ParentInverseTable is the table name for the ClinicalTrial entity.
	// It exists in this package in order to avoid circular dependency with the "clinicaltrial" package.
	ParentInverseTable = "clinical_trial"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "clinical_trial_study_eligibility"
)

// Columns holds all SQL columns for studyeligibility fields.
var Columns = []string{
	FieldID,
	FieldEligibilityCriteria,
	FieldHealthyVolunteers,
	FieldGender,
	FieldMinimumAge,
	FieldMaximumAge,
	FieldStdAgeList,
	FieldEthnicity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "study_eligibility"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"clinical_trial_study_eligibility",
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
