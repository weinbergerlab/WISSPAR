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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowReasonUpdate is the builder for updating FlowReason entities.
type FlowReasonUpdate struct {
	config
	hooks    []Hook
	mutation *FlowReasonMutation
}

// Where appends a list predicates to the FlowReasonUpdate builder.
func (fru *FlowReasonUpdate) Where(ps ...predicate.FlowReason) *FlowReasonUpdate {
	fru.mutation.Where(ps...)
	return fru
}

// SetFlowReasonGroupID sets the "flow_reason_group_id" field.
func (fru *FlowReasonUpdate) SetFlowReasonGroupID(s string) *FlowReasonUpdate {
	fru.mutation.SetFlowReasonGroupID(s)
	return fru
}

// SetFlowReasonComment sets the "flow_reason_comment" field.
func (fru *FlowReasonUpdate) SetFlowReasonComment(s string) *FlowReasonUpdate {
	fru.mutation.SetFlowReasonComment(s)
	return fru
}

// SetFlowReasonNumSubjects sets the "flow_reason_num_subjects" field.
func (fru *FlowReasonUpdate) SetFlowReasonNumSubjects(s string) *FlowReasonUpdate {
	fru.mutation.SetFlowReasonNumSubjects(s)
	return fru
}

// SetFlowReasonNumUnits sets the "flow_reason_num_units" field.
func (fru *FlowReasonUpdate) SetFlowReasonNumUnits(s string) *FlowReasonUpdate {
	fru.mutation.SetFlowReasonNumUnits(s)
	return fru
}

// SetParentID sets the "parent" edge to the FlowDropWithdraw entity by ID.
func (fru *FlowReasonUpdate) SetParentID(id int) *FlowReasonUpdate {
	fru.mutation.SetParentID(id)
	return fru
}

// SetNillableParentID sets the "parent" edge to the FlowDropWithdraw entity by ID if the given value is not nil.
func (fru *FlowReasonUpdate) SetNillableParentID(id *int) *FlowReasonUpdate {
	if id != nil {
		fru = fru.SetParentID(*id)
	}
	return fru
}

// SetParent sets the "parent" edge to the FlowDropWithdraw entity.
func (fru *FlowReasonUpdate) SetParent(f *FlowDropWithdraw) *FlowReasonUpdate {
	return fru.SetParentID(f.ID)
}

// Mutation returns the FlowReasonMutation object of the builder.
func (fru *FlowReasonUpdate) Mutation() *FlowReasonMutation {
	return fru.mutation
}

// ClearParent clears the "parent" edge to the FlowDropWithdraw entity.
func (fru *FlowReasonUpdate) ClearParent() *FlowReasonUpdate {
	fru.mutation.ClearParent()
	return fru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fru *FlowReasonUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fru.hooks) == 0 {
		affected, err = fru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fru.mutation = mutation
			affected, err = fru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fru.hooks) - 1; i >= 0; i-- {
			if fru.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fru *FlowReasonUpdate) SaveX(ctx context.Context) int {
	affected, err := fru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fru *FlowReasonUpdate) Exec(ctx context.Context) error {
	_, err := fru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fru *FlowReasonUpdate) ExecX(ctx context.Context) {
	if err := fru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fru *FlowReasonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowreason.Table,
			Columns: flowreason.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowreason.FieldID,
			},
		},
	}
	if ps := fru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fru.mutation.FlowReasonGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonGroupID,
		})
	}
	if value, ok := fru.mutation.FlowReasonComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonComment,
		})
	}
	if value, ok := fru.mutation.FlowReasonNumSubjects(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumSubjects,
		})
	}
	if value, ok := fru.mutation.FlowReasonNumUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumUnits,
		})
	}
	if fru.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowreason.ParentTable,
			Columns: []string{flowreason.ParentColumn},
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
	if nodes := fru.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowreason.ParentTable,
			Columns: []string{flowreason.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowreason.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowReasonUpdateOne is the builder for updating a single FlowReason entity.
type FlowReasonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowReasonMutation
}

