// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
)

// OutcomeOverview is the model entity for the OutcomeOverview schema.
type OutcomeOverview struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeOverviewID holds the value of the "outcome_overview_id" field.
	OutcomeOverviewID string `json:"outcome_overview_id,omitempty"`
	// OutcomeOverviewTitle holds the value of the "outcome_overview_title" field.
	OutcomeOverviewTitle string `json:"outcome_overview_title,omitempty"`
	// OutcomeOverviewDescription holds the value of the "outcome_overview_description" field.
	OutcomeOverviewDescription string `json:"outcome_overview_description,omitempty"`
	// OutcomeOverviewParticipants holds the value of the "outcome_overview_participants" field.
	OutcomeOverviewParticipants string `json:"outcome_overview_participants,omitempty"`
	// OutcomeOverviewTimeFrame holds the value of the "outcome_overview_time_frame" field.
	OutcomeOverviewTimeFrame string `json:"outcome_overview_time_frame,omitempty"`
	// OutcomeOverviewSerotype holds the value of the "outcome_overview_serotype" field.
	OutcomeOverviewSerotype string `json:"outcome_overview_serotype,omitempty"`
	// OutcomeOverviewAssay holds the value of the "outcome_overview_assay" field.
	OutcomeOverviewAssay string `json:"outcome_overview_assay,omitempty"`
	// OutcomeOverviewDoseNumber holds the value of the "outcome_overview_dose_number" field.
	OutcomeOverviewDoseNumber int64 `json:"outcome_overview_dose_number,omitempty"`
	// OutcomeOverviewValue holds the value of the "outcome_overview_value" field.
	OutcomeOverviewValue float64 `json:"outcome_overview_value,omitempty"`
	// OutcomeOverviewUpperLimit holds the value of the "outcome_overview_upper_limit" field.
	OutcomeOverviewUpperLimit string `json:"outcome_overview_upper_limit,omitempty"`
	// OutcomeOverviewLowerLimit holds the value of the "outcome_overview_lower_limit" field.
	OutcomeOverviewLowerLimit string `json:"outcome_overview_lower_limit,omitempty"`
	// OutcomeOverviewGroupID holds the value of the "outcome_overview_group_id" field.
	OutcomeOverviewGroupID string `json:"outcome_overview_group_id,omitempty"`
	// OutcomeOverviewRatio holds the value of the "outcome_overview_ratio" field.
	OutcomeOverviewRatio string `json:"outcome_overview_ratio,omitempty"`
	// OutcomeOverviewMeasureTitle holds the value of the "outcome_overview_measure_title" field.
	OutcomeOverviewMeasureTitle string `json:"outcome_overview_measure_title,omitempty"`
	// OutcomeOverviewVaccine holds the value of the "outcome_overview_vaccine" field.
	OutcomeOverviewVaccine string `json:"outcome_overview_vaccine,omitempty"`
	// OutcomeOverviewImmunocompromisedPopulation holds the value of the "outcome_overview_immunocompromised_population" field.
	OutcomeOverviewImmunocompromisedPopulation string `json:"outcome_overview_immunocompromised_population,omitempty"`
	// OutcomeOverviewManufacturer holds the value of the "outcome_overview_manufacturer" field.
	OutcomeOverviewManufacturer string `json:"outcome_overview_manufacturer,omitempty"`
	// OutcomeOverviewConfidenceInterval holds the value of the "outcome_overview_confidence_interval" field.
	OutcomeOverviewConfidenceInterval string `json:"outcome_overview_confidence_interval,omitempty"`
	// OutcomeOverviewPercentResponders holds the value of the "outcome_overview_percent_responders" field.
	OutcomeOverviewPercentResponders float64 `json:"outcome_overview_percent_responders,omitempty"`
	// OutcomeOverviewTimeFrameWeeks holds the value of the "outcome_overview_time_frame_weeks" field.
	OutcomeOverviewTimeFrameWeeks int64 `json:"outcome_overview_time_frame_weeks,omitempty"`
	// OutcomeOverviewDoseDescription holds the value of the "outcome_overview_dose_description" field.
	OutcomeOverviewDoseDescription string `json:"outcome_overview_dose_description,omitempty"`
	// OutcomeOverviewSchedule holds the value of the "outcome_overview_schedule" field.
	OutcomeOverviewSchedule string `json:"outcome_overview_schedule,omitempty"`
	// OutcomeOverviewUseCaseCode holds the value of the "outcome_overview_use_case_code" field.
	OutcomeOverviewUseCaseCode string `json:"outcome_overview_use_case_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeOverviewQuery when eager-loading is set.
	Edges                                 OutcomeOverviewEdges `json:"edges"`
	outcome_measure_outcome_overview_list *int
}

// OutcomeOverviewEdges holds the relations/edges for other nodes in the graph.
type OutcomeOverviewEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasure `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeOverviewEdges) ParentOrErr() (*OutcomeMeasure, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeOverview) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeoverview.FieldOutcomeOverviewValue, outcomeoverview.FieldOutcomeOverviewPercentResponders:
			values[i] = new(sql.NullFloat64)
		case outcomeoverview.FieldID, outcomeoverview.FieldOutcomeOverviewDoseNumber, outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks:
			values[i] = new(sql.NullInt64)
		case outcomeoverview.FieldOutcomeOverviewID, outcomeoverview.FieldOutcomeOverviewTitle, outcomeoverview.FieldOutcomeOverviewDescription, outcomeoverview.FieldOutcomeOverviewParticipants, outcomeoverview.FieldOutcomeOverviewTimeFrame, outcomeoverview.FieldOutcomeOverviewSerotype, outcomeoverview.FieldOutcomeOverviewAssay, outcomeoverview.FieldOutcomeOverviewUpperLimit, outcomeoverview.FieldOutcomeOverviewLowerLimit, outcomeoverview.FieldOutcomeOverviewGroupID, outcomeoverview.FieldOutcomeOverviewRatio, outcomeoverview.FieldOutcomeOverviewMeasureTitle, outcomeoverview.FieldOutcomeOverviewVaccine, outcomeoverview.FieldOutcomeOverviewImmunocompromisedPopulation, outcomeoverview.FieldOutcomeOverviewManufacturer, outcomeoverview.FieldOutcomeOverviewConfidenceInterval, outcomeoverview.FieldOutcomeOverviewDoseDescription, outcomeoverview.FieldOutcomeOverviewSchedule, outcomeoverview.FieldOutcomeOverviewUseCaseCode:
			values[i] = new(sql.NullString)
		case outcomeoverview.ForeignKeys[0]: // outcome_measure_outcome_overview_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeOverview", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeOverview fields.
