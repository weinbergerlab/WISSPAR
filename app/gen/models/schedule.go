// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/schedule"
)

// Schedule is the model entity for the Schedule schema.
type Schedule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Schedule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case schedule.FieldID:
			values[i] = new(sql.NullInt64)
		case schedule.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Schedule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Schedule fields.
func (s *Schedule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case schedule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case schedule.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Schedule.
// Note that you need to call Schedule.Unwrap() before calling this method if this Schedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Schedule) Update() *ScheduleUpdateOne {
	return (&ScheduleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Schedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Schedule) Unwrap() *Schedule {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("models: Schedule is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Schedule) String() string {
	var builder strings.Builder
	builder.WriteString("Schedule(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Schedules is a parsable slice of Schedule.
type Schedules []*Schedule

func (s Schedules) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
