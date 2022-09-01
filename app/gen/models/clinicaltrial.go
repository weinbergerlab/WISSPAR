// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ClinicalTrial is the model entity for the ClinicalTrial schema.
type ClinicalTrial struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StudyName holds the value of the "study_name" field.
	StudyName string `json:"study_name,omitempty"`
	// Sponsor holds the value of the "sponsor" field.
	Sponsor string `json:"sponsor,omitempty"`
	// ResponsibleParty holds the value of the "responsible_party" field.
	ResponsibleParty string `json:"responsible_party,omitempty"`
	// StudyType holds the value of the "study_type" field.
	StudyType string `json:"study_type,omitempty"`
	// Phase holds the value of the "phase" field.
	Phase string `json:"phase,omitempty"`
	// ActualEnrollment holds the value of the "actual_enrollment" field.
	ActualEnrollment string `json:"actual_enrollment,omitempty"`
	// Allocation holds the value of the "allocation" field.
	Allocation string `json:"allocation,omitempty"`
	// InterventionModel holds the value of the "intervention_model" field.
	InterventionModel string `json:"intervention_model,omitempty"`
	// Masking holds the value of the "masking" field.
	Masking string `json:"masking,omitempty"`
	// PrimaryPurpose holds the value of the "primary_purpose" field.
	PrimaryPurpose string `json:"primary_purpose,omitempty"`
	// OfficialTitle holds the value of the "official_title" field.
	OfficialTitle string `json:"official_title,omitempty"`
	// ActualStudyStartDate holds the value of the "actual_study_start_date" field.
	ActualStudyStartDate time.Time `json:"actual_study_start_date,omitempty"`
	// ActualPrimaryCompletionDate holds the value of the "actual_primary_completion_date" field.
	ActualPrimaryCompletionDate time.Time `json:"actual_primary_completion_date,omitempty"`
	// ActualStudyCompletionDate holds the value of the "actual_study_completion_date" field.
	ActualStudyCompletionDate time.Time `json:"actual_study_completion_date,omitempty"`
	// StudyID holds the value of the "study_id" field.
	StudyID string `json:"study_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClinicalTrialQuery when eager-loading is set.
	Edges ClinicalTrialEdges `json:"edges"`
}

// ClinicalTrialEdges holds the relations/edges for other nodes in the graph.
type ClinicalTrialEdges struct {
	// ResultsDefinition holds the value of the results_definition edge.
	ResultsDefinition *ResultsDefinition `json:"results_definition,omitempty"`
	// StudyLocations holds the value of the study_locations edge.
	StudyLocations []*StudyLocation `json:"study_locations,omitempty"`
	// StudyEligibility holds the value of the study_eligibility edge.
	StudyEligibility []*StudyEligibility `json:"study_eligibility,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata []*ClinicalTrialMetadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ResultsDefinitionOrErr returns the ResultsDefinition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClinicalTrialEdges) ResultsDefinitionOrErr() (*ResultsDefinition, error) {
	if e.loadedTypes[0] {
		if e.ResultsDefinition == nil {
			// The edge results_definition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resultsdefinition.Label}
		}
		return e.ResultsDefinition, nil
	}
	return nil, &NotLoadedError{edge: "results_definition"}
}

// StudyLocationsOrErr returns the StudyLocations value or an error if the edge
// was not loaded in eager-loading.
func (e ClinicalTrialEdges) StudyLocationsOrErr() ([]*StudyLocation, error) {
	if e.loadedTypes[1] {
		return e.StudyLocations, nil
	}
	return nil, &NotLoadedError{edge: "study_locations"}
}

