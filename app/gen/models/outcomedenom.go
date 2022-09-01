// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeDenom is the model entity for the OutcomeDenom schema.
type OutcomeDenom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeDenomUnits holds the value of the "outcome_denom_units" field.
	OutcomeDenomUnits string `json:"outcome_denom_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeDenomQuery when eager-loading is set.
	Edges                              OutcomeDenomEdges `json:"edges"`
	outcome_measure_outcome_denom_list *int
}

// OutcomeDenomEdges holds the relations/edges for other nodes in the graph.
type OutcomeDenomEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasure `json:"parent,omitempty"`
	// OutcomeDenomCountList holds the value of the outcome_denom_count_list edge.
	OutcomeDenomCountList []*OutcomeDenomCount `json:"outcome_denom_count_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeDenomEdges) ParentOrErr() (*OutcomeMeasure, error) {
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

// OutcomeDenomCountListOrErr returns the OutcomeDenomCountList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeDenomEdges) OutcomeDenomCountListOrErr() ([]*OutcomeDenomCount, error) {
	if e.loadedTypes[1] {
		return e.OutcomeDenomCountList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_denom_count_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeDenom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomedenom.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomedenom.FieldOutcomeDenomUnits:
			values[i] = new(sql.NullString)
		case outcomedenom.ForeignKeys[0]: // outcome_measure_outcome_denom_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeDenom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeDenom fields.
func (od *OutcomeDenom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomedenom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			od.ID = int(value.Int64)
		case outcomedenom.FieldOutcomeDenomUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_denom_units", values[i])
			} else if value.Valid {
				od.OutcomeDenomUnits = value.String
			}
		case outcomedenom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measure_outcome_denom_list", value)
			} else if value.Valid {
				od.outcome_measure_outcome_denom_list = new(int)
				*od.outcome_measure_outcome_denom_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeDenom entity.
func (od *OutcomeDenom) QueryParent() *OutcomeMeasureQuery {
	return (&OutcomeDenomClient{config: od.config}).QueryParent(od)
}

// QueryOutcomeDenomCountList queries the "outcome_denom_count_list" edge of the OutcomeDenom entity.
func (od *OutcomeDenom) QueryOutcomeDenomCountList() *OutcomeDenomCountQuery {
	return (&OutcomeDenomClient{config: od.config}).QueryOutcomeDenomCountList(od)
}

// Update returns a builder for updating this OutcomeDenom.
// Note that you need to call OutcomeDenom.Unwrap() before calling this method if this OutcomeDenom
// was returned from a transaction, and the transaction was committed or rolled back.
func (od *OutcomeDenom) Update() *OutcomeDenomUpdateOne {
	return (&OutcomeDenomClient{config: od.config}).UpdateOne(od)
}

// Unwrap unwraps the OutcomeDenom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (od *OutcomeDenom) Unwrap() *OutcomeDenom {
	tx, ok := od.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeDenom is not a transactional entity")
	}
	od.config.driver = tx.drv
	return od
}

// String implements the fmt.Stringer.
func (od *OutcomeDenom) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeDenom(")
	builder.WriteString(fmt.Sprintf("id=%v", od.ID))
	builder.WriteString(", outcome_denom_units=")
	builder.WriteString(od.OutcomeDenomUnits)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeDenoms is a parsable slice of OutcomeDenom.
type OutcomeDenoms []*OutcomeDenom

func (od OutcomeDenoms) config(cfg config) {
	for _i := range od {
		od[_i].config = cfg
	}
}
