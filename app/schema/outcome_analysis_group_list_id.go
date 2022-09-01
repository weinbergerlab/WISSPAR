package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeAnalysis -> OutcomeAnalysisGroupID
type OutcomeAnalysisGroupID struct {
	ent.Schema
}

func (OutcomeAnalysisGroupID) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_analysis_group_id"},
	}
}

// Fields of the Task.
func (OutcomeAnalysisGroupID) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_analysis_group_id"),
	}
}

// Edges of the Task.
func (OutcomeAnalysisGroupID) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeAnalysis.Type).
			Ref("outcome_analysis_group_id_list").
			Unique(),
	}
}
