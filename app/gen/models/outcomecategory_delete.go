// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeCategoryDelete is the builder for deleting a OutcomeCategory entity.
type OutcomeCategoryDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeCategoryMutation
}

// Where appends a list predicates to the OutcomeCategoryDelete builder.
func (ocd *OutcomeCategoryDelete) Where(ps ...predicate.OutcomeCategory) *OutcomeCategoryDelete {
	ocd.mutation.Where(ps...)
	return ocd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocd *OutcomeCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocd.hooks) == 0 {
		affected, err = ocd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeCategoryMutation)
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
func (ocd *OutcomeCategoryDelete) ExecX(ctx context.Context) int {
	n, err := ocd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocd *OutcomeCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomecategory.FieldID,
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

// OutcomeCategoryDeleteOne is the builder for deleting a single OutcomeCategory entity.
type OutcomeCategoryDeleteOne struct {
	ocd *OutcomeCategoryDelete
}

// Exec executes the deletion query.
func (ocdo *OutcomeCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ocdo.ocd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomecategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdo *OutcomeCategoryDeleteOne) ExecX(ctx context.Context) {
	ocdo.ocd.ExecX(ctx)
}
