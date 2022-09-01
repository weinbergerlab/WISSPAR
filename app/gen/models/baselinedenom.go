// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
)

// BaselineDenom is the model entity for the BaselineDenom schema.
type BaselineDenom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineDenomUnits holds the value of the "baseline_denom_units" field.
	BaselineDenomUnits string `json:"baseline_denom_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineDenomQuery when eager-loading is set.
	Edges                                               BaselineDenomEdges `json:"edges"`
	baseline_characteristics_module_baseline_denom_list *int
}

// BaselineDenomEdges holds the relations/edges for other nodes in the graph.
type BaselineDenomEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineCharacteristicsModule `json:"parent,omitempty"`
	// BaselineDenomCountList holds the value of the baseline_denom_count_list edge.
	BaselineDenomCountList []*BaselineDenomCount `json:"baseline_denom_count_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineDenomEdges) ParentOrErr() (*BaselineCharacteristicsModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinecharacteristicsmodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// BaselineDenomCountListOrErr returns the BaselineDenomCountList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineDenomEdges) BaselineDenomCountListOrErr() ([]*BaselineDenomCount, error) {
	if e.loadedTypes[1] {
		return e.BaselineDenomCountList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_denom_count_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineDenom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinedenom.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinedenom.FieldBaselineDenomUnits:
			values[i] = new(sql.NullString)
		case baselinedenom.ForeignKeys[0]: // baseline_characteristics_module_baseline_denom_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineDenom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineDenom fields.
func (bd *BaselineDenom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinedenom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bd.ID = int(value.Int64)
		case baselinedenom.FieldBaselineDenomUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_denom_units", values[i])
			} else if value.Valid {
				bd.BaselineDenomUnits = value.String
			}
		case baselinedenom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_characteristics_module_baseline_denom_list", value)
			} else if value.Valid {
				bd.baseline_characteristics_module_baseline_denom_list = new(int)
				*bd.baseline_characteristics_module_baseline_denom_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineDenom entity.
func (bd *BaselineDenom) QueryParent() *BaselineCharacteristicsModuleQuery {
	return (&BaselineDenomClient{config: bd.config}).QueryParent(bd)
}

// QueryBaselineDenomCountList queries the "baseline_denom_count_list" edge of the BaselineDenom entity.
func (bd *BaselineDenom) QueryBaselineDenomCountList() *BaselineDenomCountQuery {
	return (&BaselineDenomClient{config: bd.config}).QueryBaselineDenomCountList(bd)
}

// Update returns a builder for updating this BaselineDenom.
// Note that you need to call BaselineDenom.Unwrap() before calling this method if this BaselineDenom
// was returned from a transaction, and the transaction was committed or rolled back.
func (bd *BaselineDenom) Update() *BaselineDenomUpdateOne {
	return (&BaselineDenomClient{config: bd.config}).UpdateOne(bd)
}

// Unwrap unwraps the BaselineDenom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bd *BaselineDenom) Unwrap() *BaselineDenom {
	tx, ok := bd.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineDenom is not a transactional entity")
	}
	bd.config.driver = tx.drv
	return bd
}

// String implements the fmt.Stringer.
func (bd *BaselineDenom) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineDenom(")
	builder.WriteString(fmt.Sprintf("id=%v", bd.ID))
	builder.WriteString(", baseline_denom_units=")
	builder.WriteString(bd.BaselineDenomUnits)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineDenoms is a parsable slice of BaselineDenom.
type BaselineDenoms []*BaselineDenom

func (bd BaselineDenoms) config(cfg config) {
	for _i := range bd {
		bd[_i].config = cfg
	}
}
