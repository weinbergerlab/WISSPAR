package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial is a root level object
type Vaccine struct {
	ent.Schema
}

func (Vaccine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "vaccine"},
	}
}

// Fields of the Task.
func (Vaccine) Fields() []ent.Field {
	return []ent.Field{
		field.String("vaccine_name").NotEmpty().Unique(),
	}
}

// Edges of the Task.
func (Vaccine) Edges() []ent.Edge {
	return nil
}

