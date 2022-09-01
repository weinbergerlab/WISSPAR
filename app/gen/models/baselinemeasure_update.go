// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureUpdate is the builder for updating BaselineMeasure entities.
type BaselineMeasureUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureMutation
}

// Where appends a list predicates to the BaselineMeasureUpdate builder.
func (bmu *BaselineMeasureUpdate) Where(ps ...predicate.BaselineMeasure) *BaselineMeasureUpdate {
	bmu.mutation.Where(ps...)
	return bmu
}

// SetBaselineMeasureTitle sets the "baseline_measure_title" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureTitle(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureTitle(s)
	return bmu
}

// SetBaselineMeasureDescription sets the "baseline_measure_description" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureDescription(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureDescription(s)
	return bmu
}

// SetBaselineMeasurePopulationDescription sets the "baseline_measure_population_description" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasurePopulationDescription(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasurePopulationDescription(s)
	return bmu
}

// SetBaselineMeasureParamType sets the "baseline_measure_param_type" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureParamType(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureParamType(s)
	return bmu
}

// SetBaselineMeasureDispersionType sets the "baseline_measure_dispersion_type" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureDispersionType(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureDispersionType(s)
	return bmu
}

// SetBaselineMeasureUnitOfMeasure sets the "baseline_measure_unit_of_measure" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureUnitOfMeasure(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureUnitOfMeasure(s)
	return bmu
}

// SetBaselineMeasureCalculatePct sets the "baseline_measure_calculate_pct" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureCalculatePct(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureCalculatePct(s)
	return bmu
}

// SetBaselineMeasureDenomUnitsSelected sets the "baseline_measure_denom_units_selected" field.
func (bmu *BaselineMeasureUpdate) SetBaselineMeasureDenomUnitsSelected(s string) *BaselineMeasureUpdate {
	bmu.mutation.SetBaselineMeasureDenomUnitsSelected(s)
	return bmu
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bmu *BaselineMeasureUpdate) SetParentID(id int) *BaselineMeasureUpdate {
	bmu.mutation.SetParentID(id)
	return bmu
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bmu *BaselineMeasureUpdate) SetNillableParentID(id *int) *BaselineMeasureUpdate {
	if id != nil {
		bmu = bmu.SetParentID(*id)
	}
	return bmu
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bmu *BaselineMeasureUpdate) SetParent(b *BaselineCharacteristicsModule) *BaselineMeasureUpdate {
	return bmu.SetParentID(b.ID)
}

// AddBaselineMeasureDenomListIDs adds the "baseline_measure_denom_list" edge to the BaselineMeasureDenom entity by IDs.
func (bmu *BaselineMeasureUpdate) AddBaselineMeasureDenomListIDs(ids ...int) *BaselineMeasureUpdate {
	bmu.mutation.AddBaselineMeasureDenomListIDs(ids...)
	return bmu
}

// AddBaselineMeasureDenomList adds the "baseline_measure_denom_list" edges to the BaselineMeasureDenom entity.
func (bmu *BaselineMeasureUpdate) AddBaselineMeasureDenomList(b ...*BaselineMeasureDenom) *BaselineMeasureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmu.AddBaselineMeasureDenomListIDs(ids...)
}

// AddBaselineClassListIDs adds the "baseline_class_list" edge to the BaselineClass entity by IDs.
func (bmu *BaselineMeasureUpdate) AddBaselineClassListIDs(ids ...int) *BaselineMeasureUpdate {
	bmu.mutation.AddBaselineClassListIDs(ids...)
	return bmu
}

// AddBaselineClassList adds the "baseline_class_list" edges to the BaselineClass entity.
func (bmu *BaselineMeasureUpdate) AddBaselineClassList(b ...*BaselineClass) *BaselineMeasureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmu.AddBaselineClassListIDs(ids...)
}

