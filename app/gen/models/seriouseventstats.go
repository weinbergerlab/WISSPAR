// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventStats is the model entity for the SeriousEventStats schema.
type SeriousEventStats struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SeriousEventStatsGroupID holds the value of the "serious_event_stats_group_id" field.
	SeriousEventStatsGroupID string `json:"serious_event_stats_group_id,omitempty"`
	// SeriousEventStatsNumEvents holds the value of the "serious_event_stats_num_events" field.
	SeriousEventStatsNumEvents string `json:"serious_event_stats_num_events,omitempty"`
	// SeriousEventStatsNumAffected holds the value of the "serious_event_stats_num_affected" field.
	SeriousEventStatsNumAffected string `json:"serious_event_stats_num_affected,omitempty"`
	// SeriousEventStatsNumAtRisk holds the value of the "serious_event_stats_num_at_risk" field.
	SeriousEventStatsNumAtRisk string `json:"serious_event_stats_num_at_risk,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeriousEventStatsQuery when eager-loading is set.
	Edges                                  SeriousEventStatsEdges `json:"edges"`
	serious_event_serious_event_stats_list *int
}

// SeriousEventStatsEdges holds the relations/edges for other nodes in the graph.
type SeriousEventStatsEdges struct {
	// Parent holds the value of the parent edge.
	Parent *SeriousEvent `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeriousEventStatsEdges) ParentOrErr() (*SeriousEvent, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: seriousevent.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SeriousEventStats) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case seriouseventstats.FieldID:
			values[i] = new(sql.NullInt64)
		case seriouseventstats.FieldSeriousEventStatsGroupID, seriouseventstats.FieldSeriousEventStatsNumEvents, seriouseventstats.FieldSeriousEventStatsNumAffected, seriouseventstats.FieldSeriousEventStatsNumAtRisk:
			values[i] = new(sql.NullString)
		case seriouseventstats.ForeignKeys[0]: // serious_event_serious_event_stats_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SeriousEventStats", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SeriousEventStats fields.
func (ses *SeriousEventStats) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seriouseventstats.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ses.ID = int(value.Int64)
		case seriouseventstats.FieldSeriousEventStatsGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_stats_group_id", values[i])
			} else if value.Valid {
				ses.SeriousEventStatsGroupID = value.String
			}
		case seriouseventstats.FieldSeriousEventStatsNumEvents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_stats_num_events", values[i])
			} else if value.Valid {
				ses.SeriousEventStatsNumEvents = value.String
			}
		case seriouseventstats.FieldSeriousEventStatsNumAffected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_stats_num_affected", values[i])
			} else if value.Valid {
				ses.SeriousEventStatsNumAffected = value.String
			}
		case seriouseventstats.FieldSeriousEventStatsNumAtRisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serious_event_stats_num_at_risk", values[i])
			} else if value.Valid {
				ses.SeriousEventStatsNumAtRisk = value.String
			}
		case seriouseventstats.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field serious_event_serious_event_stats_list", value)
			} else if value.Valid {
				ses.serious_event_serious_event_stats_list = new(int)
				*ses.serious_event_serious_event_stats_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the SeriousEventStats entity.
func (ses *SeriousEventStats) QueryParent() *SeriousEventQuery {
	return (&SeriousEventStatsClient{config: ses.config}).QueryParent(ses)
}

// Update returns a builder for updating this SeriousEventStats.
// Note that you need to call SeriousEventStats.Unwrap() before calling this method if this SeriousEventStats
// was returned from a transaction, and the transaction was committed or rolled back.
func (ses *SeriousEventStats) Update() *SeriousEventStatsUpdateOne {
	return (&SeriousEventStatsClient{config: ses.config}).UpdateOne(ses)
}

// Unwrap unwraps the SeriousEventStats entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ses *SeriousEventStats) Unwrap() *SeriousEventStats {
	tx, ok := ses.config.driver.(*txDriver)
	if !ok {
		panic("models: SeriousEventStats is not a transactional entity")
	}
	ses.config.driver = tx.drv
	return ses
}

// String implements the fmt.Stringer.
func (ses *SeriousEventStats) String() string {
	var builder strings.Builder
	builder.WriteString("SeriousEventStats(")
	builder.WriteString(fmt.Sprintf("id=%v", ses.ID))
	builder.WriteString(", serious_event_stats_group_id=")
	builder.WriteString(ses.SeriousEventStatsGroupID)
	builder.WriteString(", serious_event_stats_num_events=")
	builder.WriteString(ses.SeriousEventStatsNumEvents)
	builder.WriteString(", serious_event_stats_num_affected=")
	builder.WriteString(ses.SeriousEventStatsNumAffected)
	builder.WriteString(", serious_event_stats_num_at_risk=")
	builder.WriteString(ses.SeriousEventStatsNumAtRisk)
	builder.WriteByte(')')
	return builder.String()
}

// SeriousEventStatsSlice is a parsable slice of SeriousEventStats.
type SeriousEventStatsSlice []*SeriousEventStats

func (ses SeriousEventStatsSlice) config(cfg config) {
	for _i := range ses {
		ses[_i].config = cfg
	}
}
