// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
)

// BaselineClassDenom is the model entity for the BaselineClassDenom schema.
type BaselineClassDenom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineClassDenomUnits holds the value of the "baseline_class_denom_units" field.
	BaselineClassDenomUnits string `json:"baseline_class_denom_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineClassDenomQuery when eager-loading is set.
	Edges                                    BaselineClassDenomEdges `json:"edges"`
	baseline_class_baseline_class_denom_list *int
}

// BaselineClassDenomEdges holds the relations/edges for other nodes in the graph.
type BaselineClassDenomEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineClass `json:"parent,omitempty"`
	// BaselineClassDenomCountList holds the value of the baseline_class_denom_count_list edge.
	BaselineClassDenomCountList []*BaselineClassDenomCount `json:"baseline_class_denom_count_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineClassDenomEdges) ParentOrErr() (*BaselineClass, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselineclass.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// BaselineClassDenomCountListOrErr returns the BaselineClassDenomCountList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineClassDenomEdges) BaselineClassDenomCountListOrErr() ([]*BaselineClassDenomCount, error) {
	if e.loadedTypes[1] {
		return e.BaselineClassDenomCountList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_class_denom_count_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineClassDenom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselineclassdenom.FieldID:
			values[i] = new(sql.NullInt64)
		case baselineclassdenom.FieldBaselineClassDenomUnits:
			values[i] = new(sql.NullString)
		case baselineclassdenom.ForeignKeys[0]: // baseline_class_baseline_class_denom_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineClassDenom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineClassDenom fields.
func (bcd *BaselineClassDenom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselineclassdenom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bcd.ID = int(value.Int64)
		case baselineclassdenom.FieldBaselineClassDenomUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_class_denom_units", values[i])
			} else if value.Valid {
				bcd.BaselineClassDenomUnits = value.String
			}
		case baselineclassdenom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_class_baseline_class_denom_list", value)
			} else if value.Valid {
				bcd.baseline_class_baseline_class_denom_list = new(int)
				*bcd.baseline_class_baseline_class_denom_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineClassDenom entity.
func (bcd *BaselineClassDenom) QueryParent() *BaselineClassQuery {
	return (&BaselineClassDenomClient{config: bcd.config}).QueryParent(bcd)
}

// QueryBaselineClassDenomCountList queries the "baseline_class_denom_count_list" edge of the BaselineClassDenom entity.
func (bcd *BaselineClassDenom) QueryBaselineClassDenomCountList() *BaselineClassDenomCountQuery {
	return (&BaselineClassDenomClient{config: bcd.config}).QueryBaselineClassDenomCountList(bcd)
}

// Update returns a builder for updating this BaselineClassDenom.
// Note that you need to call BaselineClassDenom.Unwrap() before calling this method if this BaselineClassDenom
// was returned from a transaction, and the transaction was committed or rolled back.
func (bcd *BaselineClassDenom) Update() *BaselineClassDenomUpdateOne {
	return (&BaselineClassDenomClient{config: bcd.config}).UpdateOne(bcd)
}

// Unwrap unwraps the BaselineClassDenom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bcd *BaselineClassDenom) Unwrap() *BaselineClassDenom {
	tx, ok := bcd.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineClassDenom is not a transactional entity")
	}
	bcd.config.driver = tx.drv
	return bcd
}

// String implements the fmt.Stringer.
func (bcd *BaselineClassDenom) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineClassDenom(")
	builder.WriteString(fmt.Sprintf("id=%v", bcd.ID))
	builder.WriteString(", baseline_class_denom_units=")
	builder.WriteString(bcd.BaselineClassDenomUnits)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineClassDenoms is a parsable slice of BaselineClassDenom.
type BaselineClassDenoms []*BaselineClassDenom

func (bcd BaselineClassDenoms) config(cfg config) {
	for _i := range bcd {
		bcd[_i].config = cfg
	}
}
