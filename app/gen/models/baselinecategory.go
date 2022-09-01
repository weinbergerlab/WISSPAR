// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
)

// BaselineCategory is the model entity for the BaselineCategory schema.
type BaselineCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineCategoryTitle holds the value of the "baseline_category_title" field.
	BaselineCategoryTitle string `json:"baseline_category_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineCategoryQuery when eager-loading is set.
	Edges                                 BaselineCategoryEdges `json:"edges"`
	baseline_class_baseline_category_list *int
}

// BaselineCategoryEdges holds the relations/edges for other nodes in the graph.
type BaselineCategoryEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineClass `json:"parent,omitempty"`
	// BaselineMeasurementList holds the value of the baseline_measurement_list edge.
	BaselineMeasurementList []*BaselineMeasurement `json:"baseline_measurement_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineCategoryEdges) ParentOrErr() (*BaselineClass, error) {
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

// BaselineMeasurementListOrErr returns the BaselineMeasurementList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineCategoryEdges) BaselineMeasurementListOrErr() ([]*BaselineMeasurement, error) {
	if e.loadedTypes[1] {
		return e.BaselineMeasurementList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_measurement_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinecategory.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinecategory.FieldBaselineCategoryTitle:
			values[i] = new(sql.NullString)
		case baselinecategory.ForeignKeys[0]: // baseline_class_baseline_category_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineCategory fields.
func (bc *BaselineCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinecategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bc.ID = int(value.Int64)
		case baselinecategory.FieldBaselineCategoryTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_category_title", values[i])
			} else if value.Valid {
				bc.BaselineCategoryTitle = value.String
			}
		case baselinecategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_class_baseline_category_list", value)
			} else if value.Valid {
				bc.baseline_class_baseline_category_list = new(int)
				*bc.baseline_class_baseline_category_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineCategory entity.
func (bc *BaselineCategory) QueryParent() *BaselineClassQuery {
	return (&BaselineCategoryClient{config: bc.config}).QueryParent(bc)
}

// QueryBaselineMeasurementList queries the "baseline_measurement_list" edge of the BaselineCategory entity.
func (bc *BaselineCategory) QueryBaselineMeasurementList() *BaselineMeasurementQuery {
	return (&BaselineCategoryClient{config: bc.config}).QueryBaselineMeasurementList(bc)
}

// Update returns a builder for updating this BaselineCategory.
// Note that you need to call BaselineCategory.Unwrap() before calling this method if this BaselineCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BaselineCategory) Update() *BaselineCategoryUpdateOne {
	return (&BaselineCategoryClient{config: bc.config}).UpdateOne(bc)
}

// Unwrap unwraps the BaselineCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BaselineCategory) Unwrap() *BaselineCategory {
	tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineCategory is not a transactional entity")
	}
	bc.config.driver = tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BaselineCategory) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", bc.ID))
	builder.WriteString(", baseline_category_title=")
	builder.WriteString(bc.BaselineCategoryTitle)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineCategories is a parsable slice of BaselineCategory.
type BaselineCategories []*BaselineCategory

func (bc BaselineCategories) config(cfg config) {
	for _i := range bc {
		bc[_i].config = cfg
	}
}
