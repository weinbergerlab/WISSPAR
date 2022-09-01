package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineMeasureDenom
type BaselineMeasureDenom struct {
	ent.Schema
}

func (BaselineMeasureDenom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_measure_denom"},
	}
}

// Fields of the Task.
func (BaselineMeasureDenom) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_measure_denom_units"),
	}
}

// Edges of the Task.
func (BaselineMeasureDenom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineMeasure.Type).
			Ref("baseline_measure_denom_list").
			Unique(),
		edge.To("baseline_measure_denom_count_list", BaselineMeasureDenomCount.Type),
	}
}
