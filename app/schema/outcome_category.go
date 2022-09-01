package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeClass -> OutcomeCategory
type OutcomeCategory struct {
	ent.Schema
}

func (OutcomeCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_category"},
	}
}

// Fields of the Task.
func (OutcomeCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_category_title"),
	}
}

// Edges of the Task.
func (OutcomeCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeClass.Type).
			Ref("outcome_category_list").
			Unique(),
		edge.To("outcome_measurement_list", OutcomeMeasurement.Type),
	}
}
