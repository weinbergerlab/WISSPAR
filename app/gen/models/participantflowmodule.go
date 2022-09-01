// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ParticipantFlowModule is the model entity for the ParticipantFlowModule schema.
type ParticipantFlowModule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlowPreAssignmentDetails holds the value of the "flow_pre_assignment_details" field.
	FlowPreAssignmentDetails string `json:"flow_pre_assignment_details,omitempty"`
	// FlowRecruitmentDetails holds the value of the "flow_recruitment_details" field.
	FlowRecruitmentDetails string `json:"flow_recruitment_details,omitempty"`
	// FlowTypeUnitsAnalyzed holds the value of the "flow_type_units_analyzed" field.
	FlowTypeUnitsAnalyzed string `json:"flow_type_units_analyzed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ParticipantFlowModuleQuery when eager-loading is set.
	Edges                                      ParticipantFlowModuleEdges `json:"edges"`
	results_definition_participant_flow_module *int
}

// ParticipantFlowModuleEdges holds the relations/edges for other nodes in the graph.
type ParticipantFlowModuleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ResultsDefinition `json:"parent,omitempty"`
	// FlowGroupList holds the value of the flow_group_list edge.
	FlowGroupList []*FlowGroup `json:"flow_group_list,omitempty"`
	// FlowPeriodList holds the value of the flow_period_list edge.
	FlowPeriodList []*FlowPeriod `json:"flow_period_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ParticipantFlowModuleEdges) ParentOrErr() (*ResultsDefinition, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resultsdefinition.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// FlowGroupListOrErr returns the FlowGroupList value or an error if the edge
// was not loaded in eager-loading.
func (e ParticipantFlowModuleEdges) FlowGroupListOrErr() ([]*FlowGroup, error) {
	if e.loadedTypes[1] {
		return e.FlowGroupList, nil
	}
	return nil, &NotLoadedError{edge: "flow_group_list"}
}

// FlowPeriodListOrErr returns the FlowPeriodList value or an error if the edge
// was not loaded in eager-loading.
func (e ParticipantFlowModuleEdges) FlowPeriodListOrErr() ([]*FlowPeriod, error) {
	if e.loadedTypes[2] {
		return e.FlowPeriodList, nil
	}
	return nil, &NotLoadedError{edge: "flow_period_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ParticipantFlowModule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case participantflowmodule.FieldID:
			values[i] = new(sql.NullInt64)
		case participantflowmodule.FieldFlowPreAssignmentDetails, participantflowmodule.FieldFlowRecruitmentDetails, participantflowmodule.FieldFlowTypeUnitsAnalyzed:
			values[i] = new(sql.NullString)
		case participantflowmodule.ForeignKeys[0]: // results_definition_participant_flow_module
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ParticipantFlowModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ParticipantFlowModule fields.
func (pfm *ParticipantFlowModule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case participantflowmodule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pfm.ID = int(value.Int64)
		case participantflowmodule.FieldFlowPreAssignmentDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_pre_assignment_details", values[i])
			} else if value.Valid {
				pfm.FlowPreAssignmentDetails = value.String
			}
		case participantflowmodule.FieldFlowRecruitmentDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_recruitment_details", values[i])
			} else if value.Valid {
				pfm.FlowRecruitmentDetails = value.String
			}
		case participantflowmodule.FieldFlowTypeUnitsAnalyzed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_type_units_analyzed", values[i])
			} else if value.Valid {
				pfm.FlowTypeUnitsAnalyzed = value.String
			}
		case participantflowmodule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field results_definition_participant_flow_module", value)
			} else if value.Valid {
				pfm.results_definition_participant_flow_module = new(int)
				*pfm.results_definition_participant_flow_module = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ParticipantFlowModule entity.
func (pfm *ParticipantFlowModule) QueryParent() *ResultsDefinitionQuery {
	return (&ParticipantFlowModuleClient{config: pfm.config}).QueryParent(pfm)
}

// QueryFlowGroupList queries the "flow_group_list" edge of the ParticipantFlowModule entity.
func (pfm *ParticipantFlowModule) QueryFlowGroupList() *FlowGroupQuery {
	return (&ParticipantFlowModuleClient{config: pfm.config}).QueryFlowGroupList(pfm)
}

// QueryFlowPeriodList queries the "flow_period_list" edge of the ParticipantFlowModule entity.
func (pfm *ParticipantFlowModule) QueryFlowPeriodList() *FlowPeriodQuery {
	return (&ParticipantFlowModuleClient{config: pfm.config}).QueryFlowPeriodList(pfm)
}

// Update returns a builder for updating this ParticipantFlowModule.
// Note that you need to call ParticipantFlowModule.Unwrap() before calling this method if this ParticipantFlowModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (pfm *ParticipantFlowModule) Update() *ParticipantFlowModuleUpdateOne {
	return (&ParticipantFlowModuleClient{config: pfm.config}).UpdateOne(pfm)
}

// Unwrap unwraps the ParticipantFlowModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pfm *ParticipantFlowModule) Unwrap() *ParticipantFlowModule {
	tx, ok := pfm.config.driver.(*txDriver)
	if !ok {
		panic("models: ParticipantFlowModule is not a transactional entity")
	}
	pfm.config.driver = tx.drv
	return pfm
}

// String implements the fmt.Stringer.
func (pfm *ParticipantFlowModule) String() string {
	var builder strings.Builder
	builder.WriteString("ParticipantFlowModule(")
	builder.WriteString(fmt.Sprintf("id=%v", pfm.ID))
	builder.WriteString(", flow_pre_assignment_details=")
	builder.WriteString(pfm.FlowPreAssignmentDetails)
	builder.WriteString(", flow_recruitment_details=")
	builder.WriteString(pfm.FlowRecruitmentDetails)
	builder.WriteString(", flow_type_units_analyzed=")
	builder.WriteString(pfm.FlowTypeUnitsAnalyzed)
	builder.WriteByte(')')
	return builder.String()
}

// ParticipantFlowModules is a parsable slice of ParticipantFlowModule.
type ParticipantFlowModules []*ParticipantFlowModule

func (pfm ParticipantFlowModules) config(cfg config) {
	for _i := range pfm {
		pfm[_i].config = cfg
	}
}
