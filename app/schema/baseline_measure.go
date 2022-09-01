package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure
type BaselineMeasure struct {
	ent.Schema
}

func (BaselineMeasure) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_measure"},
	}
}

// Fields of the Task.
func (BaselineMeasure) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_measure_title"),
		field.String("baseline_measure_description"),
		field.String("baseline_measure_population_description"),
		field.String("baseline_measure_param_type"),
		field.String("baseline_measure_dispersion_type"),
		field.String("baseline_measure_unit_of_measure"),
		field.String("baseline_measure_calculate_pct"),
		field.String("baseline_measure_denom_units_selected"),
	}
}

// Edges of the Task.
func (BaselineMeasure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineCharacteristicsModule.Type).
			Ref("baseline_measure_list").
			Unique(),
		edge.To("baseline_measure_denom_list", BaselineMeasureDenom.Type),
		edge.To("baseline_class_list", BaselineClass.Type),
	}
}
