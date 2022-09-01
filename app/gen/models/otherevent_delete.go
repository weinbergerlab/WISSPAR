// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventDelete is the builder for deleting a OtherEvent entity.
type OtherEventDelete struct {
	config
	hooks    []Hook
	mutation *OtherEventMutation
}

// Where appends a list predicates to the OtherEventDelete builder.
func (oed *OtherEventDelete) Where(ps ...predicate.OtherEvent) *OtherEventDelete {
	oed.mutation.Where(ps...)
	return oed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oed *OtherEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oed.hooks) == 0 {
		affected, err = oed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oed.mutation = mutation
			affected, err = oed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oed.hooks) - 1; i >= 0; i-- {
			if oed.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oed *OtherEventDelete) ExecX(ctx context.Context) int {
	n, err := oed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oed *OtherEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: otherevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: otherevent.FieldID,
			},
		},
	}
	if ps := oed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, oed.driver, _spec)
}

// OtherEventDeleteOne is the builder for deleting a single OtherEvent entity.
type OtherEventDeleteOne struct {
	oed *OtherEventDelete
}

// Exec executes the deletion query.
func (oedo *OtherEventDeleteOne) Exec(ctx context.Context) error {
	n, err := oedo.oed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{otherevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oedo *OtherEventDeleteOne) ExecX(ctx context.Context) {
	oedo.oed.ExecX(ctx)
}
