// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomCountDelete is the builder for deleting a BaselineClassDenomCount entity.
type BaselineClassDenomCountDelete struct {
	config
	hooks    []Hook
	mutation *BaselineClassDenomCountMutation
}

// Where appends a list predicates to the BaselineClassDenomCountDelete builder.
func (bcdcd *BaselineClassDenomCountDelete) Where(ps ...predicate.BaselineClassDenomCount) *BaselineClassDenomCountDelete {
	bcdcd.mutation.Where(ps...)
	return bcdcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcdcd *BaselineClassDenomCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcdcd.hooks) == 0 {
		affected, err = bcdcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcdcd.mutation = mutation
			affected, err = bcdcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcdcd.hooks) - 1; i >= 0; i-- {
			if bcdcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcd *BaselineClassDenomCountDelete) ExecX(ctx context.Context) int {
	n, err := bcdcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcdcd *BaselineClassDenomCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselineclassdenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenomcount.FieldID,
			},
		},
	}
	if ps := bcdcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bcdcd.driver, _spec)
}

// BaselineClassDenomCountDeleteOne is the builder for deleting a single BaselineClassDenomCount entity.
type BaselineClassDenomCountDeleteOne struct {
	bcdcd *BaselineClassDenomCountDelete
}

// Exec executes the deletion query.
func (bcdcdo *BaselineClassDenomCountDeleteOne) Exec(ctx context.Context) error {
	n, err := bcdcdo.bcdcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselineclassdenomcount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcdo *BaselineClassDenomCountDeleteOne) ExecX(ctx context.Context) {
	bcdcdo.bcdcd.ExecX(ctx)
}
