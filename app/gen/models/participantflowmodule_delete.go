// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ParticipantFlowModuleDelete is the builder for deleting a ParticipantFlowModule entity.
type ParticipantFlowModuleDelete struct {
	config
	hooks    []Hook
	mutation *ParticipantFlowModuleMutation
}

// Where appends a list predicates to the ParticipantFlowModuleDelete builder.
func (pfmd *ParticipantFlowModuleDelete) Where(ps ...predicate.ParticipantFlowModule) *ParticipantFlowModuleDelete {
	pfmd.mutation.Where(ps...)
	return pfmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pfmd *ParticipantFlowModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pfmd.hooks) == 0 {
		affected, err = pfmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantFlowModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pfmd.mutation = mutation
			affected, err = pfmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pfmd.hooks) - 1; i >= 0; i-- {
			if pfmd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pfmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmd *ParticipantFlowModuleDelete) ExecX(ctx context.Context) int {
	n, err := pfmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pfmd *ParticipantFlowModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: participantflowmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participantflowmodule.FieldID,
			},
		},
	}
	if ps := pfmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pfmd.driver, _spec)
}

// ParticipantFlowModuleDeleteOne is the builder for deleting a single ParticipantFlowModule entity.
type ParticipantFlowModuleDeleteOne struct {
	pfmd *ParticipantFlowModuleDelete
}

// Exec executes the deletion query.
func (pfmdo *ParticipantFlowModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := pfmdo.pfmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{participantflowmodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmdo *ParticipantFlowModuleDeleteOne) ExecX(ctx context.Context) {
	pfmdo.pfmd.ExecX(ctx)
}
