// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasureDelete is the builder for deleting a OutcomeMeasure entity.
type OutcomeMeasureDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasureMutation
}

// Where appends a list predicates to the OutcomeMeasureDelete builder.
func (omd *OutcomeMeasureDelete) Where(ps ...predicate.OutcomeMeasure) *OutcomeMeasureDelete {
	omd.mutation.Where(ps...)
	return omd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omd *OutcomeMeasureDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(omd.hooks) == 0 {
		affected, err = omd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omd.mutation = mutation
			affected, err = omd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(omd.hooks) - 1; i >= 0; i-- {
			if omd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (omd *OutcomeMeasureDelete) ExecX(ctx context.Context) int {
	n, err := omd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omd *OutcomeMeasureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomemeasure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasure.FieldID,
			},
		},
	}
	if ps := omd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, omd.driver, _spec)
}

// OutcomeMeasureDeleteOne is the builder for deleting a single OutcomeMeasure entity.
type OutcomeMeasureDeleteOne struct {
	omd *OutcomeMeasureDelete
}

// Exec executes the deletion query.
func (omdo *OutcomeMeasureDeleteOne) Exec(ctx context.Context) error {
	n, err := omdo.omd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomemeasure.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omdo *OutcomeMeasureDeleteOne) ExecX(ctx context.Context) {
	omdo.omd.ExecX(ctx)
}
