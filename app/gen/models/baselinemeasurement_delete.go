// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasurementDelete is the builder for deleting a BaselineMeasurement entity.
type BaselineMeasurementDelete struct {
	config
	hooks    []Hook
	mutation *BaselineMeasurementMutation
}

// Where appends a list predicates to the BaselineMeasurementDelete builder.
func (bmd *BaselineMeasurementDelete) Where(ps ...predicate.BaselineMeasurement) *BaselineMeasurementDelete {
	bmd.mutation.Where(ps...)
	return bmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bmd *BaselineMeasurementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmd.hooks) == 0 {
		affected, err = bmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmd.mutation = mutation
			affected, err = bmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmd.hooks) - 1; i >= 0; i-- {
			if bmd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmd *BaselineMeasurementDelete) ExecX(ctx context.Context) int {
	n, err := bmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bmd *BaselineMeasurementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinemeasurement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasurement.FieldID,
			},
		},
	}
	if ps := bmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bmd.driver, _spec)
}

// BaselineMeasurementDeleteOne is the builder for deleting a single BaselineMeasurement entity.
type BaselineMeasurementDeleteOne struct {
	bmd *BaselineMeasurementDelete
}

// Exec executes the deletion query.
func (bmdo *BaselineMeasurementDeleteOne) Exec(ctx context.Context) error {
	n, err := bmdo.bmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinemeasurement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdo *BaselineMeasurementDeleteOne) ExecX(ctx context.Context) {
	bmdo.bmd.ExecX(ctx)
}
