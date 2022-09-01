package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowPeriod -> FlowMilestone -> FlowAchievement
type FlowAchievement struct {
	ent.Schema
}

func (FlowAchievement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_achievement"},
	}
}

// Fields of the Task.
func (FlowAchievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_achievement_group_id"),
		field.String("flow_achievement_comment"),
		field.String("flow_achievement_num_subjects"),
		field.String("flow_achievement_num_units"),
	}
}

// Edges of the Task.
func (FlowAchievement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", FlowMilestone.Type).
			Ref("flow_achievement_list").
			Unique(),
	}
}
