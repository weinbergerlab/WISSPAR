// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
)

// BaselineMeasureDenomCreate is the builder for creating a BaselineMeasureDenom entity.
type BaselineMeasureDenomCreate struct {
	config
	mutation *BaselineMeasureDenomMutation
	hooks    []Hook
}

// SetBaselineMeasureDenomUnits sets the "baseline_measure_denom_units" field.
func (bmdc *BaselineMeasureDenomCreate) SetBaselineMeasureDenomUnits(s string) *BaselineMeasureDenomCreate {
	bmdc.mutation.SetBaselineMeasureDenomUnits(s)
	return bmdc
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bmdc *BaselineMeasureDenomCreate) SetParentID(id int) *BaselineMeasureDenomCreate {
	bmdc.mutation.SetParentID(id)
	return bmdc
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bmdc *BaselineMeasureDenomCreate) SetNillableParentID(id *int) *BaselineMeasureDenomCreate {
	if id != nil {
		bmdc = bmdc.SetParentID(*id)
	}
	return bmdc
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bmdc *BaselineMeasureDenomCreate) SetParent(b *BaselineMeasure) *BaselineMeasureDenomCreate {
	return bmdc.SetParentID(b.ID)
}

// AddBaselineMeasureDenomCountListIDs adds the "baseline_measure_denom_count_list" edge to the BaselineMeasureDenomCount entity by IDs.
func (bmdc *BaselineMeasureDenomCreate) AddBaselineMeasureDenomCountListIDs(ids ...int) *BaselineMeasureDenomCreate {
	bmdc.mutation.AddBaselineMeasureDenomCountListIDs(ids...)
	return bmdc
}

// AddBaselineMeasureDenomCountList adds the "baseline_measure_denom_count_list" edges to the BaselineMeasureDenomCount entity.
func (bmdc *BaselineMeasureDenomCreate) AddBaselineMeasureDenomCountList(b ...*BaselineMeasureDenomCount) *BaselineMeasureDenomCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmdc.AddBaselineMeasureDenomCountListIDs(ids...)
}

// Mutation returns the BaselineMeasureDenomMutation object of the builder.
func (bmdc *BaselineMeasureDenomCreate) Mutation() *BaselineMeasureDenomMutation {
	return bmdc.mutation
}

// Save creates the BaselineMeasureDenom in the database.
func (bmdc *BaselineMeasureDenomCreate) Save(ctx context.Context) (*BaselineMeasureDenom, error) {
	var (
		err  error
		node *BaselineMeasureDenom
	)
	if len(bmdc.hooks) == 0 {
		if err = bmdc.check(); err != nil {
			return nil, err
		}
		node, err = bmdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bmdc.check(); err != nil {
				return nil, err
			}
			bmdc.mutation = mutation
			if node, err = bmdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bmdc.hooks) - 1; i >= 0; i-- {
			if bmdc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bmdc *BaselineMeasureDenomCreate) SaveX(ctx context.Context) *BaselineMeasureDenom {
	v, err := bmdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmdc *BaselineMeasureDenomCreate) Exec(ctx context.Context) error {
	_, err := bmdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdc *BaselineMeasureDenomCreate) ExecX(ctx context.Context) {
	if err := bmdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bmdc *BaselineMeasureDenomCreate) check() error {
	if _, ok := bmdc.mutation.BaselineMeasureDenomUnits(); !ok {
		return &ValidationError{Name: "baseline_measure_denom_units", err: errors.New(`models: missing required field "BaselineMeasureDenom.baseline_measure_denom_units"`)}
	}
	return nil
}

func (bmdc *BaselineMeasureDenomCreate) sqlSave(ctx context.Context) (*BaselineMeasureDenom, error) {
	_node, _spec := bmdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bmdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bmdc *BaselineMeasureDenomCreate) createSpec() (*BaselineMeasureDenom, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineMeasureDenom{config: bmdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinemeasuredenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenom.FieldID,
			},
		}
	)
	if value, ok := bmdc.mutation.BaselineMeasureDenomUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenom.FieldBaselineMeasureDenomUnits,
		})
		_node.BaselineMeasureDenomUnits = value
	}
	if nodes := bmdc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenom.ParentTable,
			Columns: []string{baselinemeasuredenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_measure_baseline_measure_denom_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bmdc.mutation.BaselineMeasureDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
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

// BaselineMeasureDenomCreateBulk is the builder for creating many BaselineMeasureDenom entities in bulk.
type BaselineMeasureDenomCreateBulk struct {
	config
	builders []*BaselineMeasureDenomCreate
}

// Save creates the BaselineMeasureDenom entities in the database.
func (bmdcb *BaselineMeasureDenomCreateBulk) Save(ctx context.Context) ([]*BaselineMeasureDenom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bmdcb.builders))
	nodes := make([]*BaselineMeasureDenom, len(bmdcb.builders))
	mutators := make([]Mutator, len(bmdcb.builders))
	for i := range bmdcb.builders {
		func(i int, root context.Context) {
			builder := bmdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineMeasureDenomMutation)
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
					_, err = mutators[i+1].Mutate(root, bmdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bmdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bmdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bmdcb *BaselineMeasureDenomCreateBulk) SaveX(ctx context.Context) []*BaselineMeasureDenom {
	v, err := bmdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmdcb *BaselineMeasureDenomCreateBulk) Exec(ctx context.Context) error {
	_, err := bmdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcb *BaselineMeasureDenomCreateBulk) ExecX(ctx context.Context) {
	if err := bmdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
