// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasurementDelete is the builder for deleting a OutcomeMeasurement entity.
type OutcomeMeasurementDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasurementMutation
}

// Where appends a list predicates to the OutcomeMeasurementDelete builder.
func (omd *OutcomeMeasurementDelete) Where(ps ...predicate.OutcomeMeasurement) *OutcomeMeasurementDelete {
	omd.mutation.Where(ps...)
	return omd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omd *OutcomeMeasurementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(omd.hooks) == 0 {
		affected, err = omd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasurementMutation)
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
func (omd *OutcomeMeasurementDelete) ExecX(ctx context.Context) int {
	n, err := omd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omd *OutcomeMeasurementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomemeasurement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasurement.FieldID,
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

// OutcomeMeasurementDeleteOne is the builder for deleting a single OutcomeMeasurement entity.
type OutcomeMeasurementDeleteOne struct {
	omd *OutcomeMeasurementDelete
}

// Exec executes the deletion query.
func (omdo *OutcomeMeasurementDeleteOne) Exec(ctx context.Context) error {
	n, err := omdo.omd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomemeasurement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omdo *OutcomeMeasurementDeleteOne) ExecX(ctx context.Context) {
	omdo.omd.ExecX(ctx)
}
