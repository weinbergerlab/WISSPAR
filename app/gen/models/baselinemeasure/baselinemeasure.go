// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasure

const (
	// Label holds the string label denoting the baselinemeasure type in the database.
	Label = "baseline_measure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBaselineMeasureTitle holds the string denoting the baseline_measure_title field in the database.
	FieldBaselineMeasureTitle = "baseline_measure_title"
	// FieldBaselineMeasureDescription holds the string denoting the baseline_measure_description field in the database.
	FieldBaselineMeasureDescription = "baseline_measure_description"
	// FieldBaselineMeasurePopulationDescription holds the string denoting the baseline_measure_population_description field in the database.
	FieldBaselineMeasurePopulationDescription = "baseline_measure_population_description"
	// FieldBaselineMeasureParamType holds the string denoting the baseline_measure_param_type field in the database.
	FieldBaselineMeasureParamType = "baseline_measure_param_type"
	// FieldBaselineMeasureDispersionType holds the string denoting the baseline_measure_dispersion_type field in the database.
	FieldBaselineMeasureDispersionType = "baseline_measure_dispersion_type"
	// FieldBaselineMeasureUnitOfMeasure holds the string denoting the baseline_measure_unit_of_measure field in the database.
	FieldBaselineMeasureUnitOfMeasure = "baseline_measure_unit_of_measure"
	// FieldBaselineMeasureCalculatePct holds the string denoting the baseline_measure_calculate_pct field in the database.
	FieldBaselineMeasureCalculatePct = "baseline_measure_calculate_pct"
	// FieldBaselineMeasureDenomUnitsSelected holds the string denoting the baseline_measure_denom_units_selected field in the database.
	FieldBaselineMeasureDenomUnitsSelected = "baseline_measure_denom_units_selected"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeBaselineMeasureDenomList holds the string denoting the baseline_measure_denom_list edge name in mutations.
	EdgeBaselineMeasureDenomList = "baseline_measure_denom_list"
	// EdgeBaselineClassList holds the string denoting the baseline_class_list edge name in mutations.
	EdgeBaselineClassList = "baseline_class_list"
	// Table holds the table name of the baselinemeasure in the database.
	Table = "baseline_measure"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "baseline_measure"
	// ParentInverseTable is the table name for the BaselineCharacteristicsModule entity.
	// It exists in this package in order to avoid circular dependency with the "baselinecharacteristicsmodule" package.
	ParentInverseTable = "baseline_characteristics_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "baseline_characteristics_module_baseline_measure_list"
	// BaselineMeasureDenomListTable is the table that holds the baseline_measure_denom_list relation/edge.
	BaselineMeasureDenomListTable = "baseline_measure_denom"
	// BaselineMeasureDenomListInverseTable is the table name for the BaselineMeasureDenom entity.
	// It exists in this package in order to avoid circular dependency with the "baselinemeasuredenom" package.
	BaselineMeasureDenomListInverseTable = "baseline_measure_denom"
	// BaselineMeasureDenomListColumn is the table column denoting the baseline_measure_denom_list relation/edge.
	BaselineMeasureDenomListColumn = "baseline_measure_baseline_measure_denom_list"
	// BaselineClassListTable is the table that holds the baseline_class_list relation/edge.
	BaselineClassListTable = "baseline_class"
	// BaselineClassListInverseTable is the table name for the BaselineClass entity.
	// It exists in this package in order to avoid circular dependency with the "baselineclass" package.
	BaselineClassListInverseTable = "baseline_class"
	// BaselineClassListColumn is the table column denoting the baseline_class_list relation/edge.
	BaselineClassListColumn = "baseline_measure_baseline_class_list"
)

// Columns holds all SQL columns for baselinemeasure fields.
var Columns = []string{
	FieldID,
	FieldBaselineMeasureTitle,
	FieldBaselineMeasureDescription,
	FieldBaselineMeasurePopulationDescription,
	FieldBaselineMeasureParamType,
	FieldBaselineMeasureDispersionType,
	FieldBaselineMeasureUnitOfMeasure,
	FieldBaselineMeasureCalculatePct,
	FieldBaselineMeasureDenomUnitsSelected,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "baseline_measure"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"baseline_characteristics_module_baseline_measure_list",
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
