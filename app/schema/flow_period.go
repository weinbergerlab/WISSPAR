package schema

import (
	"entgo.io/ent"
"entgo.io/ent/dialect/entsql"
"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowPeriod
type FlowPeriod struct {
	ent.Schema
}

func (FlowPeriod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_period"},
	}
}

// Fields of the Task.
func (FlowPeriod) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_period_title"),
	}
}

// Edges of the Task.
func (FlowPeriod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ParticipantFlowModule.Type).
			Ref("flow_period_list").
			Unique(),
		edge.To("flow_milestone_list", FlowMilestone.Type),
		edge.To("flow_drop_withdraw_list", FlowDropWithdraw.Type),
	}
}


