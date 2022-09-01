// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowAchievementDelete is the builder for deleting a FlowAchievement entity.
type FlowAchievementDelete struct {
	config
	hooks    []Hook
	mutation *FlowAchievementMutation
}

// Where appends a list predicates to the FlowAchievementDelete builder.
func (fad *FlowAchievementDelete) Where(ps ...predicate.FlowAchievement) *FlowAchievementDelete {
	fad.mutation.Where(ps...)
	return fad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fad *FlowAchievementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fad.hooks) == 0 {
		affected, err = fad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowAchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fad.mutation = mutation
			affected, err = fad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fad.hooks) - 1; i >= 0; i-- {
			if fad.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fad *FlowAchievementDelete) ExecX(ctx context.Context) int {
	n, err := fad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fad *FlowAchievementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowachievement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowachievement.FieldID,
			},
		},
	}
	if ps := fad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fad.driver, _spec)
}

// FlowAchievementDeleteOne is the builder for deleting a single FlowAchievement entity.
type FlowAchievementDeleteOne struct {
	fad *FlowAchievementDelete
}

// Exec executes the deletion query.
func (fado *FlowAchievementDeleteOne) Exec(ctx context.Context) error {
	n, err := fado.fad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowachievement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fado *FlowAchievementDeleteOne) ExecX(ctx context.Context) {
	fado.fad.ExecX(ctx)
}
