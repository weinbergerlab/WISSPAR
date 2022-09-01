// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// PointOfContactDelete is the builder for deleting a PointOfContact entity.
type PointOfContactDelete struct {
	config
	hooks    []Hook
	mutation *PointOfContactMutation
}

// Where appends a list predicates to the PointOfContactDelete builder.
func (pocd *PointOfContactDelete) Where(ps ...predicate.PointOfContact) *PointOfContactDelete {
	pocd.mutation.Where(ps...)
	return pocd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pocd *PointOfContactDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pocd.hooks) == 0 {
		affected, err = pocd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PointOfContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pocd.mutation = mutation
			affected, err = pocd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pocd.hooks) - 1; i >= 0; i-- {
			if pocd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pocd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pocd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocd *PointOfContactDelete) ExecX(ctx context.Context) int {
	n, err := pocd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pocd *PointOfContactDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: pointofcontact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pointofcontact.FieldID,
			},
		},
	}
	if ps := pocd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pocd.driver, _spec)
}

// PointOfContactDeleteOne is the builder for deleting a single PointOfContact entity.
type PointOfContactDeleteOne struct {
	pocd *PointOfContactDelete
}

// Exec executes the deletion query.
func (pocdo *PointOfContactDeleteOne) Exec(ctx context.Context) error {
	n, err := pocdo.pocd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pointofcontact.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pocdo *PointOfContactDeleteOne) ExecX(ctx context.Context) {
	pocdo.pocd.ExecX(ctx)
}
