// Code generated (@generated) by entc, DO NOT EDIT.

package flowmilestone

const (
	// Label holds the string label denoting the flowmilestone type in the database.
	Label = "flow_milestone"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowMilestoneType holds the string denoting the flow_milestone_type field in the database.
	FieldFlowMilestoneType = "flow_milestone_type"
	// FieldFlowMilestoneComment holds the string denoting the flow_milestone_comment field in the database.
	FieldFlowMilestoneComment = "flow_milestone_comment"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeFlowAchievementList holds the string denoting the flow_achievement_list edge name in mutations.
	EdgeFlowAchievementList = "flow_achievement_list"
	// Table holds the table name of the flowmilestone in the database.
	Table = "flow_milestone"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_milestone"
	// ParentInverseTable is the table name for the FlowPeriod entity.
	// It exists in this package in order to avoid circular dependency with the "flowperiod" package.
	ParentInverseTable = "flow_period"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "flow_period_flow_milestone_list"
	// FlowAchievementListTable is the table that holds the flow_achievement_list relation/edge.
	FlowAchievementListTable = "flow_achievement"
	// FlowAchievementListInverseTable is the table name for the FlowAchievement entity.
	// It exists in this package in order to avoid circular dependency with the "flowachievement" package.
	FlowAchievementListInverseTable = "flow_achievement"
	// FlowAchievementListColumn is the table column denoting the flow_achievement_list relation/edge.
	FlowAchievementListColumn = "flow_milestone_flow_achievement_list"
)

// Columns holds all SQL columns for flowmilestone fields.
var Columns = []string{
	FieldID,
	FieldFlowMilestoneType,
	FieldFlowMilestoneComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_milestone"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"flow_period_flow_milestone_list",
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
