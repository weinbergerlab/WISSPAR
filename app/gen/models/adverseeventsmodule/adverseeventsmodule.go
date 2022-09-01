// Code generated (@generated) by entc, DO NOT EDIT.

package adverseeventsmodule

const (
	// Label holds the string label denoting the adverseeventsmodule type in the database.
	Label = "adverse_events_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventsFrequencyThreshold holds the string denoting the events_frequency_threshold field in the database.
	FieldEventsFrequencyThreshold = "events_frequency_threshold"
	// FieldEventsTimeFrame holds the string denoting the events_time_frame field in the database.
	FieldEventsTimeFrame = "events_time_frame"
	// FieldEventsDescription holds the string denoting the events_description field in the database.
	FieldEventsDescription = "events_description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeEventGroupList holds the string denoting the event_group_list edge name in mutations.
	EdgeEventGroupList = "event_group_list"
	// EdgeSeriousEventList holds the string denoting the serious_event_list edge name in mutations.
	EdgeSeriousEventList = "serious_event_list"
	// EdgeOtherEventList holds the string denoting the other_event_list edge name in mutations.
	EdgeOtherEventList = "other_event_list"
	// Table holds the table name of the adverseeventsmodule in the database.
	Table = "adverse_events_module"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "adverse_events_module"
	// ParentInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ParentInverseTable = "results_definition"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "results_definition_adverse_events_module"
	// EventGroupListTable is the table that holds the event_group_list relation/edge.
	EventGroupListTable = "event_group"
	// EventGroupListInverseTable is the table name for the EventGroup entity.
	// It exists in this package in order to avoid circular dependency with the "eventgroup" package.
	EventGroupListInverseTable = "event_group"
	// EventGroupListColumn is the table column denoting the event_group_list relation/edge.
	EventGroupListColumn = "adverse_events_module_event_group_list"
	// SeriousEventListTable is the table that holds the serious_event_list relation/edge.
	SeriousEventListTable = "serious_event"
	// SeriousEventListInverseTable is the table name for the SeriousEvent entity.
	// It exists in this package in order to avoid circular dependency with the "seriousevent" package.
	SeriousEventListInverseTable = "serious_event"
	// SeriousEventListColumn is the table column denoting the serious_event_list relation/edge.
	SeriousEventListColumn = "adverse_events_module_serious_event_list"
	// OtherEventListTable is the table that holds the other_event_list relation/edge.
	OtherEventListTable = "other_event"
	// OtherEventListInverseTable is the table name for the OtherEvent entity.
	// It exists in this package in order to avoid circular dependency with the "otherevent" package.
	OtherEventListInverseTable = "other_event"
	// OtherEventListColumn is the table column denoting the other_event_list relation/edge.
	OtherEventListColumn = "adverse_events_module_other_event_list"
)

// Columns holds all SQL columns for adverseeventsmodule fields.
var Columns = []string{
	FieldID,
	FieldEventsFrequencyThreshold,
	FieldEventsTimeFrame,
	FieldEventsDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "adverse_events_module"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"results_definition_adverse_events_module",
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
