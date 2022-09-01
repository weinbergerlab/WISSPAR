// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeOverviewDelete is the builder for deleting a OutcomeOverview entity.
type OutcomeOverviewDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeOverviewMutation
}

// Where appends a list predicates to the OutcomeOverviewDelete builder.
func (ood *OutcomeOverviewDelete) Where(ps ...predicate.OutcomeOverview) *OutcomeOverviewDelete {
	ood.mutation.Where(ps...)
	return ood
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ood *OutcomeOverviewDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ood.hooks) == 0 {
		affected, err = ood.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeOverviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ood.mutation = mutation
			affected, err = ood.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ood.hooks) - 1; i >= 0; i-- {
			if ood.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ood.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ood.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ood *OutcomeOverviewDelete) ExecX(ctx context.Context) int {
	n, err := ood.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ood *OutcomeOverviewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeoverview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeoverview.FieldID,
			},
		},
	}
	if ps := ood.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ood.driver, _spec)
}

// OutcomeOverviewDeleteOne is the builder for deleting a single OutcomeOverview entity.
type OutcomeOverviewDeleteOne struct {
	ood *OutcomeOverviewDelete
}

// Exec executes the deletion query.
func (oodo *OutcomeOverviewDeleteOne) Exec(ctx context.Context) error {
	n, err := oodo.ood.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeoverview.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oodo *OutcomeOverviewDeleteOne) ExecX(ctx context.Context) {
	oodo.ood.ExecX(ctx)
}
