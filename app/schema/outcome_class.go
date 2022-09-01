package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeClass
type OutcomeClass struct {
	ent.Schema
}

func (OutcomeClass) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_class"},
	}
}

// Fields of the Task.
func (OutcomeClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_class_title"),
	}
}

// Edges of the Task.
func (OutcomeClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasure.Type).
			Ref("outcome_class_list").
			Unique(),
		edge.To("outcome_class_denom_list", OutcomeClassDenom.Type),
		edge.To("outcome_category_list", OutcomeCategory.Type),
	}
}
