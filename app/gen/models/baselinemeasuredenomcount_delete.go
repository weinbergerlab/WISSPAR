// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomCountDelete is the builder for deleting a BaselineMeasureDenomCount entity.
type BaselineMeasureDenomCountDelete struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureDenomCountMutation
}

// Where appends a list predicates to the BaselineMeasureDenomCountDelete builder.
func (bmdcd *BaselineMeasureDenomCountDelete) Where(ps ...predicate.BaselineMeasureDenomCount) *BaselineMeasureDenomCountDelete {
	bmdcd.mutation.Where(ps...)
	return bmdcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bmdcd *BaselineMeasureDenomCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmdcd.hooks) == 0 {
		affected, err = bmdcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmdcd.mutation = mutation
			affected, err = bmdcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmdcd.hooks) - 1; i >= 0; i-- {
			if bmdcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcd *BaselineMeasureDenomCountDelete) ExecX(ctx context.Context) int {
	n, err := bmdcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bmdcd *BaselineMeasureDenomCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinemeasuredenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenomcount.FieldID,
			},
		},
	}
	if ps := bmdcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bmdcd.driver, _spec)
}

// BaselineMeasureDenomCountDeleteOne is the builder for deleting a single BaselineMeasureDenomCount entity.
type BaselineMeasureDenomCountDeleteOne struct {
	bmdcd *BaselineMeasureDenomCountDelete
}

// Exec executes the deletion query.
func (bmdcdo *BaselineMeasureDenomCountDeleteOne) Exec(ctx context.Context) error {
	n, err := bmdcdo.bmdcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinemeasuredenomcount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcdo *BaselineMeasureDenomCountDeleteOne) ExecX(ctx context.Context) {
	bmdcdo.bmdcd.ExecX(ctx)
}
