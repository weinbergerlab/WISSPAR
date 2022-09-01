// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
)

// BaselineMeasureDenom is the model entity for the BaselineMeasureDenom schema.
type BaselineMeasureDenom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineMeasureDenomUnits holds the value of the "baseline_measure_denom_units" field.
	BaselineMeasureDenomUnits string `json:"baseline_measure_denom_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineMeasureDenomQuery when eager-loading is set.
	Edges                                        BaselineMeasureDenomEdges `json:"edges"`
	baseline_measure_baseline_measure_denom_list *int
}

// BaselineMeasureDenomEdges holds the relations/edges for other nodes in the graph.
type BaselineMeasureDenomEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineMeasure `json:"parent,omitempty"`
	// BaselineMeasureDenomCountList holds the value of the baseline_measure_denom_count_list edge.
	BaselineMeasureDenomCountList []*BaselineMeasureDenomCount `json:"baseline_measure_denom_count_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineMeasureDenomEdges) ParentOrErr() (*BaselineMeasure, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinemeasure.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// BaselineMeasureDenomCountListOrErr returns the BaselineMeasureDenomCountList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineMeasureDenomEdges) BaselineMeasureDenomCountListOrErr() ([]*BaselineMeasureDenomCount, error) {
	if e.loadedTypes[1] {
		return e.BaselineMeasureDenomCountList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_measure_denom_count_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineMeasureDenom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinemeasuredenom.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinemeasuredenom.FieldBaselineMeasureDenomUnits:
			values[i] = new(sql.NullString)
		case baselinemeasuredenom.ForeignKeys[0]: // baseline_measure_baseline_measure_denom_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineMeasureDenom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineMeasureDenom fields.
func (bmd *BaselineMeasureDenom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinemeasuredenom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bmd.ID = int(value.Int64)
		case baselinemeasuredenom.FieldBaselineMeasureDenomUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_denom_units", values[i])
			} else if value.Valid {
				bmd.BaselineMeasureDenomUnits = value.String
			}
		case baselinemeasuredenom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_measure_baseline_measure_denom_list", value)
			} else if value.Valid {
				bmd.baseline_measure_baseline_measure_denom_list = new(int)
				*bmd.baseline_measure_baseline_measure_denom_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineMeasureDenom entity.
func (bmd *BaselineMeasureDenom) QueryParent() *BaselineMeasureQuery {
	return (&BaselineMeasureDenomClient{config: bmd.config}).QueryParent(bmd)
}

// QueryBaselineMeasureDenomCountList queries the "baseline_measure_denom_count_list" edge of the BaselineMeasureDenom entity.
func (bmd *BaselineMeasureDenom) QueryBaselineMeasureDenomCountList() *BaselineMeasureDenomCountQuery {
	return (&BaselineMeasureDenomClient{config: bmd.config}).QueryBaselineMeasureDenomCountList(bmd)
}

// Update returns a builder for updating this BaselineMeasureDenom.
// Note that you need to call BaselineMeasureDenom.Unwrap() before calling this method if this BaselineMeasureDenom
// was returned from a transaction, and the transaction was committed or rolled back.
func (bmd *BaselineMeasureDenom) Update() *BaselineMeasureDenomUpdateOne {
	return (&BaselineMeasureDenomClient{config: bmd.config}).UpdateOne(bmd)
}

// Unwrap unwraps the BaselineMeasureDenom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bmd *BaselineMeasureDenom) Unwrap() *BaselineMeasureDenom {
	tx, ok := bmd.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineMeasureDenom is not a transactional entity")
	}
	bmd.config.driver = tx.drv
	return bmd
}

// String implements the fmt.Stringer.
func (bmd *BaselineMeasureDenom) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineMeasureDenom(")
	builder.WriteString(fmt.Sprintf("id=%v", bmd.ID))
	builder.WriteString(", baseline_measure_denom_units=")
	builder.WriteString(bmd.BaselineMeasureDenomUnits)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineMeasureDenoms is a parsable slice of BaselineMeasureDenom.
type BaselineMeasureDenoms []*BaselineMeasureDenom

func (bmd BaselineMeasureDenoms) config(cfg config) {
	for _i := range bmd {
		bmd[_i].config = cfg
	}
}
