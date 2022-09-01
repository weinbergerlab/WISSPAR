// Code generated (@generated) by entc, DO NOT EDIT.

package eventgroup

const (
	// Label holds the string label denoting the eventgroup type in the database.
	Label = "event_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventGroupID holds the string denoting the event_group_id field in the database.
	FieldEventGroupID = "event_group_id"
	// FieldEventGroupTitle holds the string denoting the event_group_title field in the database.
	FieldEventGroupTitle = "event_group_title"
	// FieldEventGroupDescription holds the string denoting the event_group_description field in the database.
	FieldEventGroupDescription = "event_group_description"
	// FieldEventGroupDeathsNumAffected holds the string denoting the event_group_deaths_num_affected field in the database.
	FieldEventGroupDeathsNumAffected = "event_group_deaths_num_affected"
	// FieldEventGroupDeathsNumAtRisk holds the string denoting the event_group_deaths_num_at_risk field in the database.
	FieldEventGroupDeathsNumAtRisk = "event_group_deaths_num_at_risk"
	// FieldEventGroupSeriousNumAffected holds the string denoting the event_group_serious_num_affected field in the database.
	FieldEventGroupSeriousNumAffected = "event_group_serious_num_affected"
	// FieldEventGroupSeriousNumAtRisk holds the string denoting the event_group_serious_num_at_risk field in the database.
	FieldEventGroupSeriousNumAtRisk = "event_group_serious_num_at_risk"
	// FieldEventGroupOtherNumAffected holds the string denoting the event_group_other_num_affected field in the database.
	FieldEventGroupOtherNumAffected = "event_group_other_num_affected"
	// FieldEventGroupOtherNumAtRisk holds the string denoting the event_group_other_num_at_risk field in the database.
	FieldEventGroupOtherNumAtRisk = "event_group_other_num_at_risk"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the eventgroup in the database.
	Table = "event_group"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "event_group"
	// ParentInverseTable is the table name for the AdverseEventsModule entity.
	// It exists in this package in order to avoid circular dependency with the "adverseeventsmodule" package.
	ParentInverseTable = "adverse_events_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "adverse_events_module_event_group_list"
)

// Columns holds all SQL columns for eventgroup fields.
var Columns = []string{
	FieldID,
	FieldEventGroupID,
	FieldEventGroupTitle,
	FieldEventGroupDescription,
	FieldEventGroupDeathsNumAffected,
	FieldEventGroupDeathsNumAtRisk,
	FieldEventGroupSeriousNumAffected,
	FieldEventGroupSeriousNumAtRisk,
	FieldEventGroupOtherNumAffected,
	FieldEventGroupOtherNumAtRisk,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "event_group"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"adverse_events_module_event_group_list",
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
