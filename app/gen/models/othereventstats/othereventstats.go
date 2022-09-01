// Code generated (@generated) by entc, DO NOT EDIT.

package othereventstats

const (
	// Label holds the string label denoting the othereventstats type in the database.
	Label = "other_event_stats"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOtherEventStatsGroupID holds the string denoting the other_event_stats_group_id field in the database.
	FieldOtherEventStatsGroupID = "other_event_stats_group_id"
	// FieldOtherEventStatsNumEvents holds the string denoting the other_event_stats_num_events field in the database.
	FieldOtherEventStatsNumEvents = "other_event_stats_num_events"
	// FieldOtherEventStatsNumAffected holds the string denoting the other_event_stats_num_affected field in the database.
	FieldOtherEventStatsNumAffected = "other_event_stats_num_affected"
	// FieldOtherEventStatsNumAtRisk holds the string denoting the other_event_stats_num_at_risk field in the database.
	FieldOtherEventStatsNumAtRisk = "other_event_stats_num_at_risk"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the othereventstats in the database.
	Table = "other_event_stats"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "other_event_stats"
	// ParentInverseTable is the table name for the OtherEvent entity.
	// It exists in this package in order to avoid circular dependency with the "otherevent" package.
	ParentInverseTable = "other_event"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "other_event_other_event_stats_list"
)

// Columns holds all SQL columns for othereventstats fields.
var Columns = []string{
	FieldID,
	FieldOtherEventStatsGroupID,
	FieldOtherEventStatsNumEvents,
	FieldOtherEventStatsNumAffected,
	FieldOtherEventStatsNumAtRisk,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "other_event_stats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"other_event_other_event_stats_list",
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
