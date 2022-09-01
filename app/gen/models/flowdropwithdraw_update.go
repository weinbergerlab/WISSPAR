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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowDropWithdrawUpdate is the builder for updating FlowDropWithdraw entities.
type FlowDropWithdrawUpdate struct {
	config
	hooks    []Hook
	mutation *FlowDropWithdrawMutation
}

// Where appends a list predicates to the FlowDropWithdrawUpdate builder.
func (fdwu *FlowDropWithdrawUpdate) Where(ps ...predicate.FlowDropWithdraw) *FlowDropWithdrawUpdate {
	fdwu.mutation.Where(ps...)
	return fdwu
}

// SetFlowDropWithdrawType sets the "flow_drop_withdraw_type" field.
func (fdwu *FlowDropWithdrawUpdate) SetFlowDropWithdrawType(s string) *FlowDropWithdrawUpdate {
	fdwu.mutation.SetFlowDropWithdrawType(s)
	return fdwu
}

// SetFlowDropWithdrawComment sets the "flow_drop_withdraw_comment" field.
func (fdwu *FlowDropWithdrawUpdate) SetFlowDropWithdrawComment(s string) *FlowDropWithdrawUpdate {
	fdwu.mutation.SetFlowDropWithdrawComment(s)
	return fdwu
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fdwu *FlowDropWithdrawUpdate) SetParentID(id int) *FlowDropWithdrawUpdate {
	fdwu.mutation.SetParentID(id)
	return fdwu
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fdwu *FlowDropWithdrawUpdate) SetNillableParentID(id *int) *FlowDropWithdrawUpdate {
	if id != nil {
		fdwu = fdwu.SetParentID(*id)
	}
	return fdwu
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fdwu *FlowDropWithdrawUpdate) SetParent(f *FlowPeriod) *FlowDropWithdrawUpdate {
	return fdwu.SetParentID(f.ID)
}

// AddFlowReasonListIDs adds the "flow_reason_list" edge to the FlowReason entity by IDs.
func (fdwu *FlowDropWithdrawUpdate) AddFlowReasonListIDs(ids ...int) *FlowDropWithdrawUpdate {
	fdwu.mutation.AddFlowReasonListIDs(ids...)
	return fdwu
}

// AddFlowReasonList adds the "flow_reason_list" edges to the FlowReason entity.
func (fdwu *FlowDropWithdrawUpdate) AddFlowReasonList(f ...*FlowReason) *FlowDropWithdrawUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdwu.AddFlowReasonListIDs(ids...)
}

// Mutation returns the FlowDropWithdrawMutation object of the builder.
func (fdwu *FlowDropWithdrawUpdate) Mutation() *FlowDropWithdrawMutation {
	return fdwu.mutation
}

// ClearParent clears the "parent" edge to the FlowPeriod entity.
func (fdwu *FlowDropWithdrawUpdate) ClearParent() *FlowDropWithdrawUpdate {
	fdwu.mutation.ClearParent()
	return fdwu
}

// ClearFlowReasonList clears all "flow_reason_list" edges to the FlowReason entity.
func (fdwu *FlowDropWithdrawUpdate) ClearFlowReasonList() *FlowDropWithdrawUpdate {
	fdwu.mutation.ClearFlowReasonList()
	return fdwu
}

// RemoveFlowReasonListIDs removes the "flow_reason_list" edge to FlowReason entities by IDs.
func (fdwu *FlowDropWithdrawUpdate) RemoveFlowReasonListIDs(ids ...int) *FlowDropWithdrawUpdate {
	fdwu.mutation.RemoveFlowReasonListIDs(ids...)
	return fdwu
}

// RemoveFlowReasonList removes "flow_reason_list" edges to FlowReason entities.
func (fdwu *FlowDropWithdrawUpdate) RemoveFlowReasonList(f ...*FlowReason) *FlowDropWithdrawUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdwu.RemoveFlowReasonListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fdwu *FlowDropWithdrawUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fdwu.hooks) == 0 {
		affected, err = fdwu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDropWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdwu.mutation = mutation
			affected, err = fdwu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdwu.hooks) - 1; i >= 0; i-- {
			if fdwu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fdwu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdwu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fdwu *FlowDropWithdrawUpdate) SaveX(ctx context.Context) int {
	affected, err := fdwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fdwu *FlowDropWithdrawUpdate) Exec(ctx context.Context) error {
	_, err := fdwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwu *FlowDropWithdrawUpdate) ExecX(ctx context.Context) {
	if err := fdwu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fdwu *FlowDropWithdrawUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowdropwithdraw.Table,
			Columns: flowdropwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowdropwithdraw.FieldID,
			},
		},
	}
	if ps := fdwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fdwu.mutation.FlowDropWithdrawType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawType,
		})
	}
	if value, ok := fdwu.mutation.FlowDropWithdrawComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawComment,
		})
	}
	if fdwu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowdropwithdraw.ParentTable,
			Columns: []string{flowdropwithdraw.ParentColumn},
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
	if nodes := fdwu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowdropwithdraw.ParentTable,
			Columns: []string{flowdropwithdraw.ParentColumn},
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
	if fdwu.mutation.FlowReasonListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdwu.mutation.RemovedFlowReasonListIDs(); len(nodes) > 0 && !fdwu.mutation.FlowReasonListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdwu.mutation.FlowReasonListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fdwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowdropwithdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowDropWithdrawUpdateOne is the builder for updating a single FlowDropWithdraw entity.
type FlowDropWithdrawUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowDropWithdrawMutation
}

// SetFlowDropWithdrawType sets the "flow_drop_withdraw_type" field.
func (fdwuo *FlowDropWithdrawUpdateOne) SetFlowDropWithdrawType(s string) *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.SetFlowDropWithdrawType(s)
	return fdwuo
}

// SetFlowDropWithdrawComment sets the "flow_drop_withdraw_comment" field.
func (fdwuo *FlowDropWithdrawUpdateOne) SetFlowDropWithdrawComment(s string) *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.SetFlowDropWithdrawComment(s)
	return fdwuo
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fdwuo *FlowDropWithdrawUpdateOne) SetParentID(id int) *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.SetParentID(id)
	return fdwuo
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fdwuo *FlowDropWithdrawUpdateOne) SetNillableParentID(id *int) *FlowDropWithdrawUpdateOne {
	if id != nil {
		fdwuo = fdwuo.SetParentID(*id)
	}
	return fdwuo
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fdwuo *FlowDropWithdrawUpdateOne) SetParent(f *FlowPeriod) *FlowDropWithdrawUpdateOne {
	return fdwuo.SetParentID(f.ID)
}

// AddFlowReasonListIDs adds the "flow_reason_list" edge to the FlowReason entity by IDs.
func (fdwuo *FlowDropWithdrawUpdateOne) AddFlowReasonListIDs(ids ...int) *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.AddFlowReasonListIDs(ids...)
	return fdwuo
}

// AddFlowReasonList adds the "flow_reason_list" edges to the FlowReason entity.
func (fdwuo *FlowDropWithdrawUpdateOne) AddFlowReasonList(f ...*FlowReason) *FlowDropWithdrawUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdwuo.AddFlowReasonListIDs(ids...)
}

// Mutation returns the FlowDropWithdrawMutation object of the builder.
func (fdwuo *FlowDropWithdrawUpdateOne) Mutation() *FlowDropWithdrawMutation {
	return fdwuo.mutation
}

// ClearParent clears the "parent" edge to the FlowPeriod entity.
func (fdwuo *FlowDropWithdrawUpdateOne) ClearParent() *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.ClearParent()
	return fdwuo
}

// ClearFlowReasonList clears all "flow_reason_list" edges to the FlowReason entity.
func (fdwuo *FlowDropWithdrawUpdateOne) ClearFlowReasonList() *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.ClearFlowReasonList()
	return fdwuo
}

// RemoveFlowReasonListIDs removes the "flow_reason_list" edge to FlowReason entities by IDs.
func (fdwuo *FlowDropWithdrawUpdateOne) RemoveFlowReasonListIDs(ids ...int) *FlowDropWithdrawUpdateOne {
	fdwuo.mutation.RemoveFlowReasonListIDs(ids...)
	return fdwuo
}

// RemoveFlowReasonList removes "flow_reason_list" edges to FlowReason entities.
func (fdwuo *FlowDropWithdrawUpdateOne) RemoveFlowReasonList(f ...*FlowReason) *FlowDropWithdrawUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdwuo.RemoveFlowReasonListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fdwuo *FlowDropWithdrawUpdateOne) Select(field string, fields ...string) *FlowDropWithdrawUpdateOne {
	fdwuo.fields = append([]string{field}, fields...)
	return fdwuo
}

// Save executes the query and returns the updated FlowDropWithdraw entity.
func (fdwuo *FlowDropWithdrawUpdateOne) Save(ctx context.Context) (*FlowDropWithdraw, error) {
	var (
		err  error
		node *FlowDropWithdraw
	)
	if len(fdwuo.hooks) == 0 {
		node, err = fdwuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDropWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdwuo.mutation = mutation
			node, err = fdwuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fdwuo.hooks) - 1; i >= 0; i-- {
			if fdwuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fdwuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdwuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fdwuo *FlowDropWithdrawUpdateOne) SaveX(ctx context.Context) *FlowDropWithdraw {
	node, err := fdwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fdwuo *FlowDropWithdrawUpdateOne) Exec(ctx context.Context) error {
	_, err := fdwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwuo *FlowDropWithdrawUpdateOne) ExecX(ctx context.Context) {
	if err := fdwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fdwuo *FlowDropWithdrawUpdateOne) sqlSave(ctx context.Context) (_node *FlowDropWithdraw, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowdropwithdraw.Table,
			Columns: flowdropwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowdropwithdraw.FieldID,
			},
		},
	}
	id, ok := fdwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowDropWithdraw.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fdwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowdropwithdraw.FieldID)
		for _, f := range fields {
			if !flowdropwithdraw.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowdropwithdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fdwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fdwuo.mutation.FlowDropWithdrawType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawType,
		})
	}
	if value, ok := fdwuo.mutation.FlowDropWithdrawComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawComment,
		})
	}
	if fdwuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowdropwithdraw.ParentTable,
			Columns: []string{flowdropwithdraw.ParentColumn},
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
	if nodes := fdwuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowdropwithdraw.ParentTable,
			Columns: []string{flowdropwithdraw.ParentColumn},
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
	if fdwuo.mutation.FlowReasonListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdwuo.mutation.RemovedFlowReasonListIDs(); len(nodes) > 0 && !fdwuo.mutation.FlowReasonListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdwuo.mutation.FlowReasonListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdropwithdraw.FlowReasonListTable,
			Columns: []string{flowdropwithdraw.FlowReasonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowreason.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowDropWithdraw{config: fdwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fdwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowdropwithdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
