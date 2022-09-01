// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
)

// Vaccine is the model entity for the Vaccine schema.
type Vaccine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VaccineName holds the value of the "vaccine_name" field.
	VaccineName string `json:"vaccine_name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vaccine) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vaccine.FieldID:
			values[i] = new(sql.NullInt64)
		case vaccine.FieldVaccineName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vaccine", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vaccine fields.
func (v *Vaccine) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vaccine.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vaccine.FieldVaccineName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vaccine_name", values[i])
			} else if value.Valid {
				v.VaccineName = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Vaccine.
// Note that you need to call Vaccine.Unwrap() before calling this method if this Vaccine
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vaccine) Update() *VaccineUpdateOne {
	return (&VaccineClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Vaccine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vaccine) Unwrap() *Vaccine {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("models: Vaccine is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vaccine) String() string {
	var builder strings.Builder
	builder.WriteString("Vaccine(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", vaccine_name=")
	builder.WriteString(v.VaccineName)
	builder.WriteByte(')')
	return builder.String()
}

// Vaccines is a parsable slice of Vaccine.
type Vaccines []*Vaccine

func (v Vaccines) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
