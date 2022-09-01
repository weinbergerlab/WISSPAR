package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MoreInfoModule -> CertainAgreement
type CertainAgreement struct {
	ent.Schema
}

func (CertainAgreement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "certain_agreement"},
	}
}

// Fields of the Task.
func (CertainAgreement) Fields() []ent.Field {
	return []ent.Field{
		field.String("agreement_pi_sponsor_employee"),
		field.String("agreement_restriction_type"),
		field.String("agreement_restrictive_agreement"),
		field.String("agreement_other_details"),
	}
}

// Edges of the Task.
func (CertainAgreement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", MoreInfoModule.Type).Unique().Required(),
	}
}
