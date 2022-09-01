package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule is a root level object within the results definition object. Definition: https://prsinfo.clinicaltrials.gov/results_definitions.html#Result_Baseline
type BaselineCharacteristicsModule struct {
	ent.Schema
}

func (BaselineCharacteristicsModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_characteristics_module"},
	}
}

// Fields of the Task.
func (BaselineCharacteristicsModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_population_description"),
		field.String("baseline_type_units_analyzed"),
	}
}

// Edges of the Task.
func (BaselineCharacteristicsModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ResultsDefinition.Type).Ref("baseline_characteristics_module").Unique(),
		edge.To("baseline_group_list", BaselineGroup.Type),
		edge.To("baseline_denom_list", BaselineDenom.Type),
		edge.To("baseline_measure_list", BaselineMeasure.Type),
	}
}
