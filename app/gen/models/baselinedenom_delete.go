// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomDelete is the builder for deleting a BaselineDenom entity.
type BaselineDenomDelete struct {
	config
	hooks    []Hook
	mutation *BaselineDenomMutation
}

// Where appends a list predicates to the BaselineDenomDelete builder.
func (bdd *BaselineDenomDelete) Where(ps ...predicate.BaselineDenom) *BaselineDenomDelete {
	bdd.mutation.Where(ps...)
	return bdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bdd *BaselineDenomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bdd.hooks) == 0 {
		affected, err = bdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdd.mutation = mutation
			affected, err = bdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bdd.hooks) - 1; i >= 0; i-- {
			if bdd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdd *BaselineDenomDelete) ExecX(ctx context.Context) int {
	n, err := bdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bdd *BaselineDenomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinedenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenom.FieldID,
			},
		},
	}
	if ps := bdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bdd.driver, _spec)
}

// BaselineDenomDeleteOne is the builder for deleting a single BaselineDenom entity.
type BaselineDenomDeleteOne struct {
	bdd *BaselineDenomDelete
}

// Exec executes the deletion query.
func (bddo *BaselineDenomDeleteOne) Exec(ctx context.Context) error {
	n, err := bddo.bdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinedenom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bddo *BaselineDenomDeleteOne) ExecX(ctx context.Context) {
	bddo.bdd.ExecX(ctx)
}
