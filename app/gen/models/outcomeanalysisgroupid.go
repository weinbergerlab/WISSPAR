// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
)

// OutcomeAnalysisGroupID is the model entity for the OutcomeAnalysisGroupID schema.
type OutcomeAnalysisGroupID struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OutcomeAnalysisGroupID holds the value of the "outcome_analysis_group_id" field.
	OutcomeAnalysisGroupID string `json:"outcome_analysis_group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutcomeAnalysisGroupIDQuery when eager-loading is set.
	Edges                                           OutcomeAnalysisGroupIDEdges `json:"edges"`
	outcome_analysis_outcome_analysis_group_id_list *int
}

// OutcomeAnalysisGroupIDEdges holds the relations/edges for other nodes in the graph.
type OutcomeAnalysisGroupIDEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OutcomeAnalysis `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutcomeAnalysisGroupIDEdges) ParentOrErr() (*OutcomeAnalysis, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outcomeanalysis.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutcomeAnalysisGroupID) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case outcomeanalysisgroupid.FieldID:
			values[i] = new(sql.NullInt64)
		case outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID:
			values[i] = new(sql.NullString)
		case outcomeanalysisgroupid.ForeignKeys[0]: // outcome_analysis_outcome_analysis_group_id_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutcomeAnalysisGroupID", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutcomeAnalysisGroupID fields.
func (oagi *OutcomeAnalysisGroupID) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outcomeanalysisgroupid.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oagi.ID = int(value.Int64)
		case outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outcome_analysis_group_id", values[i])
			} else if value.Valid {
				oagi.OutcomeAnalysisGroupID = value.String
			}
		case outcomeanalysisgroupid.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field outcome_analysis_outcome_analysis_group_id_list", value)
			} else if value.Valid {
				oagi.outcome_analysis_outcome_analysis_group_id_list = new(int)
				*oagi.outcome_analysis_outcome_analysis_group_id_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OutcomeAnalysisGroupID entity.
func (oagi *OutcomeAnalysisGroupID) QueryParent() *OutcomeAnalysisQuery {
	return (&OutcomeAnalysisGroupIDClient{config: oagi.config}).QueryParent(oagi)
}

// Update returns a builder for updating this OutcomeAnalysisGroupID.
// Note that you need to call OutcomeAnalysisGroupID.Unwrap() before calling this method if this OutcomeAnalysisGroupID
// was returned from a transaction, and the transaction was committed or rolled back.
func (oagi *OutcomeAnalysisGroupID) Update() *OutcomeAnalysisGroupIDUpdateOne {
	return (&OutcomeAnalysisGroupIDClient{config: oagi.config}).UpdateOne(oagi)
}

// Unwrap unwraps the OutcomeAnalysisGroupID entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oagi *OutcomeAnalysisGroupID) Unwrap() *OutcomeAnalysisGroupID {
	tx, ok := oagi.config.driver.(*txDriver)
	if !ok {
		panic("models: OutcomeAnalysisGroupID is not a transactional entity")
	}
	oagi.config.driver = tx.drv
	return oagi
}

// String implements the fmt.Stringer.
func (oagi *OutcomeAnalysisGroupID) String() string {
	var builder strings.Builder
	builder.WriteString("OutcomeAnalysisGroupID(")
	builder.WriteString(fmt.Sprintf("id=%v", oagi.ID))
	builder.WriteString(", outcome_analysis_group_id=")
	builder.WriteString(oagi.OutcomeAnalysisGroupID)
	builder.WriteByte(')')
	return builder.String()
}

// OutcomeAnalysisGroupIDs is a parsable slice of OutcomeAnalysisGroupID.
type OutcomeAnalysisGroupIDs []*OutcomeAnalysisGroupID

func (oagi OutcomeAnalysisGroupIDs) config(cfg config) {
	for _i := range oagi {
		oagi[_i].config = cfg
	}
}
