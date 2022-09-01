package schema

import (
	"entgo.io/ent"
"entgo.io/ent/dialect/entsql"
"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowPeriod -> FlowDropWithdraw
type FlowDropWithdraw struct {
	ent.Schema
}

func (FlowDropWithdraw) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_drop_withdraw"},
	}
}

// Fields of the Task.
func (FlowDropWithdraw) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_drop_withdraw_type"),
		field.String("flow_drop_withdraw_comment"),
	}
}

// Edges of the Task.
func (FlowDropWithdraw) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", FlowPeriod.Type).
			Ref("flow_drop_withdraw_list").
			Unique(),
			edge.To("flow_reason_list", FlowReason.Type),
	}
}


