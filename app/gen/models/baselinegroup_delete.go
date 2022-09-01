// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineGroupDelete is the builder for deleting a BaselineGroup entity.
type BaselineGroupDelete struct {
	config
	hooks    []Hook
	mutation *BaselineGroupMutation
}

// Where appends a list predicates to the BaselineGroupDelete builder.
func (bgd *BaselineGroupDelete) Where(ps ...predicate.BaselineGroup) *BaselineGroupDelete {
	bgd.mutation.Where(ps...)
	return bgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bgd *BaselineGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bgd.hooks) == 0 {
		affected, err = bgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bgd.mutation = mutation
			affected, err = bgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bgd.hooks) - 1; i >= 0; i-- {
			if bgd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgd *BaselineGroupDelete) ExecX(ctx context.Context) int {
	n, err := bgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bgd *BaselineGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinegroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinegroup.FieldID,
			},
		},
	}
	if ps := bgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bgd.driver, _spec)
}

// BaselineGroupDeleteOne is the builder for deleting a single BaselineGroup entity.
type BaselineGroupDeleteOne struct {
	bgd *BaselineGroupDelete
}

// Exec executes the deletion query.
func (bgdo *BaselineGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := bgdo.bgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinegroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bgdo *BaselineGroupDeleteOne) ExecX(ctx context.Context) {
	bgdo.bgd.ExecX(ctx)
}
