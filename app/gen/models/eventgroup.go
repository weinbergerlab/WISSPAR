// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
)

// EventGroup is the model entity for the EventGroup schema.
type EventGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EventGroupID holds the value of the "event_group_id" field.
	EventGroupID string `json:"event_group_id,omitempty"`
	// EventGroupTitle holds the value of the "event_group_title" field.
	EventGroupTitle string `json:"event_group_title,omitempty"`
	// EventGroupDescription holds the value of the "event_group_description" field.
	EventGroupDescription string `json:"event_group_description,omitempty"`
	// EventGroupDeathsNumAffected holds the value of the "event_group_deaths_num_affected" field.
	EventGroupDeathsNumAffected string `json:"event_group_deaths_num_affected,omitempty"`
	// EventGroupDeathsNumAtRisk holds the value of the "event_group_deaths_num_at_risk" field.
	EventGroupDeathsNumAtRisk string `json:"event_group_deaths_num_at_risk,omitempty"`
	// EventGroupSeriousNumAffected holds the value of the "event_group_serious_num_affected" field.
	EventGroupSeriousNumAffected string `json:"event_group_serious_num_affected,omitempty"`
	// EventGroupSeriousNumAtRisk holds the value of the "event_group_serious_num_at_risk" field.
	EventGroupSeriousNumAtRisk string `json:"event_group_serious_num_at_risk,omitempty"`
	// EventGroupOtherNumAffected holds the value of the "event_group_other_num_affected" field.
	EventGroupOtherNumAffected string `json:"event_group_other_num_affected,omitempty"`
	// EventGroupOtherNumAtRisk holds the value of the "event_group_other_num_at_risk" field.
	EventGroupOtherNumAtRisk string `json:"event_group_other_num_at_risk,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventGroupQuery when eager-loading is set.
	Edges                                  EventGroupEdges `json:"edges"`
	adverse_events_module_event_group_list *int
}

// EventGroupEdges holds the relations/edges for other nodes in the graph.
type EventGroupEdges struct {
	// Parent holds the value of the parent edge.
	Parent *AdverseEventsModule `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventGroupEdges) ParentOrErr() (*AdverseEventsModule, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*EventGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventgroup.FieldID:
			values[i] = new(sql.NullInt64)
		case eventgroup.FieldEventGroupID, eventgroup.FieldEventGroupTitle, eventgroup.FieldEventGroupDescription, eventgroup.FieldEventGroupDeathsNumAffected, eventgroup.FieldEventGroupDeathsNumAtRisk, eventgroup.FieldEventGroupSeriousNumAffected, eventgroup.FieldEventGroupSeriousNumAtRisk, eventgroup.FieldEventGroupOtherNumAffected, eventgroup.FieldEventGroupOtherNumAtRisk:
			values[i] = new(sql.NullString)
		case eventgroup.ForeignKeys[0]: // adverse_events_module_event_group_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EventGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventGroup fields.
func (eg *EventGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			eg.ID = int(value.Int64)
		case eventgroup.FieldEventGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_id", values[i])
			} else if value.Valid {
				eg.EventGroupID = value.String
			}
		case eventgroup.FieldEventGroupTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_title", values[i])
			} else if value.Valid {
				eg.EventGroupTitle = value.String
			}
		case eventgroup.FieldEventGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_description", values[i])
			} else if value.Valid {
				eg.EventGroupDescription = value.String
			}
		case eventgroup.FieldEventGroupDeathsNumAffected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_deaths_num_affected", values[i])
			} else if value.Valid {
				eg.EventGroupDeathsNumAffected = value.String
			}
		case eventgroup.FieldEventGroupDeathsNumAtRisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_deaths_num_at_risk", values[i])
			} else if value.Valid {
				eg.EventGroupDeathsNumAtRisk = value.String
			}
		case eventgroup.FieldEventGroupSeriousNumAffected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_serious_num_affected", values[i])
			} else if value.Valid {
				eg.EventGroupSeriousNumAffected = value.String
			}
		case eventgroup.FieldEventGroupSeriousNumAtRisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_serious_num_at_risk", values[i])
			} else if value.Valid {
				eg.EventGroupSeriousNumAtRisk = value.String
			}
		case eventgroup.FieldEventGroupOtherNumAffected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_other_num_affected", values[i])
			} else if value.Valid {
				eg.EventGroupOtherNumAffected = value.String
			}
		case eventgroup.FieldEventGroupOtherNumAtRisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_group_other_num_at_risk", values[i])
			} else if value.Valid {
				eg.EventGroupOtherNumAtRisk = value.String
			}
		case eventgroup.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field adverse_events_module_event_group_list", value)
			} else if value.Valid {
				eg.adverse_events_module_event_group_list = new(int)
				*eg.adverse_events_module_event_group_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the EventGroup entity.
func (eg *EventGroup) QueryParent() *AdverseEventsModuleQuery {
	return (&EventGroupClient{config: eg.config}).QueryParent(eg)
}

// Update returns a builder for updating this EventGroup.
// Note that you need to call EventGroup.Unwrap() before calling this method if this EventGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (eg *EventGroup) Update() *EventGroupUpdateOne {
	return (&EventGroupClient{config: eg.config}).UpdateOne(eg)
}

// Unwrap unwraps the EventGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eg *EventGroup) Unwrap() *EventGroup {
	tx, ok := eg.config.driver.(*txDriver)
	if !ok {
		panic("models: EventGroup is not a transactional entity")
	}
	eg.config.driver = tx.drv
	return eg
}

// String implements the fmt.Stringer.
func (eg *EventGroup) String() string {
	var builder strings.Builder
	builder.WriteString("EventGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", eg.ID))
	builder.WriteString(", event_group_id=")
	builder.WriteString(eg.EventGroupID)
	builder.WriteString(", event_group_title=")
	builder.WriteString(eg.EventGroupTitle)
	builder.WriteString(", event_group_description=")
	builder.WriteString(eg.EventGroupDescription)
	builder.WriteString(", event_group_deaths_num_affected=")
	builder.WriteString(eg.EventGroupDeathsNumAffected)
	builder.WriteString(", event_group_deaths_num_at_risk=")
	builder.WriteString(eg.EventGroupDeathsNumAtRisk)
	builder.WriteString(", event_group_serious_num_affected=")
	builder.WriteString(eg.EventGroupSeriousNumAffected)
	builder.WriteString(", event_group_serious_num_at_risk=")
	builder.WriteString(eg.EventGroupSeriousNumAtRisk)
	builder.WriteString(", event_group_other_num_affected=")
	builder.WriteString(eg.EventGroupOtherNumAffected)
	builder.WriteString(", event_group_other_num_at_risk=")
	builder.WriteString(eg.EventGroupOtherNumAtRisk)
	builder.WriteByte(')')
	return builder.String()
}

// EventGroups is a parsable slice of EventGroup.
type EventGroups []*EventGroup

func (eg EventGroups) config(cfg config) {
	for _i := range eg {
		eg[_i].config = cfg
	}
}
