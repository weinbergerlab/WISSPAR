// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// StudyLocation is the model entity for the StudyLocation schema.
type StudyLocation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LocationFacility holds the value of the "LocationFacility" field.
	LocationFacility string `json:"LocationFacility,omitempty"`
	// LocationCity holds the value of the "LocationCity" field.
	LocationCity string `json:"LocationCity,omitempty"`
	// LocationCountry holds the value of the "LocationCountry" field.
	LocationCountry string `json:"LocationCountry,omitempty"`
	// LocationCountryCode holds the value of the "LocationCountryCode" field.
	LocationCountryCode string `json:"LocationCountryCode,omitempty"`
	// LocationContinentName holds the value of the "LocationContinentName" field.
	LocationContinentName string `json:"LocationContinentName,omitempty"`
	// LocationContinentCode holds the value of the "LocationContinentCode" field.
	LocationContinentCode string `json:"LocationContinentCode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudyLocationQuery when eager-loading is set.
	Edges                          StudyLocationEdges `json:"edges"`
	clinical_trial_study_locations *int
}

// StudyLocationEdges holds the relations/edges for other nodes in the graph.
type StudyLocationEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ClinicalTrial `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudyLocationEdges) ParentOrErr() (*ClinicalTrial, error) {
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
func (*StudyLocation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studylocation.FieldID:
			values[i] = new(sql.NullInt64)
		case studylocation.FieldLocationFacility, studylocation.FieldLocationCity, studylocation.FieldLocationCountry, studylocation.FieldLocationCountryCode, studylocation.FieldLocationContinentName, studylocation.FieldLocationContinentCode:
			values[i] = new(sql.NullString)
		case studylocation.ForeignKeys[0]: // clinical_trial_study_locations
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type StudyLocation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudyLocation fields.
func (sl *StudyLocation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studylocation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sl.ID = int(value.Int64)
		case studylocation.FieldLocationFacility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationFacility", values[i])
			} else if value.Valid {
				sl.LocationFacility = value.String
			}
		case studylocation.FieldLocationCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationCity", values[i])
			} else if value.Valid {
				sl.LocationCity = value.String
			}
		case studylocation.FieldLocationCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationCountry", values[i])
			} else if value.Valid {
				sl.LocationCountry = value.String
			}
		case studylocation.FieldLocationCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationCountryCode", values[i])
			} else if value.Valid {
				sl.LocationCountryCode = value.String
			}
		case studylocation.FieldLocationContinentName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationContinentName", values[i])
			} else if value.Valid {
				sl.LocationContinentName = value.String
			}
		case studylocation.FieldLocationContinentCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocationContinentCode", values[i])
			} else if value.Valid {
				sl.LocationContinentCode = value.String
			}
		case studylocation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field clinical_trial_study_locations", value)
			} else if value.Valid {
				sl.clinical_trial_study_locations = new(int)
				*sl.clinical_trial_study_locations = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the StudyLocation entity.
func (sl *StudyLocation) QueryParent() *ClinicalTrialQuery {
	return (&StudyLocationClient{config: sl.config}).QueryParent(sl)
}

// Update returns a builder for updating this StudyLocation.
// Note that you need to call StudyLocation.Unwrap() before calling this method if this StudyLocation
// was returned from a transaction, and the transaction was committed or rolled back.
func (sl *StudyLocation) Update() *StudyLocationUpdateOne {
	return (&StudyLocationClient{config: sl.config}).UpdateOne(sl)
}

// Unwrap unwraps the StudyLocation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sl *StudyLocation) Unwrap() *StudyLocation {
	tx, ok := sl.config.driver.(*txDriver)
	if !ok {
		panic("models: StudyLocation is not a transactional entity")
	}
	sl.config.driver = tx.drv
	return sl
}

// String implements the fmt.Stringer.
func (sl *StudyLocation) String() string {
	var builder strings.Builder
	builder.WriteString("StudyLocation(")
	builder.WriteString(fmt.Sprintf("id=%v", sl.ID))
	builder.WriteString(", LocationFacility=")
	builder.WriteString(sl.LocationFacility)
	builder.WriteString(", LocationCity=")
	builder.WriteString(sl.LocationCity)
	builder.WriteString(", LocationCountry=")
	builder.WriteString(sl.LocationCountry)
	builder.WriteString(", LocationCountryCode=")
	builder.WriteString(sl.LocationCountryCode)
	builder.WriteString(", LocationContinentName=")
	builder.WriteString(sl.LocationContinentName)
	builder.WriteString(", LocationContinentCode=")
	builder.WriteString(sl.LocationContinentCode)
	builder.WriteByte(')')
	return builder.String()
}

// StudyLocations is a parsable slice of StudyLocation.
type StudyLocations []*StudyLocation

func (sl StudyLocations) config(cfg config) {
	for _i := range sl {
		sl[_i].config = cfg
	}
}