// SetFlowReasonGroupID sets the "flow_reason_group_id" field.
func (fruo *FlowReasonUpdateOne) SetFlowReasonGroupID(s string) *FlowReasonUpdateOne {
	fruo.mutation.SetFlowReasonGroupID(s)
	return fruo
}

// SetFlowReasonComment sets the "flow_reason_comment" field.
func (fruo *FlowReasonUpdateOne) SetFlowReasonComment(s string) *FlowReasonUpdateOne {
	fruo.mutation.SetFlowReasonComment(s)
	return fruo
}

// SetFlowReasonNumSubjects sets the "flow_reason_num_subjects" field.
func (fruo *FlowReasonUpdateOne) SetFlowReasonNumSubjects(s string) *FlowReasonUpdateOne {
	fruo.mutation.SetFlowReasonNumSubjects(s)
	return fruo
}

// SetFlowReasonNumUnits sets the "flow_reason_num_units" field.
func (fruo *FlowReasonUpdateOne) SetFlowReasonNumUnits(s string) *FlowReasonUpdateOne {
	fruo.mutation.SetFlowReasonNumUnits(s)
	return fruo
}

// SetParentID sets the "parent" edge to the FlowDropWithdraw entity by ID.
func (fruo *FlowReasonUpdateOne) SetParentID(id int) *FlowReasonUpdateOne {
	fruo.mutation.SetParentID(id)
	return fruo
}

// SetNillableParentID sets the "parent" edge to the FlowDropWithdraw entity by ID if the given value is not nil.
func (fruo *FlowReasonUpdateOne) SetNillableParentID(id *int) *FlowReasonUpdateOne {
	if id != nil {
		fruo = fruo.SetParentID(*id)
	}
	return fruo
}

// SetParent sets the "parent" edge to the FlowDropWithdraw entity.
func (fruo *FlowReasonUpdateOne) SetParent(f *FlowDropWithdraw) *FlowReasonUpdateOne {
	return fruo.SetParentID(f.ID)
}

// Mutation returns the FlowReasonMutation object of the builder.
func (fruo *FlowReasonUpdateOne) Mutation() *FlowReasonMutation {
	return fruo.mutation
}

// ClearParent clears the "parent" edge to the FlowDropWithdraw entity.
func (fruo *FlowReasonUpdateOne) ClearParent() *FlowReasonUpdateOne {
	fruo.mutation.ClearParent()
	return fruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fruo *FlowReasonUpdateOne) Select(field string, fields ...string) *FlowReasonUpdateOne {
	fruo.fields = append([]string{field}, fields...)
	return fruo
}

// Save executes the query and returns the updated FlowReason entity.
func (fruo *FlowReasonUpdateOne) Save(ctx context.Context) (*FlowReason, error) {
	var (
		err  error
		node *FlowReason
	)
	if len(fruo.hooks) == 0 {
		node, err = fruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fruo.mutation = mutation
			node, err = fruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fruo.hooks) - 1; i >= 0; i-- {
			if fruo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fruo *FlowReasonUpdateOne) SaveX(ctx context.Context) *FlowReason {
	node, err := fruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fruo *FlowReasonUpdateOne) Exec(ctx context.Context) error {
	_, err := fruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fruo *FlowReasonUpdateOne) ExecX(ctx context.Context) {
	if err := fruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fruo *FlowReasonUpdateOne) sqlSave(ctx context.Context) (_node *FlowReason, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowreason.Table,
			Columns: flowreason.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowreason.FieldID,
			},
		},
	}
	id, ok := fruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowReason.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowreason.FieldID)
		for _, f := range fields {
			if !flowreason.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowreason.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fruo.mutation.FlowReasonGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonGroupID,
		})
	}
	if value, ok := fruo.mutation.FlowReasonComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonComment,
		})
	}
	if value, ok := fruo.mutation.FlowReasonNumSubjects(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumSubjects,
		})
	}
	if value, ok := fruo.mutation.FlowReasonNumUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumUnits,
		})
	}
	if fruo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowreason.ParentTable,
			Columns: []string{flowreason.ParentColumn},
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
	if nodes := fruo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowreason.ParentTable,
			Columns: []string{flowreason.ParentColumn},
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
	_node = &FlowReason{config: fruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowreason.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
