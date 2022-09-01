package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeClass -> OutcomeClassDenom -> OutcomeClassDenomCount
type OutcomeClassDenomCount struct {
	ent.Schema
}

func (OutcomeClassDenomCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_class_denom_count"},
	}
}

// Fields of the Task.
func (OutcomeClassDenomCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_class_denom_count_group_id"),
		field.String("outcome_class_denom_count_value"),
	}
}

// Edges of the Task.
func (OutcomeClassDenomCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeClassDenom.Type).
			Ref("outcome_class_denom_count_list").
			Unique(),
	}
}
