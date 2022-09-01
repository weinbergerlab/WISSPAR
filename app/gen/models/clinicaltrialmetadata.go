// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
)

// ClinicalTrialMetadata is the model entity for the ClinicalTrialMetadata schema.
type ClinicalTrialMetadata struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TagName holds the value of the "tag_name" field.
	TagName string `json:"tag_name,omitempty"`
	// TagValue holds the value of the "tag_value" field.
	TagValue string `json:"tag_value,omitempty"`
	// UseCaseCode holds the value of the "use_case_code" field.
	UseCaseCode string `json:"use_case_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClinicalTrialMetadataQuery when eager-loading is set.
	Edges                   ClinicalTrialMetadataEdges `json:"edges"`
	clinical_trial_metadata *int
}

// ClinicalTrialMetadataEdges holds the relations/edges for other nodes in the graph.
type ClinicalTrialMetadataEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ClinicalTrial `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClinicalTrialMetadataEdges) ParentOrErr() (*ClinicalTrial, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: clinicaltrial.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClinicalTrialMetadata) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case clinicaltrialmetadata.FieldID:
			values[i] = new(sql.NullInt64)
		case clinicaltrialmetadata.FieldTagName, clinicaltrialmetadata.FieldTagValue, clinicaltrialmetadata.FieldUseCaseCode:
			values[i] = new(sql.NullString)
		case clinicaltrialmetadata.ForeignKeys[0]: // clinical_trial_metadata
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ClinicalTrialMetadata", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClinicalTrialMetadata fields.
func (ctm *ClinicalTrialMetadata) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clinicaltrialmetadata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ctm.ID = int(value.Int64)
		case clinicaltrialmetadata.FieldTagName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag_name", values[i])
			} else if value.Valid {
				ctm.TagName = value.String
			}
		case clinicaltrialmetadata.FieldTagValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag_value", values[i])
			} else if value.Valid {
				ctm.TagValue = value.String
			}
		case clinicaltrialmetadata.FieldUseCaseCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field use_case_code", values[i])
			} else if value.Valid {
				ctm.UseCaseCode = value.String
			}
		case clinicaltrialmetadata.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field clinical_trial_metadata", value)
			} else if value.Valid {
				ctm.clinical_trial_metadata = new(int)
				*ctm.clinical_trial_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ClinicalTrialMetadata entity.
func (ctm *ClinicalTrialMetadata) QueryParent() *ClinicalTrialQuery {
	return (&ClinicalTrialMetadataClient{config: ctm.config}).QueryParent(ctm)
}

// Update returns a builder for updating this ClinicalTrialMetadata.
// Note that you need to call ClinicalTrialMetadata.Unwrap() before calling this method if this ClinicalTrialMetadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (ctm *ClinicalTrialMetadata) Update() *ClinicalTrialMetadataUpdateOne {
	return (&ClinicalTrialMetadataClient{config: ctm.config}).UpdateOne(ctm)
}

// Unwrap unwraps the ClinicalTrialMetadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ctm *ClinicalTrialMetadata) Unwrap() *ClinicalTrialMetadata {
	tx, ok := ctm.config.driver.(*txDriver)
	if !ok {
		panic("models: ClinicalTrialMetadata is not a transactional entity")
	}
	ctm.config.driver = tx.drv
	return ctm
}

// String implements the fmt.Stringer.
func (ctm *ClinicalTrialMetadata) String() string {
	var builder strings.Builder
	builder.WriteString("ClinicalTrialMetadata(")
	builder.WriteString(fmt.Sprintf("id=%v", ctm.ID))
	builder.WriteString(", tag_name=")
	builder.WriteString(ctm.TagName)
	builder.WriteString(", tag_value=")
	builder.WriteString(ctm.TagValue)
	builder.WriteString(", use_case_code=")
	builder.WriteString(ctm.UseCaseCode)
	builder.WriteByte(')')
	return builder.String()
}

// ClinicalTrialMetadataSlice is a parsable slice of ClinicalTrialMetadata.
type ClinicalTrialMetadataSlice []*ClinicalTrialMetadata

func (ctm ClinicalTrialMetadataSlice) config(cfg config) {
	for _i := range ctm {
		ctm[_i].config = cfg
	}
}
