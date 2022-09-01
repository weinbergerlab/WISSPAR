package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowPeriod -> FlowMilestone
type FlowMilestone struct {
	ent.Schema
}

func (FlowMilestone) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_milestone"},
	}
}

// Fields of the Task.
func (FlowMilestone) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_milestone_type"),
		field.String("flow_milestone_comment"),
	}
}

// Edges of the Task.
func (FlowMilestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", FlowPeriod.Type).
			Ref("flow_milestone_list").
			Unique(),
		edge.To("flow_achievement_list", FlowAchievement.Type),
	}
}
