// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// CertainAgreementDelete is the builder for deleting a CertainAgreement entity.
type CertainAgreementDelete struct {
	config
	hooks    []Hook
	mutation *CertainAgreementMutation
}

// Where appends a list predicates to the CertainAgreementDelete builder.
func (cad *CertainAgreementDelete) Where(ps ...predicate.CertainAgreement) *CertainAgreementDelete {
	cad.mutation.Where(ps...)
	return cad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cad *CertainAgreementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cad.hooks) == 0 {
		affected, err = cad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertainAgreementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cad.mutation = mutation
			affected, err = cad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cad.hooks) - 1; i >= 0; i-- {
			if cad.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = cad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cad *CertainAgreementDelete) ExecX(ctx context.Context) int {
	n, err := cad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cad *CertainAgreementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: certainagreement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certainagreement.FieldID,
			},
		},
	}
	if ps := cad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cad.driver, _spec)
}

// CertainAgreementDeleteOne is the builder for deleting a single CertainAgreement entity.
type CertainAgreementDeleteOne struct {
	cad *CertainAgreementDelete
}

// Exec executes the deletion query.
func (cado *CertainAgreementDeleteOne) Exec(ctx context.Context) error {
	n, err := cado.cad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{certainagreement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cado *CertainAgreementDeleteOne) ExecX(ctx context.Context) {
	cado.cad.ExecX(ctx)
}
