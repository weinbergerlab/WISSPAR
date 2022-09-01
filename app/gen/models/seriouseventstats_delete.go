// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventStatsDelete is the builder for deleting a SeriousEventStats entity.
type SeriousEventStatsDelete struct {
	config
	hooks    []Hook
	mutation *SeriousEventStatsMutation
}

// Where appends a list predicates to the SeriousEventStatsDelete builder.
func (sesd *SeriousEventStatsDelete) Where(ps ...predicate.SeriousEventStats) *SeriousEventStatsDelete {
	sesd.mutation.Where(ps...)
	return sesd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sesd *SeriousEventStatsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sesd.hooks) == 0 {
		affected, err = sesd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sesd.mutation = mutation
			affected, err = sesd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sesd.hooks) - 1; i >= 0; i-- {
			if sesd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sesd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sesd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sesd *SeriousEventStatsDelete) ExecX(ctx context.Context) int {
	n, err := sesd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sesd *SeriousEventStatsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: seriouseventstats.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriouseventstats.FieldID,
			},
		},
	}
	if ps := sesd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sesd.driver, _spec)
}

// SeriousEventStatsDeleteOne is the builder for deleting a single SeriousEventStats entity.
type SeriousEventStatsDeleteOne struct {
	sesd *SeriousEventStatsDelete
}

// Exec executes the deletion query.
func (sesdo *SeriousEventStatsDeleteOne) Exec(ctx context.Context) error {
	n, err := sesdo.sesd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{seriouseventstats.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sesdo *SeriousEventStatsDeleteOne) ExecX(ctx context.Context) {
	sesdo.sesd.ExecX(ctx)
}
