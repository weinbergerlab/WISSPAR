// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
)

// OtherEvent is the model entity for the OtherEvent schema.
type OtherEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OtherEventTerm holds the value of the "other_event_term" field.
	OtherEventTerm string `json:"other_event_term,omitempty"`
	// OtherEventOrganSystem holds the value of the "other_event_organ_system" field.
	OtherEventOrganSystem string `json:"other_event_organ_system,omitempty"`
	// OtherEventSourceVocabulary holds the value of the "other_event_source_vocabulary" field.
	OtherEventSourceVocabulary string `json:"other_event_source_vocabulary,omitempty"`
	// OtherEventAssessmentType holds the value of the "other_event_assessment_type" field.
	OtherEventAssessmentType string `json:"other_event_assessment_type,omitempty"`
	// OtherEventNotes holds the value of the "other_event_notes" field.
	OtherEventNotes string `json:"other_event_notes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OtherEventQuery when eager-loading is set.
	Edges                                  OtherEventEdges `json:"edges"`
	adverse_events_module_other_event_list *int
}

// OtherEventEdges holds the relations/edges for other nodes in the graph.
type OtherEventEdges struct {
	// Parent holds the value of the parent edge.
	Parent *AdverseEventsModule `json:"parent,omitempty"`
	// OtherEventStatsList holds the value of the other_event_stats_list edge.
	OtherEventStatsList []*OtherEventStats `json:"other_event_stats_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OtherEventEdges) ParentOrErr() (*AdverseEventsModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: adverseeventsmodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// OtherEventStatsListOrErr returns the OtherEventStatsList value or an error if the edge
// was not loaded in eager-loading.
func (e OtherEventEdges) OtherEventStatsListOrErr() ([]*OtherEventStats, error) {
	if e.loadedTypes[1] {
		return e.OtherEventStatsList, nil
	}
	return nil, &NotLoadedError{edge: "other_event_stats_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OtherEvent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case otherevent.FieldID:
			values[i] = new(sql.NullInt64)
		case otherevent.FieldOtherEventTerm, otherevent.FieldOtherEventOrganSystem, otherevent.FieldOtherEventSourceVocabulary, otherevent.FieldOtherEventAssessmentType, otherevent.FieldOtherEventNotes:
			values[i] = new(sql.NullString)
		case otherevent.ForeignKeys[0]: // adverse_events_module_other_event_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OtherEvent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OtherEvent fields.
func (oe *OtherEvent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case otherevent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oe.ID = int(value.Int64)
		case otherevent.FieldOtherEventTerm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_term", values[i])
			} else if value.Valid {
				oe.OtherEventTerm = value.String
			}
		case otherevent.FieldOtherEventOrganSystem:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_organ_system", values[i])
			} else if value.Valid {
				oe.OtherEventOrganSystem = value.String
			}
		case otherevent.FieldOtherEventSourceVocabulary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_source_vocabulary", values[i])
			} else if value.Valid {
				oe.OtherEventSourceVocabulary = value.String
			}
		case otherevent.FieldOtherEventAssessmentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_assessment_type", values[i])
			} else if value.Valid {
				oe.OtherEventAssessmentType = value.String
			}
		case otherevent.FieldOtherEventNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_notes", values[i])
			} else if value.Valid {
				oe.OtherEventNotes = value.String
			}
		case otherevent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field adverse_events_module_other_event_list", value)
			} else if value.Valid {
				oe.adverse_events_module_other_event_list = new(int)
				*oe.adverse_events_module_other_event_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OtherEvent entity.
func (oe *OtherEvent) QueryParent() *AdverseEventsModuleQuery {
	return (&OtherEventClient{config: oe.config}).QueryParent(oe)
}

// QueryOtherEventStatsList queries the "other_event_stats_list" edge of the OtherEvent entity.
func (oe *OtherEvent) QueryOtherEventStatsList() *OtherEventStatsQuery {
	return (&OtherEventClient{config: oe.config}).QueryOtherEventStatsList(oe)
}

// Update returns a builder for updating this OtherEvent.
// Note that you need to call OtherEvent.Unwrap() before calling this method if this OtherEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (oe *OtherEvent) Update() *OtherEventUpdateOne {
	return (&OtherEventClient{config: oe.config}).UpdateOne(oe)
}

// Unwrap unwraps the OtherEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oe *OtherEvent) Unwrap() *OtherEvent {
	tx, ok := oe.config.driver.(*txDriver)
	if !ok {
		panic("models: OtherEvent is not a transactional entity")
	}
	oe.config.driver = tx.drv
	return oe
}

// String implements the fmt.Stringer.
func (oe *OtherEvent) String() string {
	var builder strings.Builder
	builder.WriteString("OtherEvent(")
	builder.WriteString(fmt.Sprintf("id=%v", oe.ID))
	builder.WriteString(", other_event_term=")
	builder.WriteString(oe.OtherEventTerm)
	builder.WriteString(", other_event_organ_system=")
	builder.WriteString(oe.OtherEventOrganSystem)
	builder.WriteString(", other_event_source_vocabulary=")
	builder.WriteString(oe.OtherEventSourceVocabulary)
	builder.WriteString(", other_event_assessment_type=")
	builder.WriteString(oe.OtherEventAssessmentType)
	builder.WriteString(", other_event_notes=")
	builder.WriteString(oe.OtherEventNotes)
	builder.WriteByte(')')
	return builder.String()
}

// OtherEvents is a parsable slice of OtherEvent.
type OtherEvents []*OtherEvent

func (oe OtherEvents) config(cfg config) {
	for _i := range oe {
		oe[_i].config = cfg
	}
}
