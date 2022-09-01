package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule -> FlowGroup
type FlowGroup struct {
	ent.Schema
}

func (FlowGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "flow_group"},
	}
}

// Fields of the Task.
func (FlowGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_group_id"),
		field.String("flow_group_title"),
		field.String("flow_group_description"),
	}
}

// Edges of the Task.
func (FlowGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ParticipantFlowModule.Type).
			Ref("flow_group_list").
			Unique(),
	}
}


