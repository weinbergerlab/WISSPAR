// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomCountDelete is the builder for deleting a BaselineDenomCount entity.
type BaselineDenomCountDelete struct {
	config
	hooks    []Hook
	mutation *BaselineDenomCountMutation
}

// Where appends a list predicates to the BaselineDenomCountDelete builder.
func (bdcd *BaselineDenomCountDelete) Where(ps ...predicate.BaselineDenomCount) *BaselineDenomCountDelete {
	bdcd.mutation.Where(ps...)
	return bdcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bdcd *BaselineDenomCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bdcd.hooks) == 0 {
		affected, err = bdcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdcd.mutation = mutation
			affected, err = bdcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bdcd.hooks) - 1; i >= 0; i-- {
			if bdcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcd *BaselineDenomCountDelete) ExecX(ctx context.Context) int {
	n, err := bdcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bdcd *BaselineDenomCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinedenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenomcount.FieldID,
			},
		},
	}
	if ps := bdcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bdcd.driver, _spec)
}

// BaselineDenomCountDeleteOne is the builder for deleting a single BaselineDenomCount entity.
type BaselineDenomCountDeleteOne struct {
	bdcd *BaselineDenomCountDelete
}

// Exec executes the deletion query.
func (bdcdo *BaselineDenomCountDeleteOne) Exec(ctx context.Context) error {
	n, err := bdcdo.bdcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinedenomcount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcdo *BaselineDenomCountDeleteOne) ExecX(ctx context.Context) {
	bdcdo.bdcd.ExecX(ctx)
}
