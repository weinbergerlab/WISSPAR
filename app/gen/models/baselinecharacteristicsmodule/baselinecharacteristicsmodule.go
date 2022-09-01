// Code generated (@generated) by entc, DO NOT EDIT.

package baselinecharacteristicsmodule

const (
	// Label holds the string label denoting the baselinecharacteristicsmodule type in the database.
	Label = "baseline_characteristics_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselinePopulationDescription holds the string denoting the baseline_population_description field in the database.
	FieldBaselinePopulationDescription = "baseline_population_description"
	// FieldBaselineTypeUnitsAnalyzed holds the string denoting the baseline_type_units_analyzed field in the database.
	FieldBaselineTypeUnitsAnalyzed = "baseline_type_units_analyzed"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineGroupList holds the string denoting the baseline_group_list edge name in mutations.
	EdgeBaselineGroupList = "baseline_group_list"
	// EdgeBaselineDenomList holds the string denoting the baseline_denom_list edge name in mutations.
	EdgeBaselineDenomList = "baseline_denom_list"
	// EdgeBaselineMeasureList holds the string denoting the baseline_measure_list edge name in mutations.
	EdgeBaselineMeasureList = "baseline_measure_list"
	// Table holds the table name of the baselinecharacteristicsmodule in the database.
	Table = "baseline_characteristics_module"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_characteristics_module"
	// ParentInverseTable is the table name for the ResultsDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resultsdefinition" package.
	ParentInverseTable = "results_definition"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "results_definition_baseline_characteristics_module"
	// BaselineGroupListTable is the table that holds the baseline_group_list relation/edge.
	BaselineGroupListTable = "baseline_group"
	// BaselineGroupListInverseTable is the table name for the BaselineGroup entity.
	// It exists in this package in order to avoid circular dependency with the "baselinegroup" package.
	BaselineGroupListInverseTable = "baseline_group"
	// BaselineGroupListColumn is the table column denoting the baseline_group_list relation/edge.
	BaselineGroupListColumn = "baseline_characteristics_module_baseline_group_list"
	// BaselineDenomListTable is the table that holds the baseline_denom_list relation/edge.
	BaselineDenomListTable = "baseline_denom"
	// BaselineDenomListInverseTable is the table name for the BaselineDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselinedenom" package.
	BaselineDenomListInverseTable = "baseline_denom"
	// BaselineDenomListColumn is the table column denoting the baseline_denom_list relation/edge.
	BaselineDenomListColumn = "baseline_characteristics_module_baseline_denom_list"
	// BaselineMeasureListTable is the table that holds the baseline_measure_list relation/edge.
	BaselineMeasureListTable = "baseline_measure"
	// BaselineMeasureListInverseTable is the table name for the BaselineMeasure entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasure" package.
	BaselineMeasureListInverseTable = "baseline_measure"
	// BaselineMeasureListColumn is the table column denoting the baseline_measure_list relation/edge.
	BaselineMeasureListColumn = "baseline_characteristics_module_baseline_measure_list"
)

// Columns holds all SQL columns for baselinecharacteristicsmodule fields.
var Columns = []string{
	FieldID,
	FieldBaselinePopulationDescription,
	FieldBaselineTypeUnitsAnalyzed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_characteristics_module"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"results_definition_baseline_characteristics_module",
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
