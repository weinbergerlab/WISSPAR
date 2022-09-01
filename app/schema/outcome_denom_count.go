package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeDenom -> OutcomeDenomCount
type OutcomeDenomCount struct {
	ent.Schema
}

func (OutcomeDenomCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_denom_count"},
	}
}

// Fields of the Task.
func (OutcomeDenomCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_denom_count_group_id"),
		field.String("outcome_denom_count_value"),
	}
}

// Edges of the Task.
func (OutcomeDenomCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeDenom.Type).
			Ref("outcome_denom_count_list").
			Unique(),
	}
}
