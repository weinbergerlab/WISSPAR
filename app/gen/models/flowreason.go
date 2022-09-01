// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
)

// FlowReason is the model entity for the FlowReason schema.
type FlowReason struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowReasonGroupID holds the value of the "flow_reason_group_id" field.
	FlowReasonGroupID string `json:"flow_reason_group_id,omitempty"`
	// FlowReasonComment holds the value of the "flow_reason_comment" field.
	FlowReasonComment string `json:"flow_reason_comment,omitempty"`
	// FlowReasonNumSubjects holds the value of the "flow_reason_num_subjects" field.
	FlowReasonNumSubjects string `json:"flow_reason_num_subjects,omitempty"`
	// FlowReasonNumUnits holds the value of the "flow_reason_num_units" field.
	FlowReasonNumUnits string `json:"flow_reason_num_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowReasonQuery when eager-loading is set.
	Edges                               FlowReasonEdges `json:"edges"`
	flow_drop_withdraw_flow_reason_list *int
}

// FlowReasonEdges holds the relations/edges for other nodes in the graph.
type FlowReasonEdges struct {
	// Parent holds the value of the parent edge.
	Parent *FlowDropWithdraw `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowReasonEdges) ParentOrErr() (*FlowDropWithdraw, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flowdropwithdraw.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowReason) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowreason.FieldID:
			values[i] = new(sql.NullInt64)
		case flowreason.FieldFlowReasonGroupID, flowreason.FieldFlowReasonComment, flowreason.FieldFlowReasonNumSubjects, flowreason.FieldFlowReasonNumUnits:
			values[i] = new(sql.NullString)
		case flowreason.ForeignKeys[0]: // flow_drop_withdraw_flow_reason_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowReason", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowReason fields.
func (fr *FlowReason) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowreason.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fr.ID = int(value.Int64)
		case flowreason.FieldFlowReasonGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_reason_group_id", values[i])
			} else if value.Valid {
				fr.FlowReasonGroupID = value.String
			}
		case flowreason.FieldFlowReasonComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_reason_comment", values[i])
			} else if value.Valid {
				fr.FlowReasonComment = value.String
			}
		case flowreason.FieldFlowReasonNumSubjects:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_reason_num_subjects", values[i])
			} else if value.Valid {
				fr.FlowReasonNumSubjects = value.String
			}
		case flowreason.FieldFlowReasonNumUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_reason_num_units", values[i])
			} else if value.Valid {
				fr.FlowReasonNumUnits = value.String
			}
		case flowreason.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field flow_drop_withdraw_flow_reason_list", value)
			} else if value.Valid {
				fr.flow_drop_withdraw_flow_reason_list = new(int)
				*fr.flow_drop_withdraw_flow_reason_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowReason entity.
func (fr *FlowReason) QueryParent() *FlowDropWithdrawQuery {
	return (&FlowReasonClient{config: fr.config}).QueryParent(fr)
}

// Update returns a builder for updating this FlowReason.
// Note that you need to call FlowReason.Unwrap() before calling this method if this FlowReason
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FlowReason) Update() *FlowReasonUpdateOne {
	return (&FlowReasonClient{config: fr.config}).UpdateOne(fr)
}

// Unwrap unwraps the FlowReason entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FlowReason) Unwrap() *FlowReason {
	tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowReason is not a transactional entity")
	}
	fr.config.driver = tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FlowReason) String() string {
	var builder strings.Builder
	builder.WriteString("FlowReason(")
	builder.WriteString(fmt.Sprintf("id=%v", fr.ID))
	builder.WriteString(", flow_reason_group_id=")
	builder.WriteString(fr.FlowReasonGroupID)
	builder.WriteString(", flow_reason_comment=")
	builder.WriteString(fr.FlowReasonComment)
	builder.WriteString(", flow_reason_num_subjects=")
	builder.WriteString(fr.FlowReasonNumSubjects)
	builder.WriteString(", flow_reason_num_units=")
	builder.WriteString(fr.FlowReasonNumUnits)
	builder.WriteByte(')')
	return builder.String()
}

// FlowReasons is a parsable slice of FlowReason.
type FlowReasons []*FlowReason

func (fr FlowReasons) config(cfg config) {
	for _i := range fr {
		fr[_i].config = cfg
	}
}
