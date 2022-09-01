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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// BaselineCharacteristicsModuleUpdate is the builder for updating BaselineCharacteristicsModule entities.
type BaselineCharacteristicsModuleUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineCharacteristicsModuleMutation
}

// Where appends a list predicates to the BaselineCharacteristicsModuleUpdate builder.
func (bcmu *BaselineCharacteristicsModuleUpdate) Where(ps ...predicate.BaselineCharacteristicsModule) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.Where(ps...)
	return bcmu
}

// SetBaselinePopulationDescription sets the "baseline_population_description" field.
func (bcmu *BaselineCharacteristicsModuleUpdate) SetBaselinePopulationDescription(s string) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.SetBaselinePopulationDescription(s)
	return bcmu
}

// SetBaselineTypeUnitsAnalyzed sets the "baseline_type_units_analyzed" field.
func (bcmu *BaselineCharacteristicsModuleUpdate) SetBaselineTypeUnitsAnalyzed(s string) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.SetBaselineTypeUnitsAnalyzed(s)
	return bcmu
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (bcmu *BaselineCharacteristicsModuleUpdate) SetParentID(id int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.SetParentID(id)
	return bcmu
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (bcmu *BaselineCharacteristicsModuleUpdate) SetNillableParentID(id *int) *BaselineCharacteristicsModuleUpdate {
	if id != nil {
		bcmu = bcmu.SetParentID(*id)
	}
	return bcmu
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) SetParent(r *ResultsDefinition) *BaselineCharacteristicsModuleUpdate {
	return bcmu.SetParentID(r.ID)
}

// AddBaselineGroupListIDs adds the "baseline_group_list" edge to the BaselineGroup entity by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineGroupListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.AddBaselineGroupListIDs(ids...)
	return bcmu
}

// AddBaselineGroupList adds the "baseline_group_list" edges to the BaselineGroup entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineGroupList(b ...*BaselineGroup) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.AddBaselineGroupListIDs(ids...)
}

// AddBaselineDenomListIDs adds the "baseline_denom_list" edge to the BaselineDenom entity by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineDenomListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.AddBaselineDenomListIDs(ids...)
	return bcmu
}

// AddBaselineDenomList adds the "baseline_denom_list" edges to the BaselineDenom entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineDenomList(b ...*BaselineDenom) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.AddBaselineDenomListIDs(ids...)
}

// AddBaselineMeasureListIDs adds the "baseline_measure_list" edge to the BaselineMeasure entity by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineMeasureListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.AddBaselineMeasureListIDs(ids...)
	return bcmu
}

// AddBaselineMeasureList adds the "baseline_measure_list" edges to the BaselineMeasure entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) AddBaselineMeasureList(b ...*BaselineMeasure) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.AddBaselineMeasureListIDs(ids...)
}

// Mutation returns the BaselineCharacteristicsModuleMutation object of the builder.
func (bcmu *BaselineCharacteristicsModuleUpdate) Mutation() *BaselineCharacteristicsModuleMutation {
	return bcmu.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) ClearParent() *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.ClearParent()
	return bcmu
}

// ClearBaselineGroupList clears all "baseline_group_list" edges to the BaselineGroup entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) ClearBaselineGroupList() *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.ClearBaselineGroupList()
	return bcmu
}

// RemoveBaselineGroupListIDs removes the "baseline_group_list" edge to BaselineGroup entities by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineGroupListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.RemoveBaselineGroupListIDs(ids...)
	return bcmu
}

// RemoveBaselineGroupList removes "baseline_group_list" edges to BaselineGroup entities.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineGroupList(b ...*BaselineGroup) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.RemoveBaselineGroupListIDs(ids...)
}

// ClearBaselineDenomList clears all "baseline_denom_list" edges to the BaselineDenom entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) ClearBaselineDenomList() *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.ClearBaselineDenomList()
	return bcmu
}

// RemoveBaselineDenomListIDs removes the "baseline_denom_list" edge to BaselineDenom entities by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineDenomListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.RemoveBaselineDenomListIDs(ids...)
	return bcmu
}

// RemoveBaselineDenomList removes "baseline_denom_list" edges to BaselineDenom entities.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineDenomList(b ...*BaselineDenom) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.RemoveBaselineDenomListIDs(ids...)
}

// ClearBaselineMeasureList clears all "baseline_measure_list" edges to the BaselineMeasure entity.
func (bcmu *BaselineCharacteristicsModuleUpdate) ClearBaselineMeasureList() *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.ClearBaselineMeasureList()
	return bcmu
}

