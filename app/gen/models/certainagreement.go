// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
)

// CertainAgreement is the model entity for the CertainAgreement schema.
type CertainAgreement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AgreementPiSponsorEmployee holds the value of the "agreement_pi_sponsor_employee" field.
	AgreementPiSponsorEmployee string `json:"agreement_pi_sponsor_employee,omitempty"`
	// AgreementRestrictionType holds the value of the "agreement_restriction_type" field.
	AgreementRestrictionType string `json:"agreement_restriction_type,omitempty"`
	// AgreementRestrictiveAgreement holds the value of the "agreement_restrictive_agreement" field.
	AgreementRestrictiveAgreement string `json:"agreement_restrictive_agreement,omitempty"`
	// AgreementOtherDetails holds the value of the "agreement_other_details" field.
	AgreementOtherDetails string `json:"agreement_other_details,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CertainAgreementQuery when eager-loading is set.
	Edges                    CertainAgreementEdges `json:"edges"`
	certain_agreement_parent *int
}

// CertainAgreementEdges holds the relations/edges for other nodes in the graph.
type CertainAgreementEdges struct {
	// Parent holds the value of the parent edge.
	Parent *MoreInfoModule `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertainAgreementEdges) ParentOrErr() (*MoreInfoModule, error) {
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
func (*CertainAgreement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case certainagreement.FieldID:
			values[i] = new(sql.NullInt64)
		case certainagreement.FieldAgreementPiSponsorEmployee, certainagreement.FieldAgreementRestrictionType, certainagreement.FieldAgreementRestrictiveAgreement, certainagreement.FieldAgreementOtherDetails:
			values[i] = new(sql.NullString)
		case certainagreement.ForeignKeys[0]: // certain_agreement_parent
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CertainAgreement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertainAgreement fields.
func (ca *CertainAgreement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certainagreement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ca.ID = int(value.Int64)
		case certainagreement.FieldAgreementPiSponsorEmployee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agreement_pi_sponsor_employee", values[i])
			} else if value.Valid {
				ca.AgreementPiSponsorEmployee = value.String
			}
		case certainagreement.FieldAgreementRestrictionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agreement_restriction_type", values[i])
			} else if value.Valid {
				ca.AgreementRestrictionType = value.String
			}
		case certainagreement.FieldAgreementRestrictiveAgreement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agreement_restrictive_agreement", values[i])
			} else if value.Valid {
				ca.AgreementRestrictiveAgreement = value.String
			}
		case certainagreement.FieldAgreementOtherDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agreement_other_details", values[i])
			} else if value.Valid {
				ca.AgreementOtherDetails = value.String
			}
		case certainagreement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field certain_agreement_parent", value)
			} else if value.Valid {
				ca.certain_agreement_parent = new(int)
				*ca.certain_agreement_parent = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the CertainAgreement entity.
func (ca *CertainAgreement) QueryParent() *MoreInfoModuleQuery {
	return (&CertainAgreementClient{config: ca.config}).QueryParent(ca)
}

// Update returns a builder for updating this CertainAgreement.
// Note that you need to call CertainAgreement.Unwrap() before calling this method if this CertainAgreement
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CertainAgreement) Update() *CertainAgreementUpdateOne {
	return (&CertainAgreementClient{config: ca.config}).UpdateOne(ca)
}

// Unwrap unwraps the CertainAgreement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CertainAgreement) Unwrap() *CertainAgreement {
	tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("models: CertainAgreement is not a transactional entity")
	}
	ca.config.driver = tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CertainAgreement) String() string {
	var builder strings.Builder
	builder.WriteString("CertainAgreement(")
	builder.WriteString(fmt.Sprintf("id=%v", ca.ID))
	builder.WriteString(", agreement_pi_sponsor_employee=")
	builder.WriteString(ca.AgreementPiSponsorEmployee)
	builder.WriteString(", agreement_restriction_type=")
	builder.WriteString(ca.AgreementRestrictionType)
	builder.WriteString(", agreement_restrictive_agreement=")
	builder.WriteString(ca.AgreementRestrictiveAgreement)
	builder.WriteString(", agreement_other_details=")
	builder.WriteString(ca.AgreementOtherDetails)
	builder.WriteByte(')')
	return builder.String()
}

// CertainAgreements is a parsable slice of CertainAgreement.
type CertainAgreements []*CertainAgreement

func (ca CertainAgreements) config(cfg config) {
	for _i := range ca {
		ca[_i].config = cfg
	}
}
