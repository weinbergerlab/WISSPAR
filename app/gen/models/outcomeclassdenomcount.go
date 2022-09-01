// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
)

// OutcomeClassDenomCount is the model entity for the OutcomeClassDenomCount schema.
type OutcomeClassDenomCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeClassDenomCountGroupID holds the value of the "outcome_class_denom_count_group_id" field.
	OutcomeClassDenomCountGroupID string `json:"outcome_class_denom_count_group_id,omitempty"`
	// OutcomeClassDenomCountValue holds the value of the "outcome_class_denom_count_value" field.
	OutcomeClassDenomCountValue string `json:"outcome_class_denom_count_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeClassDenomCountQuery when eager-loading is set.
	Edges                                              OutcomeClassDenomCountEdges `json:"edges"`
	outcome_class_denom_outcome_class_denom_count_list *int
}

// OutcomeClassDenomCountEdges holds the relations/edges for other nodes in the graph.
type OutcomeClassDenomCountEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeClassDenom `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeClassDenomCountEdges) ParentOrErr() (*OutcomeClassDenom, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomeclassdenom.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeClassDenomCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeclassdenomcount.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID, outcomeclassdenomcount.FieldOutcomeClassDenomCountValue:
			values[i] = new(sql.NullString)
		case outcomeclassdenomcount.ForeignKeys[0]: // outcome_class_denom_outcome_class_denom_count_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeClassDenomCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeClassDenomCount fields.
func (ocdc *OutcomeClassDenomCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeclassdenomcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ocdc.ID = int(value.Int64)
		case outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_class_denom_count_group_id", values[i])
			} else if value.Valid {
				ocdc.OutcomeClassDenomCountGroupID = value.String
			}
		case outcomeclassdenomcount.FieldOutcomeClassDenomCountValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_class_denom_count_value", values[i])
			} else if value.Valid {
				ocdc.OutcomeClassDenomCountValue = value.String
			}
		case outcomeclassdenomcount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_class_denom_outcome_class_denom_count_list", value)
			} else if value.Valid {
				ocdc.outcome_class_denom_outcome_class_denom_count_list = new(int)
				*ocdc.outcome_class_denom_outcome_class_denom_count_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeClassDenomCount entity.
func (ocdc *OutcomeClassDenomCount) QueryParent() *OutcomeClassDenomQuery {
	return (&OutcomeClassDenomCountClient{config: ocdc.config}).QueryParent(ocdc)
}

// Update returns a builder for updating this OutcomeClassDenomCount.
// Note that you need to call OutcomeClassDenomCount.Unwrap() before calling this method if this OutcomeClassDenomCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ocdc *OutcomeClassDenomCount) Update() *OutcomeClassDenomCountUpdateOne {
	return (&OutcomeClassDenomCountClient{config: ocdc.config}).UpdateOne(ocdc)
}

// Unwrap unwraps the OutcomeClassDenomCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ocdc *OutcomeClassDenomCount) Unwrap() *OutcomeClassDenomCount {
	tx, ok := ocdc.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeClassDenomCount is not a transactional entity")
	}
	ocdc.config.driver = tx.drv
	return ocdc
}

// String implements the fmt.Stringer.
func (ocdc *OutcomeClassDenomCount) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeClassDenomCount(")
	builder.WriteString(fmt.Sprintf("id=%v", ocdc.ID))
	builder.WriteString(", outcome_class_denom_count_group_id=")
	builder.WriteString(ocdc.OutcomeClassDenomCountGroupID)
	builder.WriteString(", outcome_class_denom_count_value=")
	builder.WriteString(ocdc.OutcomeClassDenomCountValue)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeClassDenomCounts is a parsable slice of OutcomeClassDenomCount.
type OutcomeClassDenomCounts []*OutcomeClassDenomCount

func (ocdc OutcomeClassDenomCounts) config(cfg config) {
	for _i := range ocdc {
		ocdc[_i].config = cfg
	}
}
