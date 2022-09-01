// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeClass is the model entity for the OutcomeClass schema.
type OutcomeClass struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeClassTitle holds the value of the "outcome_class_title" field.
	OutcomeClassTitle string `json:"outcome_class_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeClassQuery when eager-loading is set.
	Edges                              OutcomeClassEdges `json:"edges"`
	outcome_measure_outcome_class_list *int
}

// OutcomeClassEdges holds the relations/edges for other nodes in the graph.
type OutcomeClassEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasure `json:"parent,omitempty"`
	// OutcomeClassDenomList holds the value of the outcome_class_denom_list edge.
	OutcomeClassDenomList []*OutcomeClassDenom `json:"outcome_class_denom_list,omitempty"`
	// OutcomeCategoryList holds the value of the outcome_category_list edge.
	OutcomeCategoryList []*OutcomeCategory `json:"outcome_category_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeClassEdges) ParentOrErr() (*OutcomeMeasure, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomemeasure.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// OutcomeClassDenomListOrErr returns the OutcomeClassDenomList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeClassEdges) OutcomeClassDenomListOrErr() ([]*OutcomeClassDenom, error) {
	if e.loadedTypes[1] {
		return e.OutcomeClassDenomList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_class_denom_list"}
}

// OutcomeCategoryListOrErr returns the OutcomeCategoryList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeClassEdges) OutcomeCategoryListOrErr() ([]*OutcomeCategory, error) {
	if e.loadedTypes[2] {
		return e.OutcomeCategoryList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_category_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeClass) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeclass.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomeclass.FieldOutcomeClassTitle:
			values[i] = new(sql.NullString)
		case outcomeclass.ForeignKeys[0]: // outcome_measure_outcome_class_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeClass", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeClass fields.
func (oc *OutcomeClass) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeclass.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oc.ID = int(value.Int64)
		case outcomeclass.FieldOutcomeClassTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_class_title", values[i])
			} else if value.Valid {
				oc.OutcomeClassTitle = value.String
			}
		case outcomeclass.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measure_outcome_class_list", value)
			} else if value.Valid {
				oc.outcome_measure_outcome_class_list = new(int)
				*oc.outcome_measure_outcome_class_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeClass entity.
func (oc *OutcomeClass) QueryParent() *OutcomeMeasureQuery {
	return (&OutcomeClassClient{config: oc.config}).QueryParent(oc)
}

// QueryOutcomeClassDenomList queries the "outcome_class_denom_list" edge of the OutcomeClass entity.
func (oc *OutcomeClass) QueryOutcomeClassDenomList() *OutcomeClassDenomQuery {
	return (&OutcomeClassClient{config: oc.config}).QueryOutcomeClassDenomList(oc)
}

// QueryOutcomeCategoryList queries the "outcome_category_list" edge of the OutcomeClass entity.
func (oc *OutcomeClass) QueryOutcomeCategoryList() *OutcomeCategoryQuery {
	return (&OutcomeClassClient{config: oc.config}).QueryOutcomeCategoryList(oc)
}

// Update returns a builder for updating this OutcomeClass.
// Note that you need to call OutcomeClass.Unwrap() before calling this method if this OutcomeClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (oc *OutcomeClass) Update() *OutcomeClassUpdateOne {
	return (&OutcomeClassClient{config: oc.config}).UpdateOne(oc)
}

// Unwrap unwraps the OutcomeClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oc *OutcomeClass) Unwrap() *OutcomeClass {
	tx, ok := oc.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeClass is not a transactional entity")
	}
	oc.config.driver = tx.drv
	return oc
}

// String implements the fmt.Stringer.
func (oc *OutcomeClass) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeClass(")
	builder.WriteString(fmt.Sprintf("id=%v", oc.ID))
	builder.WriteString(", outcome_class_title=")
	builder.WriteString(oc.OutcomeClassTitle)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeClasses is a parsable slice of OutcomeClass.
type OutcomeClasses []*OutcomeClass

func (oc OutcomeClasses) config(cfg config) {
	for _i := range oc {
		oc[_i].config = cfg
	}
}
