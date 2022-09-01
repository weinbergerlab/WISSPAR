// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
)

// UseCaseCreate is the builder for creating a UseCase entity.
type UseCaseCreate struct {
	config
	mutation *UseCaseMutation
	hooks    []Hook
}

// SetUseCaseName sets the "use_case_name" field.
func (ucc *UseCaseCreate) SetUseCaseName(s string) *UseCaseCreate {
	ucc.mutation.SetUseCaseName(s)
	return ucc
}

// SetUseCaseDescription sets the "use_case_description" field.
func (ucc *UseCaseCreate) SetUseCaseDescription(s string) *UseCaseCreate {
	ucc.mutation.SetUseCaseDescription(s)
	return ucc
}

// SetUseCaseCode sets the "use_case_code" field.
func (ucc *UseCaseCreate) SetUseCaseCode(s string) *UseCaseCreate {
	ucc.mutation.SetUseCaseCode(s)
	return ucc
}

// Mutation returns the UseCaseMutation object of the builder.
func (ucc *UseCaseCreate) Mutation() *UseCaseMutation {
	return ucc.mutation
}

// Save creates the UseCase in the database.
func (ucc *UseCaseCreate) Save(ctx context.Context) (*UseCase, error) {
	var (
		err  error
		node *UseCase
	)
	if len(ucc.hooks) == 0 {
		if err = ucc.check(); err != nil {
			return nil, err
		}
		node, err = ucc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UseCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucc.check(); err != nil {
				return nil, err
			}
			ucc.mutation = mutation
			if node, err = ucc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucc.hooks) - 1; i >= 0; i-- {
			if ucc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ucc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UseCaseCreate) SaveX(ctx context.Context) *UseCase {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UseCaseCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UseCaseCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UseCaseCreate) check() error {
	if _, ok := ucc.mutation.UseCaseName(); !ok {
		return &ValidationError{Name: "use_case_name", err: errors.New(`models: missing required field "UseCase.use_case_name"`)}
	}
	if _, ok := ucc.mutation.UseCaseDescription(); !ok {
		return &ValidationError{Name: "use_case_description", err: errors.New(`models: missing required field "UseCase.use_case_description"`)}
	}
	if _, ok := ucc.mutation.UseCaseCode(); !ok {
		return &ValidationError{Name: "use_case_code", err: errors.New(`models: missing required field "UseCase.use_case_code"`)}
	}
	return nil
}

func (ucc *UseCaseCreate) sqlSave(ctx context.Context) (*UseCase, error) {
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ucc *UseCaseCreate) createSpec() (*UseCase, *sqlgraph.CreateSpec) {
	var (
		_node = &UseCase{config: ucc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usecase.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usecase.FieldID,
			},
		}
	)
	if value, ok := ucc.mutation.UseCaseName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseName,
		})
		_node.UseCaseName = value
	}
	if value, ok := ucc.mutation.UseCaseDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseDescription,
		})
		_node.UseCaseDescription = value
	}
	if value, ok := ucc.mutation.UseCaseCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseCode,
		})
		_node.UseCaseCode = value
	}
	return _node, _spec
}

// UseCaseCreateBulk is the builder for creating many UseCase entities in bulk.
type UseCaseCreateBulk struct {
	config
	builders []*UseCaseCreate
}

// Save creates the UseCase entities in the database.
func (uccb *UseCaseCreateBulk) Save(ctx context.Context) ([]*UseCase, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UseCase, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UseCaseMutation)
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
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UseCaseCreateBulk) SaveX(ctx context.Context) []*UseCase {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UseCaseCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UseCaseCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}
