// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomCountDelete is the builder for deleting a OutcomeClassDenomCount entity.
type OutcomeClassDenomCountDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeClassDenomCountMutation
}

// Where appends a list predicates to the OutcomeClassDenomCountDelete builder.
func (ocdcd *OutcomeClassDenomCountDelete) Where(ps ...predicate.OutcomeClassDenomCount) *OutcomeClassDenomCountDelete {
	ocdcd.mutation.Where(ps...)
	return ocdcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocdcd *OutcomeClassDenomCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocdcd.hooks) == 0 {
		affected, err = ocdcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocdcd.mutation = mutation
			affected, err = ocdcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocdcd.hooks) - 1; i >= 0; i-- {
			if ocdcd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcd *OutcomeClassDenomCountDelete) ExecX(ctx context.Context) int {
	n, err := ocdcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocdcd *OutcomeClassDenomCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeclassdenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenomcount.FieldID,
			},
		},
	}
	if ps := ocdcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ocdcd.driver, _spec)
}

// OutcomeClassDenomCountDeleteOne is the builder for deleting a single OutcomeClassDenomCount entity.
type OutcomeClassDenomCountDeleteOne struct {
	ocdcd *OutcomeClassDenomCountDelete
}

// Exec executes the deletion query.
func (ocdcdo *OutcomeClassDenomCountDeleteOne) Exec(ctx context.Context) error {
	n, err := ocdcdo.ocdcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeclassdenomcount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcdo *OutcomeClassDenomCountDeleteOne) ExecX(ctx context.Context) {
	ocdcdo.ocdcd.ExecX(ctx)
}
