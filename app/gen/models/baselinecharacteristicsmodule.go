// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// BaselineCharacteristicsModule is the model entity for the BaselineCharacteristicsModule schema.
type BaselineCharacteristicsModule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselinePopulationDescription holds the value of the "baseline_population_description" field.
	BaselinePopulationDescription string `json:"baseline_population_description,omitempty"`
	// BaselineTypeUnitsAnalyzed holds the value of the "baseline_type_units_analyzed" field.
	BaselineTypeUnitsAnalyzed string `json:"baseline_type_units_analyzed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineCharacteristicsModuleQuery when eager-loading is set.
	Edges                                              BaselineCharacteristicsModuleEdges `json:"edges"`
	results_definition_baseline_characteristics_module *int
}

// BaselineCharacteristicsModuleEdges holds the relations/edges for other nodes in the graph.
type BaselineCharacteristicsModuleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ResultsDefinition `json:"parent,omitempty"`
	// BaselineGroupList holds the value of the baseline_group_list edge.
	BaselineGroupList []*BaselineGroup `json:"baseline_group_list,omitempty"`
	// BaselineDenomList holds the value of the baseline_denom_list edge.
	BaselineDenomList []*BaselineDenom `json:"baseline_denom_list,omitempty"`
	// BaselineMeasureList holds the value of the baseline_measure_list edge.
	BaselineMeasureList []*BaselineMeasure `json:"baseline_measure_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineCharacteristicsModuleEdges) ParentOrErr() (*ResultsDefinition, error) {
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

// BaselineGroupListOrErr returns the BaselineGroupList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineCharacteristicsModuleEdges) BaselineGroupListOrErr() ([]*BaselineGroup, error) {
	if e.loadedTypes[1] {
		return e.BaselineGroupList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_group_list"}
}

// BaselineDenomListOrErr returns the BaselineDenomList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineCharacteristicsModuleEdges) BaselineDenomListOrErr() ([]*BaselineDenom, error) {
	if e.loadedTypes[2] {
		return e.BaselineDenomList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_denom_list"}
}

// BaselineMeasureListOrErr returns the BaselineMeasureList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineCharacteristicsModuleEdges) BaselineMeasureListOrErr() ([]*BaselineMeasure, error) {
	if e.loadedTypes[3] {
		return e.BaselineMeasureList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_measure_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineCharacteristicsModule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinecharacteristicsmodule.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinecharacteristicsmodule.FieldBaselinePopulationDescription, baselinecharacteristicsmodule.FieldBaselineTypeUnitsAnalyzed:
			values[i] = new(sql.NullString)
		case baselinecharacteristicsmodule.ForeignKeys[0]: // results_definition_baseline_characteristics_module
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineCharacteristicsModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineCharacteristicsModule fields.
func (bcm *BaselineCharacteristicsModule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinecharacteristicsmodule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bcm.ID = int(value.Int64)
		case baselinecharacteristicsmodule.FieldBaselinePopulationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_population_description", values[i])
			} else if value.Valid {
				bcm.BaselinePopulationDescription = value.String
			}
		case baselinecharacteristicsmodule.FieldBaselineTypeUnitsAnalyzed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_type_units_analyzed", values[i])
			} else if value.Valid {
				bcm.BaselineTypeUnitsAnalyzed = value.String
			}
		case baselinecharacteristicsmodule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field results_definition_baseline_characteristics_module", value)
			} else if value.Valid {
				bcm.results_definition_baseline_characteristics_module = new(int)
				*bcm.results_definition_baseline_characteristics_module = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineCharacteristicsModule entity.
func (bcm *BaselineCharacteristicsModule) QueryParent() *ResultsDefinitionQuery {
	return (&BaselineCharacteristicsModuleClient{config: bcm.config}).QueryParent(bcm)
}

// QueryBaselineGroupList queries the "baseline_group_list" edge of the BaselineCharacteristicsModule entity.
func (bcm *BaselineCharacteristicsModule) QueryBaselineGroupList() *BaselineGroupQuery {
	return (&BaselineCharacteristicsModuleClient{config: bcm.config}).QueryBaselineGroupList(bcm)
}

// QueryBaselineDenomList queries the "baseline_denom_list" edge of the BaselineCharacteristicsModule entity.
func (bcm *BaselineCharacteristicsModule) QueryBaselineDenomList() *BaselineDenomQuery {
	return (&BaselineCharacteristicsModuleClient{config: bcm.config}).QueryBaselineDenomList(bcm)
}

// QueryBaselineMeasureList queries the "baseline_measure_list" edge of the BaselineCharacteristicsModule entity.
func (bcm *BaselineCharacteristicsModule) QueryBaselineMeasureList() *BaselineMeasureQuery {
	return (&BaselineCharacteristicsModuleClient{config: bcm.config}).QueryBaselineMeasureList(bcm)
}

// Update returns a builder for updating this BaselineCharacteristicsModule.
// Note that you need to call BaselineCharacteristicsModule.Unwrap() before calling this method if this BaselineCharacteristicsModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (bcm *BaselineCharacteristicsModule) Update() *BaselineCharacteristicsModuleUpdateOne {
	return (&BaselineCharacteristicsModuleClient{config: bcm.config}).UpdateOne(bcm)
}

// Unwrap unwraps the BaselineCharacteristicsModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bcm *BaselineCharacteristicsModule) Unwrap() *BaselineCharacteristicsModule {
	tx, ok := bcm.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineCharacteristicsModule is not a transactional entity")
	}
	bcm.config.driver = tx.drv
	return bcm
}

// String implements the fmt.Stringer.
func (bcm *BaselineCharacteristicsModule) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineCharacteristicsModule(")
	builder.WriteString(fmt.Sprintf("id=%v", bcm.ID))
	builder.WriteString(", baseline_population_description=")
	builder.WriteString(bcm.BaselinePopulationDescription)
	builder.WriteString(", baseline_type_units_analyzed=")
	builder.WriteString(bcm.BaselineTypeUnitsAnalyzed)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineCharacteristicsModules is a parsable slice of BaselineCharacteristicsModule.
type BaselineCharacteristicsModules []*BaselineCharacteristicsModule

func (bcm BaselineCharacteristicsModules) config(cfg config) {
	for _i := range bcm {
		bcm[_i].config = cfg
	}
}
