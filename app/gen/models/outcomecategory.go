// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
)

// OutcomeCategory is the model entity for the OutcomeCategory schema.
type OutcomeCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeCategoryTitle holds the value of the "outcome_category_title" field.
	OutcomeCategoryTitle string `json:"outcome_category_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeCategoryQuery when eager-loading is set.
	Edges                               OutcomeCategoryEdges `json:"edges"`
	outcome_class_outcome_category_list *int
}

// OutcomeCategoryEdges holds the relations/edges for other nodes in the graph.
type OutcomeCategoryEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeClass `json:"parent,omitempty"`
	// OutcomeMeasurementList holds the value of the outcome_measurement_list edge.
	OutcomeMeasurementList []*OutcomeMeasurement `json:"outcome_measurement_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeCategoryEdges) ParentOrErr() (*OutcomeClass, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomeclass.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// OutcomeMeasurementListOrErr returns the OutcomeMeasurementList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeCategoryEdges) OutcomeMeasurementListOrErr() ([]*OutcomeMeasurement, error) {
	if e.loadedTypes[1] {
		return e.OutcomeMeasurementList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_measurement_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomecategory.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomecategory.FieldOutcomeCategoryTitle:
			values[i] = new(sql.NullString)
		case outcomecategory.ForeignKeys[0]: // outcome_class_outcome_category_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeCategory fields.
func (oc *OutcomeCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomecategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oc.ID = int(value.Int64)
		case outcomecategory.FieldOutcomeCategoryTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_category_title", values[i])
			} else if value.Valid {
				oc.OutcomeCategoryTitle = value.String
			}
		case outcomecategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_class_outcome_category_list", value)
			} else if value.Valid {
				oc.outcome_class_outcome_category_list = new(int)
				*oc.outcome_class_outcome_category_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeCategory entity.
func (oc *OutcomeCategory) QueryParent() *OutcomeClassQuery {
	return (&OutcomeCategoryClient{config: oc.config}).QueryParent(oc)
}

// QueryOutcomeMeasurementList queries the "outcome_measurement_list" edge of the OutcomeCategory entity.
func (oc *OutcomeCategory) QueryOutcomeMeasurementList() *OutcomeMeasurementQuery {
	return (&OutcomeCategoryClient{config: oc.config}).QueryOutcomeMeasurementList(oc)
}

// Update returns a builder for updating this OutcomeCategory.
// Note that you need to call OutcomeCategory.Unwrap() before calling this method if this OutcomeCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (oc *OutcomeCategory) Update() *OutcomeCategoryUpdateOne {
	return (&OutcomeCategoryClient{config: oc.config}).UpdateOne(oc)
}

// Unwrap unwraps the OutcomeCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oc *OutcomeCategory) Unwrap() *OutcomeCategory {
	tx, ok := oc.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeCategory is not a transactional entity")
	}
	oc.config.driver = tx.drv
	return oc
}

// String implements the fmt.Stringer.
func (oc *OutcomeCategory) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", oc.ID))
	builder.WriteString(", outcome_category_title=")
	builder.WriteString(oc.OutcomeCategoryTitle)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeCategories is a parsable slice of OutcomeCategory.
type OutcomeCategories []*OutcomeCategory

func (oc OutcomeCategories) config(cfg config) {
	for _i := range oc {
		oc[_i].config = cfg
	}
}
