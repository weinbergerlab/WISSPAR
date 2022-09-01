// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowMilestoneUpdate is the builder for updating FlowMilestone entities.
type FlowMilestoneUpdate struct {
	config
	hooks    []Hook
	mutation *FlowMilestoneMutation
}

// Where appends a list predicates to the FlowMilestoneUpdate builder.
func (fmu *FlowMilestoneUpdate) Where(ps ...predicate.FlowMilestone) *FlowMilestoneUpdate {
	fmu.mutation.Where(ps...)
	return fmu
}

// SetFlowMilestoneType sets the "flow_milestone_type" field.
func (fmu *FlowMilestoneUpdate) SetFlowMilestoneType(s string) *FlowMilestoneUpdate {
	fmu.mutation.SetFlowMilestoneType(s)
	return fmu
}

// SetFlowMilestoneComment sets the "flow_milestone_comment" field.
func (fmu *FlowMilestoneUpdate) SetFlowMilestoneComment(s string) *FlowMilestoneUpdate {
	fmu.mutation.SetFlowMilestoneComment(s)
	return fmu
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fmu *FlowMilestoneUpdate) SetParentID(id int) *FlowMilestoneUpdate {
	fmu.mutation.SetParentID(id)
	return fmu
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fmu *FlowMilestoneUpdate) SetNillableParentID(id *int) *FlowMilestoneUpdate {
	if id != nil {
		fmu = fmu.SetParentID(*id)
	}
	return fmu
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fmu *FlowMilestoneUpdate) SetParent(f *FlowPeriod) *FlowMilestoneUpdate {
	return fmu.SetParentID(f.ID)
}

// AddFlowAchievementListIDs adds the "flow_achievement_list" edge to the FlowAchievement entity by IDs.
func (fmu *FlowMilestoneUpdate) AddFlowAchievementListIDs(ids ...int) *FlowMilestoneUpdate {
	fmu.mutation.AddFlowAchievementListIDs(ids...)
	return fmu
}

// AddFlowAchievementList adds the "flow_achievement_list" edges to the FlowAchievement entity.
func (fmu *FlowMilestoneUpdate) AddFlowAchievementList(f ...*FlowAchievement) *FlowMilestoneUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmu.AddFlowAchievementListIDs(ids...)
}

// Mutation returns the FlowMilestoneMutation object of the builder.
func (fmu *FlowMilestoneUpdate) Mutation() *FlowMilestoneMutation {
	return fmu.mutation
}

// ClearParent clears the "parent" edge to the FlowPeriod entity.
func (fmu *FlowMilestoneUpdate) ClearParent() *FlowMilestoneUpdate {
	fmu.mutation.ClearParent()
	return fmu
}

// ClearFlowAchievementList clears all "flow_achievement_list" edges to the FlowAchievement entity.
func (fmu *FlowMilestoneUpdate) ClearFlowAchievementList() *FlowMilestoneUpdate {
	fmu.mutation.ClearFlowAchievementList()
	return fmu
}

// RemoveFlowAchievementListIDs removes the "flow_achievement_list" edge to FlowAchievement entities by IDs.
func (fmu *FlowMilestoneUpdate) RemoveFlowAchievementListIDs(ids ...int) *FlowMilestoneUpdate {
	fmu.mutation.RemoveFlowAchievementListIDs(ids...)
	return fmu
}

// RemoveFlowAchievementList removes "flow_achievement_list" edges to FlowAchievement entities.
func (fmu *FlowMilestoneUpdate) RemoveFlowAchievementList(f ...*FlowAchievement) *FlowMilestoneUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmu.RemoveFlowAchievementListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fmu *FlowMilestoneUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fmu.hooks) == 0 {
		affected, err = fmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowMilestoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fmu.mutation = mutation
			affected, err = fmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fmu.hooks) - 1; i >= 0; i-- {
			if fmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fmu *FlowMilestoneUpdate) SaveX(ctx context.Context) int {
	affected, err := fmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fmu *FlowMilestoneUpdate) Exec(ctx context.Context) error {
	_, err := fmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmu *FlowMilestoneUpdate) ExecX(ctx context.Context) {
	if err := fmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fmu *FlowMilestoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowmilestone.Table,
			Columns: flowmilestone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowmilestone.FieldID,
			},
		},
	}
	if ps := fmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fmu.mutation.FlowMilestoneType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneType,
		})
	}
	if value, ok := fmu.mutation.FlowMilestoneComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneComment,
		})
	}
	if fmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowmilestone.ParentTable,
			Columns: []string{flowmilestone.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowmilestone.ParentTable,
			Columns: []string{flowmilestone.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fmu.mutation.FlowAchievementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmu.mutation.RemovedFlowAchievementListIDs(); len(nodes) > 0 && !fmu.mutation.FlowAchievementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmu.mutation.FlowAchievementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowmilestone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowMilestoneUpdateOne is the builder for updating a single FlowMilestone entity.
type FlowMilestoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowMilestoneMutation
}

// SetFlowMilestoneType sets the "flow_milestone_type" field.
func (fmuo *FlowMilestoneUpdateOne) SetFlowMilestoneType(s string) *FlowMilestoneUpdateOne {
	fmuo.mutation.SetFlowMilestoneType(s)
	return fmuo
}

// SetFlowMilestoneComment sets the "flow_milestone_comment" field.
func (fmuo *FlowMilestoneUpdateOne) SetFlowMilestoneComment(s string) *FlowMilestoneUpdateOne {
	fmuo.mutation.SetFlowMilestoneComment(s)
	return fmuo
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fmuo *FlowMilestoneUpdateOne) SetParentID(id int) *FlowMilestoneUpdateOne {
	fmuo.mutation.SetParentID(id)
	return fmuo
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fmuo *FlowMilestoneUpdateOne) SetNillableParentID(id *int) *FlowMilestoneUpdateOne {
	if id != nil {
		fmuo = fmuo.SetParentID(*id)
	}
	return fmuo
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fmuo *FlowMilestoneUpdateOne) SetParent(f *FlowPeriod) *FlowMilestoneUpdateOne {
	return fmuo.SetParentID(f.ID)
}

// AddFlowAchievementListIDs adds the "flow_achievement_list" edge to the FlowAchievement entity by IDs.
func (fmuo *FlowMilestoneUpdateOne) AddFlowAchievementListIDs(ids ...int) *FlowMilestoneUpdateOne {
	fmuo.mutation.AddFlowAchievementListIDs(ids...)
	return fmuo
}

// AddFlowAchievementList adds the "flow_achievement_list" edges to the FlowAchievement entity.
func (fmuo *FlowMilestoneUpdateOne) AddFlowAchievementList(f ...*FlowAchievement) *FlowMilestoneUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmuo.AddFlowAchievementListIDs(ids...)
}

