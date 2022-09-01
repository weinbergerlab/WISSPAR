package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeAnalysis
type OutcomeAnalysis struct {
	ent.Schema
}

func (OutcomeAnalysis) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_analysis"},
	}
}

// Fields of the Task.
func (OutcomeAnalysis) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_analysis_group_description"),
		field.String("outcome_analysis_tested_non_inferiority"),
		field.String("outcome_analysis_non_inferiority_type"),
		field.String("outcome_analysis_non_inferiority_comment"),
		field.String("outcome_analysis_p_value"),
		field.String("outcome_analysis_p_value_comment"),
		field.String("outcome_analysis_statistical_method"),
		field.String("outcome_analysis_statistical_comment"),
		field.String("outcome_analysis_param_type"),
		field.String("outcome_analysis_param_value"),
		field.String("outcome_analysis_ci_pct_value"),
		field.String("outcome_analysis_ci_num_sides"),
		field.String("outcome_analysis_ci_lower_limit"),
		field.String("outcome_analysis_ci_upper_limit"),
		field.String("outcome_analysis_ci_lower_limit_comment"),
		field.String("outcome_analysis_ci_upper_limit_comment"),
		field.String("outcome_analysis_dispersion_type"),
		field.String("outcome_analysis_dispersion_value"),
		field.String("outcome_analysis_estimate_comment"),
		field.String("outcome_analysis_other_analysis_description"),
	}
}

// Edges of the Task.
func (OutcomeAnalysis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasure.Type).
			Ref("outcome_analysis_list").
			Unique(),
		edge.To("outcome_analysis_group_id_list", OutcomeAnalysisGroupID.Type),
	}
}
