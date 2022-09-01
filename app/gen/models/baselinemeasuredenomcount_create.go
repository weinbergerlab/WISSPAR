// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
)

// BaselineMeasureDenomCountCreate is the builder for creating a BaselineMeasureDenomCount entity.
type BaselineMeasureDenomCountCreate struct {
	config
	mutation *BaselineMeasureDenomCountMutation
	hooks    []Hook
}

// SetBaselineMeasureDenomCountGroupID sets the "baseline_measure_denom_count_group_id" field.
func (bmdcc *BaselineMeasureDenomCountCreate) SetBaselineMeasureDenomCountGroupID(s string) *BaselineMeasureDenomCountCreate {
	bmdcc.mutation.SetBaselineMeasureDenomCountGroupID(s)
	return bmdcc
}

// SetBaselineMeasureDenomCountValue sets the "baseline_measure_denom_count_value" field.
func (bmdcc *BaselineMeasureDenomCountCreate) SetBaselineMeasureDenomCountValue(s string) *BaselineMeasureDenomCountCreate {
	bmdcc.mutation.SetBaselineMeasureDenomCountValue(s)
	return bmdcc
}

// SetParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID.
func (bmdcc *BaselineMeasureDenomCountCreate) SetParentID(id int) *BaselineMeasureDenomCountCreate {
	bmdcc.mutation.SetParentID(id)
	return bmdcc
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID if the given value is not nil.
func (bmdcc *BaselineMeasureDenomCountCreate) SetNillableParentID(id *int) *BaselineMeasureDenomCountCreate {
	if id != nil {
		bmdcc = bmdcc.SetParentID(*id)
	}
	return bmdcc
}

// SetParent sets the "parent" edge to the BaselineMeasureDenom entity.
func (bmdcc *BaselineMeasureDenomCountCreate) SetParent(b *BaselineMeasureDenom) *BaselineMeasureDenomCountCreate {
	return bmdcc.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasureDenomCountMutation object of the builder.
func (bmdcc *BaselineMeasureDenomCountCreate) Mutation() *BaselineMeasureDenomCountMutation {
	return bmdcc.mutation
}

// Save creates the BaselineMeasureDenomCount in the database.
func (bmdcc *BaselineMeasureDenomCountCreate) Save(ctx context.Context) (*BaselineMeasureDenomCount, error) {
	var (
		err  error
		node *BaselineMeasureDenomCount
	)
	if len(bmdcc.hooks) == 0 {
		if err = bmdcc.check(); err != nil {
			return nil, err
		}
		node, err = bmdcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bmdcc.check(); err != nil {
				return nil, err
			}
			bmdcc.mutation = mutation
			if node, err = bmdcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bmdcc.hooks) - 1; i >= 0; i-- {
			if bmdcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bmdcc *BaselineMeasureDenomCountCreate) SaveX(ctx context.Context) *BaselineMeasureDenomCount {
	v, err := bmdcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmdcc *BaselineMeasureDenomCountCreate) Exec(ctx context.Context) error {
	_, err := bmdcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcc *BaselineMeasureDenomCountCreate) ExecX(ctx context.Context) {
	if err := bmdcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bmdcc *BaselineMeasureDenomCountCreate) check() error {
	if _, ok := bmdcc.mutation.BaselineMeasureDenomCountGroupID(); !ok {
		return &ValidationError{Name: "baseline_measure_denom_count_group_id", err: errors.New(`models: missing required field "BaselineMeasureDenomCount.baseline_measure_denom_count_group_id"`)}
	}
	if _, ok := bmdcc.mutation.BaselineMeasureDenomCountValue(); !ok {
		return &ValidationError{Name: "baseline_measure_denom_count_value", err: errors.New(`models: missing required field "BaselineMeasureDenomCount.baseline_measure_denom_count_value"`)}
	}
	return nil
}

func (bmdcc *BaselineMeasureDenomCountCreate) sqlSave(ctx context.Context) (*BaselineMeasureDenomCount, error) {
	_node, _spec := bmdcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bmdcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bmdcc *BaselineMeasureDenomCountCreate) createSpec() (*BaselineMeasureDenomCount, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineMeasureDenomCount{config: bmdcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinemeasuredenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenomcount.FieldID,
			},
		}
	)
	if value, ok := bmdcc.mutation.BaselineMeasureDenomCountGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID,
		})
		_node.BaselineMeasureDenomCountGroupID = value
	}
	if value, ok := bmdcc.mutation.BaselineMeasureDenomCountValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountValue,
		})
		_node.BaselineMeasureDenomCountValue = value
	}
	if nodes := bmdcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenomcount.ParentTable,
			Columns: []string{baselinemeasuredenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_measure_denom_baseline_measure_denom_count_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BaselineMeasureDenomCountCreateBulk is the builder for creating many BaselineMeasureDenomCount entities in bulk.
type BaselineMeasureDenomCountCreateBulk struct {
	config
	builders []*BaselineMeasureDenomCountCreate
}

// Save creates the BaselineMeasureDenomCount entities in the database.
func (bmdccb *BaselineMeasureDenomCountCreateBulk) Save(ctx context.Context) ([]*BaselineMeasureDenomCount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bmdccb.builders))
	nodes := make([]*BaselineMeasureDenomCount, len(bmdccb.builders))
	mutators := make([]Mutator, len(bmdccb.builders))
	for i := range bmdccb.builders {
		func(i int, root context.Context) {
			builder := bmdccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineMeasureDenomCountMutation)
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
					_, err = mutators[i+1].Mutate(root, bmdccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bmdccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bmdccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bmdccb *BaselineMeasureDenomCountCreateBulk) SaveX(ctx context.Context) []*BaselineMeasureDenomCount {
	v, err := bmdccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmdccb *BaselineMeasureDenomCountCreateBulk) Exec(ctx context.Context) error {
	_, err := bmdccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdccb *BaselineMeasureDenomCountCreateBulk) ExecX(ctx context.Context) {
	if err := bmdccb.Exec(ctx); err != nil {
		panic(err)
	}
}
