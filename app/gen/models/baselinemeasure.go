// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
)

// BaselineMeasure is the model entity for the BaselineMeasure schema.
type BaselineMeasure struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineMeasureTitle holds the value of the "baseline_measure_title" field.
	BaselineMeasureTitle string `json:"baseline_measure_title,omitempty"`
	// BaselineMeasureDescription holds the value of the "baseline_measure_description" field.
	BaselineMeasureDescription string `json:"baseline_measure_description,omitempty"`
	// BaselineMeasurePopulationDescription holds the value of the "baseline_measure_population_description" field.
	BaselineMeasurePopulationDescription string `json:"baseline_measure_population_description,omitempty"`
	// BaselineMeasureParamType holds the value of the "baseline_measure_param_type" field.
	BaselineMeasureParamType string `json:"baseline_measure_param_type,omitempty"`
	// BaselineMeasureDispersionType holds the value of the "baseline_measure_dispersion_type" field.
	BaselineMeasureDispersionType string `json:"baseline_measure_dispersion_type,omitempty"`
	// BaselineMeasureUnitOfMeasure holds the value of the "baseline_measure_unit_of_measure" field.
	BaselineMeasureUnitOfMeasure string `json:"baseline_measure_unit_of_measure,omitempty"`
	// BaselineMeasureCalculatePct holds the value of the "baseline_measure_calculate_pct" field.
	BaselineMeasureCalculatePct string `json:"baseline_measure_calculate_pct,omitempty"`
	// BaselineMeasureDenomUnitsSelected holds the value of the "baseline_measure_denom_units_selected" field.
	BaselineMeasureDenomUnitsSelected string `json:"baseline_measure_denom_units_selected,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineMeasureQuery when eager-loading is set.
	Edges                                                 BaselineMeasureEdges `json:"edges"`
	baseline_characteristics_module_baseline_measure_list *int
}

// BaselineMeasureEdges holds the relations/edges for other nodes in the graph.
type BaselineMeasureEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineCharacteristicsModule `json:"parent,omitempty"`
	// BaselineMeasureDenomList holds the value of the baseline_measure_denom_list edge.
	BaselineMeasureDenomList []*BaselineMeasureDenom `json:"baseline_measure_denom_list,omitempty"`
	// BaselineClassList holds the value of the baseline_class_list edge.
	BaselineClassList []*BaselineClass `json:"baseline_class_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineMeasureEdges) ParentOrErr() (*BaselineCharacteristicsModule, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baselinecharacteristicsmodule.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// BaselineMeasureDenomListOrErr returns the BaselineMeasureDenomList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineMeasureEdges) BaselineMeasureDenomListOrErr() ([]*BaselineMeasureDenom, error) {
	if e.loadedTypes[1] {
		return e.BaselineMeasureDenomList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_measure_denom_list"}
}