// Mutation returns the FlowMilestoneMutation object of the builder.
func (fmuo *FlowMilestoneUpdateOne) Mutation() *FlowMilestoneMutation {
	return fmuo.mutation
}

// ClearParent clears the "parent" edge to the FlowPeriod entity.
func (fmuo *FlowMilestoneUpdateOne) ClearParent() *FlowMilestoneUpdateOne {
	fmuo.mutation.ClearParent()
	return fmuo
}

// ClearFlowAchievementList clears all "flow_achievement_list" edges to the FlowAchievement entity.
func (fmuo *FlowMilestoneUpdateOne) ClearFlowAchievementList() *FlowMilestoneUpdateOne {
	fmuo.mutation.ClearFlowAchievementList()
	return fmuo
}

// RemoveFlowAchievementListIDs removes the "flow_achievement_list" edge to FlowAchievement entities by IDs.
func (fmuo *FlowMilestoneUpdateOne) RemoveFlowAchievementListIDs(ids ...int) *FlowMilestoneUpdateOne {
	fmuo.mutation.RemoveFlowAchievementListIDs(ids...)
	return fmuo
}

// RemoveFlowAchievementList removes "flow_achievement_list" edges to FlowAchievement entities.
func (fmuo *FlowMilestoneUpdateOne) RemoveFlowAchievementList(f ...*FlowAchievement) *FlowMilestoneUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmuo.RemoveFlowAchievementListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fmuo *FlowMilestoneUpdateOne) Select(field string, fields ...string) *FlowMilestoneUpdateOne {
	fmuo.fields = append([]string{field}, fields...)
	return fmuo
}

// Save executes the query and returns the updated FlowMilestone entity.
func (fmuo *FlowMilestoneUpdateOne) Save(ctx context.Context) (*FlowMilestone, error) {
	var (
		err  error
		node *FlowMilestone
	)
	if len(fmuo.hooks) == 0 {
		node, err = fmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowMilestoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fmuo.mutation = mutation
			node, err = fmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fmuo.hooks) - 1; i >= 0; i-- {
			if fmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fmuo *FlowMilestoneUpdateOne) SaveX(ctx context.Context) *FlowMilestone {
	node, err := fmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fmuo *FlowMilestoneUpdateOne) Exec(ctx context.Context) error {
	_, err := fmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmuo *FlowMilestoneUpdateOne) ExecX(ctx context.Context) {
	if err := fmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fmuo *FlowMilestoneUpdateOne) sqlSave(ctx context.Context) (_node *FlowMilestone, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowmilestone.Table,
			Columns: flowmilestone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowmilestone.FieldID,
			},
		},
	}
	id, ok := fmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowMilestone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowmilestone.FieldID)
		for _, f := range fields {
			if !flowmilestone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowmilestone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fmuo.mutation.FlowMilestoneType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneType,
		})
	}
	if value, ok := fmuo.mutation.FlowMilestoneComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneComment,
		})
	}
	if fmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowmilestone.ParentTable,
			Columns: []string{flowmilestone.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowmilestone.ParentTable,
			Columns: []string{flowmilestone.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fmuo.mutation.FlowAchievementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmuo.mutation.RemovedFlowAchievementListIDs(); len(nodes) > 0 && !fmuo.mutation.FlowAchievementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fmuo.mutation.FlowAchievementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowmilestone.FlowAchievementListTable,
			Columns: []string{flowmilestone.FlowAchievementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowachievement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowMilestone{config: fmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowmilestone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
