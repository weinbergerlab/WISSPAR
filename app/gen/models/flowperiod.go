// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
)

// FlowPeriod is the model entity for the FlowPeriod schema.
type FlowPeriod struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowPeriodTitle holds the value of the "flow_period_title" field.
	FlowPeriodTitle string `json:"flow_period_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowPeriodQuery when eager-loading is set.
	Edges                                    FlowPeriodEdges `json:"edges"`
	participant_flow_module_flow_period_list *int
}

// FlowPeriodEdges holds the relations/edges for other nodes in the graph.
type FlowPeriodEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ParticipantFlowModule `json:"parent,omitempty"`
	// FlowMilestoneList holds the value of the flow_milestone_list edge.
	FlowMilestoneList []*FlowMilestone `json:"flow_milestone_list,omitempty"`
	// FlowDropWithdrawList holds the value of the flow_drop_withdraw_list edge.
	FlowDropWithdrawList []*FlowDropWithdraw `json:"flow_drop_withdraw_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowPeriodEdges) ParentOrErr() (*ParticipantFlowModule, error) {
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

// FlowMilestoneListOrErr returns the FlowMilestoneList value or an error if the edge
// was not loaded in eager-loading.
func (e FlowPeriodEdges) FlowMilestoneListOrErr() ([]*FlowMilestone, error) {
	if e.loadedTypes[1] {
		return e.FlowMilestoneList, nil
	}
	return nil, &NotLoadedError{edge: "flow_milestone_list"}
}

// FlowDropWithdrawListOrErr returns the FlowDropWithdrawList value or an error if the edge
// was not loaded in eager-loading.
func (e FlowPeriodEdges) FlowDropWithdrawListOrErr() ([]*FlowDropWithdraw, error) {
	if e.loadedTypes[2] {
		return e.FlowDropWithdrawList, nil
	}
	return nil, &NotLoadedError{edge: "flow_drop_withdraw_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowPeriod) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowperiod.FieldID:
			values[i] = new(sql.NullInt64)
		case flowperiod.FieldFlowPeriodTitle:
			values[i] = new(sql.NullString)
		case flowperiod.ForeignKeys[0]: // participant_flow_module_flow_period_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowPeriod", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowPeriod fields.
func (fp *FlowPeriod) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowperiod.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fp.ID = int(value.Int64)
		case flowperiod.FieldFlowPeriodTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_period_title", values[i])
			} else if value.Valid {
				fp.FlowPeriodTitle = value.String
			}
		case flowperiod.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field participant_flow_module_flow_period_list", value)
			} else if value.Valid {
				fp.participant_flow_module_flow_period_list = new(int)
				*fp.participant_flow_module_flow_period_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowPeriod entity.
func (fp *FlowPeriod) QueryParent() *ParticipantFlowModuleQuery {
	return (&FlowPeriodClient{config: fp.config}).QueryParent(fp)
}

// QueryFlowMilestoneList queries the "flow_milestone_list" edge of the FlowPeriod entity.
func (fp *FlowPeriod) QueryFlowMilestoneList() *FlowMilestoneQuery {
	return (&FlowPeriodClient{config: fp.config}).QueryFlowMilestoneList(fp)
}

// QueryFlowDropWithdrawList queries the "flow_drop_withdraw_list" edge of the FlowPeriod entity.
func (fp *FlowPeriod) QueryFlowDropWithdrawList() *FlowDropWithdrawQuery {
	return (&FlowPeriodClient{config: fp.config}).QueryFlowDropWithdrawList(fp)
}

// Update returns a builder for updating this FlowPeriod.
// Note that you need to call FlowPeriod.Unwrap() before calling this method if this FlowPeriod
// was returned from a transaction, and the transaction was committed or rolled back.
func (fp *FlowPeriod) Update() *FlowPeriodUpdateOne {
	return (&FlowPeriodClient{config: fp.config}).UpdateOne(fp)
}

// Unwrap unwraps the FlowPeriod entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fp *FlowPeriod) Unwrap() *FlowPeriod {
	tx, ok := fp.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowPeriod is not a transactional entity")
	}
	fp.config.driver = tx.drv
	return fp
}

// String implements the fmt.Stringer.
func (fp *FlowPeriod) String() string {
	var builder strings.Builder
	builder.WriteString("FlowPeriod(")
	builder.WriteString(fmt.Sprintf("id=%v", fp.ID))
	builder.WriteString(", flow_period_title=")
	builder.WriteString(fp.FlowPeriodTitle)
	builder.WriteByte(')')
	return builder.String()
}

// FlowPeriods is a parsable slice of FlowPeriod.
type FlowPeriods []*FlowPeriod

func (fp FlowPeriods) config(cfg config) {
	for _i := range fp {
		fp[_i].config = cfg
	}
}
