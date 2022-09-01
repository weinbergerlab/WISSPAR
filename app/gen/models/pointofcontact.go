// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
)

// PointOfContact is the model entity for the PointOfContact schema.
type PointOfContact struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PointOfContactTitle holds the value of the "point_of_contact_title" field.
	PointOfContactTitle string `json:"point_of_contact_title,omitempty"`
	// PointOfContactOrganization holds the value of the "point_of_contact_organization" field.
	PointOfContactOrganization string `json:"point_of_contact_organization,omitempty"`
	// PointOfContactEmail holds the value of the "point_of_contact_email" field.
	PointOfContactEmail string `json:"point_of_contact_email,omitempty"`
	// PointOfContactPhone holds the value of the "point_of_contact_phone" field.
	PointOfContactPhone string `json:"point_of_contact_phone,omitempty"`
	// PointOfContactPhoneExt holds the value of the "point_of_contact_phone_ext" field.
	PointOfContactPhoneExt string `json:"point_of_contact_phone_ext,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PointOfContactQuery when eager-loading is set.
	Edges                   PointOfContactEdges `json:"edges"`
	point_of_contact_parent *int
}

// PointOfContactEdges holds the relations/edges for other nodes in the graph.
type PointOfContactEdges struct {
	// Parent holds the value of the parent edge.
	Parent *MoreInfoModule `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PointOfContactEdges) ParentOrErr() (*MoreInfoModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moreinfomodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PointOfContact) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pointofcontact.FieldID:
			values[i] = new(sql.NullInt64)
		case pointofcontact.FieldPointOfContactTitle, pointofcontact.FieldPointOfContactOrganization, pointofcontact.FieldPointOfContactEmail, pointofcontact.FieldPointOfContactPhone, pointofcontact.FieldPointOfContactPhoneExt:
			values[i] = new(sql.NullString)
		case pointofcontact.ForeignKeys[0]: // point_of_contact_parent
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PointOfContact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PointOfContact fields.
func (poc *PointOfContact) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pointofcontact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			poc.ID = int(value.Int64)
		case pointofcontact.FieldPointOfContactTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_of_contact_title", values[i])
			} else if value.Valid {
				poc.PointOfContactTitle = value.String
			}
		case pointofcontact.FieldPointOfContactOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_of_contact_organization", values[i])
			} else if value.Valid {
				poc.PointOfContactOrganization = value.String
			}
		case pointofcontact.FieldPointOfContactEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_of_contact_email", values[i])
			} else if value.Valid {
				poc.PointOfContactEmail = value.String
			}
		case pointofcontact.FieldPointOfContactPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_of_contact_phone", values[i])
			} else if value.Valid {
				poc.PointOfContactPhone = value.String
			}
		case pointofcontact.FieldPointOfContactPhoneExt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_of_contact_phone_ext", values[i])
			} else if value.Valid {
				poc.PointOfContactPhoneExt = value.String
			}
		case pointofcontact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field point_of_contact_parent", value)
			} else if value.Valid {
				poc.point_of_contact_parent = new(int)
				*poc.point_of_contact_parent = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the PointOfContact entity.
func (poc *PointOfContact) QueryParent() *MoreInfoModuleQuery {
	return (&PointOfContactClient{config: poc.config}).QueryParent(poc)
}

// Update returns a builder for updating this PointOfContact.
// Note that you need to call PointOfContact.Unwrap() before calling this method if this PointOfContact
// was returned from a transaction, and the transaction was committed or rolled back.
func (poc *PointOfContact) Update() *PointOfContactUpdateOne {
	return (&PointOfContactClient{config: poc.config}).UpdateOne(poc)
}

// Unwrap unwraps the PointOfContact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (poc *PointOfContact) Unwrap() *PointOfContact {
	tx, ok := poc.config.driver.(*txDriver)
	if !ok {
		panic("models: PointOfContact is not a transactional entity")
	}
	poc.config.driver = tx.drv
	return poc
}

// String implements the fmt.Stringer.
func (poc *PointOfContact) String() string {
	var builder strings.Builder
	builder.WriteString("PointOfContact(")
	builder.WriteString(fmt.Sprintf("id=%v", poc.ID))
	builder.WriteString(", point_of_contact_title=")
	builder.WriteString(poc.PointOfContactTitle)
	builder.WriteString(", point_of_contact_organization=")
	builder.WriteString(poc.PointOfContactOrganization)
	builder.WriteString(", point_of_contact_email=")
	builder.WriteString(poc.PointOfContactEmail)
	builder.WriteString(", point_of_contact_phone=")
	builder.WriteString(poc.PointOfContactPhone)
	builder.WriteString(", point_of_contact_phone_ext=")
	builder.WriteString(poc.PointOfContactPhoneExt)
	builder.WriteByte(')')
	return builder.String()
}

// PointOfContacts is a parsable slice of PointOfContact.
type PointOfContacts []*PointOfContact

func (poc PointOfContacts) config(cfg config) {
	for _i := range poc {
		poc[_i].config = cfg
	}
}
