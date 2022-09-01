package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OutcomeMeasuresModule -> OutcomeMeasure -> OutcomeGroup
type OutcomeOverview struct {
	ent.Schema
}

func (OutcomeOverview) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "outcome_overview"},
	}
}

// Fields of the Task.
func (OutcomeOverview) Fields() []ent.Field {
	return []ent.Field{
		field.String("outcome_overview_id"),
		field.String("outcome_overview_title"),
		field.String("outcome_overview_description"),
		field.String("outcome_overview_participants"),
		field.String("outcome_overview_time_frame"),
		field.String("outcome_overview_serotype"),
		field.String("outcome_overview_assay"),
		field.Int64("outcome_overview_dose_number").Default(0),
		field.Float("outcome_overview_value"),
		field.String("outcome_overview_upper_limit"),
		field.String("outcome_overview_lower_limit"),
		field.String("outcome_overview_group_id").Default(""),
		field.String("outcome_overview_ratio").Default(""),
		field.String("outcome_overview_measure_title").Default(""),
		field.String("outcome_overview_vaccine").Default(""),
		field.String("outcome_overview_immunocompromised_population").Default("N/A"),
		field.String("outcome_overview_manufacturer").Default(""),
		field.String("outcome_overview_confidence_interval").Default(""),
		field.Float("outcome_overview_percent_responders").Default(0),
		field.Int64("outcome_overview_time_frame_weeks").Default(0),
		field.String("outcome_overview_dose_description").Default(""),
		field.String("outcome_overview_schedule").Default(""),
		field.String("outcome_overview_use_case_code").Default("default"),
	}
}

// Edges of the Task.
func (OutcomeOverview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", OutcomeMeasure.Type).Ref("outcome_overview_list").Unique(),
	}
}