// Mutation returns the BaselineMeasureMutation object of the builder.
func (bmu *BaselineMeasureUpdate) Mutation() *BaselineMeasureMutation {
	return bmu.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bmu *BaselineMeasureUpdate) ClearParent() *BaselineMeasureUpdate {
	bmu.mutation.ClearParent()
	return bmu
}

// ClearBaselineMeasureDenomList clears all "baseline_measure_denom_list" edges to the BaselineMeasureDenom entity.
func (bmu *BaselineMeasureUpdate) ClearBaselineMeasureDenomList() *BaselineMeasureUpdate {
	bmu.mutation.ClearBaselineMeasureDenomList()
	return bmu
}

// RemoveBaselineMeasureDenomListIDs removes the "baseline_measure_denom_list" edge to BaselineMeasureDenom entities by IDs.
func (bmu *BaselineMeasureUpdate) RemoveBaselineMeasureDenomListIDs(ids ...int) *BaselineMeasureUpdate {
	bmu.mutation.RemoveBaselineMeasureDenomListIDs(ids...)
	return bmu
}

// RemoveBaselineMeasureDenomList removes "baseline_measure_denom_list" edges to BaselineMeasureDenom entities.
func (bmu *BaselineMeasureUpdate) RemoveBaselineMeasureDenomList(b ...*BaselineMeasureDenom) *BaselineMeasureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmu.RemoveBaselineMeasureDenomListIDs(ids...)
}

// ClearBaselineClassList clears all "baseline_class_list" edges to the BaselineClass entity.
func (bmu *BaselineMeasureUpdate) ClearBaselineClassList() *BaselineMeasureUpdate {
	bmu.mutation.ClearBaselineClassList()
	return bmu
}

// RemoveBaselineClassListIDs removes the "baseline_class_list" edge to BaselineClass entities by IDs.
func (bmu *BaselineMeasureUpdate) RemoveBaselineClassListIDs(ids ...int) *BaselineMeasureUpdate {
	bmu.mutation.RemoveBaselineClassListIDs(ids...)
	return bmu
}

// RemoveBaselineClassList removes "baseline_class_list" edges to BaselineClass entities.
func (bmu *BaselineMeasureUpdate) RemoveBaselineClassList(b ...*BaselineClass) *BaselineMeasureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmu.RemoveBaselineClassListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bmu *BaselineMeasureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmu.hooks) == 0 {
		affected, err = bmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmu.mutation = mutation
			affected, err = bmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmu.hooks) - 1; i >= 0; i-- {
			if bmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmu *BaselineMeasureUpdate) SaveX(ctx context.Context) int {
	affected, err := bmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bmu *BaselineMeasureUpdate) Exec(ctx context.Context) error {
	_, err := bmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmu *BaselineMeasureUpdate) ExecX(ctx context.Context) {
	if err := bmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmu *BaselineMeasureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasure.Table,
			Columns: baselinemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasure.FieldID,
			},
		},
	}
	if ps := bmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmu.mutation.BaselineMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureTitle,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDescription,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasurePopulationDescription,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureParamType,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDispersionType,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureUnitOfMeasure(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureUnitOfMeasure,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureCalculatePct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureCalculatePct,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasureDenomUnitsSelected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDenomUnitsSelected,
		})
	}
	if bmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasure.ParentTable,
			Columns: []string{baselinemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasure.ParentTable,
			Columns: []string{baselinemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmu.mutation.BaselineMeasureDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.RemovedBaselineMeasureDenomListIDs(); len(nodes) > 0 && !bmu.mutation.BaselineMeasureDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.BaselineMeasureDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmu.mutation.BaselineClassListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.RemovedBaselineClassListIDs(); len(nodes) > 0 && !bmu.mutation.BaselineClassListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.BaselineClassListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineMeasureUpdateOne is the builder for updating a single BaselineMeasure entity.
type BaselineMeasureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineMeasureMutation
}

// SetBaselineMeasureTitle sets the "baseline_measure_title" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureTitle(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureTitle(s)
	return bmuo
}

