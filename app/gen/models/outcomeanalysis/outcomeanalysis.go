// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeanalysis

const (
	// Label holds the string label denoting the outcomeanalysis type in the database.
	Label = "outcome_analysis"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutcomeAnalysisGroupDescription holds the string denoting the outcome_analysis_group_description field in the database.
	FieldOutcomeAnalysisGroupDescription = "outcome_analysis_group_description"
	// FieldOutcomeAnalysisTestedNonInferiority holds the string denoting the outcome_analysis_tested_non_inferiority field in the database.
	FieldOutcomeAnalysisTestedNonInferiority = "outcome_analysis_tested_non_inferiority"
	// FieldOutcomeAnalysisNonInferiorityType holds the string denoting the outcome_analysis_non_inferiority_type field in the database.
	FieldOutcomeAnalysisNonInferiorityType = "outcome_analysis_non_inferiority_type"
	// FieldOutcomeAnalysisNonInferiorityComment holds the string denoting the outcome_analysis_non_inferiority_comment field in the database.
	FieldOutcomeAnalysisNonInferiorityComment = "outcome_analysis_non_inferiority_comment"
	// FieldOutcomeAnalysisPValue holds the string denoting the outcome_analysis_p_value field in the database.
	FieldOutcomeAnalysisPValue = "outcome_analysis_p_value"
	// FieldOutcomeAnalysisPValueComment holds the string denoting the outcome_analysis_p_value_comment field in the database.
	FieldOutcomeAnalysisPValueComment = "outcome_analysis_p_value_comment"
	// FieldOutcomeAnalysisStatisticalMethod holds the string denoting the outcome_analysis_statistical_method field in the database.
	FieldOutcomeAnalysisStatisticalMethod = "outcome_analysis_statistical_method"
	// FieldOutcomeAnalysisStatisticalComment holds the string denoting the outcome_analysis_statistical_comment field in the database.
	FieldOutcomeAnalysisStatisticalComment = "outcome_analysis_statistical_comment"
	// FieldOutcomeAnalysisParamType holds the string denoting the outcome_analysis_param_type field in the database.
	FieldOutcomeAnalysisParamType = "outcome_analysis_param_type"
	// FieldOutcomeAnalysisParamValue holds the string denoting the outcome_analysis_param_value field in the database.
	FieldOutcomeAnalysisParamValue = "outcome_analysis_param_value"
	// FieldOutcomeAnalysisCiPctValue holds the string denoting the outcome_analysis_ci_pct_value field in the database.
	FieldOutcomeAnalysisCiPctValue = "outcome_analysis_ci_pct_value"
	// FieldOutcomeAnalysisCiNumSides holds the string denoting the outcome_analysis_ci_num_sides field in the database.
	FieldOutcomeAnalysisCiNumSides = "outcome_analysis_ci_num_sides"
	// FieldOutcomeAnalysisCiLowerLimit holds the string denoting the outcome_analysis_ci_lower_limit field in the database.
	FieldOutcomeAnalysisCiLowerLimit = "outcome_analysis_ci_lower_limit"
	// FieldOutcomeAnalysisCiUpperLimit holds the string denoting the outcome_analysis_ci_upper_limit field in the database.
	FieldOutcomeAnalysisCiUpperLimit = "outcome_analysis_ci_upper_limit"
	// FieldOutcomeAnalysisCiLowerLimitComment holds the string denoting the outcome_analysis_ci_lower_limit_comment field in the database.
	FieldOutcomeAnalysisCiLowerLimitComment = "outcome_analysis_ci_lower_limit_comment"
	// FieldOutcomeAnalysisCiUpperLimitComment holds the string denoting the outcome_analysis_ci_upper_limit_comment field in the database.
	FieldOutcomeAnalysisCiUpperLimitComment = "outcome_analysis_ci_upper_limit_comment"
	// FieldOutcomeAnalysisDispersionType holds the string denoting the outcome_analysis_dispersion_type field in the database.
	FieldOutcomeAnalysisDispersionType = "outcome_analysis_dispersion_type"
	// FieldOutcomeAnalysisDispersionValue holds the string denoting the outcome_analysis_dispersion_value field in the database.
	FieldOutcomeAnalysisDispersionValue = "outcome_analysis_dispersion_value"
	// FieldOutcomeAnalysisEstimateComment holds the string denoting the outcome_analysis_estimate_comment field in the database.
	FieldOutcomeAnalysisEstimateComment = "outcome_analysis_estimate_comment"
	// FieldOutcomeAnalysisOtherAnalysisDescription holds the string denoting the outcome_analysis_other_analysis_description field in the database.
	FieldOutcomeAnalysisOtherAnalysisDescription = "outcome_analysis_other_analysis_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOutcomeAnalysisGroupIDList holds the string denoting the outcome_analysis_group_id_list edge name in mutations.
	EdgeOutcomeAnalysisGroupIDList = "outcome_analysis_group_id_list"
	// Table holds the table name of the outcomeanalysis in the database.
	Table = "outcome_analysis"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "outcome_analysis"
	// ParentInverseTable is the table name for the OutcomeMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "outcomemeasure" package.
	ParentInverseTable = "outcome_measure"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "outcome_measure_outcome_analysis_list"
	// OutcomeAnalysisGroupIDListTable is the table that holds the outcome_analysis_group_id_list relation/edge.
	OutcomeAnalysisGroupIDListTable = "outcome_analysis_group_id"
	// OutcomeAnalysisGroupIDListInverseTable is the table name for the OutcomeAnalysisGroupID entity.
	// It exists in this package in order to avoid circular dependency with the "outcomeanalysisgroupid" package.
	OutcomeAnalysisGroupIDListInverseTable = "outcome_analysis_group_id"
	// OutcomeAnalysisGroupIDListColumn is the table column denoting the outcome_analysis_group_id_list relation/edge.
	OutcomeAnalysisGroupIDListColumn = "outcome_analysis_outcome_analysis_group_id_list"
)

// Columns holds all SQL columns for outcomeanalysis fields.
var Columns = []string{
	FieldID,
	FieldOutcomeAnalysisGroupDescription,
	FieldOutcomeAnalysisTestedNonInferiority,
	FieldOutcomeAnalysisNonInferiorityType,
	FieldOutcomeAnalysisNonInferiorityComment,
	FieldOutcomeAnalysisPValue,
	FieldOutcomeAnalysisPValueComment,
	FieldOutcomeAnalysisStatisticalMethod,
	FieldOutcomeAnalysisStatisticalComment,
	FieldOutcomeAnalysisParamType,
	FieldOutcomeAnalysisParamValue,
	FieldOutcomeAnalysisCiPctValue,
	FieldOutcomeAnalysisCiNumSides,
	FieldOutcomeAnalysisCiLowerLimit,
	FieldOutcomeAnalysisCiUpperLimit,
	FieldOutcomeAnalysisCiLowerLimitComment,
	FieldOutcomeAnalysisCiUpperLimitComment,
	FieldOutcomeAnalysisDispersionType,
	FieldOutcomeAnalysisDispersionValue,
	FieldOutcomeAnalysisEstimateComment,
	FieldOutcomeAnalysisOtherAnalysisDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "outcome_analysis"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"outcome_measure_outcome_analysis_list",
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
