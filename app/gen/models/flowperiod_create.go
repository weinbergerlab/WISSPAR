// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
)

// FlowPeriodCreate is the builder for creating a FlowPeriod entity.
type FlowPeriodCreate struct {
	config
	mutation *FlowPeriodMutation
	hooks    []Hook
}

// SetFlowPeriodTitle sets the "flow_period_title" field.
func (fpc *FlowPeriodCreate) SetFlowPeriodTitle(s string) *FlowPeriodCreate {
	fpc.mutation.SetFlowPeriodTitle(s)
	return fpc
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fpc *FlowPeriodCreate) SetParentID(id int) *FlowPeriodCreate {
	fpc.mutation.SetParentID(id)
	return fpc
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fpc *FlowPeriodCreate) SetNillableParentID(id *int) *FlowPeriodCreate {
	if id != nil {
		fpc = fpc.SetParentID(*id)
	}
	return fpc
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fpc *FlowPeriodCreate) SetParent(p *ParticipantFlowModule) *FlowPeriodCreate {
	return fpc.SetParentID(p.ID)
}

// AddFlowMilestoneListIDs adds the "flow_milestone_list" edge to the FlowMilestone entity by IDs.
func (fpc *FlowPeriodCreate) AddFlowMilestoneListIDs(ids ...int) *FlowPeriodCreate {
	fpc.mutation.AddFlowMilestoneListIDs(ids...)
	return fpc
}

// AddFlowMilestoneList adds the "flow_milestone_list" edges to the FlowMilestone entity.
func (fpc *FlowPeriodCreate) AddFlowMilestoneList(f ...*FlowMilestone) *FlowPeriodCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpc.AddFlowMilestoneListIDs(ids...)
}

// AddFlowDropWithdrawListIDs adds the "flow_drop_withdraw_list" edge to the FlowDropWithdraw entity by IDs.
func (fpc *FlowPeriodCreate) AddFlowDropWithdrawListIDs(ids ...int) *FlowPeriodCreate {
	fpc.mutation.AddFlowDropWithdrawListIDs(ids...)
	return fpc
}

// AddFlowDropWithdrawList adds the "flow_drop_withdraw_list" edges to the FlowDropWithdraw entity.
func (fpc *FlowPeriodCreate) AddFlowDropWithdrawList(f ...*FlowDropWithdraw) *FlowPeriodCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpc.AddFlowDropWithdrawListIDs(ids...)
}

// Mutation returns the FlowPeriodMutation object of the builder.
func (fpc *FlowPeriodCreate) Mutation() *FlowPeriodMutation {
	return fpc.mutation
}

// Save creates the FlowPeriod in the database.
func (fpc *FlowPeriodCreate) Save(ctx context.Context) (*FlowPeriod, error) {
	var (
		err  error
		node *FlowPeriod
	)
	if len(fpc.hooks) == 0 {
		if err = fpc.check(); err != nil {
			return nil, err
		}
		node, err = fpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowPeriodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fpc.check(); err != nil {
				return nil, err
			}
			fpc.mutation = mutation
			if node, err = fpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fpc.hooks) - 1; i >= 0; i-- {
			if fpc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fpc *FlowPeriodCreate) SaveX(ctx context.Context) *FlowPeriod {
	v, err := fpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fpc *FlowPeriodCreate) Exec(ctx context.Context) error {
	_, err := fpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpc *FlowPeriodCreate) ExecX(ctx context.Context) {
	if err := fpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fpc *FlowPeriodCreate) check() error {
	if _, ok := fpc.mutation.FlowPeriodTitle(); !ok {
		return &ValidationError{Name: "flow_period_title", err: errors.New(`models: missing required field "FlowPeriod.flow_period_title"`)}
	}
	return nil
}

func (fpc *FlowPeriodCreate) sqlSave(ctx context.Context) (*FlowPeriod, error) {
	_node, _spec := fpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fpc *FlowPeriodCreate) createSpec() (*FlowPeriod, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowPeriod{config: fpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowperiod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowperiod.FieldID,
			},
		}
	)
	if value, ok := fpc.mutation.FlowPeriodTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowperiod.FieldFlowPeriodTitle,
		})
		_node.FlowPeriodTitle = value
	}
	if nodes := fpc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.participant_flow_module_flow_period_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fpc.mutation.FlowMilestoneListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fpc.mutation.FlowDropWithdrawListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowPeriodCreateBulk is the builder for creating many FlowPeriod entities in bulk.
type FlowPeriodCreateBulk struct {
	config
	builders []*FlowPeriodCreate
}

// Save creates the FlowPeriod entities in the database.
func (fpcb *FlowPeriodCreateBulk) Save(ctx context.Context) ([]*FlowPeriod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fpcb.builders))
	nodes := make([]*FlowPeriod, len(fpcb.builders))
	mutators := make([]Mutator, len(fpcb.builders))
	for i := range fpcb.builders {
		func(i int, root context.Context) {
			builder := fpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowPeriodMutation)
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
					_, err = mutators[i+1].Mutate(root, fpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fpcb *FlowPeriodCreateBulk) SaveX(ctx context.Context) []*FlowPeriod {
	v, err := fpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fpcb *FlowPeriodCreateBulk) Exec(ctx context.Context) error {
	_, err := fpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpcb *FlowPeriodCreateBulk) ExecX(ctx context.Context) {
	if err := fpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
