// Code generated (@generated) by entc, DO NOT EDIT.

package seriousevent

const (
	// Label holds the string label denoting the seriousevent type in the database.
	Label = "serious_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeriousEventTerm holds the string denoting the serious_event_term field in the database.
	FieldSeriousEventTerm = "serious_event_term"
	// FieldSeriousEventOrganSystem holds the string denoting the serious_event_organ_system field in the database.
	FieldSeriousEventOrganSystem = "serious_event_organ_system"
	// FieldSeriousEventSourceVocabulary holds the string denoting the serious_event_source_vocabulary field in the database.
	FieldSeriousEventSourceVocabulary = "serious_event_source_vocabulary"
	// FieldSeriousEventAssessmentType holds the string denoting the serious_event_assessment_type field in the database.
	FieldSeriousEventAssessmentType = "serious_event_assessment_type"
	// FieldSeriousEventNotes holds the string denoting the serious_event_notes field in the database.
	FieldSeriousEventNotes = "serious_event_notes"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeSeriousEventStatsList holds the string denoting the serious_event_stats_list edge name in mutations.
	EdgeSeriousEventStatsList = "serious_event_stats_list"
	// Table holds the table name of the seriousevent in the database.
	Table = "serious_event"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "serious_event"
	// ParentInverseTable is the table name for the AdverseEventsModule entity.
	// It exists in this package in order to avoid circular dependency with the "adverseeventsmodule" package.
	ParentInverseTable = "adverse_events_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "adverse_events_module_serious_event_list"
	// SeriousEventStatsListTable is the table that holds the serious_event_stats_list relation/edge.
	SeriousEventStatsListTable = "serious_event_stats"
	// SeriousEventStatsListInverseTable is the table name for the SeriousEventStats entity.
	// It exists in this package in order to avoid circular dependency with the "seriouseventstats" package.
	SeriousEventStatsListInverseTable = "serious_event_stats"
	// SeriousEventStatsListColumn is the table column denoting the serious_event_stats_list relation/edge.
	SeriousEventStatsListColumn = "serious_event_serious_event_stats_list"
)

// Columns holds all SQL columns for seriousevent fields.
var Columns = []string{
	FieldID,
	FieldSeriousEventTerm,
	FieldSeriousEventOrganSystem,
	FieldSeriousEventSourceVocabulary,
	FieldSeriousEventAssessmentType,
	FieldSeriousEventNotes,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "serious_event"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"adverse_events_module_serious_event_list",
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
