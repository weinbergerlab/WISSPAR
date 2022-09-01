// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisDelete is the builder for deleting a OutcomeAnalysis entity.
type OutcomeAnalysisDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeAnalysisMutation
}

// Where appends a list predicates to the OutcomeAnalysisDelete builder.
func (oad *OutcomeAnalysisDelete) Where(ps ...predicate.OutcomeAnalysis) *OutcomeAnalysisDelete {
	oad.mutation.Where(ps...)
	return oad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oad *OutcomeAnalysisDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oad.hooks) == 0 {
		affected, err = oad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oad.mutation = mutation
			affected, err = oad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oad.hooks) - 1; i >= 0; i-- {
			if oad.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oad *OutcomeAnalysisDelete) ExecX(ctx context.Context) int {
	n, err := oad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oad *OutcomeAnalysisDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeanalysis.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysis.FieldID,
			},
		},
	}
	if ps := oad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, oad.driver, _spec)
}

// OutcomeAnalysisDeleteOne is the builder for deleting a single OutcomeAnalysis entity.
type OutcomeAnalysisDeleteOne struct {
	oad *OutcomeAnalysisDelete
}

// Exec executes the deletion query.
func (oado *OutcomeAnalysisDeleteOne) Exec(ctx context.Context) error {
	n, err := oado.oad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeanalysis.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oado *OutcomeAnalysisDeleteOne) ExecX(ctx context.Context) {
	oado.oad.ExecX(ctx)
}
