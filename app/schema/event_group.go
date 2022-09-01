package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule -> EventGroup
type EventGroup struct {
	ent.Schema
}

func (EventGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "event_group"},
	}
}

// Fields of the Task.
func (EventGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_group_id"),
		field.String("event_group_title"),
		field.String("event_group_description"),
		field.String("event_group_deaths_num_affected"),
		field.String("event_group_deaths_num_at_risk"),
		field.String("event_group_serious_num_affected"),
		field.String("event_group_serious_num_at_risk"),
		field.String("event_group_other_num_affected"),
		field.String("event_group_other_num_at_risk"),
	}
}

// Edges of the Task.
func (EventGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", AdverseEventsModule.Type).
			Ref("event_group_list").
			Unique(),
	}
}
