package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UseCase
type UseCase struct {
	ent.Schema
}

func (UseCase) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "use_case"},
	}
}

// Fields of the Task.
func (UseCase) Fields() []ent.Field {
	return []ent.Field{
		field.String("use_case_name"),
		field.String("use_case_description"),
		field.String("use_case_code"),
	}
}

// Edges of the Task.
func (UseCase) Edges() []ent.Edge {
	return nil
}
