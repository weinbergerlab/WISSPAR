// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineCategoryDelete is the builder for deleting a BaselineCategory entity.
type BaselineCategoryDelete struct {
	config
	hooks    []Hook
	mutation *BaselineCategoryMutation
}

// Where appends a list predicates to the BaselineCategoryDelete builder.
func (bcd *BaselineCategoryDelete) Where(ps ...predicate.BaselineCategory) *BaselineCategoryDelete {
	bcd.mutation.Where(ps...)
	return bcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcd *BaselineCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcd.hooks) == 0 {
		affected, err = bcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcd.mutation = mutation
			affected, err = bcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcd.hooks) - 1; i >= 0; i-- {
			if bcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcd *BaselineCategoryDelete) ExecX(ctx context.Context) int {
	n, err := bcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcd *BaselineCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecategory.FieldID,
			},
		},
	}
	if ps := bcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bcd.driver, _spec)
}

// BaselineCategoryDeleteOne is the builder for deleting a single BaselineCategory entity.
type BaselineCategoryDeleteOne struct {
	bcd *BaselineCategoryDelete
}

// Exec executes the deletion query.
func (bcdo *BaselineCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := bcdo.bcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinecategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdo *BaselineCategoryDeleteOne) ExecX(ctx context.Context) {
	bcdo.bcd.ExecX(ctx)
}