func (oo *OutcomeOverview) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeoverview.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oo.ID = int(value.Int64)
		case outcomeoverview.FieldOutcomeOverviewID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_id", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewID = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_title", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewTitle = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_description", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewDescription = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewParticipants:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_participants", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewParticipants = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewTimeFrame:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_time_frame", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewTimeFrame = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewSerotype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_serotype", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewSerotype = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewAssay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_assay", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewAssay = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewDoseNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_dose_number", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewDoseNumber = value.Int64
			}
		case outcomeoverview.FieldOutcomeOverviewValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_value", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewValue = value.Float64
			}
		case outcomeoverview.FieldOutcomeOverviewUpperLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_upper_limit", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewUpperLimit = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewLowerLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_lower_limit", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewLowerLimit = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_group_id", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewGroupID = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewRatio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_ratio", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewRatio = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewMeasureTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_measure_title", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewMeasureTitle = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewVaccine:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_vaccine", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewVaccine = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewImmunocompromisedPopulation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_immunocompromised_population", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewImmunocompromisedPopulation = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_manufacturer", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewManufacturer = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewConfidenceInterval:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_confidence_interval", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewConfidenceInterval = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewPercentResponders:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_percent_responders", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewPercentResponders = value.Float64
			}
		case outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_time_frame_weeks", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewTimeFrameWeeks = value.Int64
			}
		case outcomeoverview.FieldOutcomeOverviewDoseDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_dose_description", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewDoseDescription = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewSchedule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_schedule", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewSchedule = value.String
			}
		case outcomeoverview.FieldOutcomeOverviewUseCaseCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_overview_use_case_code", values[i])
			} else if value.Valid {
				oo.OutcomeOverviewUseCaseCode = value.String
			}
		case outcomeoverview.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measure_outcome_overview_list", value)
			} else if value.Valid {
				oo.outcome_measure_outcome_overview_list = new(int)
				*oo.outcome_measure_outcome_overview_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeOverview entity.
func (oo *OutcomeOverview) QueryParent() *OutcomeMeasureQuery {
	return (&OutcomeOverviewClient{config: oo.config}).QueryParent(oo)
}

// Update returns a builder for updating this OutcomeOverview.
// Note that you need to call OutcomeOverview.Unwrap() before calling this method if this OutcomeOverview
// was returned from a transaction, and the transaction was committed or rolled back.
func (oo *OutcomeOverview) Update() *OutcomeOverviewUpdateOne {
	return (&OutcomeOverviewClient{config: oo.config}).UpdateOne(oo)
}

// Unwrap unwraps the OutcomeOverview entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oo *OutcomeOverview) Unwrap() *OutcomeOverview {
	tx, ok := oo.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeOverview is not a transactional entity")
	}
	oo.config.driver = tx.drv
	return oo
}

