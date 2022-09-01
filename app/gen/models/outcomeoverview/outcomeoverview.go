// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeoverview

const (
	// Label holds the string label denoting the outcomeoverview type in the database.
	Label = "outcome_overview"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeOverviewID holds the string denoting the outcome_overview_id field in the database.
	FieldOutcomeOverviewID = "outcome_overview_id"
	// FieldOutcomeOverviewTitle holds the string denoting the outcome_overview_title field in the database.
	FieldOutcomeOverviewTitle = "outcome_overview_title"
	// FieldOutcomeOverviewDescription holds the string denoting the outcome_overview_description field in the database.
	FieldOutcomeOverviewDescription = "outcome_overview_description"
	// FieldOutcomeOverviewParticipants holds the string denoting the outcome_overview_participants field in the database.
	FieldOutcomeOverviewParticipants = "outcome_overview_participants"
	// FieldOutcomeOverviewTimeFrame holds the string denoting the outcome_overview_time_frame field in the database.
	FieldOutcomeOverviewTimeFrame = "outcome_overview_time_frame"
	// FieldOutcomeOverviewSerotype holds the string denoting the outcome_overview_serotype field in the database.
	FieldOutcomeOverviewSerotype = "outcome_overview_serotype"
	// FieldOutcomeOverviewAssay holds the string denoting the outcome_overview_assay field in the database.
	FieldOutcomeOverviewAssay = "outcome_overview_assay"
	// FieldOutcomeOverviewDoseNumber holds the string denoting the outcome_overview_dose_number field in the database.
	FieldOutcomeOverviewDoseNumber = "outcome_overview_dose_number"
	// FieldOutcomeOverviewValue holds the string denoting the outcome_overview_value field in the database.
	FieldOutcomeOverviewValue = "outcome_overview_value"
	// FieldOutcomeOverviewUpperLimit holds the string denoting the outcome_overview_upper_limit field in the database.
	FieldOutcomeOverviewUpperLimit = "outcome_overview_upper_limit"
	// FieldOutcomeOverviewLowerLimit holds the string denoting the outcome_overview_lower_limit field in the database.
	FieldOutcomeOverviewLowerLimit = "outcome_overview_lower_limit"
	// FieldOutcomeOverviewGroupID holds the string denoting the outcome_overview_group_id field in the database.
	FieldOutcomeOverviewGroupID = "outcome_overview_group_id"
	// FieldOutcomeOverviewRatio holds the string denoting the outcome_overview_ratio field in the database.
	FieldOutcomeOverviewRatio = "outcome_overview_ratio"
	// FieldOutcomeOverviewMeasureTitle holds the string denoting the outcome_overview_measure_title field in the database.
	FieldOutcomeOverviewMeasureTitle = "outcome_overview_measure_title"
	// FieldOutcomeOverviewVaccine holds the string denoting the outcome_overview_vaccine field in the database.
	FieldOutcomeOverviewVaccine = "outcome_overview_vaccine"
	// FieldOutcomeOverviewImmunocompromisedPopulation holds the string denoting the outcome_overview_immunocompromised_population field in the database.
	FieldOutcomeOverviewImmunocompromisedPopulation = "outcome_overview_immunocompromised_population"
	// FieldOutcomeOverviewManufacturer holds the string denoting the outcome_overview_manufacturer field in the database.
	FieldOutcomeOverviewManufacturer = "outcome_overview_manufacturer"
	// FieldOutcomeOverviewConfidenceInterval holds the string denoting the outcome_overview_confidence_interval field in the database.
	FieldOutcomeOverviewConfidenceInterval = "outcome_overview_confidence_interval"
	// FieldOutcomeOverviewPercentResponders holds the string denoting the outcome_overview_percent_responders field in the database.
	FieldOutcomeOverviewPercentResponders = "outcome_overview_percent_responders"
	// FieldOutcomeOverviewTimeFrameWeeks holds the string denoting the outcome_overview_time_frame_weeks field in the database.
	FieldOutcomeOverviewTimeFrameWeeks = "outcome_overview_time_frame_weeks"
	// FieldOutcomeOverviewDoseDescription holds the string denoting the outcome_overview_dose_description field in the database.
	FieldOutcomeOverviewDoseDescription = "outcome_overview_dose_description"
	// FieldOutcomeOverviewSchedule holds the string denoting the outcome_overview_schedule field in the database.
	FieldOutcomeOverviewSchedule = "outcome_overview_schedule"
	// FieldOutcomeOverviewUseCaseCode holds the string denoting the outcome_overview_use_case_code field in the database.
	FieldOutcomeOverviewUseCaseCode = "outcome_overview_use_case_code"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the outcomeoverview in the database.
	Table = "outcome_overview"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_overview"
	// ParentInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	ParentInverseTable = "outcome_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measure_outcome_overview_list"
)

