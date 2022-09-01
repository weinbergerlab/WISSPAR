// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
)

// OtherEventStats is the model entity for the OtherEventStats schema.
type OtherEventStats struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OtherEventStatsGroupID holds the value of the "other_event_stats_group_id" field.
	OtherEventStatsGroupID string `json:"other_event_stats_group_id,omitempty"`
	// OtherEventStatsNumEvents holds the value of the "other_event_stats_num_events" field.
	OtherEventStatsNumEvents string `json:"other_event_stats_num_events,omitempty"`
	// OtherEventStatsNumAffected holds the value of the "other_event_stats_num_affected" field.
	OtherEventStatsNumAffected string `json:"other_event_stats_num_affected,omitempty"`
	// OtherEventStatsNumAtRisk holds the value of the "other_event_stats_num_at_risk" field.
	OtherEventStatsNumAtRisk string `json:"other_event_stats_num_at_risk,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OtherEventStatsQuery when eager-loading is set.
	Edges                              OtherEventStatsEdges `json:"edges"`
	other_event_other_event_stats_list *int
}

// OtherEventStatsEdges holds the relations/edges for other nodes in the graph.
type OtherEventStatsEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OtherEvent `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OtherEventStatsEdges) ParentOrErr() (*OtherEvent, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: otherevent.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OtherEventStats) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case othereventstats.FieldID:
			values[i] = new(sql.NullInt64)
		case othereventstats.FieldOtherEventStatsGroupID, othereventstats.FieldOtherEventStatsNumEvents, othereventstats.FieldOtherEventStatsNumAffected, othereventstats.FieldOtherEventStatsNumAtRisk:
			values[i] = new(sql.NullString)
		case othereventstats.ForeignKeys[0]: // other_event_other_event_stats_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OtherEventStats", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OtherEventStats fields.
func (oes *OtherEventStats) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case othereventstats.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oes.ID = int(value.Int64)
		case othereventstats.FieldOtherEventStatsGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_stats_group_id", values[i])
			} else if value.Valid {
				oes.OtherEventStatsGroupID = value.String
			}
		case othereventstats.FieldOtherEventStatsNumEvents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_stats_num_events", values[i])
			} else if value.Valid {
				oes.OtherEventStatsNumEvents = value.String
			}
		case othereventstats.FieldOtherEventStatsNumAffected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_stats_num_affected", values[i])
			} else if value.Valid {
				oes.OtherEventStatsNumAffected = value.String
			}
		case othereventstats.FieldOtherEventStatsNumAtRisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_event_stats_num_at_risk", values[i])
			} else if value.Valid {
				oes.OtherEventStatsNumAtRisk = value.String
			}
		case othereventstats.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field other_event_other_event_stats_list", value)
			} else if value.Valid {
				oes.other_event_other_event_stats_list = new(int)
				*oes.other_event_other_event_stats_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OtherEventStats entity.
func (oes *OtherEventStats) QueryParent() *OtherEventQuery {
	return (&OtherEventStatsClient{config: oes.config}).QueryParent(oes)
}

// Update returns a builder for updating this OtherEventStats.
// Note that you need to call OtherEventStats.Unwrap() before calling this method if this OtherEventStats
// was returned from a transaction, and the transaction was committed or rolled back.
func (oes *OtherEventStats) Update() *OtherEventStatsUpdateOne {
	return (&OtherEventStatsClient{config: oes.config}).UpdateOne(oes)
}

// Unwrap unwraps the OtherEventStats entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oes *OtherEventStats) Unwrap() *OtherEventStats {
	tx, ok := oes.config.driver.(*txDriver)
	if !ok {
		panic("models: OtherEventStats is not a transactional entity")
	}
	oes.config.driver = tx.drv
	return oes
}

// String implements the fmt.Stringer.
func (oes *OtherEventStats) String() string {
	var builder strings.Builder
	builder.WriteString("OtherEventStats(")
	builder.WriteString(fmt.Sprintf("id=%v", oes.ID))
	builder.WriteString(", other_event_stats_group_id=")
	builder.WriteString(oes.OtherEventStatsGroupID)
	builder.WriteString(", other_event_stats_num_events=")
	builder.WriteString(oes.OtherEventStatsNumEvents)
	builder.WriteString(", other_event_stats_num_affected=")
	builder.WriteString(oes.OtherEventStatsNumAffected)
	builder.WriteString(", other_event_stats_num_at_risk=")
	builder.WriteString(oes.OtherEventStatsNumAtRisk)
	builder.WriteByte(')')
	return builder.String()
}

// OtherEventStatsSlice is a parsable slice of OtherEventStats.
type OtherEventStatsSlice []*OtherEventStats

func (oes OtherEventStatsSlice) config(cfg config) {
	for _i := range oes {
		oes[_i].config = cfg
	}
}
