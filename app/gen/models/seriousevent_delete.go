// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
)

// SeriousEventDelete is the builder for deleting a SeriousEvent entity.
type SeriousEventDelete struct {
	config
	hooks    []Hook
	mutation *SeriousEventMutation
}

// Where appends a list predicates to the SeriousEventDelete builder.
func (sed *SeriousEventDelete) Where(ps ...predicate.SeriousEvent) *SeriousEventDelete {
	sed.mutation.Where(ps...)
	return sed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sed *SeriousEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sed.hooks) == 0 {
		affected, err = sed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sed.mutation = mutation
			affected, err = sed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sed.hooks) - 1; i >= 0; i-- {
			if sed.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sed *SeriousEventDelete) ExecX(ctx context.Context) int {
	n, err := sed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sed *SeriousEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: seriousevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriousevent.FieldID,
			},
		},
	}
	if ps := sed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sed.driver, _spec)
}

// SeriousEventDeleteOne is the builder for deleting a single SeriousEvent entity.
type SeriousEventDeleteOne struct {
	sed *SeriousEventDelete
}

// Exec executes the deletion query.
func (sedo *SeriousEventDeleteOne) Exec(ctx context.Context) error {
	n, err := sedo.sed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{seriousevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sedo *SeriousEventDeleteOne) ExecX(ctx context.Context) {
	sedo.sed.ExecX(ctx)
}
