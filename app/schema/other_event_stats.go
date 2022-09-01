package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule -> OtherEvent -> OtherEventStats
type OtherEventStats struct {
	ent.Schema
}

func (OtherEventStats) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "other_event_stats"},
	}
}

// Fields of the Task.
func (OtherEventStats) Fields() []ent.Field {
	return []ent.Field{
		field.String("other_event_stats_group_id"),
		field.String("other_event_stats_num_events"),
		field.String("other_event_stats_num_affected"),
		field.String("other_event_stats_num_at_risk"),
	}
}

// Edges of the Task.
func (OtherEventStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OtherEvent.Type).
			Ref("other_event_stats_list").
			Unique(),
	}
}
