package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClinicalTrial -> StudyLocation houses the locations from the clinical trial
type StudyLocation struct {
	ent.Schema
}

func (StudyLocation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "study_location"},
	}
}

// Fields of the Task.
func (StudyLocation) Fields() []ent.Field {
	return []ent.Field{
		field.String("LocationFacility"),
		field.String("LocationCity"),
		field.String("LocationCountry"),
		field.String("LocationCountryCode"),
		field.String("LocationContinentName"),
		field.String("LocationContinentCode"),
	}
}

// Edges of the Task.
func (StudyLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ClinicalTrial.Type).Ref("study_locations").Unique(),
	}
}
