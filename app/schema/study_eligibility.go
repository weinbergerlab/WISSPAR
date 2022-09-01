package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial -> StudyEligibility houses the locations from the clinical trial
type StudyEligibility struct {
	ent.Schema
}

func (StudyEligibility) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "study_eligibility"},
	}
}

// Fields of the Task.
func (StudyEligibility) Fields() []ent.Field {
	return []ent.Field{
		field.String("EligibilityCriteria"),
		field.String("HealthyVolunteers"),
		field.String("Gender"),
		field.String("MinimumAge"),
		field.String("MaximumAge"),
		field.String("StdAgeList"),
		field.String("Ethnicity"),
	}
}

// Edges of the Task.
func (StudyEligibility) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ClinicalTrial.Type).Ref("study_eligibility").Unique(),
	}
}
