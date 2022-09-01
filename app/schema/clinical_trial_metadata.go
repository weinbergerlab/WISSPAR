package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClinicalTrialMetadata is a root level object
type ClinicalTrialMetadata struct {
	ent.Schema
}

func (ClinicalTrialMetadata) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "clinical_trial_metadata"},
	}
}

// Fields of the Task.
func (ClinicalTrialMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag_name"),
		field.String("tag_value"),
		field.String("use_case_code").Default("default"),
	}
}

// Edges of the Task.
func (ClinicalTrialMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ClinicalTrial.Type).Ref("metadata").Unique(),
	}
}
