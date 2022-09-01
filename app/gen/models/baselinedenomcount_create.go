// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
)

// BaselineDenomCountCreate is the builder for creating a BaselineDenomCount entity.
type BaselineDenomCountCreate struct {
	config
	mutation *BaselineDenomCountMutation
	hooks    []Hook
}

// SetBaselineDenomCountGroupID sets the "baseline_denom_count_group_id" field.
func (bdcc *BaselineDenomCountCreate) SetBaselineDenomCountGroupID(s string) *BaselineDenomCountCreate {
	bdcc.mutation.SetBaselineDenomCountGroupID(s)
	return bdcc
}

// SetBaselineDenomCountValue sets the "baseline_denom_count_value" field.
func (bdcc *BaselineDenomCountCreate) SetBaselineDenomCountValue(s string) *BaselineDenomCountCreate {
	bdcc.mutation.SetBaselineDenomCountValue(s)
	return bdcc
}

// SetParentID sets the "parent" edge to the BaselineDenom entity by ID.
func (bdcc *BaselineDenomCountCreate) SetParentID(id int) *BaselineDenomCountCreate {
	bdcc.mutation.SetParentID(id)
	return bdcc
}

// SetNillableParentID sets the "parent" edge to the BaselineDenom entity by ID if the given value is not nil.
func (bdcc *BaselineDenomCountCreate) SetNillableParentID(id *int) *BaselineDenomCountCreate {
	if id != nil {
		bdcc = bdcc.SetParentID(*id)
	}
	return bdcc
}

// SetParent sets the "parent" edge to the BaselineDenom entity.
func (bdcc *BaselineDenomCountCreate) SetParent(b *BaselineDenom) *BaselineDenomCountCreate {
	return bdcc.SetParentID(b.ID)
}

// Mutation returns the BaselineDenomCountMutation object of the builder.
func (bdcc *BaselineDenomCountCreate) Mutation() *BaselineDenomCountMutation {
	return bdcc.mutation
}

// Save creates the BaselineDenomCount in the database.
func (bdcc *BaselineDenomCountCreate) Save(ctx context.Context) (*BaselineDenomCount, error) {
	var (
		err  error
		node *BaselineDenomCount
	)
	if len(bdcc.hooks) == 0 {
		if err = bdcc.check(); err != nil {
			return nil, err
		}
		node, err = bdcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bdcc.check(); err != nil {
				return nil, err
			}
			bdcc.mutation = mutation
			if node, err = bdcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bdcc.hooks) - 1; i >= 0; i-- {
			if bdcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bdcc *BaselineDenomCountCreate) SaveX(ctx context.Context) *BaselineDenomCount {
	v, err := bdcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdcc *BaselineDenomCountCreate) Exec(ctx context.Context) error {
	_, err := bdcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcc *BaselineDenomCountCreate) ExecX(ctx context.Context) {
	if err := bdcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdcc *BaselineDenomCountCreate) check() error {
	if _, ok := bdcc.mutation.BaselineDenomCountGroupID(); !ok {
		return &ValidationError{Name: "baseline_denom_count_group_id", err: errors.New(`models: missing required field "BaselineDenomCount.baseline_denom_count_group_id"`)}
	}
	if _, ok := bdcc.mutation.BaselineDenomCountValue(); !ok {
		return &ValidationError{Name: "baseline_denom_count_value", err: errors.New(`models: missing required field "BaselineDenomCount.baseline_denom_count_value"`)}
	}
	return nil
}

func (bdcc *BaselineDenomCountCreate) sqlSave(ctx context.Context) (*BaselineDenomCount, error) {
	_node, _spec := bdcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bdcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bdcc *BaselineDenomCountCreate) createSpec() (*BaselineDenomCount, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineDenomCount{config: bdcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinedenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenomcount.FieldID,
			},
		}
	)
	if value, ok := bdcc.mutation.BaselineDenomCountGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountGroupID,
		})
		_node.BaselineDenomCountGroupID = value
	}
	if value, ok := bdcc.mutation.BaselineDenomCountValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountValue,
		})
		_node.BaselineDenomCountValue = value
	}
	if nodes := bdcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenomcount.ParentTable,
			Columns: []string{baselinedenomcount.ParentColumn},
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
		_node.baseline_denom_baseline_denom_count_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BaselineDenomCountCreateBulk is the builder for creating many BaselineDenomCount entities in bulk.
type BaselineDenomCountCreateBulk struct {
	config
	builders []*BaselineDenomCountCreate
}

// Save creates the BaselineDenomCount entities in the database.
func (bdccb *BaselineDenomCountCreateBulk) Save(ctx context.Context) ([]*BaselineDenomCount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bdccb.builders))
	nodes := make([]*BaselineDenomCount, len(bdccb.builders))
	mutators := make([]Mutator, len(bdccb.builders))
	for i := range bdccb.builders {
		func(i int, root context.Context) {
			builder := bdccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineDenomCountMutation)
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
					_, err = mutators[i+1].Mutate(root, bdccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bdccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bdccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bdccb *BaselineDenomCountCreateBulk) SaveX(ctx context.Context) []*BaselineDenomCount {
	v, err := bdccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdccb *BaselineDenomCountCreateBulk) Exec(ctx context.Context) error {
	_, err := bdccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdccb *BaselineDenomCountCreateBulk) ExecX(ctx context.Context) {
	if err := bdccb.Exec(ctx); err != nil {
		panic(err)
	}
}
