// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
)

// StudyEligibilityDelete is the builder for deleting a StudyEligibility entity.
type StudyEligibilityDelete struct {
	config
	hooks    []Hook
	mutation *StudyEligibilityMutation
}

// Where appends a list predicates to the StudyEligibilityDelete builder.
func (sed *StudyEligibilityDelete) Where(ps ...predicate.StudyEligibility) *StudyEligibilityDelete {
	sed.mutation.Where(ps...)
	return sed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sed *StudyEligibilityDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sed.hooks) == 0 {
		affected, err = sed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyEligibilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sed.mutation = mutation
			affected, err = sed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sed.hooks) - 1; i >= 0; i-- {
			if sed.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sed *StudyEligibilityDelete) ExecX(ctx context.Context) int {
	n, err := sed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sed *StudyEligibilityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: studyeligibility.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyeligibility.FieldID,
			},
		},
	}
	if ps := sed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sed.driver, _spec)
}

// StudyEligibilityDeleteOne is the builder for deleting a single StudyEligibility entity.
type StudyEligibilityDeleteOne struct {
	sed *StudyEligibilityDelete
}

// Exec executes the deletion query.
func (sedo *StudyEligibilityDeleteOne) Exec(ctx context.Context) error {
	n, err := sedo.sed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{studyeligibility.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sedo *StudyEligibilityDeleteOne) ExecX(ctx context.Context) {
	sedo.sed.ExecX(ctx)
}
