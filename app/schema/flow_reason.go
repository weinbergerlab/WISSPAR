package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowPeriod -> FlowDropWithdraw -> FlowReason
type FlowReason struct {
	ent.Schema
}

func (FlowReason) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_reason"},
	}
}

// Fields of the Task.
func (FlowReason) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_reason_group_id"),
		field.String("flow_reason_comment"),
		field.String("flow_reason_num_subjects"),
		field.String("flow_reason_num_units"),
	}
}

// Edges of the Task.
func (FlowReason) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", FlowDropWithdraw.Type).
			Ref("flow_reason_list").
			Unique(),
	}
}
