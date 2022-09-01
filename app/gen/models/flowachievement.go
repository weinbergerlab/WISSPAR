// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
)

// FlowAchievement is the model entity for the FlowAchievement schema.
type FlowAchievement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowAchievementGroupID holds the value of the "flow_achievement_group_id" field.
	FlowAchievementGroupID string `json:"flow_achievement_group_id,omitempty"`
	// FlowAchievementComment holds the value of the "flow_achievement_comment" field.
	FlowAchievementComment string `json:"flow_achievement_comment,omitempty"`
	// FlowAchievementNumSubjects holds the value of the "flow_achievement_num_subjects" field.
	FlowAchievementNumSubjects string `json:"flow_achievement_num_subjects,omitempty"`
	// FlowAchievementNumUnits holds the value of the "flow_achievement_num_units" field.
	FlowAchievementNumUnits string `json:"flow_achievement_num_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowAchievementQuery when eager-loading is set.
	Edges                                FlowAchievementEdges `json:"edges"`
	flow_milestone_flow_achievement_list *int
}

// FlowAchievementEdges holds the relations/edges for other nodes in the graph.
type FlowAchievementEdges struct {
	// Parent holds the value of the parent edge.
	Parent *FlowMilestone `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowAchievementEdges) ParentOrErr() (*FlowMilestone, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flowmilestone.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowAchievement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowachievement.FieldID:
			values[i] = new(sql.NullInt64)
		case flowachievement.FieldFlowAchievementGroupID, flowachievement.FieldFlowAchievementComment, flowachievement.FieldFlowAchievementNumSubjects, flowachievement.FieldFlowAchievementNumUnits:
			values[i] = new(sql.NullString)
		case flowachievement.ForeignKeys[0]: // flow_milestone_flow_achievement_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowAchievement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowAchievement fields.
func (fa *FlowAchievement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowachievement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fa.ID = int(value.Int64)
		case flowachievement.FieldFlowAchievementGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_achievement_group_id", values[i])
			} else if value.Valid {
				fa.FlowAchievementGroupID = value.String
			}
		case flowachievement.FieldFlowAchievementComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_achievement_comment", values[i])
			} else if value.Valid {
				fa.FlowAchievementComment = value.String
			}
		case flowachievement.FieldFlowAchievementNumSubjects:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_achievement_num_subjects", values[i])
			} else if value.Valid {
				fa.FlowAchievementNumSubjects = value.String
			}
		case flowachievement.FieldFlowAchievementNumUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_achievement_num_units", values[i])
			} else if value.Valid {
				fa.FlowAchievementNumUnits = value.String
			}
		case flowachievement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field flow_milestone_flow_achievement_list", value)
			} else if value.Valid {
				fa.flow_milestone_flow_achievement_list = new(int)
				*fa.flow_milestone_flow_achievement_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the FlowAchievement entity.
func (fa *FlowAchievement) QueryParent() *FlowMilestoneQuery {
	return (&FlowAchievementClient{config: fa.config}).QueryParent(fa)
}

// Update returns a builder for updating this FlowAchievement.
// Note that you need to call FlowAchievement.Unwrap() before calling this method if this FlowAchievement
// was returned from a transaction, and the transaction was committed or rolled back.
func (fa *FlowAchievement) Update() *FlowAchievementUpdateOne {
	return (&FlowAchievementClient{config: fa.config}).UpdateOne(fa)
}

// Unwrap unwraps the FlowAchievement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fa *FlowAchievement) Unwrap() *FlowAchievement {
	tx, ok := fa.config.driver.(*txDriver)
	if !ok {
		panic("models: FlowAchievement is not a transactional entity")
	}
	fa.config.driver = tx.drv
	return fa
}

// String implements the fmt.Stringer.
func (fa *FlowAchievement) String() string {
	var builder strings.Builder
	builder.WriteString("FlowAchievement(")
	builder.WriteString(fmt.Sprintf("id=%v", fa.ID))
	builder.WriteString(", flow_achievement_group_id=")
	builder.WriteString(fa.FlowAchievementGroupID)
	builder.WriteString(", flow_achievement_comment=")
	builder.WriteString(fa.FlowAchievementComment)
	builder.WriteString(", flow_achievement_num_subjects=")
	builder.WriteString(fa.FlowAchievementNumSubjects)
	builder.WriteString(", flow_achievement_num_units=")
	builder.WriteString(fa.FlowAchievementNumUnits)
	builder.WriteByte(')')
	return builder.String()
}

// FlowAchievements is a parsable slice of FlowAchievement.
type FlowAchievements []*FlowAchievement

func (fa FlowAchievements) config(cfg config) {
	for _i := range fa {
		fa[_i].config = cfg
	}
}
