// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowReasonDelete is the builder for deleting a FlowReason entity.
type FlowReasonDelete struct {
	config
	hooks    []Hook
	mutation *FlowReasonMutation
}

// Where appends a list predicates to the FlowReasonDelete builder.
func (frd *FlowReasonDelete) Where(ps ...predicate.FlowReason) *FlowReasonDelete {
	frd.mutation.Where(ps...)
	return frd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (frd *FlowReasonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(frd.hooks) == 0 {
		affected, err = frd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			frd.mutation = mutation
			affected, err = frd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(frd.hooks) - 1; i >= 0; i-- {
			if frd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = frd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, frd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (frd *FlowReasonDelete) ExecX(ctx context.Context) int {
	n, err := frd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (frd *FlowReasonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowreason.FieldID,
			},
		},
	}
	if ps := frd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, frd.driver, _spec)
}

// FlowReasonDeleteOne is the builder for deleting a single FlowReason entity.
type FlowReasonDeleteOne struct {
	frd *FlowReasonDelete
}

// Exec executes the deletion query.
func (frdo *FlowReasonDeleteOne) Exec(ctx context.Context) error {
	n, err := frdo.frd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowreason.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (frdo *FlowReasonDeleteOne) ExecX(ctx context.Context) {
	frdo.frd.ExecX(ctx)
}
