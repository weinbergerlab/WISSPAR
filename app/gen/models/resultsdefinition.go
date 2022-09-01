// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ResultsDefinition is the model entity for the ResultsDefinition schema.
type ResultsDefinition struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResultsDefinitionQuery when eager-loading is set.
	Edges                             ResultsDefinitionEdges `json:"edges"`
	clinical_trial_results_definition *int
}

// ResultsDefinitionEdges holds the relations/edges for other nodes in the graph.
type ResultsDefinitionEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ClinicalTrial `json:"parent,omitempty"`
	// ParticipantFlowModule holds the value of the participant_flow_module edge.
	ParticipantFlowModule *ParticipantFlowModule `json:"participant_flow_module,omitempty"`
	// BaselineCharacteristicsModule holds the value of the baseline_characteristics_module edge.
	BaselineCharacteristicsModule *BaselineCharacteristicsModule `json:"baseline_characteristics_module,omitempty"`
	// OutcomeMeasuresModule holds the value of the outcome_measures_module edge.
	OutcomeMeasuresModule *OutcomeMeasuresModule `json:"outcome_measures_module,omitempty"`
	// AdverseEventsModule holds the value of the adverse_events_module edge.
	AdverseEventsModule *AdverseEventsModule `json:"adverse_events_module,omitempty"`
	// MoreInfoModule holds the value of the more_info_module edge.
	MoreInfoModule *MoreInfoModule `json:"more_info_module,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) ParentOrErr() (*ClinicalTrial, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: clinicaltrial.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ParticipantFlowModuleOrErr returns the ParticipantFlowModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) ParticipantFlowModuleOrErr() (*ParticipantFlowModule, error) {
	if e.loadedTypes[1] {
		if e.ParticipantFlowModule == nil {
			// The edge participant_flow_module was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: participantflowmodule.Label}
		}
		return e.ParticipantFlowModule, nil
	}
	return nil, &NotLoadedError{edge: "participant_flow_module"}
}

// BaselineCharacteristicsModuleOrErr returns the BaselineCharacteristicsModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) BaselineCharacteristicsModuleOrErr() (*BaselineCharacteristicsModule, error) {
	if e.loadedTypes[2] {
		if e.BaselineCharacteristicsModule == nil {
			// The edge baseline_characteristics_module was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinecharacteristicsmodule.Label}
		}
		return e.BaselineCharacteristicsModule, nil
	}
	return nil, &NotLoadedError{edge: "baseline_characteristics_module"}
}

// OutcomeMeasuresModuleOrErr returns the OutcomeMeasuresModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) OutcomeMeasuresModuleOrErr() (*OutcomeMeasuresModule, error) {
	if e.loadedTypes[3] {
		if e.OutcomeMeasuresModule == nil {
			// The edge outcome_measures_module was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomemeasuresmodule.Label}
		}
		return e.OutcomeMeasuresModule, nil
	}
	return nil, &NotLoadedError{edge: "outcome_measures_module"}
}

// AdverseEventsModuleOrErr returns the AdverseEventsModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) AdverseEventsModuleOrErr() (*AdverseEventsModule, error) {
	if e.loadedTypes[4] {
		if e.AdverseEventsModule == nil {
			// The edge adverse_events_module was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: adverseeventsmodule.Label}
		}
		return e.AdverseEventsModule, nil
	}
	return nil, &NotLoadedError{edge: "adverse_events_module"}
}

// MoreInfoModuleOrErr returns the MoreInfoModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultsDefinitionEdges) MoreInfoModuleOrErr() (*MoreInfoModule, error) {
	if e.loadedTypes[5] {
		if e.MoreInfoModule == nil {
			// The edge more_info_module was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moreinfomodule.Label}
		}
		return e.MoreInfoModule, nil
	}
	return nil, &NotLoadedError{edge: "more_info_module"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResultsDefinition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case resultsdefinition.FieldID:
			values[i] = new(sql.NullInt64)
		case resultsdefinition.ForeignKeys[0]: // clinical_trial_results_definition
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ResultsDefinition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResultsDefinition fields.
func (rd *ResultsDefinition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resultsdefinition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rd.ID = int(value.Int64)
		case resultsdefinition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field clinical_trial_results_definition", value)
			} else if value.Valid {
				rd.clinical_trial_results_definition = new(int)
				*rd.clinical_trial_results_definition = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryParent() *ClinicalTrialQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryParent(rd)
}

// QueryParticipantFlowModule queries the "participant_flow_module" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryParticipantFlowModule() *ParticipantFlowModuleQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryParticipantFlowModule(rd)
}

// QueryBaselineCharacteristicsModule queries the "baseline_characteristics_module" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryBaselineCharacteristicsModule() *BaselineCharacteristicsModuleQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryBaselineCharacteristicsModule(rd)
}

// QueryOutcomeMeasuresModule queries the "outcome_measures_module" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryOutcomeMeasuresModule() *OutcomeMeasuresModuleQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryOutcomeMeasuresModule(rd)
}

// QueryAdverseEventsModule queries the "adverse_events_module" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryAdverseEventsModule() *AdverseEventsModuleQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryAdverseEventsModule(rd)
}

// QueryMoreInfoModule queries the "more_info_module" edge of the ResultsDefinition entity.
func (rd *ResultsDefinition) QueryMoreInfoModule() *MoreInfoModuleQuery {
	return (&ResultsDefinitionClient{config: rd.config}).QueryMoreInfoModule(rd)
}

// Update returns a builder for updating this ResultsDefinition.
// Note that you need to call ResultsDefinition.Unwrap() before calling this method if this ResultsDefinition
// was returned from a transaction, and the transaction was committed or rolled back.
func (rd *ResultsDefinition) Update() *ResultsDefinitionUpdateOne {
	return (&ResultsDefinitionClient{config: rd.config}).UpdateOne(rd)
}

// Unwrap unwraps the ResultsDefinition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rd *ResultsDefinition) Unwrap() *ResultsDefinition {
	tx, ok := rd.config.driver.(*txDriver)
	if !ok {
		panic("models: ResultsDefinition is not a transactional entity")
	}
	rd.config.driver = tx.drv
	return rd
}

// String implements the fmt.Stringer.
func (rd *ResultsDefinition) String() string {
	var builder strings.Builder
	builder.WriteString("ResultsDefinition(")
	builder.WriteString(fmt.Sprintf("id=%v", rd.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ResultsDefinitions is a parsable slice of ResultsDefinition.
type ResultsDefinitions []*ResultsDefinition

func (rd ResultsDefinitions) config(cfg config) {
	for _i := range rd {
		rd[_i].config = cfg
	}
}
