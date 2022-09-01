// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowPeriodUpdate is the builder for updating FlowPeriod entities.
type FlowPeriodUpdate struct {
	config
	hooks    []Hook
	mutation *FlowPeriodMutation
}

// Where appends a list predicates to the FlowPeriodUpdate builder.
func (fpu *FlowPeriodUpdate) Where(ps ...predicate.FlowPeriod) *FlowPeriodUpdate {
	fpu.mutation.Where(ps...)
	return fpu
}

// SetFlowPeriodTitle sets the "flow_period_title" field.
func (fpu *FlowPeriodUpdate) SetFlowPeriodTitle(s string) *FlowPeriodUpdate {
	fpu.mutation.SetFlowPeriodTitle(s)
	return fpu
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fpu *FlowPeriodUpdate) SetParentID(id int) *FlowPeriodUpdate {
	fpu.mutation.SetParentID(id)
	return fpu
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fpu *FlowPeriodUpdate) SetNillableParentID(id *int) *FlowPeriodUpdate {
	if id != nil {
		fpu = fpu.SetParentID(*id)
	}
	return fpu
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fpu *FlowPeriodUpdate) SetParent(p *ParticipantFlowModule) *FlowPeriodUpdate {
	return fpu.SetParentID(p.ID)
}

// AddFlowMilestoneListIDs adds the "flow_milestone_list" edge to the FlowMilestone entity by IDs.
func (fpu *FlowPeriodUpdate) AddFlowMilestoneListIDs(ids ...int) *FlowPeriodUpdate {
	fpu.mutation.AddFlowMilestoneListIDs(ids...)
	return fpu
}

// AddFlowMilestoneList adds the "flow_milestone_list" edges to the FlowMilestone entity.
func (fpu *FlowPeriodUpdate) AddFlowMilestoneList(f ...*FlowMilestone) *FlowPeriodUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.AddFlowMilestoneListIDs(ids...)
}

// AddFlowDropWithdrawListIDs adds the "flow_drop_withdraw_list" edge to the FlowDropWithdraw entity by IDs.
func (fpu *FlowPeriodUpdate) AddFlowDropWithdrawListIDs(ids ...int) *FlowPeriodUpdate {
	fpu.mutation.AddFlowDropWithdrawListIDs(ids...)
	return fpu
}

// AddFlowDropWithdrawList adds the "flow_drop_withdraw_list" edges to the FlowDropWithdraw entity.
func (fpu *FlowPeriodUpdate) AddFlowDropWithdrawList(f ...*FlowDropWithdraw) *FlowPeriodUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.AddFlowDropWithdrawListIDs(ids...)
}

// Mutation returns the FlowPeriodMutation object of the builder.
func (fpu *FlowPeriodUpdate) Mutation() *FlowPeriodMutation {
	return fpu.mutation
}

// ClearParent clears the "parent" edge to the ParticipantFlowModule entity.
func (fpu *FlowPeriodUpdate) ClearParent() *FlowPeriodUpdate {
	fpu.mutation.ClearParent()
	return fpu
}

// ClearFlowMilestoneList clears all "flow_milestone_list" edges to the FlowMilestone entity.
func (fpu *FlowPeriodUpdate) ClearFlowMilestoneList() *FlowPeriodUpdate {
	fpu.mutation.ClearFlowMilestoneList()
	return fpu
}

// RemoveFlowMilestoneListIDs removes the "flow_milestone_list" edge to FlowMilestone entities by IDs.
func (fpu *FlowPeriodUpdate) RemoveFlowMilestoneListIDs(ids ...int) *FlowPeriodUpdate {
	fpu.mutation.RemoveFlowMilestoneListIDs(ids...)
	return fpu
}

// RemoveFlowMilestoneList removes "flow_milestone_list" edges to FlowMilestone entities.
func (fpu *FlowPeriodUpdate) RemoveFlowMilestoneList(f ...*FlowMilestone) *FlowPeriodUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.RemoveFlowMilestoneListIDs(ids...)
}

// ClearFlowDropWithdrawList clears all "flow_drop_withdraw_list" edges to the FlowDropWithdraw entity.
func (fpu *FlowPeriodUpdate) ClearFlowDropWithdrawList() *FlowPeriodUpdate {
	fpu.mutation.ClearFlowDropWithdrawList()
	return fpu
}

// RemoveFlowDropWithdrawListIDs removes the "flow_drop_withdraw_list" edge to FlowDropWithdraw entities by IDs.
func (fpu *FlowPeriodUpdate) RemoveFlowDropWithdrawListIDs(ids ...int) *FlowPeriodUpdate {
	fpu.mutation.RemoveFlowDropWithdrawListIDs(ids...)
	return fpu
}

