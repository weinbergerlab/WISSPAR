// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
)

// UseCase is the model entity for the UseCase schema.
type UseCase struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UseCaseName holds the value of the "use_case_name" field.
	UseCaseName string `json:"use_case_name,omitempty"`
	// UseCaseDescription holds the value of the "use_case_description" field.
	UseCaseDescription string `json:"use_case_description,omitempty"`
	// UseCaseCode holds the value of the "use_case_code" field.
	UseCaseCode string `json:"use_case_code,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UseCase) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case usecase.FieldID:
			values[i] = new(sql.NullInt64)
		case usecase.FieldUseCaseName, usecase.FieldUseCaseDescription, usecase.FieldUseCaseCode:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UseCase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UseCase fields.
func (uc *UseCase) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usecase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int(value.Int64)
		case usecase.FieldUseCaseName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field use_case_name", values[i])
			} else if value.Valid {
				uc.UseCaseName = value.String
			}
		case usecase.FieldUseCaseDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field use_case_description", values[i])
			} else if value.Valid {
				uc.UseCaseDescription = value.String
			}
		case usecase.FieldUseCaseCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field use_case_code", values[i])
			} else if value.Valid {
				uc.UseCaseCode = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UseCase.
// Note that you need to call UseCase.Unwrap() before calling this method if this UseCase
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UseCase) Update() *UseCaseUpdateOne {
	return (&UseCaseClient{config: uc.config}).UpdateOne(uc)
}

// Unwrap unwraps the UseCase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UseCase) Unwrap() *UseCase {
	tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("models: UseCase is not a transactional entity")
	}
	uc.config.driver = tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UseCase) String() string {
	var builder strings.Builder
	builder.WriteString("UseCase(")
	builder.WriteString(fmt.Sprintf("id=%v", uc.ID))
	builder.WriteString(", use_case_name=")
	builder.WriteString(uc.UseCaseName)
	builder.WriteString(", use_case_description=")
	builder.WriteString(uc.UseCaseDescription)
	builder.WriteString(", use_case_code=")
	builder.WriteString(uc.UseCaseCode)
	builder.WriteByte(')')
	return builder.String()
}

// UseCases is a parsable slice of UseCase.
type UseCases []*UseCase

func (uc UseCases) config(cfg config) {
	for _i := range uc {
		uc[_i].config = cfg
	}
}
