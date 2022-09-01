// Code generated (@generated) by entc, DO NOT EDIT.

package flowachievement

const (
	// Label holds the string label denoting the flowachievement type in the database.
	Label = "flow_achievement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFlowAchievementGroupID holds the string denoting the flow_achievement_group_id field in the database.
	FieldFlowAchievementGroupID = "flow_achievement_group_id"
	// FieldFlowAchievementComment holds the string denoting the flow_achievement_comment field in the database.
	FieldFlowAchievementComment = "flow_achievement_comment"
	// FieldFlowAchievementNumSubjects holds the string denoting the flow_achievement_num_subjects field in the database.
	FieldFlowAchievementNumSubjects = "flow_achievement_num_subjects"
	// FieldFlowAchievementNumUnits holds the string denoting the flow_achievement_num_units field in the database.
	FieldFlowAchievementNumUnits = "flow_achievement_num_units"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the flowachievement in the database.
	Table = "flow_achievement"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "flow_achievement"
	// ParentInverseTable is the table name for the FlowMilestone entity.
	// It exists in this package in order to avoid circular dependency with the "flowmilestone" package.
	ParentInverseTable = "flow_milestone"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "flow_milestone_flow_achievement_list"
)

// Columns holds all SQL columns for flowachievement fields.
var Columns = []string{
	FieldID,
	FieldFlowAchievementGroupID,
	FieldFlowAchievementComment,
	FieldFlowAchievementNumSubjects,
	FieldFlowAchievementNumUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flow_achievement"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"flow_milestone_flow_achievement_list",
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
