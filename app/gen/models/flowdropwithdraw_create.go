// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
)

// FlowDropWithdrawCreate is the builder for creating a FlowDropWithdraw entity.
type FlowDropWithdrawCreate struct {
	config
	mutation *FlowDropWithdrawMutation
	hooks    []Hook
}

// SetFlowDropWithdrawType sets the "flow_drop_withdraw_type" field.
func (fdwc *FlowDropWithdrawCreate) SetFlowDropWithdrawType(s string) *FlowDropWithdrawCreate {
	fdwc.mutation.SetFlowDropWithdrawType(s)
	return fdwc
}

// SetFlowDropWithdrawComment sets the "flow_drop_withdraw_comment" field.
func (fdwc *FlowDropWithdrawCreate) SetFlowDropWithdrawComment(s string) *FlowDropWithdrawCreate {
	fdwc.mutation.SetFlowDropWithdrawComment(s)
	return fdwc
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fdwc *FlowDropWithdrawCreate) SetParentID(id int) *FlowDropWithdrawCreate {
	fdwc.mutation.SetParentID(id)
	return fdwc
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fdwc *FlowDropWithdrawCreate) SetNillableParentID(id *int) *FlowDropWithdrawCreate {
	if id != nil {
		fdwc = fdwc.SetParentID(*id)
	}
	return fdwc
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fdwc *FlowDropWithdrawCreate) SetParent(f *FlowPeriod) *FlowDropWithdrawCreate {
	return fdwc.SetParentID(f.ID)
}

// AddFlowReasonListIDs adds the "flow_reason_list" edge to the FlowReason entity by IDs.
func (fdwc *FlowDropWithdrawCreate) AddFlowReasonListIDs(ids ...int) *FlowDropWithdrawCreate {
	fdwc.mutation.AddFlowReasonListIDs(ids...)
	return fdwc
}

// AddFlowReasonList adds the "flow_reason_list" edges to the FlowReason entity.
func (fdwc *FlowDropWithdrawCreate) AddFlowReasonList(f ...*FlowReason) *FlowDropWithdrawCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdwc.AddFlowReasonListIDs(ids...)
}

// Mutation returns the FlowDropWithdrawMutation object of the builder.
func (fdwc *FlowDropWithdrawCreate) Mutation() *FlowDropWithdrawMutation {
	return fdwc.mutation
}

// Save creates the FlowDropWithdraw in the database.
func (fdwc *FlowDropWithdrawCreate) Save(ctx context.Context) (*FlowDropWithdraw, error) {
	var (
		err  error
		node *FlowDropWithdraw
	)
	if len(fdwc.hooks) == 0 {
		if err = fdwc.check(); err != nil {
			return nil, err
		}
		node, err = fdwc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDropWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fdwc.check(); err != nil {
				return nil, err
			}
			fdwc.mutation = mutation
			if node, err = fdwc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fdwc.hooks) - 1; i >= 0; i-- {
			if fdwc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fdwc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdwc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fdwc *FlowDropWithdrawCreate) SaveX(ctx context.Context) *FlowDropWithdraw {
	v, err := fdwc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fdwc *FlowDropWithdrawCreate) Exec(ctx context.Context) error {
	_, err := fdwc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwc *FlowDropWithdrawCreate) ExecX(ctx context.Context) {
	if err := fdwc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fdwc *FlowDropWithdrawCreate) check() error {
	if _, ok := fdwc.mutation.FlowDropWithdrawType(); !ok {
		return &ValidationError{Name: "flow_drop_withdraw_type", err: errors.New(`models: missing required field "FlowDropWithdraw.flow_drop_withdraw_type"`)}
	}
	if _, ok := fdwc.mutation.FlowDropWithdrawComment(); !ok {
		return &ValidationError{Name: "flow_drop_withdraw_comment", err: errors.New(`models: missing required field "FlowDropWithdraw.flow_drop_withdraw_comment"`)}
	}
	return nil
}

func (fdwc *FlowDropWithdrawCreate) sqlSave(ctx context.Context) (*FlowDropWithdraw, error) {
	_node, _spec := fdwc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fdwc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fdwc *FlowDropWithdrawCreate) createSpec() (*FlowDropWithdraw, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowDropWithdraw{config: fdwc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowdropwithdraw.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowdropwithdraw.FieldID,
			},
		}
	)
	if value, ok := fdwc.mutation.FlowDropWithdrawType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawType,
		})
		_node.FlowDropWithdrawType = value
	}
	if value, ok := fdwc.mutation.FlowDropWithdrawComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdropwithdraw.FieldFlowDropWithdrawComment,
		})
		_node.FlowDropWithdrawComment = value
	}
	if nodes := fdwc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.flow_period_flow_drop_withdraw_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fdwc.mutation.FlowReasonListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowDropWithdrawCreateBulk is the builder for creating many FlowDropWithdraw entities in bulk.
type FlowDropWithdrawCreateBulk struct {
	config
	builders []*FlowDropWithdrawCreate
}

// Save creates the FlowDropWithdraw entities in the database.
func (fdwcb *FlowDropWithdrawCreateBulk) Save(ctx context.Context) ([]*FlowDropWithdraw, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fdwcb.builders))
	nodes := make([]*FlowDropWithdraw, len(fdwcb.builders))
	mutators := make([]Mutator, len(fdwcb.builders))
	for i := range fdwcb.builders {
		func(i int, root context.Context) {
			builder := fdwcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowDropWithdrawMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fdwcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fdwcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fdwcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fdwcb *FlowDropWithdrawCreateBulk) SaveX(ctx context.Context) []*FlowDropWithdraw {
	v, err := fdwcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fdwcb *FlowDropWithdrawCreateBulk) Exec(ctx context.Context) error {
	_, err := fdwcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwcb *FlowDropWithdrawCreateBulk) ExecX(ctx context.Context) {
	if err := fdwcb.Exec(ctx); err != nil {
		panic(err)
	}
}
