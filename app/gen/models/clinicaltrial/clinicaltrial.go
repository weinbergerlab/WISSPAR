// Code generated (@generated) by entc, DO NOT EDIT.

package clinicaltrial

const (
	// Label holds the string label denoting the clinicaltrial type in the database.
	Label = "clinical_trial"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStudyName holds the string denoting the study_name field in the database.
	FieldStudyName = "study_name"
	// FieldSponsor holds the string denoting the sponsor field in the database.
	FieldSponsor = "sponsor"
	// FieldResponsibleParty holds the string denoting the responsible_party field in the database.
	FieldResponsibleParty = "responsible_party"
	// FieldStudyType holds the string denoting the study_type field in the database.
	FieldStudyType = "study_type"
	// FieldPhase holds the string denoting the phase field in the database.
	FieldPhase = "phase"
	// FieldActualEnrollment holds the string denoting the actual_enrollment field in the database.
	FieldActualEnrollment = "actual_enrollment"
	// FieldAllocation holds the string denoting the allocation field in the database.
	FieldAllocation = "allocation"
	// FieldInterventionModel holds the string denoting the intervention_model field in the database.
	FieldInterventionModel = "intervention_model"
	// FieldMasking holds the string denoting the masking field in the database.
	FieldMasking = "masking"
	// FieldPrimaryPurpose holds the string denoting the primary_purpose field in the database.
	FieldPrimaryPurpose = "primary_purpose"
	// FieldOfficialTitle holds the string denoting the official_title field in the database.
	FieldOfficialTitle = "official_title"
	// FieldActualStudyStartDate holds the string denoting the actual_study_start_date field in the database.
	FieldActualStudyStartDate = "actual_study_start_date"
	// FieldActualPrimaryCompletionDate holds the string denoting the actual_primary_completion_date field in the database.
	FieldActualPrimaryCompletionDate = "actual_primary_completion_date"
	// FieldActualStudyCompletionDate holds the string denoting the actual_study_completion_date field in the database.
	FieldActualStudyCompletionDate = "actual_study_completion_date"
	// FieldStudyID holds the string denoting the study_id field in the database.
	FieldStudyID = "study_id"
	// EdgeResultsDefinition holds the string denoting the results_definition edge name in mutations.
	EdgeResultsDefinition = "results_definition"
	// EdgeStudyLocations holds the string denoting the study_locations edge name in mutations.
	EdgeStudyLocations = "study_locations"
	// EdgeStudyEligibility holds the string denoting the study_eligibility edge name in mutations.
	EdgeStudyEligibility = "study_eligibility"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// Table holds the table name of the clinicaltrial in the database.
	Table = "clinical_trial"
	// ResultsDefinitionTable is the table that holds the results_definition relation/edge.
	ResultsDefinitionTable = "results_definition"
	// ResultsDefinitionInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ResultsDefinitionInverseTable = "results_definition"
	// ResultsDefinitionColumn is the table column denoting the results_definition relation/edge.
	ResultsDefinitionColumn = "clinical_trial_results_definition"
	// StudyLocationsTable is the table that holds the study_locations relation/edge.
	StudyLocationsTable = "study_location"
	// StudyLocationsInverseTable is the table name for the StudyLocation entity.
	// It exists in this package in order to avoid circular dependency with the "studylocation" package.
	StudyLocationsInverseTable = "study_location"
	// StudyLocationsColumn is the table column denoting the study_locations relation/edge.
	StudyLocationsColumn = "clinical_trial_study_locations"
	// StudyEligibilityTable is the table that holds the study_eligibility relation/edge.
	StudyEligibilityTable = "study_eligibility"
	// StudyEligibilityInverseTable is the table name for the StudyEligibility entity.
	// It exists in this package in order to avoid circular dependency with the "studyeligibility" package.
	StudyEligibilityInverseTable = "study_eligibility"
	// StudyEligibilityColumn is the table column denoting the study_eligibility relation/edge.
	StudyEligibilityColumn = "clinical_trial_study_eligibility"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "clinical_trial_metadata"
	// MetadataInverseTable is the table name for the ClinicalTrialMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "clinicaltrialmetadata" package.
	MetadataInverseTable = "clinical_trial_metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "clinical_trial_metadata"
)

// Columns holds all SQL columns for clinicaltrial fields.
var Columns = []string{
	FieldID,
	FieldStudyName,
	FieldSponsor,
	FieldResponsibleParty,
	FieldStudyType,
	FieldPhase,
	FieldActualEnrollment,
	FieldAllocation,
	FieldInterventionModel,
	FieldMasking,
	FieldPrimaryPurpose,
	FieldOfficialTitle,
	FieldActualStudyStartDate,
	FieldActualPrimaryCompletionDate,
	FieldActualStudyCompletionDate,
	FieldStudyID,
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
