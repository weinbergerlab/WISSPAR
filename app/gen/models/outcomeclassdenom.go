// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
)

// OutcomeClassDenom is the model entity for the OutcomeClassDenom schema.
type OutcomeClassDenom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeClassDenomUnits holds the value of the "outcome_class_denom_units" field.
	OutcomeClassDenomUnits string `json:"outcome_class_denom_units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeClassDenomQuery when eager-loading is set.
	Edges                                  OutcomeClassDenomEdges `json:"edges"`
	outcome_class_outcome_class_denom_list *int
}

// OutcomeClassDenomEdges holds the relations/edges for other nodes in the graph.
type OutcomeClassDenomEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeClass `json:"parent,omitempty"`
	// OutcomeClassDenomCountList holds the value of the outcome_class_denom_count_list edge.
	OutcomeClassDenomCountList []*OutcomeClassDenomCount `json:"outcome_class_denom_count_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeClassDenomEdges) ParentOrErr() (*OutcomeClass, error) {
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

// OutcomeClassDenomCountListOrErr returns the OutcomeClassDenomCountList value or an error if the edge
// was not loaded in eager-loading.
func (e OutcomeClassDenomEdges) OutcomeClassDenomCountListOrErr() ([]*OutcomeClassDenomCount, error) {
	if e.loadedTypes[1] {
		return e.OutcomeClassDenomCountList, nil
	}
	return nil, &NotLoadedError{edge: "outcome_class_denom_count_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeClassDenom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeclassdenom.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomeclassdenom.FieldOutcomeClassDenomUnits:
			values[i] = new(sql.NullString)
		case outcomeclassdenom.ForeignKeys[0]: // outcome_class_outcome_class_denom_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeClassDenom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeClassDenom fields.
func (ocd *OutcomeClassDenom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeclassdenom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ocd.ID = int(value.Int64)
		case outcomeclassdenom.FieldOutcomeClassDenomUnits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_class_denom_units", values[i])
			} else if value.Valid {
				ocd.OutcomeClassDenomUnits = value.String
			}
		case outcomeclassdenom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_class_outcome_class_denom_list", value)
			} else if value.Valid {
				ocd.outcome_class_outcome_class_denom_list = new(int)
				*ocd.outcome_class_outcome_class_denom_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeClassDenom entity.
func (ocd *OutcomeClassDenom) QueryParent() *OutcomeClassQuery {
	return (&OutcomeClassDenomClient{config: ocd.config}).QueryParent(ocd)
}

// QueryOutcomeClassDenomCountList queries the "outcome_class_denom_count_list" edge of the OutcomeClassDenom entity.
func (ocd *OutcomeClassDenom) QueryOutcomeClassDenomCountList() *OutcomeClassDenomCountQuery {
	return (&OutcomeClassDenomClient{config: ocd.config}).QueryOutcomeClassDenomCountList(ocd)
}

// Update returns a builder for updating this OutcomeClassDenom.
// Note that you need to call OutcomeClassDenom.Unwrap() before calling this method if this OutcomeClassDenom
// was returned from a transaction, and the transaction was committed or rolled back.
func (ocd *OutcomeClassDenom) Update() *OutcomeClassDenomUpdateOne {
	return (&OutcomeClassDenomClient{config: ocd.config}).UpdateOne(ocd)
}

// Unwrap unwraps the OutcomeClassDenom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ocd *OutcomeClassDenom) Unwrap() *OutcomeClassDenom {
	tx, ok := ocd.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeClassDenom is not a transactional entity")
	}
	ocd.config.driver = tx.drv
	return ocd
}

// String implements the fmt.Stringer.
func (ocd *OutcomeClassDenom) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeClassDenom(")
	builder.WriteString(fmt.Sprintf("id=%v", ocd.ID))
	builder.WriteString(", outcome_class_denom_units=")
	builder.WriteString(ocd.OutcomeClassDenomUnits)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeClassDenoms is a parsable slice of OutcomeClassDenom.
type OutcomeClassDenoms []*OutcomeClassDenom

func (ocd OutcomeClassDenoms) config(cfg config) {
	for _i := range ocd {
		ocd[_i].config = cfg
	}
}