// RemoveBaselineMeasureListIDs removes the "baseline_measure_list" edge to BaselineMeasure entities by IDs.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineMeasureListIDs(ids ...int) *BaselineCharacteristicsModuleUpdate {
	bcmu.mutation.RemoveBaselineMeasureListIDs(ids...)
	return bcmu
}

// RemoveBaselineMeasureList removes "baseline_measure_list" edges to BaselineMeasure entities.
func (bcmu *BaselineCharacteristicsModuleUpdate) RemoveBaselineMeasureList(b ...*BaselineMeasure) *BaselineCharacteristicsModuleUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmu.RemoveBaselineMeasureListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcmu *BaselineCharacteristicsModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcmu.hooks) == 0 {
		affected, err = bcmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCharacteristicsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcmu.mutation = mutation
			affected, err = bcmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcmu.hooks) - 1; i >= 0; i-- {
			if bcmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcmu *BaselineCharacteristicsModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := bcmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcmu *BaselineCharacteristicsModuleUpdate) Exec(ctx context.Context) error {
	_, err := bcmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmu *BaselineCharacteristicsModuleUpdate) ExecX(ctx context.Context) {
	if err := bcmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcmu *BaselineCharacteristicsModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecharacteristicsmodule.Table,
			Columns: baselinecharacteristicsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecharacteristicsmodule.FieldID,
			},
		},
	}
	if ps := bcmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcmu.mutation.BaselinePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselinePopulationDescription,
		})
	}
	if value, ok := bcmu.mutation.BaselineTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselineTypeUnitsAnalyzed,
		})
	}
	if bcmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   baselinecharacteristicsmodule.ParentTable,
			Columns: []string{baselinecharacteristicsmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   baselinecharacteristicsmodule.ParentTable,
			Columns: []string{baselinecharacteristicsmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmu.mutation.BaselineGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.RemovedBaselineGroupListIDs(); len(nodes) > 0 && !bcmu.mutation.BaselineGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.BaselineGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmu.mutation.BaselineDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.RemovedBaselineDenomListIDs(); len(nodes) > 0 && !bcmu.mutation.BaselineDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.BaselineDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmu.mutation.BaselineMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.RemovedBaselineMeasureListIDs(); len(nodes) > 0 && !bcmu.mutation.BaselineMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmu.mutation.BaselineMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinecharacteristicsmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineCharacteristicsModuleUpdateOne is the builder for updating a single BaselineCharacteristicsModule entity.
type BaselineCharacteristicsModuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineCharacteristicsModuleMutation
}

// SetBaselinePopulationDescription sets the "baseline_population_description" field.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SetBaselinePopulationDescription(s string) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.SetBaselinePopulationDescription(s)
	return bcmuo
}

// SetBaselineTypeUnitsAnalyzed sets the "baseline_type_units_analyzed" field.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SetBaselineTypeUnitsAnalyzed(s string) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.SetBaselineTypeUnitsAnalyzed(s)
	return bcmuo
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SetParentID(id int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.SetParentID(id)
	return bcmuo
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SetNillableParentID(id *int) *BaselineCharacteristicsModuleUpdateOne {
	if id != nil {
		bcmuo = bcmuo.SetParentID(*id)
	}
	return bcmuo
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SetParent(r *ResultsDefinition) *BaselineCharacteristicsModuleUpdateOne {
	return bcmuo.SetParentID(r.ID)
}

// AddBaselineGroupListIDs adds the "baseline_group_list" edge to the BaselineGroup entity by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineGroupListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.AddBaselineGroupListIDs(ids...)
	return bcmuo
}

// AddBaselineGroupList adds the "baseline_group_list" edges to the BaselineGroup entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineGroupList(b ...*BaselineGroup) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.AddBaselineGroupListIDs(ids...)
}

// AddBaselineDenomListIDs adds the "baseline_denom_list" edge to the BaselineDenom entity by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineDenomListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.AddBaselineDenomListIDs(ids...)
	return bcmuo
}

// AddBaselineDenomList adds the "baseline_denom_list" edges to the BaselineDenom entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineDenomList(b ...*BaselineDenom) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.AddBaselineDenomListIDs(ids...)
}

// AddBaselineMeasureListIDs adds the "baseline_measure_list" edge to the BaselineMeasure entity by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineMeasureListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.AddBaselineMeasureListIDs(ids...)
	return bcmuo
}

// AddBaselineMeasureList adds the "baseline_measure_list" edges to the BaselineMeasure entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) AddBaselineMeasureList(b ...*BaselineMeasure) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.AddBaselineMeasureListIDs(ids...)
}