// Columns holds all SQL columns for outcomeoverview fields.
var Columns = []string{
	FieldID,
	FieldOutcomeOverviewID,
	FieldOutcomeOverviewTitle,
	FieldOutcomeOverviewDescription,
	FieldOutcomeOverviewParticipants,
	FieldOutcomeOverviewTimeFrame,
	FieldOutcomeOverviewSerotype,
	FieldOutcomeOverviewAssay,
	FieldOutcomeOverviewDoseNumber,
	FieldOutcomeOverviewValue,
	FieldOutcomeOverviewUpperLimit,
	FieldOutcomeOverviewLowerLimit,
	FieldOutcomeOverviewGroupID,
	FieldOutcomeOverviewRatio,
	FieldOutcomeOverviewMeasureTitle,
	FieldOutcomeOverviewVaccine,
	FieldOutcomeOverviewImmunocompromisedPopulation,
	FieldOutcomeOverviewManufacturer,
	FieldOutcomeOverviewConfidenceInterval,
	FieldOutcomeOverviewPercentResponders,
	FieldOutcomeOverviewTimeFrameWeeks,
	FieldOutcomeOverviewDoseDescription,
	FieldOutcomeOverviewSchedule,
	FieldOutcomeOverviewUseCaseCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_overview"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measure_outcome_overview_list",
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
	// DefaultOutcomeOverviewDoseNumber holds the default value on creation for the "outcome_overview_dose_number" field.
	DefaultOutcomeOverviewDoseNumber int64
	// DefaultOutcomeOverviewGroupID holds the default value on creation for the "outcome_overview_group_id" field.
	DefaultOutcomeOverviewGroupID string
	// DefaultOutcomeOverviewRatio holds the default value on creation for the "outcome_overview_ratio" field.
	DefaultOutcomeOverviewRatio string
	// DefaultOutcomeOverviewMeasureTitle holds the default value on creation for the "outcome_overview_measure_title" field.
	DefaultOutcomeOverviewMeasureTitle string
	// DefaultOutcomeOverviewVaccine holds the default value on creation for the "outcome_overview_vaccine" field.
	DefaultOutcomeOverviewVaccine string
	// DefaultOutcomeOverviewImmunocompromisedPopulation holds the default value on creation for the "outcome_overview_immunocompromised_population" field.
	DefaultOutcomeOverviewImmunocompromisedPopulation string
	// DefaultOutcomeOverviewManufacturer holds the default value on creation for the "outcome_overview_manufacturer" field.
	DefaultOutcomeOverviewManufacturer string
	// DefaultOutcomeOverviewConfidenceInterval holds the default value on creation for the "outcome_overview_confidence_interval" field.
	DefaultOutcomeOverviewConfidenceInterval string
	// DefaultOutcomeOverviewPercentResponders holds the default value on creation for the "outcome_overview_percent_responders" field.
	DefaultOutcomeOverviewPercentResponders float64
	// DefaultOutcomeOverviewTimeFrameWeeks holds the default value on creation for the "outcome_overview_time_frame_weeks" field.
	DefaultOutcomeOverviewTimeFrameWeeks int64
	// DefaultOutcomeOverviewDoseDescription holds the default value on creation for the "outcome_overview_dose_description" field.
	DefaultOutcomeOverviewDoseDescription string
	// DefaultOutcomeOverviewSchedule holds the default value on creation for the "outcome_overview_schedule" field.
	DefaultOutcomeOverviewSchedule string
	// DefaultOutcomeOverviewUseCaseCode holds the default value on creation for the "outcome_overview_use_case_code" field.
	DefaultOutcomeOverviewUseCaseCode string
)
