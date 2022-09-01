// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ClinicalTrialDelete is the builder for deleting a ClinicalTrial entity.
type ClinicalTrialDelete struct {
	config
	hooks    []Hook
	mutation *ClinicalTrialMutation
}

// Where appends a list predicates to the ClinicalTrialDelete builder.
func (ctd *ClinicalTrialDelete) Where(ps ...predicate.ClinicalTrial) *ClinicalTrialDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *ClinicalTrialDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctd.hooks) == 0 {
		affected, err = ctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctd.mutation = mutation
			affected, err = ctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctd.hooks) - 1; i >= 0; i-- {
			if ctd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *ClinicalTrialDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *ClinicalTrialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: clinicaltrial.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrial.FieldID,
			},
		},
	}
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
}

// ClinicalTrialDeleteOne is the builder for deleting a single ClinicalTrial entity.
type ClinicalTrialDeleteOne struct {
	ctd *ClinicalTrialDelete
}

// Exec executes the deletion query.
func (ctdo *ClinicalTrialDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{clinicaltrial.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *ClinicalTrialDeleteOne) ExecX(ctx context.Context) {
	ctdo.ctd.ExecX(ctx)
}