// SetBaselineMeasureDescription sets the "baseline_measure_description" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureDescription(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureDescription(s)
	return bmuo
}

// SetBaselineMeasurePopulationDescription sets the "baseline_measure_population_description" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasurePopulationDescription(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasurePopulationDescription(s)
	return bmuo
}

// SetBaselineMeasureParamType sets the "baseline_measure_param_type" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureParamType(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureParamType(s)
	return bmuo
}

// SetBaselineMeasureDispersionType sets the "baseline_measure_dispersion_type" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureDispersionType(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureDispersionType(s)
	return bmuo
}

// SetBaselineMeasureUnitOfMeasure sets the "baseline_measure_unit_of_measure" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureUnitOfMeasure(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureUnitOfMeasure(s)
	return bmuo
}

// SetBaselineMeasureCalculatePct sets the "baseline_measure_calculate_pct" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureCalculatePct(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureCalculatePct(s)
	return bmuo
}

// SetBaselineMeasureDenomUnitsSelected sets the "baseline_measure_denom_units_selected" field.
func (bmuo *BaselineMeasureUpdateOne) SetBaselineMeasureDenomUnitsSelected(s string) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetBaselineMeasureDenomUnitsSelected(s)
	return bmuo
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bmuo *BaselineMeasureUpdateOne) SetParentID(id int) *BaselineMeasureUpdateOne {
	bmuo.mutation.SetParentID(id)
	return bmuo
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bmuo *BaselineMeasureUpdateOne) SetNillableParentID(id *int) *BaselineMeasureUpdateOne {
	if id != nil {
		bmuo = bmuo.SetParentID(*id)
	}
	return bmuo
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bmuo *BaselineMeasureUpdateOne) SetParent(b *BaselineCharacteristicsModule) *BaselineMeasureUpdateOne {
	return bmuo.SetParentID(b.ID)
}

// AddBaselineMeasureDenomListIDs adds the "baseline_measure_denom_list" edge to the BaselineMeasureDenom entity by IDs.
func (bmuo *BaselineMeasureUpdateOne) AddBaselineMeasureDenomListIDs(ids ...int) *BaselineMeasureUpdateOne {
	bmuo.mutation.AddBaselineMeasureDenomListIDs(ids...)
	return bmuo
}

// AddBaselineMeasureDenomList adds the "baseline_measure_denom_list" edges to the BaselineMeasureDenom entity.
func (bmuo *BaselineMeasureUpdateOne) AddBaselineMeasureDenomList(b ...*BaselineMeasureDenom) *BaselineMeasureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmuo.AddBaselineMeasureDenomListIDs(ids...)
}

// AddBaselineClassListIDs adds the "baseline_class_list" edge to the BaselineClass entity by IDs.
func (bmuo *BaselineMeasureUpdateOne) AddBaselineClassListIDs(ids ...int) *BaselineMeasureUpdateOne {
	bmuo.mutation.AddBaselineClassListIDs(ids...)
	return bmuo
}

// AddBaselineClassList adds the "baseline_class_list" edges to the BaselineClass entity.
func (bmuo *BaselineMeasureUpdateOne) AddBaselineClassList(b ...*BaselineClass) *BaselineMeasureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmuo.AddBaselineClassListIDs(ids...)
}

// Mutation returns the BaselineMeasureMutation object of the builder.
func (bmuo *BaselineMeasureUpdateOne) Mutation() *BaselineMeasureMutation {
	return bmuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bmuo *BaselineMeasureUpdateOne) ClearParent() *BaselineMeasureUpdateOne {
	bmuo.mutation.ClearParent()
	return bmuo
}

// ClearBaselineMeasureDenomList clears all "baseline_measure_denom_list" edges to the BaselineMeasureDenom entity.
func (bmuo *BaselineMeasureUpdateOne) ClearBaselineMeasureDenomList() *BaselineMeasureUpdateOne {
	bmuo.mutation.ClearBaselineMeasureDenomList()
	return bmuo
}

