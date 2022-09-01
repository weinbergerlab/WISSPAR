// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
)

// DoseDescription is the model entity for the DoseDescription schema.
type DoseDescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DoseDescription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dosedescription.FieldID:
			values[i] = new(sql.NullInt64)
		case dosedescription.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DoseDescription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DoseDescription fields.
func (dd *DoseDescription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dosedescription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dd.ID = int(value.Int64)
		case dosedescription.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dd.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DoseDescription.
// Note that you need to call DoseDescription.Unwrap() before calling this method if this DoseDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (dd *DoseDescription) Update() *DoseDescriptionUpdateOne {
	return (&DoseDescriptionClient{config: dd.config}).UpdateOne(dd)
}

// Unwrap unwraps the DoseDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dd *DoseDescription) Unwrap() *DoseDescription {
	tx, ok := dd.config.driver.(*txDriver)
	if !ok {
		panic("models: DoseDescription is not a transactional entity")
	}
	dd.config.driver = tx.drv
	return dd
}

// String implements the fmt.Stringer.
func (dd *DoseDescription) String() string {
	var builder strings.Builder
	builder.WriteString("DoseDescription(")
	builder.WriteString(fmt.Sprintf("id=%v", dd.ID))
	builder.WriteString(", name=")
	builder.WriteString(dd.Name)
	builder.WriteByte(')')
	return builder.String()
}

// DoseDescriptions is a parsable slice of DoseDescription.
type DoseDescriptions []*DoseDescription

func (dd DoseDescriptions) config(cfg config) {
	for _i := range dd {
		dd[_i].config = cfg
	}
}
