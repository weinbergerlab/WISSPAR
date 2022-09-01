// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeAnalysis is the model entity for the OutcomeAnalysis schema.
type OutcomeAnalysis struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeAnalysisGroupDescription holds the value of the "outcome_analysis_group_description" field.
	OutcomeAnalysisGroupDescription string `json:"outcome_analysis_group_description,omitempty"`
	// OutcomeAnalysisTestedNonInferiority holds the value of the "outcome_analysis_tested_non_inferiority" field.
	OutcomeAnalysisTestedNonInferiority string `json:"outcome_analysis_tested_non_inferiority,omitempty"`
	// OutcomeAnalysisNonInferiorityType holds the value of the "outcome_analysis_non_inferiority_type" field.
	OutcomeAnalysisNonInferiorityType string `json:"outcome_analysis_non_inferiority_type,omitempty"`
	// OutcomeAnalysisNonInferiorityComment holds the value of the "outcome_analysis_non_inferiority_comment" field.
	OutcomeAnalysisNonInferiorityComment string `json:"outcome_analysis_non_inferiority_comment,omitempty"`
	// OutcomeAnalysisPValue holds the value of the "outcome_analysis_p_value" field.
	OutcomeAnalysisPValue string `json:"outcome_analysis_p_value,omitempty"`
	// OutcomeAnalysisPValueComment holds the value of the "outcome_analysis_p_value_comment" field.
	OutcomeAnalysisPValueComment string `json:"outcome_analysis_p_value_comment,omitempty"`
	// OutcomeAnalysisStatisticalMethod holds the value of the "outcome_analysis_statistical_method" field.
	OutcomeAnalysisStatisticalMethod string `json:"outcome_analysis_statistical_method,omitempty"`
	// OutcomeAnalysisStatisticalComment holds the value of the "outcome_analysis_statistical_comment" field.
	OutcomeAnalysisStatisticalComment string `json:"outcome_analysis_statistical_comment,omitempty"`
	// OutcomeAnalysisParamType holds the value of the "outcome_analysis_param_type" field.
	OutcomeAnalysisParamType string `json:"outcome_analysis_param_type,omitempty"`
	// OutcomeAnalysisParamValue holds the value of the "outcome_analysis_param_value" field.
	OutcomeAnalysisParamValue string `json:"outcome_analysis_param_value,omitempty"`
	// OutcomeAnalysisCiPctValue holds the value of the "outcome_analysis_ci_pct_value" field.
	OutcomeAnalysisCiPctValue string `json:"outcome_analysis_ci_pct_value,omitempty"`
	// OutcomeAnalysisCiNumSides holds the value of the "outcome_analysis_ci_num_sides" field.
	OutcomeAnalysisCiNumSides string `json:"outcome_analysis_ci_num_sides,omitempty"`
	// OutcomeAnalysisCiLowerLimit holds the value of the "outcome_analysis_ci_lower_limit" field.
	OutcomeAnalysisCiLowerLimit string `json:"outcome_analysis_ci_lower_limit,omitempty"`
	// OutcomeAnalysisCiUpperLimit holds the value of the "outcome_analysis_ci_upper_limit" field.
	OutcomeAnalysisCiUpperLimit string `json:"outcome_analysis_ci_upper_limit,omitempty"`
	// OutcomeAnalysisCiLowerLimitComment holds the value of the "outcome_analysis_ci_lower_limit_comment" field.
	OutcomeAnalysisCiLowerLimitComment string `json:"outcome_analysis_ci_lower_limit_comment,omitempty"`
	// OutcomeAnalysisCiUpperLimitComment holds the value of the "outcome_analysis_ci_upper_limit_comment" field.
	OutcomeAnalysisCiUpperLimitComment string `json:"outcome_analysis_ci_upper_limit_comment,omitempty"`
	// OutcomeAnalysisDispersionType holds the value of the "outcome_analysis_dispersion_type" field.
	OutcomeAnalysisDispersionType string `json:"outcome_analysis_dispersion_type,omitempty"`
	// OutcomeAnalysisDispersionValue holds the value of the "outcome_analysis_dispersion_value" field.
	OutcomeAnalysisDispersionValue string `json:"outcome_analysis_dispersion_value,omitempty"`
	// OutcomeAnalysisEstimateComment holds the value of the "outcome_analysis_estimate_comment" field.
	OutcomeAnalysisEstimateComment string `json:"outcome_analysis_estimate_comment,omitempty"`
	// OutcomeAnalysisOtherAnalysisDescription holds the value of the "outcome_analysis_other_analysis_description" field.
	OutcomeAnalysisOtherAnalysisDescription string `json:"outcome_analysis_other_analysis_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeAnalysisQuery when eager-loading is set.
	Edges                                 OutcomeAnalysisEdges `json:"edges"`
	outcome_measure_outcome_analysis_list *int
}

// OutcomeAnalysisEdges holds the relations/edges for other nodes in the graph.
type OutcomeAnalysisEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasure `json:"parent,omitempty"`
	// OutcomeAnalysisGroupIDList holds the value of the outcome_analysis_group_id_list edge.
	OutcomeAnalysisGroupIDList []*OutcomeAnalysisGroupID `json:"outcome_analysis_group_id_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeAnalysisEdges) ParentOrErr() (*OutcomeMeasure, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomemeasure.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// OutcomeAnalysisGroupIDListOrErr returns the OutcomeAnalysisGroupIDList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeAnalysisEdges) OutcomeAnalysisGroupIDListOrErr() ([]*OutcomeAnalysisGroupID, error) {
	if e.loadedTypes[1] {
		return e.OutcomeAnalysisGroupIDList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_analysis_group_id_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeAnalysis) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeanalysis.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomeanalysis.FieldOutcomeAnalysisGroupDescription, outcomeanalysis.FieldOutcomeAnalysisTestedNonInferiority, outcomeanalysis.FieldOutcomeAnalysisNonInferiorityType, outcomeanalysis.FieldOutcomeAnalysisNonInferiorityComment, outcomeanalysis.FieldOutcomeAnalysisPValue, outcomeanalysis.FieldOutcomeAnalysisPValueComment, outcomeanalysis.FieldOutcomeAnalysisStatisticalMethod, outcomeanalysis.FieldOutcomeAnalysisStatisticalComment, outcomeanalysis.FieldOutcomeAnalysisParamType, outcomeanalysis.FieldOutcomeAnalysisParamValue, outcomeanalysis.FieldOutcomeAnalysisCiPctValue, outcomeanalysis.FieldOutcomeAnalysisCiNumSides, outcomeanalysis.FieldOutcomeAnalysisCiLowerLimit, outcomeanalysis.FieldOutcomeAnalysisCiUpperLimit, outcomeanalysis.FieldOutcomeAnalysisCiLowerLimitComment, outcomeanalysis.FieldOutcomeAnalysisCiUpperLimitComment, outcomeanalysis.FieldOutcomeAnalysisDispersionType, outcomeanalysis.FieldOutcomeAnalysisDispersionValue, outcomeanalysis.FieldOutcomeAnalysisEstimateComment, outcomeanalysis.FieldOutcomeAnalysisOtherAnalysisDescription:
			values[i] = new(sql.NullString)
		case outcomeanalysis.ForeignKeys[0]: // outcome_measure_outcome_analysis_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeAnalysis", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeAnalysis fields.
