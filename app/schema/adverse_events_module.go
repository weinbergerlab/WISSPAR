package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdverseEventsModule is a root level object within the results definition object. Definition: https://prsinfo.clinicaltrials.gov/results_definitions.html#Result_AdverseEvents
type AdverseEventsModule struct {
	ent.Schema
}

func (AdverseEventsModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "adverse_events_module"},
	}
}

// Fields of the Task.
func (AdverseEventsModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("events_frequency_threshold"),
		field.String("events_time_frame"),
		field.String("events_description"),
	}
}

// Edges of the Task.
func (AdverseEventsModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ResultsDefinition.Type).Ref("adverse_events_module").Unique(),
		edge.To("event_group_list", EventGroup.Type),
		edge.To("serious_event_list", SeriousEvent.Type),
		edge.To("other_event_list", OtherEvent.Type),
	}
}
