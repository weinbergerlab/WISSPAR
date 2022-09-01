// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomDelete is the builder for deleting a OutcomeDenom entity.
type OutcomeDenomDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeDenomMutation
}

// Where appends a list predicates to the OutcomeDenomDelete builder.
func (odd *OutcomeDenomDelete) Where(ps ...predicate.OutcomeDenom) *OutcomeDenomDelete {
	odd.mutation.Where(ps...)
	return odd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odd *OutcomeDenomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odd.hooks) == 0 {
		affected, err = odd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odd.mutation = mutation
			affected, err = odd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odd.hooks) - 1; i >= 0; i-- {
			if odd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (odd *OutcomeDenomDelete) ExecX(ctx context.Context) int {
	n, err := odd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odd *OutcomeDenomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomedenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenom.FieldID,
			},
		},
	}
	if ps := odd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, odd.driver, _spec)
}

// OutcomeDenomDeleteOne is the builder for deleting a single OutcomeDenom entity.
type OutcomeDenomDeleteOne struct {
	odd *OutcomeDenomDelete
}

// Exec executes the deletion query.
func (oddo *OutcomeDenomDeleteOne) Exec(ctx context.Context) error {
	n, err := oddo.odd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomedenom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oddo *OutcomeDenomDeleteOne) ExecX(ctx context.Context) {
	oddo.odd.ExecX(ctx)
}
