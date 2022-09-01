// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
)

// BaselineGroup is the model entity for the BaselineGroup schema.
type BaselineGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BaselineGroupID holds the value of the "baseline_group_id" field.
	BaselineGroupID string `json:"baseline_group_id,omitempty"`
	// BaselineGroupTitle holds the value of the "baseline_group_title" field.
	BaselineGroupTitle string `json:"baseline_group_title,omitempty"`
	// BaselineGroupDescription holds the value of the "baseline_group_description" field.
	BaselineGroupDescription string `json:"baseline_group_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BaselineGroupQuery when eager-loading is set.
	Edges                                               BaselineGroupEdges `json:"edges"`
	baseline_characteristics_module_baseline_group_list *int
}

// BaselineGroupEdges holds the relations/edges for other nodes in the graph.
type BaselineGroupEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BaselineCharacteristicsModule `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaselineGroupEdges) ParentOrErr() (*BaselineCharacteristicsModule, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*BaselineGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baselinegroup.FieldID:
			values[i] = new(sql.NullInt64)
		case baselinegroup.FieldBaselineGroupID, baselinegroup.FieldBaselineGroupTitle, baselinegroup.FieldBaselineGroupDescription:
			values[i] = new(sql.NullString)
		case baselinegroup.ForeignKeys[0]: // baseline_characteristics_module_baseline_group_list
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BaselineGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaselineGroup fields.
func (bg *BaselineGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baselinegroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bg.ID = int(value.Int64)
		case baselinegroup.FieldBaselineGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_group_id", values[i])
			} else if value.Valid {
				bg.BaselineGroupID = value.String
			}
		case baselinegroup.FieldBaselineGroupTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_group_title", values[i])
			} else if value.Valid {
				bg.BaselineGroupTitle = value.String
			}
		case baselinegroup.FieldBaselineGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseline_group_description", values[i])
			} else if value.Valid {
				bg.BaselineGroupDescription = value.String
			}
		case baselinegroup.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field baseline_characteristics_module_baseline_group_list", value)
			} else if value.Valid {
				bg.baseline_characteristics_module_baseline_group_list = new(int)
				*bg.baseline_characteristics_module_baseline_group_list = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the BaselineGroup entity.
func (bg *BaselineGroup) QueryParent() *BaselineCharacteristicsModuleQuery {
	return (&BaselineGroupClient{config: bg.config}).QueryParent(bg)
}

// Update returns a builder for updating this BaselineGroup.
// Note that you need to call BaselineGroup.Unwrap() before calling this method if this BaselineGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (bg *BaselineGroup) Update() *BaselineGroupUpdateOne {
	return (&BaselineGroupClient{config: bg.config}).UpdateOne(bg)
}

// Unwrap unwraps the BaselineGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bg *BaselineGroup) Unwrap() *BaselineGroup {
	tx, ok := bg.config.driver.(*txDriver)
	if !ok {
		panic("models: BaselineGroup is not a transactional entity")
	}
	bg.config.driver = tx.drv
	return bg
}

// String implements the fmt.Stringer.
func (bg *BaselineGroup) String() string {
	var builder strings.Builder
	builder.WriteString("BaselineGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", bg.ID))
	builder.WriteString(", baseline_group_id=")
	builder.WriteString(bg.BaselineGroupID)
	builder.WriteString(", baseline_group_title=")
	builder.WriteString(bg.BaselineGroupTitle)
	builder.WriteString(", baseline_group_description=")
	builder.WriteString(bg.BaselineGroupDescription)
	builder.WriteByte(')')
	return builder.String()
}

// BaselineGroups is a parsable slice of BaselineGroup.
type BaselineGroups []*BaselineGroup

func (bg BaselineGroups) config(cfg config) {
	for _i := range bg {
		bg[_i].config = cfg
	}
}
