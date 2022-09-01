package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BaselineCharacteristicsModule -> BaselineGroup
type BaselineGroup struct {
	ent.Schema
}

func (BaselineGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseline_group"},
	}
}

// Fields of the Task.
func (BaselineGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("baseline_group_id"),
		field.String("baseline_group_title"),
		field.String("baseline_group_description"),
	}
}

// Edges of the Task.
func (BaselineGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BaselineCharacteristicsModule.Type).
			Ref("baseline_group_list").
			Unique(),
	}
}
