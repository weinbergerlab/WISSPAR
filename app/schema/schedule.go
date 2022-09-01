package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial is a root level object
type Schedule struct {
	ent.Schema
}

func (Schedule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "schedule"},
	}
}

// Fields of the Task.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Task.
func (Schedule) Edges() []ent.Edge {
	return nil
}

