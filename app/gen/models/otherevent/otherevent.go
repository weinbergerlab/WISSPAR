// Code generated (@generated) by entc, DO NOT EDIT.

package otherevent

const (
	// Label holds the string label denoting the otherevent type in the database.
	Label = "other_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOtherEventTerm holds the string denoting the other_event_term field in the database.
	FieldOtherEventTerm = "other_event_term"
	// FieldOtherEventOrganSystem holds the string denoting the other_event_organ_system field in the database.
	FieldOtherEventOrganSystem = "other_event_organ_system"
	// FieldOtherEventSourceVocabulary holds the string denoting the other_event_source_vocabulary field in the database.
	FieldOtherEventSourceVocabulary = "other_event_source_vocabulary"
	// FieldOtherEventAssessmentType holds the string denoting the other_event_assessment_type field in the database.
	FieldOtherEventAssessmentType = "other_event_assessment_type"
	// FieldOtherEventNotes holds the string denoting the other_event_notes field in the database.
	FieldOtherEventNotes = "other_event_notes"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeOtherEventStatsList holds the string denoting the other_event_stats_list edge name in mutations.
	EdgeOtherEventStatsList = "other_event_stats_list"
	// Table holds the table name of the otherevent in the database.
	Table = "other_event"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "other_event"
	// ParentInverseTable is the table name for the AdverseEventsModule entity.
	// It exists in this package in order to avoid circular dependency with the "adverseeventsmodule" package.
	ParentInverseTable = "adverse_events_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "adverse_events_module_other_event_list"
	// OtherEventStatsListTable is the table that holds the other_event_stats_list relation/edge.
	OtherEventStatsListTable = "other_event_stats"
	// OtherEventStatsListInverseTable is the table name for the OtherEventStats entity.
	// It exists in this package in order to avoid circular dependency with the "othereventstats" package.
	OtherEventStatsListInverseTable = "other_event_stats"
	// OtherEventStatsListColumn is the table column denoting the other_event_stats_list relation/edge.
	OtherEventStatsListColumn = "other_event_other_event_stats_list"
)

// Columns holds all SQL columns for otherevent fields.
var Columns = []string{
	FieldID,
	FieldOtherEventTerm,
	FieldOtherEventOrganSystem,
	FieldOtherEventSourceVocabulary,
	FieldOtherEventAssessmentType,
	FieldOtherEventNotes,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "other_event"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"adverse_events_module_other_event_list",
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
