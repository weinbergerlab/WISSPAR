// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
)

// FlowGroup is the model entity for the FlowGroup schema.
type FlowGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowGroupID holds the value of the "flow_group_id" field.
	FlowGroupID string `json:"flow_group_id,omitempty"`
	// FlowGroupTitle holds the value of the "flow_group_title" field.
	FlowGroupTitle string `json:"flow_group_title,omitempty"`
	// FlowGroupDescription holds the value of the "flow_group_description" field.
	FlowGroupDescription string `json:"flow_group_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowGroupQuery when eager-loading is set.
	Edges                                   FlowGroupEdges `json:"edges"`
	participant_flow_module_flow_group_list *int
}

// FlowGroupEdges holds the relations/edges for other nodes in the graph.
type FlowGroupEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ParticipantFlowModule `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowGroupEdges) ParentOrErr() (*ParticipantFlowModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: participantflowmodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowgroup.FieldID:
			values[i] = new(sql.NullInt64)
		case flowgroup.FieldFlowGroupID, flowgroup.FieldFlowGroupTitle, flowgroup.FieldFlowGroupDescription:
			values[i] = new(sql.NullString)
		case flowgroup.ForeignKeys[0]: // participant_flow_module_flow_group_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowGroup fields.
func (fg *FlowGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fg.ID = int(value.Int64)
		case flowgroup.FieldFlowGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_group_id", values[i])
			} else if value.Valid {
				fg.FlowGroupID = value.String
			}
		case flowgroup.FieldFlowGroupTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_group_title", values[i])
			} else if value.Valid {
				fg.FlowGroupTitle = value.String
			}
		case flowgroup.FieldFlowGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_group_description", values[i])
			} else if value.Valid {
				fg.FlowGroupDescription = value.String
			}
		case flowgroup.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field participant_flow_module_flow_group_list", value)
			} else if value.Valid {
				fg.participant_flow_module_flow_group_list = new(int)
				*fg.participant_flow_module_flow_group_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowGroup entity.
func (fg *FlowGroup) QueryParent() *ParticipantFlowModuleQuery {
	return (&FlowGroupClient{config: fg.config}).QueryParent(fg)
}

// Update returns a builder for updating this FlowGroup.
// Note that you need to call FlowGroup.Unwrap() before calling this method if this FlowGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (fg *FlowGroup) Update() *FlowGroupUpdateOne {
	return (&FlowGroupClient{config: fg.config}).UpdateOne(fg)
}

// Unwrap unwraps the FlowGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fg *FlowGroup) Unwrap() *FlowGroup {
	tx, ok := fg.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowGroup is not a transactional entity")
	}
	fg.config.driver = tx.drv
	return fg
}

// String implements the fmt.Stringer.
func (fg *FlowGroup) String() string {
	var builder strings.Builder
	builder.WriteString("FlowGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", fg.ID))
	builder.WriteString(", flow_group_id=")
	builder.WriteString(fg.FlowGroupID)
	builder.WriteString(", flow_group_title=")
	builder.WriteString(fg.FlowGroupTitle)
	builder.WriteString(", flow_group_description=")
	builder.WriteString(fg.FlowGroupDescription)
	builder.WriteByte(')')
	return builder.String()
}

// FlowGroups is a parsable slice of FlowGroup.
type FlowGroups []*FlowGroup

func (fg FlowGroups) config(cfg config) {
	for _i := range fg {
		fg[_i].config = cfg
	}
}
