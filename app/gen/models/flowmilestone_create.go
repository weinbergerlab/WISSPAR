// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
)

// FlowMilestoneCreate is the builder for creating a FlowMilestone entity.
type FlowMilestoneCreate struct {
	config
	mutation *FlowMilestoneMutation
	hooks    []Hook
}

// SetFlowMilestoneType sets the "flow_milestone_type" field.
func (fmc *FlowMilestoneCreate) SetFlowMilestoneType(s string) *FlowMilestoneCreate {
	fmc.mutation.SetFlowMilestoneType(s)
	return fmc
}

// SetFlowMilestoneComment sets the "flow_milestone_comment" field.
func (fmc *FlowMilestoneCreate) SetFlowMilestoneComment(s string) *FlowMilestoneCreate {
	fmc.mutation.SetFlowMilestoneComment(s)
	return fmc
}

// SetParentID sets the "parent" edge to the FlowPeriod entity by ID.
func (fmc *FlowMilestoneCreate) SetParentID(id int) *FlowMilestoneCreate {
	fmc.mutation.SetParentID(id)
	return fmc
}

// SetNillableParentID sets the "parent" edge to the FlowPeriod entity by ID if the given value is not nil.
func (fmc *FlowMilestoneCreate) SetNillableParentID(id *int) *FlowMilestoneCreate {
	if id != nil {
		fmc = fmc.SetParentID(*id)
	}
	return fmc
}

// SetParent sets the "parent" edge to the FlowPeriod entity.
func (fmc *FlowMilestoneCreate) SetParent(f *FlowPeriod) *FlowMilestoneCreate {
	return fmc.SetParentID(f.ID)
}

// AddFlowAchievementListIDs adds the "flow_achievement_list" edge to the FlowAchievement entity by IDs.
func (fmc *FlowMilestoneCreate) AddFlowAchievementListIDs(ids ...int) *FlowMilestoneCreate {
	fmc.mutation.AddFlowAchievementListIDs(ids...)
	return fmc
}

// AddFlowAchievementList adds the "flow_achievement_list" edges to the FlowAchievement entity.
func (fmc *FlowMilestoneCreate) AddFlowAchievementList(f ...*FlowAchievement) *FlowMilestoneCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmc.AddFlowAchievementListIDs(ids...)
}

// Mutation returns the FlowMilestoneMutation object of the builder.
func (fmc *FlowMilestoneCreate) Mutation() *FlowMilestoneMutation {
	return fmc.mutation
}

// Save creates the FlowMilestone in the database.
func (fmc *FlowMilestoneCreate) Save(ctx context.Context) (*FlowMilestone, error) {
	var (
		err  error
		node *FlowMilestone
	)
	if len(fmc.hooks) == 0 {
		if err = fmc.check(); err != nil {
			return nil, err
		}
		node, err = fmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowMilestoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fmc.check(); err != nil {
				return nil, err
			}
			fmc.mutation = mutation
			if node, err = fmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fmc.hooks) - 1; i >= 0; i-- {
			if fmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fmc *FlowMilestoneCreate) SaveX(ctx context.Context) *FlowMilestone {
	v, err := fmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fmc *FlowMilestoneCreate) Exec(ctx context.Context) error {
	_, err := fmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmc *FlowMilestoneCreate) ExecX(ctx context.Context) {
	if err := fmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fmc *FlowMilestoneCreate) check() error {
	if _, ok := fmc.mutation.FlowMilestoneType(); !ok {
		return &ValidationError{Name: "flow_milestone_type", err: errors.New(`models: missing required field "FlowMilestone.flow_milestone_type"`)}
	}
	if _, ok := fmc.mutation.FlowMilestoneComment(); !ok {
		return &ValidationError{Name: "flow_milestone_comment", err: errors.New(`models: missing required field "FlowMilestone.flow_milestone_comment"`)}
	}
	return nil
}

func (fmc *FlowMilestoneCreate) sqlSave(ctx context.Context) (*FlowMilestone, error) {
	_node, _spec := fmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fmc *FlowMilestoneCreate) createSpec() (*FlowMilestone, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowMilestone{config: fmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowmilestone.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowmilestone.FieldID,
			},
		}
	)
	if value, ok := fmc.mutation.FlowMilestoneType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneType,
		})
		_node.FlowMilestoneType = value
	}
	if value, ok := fmc.mutation.FlowMilestoneComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowmilestone.FieldFlowMilestoneComment,
		})
		_node.FlowMilestoneComment = value
	}
	if nodes := fmc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.flow_period_flow_milestone_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fmc.mutation.FlowAchievementListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowMilestoneCreateBulk is the builder for creating many FlowMilestone entities in bulk.
type FlowMilestoneCreateBulk struct {
	config
	builders []*FlowMilestoneCreate
}

// Save creates the FlowMilestone entities in the database.
func (fmcb *FlowMilestoneCreateBulk) Save(ctx context.Context) ([]*FlowMilestone, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fmcb.builders))
	nodes := make([]*FlowMilestone, len(fmcb.builders))
	mutators := make([]Mutator, len(fmcb.builders))
	for i := range fmcb.builders {
		func(i int, root context.Context) {
			builder := fmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowMilestoneMutation)
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
					_, err = mutators[i+1].Mutate(root, fmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fmcb *FlowMilestoneCreateBulk) SaveX(ctx context.Context) []*FlowMilestone {
	v, err := fmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fmcb *FlowMilestoneCreateBulk) Exec(ctx context.Context) error {
	_, err := fmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmcb *FlowMilestoneCreateBulk) ExecX(ctx context.Context) {
	if err := fmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
