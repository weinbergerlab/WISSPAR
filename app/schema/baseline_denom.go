package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineDenom
type BaselineDenom struct {
	ent.Schema
}

func (BaselineDenom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_denom"},
	}
}

// Fields of the Task.
func (BaselineDenom) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_denom_units"),
	}
}

// Edges of the Task.
func (BaselineDenom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineCharacteristicsModule.Type).
			Ref("baseline_denom_list").
			Unique(),
		edge.To("baseline_denom_count_list", BaselineDenomCount.Type),
	}
}
