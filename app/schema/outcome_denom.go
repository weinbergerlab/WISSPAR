package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeDenom
type OutcomeDenom struct {
	ent.Schema
}

func (OutcomeDenom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_denom"},
	}
}

// Fields of the Task.
func (OutcomeDenom) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_denom_units"),
	}
}

// Edges of the Task.
func (OutcomeDenom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasure.Type).
			Ref("outcome_denom_list").
			Unique(),
		edge.To("outcome_denom_count_list", OutcomeDenomCount.Type),
	}
}
