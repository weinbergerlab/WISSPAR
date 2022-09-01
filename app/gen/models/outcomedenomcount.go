// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
)

// OutcomeDenomCount is the model entity for the OutcomeDenomCount schema.
type OutcomeDenomCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeDenomCountGroupID holds the value of the "outcome_denom_count_group_id" field.
	OutcomeDenomCountGroupID string `json:"outcome_denom_count_group_id,omitempty"`
	// OutcomeDenomCountValue holds the value of the "outcome_denom_count_value" field.
	OutcomeDenomCountValue string `json:"outcome_denom_count_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeDenomCountQuery when eager-loading is set.
	Edges                                  OutcomeDenomCountEdges `json:"edges"`
	outcome_denom_outcome_denom_count_list *int
}

// OutcomeDenomCountEdges holds the relations/edges for other nodes in the graph.
type OutcomeDenomCountEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeDenom `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeDenomCountEdges) ParentOrErr() (*OutcomeDenom, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomedenom.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeDenomCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomedenomcount.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomedenomcount.FieldOutcomeDenomCountGroupID, outcomedenomcount.FieldOutcomeDenomCountValue:
			values[i] = new(sql.NullString)
		case outcomedenomcount.ForeignKeys[0]: // outcome_denom_outcome_denom_count_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeDenomCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeDenomCount fields.
func (odc *OutcomeDenomCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomedenomcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			odc.ID = int(value.Int64)
		case outcomedenomcount.FieldOutcomeDenomCountGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_denom_count_group_id", values[i])
			} else if value.Valid {
				odc.OutcomeDenomCountGroupID = value.String
			}
		case outcomedenomcount.FieldOutcomeDenomCountValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_denom_count_value", values[i])
			} else if value.Valid {
				odc.OutcomeDenomCountValue = value.String
			}
		case outcomedenomcount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_denom_outcome_denom_count_list", value)
			} else if value.Valid {
				odc.outcome_denom_outcome_denom_count_list = new(int)
				*odc.outcome_denom_outcome_denom_count_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeDenomCount entity.
func (odc *OutcomeDenomCount) QueryParent() *OutcomeDenomQuery {
	return (&OutcomeDenomCountClient{config: odc.config}).QueryParent(odc)
}

// Update returns a builder for updating this OutcomeDenomCount.
// Note that you need to call OutcomeDenomCount.Unwrap() before calling this method if this OutcomeDenomCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (odc *OutcomeDenomCount) Update() *OutcomeDenomCountUpdateOne {
	return (&OutcomeDenomCountClient{config: odc.config}).UpdateOne(odc)
}

// Unwrap unwraps the OutcomeDenomCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (odc *OutcomeDenomCount) Unwrap() *OutcomeDenomCount {
	tx, ok := odc.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeDenomCount is not a transactional entity")
	}
	odc.config.driver = tx.drv
	return odc
}

// String implements the fmt.Stringer.
func (odc *OutcomeDenomCount) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeDenomCount(")
	builder.WriteString(fmt.Sprintf("id=%v", odc.ID))
	builder.WriteString(", outcome_denom_count_group_id=")
	builder.WriteString(odc.OutcomeDenomCountGroupID)
	builder.WriteString(", outcome_denom_count_value=")
	builder.WriteString(odc.OutcomeDenomCountValue)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeDenomCounts is a parsable slice of OutcomeDenomCount.
type OutcomeDenomCounts []*OutcomeDenomCount

func (odc OutcomeDenomCounts) config(cfg config) {
	for _i := range odc {
		odc[_i].config = cfg
	}
}
