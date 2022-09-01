// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
)

// FlowDropWithdraw is the model entity for the FlowDropWithdraw schema.
type FlowDropWithdraw struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowDropWithdrawType holds the value of the "flow_drop_withdraw_type" field.
	FlowDropWithdrawType string `json:"flow_drop_withdraw_type,omitempty"`
	// FlowDropWithdrawComment holds the value of the "flow_drop_withdraw_comment" field.
	FlowDropWithdrawComment string `json:"flow_drop_withdraw_comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowDropWithdrawQuery when eager-loading is set.
	Edges                               FlowDropWithdrawEdges `json:"edges"`
	flow_period_flow_drop_withdraw_list *int
}

// FlowDropWithdrawEdges holds the relations/edges for other nodes in the graph.
type FlowDropWithdrawEdges struct {
	// Parent holds the value of the parent edge.
	Parent *FlowPeriod `json:"parent,omitempty"`
	// FlowReasonList holds the value of the flow_reason_list edge.
	FlowReasonList []*FlowReason `json:"flow_reason_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowDropWithdrawEdges) ParentOrErr() (*FlowPeriod, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flowperiod.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// FlowReasonListOrErr returns the FlowReasonList value or an error if the edge
// was not loaded in eager-loading.
func (e FlowDropWithdrawEdges) FlowReasonListOrErr() ([]*FlowReason, error) {
	if e.loadedTypes[1] {
		return e.FlowReasonList, nil
	}
	return nil, &NotLoadedError{edge: "flow_reason_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowDropWithdraw) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowdropwithdraw.FieldID:
			values[i] = new(sql.NullInt64)
		case flowdropwithdraw.FieldFlowDropWithdrawType, flowdropwithdraw.FieldFlowDropWithdrawComment:
			values[i] = new(sql.NullString)
		case flowdropwithdraw.ForeignKeys[0]: // flow_period_flow_drop_withdraw_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowDropWithdraw", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowDropWithdraw fields.
func (fdw *FlowDropWithdraw) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowdropwithdraw.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fdw.ID = int(value.Int64)
		case flowdropwithdraw.FieldFlowDropWithdrawType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_drop_withdraw_type", values[i])
			} else if value.Valid {
				fdw.FlowDropWithdrawType = value.String
			}
		case flowdropwithdraw.FieldFlowDropWithdrawComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_drop_withdraw_comment", values[i])
			} else if value.Valid {
				fdw.FlowDropWithdrawComment = value.String
			}
		case flowdropwithdraw.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field flow_period_flow_drop_withdraw_list", value)
			} else if value.Valid {
				fdw.flow_period_flow_drop_withdraw_list = new(int)
				*fdw.flow_period_flow_drop_withdraw_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowDropWithdraw entity.
func (fdw *FlowDropWithdraw) QueryParent() *FlowPeriodQuery {
	return (&FlowDropWithdrawClient{config: fdw.config}).QueryParent(fdw)
}

// QueryFlowReasonList queries the "flow_reason_list" edge of the FlowDropWithdraw entity.
func (fdw *FlowDropWithdraw) QueryFlowReasonList() *FlowReasonQuery {
	return (&FlowDropWithdrawClient{config: fdw.config}).QueryFlowReasonList(fdw)
}

// Update returns a builder for updating this FlowDropWithdraw.
// Note that you need to call FlowDropWithdraw.Unwrap() before calling this method if this FlowDropWithdraw
// was returned from a transaction, and the transaction was committed or rolled back.
func (fdw *FlowDropWithdraw) Update() *FlowDropWithdrawUpdateOne {
	return (&FlowDropWithdrawClient{config: fdw.config}).UpdateOne(fdw)
}

// Unwrap unwraps the FlowDropWithdraw entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fdw *FlowDropWithdraw) Unwrap() *FlowDropWithdraw {
	tx, ok := fdw.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowDropWithdraw is not a transactional entity")
	}
	fdw.config.driver = tx.drv
	return fdw
}

// String implements the fmt.Stringer.
func (fdw *FlowDropWithdraw) String() string {
	var builder strings.Builder
	builder.WriteString("FlowDropWithdraw(")
	builder.WriteString(fmt.Sprintf("id=%v", fdw.ID))
	builder.WriteString(", flow_drop_withdraw_type=")
	builder.WriteString(fdw.FlowDropWithdrawType)
	builder.WriteString(", flow_drop_withdraw_comment=")
	builder.WriteString(fdw.FlowDropWithdrawComment)
	builder.WriteByte(')')
	return builder.String()
}

// FlowDropWithdraws is a parsable slice of FlowDropWithdraw.
type FlowDropWithdraws []*FlowDropWithdraw

func (fdw FlowDropWithdraws) config(cfg config) {
	for _i := range fdw {
		fdw[_i].config = cfg
	}
}
