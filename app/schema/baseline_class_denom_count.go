package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineClass -> BaselineClassDenom -> BaselineClassDenomCount
type BaselineClassDenomCount struct {
	ent.Schema
}

func (BaselineClassDenomCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_class_denom_count"},
	}
}

// Fields of the Task.
func (BaselineClassDenomCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_class_denom_count_group_id"),
		field.String("baseline_class_denom_count_value"),
	}
}

// Edges of the Task.
func (BaselineClassDenomCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineClassDenom.Type).
			Ref("baseline_class_denom_count_list").
			Unique(),
	}
}
