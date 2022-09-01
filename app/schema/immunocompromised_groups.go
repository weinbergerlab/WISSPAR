package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ImmunocompromisedGroups is a root level object
type ImmunocompromisedGroups struct {
	ent.Schema
}

func (ImmunocompromisedGroups) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "immunocompromised_groups"},
	}
}

// Fields of the Task.
func (ImmunocompromisedGroups) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_name").NotEmpty().Unique(),
	}
}

// Edges of the Task.
func (ImmunocompromisedGroups) Edges() []ent.Edge {
	return nil
}
