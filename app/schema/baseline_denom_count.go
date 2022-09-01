package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineDenom -> BaselineDenomCount
type BaselineDenomCount struct {
	ent.Schema
}

func (BaselineDenomCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_denom_count"},
	}
}

// Fields of the Task.
func (BaselineDenomCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_denom_count_group_id"),
		field.String("baseline_denom_count_value"),
	}
}

// Edges of the Task.
func (BaselineDenomCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineDenom.Type).
			Ref("baseline_denom_count_list").
			Unique(),
	}
}