// Mutation returns the BaselineCharacteristicsModuleMutation object of the builder.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) Mutation() *BaselineCharacteristicsModuleMutation {
	return bcmuo.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) ClearParent() *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.ClearParent()
	return bcmuo
}

// ClearBaselineGroupList clears all "baseline_group_list" edges to the BaselineGroup entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) ClearBaselineGroupList() *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.ClearBaselineGroupList()
	return bcmuo
}

// RemoveBaselineGroupListIDs removes the "baseline_group_list" edge to BaselineGroup entities by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineGroupListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.RemoveBaselineGroupListIDs(ids...)
	return bcmuo
}

// RemoveBaselineGroupList removes "baseline_group_list" edges to BaselineGroup entities.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineGroupList(b ...*BaselineGroup) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.RemoveBaselineGroupListIDs(ids...)
}

// ClearBaselineDenomList clears all "baseline_denom_list" edges to the BaselineDenom entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) ClearBaselineDenomList() *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.ClearBaselineDenomList()
	return bcmuo
}

// RemoveBaselineDenomListIDs removes the "baseline_denom_list" edge to BaselineDenom entities by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineDenomListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.RemoveBaselineDenomListIDs(ids...)
	return bcmuo
}

// RemoveBaselineDenomList removes "baseline_denom_list" edges to BaselineDenom entities.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineDenomList(b ...*BaselineDenom) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.RemoveBaselineDenomListIDs(ids...)
}

// ClearBaselineMeasureList clears all "baseline_measure_list" edges to the BaselineMeasure entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) ClearBaselineMeasureList() *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.ClearBaselineMeasureList()
	return bcmuo
}

// RemoveBaselineMeasureListIDs removes the "baseline_measure_list" edge to BaselineMeasure entities by IDs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineMeasureListIDs(ids ...int) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.mutation.RemoveBaselineMeasureListIDs(ids...)
	return bcmuo
}

// RemoveBaselineMeasureList removes "baseline_measure_list" edges to BaselineMeasure entities.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) RemoveBaselineMeasureList(b ...*BaselineMeasure) *BaselineCharacteristicsModuleUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmuo.RemoveBaselineMeasureListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) Select(field string, fields ...string) *BaselineCharacteristicsModuleUpdateOne {
	bcmuo.fields = append([]string{field}, fields...)
	return bcmuo
}

// Save executes the query and returns the updated BaselineCharacteristicsModule entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) Save(ctx context.Context) (*BaselineCharacteristicsModule, error) {
	var (
		err  error
		node *BaselineCharacteristicsModule
	)
	if len(bcmuo.hooks) == 0 {
		node, err = bcmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCharacteristicsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcmuo.mutation = mutation
			node, err = bcmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcmuo.hooks) - 1; i >= 0; i-- {
			if bcmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) SaveX(ctx context.Context) *BaselineCharacteristicsModule {
	node, err := bcmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := bcmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmuo *BaselineCharacteristicsModuleUpdateOne) ExecX(ctx context.Context) {
	if err := bcmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcmuo *BaselineCharacteristicsModuleUpdateOne) sqlSave(ctx context.Context) (_node *BaselineCharacteristicsModule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecharacteristicsmodule.Table,
			Columns: baselinecharacteristicsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecharacteristicsmodule.FieldID,
			},
		},
	}
	id, ok := bcmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineCharacteristicsModule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecharacteristicsmodule.FieldID)
		for _, f := range fields {
			if !baselinecharacteristicsmodule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinecharacteristicsmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcmuo.mutation.BaselinePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselinePopulationDescription,
		})
	}
	if value, ok := bcmuo.mutation.BaselineTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselineTypeUnitsAnalyzed,
		})
	}
	if bcmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   baselinecharacteristicsmodule.ParentTable,
			Columns: []string{baselinecharacteristicsmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   baselinecharacteristicsmodule.ParentTable,
			Columns: []string{baselinecharacteristicsmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmuo.mutation.BaselineGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.RemovedBaselineGroupListIDs(); len(nodes) > 0 && !bcmuo.mutation.BaselineGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.BaselineGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmuo.mutation.BaselineDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.RemovedBaselineDenomListIDs(); len(nodes) > 0 && !bcmuo.mutation.BaselineDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.BaselineDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcmuo.mutation.BaselineMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.RemovedBaselineMeasureListIDs(); len(nodes) > 0 && !bcmuo.mutation.BaselineMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcmuo.mutation.BaselineMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineCharacteristicsModule{config: bcmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinecharacteristicsmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
