// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ClinicalTrialMetadataDelete is the builder for deleting a ClinicalTrialMetadata entity.
type ClinicalTrialMetadataDelete struct {
	config
	hooks    []Hook
	mutation *ClinicalTrialMetadataMutation
}

// Where appends a list predicates to the ClinicalTrialMetadataDelete builder.
func (ctmd *ClinicalTrialMetadataDelete) Where(ps ...predicate.ClinicalTrialMetadata) *ClinicalTrialMetadataDelete {
	ctmd.mutation.Where(ps...)
	return ctmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctmd *ClinicalTrialMetadataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctmd.hooks) == 0 {
		affected, err = ctmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctmd.mutation = mutation
			affected, err = ctmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctmd.hooks) - 1; i >= 0; i-- {
			if ctmd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmd *ClinicalTrialMetadataDelete) ExecX(ctx context.Context) int {
	n, err := ctmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctmd *ClinicalTrialMetadataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: clinicaltrialmetadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrialmetadata.FieldID,
			},
		},
	}
	if ps := ctmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ctmd.driver, _spec)
}

// ClinicalTrialMetadataDeleteOne is the builder for deleting a single ClinicalTrialMetadata entity.
type ClinicalTrialMetadataDeleteOne struct {
	ctmd *ClinicalTrialMetadataDelete
}

// Exec executes the deletion query.
func (ctmdo *ClinicalTrialMetadataDeleteOne) Exec(ctx context.Context) error {
	n, err := ctmdo.ctmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{clinicaltrialmetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmdo *ClinicalTrialMetadataDeleteOne) ExecX(ctx context.Context) {
	ctmdo.ctmd.ExecX(ctx)
}
