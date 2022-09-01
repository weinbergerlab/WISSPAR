package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// ClinicalTrial -> ResultsDefinition houses the modules from the clinical trial
type ResultsDefinition struct {
	ent.Schema
}

func (ResultsDefinition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "results_definition"},
	}
}

// Fields of the Task.
func (ResultsDefinition) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Task.
func (ResultsDefinition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ClinicalTrial.Type).Ref("results_definition").Unique(),
		edge.To("participant_flow_module", ParticipantFlowModule.Type).Unique(),
		edge.To("baseline_characteristics_module", BaselineCharacteristicsModule.Type).Unique(),
		edge.To("outcome_measures_module", OutcomeMeasuresModule.Type).Unique(),
		edge.To("adverse_events_module", AdverseEventsModule.Type).Unique(),
		edge.To("more_info_module", MoreInfoModule.Type).Unique(),
	}
}