// RemoveBaselineMeasureDenomListIDs removes the "baseline_measure_denom_list" edge to BaselineMeasureDenom entities by IDs.
func (bmuo *BaselineMeasureUpdateOne) RemoveBaselineMeasureDenomListIDs(ids ...int) *BaselineMeasureUpdateOne {
	bmuo.mutation.RemoveBaselineMeasureDenomListIDs(ids...)
	return bmuo
}

// RemoveBaselineMeasureDenomList removes "baseline_measure_denom_list" edges to BaselineMeasureDenom entities.
func (bmuo *BaselineMeasureUpdateOne) RemoveBaselineMeasureDenomList(b ...*BaselineMeasureDenom) *BaselineMeasureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmuo.RemoveBaselineMeasureDenomListIDs(ids...)
}

// ClearBaselineClassList clears all "baseline_class_list" edges to the BaselineClass entity.
func (bmuo *BaselineMeasureUpdateOne) ClearBaselineClassList() *BaselineMeasureUpdateOne {
	bmuo.mutation.ClearBaselineClassList()
	return bmuo
}

// RemoveBaselineClassListIDs removes the "baseline_class_list" edge to BaselineClass entities by IDs.
func (bmuo *BaselineMeasureUpdateOne) RemoveBaselineClassListIDs(ids ...int) *BaselineMeasureUpdateOne {
	bmuo.mutation.RemoveBaselineClassListIDs(ids...)
	return bmuo
}

// RemoveBaselineClassList removes "baseline_class_list" edges to BaselineClass entities.
func (bmuo *BaselineMeasureUpdateOne) RemoveBaselineClassList(b ...*BaselineClass) *BaselineMeasureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmuo.RemoveBaselineClassListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bmuo *BaselineMeasureUpdateOne) Select(field string, fields ...string) *BaselineMeasureUpdateOne {
	bmuo.fields = append([]string{field}, fields...)
	return bmuo
}

// Save executes the query and returns the updated BaselineMeasure entity.
func (bmuo *BaselineMeasureUpdateOne) Save(ctx context.Context) (*BaselineMeasure, error) {
	var (
		err  error
		node *BaselineMeasure
	)
	if len(bmuo.hooks) == 0 {
		node, err = bmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmuo.mutation = mutation
			node, err = bmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bmuo.hooks) - 1; i >= 0; i-- {
			if bmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmuo *BaselineMeasureUpdateOne) SaveX(ctx context.Context) *BaselineMeasure {
	node, err := bmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bmuo *BaselineMeasureUpdateOne) Exec(ctx context.Context) error {
	_, err := bmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmuo *BaselineMeasureUpdateOne) ExecX(ctx context.Context) {
	if err := bmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmuo *BaselineMeasureUpdateOne) sqlSave(ctx context.Context) (_node *BaselineMeasure, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasure.Table,
			Columns: baselinemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasure.FieldID,
			},
		},
	}
	id, ok := bmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineMeasure.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasure.FieldID)
		for _, f := range fields {
			if !baselinemeasure.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinemeasure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmuo.mutation.BaselineMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureTitle,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDescription,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasurePopulationDescription,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureParamType,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDispersionType,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureUnitOfMeasure(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureUnitOfMeasure,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureCalculatePct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureCalculatePct,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasureDenomUnitsSelected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDenomUnitsSelected,
		})
	}
	if bmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasure.ParentTable,
			Columns: []string{baselinemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasure.ParentTable,
			Columns: []string{baselinemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmuo.mutation.BaselineMeasureDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.RemovedBaselineMeasureDenomListIDs(); len(nodes) > 0 && !bmuo.mutation.BaselineMeasureDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.BaselineMeasureDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmuo.mutation.BaselineClassListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.RemovedBaselineClassListIDs(); len(nodes) > 0 && !bmuo.mutation.BaselineClassListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.BaselineClassListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineMeasure{config: bmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
