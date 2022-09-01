package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineClass
type BaselineClass struct {
	ent.Schema
}

func (BaselineClass) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_class"},
	}
}

// Fields of the Task.
func (BaselineClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_class_title"),
	}
}

// Edges of the Task.
func (BaselineClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineMeasure.Type).
			Ref("baseline_class_list").
			Unique(),
		edge.To("baseline_class_denom_list", BaselineClassDenom.Type),
		edge.To("baseline_category_list", BaselineCategory.Type),
	}
}
