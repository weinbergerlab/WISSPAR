package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineClass -> BaselineClassDenom
type BaselineClassDenom struct {
	ent.Schema
}

func (BaselineClassDenom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_class_denom"},
	}
}

// Fields of the Task.
func (BaselineClassDenom) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_class_denom_units"),
	}
}

// Edges of the Task.
func (BaselineClassDenom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineClass.Type).
			Ref("baseline_class_denom_list").
			Unique(),
		edge.To("baseline_class_denom_count_list", BaselineClassDenomCount.Type),
	}
}
