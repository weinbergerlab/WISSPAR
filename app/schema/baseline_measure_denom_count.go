package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineMeasureDenom -> BaselineMeasureDenomCount
type BaselineMeasureDenomCount struct {
	ent.Schema
}

func (BaselineMeasureDenomCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_measure_denom_count"},
	}
}

// Fields of the Task.
func (BaselineMeasureDenomCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_measure_denom_count_group_id"),
		field.String("baseline_measure_denom_count_value"),
	}
}

// Edges of the Task.
func (BaselineMeasureDenomCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineMeasureDenom.Type).
			Ref("baseline_measure_denom_count_list").
			Unique(),
	}
}
