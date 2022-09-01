// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/manufacturer"
)

// ManufacturerCreate is the builder for creating a Manufacturer entity.
type ManufacturerCreate struct {
	config
	mutation *ManufacturerMutation
	hooks    []Hook
}

// SetManufacturerName sets the "manufacturer_name" field.
func (mc *ManufacturerCreate) SetManufacturerName(s string) *ManufacturerCreate {
	mc.mutation.SetManufacturerName(s)
	return mc
}

// Mutation returns the ManufacturerMutation object of the builder.
func (mc *ManufacturerCreate) Mutation() *ManufacturerMutation {
	return mc.mutation
}

// Save creates the Manufacturer in the database.
func (mc *ManufacturerCreate) Save(ctx context.Context) (*Manufacturer, error) {
	var (
		err  error
		node *Manufacturer
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ManufacturerCreate) SaveX(ctx context.Context) *Manufacturer {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ManufacturerCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ManufacturerCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ManufacturerCreate) check() error {
	if _, ok := mc.mutation.ManufacturerName(); !ok {
		return &ValidationError{Name: "manufacturer_name", err: errors.New(`models: missing required field "Manufacturer.manufacturer_name"`)}
	}
	if v, ok := mc.mutation.ManufacturerName(); ok {
		if err := manufacturer.ManufacturerNameValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer_name", err: fmt.Errorf(`models: validator failed for field "Manufacturer.manufacturer_name": %w`, err)}
		}
	}
	return nil
}

func (mc *ManufacturerCreate) sqlSave(ctx context.Context) (*Manufacturer, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *ManufacturerCreate) createSpec() (*Manufacturer, *sqlgraph.CreateSpec) {
	var (
		_node = &Manufacturer{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: manufacturer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manufacturer.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.ManufacturerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldManufacturerName,
		})
		_node.ManufacturerName = value
	}
	return _node, _spec
}

// ManufacturerCreateBulk is the builder for creating many Manufacturer entities in bulk.
type ManufacturerCreateBulk struct {
	config
	builders []*ManufacturerCreate
}

// Save creates the Manufacturer entities in the database.
func (mcb *ManufacturerCreateBulk) Save(ctx context.Context) ([]*Manufacturer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Manufacturer, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ManufacturerMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *ManufacturerCreateBulk) SaveX(ctx context.Context) []*Manufacturer {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ManufacturerCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ManufacturerCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
