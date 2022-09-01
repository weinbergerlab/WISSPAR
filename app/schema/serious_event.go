package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule -> SeriousEvent
type SeriousEvent struct {
	ent.Schema
}

func (SeriousEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "serious_event"},
	}
}

// Fields of the Task.
func (SeriousEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("serious_event_term"),
		field.String("serious_event_organ_system"),
		field.String("serious_event_source_vocabulary"),
		field.String("serious_event_assessment_type"),
		field.String("serious_event_notes"),
	}
}

// Edges of the Task.
func (SeriousEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", AdverseEventsModule.Type).
			Ref("serious_event_list").
			Unique(),
		edge.To("serious_event_stats_list", SeriousEventStats.Type),
	}
}
