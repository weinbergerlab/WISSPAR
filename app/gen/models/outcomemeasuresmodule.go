// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// OutcomeMeasuresModule is the model entity for the OutcomeMeasuresModule schema.
type OutcomeMeasuresModule struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeMeasuresModuleQuery when eager-loading is set.
	Edges                                      OutcomeMeasuresModuleEdges `json:"edges"`
	results_definition_outcome_measures_module *int
}

// OutcomeMeasuresModuleEdges holds the relations/edges for other nodes in the graph.
type OutcomeMeasuresModuleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ResultsDefinition `json:"parent,omitempty"`
	// OutcomeMeasureList holds the value of the outcome_measure_list edge.
	OutcomeMeasureList []*OutcomeMeasure `json:"outcome_measure_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeMeasuresModuleEdges) ParentOrErr() (*ResultsDefinition, error) {
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

// OutcomeMeasureListOrErr returns the OutcomeMeasureList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeMeasuresModuleEdges) OutcomeMeasureListOrErr() ([]*OutcomeMeasure, error) {
	if e.loadedTypes[1] {
		return e.OutcomeMeasureList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_measure_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeMeasuresModule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomemeasuresmodule.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomemeasuresmodule.ForeignKeys[0]: // results_definition_outcome_measures_module
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeMeasuresModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeMeasuresModule fields.
func (omm *OutcomeMeasuresModule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomemeasuresmodule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			omm.ID = int(value.Int64)
		case outcomemeasuresmodule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field results_definition_outcome_measures_module", value)
			} else if value.Valid {
				omm.results_definition_outcome_measures_module = new(int)
				*omm.results_definition_outcome_measures_module = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeMeasuresModule entity.
func (omm *OutcomeMeasuresModule) QueryParent() *ResultsDefinitionQuery {
	return (&OutcomeMeasuresModuleClient{config: omm.config}).QueryParent(omm)
}

// QueryOutcomeMeasureList queries the "outcome_measure_list" edge of the OutcomeMeasuresModule entity.
func (omm *OutcomeMeasuresModule) QueryOutcomeMeasureList() *OutcomeMeasureQuery {
	return (&OutcomeMeasuresModuleClient{config: omm.config}).QueryOutcomeMeasureList(omm)
}

// Update returns a builder for updating this OutcomeMeasuresModule.
// Note that you need to call OutcomeMeasuresModule.Unwrap() before calling this method if this OutcomeMeasuresModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (omm *OutcomeMeasuresModule) Update() *OutcomeMeasuresModuleUpdateOne {
	return (&OutcomeMeasuresModuleClient{config: omm.config}).UpdateOne(omm)
}

// Unwrap unwraps the OutcomeMeasuresModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (omm *OutcomeMeasuresModule) Unwrap() *OutcomeMeasuresModule {
	tx, ok := omm.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeMeasuresModule is not a transactional entity")
	}
	omm.config.driver = tx.drv
	return omm
}

// String implements the fmt.Stringer.
func (omm *OutcomeMeasuresModule) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeMeasuresModule(")
	builder.WriteString(fmt.Sprintf("id=%v", omm.ID))
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeMeasuresModules is a parsable slice of OutcomeMeasuresModule.
type OutcomeMeasuresModules []*OutcomeMeasuresModule

func (omm OutcomeMeasuresModules) config(cfg config) {
	for _i := range omm {
		omm[_i].config = cfg
	}
}
