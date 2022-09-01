// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
)

// OutcomeMeasure is the model entity for the OutcomeMeasure schema.
type OutcomeMeasure struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeMeasureType holds the value of the "outcome_measure_type" field.
	OutcomeMeasureType string `json:"outcome_measure_type,omitempty"`
	// OutcomeMeasureTitle holds the value of the "outcome_measure_title" field.
	OutcomeMeasureTitle string `json:"outcome_measure_title,omitempty"`
	// OutcomeMeasureDescription holds the value of the "outcome_measure_description" field.
	OutcomeMeasureDescription string `json:"outcome_measure_description,omitempty"`
	// OutcomeMeasurePopulationDescription holds the value of the "outcome_measure_population_description" field.
	OutcomeMeasurePopulationDescription string `json:"outcome_measure_population_description,omitempty"`
	// OutcomeMeasureReportingStatus holds the value of the "outcome_measure_reporting_status" field.
	OutcomeMeasureReportingStatus string `json:"outcome_measure_reporting_status,omitempty"`
	// OutcomeMeasureAnticipatedPostingDate holds the value of the "outcome_measure_anticipated_posting_date" field.
	OutcomeMeasureAnticipatedPostingDate string `json:"outcome_measure_anticipated_posting_date,omitempty"`
	// OutcomeMeasureParamType holds the value of the "outcome_measure_param_type" field.
	OutcomeMeasureParamType string `json:"outcome_measure_param_type,omitempty"`
	// OutcomeMeasureDispersionType holds the value of the "outcome_measure_dispersion_type" field.
	OutcomeMeasureDispersionType string `json:"outcome_measure_dispersion_type,omitempty"`
	// OutcomeMeasureUnitOfMeasure holds the value of the "outcome_measure_unit_of_measure" field.
	OutcomeMeasureUnitOfMeasure string `json:"outcome_measure_unit_of_measure,omitempty"`
	// OutcomeMeasureCalculatePct holds the value of the "outcome_measure_calculate_pct" field.
	OutcomeMeasureCalculatePct string `json:"outcome_measure_calculate_pct,omitempty"`
	// OutcomeMeasureTimeFrame holds the value of the "outcome_measure_time_frame" field.
	OutcomeMeasureTimeFrame string `json:"outcome_measure_time_frame,omitempty"`
	// OutcomeMeasureTypeUnitsAnalyzed holds the value of the "outcome_measure_type_units_analyzed" field.
	OutcomeMeasureTypeUnitsAnalyzed string `json:"outcome_measure_type_units_analyzed,omitempty"`
	// OutcomeMeasureDenomUnitsSelected holds the value of the "outcome_measure_denom_units_selected" field.
	OutcomeMeasureDenomUnitsSelected string `json:"outcome_measure_denom_units_selected,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeMeasureQuery when eager-loading is set.
	Edges                                        OutcomeMeasureEdges `json:"edges"`
	outcome_measures_module_outcome_measure_list *int
}

// OutcomeMeasureEdges holds the relations/edges for other nodes in the graph.
type OutcomeMeasureEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasuresModule `json:"parent,omitempty"`
	// OutcomeGroupList holds the value of the outcome_group_list edge.
	OutcomeGroupList []*OutcomeGroup `json:"outcome_group_list,omitempty"`
	// OutcomeOverviewList holds the value of the outcome_overview_list edge.
	OutcomeOverviewList []*OutcomeOverview `json:"outcome_overview_list,omitempty"`
	// OutcomeDenomList holds the value of the outcome_denom_list edge.
	OutcomeDenomList []*OutcomeDenom `json:"outcome_denom_list,omitempty"`
	// OutcomeClassList holds the value of the outcome_class_list edge.
	OutcomeClassList []*OutcomeClass `json:"outcome_class_list,omitempty"`
	// OutcomeAnalysisList holds the value of the outcome_analysis_list edge.
	OutcomeAnalysisList []*OutcomeAnalysis `json:"outcome_analysis_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeMeasureEdges) ParentOrErr() (*OutcomeMeasuresModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomemeasuresmodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// OutcomeGroupListOrErr returns the OutcomeGroupList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasureEdges) OutcomeGroupListOrErr() ([]*OutcomeGroup, error) {
	if e.loadedTypes[1] {
		return e.OutcomeGroupList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_group_list"}
}