// BaselineClassListOrErr returns the BaselineClassList value or an error if the edge
// was not loaded in eager-loading.
func (e BaselineMeasureEdges) BaselineClassListOrErr() ([]*BaselineClass, error) {
	if e.loadedTypes[2] {
		return e.BaselineClassList, nil
	}
	return nil, &NotLoadedError{edge: "baseline_class_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineMeasure) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinemeasure.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinemeasure.FieldBaselineMeasureTitle, baselinemeasure.FieldBaselineMeasureDescription, baselinemeasure.FieldBaselineMeasurePopulationDescription, baselinemeasure.FieldBaselineMeasureParamType, baselinemeasure.FieldBaselineMeasureDispersionType, baselinemeasure.FieldBaselineMeasureUnitOfMeasure, baselinemeasure.FieldBaselineMeasureCalculatePct, baselinemeasure.FieldBaselineMeasureDenomUnitsSelected:
			values[i] = new(sql.NullString)
		case baselinemeasure.ForeignKeys[0]: // baseline_characteristics_module_baseline_measure_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineMeasure", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineMeasure fields.
func (bm *BaselineMeasure) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinemeasure.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bm.ID = int(value.Int64)
		case baselinemeasure.FieldBaselineMeasureTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_title", values[i])
			} else if value.Valid {
				bm.BaselineMeasureTitle = value.String
			}
		case baselinemeasure.FieldBaselineMeasureDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_description", values[i])
			} else if value.Valid {
				bm.BaselineMeasureDescription = value.String
			}
		case baselinemeasure.FieldBaselineMeasurePopulationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_population_description", values[i])
			} else if value.Valid {
				bm.BaselineMeasurePopulationDescription = value.String
			}
		case baselinemeasure.FieldBaselineMeasureParamType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_param_type", values[i])
			} else if value.Valid {
				bm.BaselineMeasureParamType = value.String
			}
		case baselinemeasure.FieldBaselineMeasureDispersionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_dispersion_type", values[i])
			} else if value.Valid {
				bm.BaselineMeasureDispersionType = value.String
			}
		case baselinemeasure.FieldBaselineMeasureUnitOfMeasure:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_unit_of_measure", values[i])
			} else if value.Valid {
				bm.BaselineMeasureUnitOfMeasure = value.String
			}
		case baselinemeasure.FieldBaselineMeasureCalculatePct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_calculate_pct", values[i])
			} else if value.Valid {
				bm.BaselineMeasureCalculatePct = value.String
			}
		case baselinemeasure.FieldBaselineMeasureDenomUnitsSelected:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_measure_denom_units_selected", values[i])
			} else if value.Valid {
				bm.BaselineMeasureDenomUnitsSelected = value.String
			}
		case baselinemeasure.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_characteristics_module_baseline_measure_list", value)
			} else if value.Valid {
				bm.baseline_characteristics_module_baseline_measure_list = new(int)
				*bm.baseline_characteristics_module_baseline_measure_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineMeasure entity.
func (bm *BaselineMeasure) QueryParent() *BaselineCharacteristicsModuleQuery {
	return (&BaselineMeasureClient{config: bm.config}).QueryParent(bm)
}

// QueryBaselineMeasureDenomList queries the "baseline_measure_denom_list" edge of the BaselineMeasure entity.
func (bm *BaselineMeasure) QueryBaselineMeasureDenomList() *BaselineMeasureDenomQuery {
	return (&BaselineMeasureClient{config: bm.config}).QueryBaselineMeasureDenomList(bm)
}

// QueryBaselineClassList queries the "baseline_class_list" edge of the BaselineMeasure entity.
func (bm *BaselineMeasure) QueryBaselineClassList() *BaselineClassQuery {
	return (&BaselineMeasureClient{config: bm.config}).QueryBaselineClassList(bm)
}

// Update returns a builder for updating this BaselineMeasure.
// Note that you need to call BaselineMeasure.Unwrap() before calling this method if this BaselineMeasure
// was returned from a transaction, and the transaction was committed or rolled back.
func (bm *BaselineMeasure) Update() *BaselineMeasureUpdateOne {
	return (&BaselineMeasureClient{config: bm.config}).UpdateOne(bm)
}

// Unwrap unwraps the BaselineMeasure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bm *BaselineMeasure) Unwrap() *BaselineMeasure {
	tx, ok := bm.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineMeasure is not a transactional entity")
	}
	bm.config.driver = tx.drv
	return bm
}

// String implements the fmt.Stringer.
func (bm *BaselineMeasure) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineMeasure(")
	builder.WriteString(fmt.Sprintf("id=%v", bm.ID))
	builder.WriteString(", baseline_measure_title=")
	builder.WriteString(bm.BaselineMeasureTitle)
	builder.WriteString(", baseline_measure_description=")
	builder.WriteString(bm.BaselineMeasureDescription)
	builder.WriteString(", baseline_measure_population_description=")
	builder.WriteString(bm.BaselineMeasurePopulationDescription)
	builder.WriteString(", baseline_measure_param_type=")
	builder.WriteString(bm.BaselineMeasureParamType)
	builder.WriteString(", baseline_measure_dispersion_type=")
	builder.WriteString(bm.BaselineMeasureDispersionType)
	builder.WriteString(", baseline_measure_unit_of_measure=")
	builder.WriteString(bm.BaselineMeasureUnitOfMeasure)
	builder.WriteString(", baseline_measure_calculate_pct=")
	builder.WriteString(bm.BaselineMeasureCalculatePct)
	builder.WriteString(", baseline_measure_denom_units_selected=")
	builder.WriteString(bm.BaselineMeasureDenomUnitsSelected)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineMeasures is a parsable slice of BaselineMeasure.
type BaselineMeasures []*BaselineMeasure

func (bm BaselineMeasures) config(cfg config) {
	for _i := range bm {
		bm[_i].config = cfg
	}
}
