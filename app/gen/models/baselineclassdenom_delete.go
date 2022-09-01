// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomDelete is the builder for deleting a BaselineClassDenom entity.
type BaselineClassDenomDelete struct {
	config
	hooks    []Hook
	mutation *BaselineClassDenomMutation
}

// Where appends a list predicates to the BaselineClassDenomDelete builder.
func (bcdd *BaselineClassDenomDelete) Where(ps ...predicate.BaselineClassDenom) *BaselineClassDenomDelete {
	bcdd.mutation.Where(ps...)
	return bcdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcdd *BaselineClassDenomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcdd.hooks) == 0 {
		affected, err = bcdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcdd.mutation = mutation
			affected, err = bcdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcdd.hooks) - 1; i >= 0; i-- {
			if bcdd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdd *BaselineClassDenomDelete) ExecX(ctx context.Context) int {
	n, err := bcdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcdd *BaselineClassDenomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselineclassdenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenom.FieldID,
			},
		},
	}
	if ps := bcdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bcdd.driver, _spec)
}

// BaselineClassDenomDeleteOne is the builder for deleting a single BaselineClassDenom entity.
type BaselineClassDenomDeleteOne struct {
	bcdd *BaselineClassDenomDelete
}

// Exec executes the deletion query.
func (bcddo *BaselineClassDenomDeleteOne) Exec(ctx context.Context) error {
	n, err := bcddo.bcdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselineclassdenom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcddo *BaselineClassDenomDeleteOne) ExecX(ctx context.Context) {
	bcddo.bcdd.ExecX(ctx)
}
