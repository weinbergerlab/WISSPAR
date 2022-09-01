// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowDropWithdrawDelete is the builder for deleting a FlowDropWithdraw entity.
type FlowDropWithdrawDelete struct {
	config
	hooks    []Hook
	mutation *FlowDropWithdrawMutation
}

// Where appends a list predicates to the FlowDropWithdrawDelete builder.
func (fdwd *FlowDropWithdrawDelete) Where(ps ...predicate.FlowDropWithdraw) *FlowDropWithdrawDelete {
	fdwd.mutation.Where(ps...)
	return fdwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fdwd *FlowDropWithdrawDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fdwd.hooks) == 0 {
		affected, err = fdwd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDropWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdwd.mutation = mutation
			affected, err = fdwd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdwd.hooks) - 1; i >= 0; i-- {
			if fdwd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fdwd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdwd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwd *FlowDropWithdrawDelete) ExecX(ctx context.Context) int {
	n, err := fdwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fdwd *FlowDropWithdrawDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowdropwithdraw.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowdropwithdraw.FieldID,
			},
		},
	}
	if ps := fdwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fdwd.driver, _spec)
}

// FlowDropWithdrawDeleteOne is the builder for deleting a single FlowDropWithdraw entity.
type FlowDropWithdrawDeleteOne struct {
	fdwd *FlowDropWithdrawDelete
}

// Exec executes the deletion query.
func (fdwdo *FlowDropWithdrawDeleteOne) Exec(ctx context.Context) error {
	n, err := fdwdo.fdwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowdropwithdraw.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdwdo *FlowDropWithdrawDeleteOne) ExecX(ctx context.Context) {
	fdwdo.fdwd.ExecX(ctx)
}
