package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MoreInfoModule -> PointOfContact
type PointOfContact struct {
	ent.Schema
}

func (PointOfContact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_of_contact"},
	}
}

// Fields of the Task.
func (PointOfContact) Fields() []ent.Field {
	return []ent.Field{
		field.String("point_of_contact_title"),
		field.String("point_of_contact_organization"),
		field.String("point_of_contact_email"),
		field.String("point_of_contact_phone"),
		field.String("point_of_contact_phone_ext"),
	}
}

// Edges of the Task.
func (PointOfContact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", MoreInfoModule.Type).Unique().Required(),
	}
}
