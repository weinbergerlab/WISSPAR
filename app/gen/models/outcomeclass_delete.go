// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDelete is the builder for deleting a OutcomeClass entity.
type OutcomeClassDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeClassMutation
}

// Where appends a list predicates to the OutcomeClassDelete builder.
func (ocd *OutcomeClassDelete) Where(ps ...predicate.OutcomeClass) *OutcomeClassDelete {
	ocd.mutation.Where(ps...)
	return ocd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocd *OutcomeClassDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocd.hooks) == 0 {
		affected, err = ocd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocd.mutation = mutation
			affected, err = ocd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocd.hooks) - 1; i >= 0; i-- {
			if ocd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocd *OutcomeClassDelete) ExecX(ctx context.Context) int {
	n, err := ocd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocd *OutcomeClassDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclass.FieldID,
			},
		},
	}
	if ps := ocd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ocd.driver, _spec)
}

// OutcomeClassDeleteOne is the builder for deleting a single OutcomeClass entity.
type OutcomeClassDeleteOne struct {
	ocd *OutcomeClassDelete
}

// Exec executes the deletion query.
func (ocdo *OutcomeClassDeleteOne) Exec(ctx context.Context) error {
	n, err := ocdo.ocd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeclass.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdo *OutcomeClassDeleteOne) ExecX(ctx context.Context) {
	ocdo.ocd.ExecX(ctx)
}
