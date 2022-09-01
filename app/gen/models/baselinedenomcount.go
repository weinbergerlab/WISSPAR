// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
)

// BaselineDenomCount is the model entity for the BaselineDenomCount schema.
type BaselineDenomCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineDenomCountGroupID holds the value of the "baseline_denom_count_group_id" field.
	BaselineDenomCountGroupID string `json:"baseline_denom_count_group_id,omitempty"`
	// BaselineDenomCountValue holds the value of the "baseline_denom_count_value" field.
	BaselineDenomCountValue string `json:"baseline_denom_count_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineDenomCountQuery when eager-loading is set.
	Edges                                    BaselineDenomCountEdges `json:"edges"`
	baseline_denom_baseline_denom_count_list *int
}

// BaselineDenomCountEdges holds the relations/edges for other nodes in the graph.
type BaselineDenomCountEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineDenom `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineDenomCountEdges) ParentOrErr() (*BaselineDenom, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinedenom.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineDenomCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinedenomcount.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinedenomcount.FieldBaselineDenomCountGroupID, baselinedenomcount.FieldBaselineDenomCountValue:
			values[i] = new(sql.NullString)
		case baselinedenomcount.ForeignKeys[0]: // baseline_denom_baseline_denom_count_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineDenomCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineDenomCount fields.
func (bdc *BaselineDenomCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinedenomcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bdc.ID = int(value.Int64)
		case baselinedenomcount.FieldBaselineDenomCountGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_denom_count_group_id", values[i])
			} else if value.Valid {
				bdc.BaselineDenomCountGroupID = value.String
			}
		case baselinedenomcount.FieldBaselineDenomCountValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_denom_count_value", values[i])
			} else if value.Valid {
				bdc.BaselineDenomCountValue = value.String
			}
		case baselinedenomcount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_denom_baseline_denom_count_list", value)
			} else if value.Valid {
				bdc.baseline_denom_baseline_denom_count_list = new(int)
				*bdc.baseline_denom_baseline_denom_count_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineDenomCount entity.
func (bdc *BaselineDenomCount) QueryParent() *BaselineDenomQuery {
	return (&BaselineDenomCountClient{config: bdc.config}).QueryParent(bdc)
}

// Update returns a builder for updating this BaselineDenomCount.
// Note that you need to call BaselineDenomCount.Unwrap() before calling this method if this BaselineDenomCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (bdc *BaselineDenomCount) Update() *BaselineDenomCountUpdateOne {
	return (&BaselineDenomCountClient{config: bdc.config}).UpdateOne(bdc)
}

// Unwrap unwraps the BaselineDenomCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bdc *BaselineDenomCount) Unwrap() *BaselineDenomCount {
	tx, ok := bdc.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineDenomCount is not a transactional entity")
	}
	bdc.config.driver = tx.drv
	return bdc
}

// String implements the fmt.Stringer.
func (bdc *BaselineDenomCount) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineDenomCount(")
	builder.WriteString(fmt.Sprintf("id=%v", bdc.ID))
	builder.WriteString(", baseline_denom_count_group_id=")
	builder.WriteString(bdc.BaselineDenomCountGroupID)
	builder.WriteString(", baseline_denom_count_value=")
	builder.WriteString(bdc.BaselineDenomCountValue)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineDenomCounts is a parsable slice of BaselineDenomCount.
type BaselineDenomCounts []*BaselineDenomCount

func (bdc BaselineDenomCounts) config(cfg config) {
	for _i := range bdc {
		bdc[_i].config = cfg
	}
}
