// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
)

// BaselineMeasurement is the model entity for the BaselineMeasurement schema.
type BaselineMeasurement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineMeasurementGroupID holds the value of the "baseline_measurement_group_id" field.
	BaselineMeasurementGroupID string `json:"baseline_measurement_group_id,omitempty"`
	// BaselineMeasurementValue holds the value of the "baseline_measurement_value" field.
	BaselineMeasurementValue string `json:"baseline_measurement_value,omitempty"`
	// BaselineMeasurementSpread holds the value of the "baseline_measurement_spread" field.
	BaselineMeasurementSpread string `json:"baseline_measurement_spread,omitempty"`
	// BaselineMeasurementLowerLimit holds the value of the "baseline_measurement_lower_limit" field.
	BaselineMeasurementLowerLimit string `json:"baseline_measurement_lower_limit,omitempty"`
	// BaselineMeasurementUpperLimit holds the value of the "baseline_measurement_upper_limit" field.
	BaselineMeasurementUpperLimit string `json:"baseline_measurement_upper_limit,omitempty"`
	// BaselineMeasurementComment holds the value of the "baseline_measurement_comment" field.
	BaselineMeasurementComment string `json:"baseline_measurement_comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineMeasurementQuery when eager-loading is set.
	Edges                                       BaselineMeasurementEdges `json:"edges"`
	baseline_category_baseline_measurement_list *int
}

// BaselineMeasurementEdges holds the relations/edges for other nodes in the graph.
type BaselineMeasurementEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineCategory `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineMeasurementEdges) ParentOrErr() (*BaselineCategory, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinecategory.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineMeasurement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinemeasurement.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinemeasurement.FieldBaselineMeasurementGroupID, baselinemeasurement.FieldBaselineMeasurementValue, baselinemeasurement.FieldBaselineMeasurementSpread, baselinemeasurement.FieldBaselineMeasurementLowerLimit, baselinemeasurement.FieldBaselineMeasurementUpperLimit, baselinemeasurement.FieldBaselineMeasurementComment:
			values[i] = new(sql.NullString)
		case baselinemeasurement.ForeignKeys[0]: // baseline_category_baseline_measurement_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineMeasurement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineMeasurement fields.
func (bm *BaselineMeasurement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinemeasurement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bm.ID = int(value.Int64)
		case baselinemeasurement.FieldBaselineMeasurementGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_group_id", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementGroupID = value.String
			}
		case baselinemeasurement.FieldBaselineMeasurementValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_value", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementValue = value.String
			}
		case baselinemeasurement.FieldBaselineMeasurementSpread:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_spread", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementSpread = value.String
			}
		case baselinemeasurement.FieldBaselineMeasurementLowerLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_lower_limit", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementLowerLimit = value.String
			}
		case baselinemeasurement.FieldBaselineMeasurementUpperLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_upper_limit", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementUpperLimit = value.String
			}
		case baselinemeasurement.FieldBaselineMeasurementComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measurement_comment", values[i])
			} else if value.Valid {
				bm.BaselineMeasurementComment = value.String
			}
		case baselinemeasurement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_category_baseline_measurement_list", value)
			} else if value.Valid {
				bm.baseline_category_baseline_measurement_list = new(int)
				*bm.baseline_category_baseline_measurement_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineMeasurement entity.
func (bm *BaselineMeasurement) QueryParent() *BaselineCategoryQuery {
	return (&BaselineMeasurementClient{config: bm.config}).QueryParent(bm)
}

// Update returns a builder for updating this BaselineMeasurement.
// Note that you need to call BaselineMeasurement.Unwrap() before calling this method if this BaselineMeasurement
// was returned from a transaction, and the transaction was committed or rolled back.
func (bm *BaselineMeasurement) Update() *BaselineMeasurementUpdateOne {
	return (&BaselineMeasurementClient{config: bm.config}).UpdateOne(bm)
}

// Unwrap unwraps the BaselineMeasurement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bm *BaselineMeasurement) Unwrap() *BaselineMeasurement {
	tx, ok := bm.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineMeasurement is not a transactional entity")
	}
	bm.config.driver = tx.drv
	return bm
}

// String implements the fmt.Stringer.
func (bm *BaselineMeasurement) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineMeasurement(")
	builder.WriteString(fmt.Sprintf("id=%v", bm.ID))
	builder.WriteString(", baseline_measurement_group_id=")
	builder.WriteString(bm.BaselineMeasurementGroupID)
	builder.WriteString(", baseline_measurement_value=")
	builder.WriteString(bm.BaselineMeasurementValue)
	builder.WriteString(", baseline_measurement_spread=")
	builder.WriteString(bm.BaselineMeasurementSpread)
	builder.WriteString(", baseline_measurement_lower_limit=")
	builder.WriteString(bm.BaselineMeasurementLowerLimit)
	builder.WriteString(", baseline_measurement_upper_limit=")
	builder.WriteString(bm.BaselineMeasurementUpperLimit)
	builder.WriteString(", baseline_measurement_comment=")
	builder.WriteString(bm.BaselineMeasurementComment)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineMeasurements is a parsable slice of BaselineMeasurement.
type BaselineMeasurements []*BaselineMeasurement

func (bm BaselineMeasurements) config(cfg config) {
	for _i := range bm {
		bm[_i].config = cfg
	}
}
