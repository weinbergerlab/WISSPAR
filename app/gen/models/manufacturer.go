// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/manufacturer"
)

// Manufacturer is the model entity for the Manufacturer schema.
type Manufacturer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ManufacturerName holds the value of the "manufacturer_name" field.
	ManufacturerName string `json:"manufacturer_name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Manufacturer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case manufacturer.FieldID:
			values[i] = new(sql.NullInt64)
		case manufacturer.FieldManufacturerName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Manufacturer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Manufacturer fields.
func (m *Manufacturer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case manufacturer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case manufacturer.FieldManufacturerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer_name", values[i])
			} else if value.Valid {
				m.ManufacturerName = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Manufacturer.
// Note that you need to call Manufacturer.Unwrap() before calling this method if this Manufacturer
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Manufacturer) Update() *ManufacturerUpdateOne {
	return (&ManufacturerClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Manufacturer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Manufacturer) Unwrap() *Manufacturer {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("models: Manufacturer is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Manufacturer) String() string {
	var builder strings.Builder
	builder.WriteString("Manufacturer(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", manufacturer_name=")
	builder.WriteString(m.ManufacturerName)
	builder.WriteByte(')')
	return builder.String()
}

// Manufacturers is a parsable slice of Manufacturer.
type Manufacturers []*Manufacturer

func (m Manufacturers) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
