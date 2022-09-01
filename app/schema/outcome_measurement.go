package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeClass -> OutcomeCategory -> OutcomeMeasurement
type OutcomeMeasurement struct {
	ent.Schema
}

func (OutcomeMeasurement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_measurement"},
	}
}

// Fields of the Task.
func (OutcomeMeasurement) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_measurement_group_id"),
		field.String("outcome_measurement_value"),
		field.String("outcome_measurement_spread"),
		field.String("outcome_measurement_lower_limit"),
		field.String("outcome_measurement_upper_limit"),
		field.String("outcome_measurement_comment"),
	}
}

// Edges of the Task.
func (OutcomeMeasurement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeCategory.Type).
			Ref("outcome_measurement_list").
			Unique(),
	}
}
