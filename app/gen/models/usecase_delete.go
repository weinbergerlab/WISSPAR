// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
)

// UseCaseDelete is the builder for deleting a UseCase entity.
type UseCaseDelete struct {
	config
	hooks    []Hook
	mutation *UseCaseMutation
}

// Where appends a list predicates to the UseCaseDelete builder.
func (ucd *UseCaseDelete) Where(ps ...predicate.UseCase) *UseCaseDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UseCaseDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucd.hooks) == 0 {
		affected, err = ucd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UseCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucd.mutation = mutation
			affected, err = ucd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucd.hooks) - 1; i >= 0; i-- {
			if ucd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ucd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UseCaseDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UseCaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usecase.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usecase.FieldID,
			},
		},
	}
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
}

// UseCaseDeleteOne is the builder for deleting a single UseCase entity.
type UseCaseDeleteOne struct {
	ucd *UseCaseDelete
}

// Exec executes the deletion query.
func (ucdo *UseCaseDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usecase.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UseCaseDeleteOne) ExecX(ctx context.Context) {
	ucdo.ucd.ExecX(ctx)
}
