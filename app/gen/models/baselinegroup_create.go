// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
)

// BaselineGroupCreate is the builder for creating a BaselineGroup entity.
type BaselineGroupCreate struct {
	config
	mutation *BaselineGroupMutation
	hooks    []Hook
}

// SetBaselineGroupID sets the "baseline_group_id" field.
func (bgc *BaselineGroupCreate) SetBaselineGroupID(s string) *BaselineGroupCreate {
	bgc.mutation.SetBaselineGroupID(s)
	return bgc
}

// SetBaselineGroupTitle sets the "baseline_group_title" field.
func (bgc *BaselineGroupCreate) SetBaselineGroupTitle(s string) *BaselineGroupCreate {
	bgc.mutation.SetBaselineGroupTitle(s)
	return bgc
}

// SetBaselineGroupDescription sets the "baseline_group_description" field.
func (bgc *BaselineGroupCreate) SetBaselineGroupDescription(s string) *BaselineGroupCreate {
	bgc.mutation.SetBaselineGroupDescription(s)
	return bgc
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bgc *BaselineGroupCreate) SetParentID(id int) *BaselineGroupCreate {
	bgc.mutation.SetParentID(id)
	return bgc
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bgc *BaselineGroupCreate) SetNillableParentID(id *int) *BaselineGroupCreate {
	if id != nil {
		bgc = bgc.SetParentID(*id)
	}
	return bgc
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bgc *BaselineGroupCreate) SetParent(b *BaselineCharacteristicsModule) *BaselineGroupCreate {
	return bgc.SetParentID(b.ID)
}

// Mutation returns the BaselineGroupMutation object of the builder.
func (bgc *BaselineGroupCreate) Mutation() *BaselineGroupMutation {
	return bgc.mutation
}

// Save creates the BaselineGroup in the database.
func (bgc *BaselineGroupCreate) Save(ctx context.Context) (*BaselineGroup, error) {
	var (
		err  error
		node *BaselineGroup
	)
	if len(bgc.hooks) == 0 {
		if err = bgc.check(); err != nil {
			return nil, err
		}
		node, err = bgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bgc.check(); err != nil {
				return nil, err
			}
			bgc.mutation = mutation
			if node, err = bgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bgc.hooks) - 1; i >= 0; i-- {
			if bgc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bgc *BaselineGroupCreate) SaveX(ctx context.Context) *BaselineGroup {
	v, err := bgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bgc *BaselineGroupCreate) Exec(ctx context.Context) error {
	_, err := bgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgc *BaselineGroupCreate) ExecX(ctx context.Context) {
	if err := bgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bgc *BaselineGroupCreate) check() error {
	if _, ok := bgc.mutation.BaselineGroupID(); !ok {
		return &ValidationError{Name: "baseline_group_id", err: errors.New(`models: missing required field "BaselineGroup.baseline_group_id"`)}
	}
	if _, ok := bgc.mutation.BaselineGroupTitle(); !ok {
		return &ValidationError{Name: "baseline_group_title", err: errors.New(`models: missing required field "BaselineGroup.baseline_group_title"`)}
	}
	if _, ok := bgc.mutation.BaselineGroupDescription(); !ok {
		return &ValidationError{Name: "baseline_group_description", err: errors.New(`models: missing required field "BaselineGroup.baseline_group_description"`)}
	}
	return nil
}

func (bgc *BaselineGroupCreate) sqlSave(ctx context.Context) (*BaselineGroup, error) {
	_node, _spec := bgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bgc *BaselineGroupCreate) createSpec() (*BaselineGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineGroup{config: bgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinegroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinegroup.FieldID,
			},
		}
	)
	if value, ok := bgc.mutation.BaselineGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupID,
		})
		_node.BaselineGroupID = value
	}
	if value, ok := bgc.mutation.BaselineGroupTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupTitle,
		})
		_node.BaselineGroupTitle = value
	}
	if value, ok := bgc.mutation.BaselineGroupDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupDescription,
		})
		_node.BaselineGroupDescription = value
	}
	if nodes := bgc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinegroup.ParentTable,
			Columns: []string{baselinegroup.ParentColumn},
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
		_node.baseline_characteristics_module_baseline_group_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BaselineGroupCreateBulk is the builder for creating many BaselineGroup entities in bulk.
type BaselineGroupCreateBulk struct {
	config
	builders []*BaselineGroupCreate
}

// Save creates the BaselineGroup entities in the database.
func (bgcb *BaselineGroupCreateBulk) Save(ctx context.Context) ([]*BaselineGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bgcb.builders))
	nodes := make([]*BaselineGroup, len(bgcb.builders))
	mutators := make([]Mutator, len(bgcb.builders))
	for i := range bgcb.builders {
		func(i int, root context.Context) {
			builder := bgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, bgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bgcb *BaselineGroupCreateBulk) SaveX(ctx context.Context) []*BaselineGroup {
	v, err := bgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bgcb *BaselineGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := bgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgcb *BaselineGroupCreateBulk) ExecX(ctx context.Context) {
	if err := bgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
