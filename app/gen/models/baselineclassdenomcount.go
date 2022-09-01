// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
)

// BaselineClassDenomCount is the model entity for the BaselineClassDenomCount schema.
type BaselineClassDenomCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineClassDenomCountGroupID holds the value of the "baseline_class_denom_count_group_id" field.
	BaselineClassDenomCountGroupID string `json:"baseline_class_denom_count_group_id,omitempty"`
	// BaselineClassDenomCountValue holds the value of the "baseline_class_denom_count_value" field.
	BaselineClassDenomCountValue string `json:"baseline_class_denom_count_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineClassDenomCountQuery when eager-loading is set.
	Edges                                                BaselineClassDenomCountEdges `json:"edges"`
	baseline_class_denom_baseline_class_denom_count_list *int
}

// BaselineClassDenomCountEdges holds the relations/edges for other nodes in the graph.
type BaselineClassDenomCountEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineClassDenom `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineClassDenomCountEdges) ParentOrErr() (*BaselineClassDenom, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselineclassdenom.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineClassDenomCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselineclassdenomcount.FieldID:
			values[i] = new(sql.NullInt64)
		case baselineclassdenomcount.FieldBaselineClassDenomCountGroupID, baselineclassdenomcount.FieldBaselineClassDenomCountValue:
			values[i] = new(sql.NullString)
		case baselineclassdenomcount.ForeignKeys[0]: // baseline_class_denom_baseline_class_denom_count_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineClassDenomCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineClassDenomCount fields.
func (bcdc *BaselineClassDenomCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselineclassdenomcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bcdc.ID = int(value.Int64)
		case baselineclassdenomcount.FieldBaselineClassDenomCountGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_class_denom_count_group_id", values[i])
			} else if value.Valid {
				bcdc.BaselineClassDenomCountGroupID = value.String
			}
		case baselineclassdenomcount.FieldBaselineClassDenomCountValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_class_denom_count_value", values[i])
			} else if value.Valid {
				bcdc.BaselineClassDenomCountValue = value.String
			}
		case baselineclassdenomcount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_class_denom_baseline_class_denom_count_list", value)
			} else if value.Valid {
				bcdc.baseline_class_denom_baseline_class_denom_count_list = new(int)
				*bcdc.baseline_class_denom_baseline_class_denom_count_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineClassDenomCount entity.
func (bcdc *BaselineClassDenomCount) QueryParent() *BaselineClassDenomQuery {
	return (&BaselineClassDenomCountClient{config: bcdc.config}).QueryParent(bcdc)
}

// Update returns a builder for updating this BaselineClassDenomCount.
// Note that you need to call BaselineClassDenomCount.Unwrap() before calling this method if this BaselineClassDenomCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (bcdc *BaselineClassDenomCount) Update() *BaselineClassDenomCountUpdateOne {
	return (&BaselineClassDenomCountClient{config: bcdc.config}).UpdateOne(bcdc)
}

// Unwrap unwraps the BaselineClassDenomCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bcdc *BaselineClassDenomCount) Unwrap() *BaselineClassDenomCount {
	tx, ok := bcdc.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineClassDenomCount is not a transactional entity")
	}
	bcdc.config.driver = tx.drv
	return bcdc
}

// String implements the fmt.Stringer.
func (bcdc *BaselineClassDenomCount) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineClassDenomCount(")
	builder.WriteString(fmt.Sprintf("id=%v", bcdc.ID))
	builder.WriteString(", baseline_class_denom_count_group_id=")
	builder.WriteString(bcdc.BaselineClassDenomCountGroupID)
	builder.WriteString(", baseline_class_denom_count_value=")
	builder.WriteString(bcdc.BaselineClassDenomCountValue)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineClassDenomCounts is a parsable slice of BaselineClassDenomCount.
type BaselineClassDenomCounts []*BaselineClassDenomCount

func (bcdc BaselineClassDenomCounts) config(cfg config) {
	for _i := range bcdc {
		bcdc[_i].config = cfg
	}
}
