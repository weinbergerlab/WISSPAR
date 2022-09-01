package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeClass -> OutcomeClassDenom
type OutcomeClassDenom struct {
	ent.Schema
}

func (OutcomeClassDenom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_class_denom"},
	}
}

// Fields of the Task.
func (OutcomeClassDenom) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_class_denom_units"),
	}
}

// Edges of the Task.
func (OutcomeClassDenom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeClass.Type).
			Ref("outcome_class_denom_list").
			Unique(),
		edge.To("outcome_class_denom_count_list", OutcomeClassDenomCount.Type),
	}
}
