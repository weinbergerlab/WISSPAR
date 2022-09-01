// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDelete is the builder for deleting a BaselineMeasure entity.
type BaselineMeasureDelete struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureMutation
}

// Where appends a list predicates to the BaselineMeasureDelete builder.
func (bmd *BaselineMeasureDelete) Where(ps ...predicate.BaselineMeasure) *BaselineMeasureDelete {
	bmd.mutation.Where(ps...)
	return bmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bmd *BaselineMeasureDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmd.hooks) == 0 {
		affected, err = bmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureMutation)
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
func (bmd *BaselineMeasureDelete) ExecX(ctx context.Context) int {
	n, err := bmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bmd *BaselineMeasureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinemeasure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasure.FieldID,
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

// BaselineMeasureDeleteOne is the builder for deleting a single BaselineMeasure entity.
type BaselineMeasureDeleteOne struct {
	bmd *BaselineMeasureDelete
}

// Exec executes the deletion query.
func (bmdo *BaselineMeasureDeleteOne) Exec(ctx context.Context) error {
	n, err := bmdo.bmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinemeasure.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdo *BaselineMeasureDeleteOne) ExecX(ctx context.Context) {
	bmdo.bmd.ExecX(ctx)
}
