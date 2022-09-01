package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineMeasure -> BaselineClass -> BaselineCategory -> BaselineMeasurement
type BaselineMeasurement struct {
	ent.Schema
}

func (BaselineMeasurement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_measurement"},
	}
}

// Fields of the Task.
func (BaselineMeasurement) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_measurement_group_id"),
		field.String("baseline_measurement_value"),
		field.String("baseline_measurement_spread"),
		field.String("baseline_measurement_lower_limit"),
		field.String("baseline_measurement_upper_limit"),
		field.String("baseline_measurement_comment"),
	}
}

// Edges of the Task.
func (BaselineMeasurement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineCategory.Type).
			Ref("baseline_measurement_list").
			Unique(),
	}
}
