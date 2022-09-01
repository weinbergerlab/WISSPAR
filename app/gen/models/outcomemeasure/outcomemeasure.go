// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasure

const (
	// Label holds the string label denoting the outcomemeasure type in the database.
	Label = "outcome_measure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeMeasureType holds the string denoting the outcome_measure_type field in the database.
	FieldOutcomeMeasureType = "outcome_measure_type"
	// FieldOutcomeMeasureTitle holds the string denoting the outcome_measure_title field in the database.
	FieldOutcomeMeasureTitle = "outcome_measure_title"
	// FieldOutcomeMeasureDescription holds the string denoting the outcome_measure_description field in the database.
	FieldOutcomeMeasureDescription = "outcome_measure_description"
	// FieldOutcomeMeasurePopulationDescription holds the string denoting the outcome_measure_population_description field in the database.
	FieldOutcomeMeasurePopulationDescription = "outcome_measure_population_description"
	// FieldOutcomeMeasureReportingStatus holds the string denoting the outcome_measure_reporting_status field in the database.
	FieldOutcomeMeasureReportingStatus = "outcome_measure_reporting_status"
	// FieldOutcomeMeasureAnticipatedPostingDate holds the string denoting the outcome_measure_anticipated_posting_date field in the database.
	FieldOutcomeMeasureAnticipatedPostingDate = "outcome_measure_anticipated_posting_date"
	// FieldOutcomeMeasureParamType holds the string denoting the outcome_measure_param_type field in the database.
	FieldOutcomeMeasureParamType = "outcome_measure_param_type"
	// FieldOutcomeMeasureDispersionType holds the string denoting the outcome_measure_dispersion_type field in the database.
	FieldOutcomeMeasureDispersionType = "outcome_measure_dispersion_type"
	// FieldOutcomeMeasureUnitOfMeasure holds the string denoting the outcome_measure_unit_of_measure field in the database.
	FieldOutcomeMeasureUnitOfMeasure = "outcome_measure_unit_of_measure"
	// FieldOutcomeMeasureCalculatePct holds the string denoting the outcome_measure_calculate_pct field in the database.
	FieldOutcomeMeasureCalculatePct = "outcome_measure_calculate_pct"
	// FieldOutcomeMeasureTimeFrame holds the string denoting the outcome_measure_time_frame field in the database.
	FieldOutcomeMeasureTimeFrame = "outcome_measure_time_frame"
	// FieldOutcomeMeasureTypeUnitsAnalyzed holds the string denoting the outcome_measure_type_units_analyzed field in the database.
	FieldOutcomeMeasureTypeUnitsAnalyzed = "outcome_measure_type_units_analyzed"
	// FieldOutcomeMeasureDenomUnitsSelected holds the string denoting the outcome_measure_denom_units_selected field in the database.
	FieldOutcomeMeasureDenomUnitsSelected = "outcome_measure_denom_units_selected"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeGroupList holds the string denoting the outcome_group_list edge name in mutations.
	EdgeOutcomeGroupList = "outcome_group_list"
	// EdgeOutcomeOverviewList holds the string denoting the outcome_overview_list edge name in mutations.
	EdgeOutcomeOverviewList = "outcome_overview_list"
	// EdgeOutcomeDenomList holds the string denoting the outcome_denom_list edge name in mutations.
	EdgeOutcomeDenomList = "outcome_denom_list"
	// EdgeOutcomeClassList holds the string denoting the outcome_class_list edge name in mutations.
	EdgeOutcomeClassList = "outcome_class_list"
	// EdgeOutcomeAnalysisList holds the string denoting the outcome_analysis_list edge name in mutations.
	EdgeOutcomeAnalysisList = "outcome_analysis_list"
	// Table holds the table name of the outcomemeasure in the database.
	Table = "outcome_measure"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_measure"
	// ParentInverseTable is the table name for the OutcomeMeasuresModule entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasuresmodule" package.
	ParentInverseTable = "outcome_measures_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measures_module_outcome_measure_list"
	// OutcomeGroupListTable is the table that holds the outcome_group_list relation/edge.
	OutcomeGroupListTable = "outcome_group"
	// OutcomeGroupListInverseTable is the table name for the OutcomeGroup entity.
	// It exists in this package in order to avoid circular dependency with the "outcomegroup" package.
	OutcomeGroupListInverseTable = "outcome_group"
	// OutcomeGroupListColumn is the table column denoting the outcome_group_list relation/edge.
	OutcomeGroupListColumn = "outcome_measure_outcome_group_list"
	// OutcomeOverviewListTable is the table that holds the outcome_overview_list relation/edge.
	OutcomeOverviewListTable = "outcome_overview"
	// OutcomeOverviewListInverseTable is the table name for the OutcomeOverview entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeoverview" package.
	OutcomeOverviewListInverseTable = "outcome_overview"
	// OutcomeOverviewListColumn is the table column denoting the outcome_overview_list relation/edge.
	OutcomeOverviewListColumn = "outcome_measure_outcome_overview_list"
	// OutcomeDenomListTable is the table that holds the outcome_denom_list relation/edge.
	OutcomeDenomListTable = "outcome_denom"
	// OutcomeDenomListInverseTable is the table name for the OutcomeDenom entity.
	// It exists in this package in order to avoid circular dependency with the "outcomedenom" package.
	OutcomeDenomListInverseTable = "outcome_denom"
	// OutcomeDenomListColumn is the table column denoting the outcome_denom_list relation/edge.
	OutcomeDenomListColumn = "outcome_measure_outcome_denom_list"
	// OutcomeClassListTable is the table that holds the outcome_class_list relation/edge.
	OutcomeClassListTable = "outcome_class"
	// OutcomeClassListInverseTable is the table name for the OutcomeClass entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeclass" package.
	OutcomeClassListInverseTable = "outcome_class"
	// OutcomeClassListColumn is the table column denoting the outcome_class_list relation/edge.
	OutcomeClassListColumn = "outcome_measure_outcome_class_list"
	// OutcomeAnalysisListTable is the table that holds the outcome_analysis_list relation/edge.
	OutcomeAnalysisListTable = "outcome_analysis"
	// OutcomeAnalysisListInverseTable is the table name for the OutcomeAnalysis entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeanalysis" package.
	OutcomeAnalysisListInverseTable = "outcome_analysis"
	// OutcomeAnalysisListColumn is the table column denoting the outcome_analysis_list relation/edge.
	OutcomeAnalysisListColumn = "outcome_measure_outcome_analysis_list"
)

// Columns holds all SQL columns for outcomemeasure fields.
var Columns = []string{
	FieldID,
	FieldOutcomeMeasureType,
	FieldOutcomeMeasureTitle,
	FieldOutcomeMeasureDescription,
	FieldOutcomeMeasurePopulationDescription,
	FieldOutcomeMeasureReportingStatus,
	FieldOutcomeMeasureAnticipatedPostingDate,
	FieldOutcomeMeasureParamType,
	FieldOutcomeMeasureDispersionType,
	FieldOutcomeMeasureUnitOfMeasure,
	FieldOutcomeMeasureCalculatePct,
	FieldOutcomeMeasureTimeFrame,
	FieldOutcomeMeasureTypeUnitsAnalyzed,
	FieldOutcomeMeasureDenomUnitsSelected,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_measure"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measures_module_outcome_measure_list",
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
