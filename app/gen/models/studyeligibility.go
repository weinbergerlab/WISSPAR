// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
)

// StudyEligibility is the model entity for the StudyEligibility schema.
type StudyEligibility struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EligibilityCriteria holds the value of the "EligibilityCriteria" field.
	EligibilityCriteria string `json:"EligibilityCriteria,omitempty"`
	// HealthyVolunteers holds the value of the "HealthyVolunteers" field.
	HealthyVolunteers string `json:"HealthyVolunteers,omitempty"`
	// Gender holds the value of the "Gender" field.
	Gender string `json:"Gender,omitempty"`
	// MinimumAge holds the value of the "MinimumAge" field.
	MinimumAge string `json:"MinimumAge,omitempty"`
	// MaximumAge holds the value of the "MaximumAge" field.
	MaximumAge string `json:"MaximumAge,omitempty"`
	// StdAgeList holds the value of the "StdAgeList" field.
	StdAgeList string `json:"StdAgeList,omitempty"`
	// Ethnicity holds the value of the "Ethnicity" field.
	Ethnicity string `json:"Ethnicity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudyEligibilityQuery when eager-loading is set.
	Edges                            StudyEligibilityEdges `json:"edges"`
	clinical_trial_study_eligibility *int
}

// StudyEligibilityEdges holds the relations/edges for other nodes in the graph.
type StudyEligibilityEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ClinicalTrial `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudyEligibilityEdges) ParentOrErr() (*ClinicalTrial, error) {
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
func (*StudyEligibility) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studyeligibility.FieldID:
			values[i] = new(sql.NullInt64)
		case studyeligibility.FieldEligibilityCriteria, studyeligibility.FieldHealthyVolunteers, studyeligibility.FieldGender, studyeligibility.FieldMinimumAge, studyeligibility.FieldMaximumAge, studyeligibility.FieldStdAgeList, studyeligibility.FieldEthnicity:
			values[i] = new(sql.NullString)
		case studyeligibility.ForeignKeys[0]: // clinical_trial_study_eligibility
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type StudyEligibility", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudyEligibility fields.
func (se *StudyEligibility) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studyeligibility.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			se.ID = int(value.Int64)
		case studyeligibility.FieldEligibilityCriteria:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligibilityCriteria", values[i])
			} else if value.Valid {
				se.EligibilityCriteria = value.String
			}
		case studyeligibility.FieldHealthyVolunteers:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field HealthyVolunteers", values[i])
			} else if value.Valid {
				se.HealthyVolunteers = value.String
			}
		case studyeligibility.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Gender", values[i])
			} else if value.Valid {
				se.Gender = value.String
			}
		case studyeligibility.FieldMinimumAge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MinimumAge", values[i])
			} else if value.Valid {
				se.MinimumAge = value.String
			}
		case studyeligibility.FieldMaximumAge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MaximumAge", values[i])
			} else if value.Valid {
				se.MaximumAge = value.String
			}
		case studyeligibility.FieldStdAgeList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field StdAgeList", values[i])
			} else if value.Valid {
				se.StdAgeList = value.String
			}
		case studyeligibility.FieldEthnicity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Ethnicity", values[i])
			} else if value.Valid {
				se.Ethnicity = value.String
			}
		case studyeligibility.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field clinical_trial_study_eligibility", value)
			} else if value.Valid {
				se.clinical_trial_study_eligibility = new(int)
				*se.clinical_trial_study_eligibility = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the StudyEligibility entity.
func (se *StudyEligibility) QueryParent() *ClinicalTrialQuery {
	return (&StudyEligibilityClient{config: se.config}).QueryParent(se)
}

// Update returns a builder for updating this StudyEligibility.
// Note that you need to call StudyEligibility.Unwrap() before calling this method if this StudyEligibility
// was returned from a transaction, and the transaction was committed or rolled back.
func (se *StudyEligibility) Update() *StudyEligibilityUpdateOne {
	return (&StudyEligibilityClient{config: se.config}).UpdateOne(se)
}

// Unwrap unwraps the StudyEligibility entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (se *StudyEligibility) Unwrap() *StudyEligibility {
	tx, ok := se.config.driver.(*txDriver)
	if !ok {
		panic("models: StudyEligibility is not a transactional entity")
	}
	se.config.driver = tx.drv
	return se
}

// String implements the fmt.Stringer.
func (se *StudyEligibility) String() string {
	var builder strings.Builder
	builder.WriteString("StudyEligibility(")
	builder.WriteString(fmt.Sprintf("id=%v", se.ID))
	builder.WriteString(", EligibilityCriteria=")
	builder.WriteString(se.EligibilityCriteria)
	builder.WriteString(", HealthyVolunteers=")
	builder.WriteString(se.HealthyVolunteers)
	builder.WriteString(", Gender=")
	builder.WriteString(se.Gender)
	builder.WriteString(", MinimumAge=")
	builder.WriteString(se.MinimumAge)
	builder.WriteString(", MaximumAge=")
	builder.WriteString(se.MaximumAge)
	builder.WriteString(", StdAgeList=")
	builder.WriteString(se.StdAgeList)
	builder.WriteString(", Ethnicity=")
	builder.WriteString(se.Ethnicity)
	builder.WriteByte(')')
	return builder.String()
}

// StudyEligibilities is a parsable slice of StudyEligibility.
type StudyEligibilities []*StudyEligibility

func (se StudyEligibilities) config(cfg config) {
	for _i := range se {
		se[_i].config = cfg
	}
}
