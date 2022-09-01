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
)

// FlowAchievementCreate is the builder for creating a FlowAchievement entity.
type FlowAchievementCreate struct {
	config
	mutation *FlowAchievementMutation
	hooks    []Hook
}

// SetFlowAchievementGroupID sets the "flow_achievement_group_id" field.
func (fac *FlowAchievementCreate) SetFlowAchievementGroupID(s string) *FlowAchievementCreate {
	fac.mutation.SetFlowAchievementGroupID(s)
	return fac
}

// SetFlowAchievementComment sets the "flow_achievement_comment" field.
func (fac *FlowAchievementCreate) SetFlowAchievementComment(s string) *FlowAchievementCreate {
	fac.mutation.SetFlowAchievementComment(s)
	return fac
}

// SetFlowAchievementNumSubjects sets the "flow_achievement_num_subjects" field.
func (fac *FlowAchievementCreate) SetFlowAchievementNumSubjects(s string) *FlowAchievementCreate {
	fac.mutation.SetFlowAchievementNumSubjects(s)
	return fac
}

// SetFlowAchievementNumUnits sets the "flow_achievement_num_units" field.
func (fac *FlowAchievementCreate) SetFlowAchievementNumUnits(s string) *FlowAchievementCreate {
	fac.mutation.SetFlowAchievementNumUnits(s)
	return fac
}

// SetParentID sets the "parent" edge to the FlowMilestone entity by ID.
func (fac *FlowAchievementCreate) SetParentID(id int) *FlowAchievementCreate {
	fac.mutation.SetParentID(id)
	return fac
}

// SetNillableParentID sets the "parent" edge to the FlowMilestone entity by ID if the given value is not nil.
func (fac *FlowAchievementCreate) SetNillableParentID(id *int) *FlowAchievementCreate {
	if id != nil {
		fac = fac.SetParentID(*id)
	}
	return fac
}

// SetParent sets the "parent" edge to the FlowMilestone entity.
func (fac *FlowAchievementCreate) SetParent(f *FlowMilestone) *FlowAchievementCreate {
	return fac.SetParentID(f.ID)
}

// Mutation returns the FlowAchievementMutation object of the builder.
func (fac *FlowAchievementCreate) Mutation() *FlowAchievementMutation {
	return fac.mutation
}

// Save creates the FlowAchievement in the database.
func (fac *FlowAchievementCreate) Save(ctx context.Context) (*FlowAchievement, error) {
	var (
		err  error
		node *FlowAchievement
	)
	if len(fac.hooks) == 0 {
		if err = fac.check(); err != nil {
			return nil, err
		}
		node, err = fac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowAchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fac.check(); err != nil {
				return nil, err
			}
			fac.mutation = mutation
			if node, err = fac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fac.hooks) - 1; i >= 0; i-- {
			if fac.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fac *FlowAchievementCreate) SaveX(ctx context.Context) *FlowAchievement {
	v, err := fac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fac *FlowAchievementCreate) Exec(ctx context.Context) error {
	_, err := fac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fac *FlowAchievementCreate) ExecX(ctx context.Context) {
	if err := fac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fac *FlowAchievementCreate) check() error {
	if _, ok := fac.mutation.FlowAchievementGroupID(); !ok {
		return &ValidationError{Name: "flow_achievement_group_id", err: errors.New(`models: missing required field "FlowAchievement.flow_achievement_group_id"`)}
	}
	if _, ok := fac.mutation.FlowAchievementComment(); !ok {
		return &ValidationError{Name: "flow_achievement_comment", err: errors.New(`models: missing required field "FlowAchievement.flow_achievement_comment"`)}
	}
	if _, ok := fac.mutation.FlowAchievementNumSubjects(); !ok {
		return &ValidationError{Name: "flow_achievement_num_subjects", err: errors.New(`models: missing required field "FlowAchievement.flow_achievement_num_subjects"`)}
	}
	if _, ok := fac.mutation.FlowAchievementNumUnits(); !ok {
		return &ValidationError{Name: "flow_achievement_num_units", err: errors.New(`models: missing required field "FlowAchievement.flow_achievement_num_units"`)}
	}
	return nil
}

func (fac *FlowAchievementCreate) sqlSave(ctx context.Context) (*FlowAchievement, error) {
	_node, _spec := fac.createSpec()
	if err := sqlgraph.CreateNode(ctx, fac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fac *FlowAchievementCreate) createSpec() (*FlowAchievement, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowAchievement{config: fac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowachievement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowachievement.FieldID,
			},
		}
	)
	if value, ok := fac.mutation.FlowAchievementGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementGroupID,
		})
		_node.FlowAchievementGroupID = value
	}
	if value, ok := fac.mutation.FlowAchievementComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementComment,
		})
		_node.FlowAchievementComment = value
	}
	if value, ok := fac.mutation.FlowAchievementNumSubjects(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumSubjects,
		})
		_node.FlowAchievementNumSubjects = value
	}
	if value, ok := fac.mutation.FlowAchievementNumUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumUnits,
		})
		_node.FlowAchievementNumUnits = value
	}
	if nodes := fac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowachievement.ParentTable,
			Columns: []string{flowachievement.ParentColumn},
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
		_node.flow_milestone_flow_achievement_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowAchievementCreateBulk is the builder for creating many FlowAchievement entities in bulk.
type FlowAchievementCreateBulk struct {
	config
	builders []*FlowAchievementCreate
}

// Save creates the FlowAchievement entities in the database.
func (facb *FlowAchievementCreateBulk) Save(ctx context.Context) ([]*FlowAchievement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(facb.builders))
	nodes := make([]*FlowAchievement, len(facb.builders))
	mutators := make([]Mutator, len(facb.builders))
	for i := range facb.builders {
		func(i int, root context.Context) {
			builder := facb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowAchievementMutation)
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
					_, err = mutators[i+1].Mutate(root, facb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, facb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, facb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (facb *FlowAchievementCreateBulk) SaveX(ctx context.Context) []*FlowAchievement {
	v, err := facb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (facb *FlowAchievementCreateBulk) Exec(ctx context.Context) error {
	_, err := facb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (facb *FlowAchievementCreateBulk) ExecX(ctx context.Context) {
	if err := facb.Exec(ctx); err != nil {
		panic(err)
	}
}
