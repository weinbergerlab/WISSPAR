package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial is a root level object
type ClinicalTrial struct {
	ent.Schema
}

func (ClinicalTrial) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "clinical_trial"},
	}
}

// Fields of the Task.
func (ClinicalTrial) Fields() []ent.Field {
	return []ent.Field{
		field.String("study_name"),
		field.String("sponsor"),
		field.String("responsible_party"),
		field.String("study_type"),
		field.String("phase"),
		field.String("actual_enrollment"),
		field.String("allocation"),
		field.String("intervention_model"),
		field.String("masking"),
		field.String("primary_purpose"),
		field.String("official_title"),
		field.Time("actual_study_start_date"),
		field.Time("actual_primary_completion_date"),
		field.Time("actual_study_completion_date"),
		field.String("study_id"),
	}
}

// Edges of the Task.
func (ClinicalTrial) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("results_definition", ResultsDefinition.Type).Unique(),
		edge.To("study_locations", StudyLocation.Type),
		edge.To("study_eligibility", StudyEligibility.Type),
		edge.To("metadata", ClinicalTrialMetadata.Type),
	}
}
