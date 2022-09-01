package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure
type OutcomeMeasure struct {
	ent.Schema
}

func (OutcomeMeasure) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_measure"},
	}
}

// Fields of the Task.
func (OutcomeMeasure) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_measure_type"),
		field.String("outcome_measure_title"),
		field.String("outcome_measure_description"),
		field.String("outcome_measure_population_description"),
		field.String("outcome_measure_reporting_status"),
		field.String("outcome_measure_anticipated_posting_date"),
		field.String("outcome_measure_param_type"),
		field.String("outcome_measure_dispersion_type"),
		field.String("outcome_measure_unit_of_measure"),
		field.String("outcome_measure_calculate_pct"),
		field.String("outcome_measure_time_frame"),
		field.String("outcome_measure_type_units_analyzed"),
		field.String("outcome_measure_denom_units_selected"),
	}
}

// Edges of the Task.
func (OutcomeMeasure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasuresModule.Type).
			Ref("outcome_measure_list").
			Unique(),
		edge.To("outcome_group_list", OutcomeGroup.Type),
		edge.To("outcome_overview_list", OutcomeOverview.Type),
		edge.To("outcome_denom_list", OutcomeDenom.Type),
		edge.To("outcome_class_list", OutcomeClass.Type),
		edge.To("outcome_analysis_list", OutcomeAnalysis.Type),
	}
}
