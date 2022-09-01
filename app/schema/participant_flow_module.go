package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParticipantFlowModule is a root level object within the results definition object. Definition: https://prsinfo.clinicaltrials.gov/results_definitions.html#Result_ParticipantFlow
type ParticipantFlowModule struct {
	ent.Schema
}

func (ParticipantFlowModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "participant_flow_module"},
	}
}

// Fields of the Task.
func (ParticipantFlowModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_pre_assignment_details"),
		field.String("flow_recruitment_details"),
		field.String("flow_type_units_analyzed"),
	}
}

// Edges of the Task.
func (ParticipantFlowModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", ResultsDefinition.Type).Ref("participant_flow_module").Unique(),
		edge.To("flow_group_list", FlowGroup.Type),
		edge.To("flow_period_list", FlowPeriod.Type),
	}
}
