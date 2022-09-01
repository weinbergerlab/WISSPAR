// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
)

// ImmunocompromisedGroups is the model entity for the ImmunocompromisedGroups schema.
type ImmunocompromisedGroups struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GroupName holds the value of the "group_name" field.
	GroupName string `json:"group_name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImmunocompromisedGroups) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case immunocompromisedgroups.FieldID:
			values[i] = new(sql.NullInt64)
		case immunocompromisedgroups.FieldGroupName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ImmunocompromisedGroups", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImmunocompromisedGroups fields.
func (ig *ImmunocompromisedGroups) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case immunocompromisedgroups.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ig.ID = int(value.Int64)
		case immunocompromisedgroups.FieldGroupName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_name", values[i])
			} else if value.Valid {
				ig.GroupName = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ImmunocompromisedGroups.
// Note that you need to call ImmunocompromisedGroups.Unwrap() before calling this method if this ImmunocompromisedGroups
// was returned from a transaction, and the transaction was committed or rolled back.
func (ig *ImmunocompromisedGroups) Update() *ImmunocompromisedGroupsUpdateOne {
	return (&ImmunocompromisedGroupsClient{config: ig.config}).UpdateOne(ig)
}

// Unwrap unwraps the ImmunocompromisedGroups entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ig *ImmunocompromisedGroups) Unwrap() *ImmunocompromisedGroups {
	tx, ok := ig.config.driver.(*txDriver)
	if !ok {
		panic("models: ImmunocompromisedGroups is not a transactional entity")
	}
	ig.config.driver = tx.drv
	return ig
}

// String implements the fmt.Stringer.
func (ig *ImmunocompromisedGroups) String() string {
	var builder strings.Builder
	builder.WriteString("ImmunocompromisedGroups(")
	builder.WriteString(fmt.Sprintf("id=%v", ig.ID))
	builder.WriteString(", group_name=")
	builder.WriteString(ig.GroupName)
	builder.WriteByte(')')
	return builder.String()
}

// ImmunocompromisedGroupsSlice is a parsable slice of ImmunocompromisedGroups.
type ImmunocompromisedGroupsSlice []*ImmunocompromisedGroups

func (ig ImmunocompromisedGroupsSlice) config(cfg config) {
	for _i := range ig {
		ig[_i].config = cfg
	}
}
