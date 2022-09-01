// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// AdverseEventsModule is the model entity for the AdverseEventsModule schema.
type AdverseEventsModule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EventsFrequencyThreshold holds the value of the "events_frequency_threshold" field.
	EventsFrequencyThreshold string `json:"events_frequency_threshold,omitempty"`
	// EventsTimeFrame holds the value of the "events_time_frame" field.
	EventsTimeFrame string `json:"events_time_frame,omitempty"`
	// EventsDescription holds the value of the "events_description" field.
	EventsDescription string `json:"events_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdverseEventsModuleQuery when eager-loading is set.
	Edges                                    AdverseEventsModuleEdges `json:"edges"`
	results_definition_adverse_events_module *int
}

// AdverseEventsModuleEdges holds the relations/edges for other nodes in the graph.
type AdverseEventsModuleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ResultsDefinition `json:"parent,omitempty"`
	// EventGroupList holds the value of the event_group_list edge.
	EventGroupList []*EventGroup `json:"event_group_list,omitempty"`
	// SeriousEventList holds the value of the serious_event_list edge.
	SeriousEventList []*SeriousEvent `json:"serious_event_list,omitempty"`
	// OtherEventList holds the value of the other_event_list edge.
	OtherEventList []*OtherEvent `json:"other_event_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdverseEventsModuleEdges) ParentOrErr() (*ResultsDefinition, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resultsdefinition.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// EventGroupListOrErr returns the EventGroupList value or an error if the edge
// was not loaded in eager-loading.
func (e AdverseEventsModuleEdges) EventGroupListOrErr() ([]*EventGroup, error) {
	if e.loadedTypes[1] {
		return e.EventGroupList, nil
	}
	return nil, &NotLoadedError{edge: "event_group_list"}
}

// SeriousEventListOrErr returns the SeriousEventList value or an error if the edge
// was not loaded in eager-loading.
func (e AdverseEventsModuleEdges) SeriousEventListOrErr() ([]*SeriousEvent, error) {
	if e.loadedTypes[2] {
		return e.SeriousEventList, nil
	}
	return nil, &NotLoadedError{edge: "serious_event_list"}
}

// OtherEventListOrErr returns the OtherEventList value or an error if the edge
// was not loaded in eager-loading.
func (e AdverseEventsModuleEdges) OtherEventListOrErr() ([]*OtherEvent, error) {
	if e.loadedTypes[3] {
		return e.OtherEventList, nil
	}
	return nil, &NotLoadedError{edge: "other_event_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdverseEventsModule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case adverseeventsmodule.FieldID:
			values[i] = new(sql.NullInt64)
		case adverseeventsmodule.FieldEventsFrequencyThreshold, adverseeventsmodule.FieldEventsTimeFrame, adverseeventsmodule.FieldEventsDescription:
			values[i] = new(sql.NullString)
		case adverseeventsmodule.ForeignKeys[0]: // results_definition_adverse_events_module
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdverseEventsModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdverseEventsModule fields.
func (aem *AdverseEventsModule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adverseeventsmodule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aem.ID = int(value.Int64)
		case adverseeventsmodule.FieldEventsFrequencyThreshold:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field events_frequency_threshold", values[i])
			} else if value.Valid {
				aem.EventsFrequencyThreshold = value.String
			}
		case adverseeventsmodule.FieldEventsTimeFrame:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field events_time_frame", values[i])
			} else if value.Valid {
				aem.EventsTimeFrame = value.String
			}
		case adverseeventsmodule.FieldEventsDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field events_description", values[i])
			} else if value.Valid {
				aem.EventsDescription = value.String
			}
		case adverseeventsmodule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field results_definition_adverse_events_module", value)
			} else if value.Valid {
				aem.results_definition_adverse_events_module = new(int)
				*aem.results_definition_adverse_events_module = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the AdverseEventsModule entity.
func (aem *AdverseEventsModule) QueryParent() *ResultsDefinitionQuery {
	return (&AdverseEventsModuleClient{config: aem.config}).QueryParent(aem)
}

// QueryEventGroupList queries the "event_group_list" edge of the AdverseEventsModule entity.
func (aem *AdverseEventsModule) QueryEventGroupList() *EventGroupQuery {
	return (&AdverseEventsModuleClient{config: aem.config}).QueryEventGroupList(aem)
}

// QuerySeriousEventList queries the "serious_event_list" edge of the AdverseEventsModule entity.
func (aem *AdverseEventsModule) QuerySeriousEventList() *SeriousEventQuery {
	return (&AdverseEventsModuleClient{config: aem.config}).QuerySeriousEventList(aem)
}

// QueryOtherEventList queries the "other_event_list" edge of the AdverseEventsModule entity.
func (aem *AdverseEventsModule) QueryOtherEventList() *OtherEventQuery {
	return (&AdverseEventsModuleClient{config: aem.config}).QueryOtherEventList(aem)
}

// Update returns a builder for updating this AdverseEventsModule.
// Note that you need to call AdverseEventsModule.Unwrap() before calling this method if this AdverseEventsModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (aem *AdverseEventsModule) Update() *AdverseEventsModuleUpdateOne {
	return (&AdverseEventsModuleClient{config: aem.config}).UpdateOne(aem)
}

// Unwrap unwraps the AdverseEventsModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aem *AdverseEventsModule) Unwrap() *AdverseEventsModule {
	tx, ok := aem.config.driver.(*txDriver)
	if !ok {
		panic("models: AdverseEventsModule is not a transactional entity")
	}
	aem.config.driver = tx.drv
	return aem
}

// String implements the fmt.Stringer.
func (aem *AdverseEventsModule) String() string {
	var builder strings.Builder
	builder.WriteString("AdverseEventsModule(")
	builder.WriteString(fmt.Sprintf("id=%v", aem.ID))
	builder.WriteString(", events_frequency_threshold=")
	builder.WriteString(aem.EventsFrequencyThreshold)
	builder.WriteString(", events_time_frame=")
	builder.WriteString(aem.EventsTimeFrame)
	builder.WriteString(", events_description=")
	builder.WriteString(aem.EventsDescription)
	builder.WriteByte(')')
	return builder.String()
}

// AdverseEventsModules is a parsable slice of AdverseEventsModule.
type AdverseEventsModules []*AdverseEventsModule

func (aem AdverseEventsModules) config(cfg config) {
	for _i := range aem {
		aem[_i].config = cfg
	}
}
