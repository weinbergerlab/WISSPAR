package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule -> OtherEvent
type OtherEvent struct {
	ent.Schema
}

func (OtherEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "other_event"},
	}
}

// Fields of the Task.
func (OtherEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("other_event_term"),
		field.String("other_event_organ_system"),
		field.String("other_event_source_vocabulary"),
		field.String("other_event_assessment_type"),
		field.String("other_event_notes"),
	}
}

// Edges of the Task.
func (OtherEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", AdverseEventsModule.Type).
			Ref("other_event_list").
			Unique(),
		edge.To("other_event_stats_list", OtherEventStats.Type),
	}
}
