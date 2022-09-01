package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeGroup
type OutcomeGroup struct {
	ent.Schema
}

func (OutcomeGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_group"},
	}
}

// Fields of the Task.
func (OutcomeGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_group_id"),
		field.String("outcome_group_title"),
		field.String("outcome_group_description"),
	}
}

// Edges of the Task.
func (OutcomeGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasure.Type).
			Ref("outcome_group_list").
			Unique(),
	}
}
