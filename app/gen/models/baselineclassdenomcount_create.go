// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
)

// BaselineClassDenomCountCreate is the builder for creating a BaselineClassDenomCount entity.
type BaselineClassDenomCountCreate struct {
	config
	mutation *BaselineClassDenomCountMutation
	hooks    []Hook
}

// SetBaselineClassDenomCountGroupID sets the "baseline_class_denom_count_group_id" field.
func (bcdcc *BaselineClassDenomCountCreate) SetBaselineClassDenomCountGroupID(s string) *BaselineClassDenomCountCreate {
	bcdcc.mutation.SetBaselineClassDenomCountGroupID(s)
	return bcdcc
}

// SetBaselineClassDenomCountValue sets the "baseline_class_denom_count_value" field.
func (bcdcc *BaselineClassDenomCountCreate) SetBaselineClassDenomCountValue(s string) *BaselineClassDenomCountCreate {
	bcdcc.mutation.SetBaselineClassDenomCountValue(s)
	return bcdcc
}

// SetParentID sets the "parent" edge to the BaselineClassDenom entity by ID.
func (bcdcc *BaselineClassDenomCountCreate) SetParentID(id int) *BaselineClassDenomCountCreate {
	bcdcc.mutation.SetParentID(id)
	return bcdcc
}

// SetNillableParentID sets the "parent" edge to the BaselineClassDenom entity by ID if the given value is not nil.
func (bcdcc *BaselineClassDenomCountCreate) SetNillableParentID(id *int) *BaselineClassDenomCountCreate {
	if id != nil {
		bcdcc = bcdcc.SetParentID(*id)
	}
	return bcdcc
}

// SetParent sets the "parent" edge to the BaselineClassDenom entity.
func (bcdcc *BaselineClassDenomCountCreate) SetParent(b *BaselineClassDenom) *BaselineClassDenomCountCreate {
	return bcdcc.SetParentID(b.ID)
}

// Mutation returns the BaselineClassDenomCountMutation object of the builder.
func (bcdcc *BaselineClassDenomCountCreate) Mutation() *BaselineClassDenomCountMutation {
	return bcdcc.mutation
}

// Save creates the BaselineClassDenomCount in the database.
func (bcdcc *BaselineClassDenomCountCreate) Save(ctx context.Context) (*BaselineClassDenomCount, error) {
	var (
		err  error
		node *BaselineClassDenomCount
	)
	if len(bcdcc.hooks) == 0 {
		if err = bcdcc.check(); err != nil {
			return nil, err
		}
		node, err = bcdcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcdcc.check(); err != nil {
				return nil, err
			}
			bcdcc.mutation = mutation
			if node, err = bcdcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcdcc.hooks) - 1; i >= 0; i-- {
			if bcdcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcdcc *BaselineClassDenomCountCreate) SaveX(ctx context.Context) *BaselineClassDenomCount {
	v, err := bcdcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcdcc *BaselineClassDenomCountCreate) Exec(ctx context.Context) error {
	_, err := bcdcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcc *BaselineClassDenomCountCreate) ExecX(ctx context.Context) {
	if err := bcdcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcdcc *BaselineClassDenomCountCreate) check() error {
	if _, ok := bcdcc.mutation.BaselineClassDenomCountGroupID(); !ok {
		return &ValidationError{Name: "baseline_class_denom_count_group_id", err: errors.New(`models: missing required field "BaselineClassDenomCount.baseline_class_denom_count_group_id"`)}
	}
	if _, ok := bcdcc.mutation.BaselineClassDenomCountValue(); !ok {
		return &ValidationError{Name: "baseline_class_denom_count_value", err: errors.New(`models: missing required field "BaselineClassDenomCount.baseline_class_denom_count_value"`)}
	}
	return nil
}

func (bcdcc *BaselineClassDenomCountCreate) sqlSave(ctx context.Context) (*BaselineClassDenomCount, error) {
	_node, _spec := bcdcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcdcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bcdcc *BaselineClassDenomCountCreate) createSpec() (*BaselineClassDenomCount, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineClassDenomCount{config: bcdcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselineclassdenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenomcount.FieldID,
			},
		}
	)
	if value, ok := bcdcc.mutation.BaselineClassDenomCountGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountGroupID,
		})
		_node.BaselineClassDenomCountGroupID = value
	}
	if value, ok := bcdcc.mutation.BaselineClassDenomCountValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountValue,
		})
		_node.BaselineClassDenomCountValue = value
	}
	if nodes := bcdcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenomcount.ParentTable,
			Columns: []string{baselineclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_class_denom_baseline_class_denom_count_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BaselineClassDenomCountCreateBulk is the builder for creating many BaselineClassDenomCount entities in bulk.
type BaselineClassDenomCountCreateBulk struct {
	config
	builders []*BaselineClassDenomCountCreate
}

// Save creates the BaselineClassDenomCount entities in the database.
func (bcdccb *BaselineClassDenomCountCreateBulk) Save(ctx context.Context) ([]*BaselineClassDenomCount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcdccb.builders))
	nodes := make([]*BaselineClassDenomCount, len(bcdccb.builders))
	mutators := make([]Mutator, len(bcdccb.builders))
	for i := range bcdccb.builders {
		func(i int, root context.Context) {
			builder := bcdccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineClassDenomCountMutation)
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
					_, err = mutators[i+1].Mutate(root, bcdccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcdccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcdccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcdccb *BaselineClassDenomCountCreateBulk) SaveX(ctx context.Context) []*BaselineClassDenomCount {
	v, err := bcdccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcdccb *BaselineClassDenomCountCreateBulk) Exec(ctx context.Context) error {
	_, err := bcdccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdccb *BaselineClassDenomCountCreateBulk) ExecX(ctx context.Context) {
	if err := bcdccb.Exec(ctx); err != nil {
		panic(err)
	}
}
