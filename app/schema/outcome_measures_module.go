package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// OutcomeMeasuresModule is a root level object within the results definition object. Definition: https://prsinfo.clinicaltrials.gov/results_definitions.html#Result_Outcome_Measure
type OutcomeMeasuresModule struct {
	ent.Schema
}

func (OutcomeMeasuresModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_measures_module"},
	}
}

// Fields of the Task.
func (OutcomeMeasuresModule) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Task.
func (OutcomeMeasuresModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ResultsDefinition.Type).Ref("outcome_measures_module").Unique(),
		edge.To("outcome_measure_list", OutcomeMeasure.Type),
	}
}
