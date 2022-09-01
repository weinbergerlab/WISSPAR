package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial is a root level object
type Manufacturer struct {
	ent.Schema
}

func (Manufacturer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "manufacturer"},
	}
}

// Fields of the Task.
func (Manufacturer) Fields() []ent.Field {
	return []ent.Field{
		field.String("manufacturer_name").NotEmpty().Unique(),
	}
}

// Edges of the Task.
func (Manufacturer) Edges() []ent.Edge {
	return nil
}

