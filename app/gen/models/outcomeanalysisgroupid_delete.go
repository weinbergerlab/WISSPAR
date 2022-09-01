// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisGroupIDDelete is the builder for deleting a OutcomeAnalysisGroupID entity.
type OutcomeAnalysisGroupIDDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeAnalysisGroupIDMutation
}

// Where appends a list predicates to the OutcomeAnalysisGroupIDDelete builder.
func (oagid *OutcomeAnalysisGroupIDDelete) Where(ps ...predicate.OutcomeAnalysisGroupID) *OutcomeAnalysisGroupIDDelete {
	oagid.mutation.Where(ps...)
	return oagid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oagid *OutcomeAnalysisGroupIDDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oagid.hooks) == 0 {
		affected, err = oagid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisGroupIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oagid.mutation = mutation
			affected, err = oagid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oagid.hooks) - 1; i >= 0; i-- {
			if oagid.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oagid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oagid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oagid *OutcomeAnalysisGroupIDDelete) ExecX(ctx context.Context) int {
	n, err := oagid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oagid *OutcomeAnalysisGroupIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeanalysisgroupid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysisgroupid.FieldID,
			},
		},
	}
	if ps := oagid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, oagid.driver, _spec)
}

// OutcomeAnalysisGroupIDDeleteOne is the builder for deleting a single OutcomeAnalysisGroupID entity.
type OutcomeAnalysisGroupIDDeleteOne struct {
	oagid *OutcomeAnalysisGroupIDDelete
}

// Exec executes the deletion query.
func (oagido *OutcomeAnalysisGroupIDDeleteOne) Exec(ctx context.Context) error {
	n, err := oagido.oagid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeanalysisgroupid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oagido *OutcomeAnalysisGroupIDDeleteOne) ExecX(ctx context.Context) {
	oagido.oagid.ExecX(ctx)
}