// String implements the fmt.Stringer.
func (oo *OutcomeOverview) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeOverview(")
	builder.WriteString(fmt.Sprintf("id=%v", oo.ID))
	builder.WriteString(", outcome_overview_id=")
	builder.WriteString(oo.OutcomeOverviewID)
	builder.WriteString(", outcome_overview_title=")
	builder.WriteString(oo.OutcomeOverviewTitle)
	builder.WriteString(", outcome_overview_description=")
	builder.WriteString(oo.OutcomeOverviewDescription)
	builder.WriteString(", outcome_overview_participants=")
	builder.WriteString(oo.OutcomeOverviewParticipants)
	builder.WriteString(", outcome_overview_time_frame=")
	builder.WriteString(oo.OutcomeOverviewTimeFrame)
	builder.WriteString(", outcome_overview_serotype=")
	builder.WriteString(oo.OutcomeOverviewSerotype)
	builder.WriteString(", outcome_overview_assay=")
	builder.WriteString(oo.OutcomeOverviewAssay)
	builder.WriteString(", outcome_overview_dose_number=")
	builder.WriteString(fmt.Sprintf("%v", oo.OutcomeOverviewDoseNumber))
	builder.WriteString(", outcome_overview_value=")
	builder.WriteString(fmt.Sprintf("%v", oo.OutcomeOverviewValue))
	builder.WriteString(", outcome_overview_upper_limit=")
	builder.WriteString(oo.OutcomeOverviewUpperLimit)
	builder.WriteString(", outcome_overview_lower_limit=")
	builder.WriteString(oo.OutcomeOverviewLowerLimit)
	builder.WriteString(", outcome_overview_group_id=")
	builder.WriteString(oo.OutcomeOverviewGroupID)
	builder.WriteString(", outcome_overview_ratio=")
	builder.WriteString(oo.OutcomeOverviewRatio)
	builder.WriteString(", outcome_overview_measure_title=")
	builder.WriteString(oo.OutcomeOverviewMeasureTitle)
	builder.WriteString(", outcome_overview_vaccine=")
	builder.WriteString(oo.OutcomeOverviewVaccine)
	builder.WriteString(", outcome_overview_immunocompromised_population=")
	builder.WriteString(oo.OutcomeOverviewImmunocompromisedPopulation)
	builder.WriteString(", outcome_overview_manufacturer=")
	builder.WriteString(oo.OutcomeOverviewManufacturer)
	builder.WriteString(", outcome_overview_confidence_interval=")
	builder.WriteString(oo.OutcomeOverviewConfidenceInterval)
	builder.WriteString(", outcome_overview_percent_responders=")
	builder.WriteString(fmt.Sprintf("%v", oo.OutcomeOverviewPercentResponders))
	builder.WriteString(", outcome_overview_time_frame_weeks=")
	builder.WriteString(fmt.Sprintf("%v", oo.OutcomeOverviewTimeFrameWeeks))
	builder.WriteString(", outcome_overview_dose_description=")
	builder.WriteString(oo.OutcomeOverviewDoseDescription)
	builder.WriteString(", outcome_overview_schedule=")
	builder.WriteString(oo.OutcomeOverviewSchedule)
	builder.WriteString(", outcome_overview_use_case_code=")
	builder.WriteString(oo.OutcomeOverviewUseCaseCode)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeOverviews is a parsable slice of OutcomeOverview.
type OutcomeOverviews []*OutcomeOverview

func (oo OutcomeOverviews) config(cfg config) {
	for _i := range oo {
		oo[_i].config = cfg
	}
}
