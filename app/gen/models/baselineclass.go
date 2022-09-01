// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
)

// BaselineClass is the model entity for the BaselineClass schema.
type BaselineClass struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineClassTitle holds the value of the "baseline_class_title" field.
	BaselineClassTitle string `json:"baseline_class_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineClassQuery when eager-loading is set.
	Edges                                BaselineClassEdges `json:"edges"`
	baseline_measure_baseline_class_list *int
}

// BaselineClassEdges holds the relations/edges for other nodes in the graph.
type BaselineClassEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineMeasure `json:"parent,omitempty"`
	// BaselineClassDenomList holds the value of the baseline_class_denom_list edge.
	BaselineClassDenomList []*BaselineClassDenom `json:"baseline_class_denom_list,omitempty"`
	// BaselineCategoryList holds the value of the baseline_category_list edge.
	BaselineCategoryList []*BaselineCategory `json:"baseline_category_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineClassEdges) ParentOrErr() (*BaselineMeasure, error) {
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

// BaselineClassDenomListOrErr returns the BaselineClassDenomList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineClassEdges) BaselineClassDenomListOrErr() ([]*BaselineClassDenom, error) {
	if e.loadedTypes[1] {
		return e.BaselineClassDenomList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_class_denom_list"}
}

// BaselineCategoryListOrErr returns the BaselineCategoryList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineClassEdges) BaselineCategoryListOrErr() ([]*BaselineCategory, error) {
	if e.loadedTypes[2] {
		return e.BaselineCategoryList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_category_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineClass) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselineclass.FieldID:
			values[i] = new(sql.NullInt64)
		case baselineclass.FieldBaselineClassTitle:
			values[i] = new(sql.NullString)
		case baselineclass.ForeignKeys[0]: // baseline_measure_baseline_class_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineClass", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineClass fields.
func (bc *BaselineClass) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselineclass.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bc.ID = int(value.Int64)
		case baselineclass.FieldBaselineClassTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_class_title", values[i])
			} else if value.Valid {
				bc.BaselineClassTitle = value.String
			}
		case baselineclass.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_measure_baseline_class_list", value)
			} else if value.Valid {
				bc.baseline_measure_baseline_class_list = new(int)
				*bc.baseline_measure_baseline_class_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineClass entity.
func (bc *BaselineClass) QueryParent() *BaselineMeasureQuery {
	return (&BaselineClassClient{config: bc.config}).QueryParent(bc)
}

// QueryBaselineClassDenomList queries the "baseline_class_denom_list" edge of the BaselineClass entity.
func (bc *BaselineClass) QueryBaselineClassDenomList() *BaselineClassDenomQuery {
	return (&BaselineClassClient{config: bc.config}).QueryBaselineClassDenomList(bc)
}

// QueryBaselineCategoryList queries the "baseline_category_list" edge of the BaselineClass entity.
func (bc *BaselineClass) QueryBaselineCategoryList() *BaselineCategoryQuery {
	return (&BaselineClassClient{config: bc.config}).QueryBaselineCategoryList(bc)
}

// Update returns a builder for updating this BaselineClass.
// Note that you need to call BaselineClass.Unwrap() before calling this method if this BaselineClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BaselineClass) Update() *BaselineClassUpdateOne {
	return (&BaselineClassClient{config: bc.config}).UpdateOne(bc)
}

// Unwrap unwraps the BaselineClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BaselineClass) Unwrap() *BaselineClass {
	tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineClass is not a transactional entity")
	}
	bc.config.driver = tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BaselineClass) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineClass(")
	builder.WriteString(fmt.Sprintf("id=%v", bc.ID))
	builder.WriteString(", baseline_class_title=")
	builder.WriteString(bc.BaselineClassTitle)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineClasses is a parsable slice of BaselineClass.
type BaselineClasses []*BaselineClass

func (bc BaselineClasses) config(cfg config) {
	for _i := range bc {
		bc[_i].config = cfg
	}
}
