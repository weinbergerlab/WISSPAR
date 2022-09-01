// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
)

// BaselineDenomCreate is the builder for creating a BaselineDenom entity.
type BaselineDenomCreate struct {
	config
	mutation *BaselineDenomMutation
	hooks    []Hook
}

// SetBaselineDenomUnits sets the "baseline_denom_units" field.
func (bdc *BaselineDenomCreate) SetBaselineDenomUnits(s string) *BaselineDenomCreate {
	bdc.mutation.SetBaselineDenomUnits(s)
	return bdc
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bdc *BaselineDenomCreate) SetParentID(id int) *BaselineDenomCreate {
	bdc.mutation.SetParentID(id)
	return bdc
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bdc *BaselineDenomCreate) SetNillableParentID(id *int) *BaselineDenomCreate {
	if id != nil {
		bdc = bdc.SetParentID(*id)
	}
	return bdc
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bdc *BaselineDenomCreate) SetParent(b *BaselineCharacteristicsModule) *BaselineDenomCreate {
	return bdc.SetParentID(b.ID)
}

// AddBaselineDenomCountListIDs adds the "baseline_denom_count_list" edge to the BaselineDenomCount entity by IDs.
func (bdc *BaselineDenomCreate) AddBaselineDenomCountListIDs(ids ...int) *BaselineDenomCreate {
	bdc.mutation.AddBaselineDenomCountListIDs(ids...)
	return bdc
}

// AddBaselineDenomCountList adds the "baseline_denom_count_list" edges to the BaselineDenomCount entity.
func (bdc *BaselineDenomCreate) AddBaselineDenomCountList(b ...*BaselineDenomCount) *BaselineDenomCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bdc.AddBaselineDenomCountListIDs(ids...)
}

// Mutation returns the BaselineDenomMutation object of the builder.
func (bdc *BaselineDenomCreate) Mutation() *BaselineDenomMutation {
	return bdc.mutation
}

// Save creates the BaselineDenom in the database.
func (bdc *BaselineDenomCreate) Save(ctx context.Context) (*BaselineDenom, error) {
	var (
		err  error
		node *BaselineDenom
	)
	if len(bdc.hooks) == 0 {
		if err = bdc.check(); err != nil {
			return nil, err
		}
		node, err = bdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bdc.check(); err != nil {
				return nil, err
			}
			bdc.mutation = mutation
			if node, err = bdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bdc.hooks) - 1; i >= 0; i-- {
			if bdc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bdc *BaselineDenomCreate) SaveX(ctx context.Context) *BaselineDenom {
	v, err := bdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdc *BaselineDenomCreate) Exec(ctx context.Context) error {
	_, err := bdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdc *BaselineDenomCreate) ExecX(ctx context.Context) {
	if err := bdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdc *BaselineDenomCreate) check() error {
	if _, ok := bdc.mutation.BaselineDenomUnits(); !ok {
		return &ValidationError{Name: "baseline_denom_units", err: errors.New(`models: missing required field "BaselineDenom.baseline_denom_units"`)}
	}
	return nil
}

func (bdc *BaselineDenomCreate) sqlSave(ctx context.Context) (*BaselineDenom, error) {
	_node, _spec := bdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bdc *BaselineDenomCreate) createSpec() (*BaselineDenom, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineDenom{config: bdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinedenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenom.FieldID,
			},
		}
	)
	if value, ok := bdc.mutation.BaselineDenomUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenom.FieldBaselineDenomUnits,
		})
		_node.BaselineDenomUnits = value
	}
	if nodes := bdc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenom.ParentTable,
			Columns: []string{baselinedenom.ParentColumn},
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
		_node.baseline_characteristics_module_baseline_denom_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bdc.mutation.BaselineDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
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

// BaselineDenomCreateBulk is the builder for creating many BaselineDenom entities in bulk.
type BaselineDenomCreateBulk struct {
	config
	builders []*BaselineDenomCreate
}

// Save creates the BaselineDenom entities in the database.
func (bdcb *BaselineDenomCreateBulk) Save(ctx context.Context) ([]*BaselineDenom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bdcb.builders))
	nodes := make([]*BaselineDenom, len(bdcb.builders))
	mutators := make([]Mutator, len(bdcb.builders))
	for i := range bdcb.builders {
		func(i int, root context.Context) {
			builder := bdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineDenomMutation)
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
					_, err = mutators[i+1].Mutate(root, bdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bdcb *BaselineDenomCreateBulk) SaveX(ctx context.Context) []*BaselineDenom {
	v, err := bdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdcb *BaselineDenomCreateBulk) Exec(ctx context.Context) error {
	_, err := bdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcb *BaselineDenomCreateBulk) ExecX(ctx context.Context) {
	if err := bdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
