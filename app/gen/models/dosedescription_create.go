// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
)

// DoseDescriptionCreate is the builder for creating a DoseDescription entity.
type DoseDescriptionCreate struct {
	config
	mutation *DoseDescriptionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ddc *DoseDescriptionCreate) SetName(s string) *DoseDescriptionCreate {
	ddc.mutation.SetName(s)
	return ddc
}

// Mutation returns the DoseDescriptionMutation object of the builder.
func (ddc *DoseDescriptionCreate) Mutation() *DoseDescriptionMutation {
	return ddc.mutation
}

// Save creates the DoseDescription in the database.
func (ddc *DoseDescriptionCreate) Save(ctx context.Context) (*DoseDescription, error) {
	var (
		err  error
		node *DoseDescription
	)
	if len(ddc.hooks) == 0 {
		if err = ddc.check(); err != nil {
			return nil, err
		}
		node, err = ddc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoseDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddc.check(); err != nil {
				return nil, err
			}
			ddc.mutation = mutation
			if node, err = ddc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ddc.hooks) - 1; i >= 0; i-- {
			if ddc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ddc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ddc *DoseDescriptionCreate) SaveX(ctx context.Context) *DoseDescription {
	v, err := ddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddc *DoseDescriptionCreate) Exec(ctx context.Context) error {
	_, err := ddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddc *DoseDescriptionCreate) ExecX(ctx context.Context) {
	if err := ddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddc *DoseDescriptionCreate) check() error {
	if _, ok := ddc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`models: missing required field "DoseDescription.name"`)}
	}
	if v, ok := ddc.mutation.Name(); ok {
		if err := dosedescription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "DoseDescription.name": %w`, err)}
		}
	}
	return nil
}

func (ddc *DoseDescriptionCreate) sqlSave(ctx context.Context) (*DoseDescription, error) {
	_node, _spec := ddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ddc *DoseDescriptionCreate) createSpec() (*DoseDescription, *sqlgraph.CreateSpec) {
	var (
		_node = &DoseDescription{config: ddc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dosedescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dosedescription.FieldID,
			},
		}
	)
	if value, ok := ddc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dosedescription.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// DoseDescriptionCreateBulk is the builder for creating many DoseDescription entities in bulk.
type DoseDescriptionCreateBulk struct {
	config
	builders []*DoseDescriptionCreate
}

// Save creates the DoseDescription entities in the database.
func (ddcb *DoseDescriptionCreateBulk) Save(ctx context.Context) ([]*DoseDescription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ddcb.builders))
	nodes := make([]*DoseDescription, len(ddcb.builders))
	mutators := make([]Mutator, len(ddcb.builders))
	for i := range ddcb.builders {
		func(i int, root context.Context) {
			builder := ddcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DoseDescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, ddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddcb *DoseDescriptionCreateBulk) SaveX(ctx context.Context) []*DoseDescription {
	v, err := ddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddcb *DoseDescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := ddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddcb *DoseDescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := ddcb.Exec(ctx); err != nil {
		panic(err)
	}
}