func (oa *OutcomeAnalysis) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeanalysis.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oa.ID = int(value.Int64)
		case outcomeanalysis.FieldOutcomeAnalysisGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_group_description", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisGroupDescription = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisTestedNonInferiority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_tested_non_inferiority", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisTestedNonInferiority = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisNonInferiorityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_non_inferiority_type", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisNonInferiorityType = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisNonInferiorityComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_non_inferiority_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisNonInferiorityComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisPValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_p_value", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisPValue = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisPValueComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_p_value_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisPValueComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisStatisticalMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_statistical_method", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisStatisticalMethod = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisStatisticalComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_statistical_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisStatisticalComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisParamType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_param_type", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisParamType = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisParamValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_param_value", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisParamValue = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiPctValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_pct_value", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiPctValue = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiNumSides:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_num_sides", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiNumSides = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiLowerLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_lower_limit", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiLowerLimit = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiUpperLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_upper_limit", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiUpperLimit = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiLowerLimitComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_lower_limit_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiLowerLimitComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisCiUpperLimitComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_ci_upper_limit_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisCiUpperLimitComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisDispersionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_dispersion_type", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisDispersionType = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisDispersionValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_dispersion_value", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisDispersionValue = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisEstimateComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_estimate_comment", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisEstimateComment = value.String
			}
		case outcomeanalysis.FieldOutcomeAnalysisOtherAnalysisDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_other_analysis_description", values[i])
			} else if value.Valid {
				oa.OutcomeAnalysisOtherAnalysisDescription = value.String
			}
		case outcomeanalysis.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measure_outcome_analysis_list", value)
			} else if value.Valid {
				oa.outcome_measure_outcome_analysis_list = new(int)
				*oa.outcome_measure_outcome_analysis_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeAnalysis entity.
func (oa *OutcomeAnalysis) QueryParent() *OutcomeMeasureQuery {
	return (&OutcomeAnalysisClient{config: oa.config}).QueryParent(oa)
}

// QueryOutcomeAnalysisGroupIDList queries the "outcome_analysis_group_id_list" edge of the OutcomeAnalysis entity.
func (oa *OutcomeAnalysis) QueryOutcomeAnalysisGroupIDList() *OutcomeAnalysisGroupIDQuery {
	return (&OutcomeAnalysisClient{config: oa.config}).QueryOutcomeAnalysisGroupIDList(oa)
}

// Update returns a builder for updating this OutcomeAnalysis.
// Note that you need to call OutcomeAnalysis.Unwrap() before calling this method if this OutcomeAnalysis
// was returned from a transaction, and the transaction was committed or rolled back.
func (oa *OutcomeAnalysis) Update() *OutcomeAnalysisUpdateOne {
	return (&OutcomeAnalysisClient{config: oa.config}).UpdateOne(oa)
}

// Unwrap unwraps the OutcomeAnalysis entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oa *OutcomeAnalysis) Unwrap() *OutcomeAnalysis {
	tx, ok := oa.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeAnalysis is not a transactional entity")
	}
	oa.config.driver = tx.drv
	return oa
}

