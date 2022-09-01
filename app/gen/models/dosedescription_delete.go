// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// DoseDescriptionDelete is the builder for deleting a DoseDescription entity.
type DoseDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *DoseDescriptionMutation
}

// Where appends a list predicates to the DoseDescriptionDelete builder.
func (ddd *DoseDescriptionDelete) Where(ps ...predicate.DoseDescription) *DoseDescriptionDelete {
	ddd.mutation.Where(ps...)
	return ddd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ddd *DoseDescriptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ddd.hooks) == 0 {
		affected, err = ddd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoseDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ddd.mutation = mutation
			affected, err = ddd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ddd.hooks) - 1; i >= 0; i-- {
			if ddd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ddd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddd *DoseDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := ddd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ddd *DoseDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dosedescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dosedescription.FieldID,
			},
		},
	}
	if ps := ddd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ddd.driver, _spec)
}

// DoseDescriptionDeleteOne is the builder for deleting a single DoseDescription entity.
type DoseDescriptionDeleteOne struct {
	ddd *DoseDescriptionDelete
}

// Exec executes the deletion query.
func (dddo *DoseDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := dddo.ddd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dosedescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dddo *DoseDescriptionDeleteOne) ExecX(ctx context.Context) {
	dddo.ddd.ExecX(ctx)
}