// RemoveFlowDropWithdrawList removes "flow_drop_withdraw_list" edges to FlowDropWithdraw entities.
func (fpu *FlowPeriodUpdate) RemoveFlowDropWithdrawList(f ...*FlowDropWithdraw) *FlowPeriodUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.RemoveFlowDropWithdrawListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fpu *FlowPeriodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fpu.hooks) == 0 {
		affected, err = fpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowPeriodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fpu.mutation = mutation
			affected, err = fpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fpu.hooks) - 1; i >= 0; i-- {
			if fpu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fpu *FlowPeriodUpdate) SaveX(ctx context.Context) int {
	affected, err := fpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fpu *FlowPeriodUpdate) Exec(ctx context.Context) error {
	_, err := fpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpu *FlowPeriodUpdate) ExecX(ctx context.Context) {
	if err := fpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fpu *FlowPeriodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowperiod.Table,
			Columns: flowperiod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowperiod.FieldID,
			},
		},
	}
	if ps := fpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fpu.mutation.FlowPeriodTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowperiod.FieldFlowPeriodTitle,
		})
	}
	if fpu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowperiod.ParentTable,
			Columns: []string{flowperiod.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowperiod.ParentTable,
			Columns: []string{flowperiod.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fpu.mutation.FlowMilestoneListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.RemovedFlowMilestoneListIDs(); len(nodes) > 0 && !fpu.mutation.FlowMilestoneListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.FlowMilestoneListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fpu.mutation.FlowDropWithdrawListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.RemovedFlowDropWithdrawListIDs(); len(nodes) > 0 && !fpu.mutation.FlowDropWithdrawListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.FlowDropWithdrawListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowperiod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowPeriodUpdateOne is the builder for updating a single FlowPeriod entity.
type FlowPeriodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowPeriodMutation
}

// SetFlowPeriodTitle sets the "flow_period_title" field.
func (fpuo *FlowPeriodUpdateOne) SetFlowPeriodTitle(s string) *FlowPeriodUpdateOne {
	fpuo.mutation.SetFlowPeriodTitle(s)
	return fpuo
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fpuo *FlowPeriodUpdateOne) SetParentID(id int) *FlowPeriodUpdateOne {
	fpuo.mutation.SetParentID(id)
	return fpuo
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fpuo *FlowPeriodUpdateOne) SetNillableParentID(id *int) *FlowPeriodUpdateOne {
	if id != nil {
		fpuo = fpuo.SetParentID(*id)
	}
	return fpuo
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fpuo *FlowPeriodUpdateOne) SetParent(p *ParticipantFlowModule) *FlowPeriodUpdateOne {
	return fpuo.SetParentID(p.ID)
}

// AddFlowMilestoneListIDs adds the "flow_milestone_list" edge to the FlowMilestone entity by IDs.
func (fpuo *FlowPeriodUpdateOne) AddFlowMilestoneListIDs(ids ...int) *FlowPeriodUpdateOne {
	fpuo.mutation.AddFlowMilestoneListIDs(ids...)
	return fpuo
}

// AddFlowMilestoneList adds the "flow_milestone_list" edges to the FlowMilestone entity.
func (fpuo *FlowPeriodUpdateOne) AddFlowMilestoneList(f ...*FlowMilestone) *FlowPeriodUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.AddFlowMilestoneListIDs(ids...)
}

// AddFlowDropWithdrawListIDs adds the "flow_drop_withdraw_list" edge to the FlowDropWithdraw entity by IDs.
func (fpuo *FlowPeriodUpdateOne) AddFlowDropWithdrawListIDs(ids ...int) *FlowPeriodUpdateOne {
	fpuo.mutation.AddFlowDropWithdrawListIDs(ids...)
	return fpuo
}

// AddFlowDropWithdrawList adds the "flow_drop_withdraw_list" edges to the FlowDropWithdraw entity.
func (fpuo *FlowPeriodUpdateOne) AddFlowDropWithdrawList(f ...*FlowDropWithdraw) *FlowPeriodUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.AddFlowDropWithdrawListIDs(ids...)
}

// Mutation returns the FlowPeriodMutation object of the builder.
func (fpuo *FlowPeriodUpdateOne) Mutation() *FlowPeriodMutation {
	return fpuo.mutation
}

