// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasuresModuleDelete is the builder for deleting a OutcomeMeasuresModule entity.
type OutcomeMeasuresModuleDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasuresModuleMutation
}

// Where appends a list predicates to the OutcomeMeasuresModuleDelete builder.
func (ommd *OutcomeMeasuresModuleDelete) Where(ps ...predicate.OutcomeMeasuresModule) *OutcomeMeasuresModuleDelete {
	ommd.mutation.Where(ps...)
	return ommd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ommd *OutcomeMeasuresModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ommd.hooks) == 0 {
		affected, err = ommd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasuresModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ommd.mutation = mutation
			affected, err = ommd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ommd.hooks) - 1; i >= 0; i-- {
			if ommd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ommd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ommd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ommd *OutcomeMeasuresModuleDelete) ExecX(ctx context.Context) int {
	n, err := ommd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ommd *OutcomeMeasuresModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomemeasuresmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasuresmodule.FieldID,
			},
		},
	}
	if ps := ommd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ommd.driver, _spec)
}

// OutcomeMeasuresModuleDeleteOne is the builder for deleting a single OutcomeMeasuresModule entity.
type OutcomeMeasuresModuleDeleteOne struct {
	ommd *OutcomeMeasuresModuleDelete
}

// Exec executes the deletion query.
func (ommdo *OutcomeMeasuresModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := ommdo.ommd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomemeasuresmodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ommdo *OutcomeMeasuresModuleDeleteOne) ExecX(ctx context.Context) {
	ommdo.ommd.ExecX(ctx)
}
