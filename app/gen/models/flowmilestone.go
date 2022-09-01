// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
)

// FlowMilestone is the model entity for the FlowMilestone schema.
type FlowMilestone struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowMilestoneType holds the value of the "flow_milestone_type" field.
	FlowMilestoneType string `json:"flow_milestone_type,omitempty"`
	// FlowMilestoneComment holds the value of the "flow_milestone_comment" field.
	FlowMilestoneComment string `json:"flow_milestone_comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowMilestoneQuery when eager-loading is set.
	Edges                           FlowMilestoneEdges `json:"edges"`
	flow_period_flow_milestone_list *int
}

// FlowMilestoneEdges holds the relations/edges for other nodes in the graph.
type FlowMilestoneEdges struct {
	// Parent holds the value of the parent edge.
	Parent *FlowPeriod `json:"parent,omitempty"`
	// FlowAchievementList holds the value of the flow_achievement_list edge.
	FlowAchievementList []*FlowAchievement `json:"flow_achievement_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowMilestoneEdges) ParentOrErr() (*FlowPeriod, error) {
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

// FlowAchievementListOrErr returns the FlowAchievementList value or an error if the edge
// was not loaded in eager-loading.
func (e FlowMilestoneEdges) FlowAchievementListOrErr() ([]*FlowAchievement, error) {
	if e.loadedTypes[1] {
		return e.FlowAchievementList, nil
	}
	return nil, &NotLoadedError{edge: "flow_achievement_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowMilestone) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowmilestone.FieldID:
			values[i] = new(sql.NullInt64)
		case flowmilestone.FieldFlowMilestoneType, flowmilestone.FieldFlowMilestoneComment:
			values[i] = new(sql.NullString)
		case flowmilestone.ForeignKeys[0]: // flow_period_flow_milestone_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowMilestone", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowMilestone fields.
func (fm *FlowMilestone) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowmilestone.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fm.ID = int(value.Int64)
		case flowmilestone.FieldFlowMilestoneType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_milestone_type", values[i])
			} else if value.Valid {
				fm.FlowMilestoneType = value.String
			}
		case flowmilestone.FieldFlowMilestoneComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_milestone_comment", values[i])
			} else if value.Valid {
				fm.FlowMilestoneComment = value.String
			}
		case flowmilestone.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field flow_period_flow_milestone_list", value)
			} else if value.Valid {
				fm.flow_period_flow_milestone_list = new(int)
				*fm.flow_period_flow_milestone_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowMilestone entity.
func (fm *FlowMilestone) QueryParent() *FlowPeriodQuery {
	return (&FlowMilestoneClient{config: fm.config}).QueryParent(fm)
}

// QueryFlowAchievementList queries the "flow_achievement_list" edge of the FlowMilestone entity.
func (fm *FlowMilestone) QueryFlowAchievementList() *FlowAchievementQuery {
	return (&FlowMilestoneClient{config: fm.config}).QueryFlowAchievementList(fm)
}

// Update returns a builder for updating this FlowMilestone.
// Note that you need to call FlowMilestone.Unwrap() before calling this method if this FlowMilestone
// was returned from a transaction, and the transaction was committed or rolled back.
func (fm *FlowMilestone) Update() *FlowMilestoneUpdateOne {
	return (&FlowMilestoneClient{config: fm.config}).UpdateOne(fm)
}

// Unwrap unwraps the FlowMilestone entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fm *FlowMilestone) Unwrap() *FlowMilestone {
	tx, ok := fm.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowMilestone is not a transactional entity")
	}
	fm.config.driver = tx.drv
	return fm
}

// String implements the fmt.Stringer.
func (fm *FlowMilestone) String() string {
	var builder strings.Builder
	builder.WriteString("FlowMilestone(")
	builder.WriteString(fmt.Sprintf("id=%v", fm.ID))
	builder.WriteString(", flow_milestone_type=")
	builder.WriteString(fm.FlowMilestoneType)
	builder.WriteString(", flow_milestone_comment=")
	builder.WriteString(fm.FlowMilestoneComment)
	builder.WriteByte(')')
	return builder.String()
}

// FlowMilestones is a parsable slice of FlowMilestone.
type FlowMilestones []*FlowMilestone

func (fm FlowMilestones) config(cfg config) {
	for _i := range fm {
		fm[_i].config = cfg
	}
}
