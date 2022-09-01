// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// EventGroupDelete is the builder for deleting a EventGroup entity.
type EventGroupDelete struct {
	config
	hooks    []Hook
	mutation *EventGroupMutation
}

// Where appends a list predicates to the EventGroupDelete builder.
func (egd *EventGroupDelete) Where(ps ...predicate.EventGroup) *EventGroupDelete {
	egd.mutation.Where(ps...)
	return egd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (egd *EventGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(egd.hooks) == 0 {
		affected, err = egd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			egd.mutation = mutation
			affected, err = egd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(egd.hooks) - 1; i >= 0; i-- {
			if egd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = egd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, egd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (egd *EventGroupDelete) ExecX(ctx context.Context) int {
	n, err := egd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (egd *EventGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: eventgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventgroup.FieldID,
			},
		},
	}
	if ps := egd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, egd.driver, _spec)
}

// EventGroupDeleteOne is the builder for deleting a single EventGroup entity.
type EventGroupDeleteOne struct {
	egd *EventGroupDelete
}

// Exec executes the deletion query.
func (egdo *EventGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := egdo.egd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{eventgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (egdo *EventGroupDeleteOne) ExecX(ctx context.Context) {
	egdo.egd.ExecX(ctx)
}
