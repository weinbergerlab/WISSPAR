package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial is a root level object
type DoseDescription struct {
	ent.Schema
}

func (DoseDescription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "dose_description"},
	}
}

// Fields of the Task.
func (DoseDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Task.
func (DoseDescription) Edges() []ent.Edge {
	return nil
}