// String implements the fmt.Stringer.
func (oa *OutcomeAnalysis) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeAnalysis(")
	builder.WriteString(fmt.Sprintf("id=%v", oa.ID))
	builder.WriteString(", outcome_analysis_group_description=")
	builder.WriteString(oa.OutcomeAnalysisGroupDescription)
	builder.WriteString(", outcome_analysis_tested_non_inferiority=")
	builder.WriteString(oa.OutcomeAnalysisTestedNonInferiority)
	builder.WriteString(", outcome_analysis_non_inferiority_type=")
	builder.WriteString(oa.OutcomeAnalysisNonInferiorityType)
	builder.WriteString(", outcome_analysis_non_inferiority_comment=")
	builder.WriteString(oa.OutcomeAnalysisNonInferiorityComment)
	builder.WriteString(", outcome_analysis_p_value=")
	builder.WriteString(oa.OutcomeAnalysisPValue)
	builder.WriteString(", outcome_analysis_p_value_comment=")
	builder.WriteString(oa.OutcomeAnalysisPValueComment)
	builder.WriteString(", outcome_analysis_statistical_method=")
	builder.WriteString(oa.OutcomeAnalysisStatisticalMethod)
	builder.WriteString(", outcome_analysis_statistical_comment=")
	builder.WriteString(oa.OutcomeAnalysisStatisticalComment)
	builder.WriteString(", outcome_analysis_param_type=")
	builder.WriteString(oa.OutcomeAnalysisParamType)
	builder.WriteString(", outcome_analysis_param_value=")
	builder.WriteString(oa.OutcomeAnalysisParamValue)
	builder.WriteString(", outcome_analysis_ci_pct_value=")
	builder.WriteString(oa.OutcomeAnalysisCiPctValue)
	builder.WriteString(", outcome_analysis_ci_num_sides=")
	builder.WriteString(oa.OutcomeAnalysisCiNumSides)
	builder.WriteString(", outcome_analysis_ci_lower_limit=")
	builder.WriteString(oa.OutcomeAnalysisCiLowerLimit)
	builder.WriteString(", outcome_analysis_ci_upper_limit=")
	builder.WriteString(oa.OutcomeAnalysisCiUpperLimit)
	builder.WriteString(", outcome_analysis_ci_lower_limit_comment=")
	builder.WriteString(oa.OutcomeAnalysisCiLowerLimitComment)
	builder.WriteString(", outcome_analysis_ci_upper_limit_comment=")
	builder.WriteString(oa.OutcomeAnalysisCiUpperLimitComment)
	builder.WriteString(", outcome_analysis_dispersion_type=")
	builder.WriteString(oa.OutcomeAnalysisDispersionType)
	builder.WriteString(", outcome_analysis_dispersion_value=")
	builder.WriteString(oa.OutcomeAnalysisDispersionValue)
	builder.WriteString(", outcome_analysis_estimate_comment=")
	builder.WriteString(oa.OutcomeAnalysisEstimateComment)
	builder.WriteString(", outcome_analysis_other_analysis_description=")
	builder.WriteString(oa.OutcomeAnalysisOtherAnalysisDescription)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeAnalyses is a parsable slice of OutcomeAnalysis.
type OutcomeAnalyses []*OutcomeAnalysis

func (oa OutcomeAnalyses) config(cfg config) {
	for _i := range oa {
		oa[_i].config = cfg
	}
}
