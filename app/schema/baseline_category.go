package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineClass -> BaselineCategory
type BaselineCategory struct {
	ent.Schema
}

func (BaselineCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_category"},
	}
}

// Fields of the Task.
func (BaselineCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_category_title"),
	}
}

// Edges of the Task.
func (BaselineCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineClass.Type).
			Ref("baseline_category_list").
			Unique(),
		edge.To("baseline_measurement_list", BaselineMeasurement.Type),
	}
}