// ClearParent clears the "parent" edge to the ParticipantFlowModule entity.
func (fpuo *FlowPeriodUpdateOne) ClearParent() *FlowPeriodUpdateOne {
	fpuo.mutation.ClearParent()
	return fpuo
}

// ClearFlowMilestoneList clears all "flow_milestone_list" edges to the FlowMilestone entity.
func (fpuo *FlowPeriodUpdateOne) ClearFlowMilestoneList() *FlowPeriodUpdateOne {
	fpuo.mutation.ClearFlowMilestoneList()
	return fpuo
}

// RemoveFlowMilestoneListIDs removes the "flow_milestone_list" edge to FlowMilestone entities by IDs.
func (fpuo *FlowPeriodUpdateOne) RemoveFlowMilestoneListIDs(ids ...int) *FlowPeriodUpdateOne {
	fpuo.mutation.RemoveFlowMilestoneListIDs(ids...)
	return fpuo
}

// RemoveFlowMilestoneList removes "flow_milestone_list" edges to FlowMilestone entities.
func (fpuo *FlowPeriodUpdateOne) RemoveFlowMilestoneList(f ...*FlowMilestone) *FlowPeriodUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.RemoveFlowMilestoneListIDs(ids...)
}

// ClearFlowDropWithdrawList clears all "flow_drop_withdraw_list" edges to the FlowDropWithdraw entity.
func (fpuo *FlowPeriodUpdateOne) ClearFlowDropWithdrawList() *FlowPeriodUpdateOne {
	fpuo.mutation.ClearFlowDropWithdrawList()
	return fpuo
}

// RemoveFlowDropWithdrawListIDs removes the "flow_drop_withdraw_list" edge to FlowDropWithdraw entities by IDs.
func (fpuo *FlowPeriodUpdateOne) RemoveFlowDropWithdrawListIDs(ids ...int) *FlowPeriodUpdateOne {
	fpuo.mutation.RemoveFlowDropWithdrawListIDs(ids...)
	return fpuo
}

// RemoveFlowDropWithdrawList removes "flow_drop_withdraw_list" edges to FlowDropWithdraw entities.
func (fpuo *FlowPeriodUpdateOne) RemoveFlowDropWithdrawList(f ...*FlowDropWithdraw) *FlowPeriodUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.RemoveFlowDropWithdrawListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fpuo *FlowPeriodUpdateOne) Select(field string, fields ...string) *FlowPeriodUpdateOne {
	fpuo.fields = append([]string{field}, fields...)
	return fpuo
}

// Save executes the query and returns the updated FlowPeriod entity.
func (fpuo *FlowPeriodUpdateOne) Save(ctx context.Context) (*FlowPeriod, error) {
	var (
		err  error
		node *FlowPeriod
	)
	if len(fpuo.hooks) == 0 {
		node, err = fpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowPeriodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fpuo.mutation = mutation
			node, err = fpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fpuo.hooks) - 1; i >= 0; i-- {
			if fpuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fpuo *FlowPeriodUpdateOne) SaveX(ctx context.Context) *FlowPeriod {
	node, err := fpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fpuo *FlowPeriodUpdateOne) Exec(ctx context.Context) error {
	_, err := fpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpuo *FlowPeriodUpdateOne) ExecX(ctx context.Context) {
	if err := fpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fpuo *FlowPeriodUpdateOne) sqlSave(ctx context.Context) (_node *FlowPeriod, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowperiod.Table,
			Columns: flowperiod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowperiod.FieldID,
			},
		},
	}
	id, ok := fpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowPeriod.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowperiod.FieldID)
		for _, f := range fields {
			if !flowperiod.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowperiod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fpuo.mutation.FlowPeriodTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowperiod.FieldFlowPeriodTitle,
		})
	}
	if fpuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowperiod.ParentTable,
			Columns: []string{flowperiod.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowperiod.ParentTable,
			Columns: []string{flowperiod.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fpuo.mutation.FlowMilestoneListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.RemovedFlowMilestoneListIDs(); len(nodes) > 0 && !fpuo.mutation.FlowMilestoneListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.FlowMilestoneListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowMilestoneListTable,
			Columns: []string{flowperiod.FlowMilestoneListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fpuo.mutation.FlowDropWithdrawListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.RemovedFlowDropWithdrawListIDs(); len(nodes) > 0 && !fpuo.mutation.FlowDropWithdrawListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.FlowDropWithdrawListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowperiod.FlowDropWithdrawListTable,
			Columns: []string{flowperiod.FlowDropWithdrawListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowPeriod{config: fpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowperiod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
