// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowPeriodDelete is the builder for deleting a FlowPeriod entity.
type FlowPeriodDelete struct {
	config
	hooks    []Hook
	mutation *FlowPeriodMutation
}

// Where appends a list predicates to the FlowPeriodDelete builder.
func (fpd *FlowPeriodDelete) Where(ps ...predicate.FlowPeriod) *FlowPeriodDelete {
	fpd.mutation.Where(ps...)
	return fpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fpd *FlowPeriodDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fpd.hooks) == 0 {
		affected, err = fpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowPeriodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fpd.mutation = mutation
			affected, err = fpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fpd.hooks) - 1; i >= 0; i-- {
			if fpd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpd *FlowPeriodDelete) ExecX(ctx context.Context) int {
	n, err := fpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fpd *FlowPeriodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowperiod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowperiod.FieldID,
			},
		},
	}
	if ps := fpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fpd.driver, _spec)
}

// FlowPeriodDeleteOne is the builder for deleting a single FlowPeriod entity.
type FlowPeriodDeleteOne struct {
	fpd *FlowPeriodDelete
}

// Exec executes the deletion query.
func (fpdo *FlowPeriodDeleteOne) Exec(ctx context.Context) error {
	n, err := fpdo.fpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowperiod.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fpdo *FlowPeriodDeleteOne) ExecX(ctx context.Context) {
	fpdo.fpd.ExecX(ctx)
}