// StudyEligibilityOrErr returns the StudyEligibility value or an error if the edge
// was not loaded in eager-loading.
func (e ClinicalTrialEdges) StudyEligibilityOrErr() ([]*StudyEligibility, error) {
	if e.loadedTypes[2] {
		return e.StudyEligibility, nil
	}
	return nil, &NotLoadedError{edge: "study_eligibility"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading.
func (e ClinicalTrialEdges) MetadataOrErr() ([]*ClinicalTrialMetadata, error) {
	if e.loadedTypes[3] {
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClinicalTrial) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case clinicaltrial.FieldID:
			values[i] = new(sql.NullInt64)
		case clinicaltrial.FieldStudyName, clinicaltrial.FieldSponsor, clinicaltrial.FieldResponsibleParty, clinicaltrial.FieldStudyType, clinicaltrial.FieldPhase, clinicaltrial.FieldActualEnrollment, clinicaltrial.FieldAllocation, clinicaltrial.FieldInterventionModel, clinicaltrial.FieldMasking, clinicaltrial.FieldPrimaryPurpose, clinicaltrial.FieldOfficialTitle, clinicaltrial.FieldStudyID:
			values[i] = new(sql.NullString)
		case clinicaltrial.FieldActualStudyStartDate, clinicaltrial.FieldActualPrimaryCompletionDate, clinicaltrial.FieldActualStudyCompletionDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ClinicalTrial", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClinicalTrial fields.
func (ct *ClinicalTrial) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clinicaltrial.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case clinicaltrial.FieldStudyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field study_name", values[i])
			} else if value.Valid {
				ct.StudyName = value.String
			}
		case clinicaltrial.FieldSponsor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sponsor", values[i])
			} else if value.Valid {
				ct.Sponsor = value.String
			}
		case clinicaltrial.FieldResponsibleParty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field responsible_party", values[i])
			} else if value.Valid {
				ct.ResponsibleParty = value.String
			}
		case clinicaltrial.FieldStudyType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field study_type", values[i])
			} else if value.Valid {
				ct.StudyType = value.String
			}
		case clinicaltrial.FieldPhase:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phase", values[i])
			} else if value.Valid {
				ct.Phase = value.String
			}
		case clinicaltrial.FieldActualEnrollment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actual_enrollment", values[i])
			} else if value.Valid {
				ct.ActualEnrollment = value.String
			}
		case clinicaltrial.FieldAllocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field allocation", values[i])
			} else if value.Valid {
				ct.Allocation = value.String
			}
		case clinicaltrial.FieldInterventionModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intervention_model", values[i])
			} else if value.Valid {
				ct.InterventionModel = value.String
			}
		case clinicaltrial.FieldMasking:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field masking", values[i])
			} else if value.Valid {
				ct.Masking = value.String
			}
		case clinicaltrial.FieldPrimaryPurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_purpose", values[i])
			} else if value.Valid {
				ct.PrimaryPurpose = value.String
			}
		case clinicaltrial.FieldOfficialTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field official_title", values[i])
			} else if value.Valid {
				ct.OfficialTitle = value.String
			}
		case clinicaltrial.FieldActualStudyStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field actual_study_start_date", values[i])
			} else if value.Valid {
				ct.ActualStudyStartDate = value.Time
			}
		case clinicaltrial.FieldActualPrimaryCompletionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field actual_primary_completion_date", values[i])
			} else if value.Valid {
				ct.ActualPrimaryCompletionDate = value.Time
			}
		case clinicaltrial.FieldActualStudyCompletionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field actual_study_completion_date", values[i])
			} else if value.Valid {
				ct.ActualStudyCompletionDate = value.Time
			}
		case clinicaltrial.FieldStudyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field study_id", values[i])
			} else if value.Valid {
				ct.StudyID = value.String
			}
		}
	}
	return nil
}

// QueryResultsDefinition queries the "results_definition" edge of the ClinicalTrial entity.
func (ct *ClinicalTrial) QueryResultsDefinition() *ResultsDefinitionQuery {
	return (&ClinicalTrialClient{config: ct.config}).QueryResultsDefinition(ct)
}

// QueryStudyLocations queries the "study_locations" edge of the ClinicalTrial entity.
func (ct *ClinicalTrial) QueryStudyLocations() *StudyLocationQuery {
	return (&ClinicalTrialClient{config: ct.config}).QueryStudyLocations(ct)
}

// QueryStudyEligibility queries the "study_eligibility" edge of the ClinicalTrial entity.
func (ct *ClinicalTrial) QueryStudyEligibility() *StudyEligibilityQuery {
	return (&ClinicalTrialClient{config: ct.config}).QueryStudyEligibility(ct)
}

// QueryMetadata queries the "metadata" edge of the ClinicalTrial entity.
func (ct *ClinicalTrial) QueryMetadata() *ClinicalTrialMetadataQuery {
	return (&ClinicalTrialClient{config: ct.config}).QueryMetadata(ct)
}

// Update returns a builder for updating this ClinicalTrial.
// Note that you need to call ClinicalTrial.Unwrap() before calling this method if this ClinicalTrial
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *ClinicalTrial) Update() *ClinicalTrialUpdateOne {
	return (&ClinicalTrialClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the ClinicalTrial entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *ClinicalTrial) Unwrap() *ClinicalTrial {
	tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("models: ClinicalTrial is not a transactional entity")
	}
	ct.config.driver = tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *ClinicalTrial) String() string {
	var builder strings.Builder
	builder.WriteString("ClinicalTrial(")
	builder.WriteString(fmt.Sprintf("id=%v", ct.ID))
	builder.WriteString(", study_name=")
	builder.WriteString(ct.StudyName)
	builder.WriteString(", sponsor=")
	builder.WriteString(ct.Sponsor)
	builder.WriteString(", responsible_party=")
	builder.WriteString(ct.ResponsibleParty)
	builder.WriteString(", study_type=")
	builder.WriteString(ct.StudyType)
	builder.WriteString(", phase=")
	builder.WriteString(ct.Phase)
	builder.WriteString(", actual_enrollment=")
	builder.WriteString(ct.ActualEnrollment)
	builder.WriteString(", allocation=")
	builder.WriteString(ct.Allocation)
	builder.WriteString(", intervention_model=")
	builder.WriteString(ct.InterventionModel)
	builder.WriteString(", masking=")
	builder.WriteString(ct.Masking)
	builder.WriteString(", primary_purpose=")
	builder.WriteString(ct.PrimaryPurpose)
	builder.WriteString(", official_title=")
	builder.WriteString(ct.OfficialTitle)
	builder.WriteString(", actual_study_start_date=")
	builder.WriteString(ct.ActualStudyStartDate.Format(time.ANSIC))
	builder.WriteString(", actual_primary_completion_date=")
	builder.WriteString(ct.ActualPrimaryCompletionDate.Format(time.ANSIC))
	builder.WriteString(", actual_study_completion_date=")
	builder.WriteString(ct.ActualStudyCompletionDate.Format(time.ANSIC))
	builder.WriteString(", study_id=")
	builder.WriteString(ct.StudyID)
	builder.WriteByte(')')
	return builder.String()
}

// ClinicalTrials is a parsable slice of ClinicalTrial.
type ClinicalTrials []*ClinicalTrial

func (ct ClinicalTrials) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
