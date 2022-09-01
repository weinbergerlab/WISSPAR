// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
)

// SeriousEvent is the model entity for the SeriousEvent schema.
type SeriousEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SeriousEventTerm holds the value of the "serious_event_term" field.
	SeriousEventTerm string `json:"serious_event_term,omitempty"`
	// SeriousEventOrganSystem holds the value of the "serious_event_organ_system" field.
	SeriousEventOrganSystem string `json:"serious_event_organ_system,omitempty"`
	// SeriousEventSourceVocabulary holds the value of the "serious_event_source_vocabulary" field.
	SeriousEventSourceVocabulary string `json:"serious_event_source_vocabulary,omitempty"`
	// SeriousEventAssessmentType holds the value of the "serious_event_assessment_type" field.
	SeriousEventAssessmentType string `json:"serious_event_assessment_type,omitempty"`
	// SeriousEventNotes holds the value of the "serious_event_notes" field.
	SeriousEventNotes string `json:"serious_event_notes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeriousEventQuery when eager-loading is set.
	Edges                                    SeriousEventEdges `json:"edges"`
	adverse_events_module_serious_event_list *int
}

// SeriousEventEdges holds the relations/edges for other nodes in the graph.
type SeriousEventEdges struct {
	// Parent holds the value of the parent edge.
	Parent *AdverseEventsModule `json:"parent,omitempty"`
	// SeriousEventStatsList holds the value of the serious_event_stats_list edge.
	SeriousEventStatsList []*SeriousEventStats `json:"serious_event_stats_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeriousEventEdges) ParentOrErr() (*AdverseEventsModule, error) {
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

// SeriousEventStatsListOrErr returns the SeriousEventStatsList value or an error if the edge
// was not loaded in eager-loading.
func (e SeriousEventEdges) SeriousEventStatsListOrErr() ([]*SeriousEventStats, error) {
	if e.loadedTypes[1] {
		return e.SeriousEventStatsList, nil
	}
	return nil, &NotLoadedError{edge: "serious_event_stats_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SeriousEvent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case seriousevent.FieldID:
			values[i] = new(sql.NullInt64)
		case seriousevent.FieldSeriousEventTerm, seriousevent.FieldSeriousEventOrganSystem, seriousevent.FieldSeriousEventSourceVocabulary, seriousevent.FieldSeriousEventAssessmentType, seriousevent.FieldSeriousEventNotes:
			values[i] = new(sql.NullString)
		case seriousevent.ForeignKeys[0]: // adverse_events_module_serious_event_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SeriousEvent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SeriousEvent fields.
func (se *SeriousEvent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seriousevent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			se.ID = int(value.Int64)
		case seriousevent.FieldSeriousEventTerm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_term", values[i])
			} else if value.Valid {
				se.SeriousEventTerm = value.String
			}
		case seriousevent.FieldSeriousEventOrganSystem:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_organ_system", values[i])
			} else if value.Valid {
				se.SeriousEventOrganSystem = value.String
			}
		case seriousevent.FieldSeriousEventSourceVocabulary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_source_vocabulary", values[i])
			} else if value.Valid {
				se.SeriousEventSourceVocabulary = value.String
			}
		case seriousevent.FieldSeriousEventAssessmentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_assessment_type", values[i])
			} else if value.Valid {
				se.SeriousEventAssessmentType = value.String
			}
		case seriousevent.FieldSeriousEventNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_notes", values[i])
			} else if value.Valid {
				se.SeriousEventNotes = value.String
			}
		case seriousevent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field adverse_events_module_serious_event_list", value)
			} else if value.Valid {
				se.adverse_events_module_serious_event_list = new(int)
				*se.adverse_events_module_serious_event_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the SeriousEvent entity.
func (se *SeriousEvent) QueryParent() *AdverseEventsModuleQuery {
	return (&SeriousEventClient{config: se.config}).QueryParent(se)
}

// QuerySeriousEventStatsList queries the "serious_event_stats_list" edge of the SeriousEvent entity.
func (se *SeriousEvent) QuerySeriousEventStatsList() *SeriousEventStatsQuery {
	return (&SeriousEventClient{config: se.config}).QuerySeriousEventStatsList(se)
}

// Update returns a builder for updating this SeriousEvent.
// Note that you need to call SeriousEvent.Unwrap() before calling this method if this SeriousEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (se *SeriousEvent) Update() *SeriousEventUpdateOne {
	return (&SeriousEventClient{config: se.config}).UpdateOne(se)
}

// Unwrap unwraps the SeriousEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (se *SeriousEvent) Unwrap() *SeriousEvent {
	tx, ok := se.config.driver.(*txDriver)
	if !ok {
		panic("models: SeriousEvent is not a transactional entity")
	}
	se.config.driver = tx.drv
	return se
}

// String implements the fmt.Stringer.
func (se *SeriousEvent) String() string {
	var builder strings.Builder
	builder.WriteString("SeriousEvent(")
	builder.WriteString(fmt.Sprintf("id=%v", se.ID))
	builder.WriteString(", serious_event_term=")
	builder.WriteString(se.SeriousEventTerm)
	builder.WriteString(", serious_event_organ_system=")
	builder.WriteString(se.SeriousEventOrganSystem)
	builder.WriteString(", serious_event_source_vocabulary=")
	builder.WriteString(se.SeriousEventSourceVocabulary)
	builder.WriteString(", serious_event_assessment_type=")
	builder.WriteString(se.SeriousEventAssessmentType)
	builder.WriteString(", serious_event_notes=")
	builder.WriteString(se.SeriousEventNotes)
	builder.WriteByte(')')
	return builder.String()
}

// SeriousEvents is a parsable slice of SeriousEvent.
type SeriousEvents []*SeriousEvent

func (se SeriousEvents) config(cfg config) {
	for _i := range se {
		se[_i].config = cfg
	}
}
