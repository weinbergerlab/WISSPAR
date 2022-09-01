// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventStatsDelete is the builder for deleting a OtherEventStats entity.
type OtherEventStatsDelete struct {
	config
	hooks    []Hook
	mutation *OtherEventStatsMutation
}

// Where appends a list predicates to the OtherEventStatsDelete builder.
func (oesd *OtherEventStatsDelete) Where(ps ...predicate.OtherEventStats) *OtherEventStatsDelete {
	oesd.mutation.Where(ps...)
	return oesd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oesd *OtherEventStatsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oesd.hooks) == 0 {
		affected, err = oesd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oesd.mutation = mutation
			affected, err = oesd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oesd.hooks) - 1; i >= 0; i-- {
			if oesd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oesd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oesd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oesd *OtherEventStatsDelete) ExecX(ctx context.Context) int {
	n, err := oesd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oesd *OtherEventStatsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: othereventstats.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: othereventstats.FieldID,
			},
		},
	}
	if ps := oesd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, oesd.driver, _spec)
}

// OtherEventStatsDeleteOne is the builder for deleting a single OtherEventStats entity.
type OtherEventStatsDeleteOne struct {
	oesd *OtherEventStatsDelete
}

// Exec executes the deletion query.
func (oesdo *OtherEventStatsDeleteOne) Exec(ctx context.Context) error {
	n, err := oesdo.oesd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{othereventstats.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oesdo *OtherEventStatsDeleteOne) ExecX(ctx context.Context) {
	oesdo.oesd.ExecX(ctx)
}
