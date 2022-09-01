// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
)

// VaccineCreate is the builder for creating a Vaccine entity.
type VaccineCreate struct {
	config
	mutation *VaccineMutation
	hooks    []Hook
}

// SetVaccineName sets the "vaccine_name" field.
func (vc *VaccineCreate) SetVaccineName(s string) *VaccineCreate {
	vc.mutation.SetVaccineName(s)
	return vc
}

// Mutation returns the VaccineMutation object of the builder.
func (vc *VaccineCreate) Mutation() *VaccineMutation {
	return vc.mutation
}

// Save creates the Vaccine in the database.
func (vc *VaccineCreate) Save(ctx context.Context) (*Vaccine, error) {
	var (
		err  error
		node *Vaccine
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VaccineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VaccineCreate) SaveX(ctx context.Context) *Vaccine {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VaccineCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VaccineCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VaccineCreate) check() error {
	if _, ok := vc.mutation.VaccineName(); !ok {
		return &ValidationError{Name: "vaccine_name", err: errors.New(`models: missing required field "Vaccine.vaccine_name"`)}
	}
	if v, ok := vc.mutation.VaccineName(); ok {
		if err := vaccine.VaccineNameValidator(v); err != nil {
			return &ValidationError{Name: "vaccine_name", err: fmt.Errorf(`models: validator failed for field "Vaccine.vaccine_name": %w`, err)}
		}
	}
	return nil
}

func (vc *VaccineCreate) sqlSave(ctx context.Context) (*Vaccine, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VaccineCreate) createSpec() (*Vaccine, *sqlgraph.CreateSpec) {
	var (
		_node = &Vaccine{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vaccine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vaccine.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.VaccineName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vaccine.FieldVaccineName,
		})
		_node.VaccineName = value
	}
	return _node, _spec
}

// VaccineCreateBulk is the builder for creating many Vaccine entities in bulk.
type VaccineCreateBulk struct {
	config
	builders []*VaccineCreate
}

// Save creates the Vaccine entities in the database.
func (vcb *VaccineCreateBulk) Save(ctx context.Context) ([]*Vaccine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vaccine, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VaccineMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VaccineCreateBulk) SaveX(ctx context.Context) []*Vaccine {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VaccineCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VaccineCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
