// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
)

// OutcomeMeasurement is the model entity for the OutcomeMeasurement schema.
type OutcomeMeasurement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeMeasurementGroupID holds the value of the "outcome_measurement_group_id" field.
	OutcomeMeasurementGroupID string `json:"outcome_measurement_group_id,omitempty"`
	// OutcomeMeasurementValue holds the value of the "outcome_measurement_value" field.
	OutcomeMeasurementValue string `json:"outcome_measurement_value,omitempty"`
	// OutcomeMeasurementSpread holds the value of the "outcome_measurement_spread" field.
	OutcomeMeasurementSpread string `json:"outcome_measurement_spread,omitempty"`
	// OutcomeMeasurementLowerLimit holds the value of the "outcome_measurement_lower_limit" field.
	OutcomeMeasurementLowerLimit string `json:"outcome_measurement_lower_limit,omitempty"`
	// OutcomeMeasurementUpperLimit holds the value of the "outcome_measurement_upper_limit" field.
	OutcomeMeasurementUpperLimit string `json:"outcome_measurement_upper_limit,omitempty"`
	// OutcomeMeasurementComment holds the value of the "outcome_measurement_comment" field.
	OutcomeMeasurementComment string `json:"outcome_measurement_comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeMeasurementQuery when eager-loading is set.
	Edges                                     OutcomeMeasurementEdges `json:"edges"`
	outcome_category_outcome_measurement_list *int
}

// OutcomeMeasurementEdges holds the relations/edges for other nodes in the graph.
type OutcomeMeasurementEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeCategory `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeMeasurementEdges) ParentOrErr() (*OutcomeCategory, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomecategory.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeMeasurement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomemeasurement.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomemeasurement.FieldOutcomeMeasurementGroupID, outcomemeasurement.FieldOutcomeMeasurementValue, outcomemeasurement.FieldOutcomeMeasurementSpread, outcomemeasurement.FieldOutcomeMeasurementLowerLimit, outcomemeasurement.FieldOutcomeMeasurementUpperLimit, outcomemeasurement.FieldOutcomeMeasurementComment:
			values[i] = new(sql.NullString)
		case outcomemeasurement.ForeignKeys[0]: // outcome_category_outcome_measurement_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeMeasurement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeMeasurement fields.
func (om *OutcomeMeasurement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomemeasurement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			om.ID = int(value.Int64)
		case outcomemeasurement.FieldOutcomeMeasurementGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_group_id", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementGroupID = value.String
			}
		case outcomemeasurement.FieldOutcomeMeasurementValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_value", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementValue = value.String
			}
		case outcomemeasurement.FieldOutcomeMeasurementSpread:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_spread", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementSpread = value.String
			}
		case outcomemeasurement.FieldOutcomeMeasurementLowerLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_lower_limit", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementLowerLimit = value.String
			}
		case outcomemeasurement.FieldOutcomeMeasurementUpperLimit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_upper_limit", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementUpperLimit = value.String
			}
		case outcomemeasurement.FieldOutcomeMeasurementComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_measurement_comment", values[i])
			} else if value.Valid {
				om.OutcomeMeasurementComment = value.String
			}
		case outcomemeasurement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_category_outcome_measurement_list", value)
			} else if value.Valid {
				om.outcome_category_outcome_measurement_list = new(int)
				*om.outcome_category_outcome_measurement_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeMeasurement entity.
func (om *OutcomeMeasurement) QueryParent() *OutcomeCategoryQuery {
	return (&OutcomeMeasurementClient{config: om.config}).QueryParent(om)
}

// Update returns a builder for updating this OutcomeMeasurement.
// Note that you need to call OutcomeMeasurement.Unwrap() before calling this method if this OutcomeMeasurement
// was returned from a transaction, and the transaction was committed or rolled back.
func (om *OutcomeMeasurement) Update() *OutcomeMeasurementUpdateOne {
	return (&OutcomeMeasurementClient{config: om.config}).UpdateOne(om)
}

// Unwrap unwraps the OutcomeMeasurement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (om *OutcomeMeasurement) Unwrap() *OutcomeMeasurement {
	tx, ok := om.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeMeasurement is not a transactional entity")
	}
	om.config.driver = tx.drv
	return om
}

// String implements the fmt.Stringer.
func (om *OutcomeMeasurement) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeMeasurement(")
	builder.WriteString(fmt.Sprintf("id=%v", om.ID))
	builder.WriteString(", outcome_measurement_group_id=")
	builder.WriteString(om.OutcomeMeasurementGroupID)
	builder.WriteString(", outcome_measurement_value=")
	builder.WriteString(om.OutcomeMeasurementValue)
	builder.WriteString(", outcome_measurement_spread=")
	builder.WriteString(om.OutcomeMeasurementSpread)
	builder.WriteString(", outcome_measurement_lower_limit=")
	builder.WriteString(om.OutcomeMeasurementLowerLimit)
	builder.WriteString(", outcome_measurement_upper_limit=")
	builder.WriteString(om.OutcomeMeasurementUpperLimit)
	builder.WriteString(", outcome_measurement_comment=")
	builder.WriteString(om.OutcomeMeasurementComment)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeMeasurements is a parsable slice of OutcomeMeasurement.
type OutcomeMeasurements []*OutcomeMeasurement

func (om OutcomeMeasurements) config(cfg config) {
	for _i := range om {
		om[_i].config = cfg
	}
}
