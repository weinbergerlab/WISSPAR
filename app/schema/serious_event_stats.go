package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule -> SeriousEvent -> SeriousEventStats
type SeriousEventStats struct {
	ent.Schema
}

func (SeriousEventStats) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "serious_event_stats"},
	}
}

// Fields of the Task.
func (SeriousEventStats) Fields() []ent.Field {
	return []ent.Field{
		field.String("serious_event_stats_group_id"),
		field.String("serious_event_stats_num_events"),
		field.String("serious_event_stats_num_affected"),
		field.String("serious_event_stats_num_at_risk"),
	}
}

// Edges of the Task.
func (SeriousEventStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", SeriousEvent.Type).
			Ref("serious_event_stats_list").
			Unique(),
	}
}
