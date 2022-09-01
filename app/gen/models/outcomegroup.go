// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeGroup is the model entity for the OutcomeGroup schema.
type OutcomeGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeGroupID holds the value of the "outcome_group_id" field.
	OutcomeGroupID string `json:"outcome_group_id,omitempty"`
	// OutcomeGroupTitle holds the value of the "outcome_group_title" field.
	OutcomeGroupTitle string `json:"outcome_group_title,omitempty"`
	// OutcomeGroupDescription holds the value of the "outcome_group_description" field.
	OutcomeGroupDescription string `json:"outcome_group_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeGroupQuery when eager-loading is set.
	Edges                              OutcomeGroupEdges `json:"edges"`
	outcome_measure_outcome_group_list *int
}

// OutcomeGroupEdges holds the relations/edges for other nodes in the graph.
type OutcomeGroupEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeMeasure `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeGroupEdges) ParentOrErr() (*OutcomeMeasure, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomegroup.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomegroup.FieldOutcomeGroupID, outcomegroup.FieldOutcomeGroupTitle, outcomegroup.FieldOutcomeGroupDescription:
			values[i] = new(sql.NullString)
		case outcomegroup.ForeignKeys[0]: // outcome_measure_outcome_group_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeGroup fields.
func (og *OutcomeGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomegroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			og.ID = int(value.Int64)
		case outcomegroup.FieldOutcomeGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_group_id", values[i])
			} else if value.Valid {
				og.OutcomeGroupID = value.String
			}
		case outcomegroup.FieldOutcomeGroupTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_group_title", values[i])
			} else if value.Valid {
				og.OutcomeGroupTitle = value.String
			}
		case outcomegroup.FieldOutcomeGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_group_description", values[i])
			} else if value.Valid {
				og.OutcomeGroupDescription = value.String
			}
		case outcomegroup.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_measure_outcome_group_list", value)
			} else if value.Valid {
				og.outcome_measure_outcome_group_list = new(int)
				*og.outcome_measure_outcome_group_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeGroup entity.
func (og *OutcomeGroup) QueryParent() *OutcomeMeasureQuery {
	return (&OutcomeGroupClient{config: og.config}).QueryParent(og)
}

// Update returns a builder for updating this OutcomeGroup.
// Note that you need to call OutcomeGroup.Unwrap() before calling this method if this OutcomeGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (og *OutcomeGroup) Update() *OutcomeGroupUpdateOne {
	return (&OutcomeGroupClient{config: og.config}).UpdateOne(og)
}

// Unwrap unwraps the OutcomeGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (og *OutcomeGroup) Unwrap() *OutcomeGroup {
	tx, ok := og.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeGroup is not a transactional entity")
	}
	og.config.driver = tx.drv
	return og
}

// String implements the fmt.Stringer.
func (og *OutcomeGroup) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", og.ID))
	builder.WriteString(", outcome_group_id=")
	builder.WriteString(og.OutcomeGroupID)
	builder.WriteString(", outcome_group_title=")
	builder.WriteString(og.OutcomeGroupTitle)
	builder.WriteString(", outcome_group_description=")
	builder.WriteString(og.OutcomeGroupDescription)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeGroups is a parsable slice of OutcomeGroup.
type OutcomeGroups []*OutcomeGroup

func (og OutcomeGroups) config(cfg config) {
	for _i := range og {
		og[_i].config = cfg
	}
}
