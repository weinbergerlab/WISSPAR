package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MoreInfoModule is a root level object within the results definition object.
type MoreInfoModule struct {
	ent.Schema
}

func (MoreInfoModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "more_info_module"},
	}
}

// Fields of the Task.
func (MoreInfoModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("limitations_and_caveats_description"),
	}
}

// Edges of the Task.
func (MoreInfoModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ResultsDefinition.Type).Ref("more_info_module").Unique(),
		edge.To("certain_agreement", CertainAgreement.Type).Unique(),
		edge.To("point_of_contact", PointOfContact.Type).Unique(),
	}
}
