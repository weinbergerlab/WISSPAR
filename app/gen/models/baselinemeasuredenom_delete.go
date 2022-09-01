// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomDelete is the builder for deleting a BaselineMeasureDenom entity.
type BaselineMeasureDenomDelete struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureDenomMutation
}

// Where appends a list predicates to the BaselineMeasureDenomDelete builder.
func (bmdd *BaselineMeasureDenomDelete) Where(ps ...predicate.BaselineMeasureDenom) *BaselineMeasureDenomDelete {
	bmdd.mutation.Where(ps...)
	return bmdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bmdd *BaselineMeasureDenomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmdd.hooks) == 0 {
		affected, err = bmdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmdd.mutation = mutation
			affected, err = bmdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmdd.hooks) - 1; i >= 0; i-- {
			if bmdd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdd *BaselineMeasureDenomDelete) ExecX(ctx context.Context) int {
	n, err := bmdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bmdd *BaselineMeasureDenomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinemeasuredenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenom.FieldID,
			},
		},
	}
	if ps := bmdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bmdd.driver, _spec)
}

// BaselineMeasureDenomDeleteOne is the builder for deleting a single BaselineMeasureDenom entity.
type BaselineMeasureDenomDeleteOne struct {
	bmdd *BaselineMeasureDenomDelete
}

// Exec executes the deletion query.
func (bmddo *BaselineMeasureDenomDeleteOne) Exec(ctx context.Context) error {
	n, err := bmddo.bmdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinemeasuredenom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bmddo *BaselineMeasureDenomDeleteOne) ExecX(ctx context.Context) {
	bmddo.bmdd.ExecX(ctx)
}
