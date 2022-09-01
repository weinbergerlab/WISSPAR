// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
)

// BaselineMeasureDenomCount is the model entity for the BaselineMeasureDenomCount schema.
type BaselineMeasureDenomCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineMeasureDenomCountGroupID holds the value of the "baseline_measure_denom_count_group_id" field.
	BaselineMeasureDenomCountGroupID string `json:"baseline_measure_denom_count_group_id,omitempty"`
	// BaselineMeasureDenomCountValue holds the value of the "baseline_measure_denom_count_value" field.
	BaselineMeasureDenomCountValue string `json:"baseline_measure_denom_count_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineMeasureDenomCountQuery when eager-loading is set.
	Edges                                                    BaselineMeasureDenomCountEdges `json:"edges"`
	baseline_measure_denom_baseline_measure_denom_count_list *int
}

// BaselineMeasureDenomCountEdges holds the relations/edges for other nodes in the graph.
type BaselineMeasureDenomCountEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineMeasureDenom `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineMeasureDenomCountEdges) ParentOrErr() (*BaselineMeasureDenom, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinemeasuredenom.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineMeasureDenomCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinemeasuredenomcount.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID, baselinemeasuredenomcount.FieldBaselineMeasureDenomCountValue:
			values[i] = new(sql.NullString)
		case baselinemeasuredenomcount.ForeignKeys[0]: // baseline_measure_denom_baseline_measure_denom_count_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineMeasureDenomCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineMeasureDenomCount fields.
func (bmdc *BaselineMeasureDenomCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinemeasuredenomcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bmdc.ID = int(value.Int64)
		case baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_denom_count_group_id", values[i])
			} else if value.Valid {
				bmdc.BaselineMeasureDenomCountGroupID = value.String
			}
		case baselinemeasuredenomcount.FieldBaselineMeasureDenomCountValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_denom_count_value", values[i])
			} else if value.Valid {
				bmdc.BaselineMeasureDenomCountValue = value.String
			}
		case baselinemeasuredenomcount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_measure_denom_baseline_measure_denom_count_list", value)
			} else if value.Valid {
				bmdc.baseline_measure_denom_baseline_measure_denom_count_list = new(int)
				*bmdc.baseline_measure_denom_baseline_measure_denom_count_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineMeasureDenomCount entity.
func (bmdc *BaselineMeasureDenomCount) QueryParent() *BaselineMeasureDenomQuery {
	return (&BaselineMeasureDenomCountClient{config: bmdc.config}).QueryParent(bmdc)
}

// Update returns a builder for updating this BaselineMeasureDenomCount.
// Note that you need to call BaselineMeasureDenomCount.Unwrap() before calling this method if this BaselineMeasureDenomCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (bmdc *BaselineMeasureDenomCount) Update() *BaselineMeasureDenomCountUpdateOne {
	return (&BaselineMeasureDenomCountClient{config: bmdc.config}).UpdateOne(bmdc)
}

// Unwrap unwraps the BaselineMeasureDenomCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bmdc *BaselineMeasureDenomCount) Unwrap() *BaselineMeasureDenomCount {
	tx, ok := bmdc.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineMeasureDenomCount is not a transactional entity")
	}
	bmdc.config.driver = tx.drv
	return bmdc
}

// String implements the fmt.Stringer.
func (bmdc *BaselineMeasureDenomCount) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineMeasureDenomCount(")
	builder.WriteString(fmt.Sprintf("id=%v", bmdc.ID))
	builder.WriteString(", baseline_measure_denom_count_group_id=")
	builder.WriteString(bmdc.BaselineMeasureDenomCountGroupID)
	builder.WriteString(", baseline_measure_denom_count_value=")
	builder.WriteString(bmdc.BaselineMeasureDenomCountValue)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineMeasureDenomCounts is a parsable slice of BaselineMeasureDenomCount.
type BaselineMeasureDenomCounts []*BaselineMeasureDenomCount

func (bmdc BaselineMeasureDenomCounts) config(cfg config) {
	for _i := range bmdc {
		bmdc[_i].config = cfg
	}
}
