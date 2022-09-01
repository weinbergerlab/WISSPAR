// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// MoreInfoModule is the model entity for the MoreInfoModule schema.
type MoreInfoModule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LimitationsAndCaveatsDescription holds the value of the "limitations_and_caveats_description" field.
	LimitationsAndCaveatsDescription string `json:"limitations_and_caveats_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MoreInfoModuleQuery when eager-loading is set.
	Edges                               MoreInfoModuleEdges `json:"edges"`
	more_info_module_certain_agreement  *int
	more_info_module_point_of_contact   *int
	results_definition_more_info_module *int
}

// MoreInfoModuleEdges holds the relations/edges for other nodes in the graph.
type MoreInfoModuleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ResultsDefinition `json:"parent,omitempty"`
	// CertainAgreement holds the value of the certain_agreement edge.
	CertainAgreement *CertainAgreement `json:"certain_agreement,omitempty"`
	// PointOfContact holds the value of the point_of_contact edge.
	PointOfContact *PointOfContact `json:"point_of_contact,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MoreInfoModuleEdges) ParentOrErr() (*ResultsDefinition, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resultsdefinition.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// CertainAgreementOrErr returns the CertainAgreement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MoreInfoModuleEdges) CertainAgreementOrErr() (*CertainAgreement, error) {
	if e.loadedTypes[1] {
		if e.CertainAgreement == nil {
			// The edge certain_agreement was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: certainagreement.Label}
		}
		return e.CertainAgreement, nil
	}
	return nil, &NotLoadedError{edge: "certain_agreement"}
}

// PointOfContactOrErr returns the PointOfContact value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MoreInfoModuleEdges) PointOfContactOrErr() (*PointOfContact, error) {
	if e.loadedTypes[2] {
		if e.PointOfContact == nil {
			// The edge point_of_contact was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pointofcontact.Label}
		}
		return e.PointOfContact, nil
	}
	return nil, &NotLoadedError{edge: "point_of_contact"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MoreInfoModule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case moreinfomodule.FieldID:
			values[i] = new(sql.NullInt64)
		case moreinfomodule.FieldLimitationsAndCaveatsDescription:
			values[i] = new(sql.NullString)
		case moreinfomodule.ForeignKeys[0]: // more_info_module_certain_agreement
			values[i] = new(sql.NullInt64)
		case moreinfomodule.ForeignKeys[1]: // more_info_module_point_of_contact
			values[i] = new(sql.NullInt64)
		case moreinfomodule.ForeignKeys[2]: // results_definition_more_info_module
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MoreInfoModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MoreInfoModule fields.
func (mim *MoreInfoModule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case moreinfomodule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mim.ID = int(value.Int64)
		case moreinfomodule.FieldLimitationsAndCaveatsDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field limitations_and_caveats_description", values[i])
			} else if value.Valid {
				mim.LimitationsAndCaveatsDescription = value.String
			}
		case moreinfomodule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field more_info_module_certain_agreement", value)
			} else if value.Valid {
				mim.more_info_module_certain_agreement = new(int)
				*mim.more_info_module_certain_agreement = int(value.Int64)
			}
		case moreinfomodule.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field more_info_module_point_of_contact", value)
			} else if value.Valid {
				mim.more_info_module_point_of_contact = new(int)
				*mim.more_info_module_point_of_contact = int(value.Int64)
			}
		case moreinfomodule.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field results_definition_more_info_module", value)
			} else if value.Valid {
				mim.results_definition_more_info_module = new(int)
				*mim.results_definition_more_info_module = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the MoreInfoModule entity.
func (mim *MoreInfoModule) QueryParent() *ResultsDefinitionQuery {
	return (&MoreInfoModuleClient{config: mim.config}).QueryParent(mim)
}

// QueryCertainAgreement queries the "certain_agreement" edge of the MoreInfoModule entity.
func (mim *MoreInfoModule) QueryCertainAgreement() *CertainAgreementQuery {
	return (&MoreInfoModuleClient{config: mim.config}).QueryCertainAgreement(mim)
}

// QueryPointOfContact queries the "point_of_contact" edge of the MoreInfoModule entity.
func (mim *MoreInfoModule) QueryPointOfContact() *PointOfContactQuery {
	return (&MoreInfoModuleClient{config: mim.config}).QueryPointOfContact(mim)
}

// Update returns a builder for updating this MoreInfoModule.
// Note that you need to call MoreInfoModule.Unwrap() before calling this method if this MoreInfoModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (mim *MoreInfoModule) Update() *MoreInfoModuleUpdateOne {
	return (&MoreInfoModuleClient{config: mim.config}).UpdateOne(mim)
}

// Unwrap unwraps the MoreInfoModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mim *MoreInfoModule) Unwrap() *MoreInfoModule {
	tx, ok := mim.config.driver.(*txDriver)
	if !ok {
		panic("models: MoreInfoModule is not a transactional entity")
	}
	mim.config.driver = tx.drv
	return mim
}

// String implements the fmt.Stringer.
func (mim *MoreInfoModule) String() string {
	var builder strings.Builder
	builder.WriteString("MoreInfoModule(")
	builder.WriteString(fmt.Sprintf("id=%v", mim.ID))
	builder.WriteString(", limitations_and_caveats_description=")
	builder.WriteString(mim.LimitationsAndCaveatsDescription)
	builder.WriteByte(')')
	return builder.String()
}

// MoreInfoModules is a parsable slice of MoreInfoModule.
type MoreInfoModules []*MoreInfoModule

func (mim MoreInfoModules) config(cfg config) {
	for _i := range mim {
		mim[_i].config = cfg
	}
}