// OutcomeOverviewListOrErr returns the OutcomeOverviewList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasureEdges) OutcomeOverviewListOrErr() ([]*OutcomeOverview, error) {
	if e.loadedTypes[2] {
		return e.OutcomeOverviewList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_overview_list"}
}

// OutcomeDenomListOrErr returns the OutcomeDenomList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasureEdges) OutcomeDenomListOrErr() ([]*OutcomeDenom, error) {
	if e.loadedTypes[3] {
		return e.OutcomeDenomList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_denom_list"}
}

// OutcomeClassListOrErr returns the OutcomeClassList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasureEdges) OutcomeClassListOrErr() ([]*OutcomeClass, error) {
	if e.loadedTypes[4] {
		return e.OutcomeClassList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_class_list"}
}

// OutcomeAnalysisListOrErr returns the OutcomeAnalysisList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasureEdges) OutcomeAnalysisListOrErr() ([]*OutcomeAnalysis, error) {
	if e.loadedTypes[5] {
		return e.OutcomeAnalysisList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_analysis_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeMeasure) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomemeasure.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomemeasure.FieldOutcomeMeasureType, outcomemeasure.FieldOutcomeMeasureTitle, outcomemeasure.FieldOutcomeMeasureDescription, outcomemeasure.FieldOutcomeMeasurePopulationDescription, outcomemeasure.FieldOutcomeMeasureReportingStatus, outcomemeasure.FieldOutcomeMeasureAnticipatedPostingDate, outcomemeasure.FieldOutcomeMeasureParamType, outcomemeasure.FieldOutcomeMeasureDispersionType, outcomemeasure.FieldOutcomeMeasureUnitOfMeasure, outcomemeasure.FieldOutcomeMeasureCalculatePct, outcomemeasure.FieldOutcomeMeasureTimeFrame, outcomemeasure.FieldOutcomeMeasureTypeUnitsAnalyzed, outcomemeasure.FieldOutcomeMeasureDenomUnitsSelected:
			values[i] = new(sql.NullString)
		case outcomemeasure.ForeignKeys[0]: // outcome_measures_module_outcome_measure_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeMeasure", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeMeasure fields.
func (om *OutcomeMeasure) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomemeasure.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			om.ID = int(value.Int64)
		case outcomemeasure.FieldOutcomeMeasureType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_type", values[i])
			} else if value.Valid {
				om.OutcomeMeasureType = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_title", values[i])
			} else if value.Valid {
				om.OutcomeMeasureTitle = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_description", values[i])
			} else if value.Valid {
				om.OutcomeMeasureDescription = value.String
			}
		case outcomemeasure.FieldOutcomeMeasurePopulationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_population_description", values[i])
			} else if value.Valid {
				om.OutcomeMeasurePopulationDescription = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureReportingStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_reporting_status", values[i])
			} else if value.Valid {
				om.OutcomeMeasureReportingStatus = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureAnticipatedPostingDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_anticipated_posting_date", values[i])
			} else if value.Valid {
				om.OutcomeMeasureAnticipatedPostingDate = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureParamType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_param_type", values[i])
			} else if value.Valid {
				om.OutcomeMeasureParamType = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureDispersionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_dispersion_type", values[i])
			} else if value.Valid {
				om.OutcomeMeasureDispersionType = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureUnitOfMeasure:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_unit_of_measure", values[i])
			} else if value.Valid {
				om.OutcomeMeasureUnitOfMeasure = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureCalculatePct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_calculate_pct", values[i])
			} else if value.Valid {
				om.OutcomeMeasureCalculatePct = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureTimeFrame:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_time_frame", values[i])
			} else if value.Valid {
				om.OutcomeMeasureTimeFrame = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureTypeUnitsAnalyzed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_type_units_analyzed", values[i])
			} else if value.Valid {
				om.OutcomeMeasureTypeUnitsAnalyzed = value.String
			}
		case outcomemeasure.FieldOutcomeMeasureDenomUnitsSelected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measure_denom_units_selected", values[i])
			} else if value.Valid {
				om.OutcomeMeasureDenomUnitsSelected = value.String
			}
		case outcomemeasure.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measures_module_outcome_measure_list", value)
			} else if value.Valid {
				om.outcome_measures_module_outcome_measure_list = new(int)
				*om.outcome_measures_module_outcome_measure_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryParent() *OutcomeMeasuresModuleQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryParent(om)
}

