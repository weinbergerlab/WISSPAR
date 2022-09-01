// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomCountDelete is the builder for deleting a OutcomeDenomCount entity.
type OutcomeDenomCountDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeDenomCountMutation
}

// Where appends a list predicates to the OutcomeDenomCountDelete builder.
func (odcd *OutcomeDenomCountDelete) Where(ps ...predicate.OutcomeDenomCount) *OutcomeDenomCountDelete {
	odcd.mutation.Where(ps...)
	return odcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odcd *OutcomeDenomCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odcd.hooks) == 0 {
		affected, err = odcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odcd.mutation = mutation
			affected, err = odcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odcd.hooks) - 1; i >= 0; i-- {
			if odcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcd *OutcomeDenomCountDelete) ExecX(ctx context.Context) int {
	n, err := odcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odcd *OutcomeDenomCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomedenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenomcount.FieldID,
			},
		},
	}
	if ps := odcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, odcd.driver, _spec)
}

// OutcomeDenomCountDeleteOne is the builder for deleting a single OutcomeDenomCount entity.
type OutcomeDenomCountDeleteOne struct {
	odcd *OutcomeDenomCountDelete
}

// Exec executes the deletion query.
func (odcdo *OutcomeDenomCountDeleteOne) Exec(ctx context.Context) error {
	n, err := odcdo.odcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomedenomcount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odcdo *OutcomeDenomCountDeleteOne) ExecX(ctx context.Context) {
	odcdo.odcd.ExecX(ctx)
}
