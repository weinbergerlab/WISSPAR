// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowMilestoneDelete is the builder for deleting a FlowMilestone entity.
type FlowMilestoneDelete struct {
	config
	hooks    []Hook
	mutation *FlowMilestoneMutation
}

// Where appends a list predicates to the FlowMilestoneDelete builder.
func (fmd *FlowMilestoneDelete) Where(ps ...predicate.FlowMilestone) *FlowMilestoneDelete {
	fmd.mutation.Where(ps...)
	return fmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fmd *FlowMilestoneDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fmd.hooks) == 0 {
		affected, err = fmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowMilestoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fmd.mutation = mutation
			affected, err = fmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fmd.hooks) - 1; i >= 0; i-- {
			if fmd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmd *FlowMilestoneDelete) ExecX(ctx context.Context) int {
	n, err := fmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fmd *FlowMilestoneDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowmilestone.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowmilestone.FieldID,
			},
		},
	}
	if ps := fmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fmd.driver, _spec)
}

// FlowMilestoneDeleteOne is the builder for deleting a single FlowMilestone entity.
type FlowMilestoneDeleteOne struct {
	fmd *FlowMilestoneDelete
}

// Exec executes the deletion query.
func (fmdo *FlowMilestoneDeleteOne) Exec(ctx context.Context) error {
	n, err := fmdo.fmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowmilestone.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fmdo *FlowMilestoneDeleteOne) ExecX(ctx context.Context) {
	fmdo.fmd.ExecX(ctx)
}