// QueryOutcomeGroupList queries the "outcome_group_list" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryOutcomeGroupList() *OutcomeGroupQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryOutcomeGroupList(om)
}

// QueryOutcomeOverviewList queries the "outcome_overview_list" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryOutcomeOverviewList() *OutcomeOverviewQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryOutcomeOverviewList(om)
}

// QueryOutcomeDenomList queries the "outcome_denom_list" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryOutcomeDenomList() *OutcomeDenomQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryOutcomeDenomList(om)
}

// QueryOutcomeClassList queries the "outcome_class_list" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryOutcomeClassList() *OutcomeClassQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryOutcomeClassList(om)
}

// QueryOutcomeAnalysisList queries the "outcome_analysis_list" edge of the OutcomeMeasure entity.
func (om *OutcomeMeasure) QueryOutcomeAnalysisList() *OutcomeAnalysisQuery {
	return (&OutcomeMeasureClient{config: om.config}).QueryOutcomeAnalysisList(om)
}

// Update returns a builder for updating this OutcomeMeasure.
// Note that you need to call OutcomeMeasure.Unwrap() before calling this method if this OutcomeMeasure
// was returned from a transaction, and the transaction was committed or rolled back.
func (om *OutcomeMeasure) Update() *OutcomeMeasureUpdateOne {
	return (&OutcomeMeasureClient{config: om.config}).UpdateOne(om)
}

// Unwrap unwraps the OutcomeMeasure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (om *OutcomeMeasure) Unwrap() *OutcomeMeasure {
	tx, ok := om.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeMeasure is not a transactional entity")
	}
	om.config.driver = tx.drv
	return om
}

// String implements the fmt.Stringer.
func (om *OutcomeMeasure) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeMeasure(")
	builder.WriteString(fmt.Sprintf("id=%v", om.ID))
	builder.WriteString(", outcome_measure_type=")
	builder.WriteString(om.OutcomeMeasureType)
	builder.WriteString(", outcome_measure_title=")
	builder.WriteString(om.OutcomeMeasureTitle)
	builder.WriteString(", outcome_measure_description=")
	builder.WriteString(om.OutcomeMeasureDescription)
	builder.WriteString(", outcome_measure_population_description=")
	builder.WriteString(om.OutcomeMeasurePopulationDescription)
	builder.WriteString(", outcome_measure_reporting_status=")
	builder.WriteString(om.OutcomeMeasureReportingStatus)
	builder.WriteString(", outcome_measure_anticipated_posting_date=")
	builder.WriteString(om.OutcomeMeasureAnticipatedPostingDate)
	builder.WriteString(", outcome_measure_param_type=")
	builder.WriteString(om.OutcomeMeasureParamType)
	builder.WriteString(", outcome_measure_dispersion_type=")
	builder.WriteString(om.OutcomeMeasureDispersionType)
	builder.WriteString(", outcome_measure_unit_of_measure=")
	builder.WriteString(om.OutcomeMeasureUnitOfMeasure)
	builder.WriteString(", outcome_measure_calculate_pct=")
	builder.WriteString(om.OutcomeMeasureCalculatePct)
	builder.WriteString(", outcome_measure_time_frame=")
	builder.WriteString(om.OutcomeMeasureTimeFrame)
	builder.WriteString(", outcome_measure_type_units_analyzed=")
	builder.WriteString(om.OutcomeMeasureTypeUnitsAnalyzed)
	builder.WriteString(", outcome_measure_denom_units_selected=")
	builder.WriteString(om.OutcomeMeasureDenomUnitsSelected)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeMeasures is a parsable slice of OutcomeMeasure.
type OutcomeMeasures []*OutcomeMeasure

func (om OutcomeMeasures) config(cfg config) {
	for _i := range om {
		om[_i].config = cfg
	}
}
