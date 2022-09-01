// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomDelete is the builder for deleting a OutcomeClassDenom entity.
type OutcomeClassDenomDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeClassDenomMutation
}

// Where appends a list predicates to the OutcomeClassDenomDelete builder.
func (ocdd *OutcomeClassDenomDelete) Where(ps ...predicate.OutcomeClassDenom) *OutcomeClassDenomDelete {
	ocdd.mutation.Where(ps...)
	return ocdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocdd *OutcomeClassDenomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocdd.hooks) == 0 {
		affected, err = ocdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocdd.mutation = mutation
			affected, err = ocdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocdd.hooks) - 1; i >= 0; i-- {
			if ocdd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdd *OutcomeClassDenomDelete) ExecX(ctx context.Context) int {
	n, err := ocdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocdd *OutcomeClassDenomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomeclassdenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenom.FieldID,
			},
		},
	}
	if ps := ocdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ocdd.driver, _spec)
}

// OutcomeClassDenomDeleteOne is the builder for deleting a single OutcomeClassDenom entity.
type OutcomeClassDenomDeleteOne struct {
	ocdd *OutcomeClassDenomDelete
}

// Exec executes the deletion query.
func (ocddo *OutcomeClassDenomDeleteOne) Exec(ctx context.Context) error {
	n, err := ocddo.ocdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomeclassdenom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocddo *OutcomeClassDenomDeleteOne) ExecX(ctx context.Context) {
	ocddo.ocdd.ExecX(ctx)
}
