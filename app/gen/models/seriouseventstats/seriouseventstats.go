// Code generated (@generated) by entc, DO NOT EDIT.

package seriouseventstats

const (
	// Label holds the string label denoting the seriouseventstats type in the database.
	Label = "serious_event_stats"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeriousEventStatsGroupID holds the string denoting the serious_event_stats_group_id field in the database.
	FieldSeriousEventStatsGroupID = "serious_event_stats_group_id"
	// FieldSeriousEventStatsNumEvents holds the string denoting the serious_event_stats_num_events field in the database.
	FieldSeriousEventStatsNumEvents = "serious_event_stats_num_events"
	// FieldSeriousEventStatsNumAffected holds the string denoting the serious_event_stats_num_affected field in the database.
	FieldSeriousEventStatsNumAffected = "serious_event_stats_num_affected"
	// FieldSeriousEventStatsNumAtRisk holds the string denoting the serious_event_stats_num_at_risk field in the database.
	FieldSeriousEventStatsNumAtRisk = "serious_event_stats_num_at_risk"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the seriouseventstats in the database.
	Table = "serious_event_stats"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "serious_event_stats"
	// ParentInverseTable is the table name for the SeriousEvent entity.
	// It exists in this package in order to avoid circular dependency with the "seriousevent" package.
	ParentInverseTable = "serious_event"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "serious_event_serious_event_stats_list"
)

// Columns holds all SQL columns for seriouseventstats fields.
var Columns = []string{
	FieldID,
	FieldSeriousEventStatsGroupID,
	FieldSeriousEventStatsNumEvents,
	FieldSeriousEventStatsNumAffected,
	FieldSeriousEventStatsNumAtRisk,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "serious_event_stats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"serious_event_serious_event_stats_list",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
