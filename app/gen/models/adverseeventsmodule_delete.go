// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// AdverseEventsModuleDelete is the builder for deleting a AdverseEventsModule entity.
type AdverseEventsModuleDelete struct {
	config
	hooks    []Hook
	mutation *AdverseEventsModuleMutation
}

// Where appends a list predicates to the AdverseEventsModuleDelete builder.
func (aemd *AdverseEventsModuleDelete) Where(ps ...predicate.AdverseEventsModule) *AdverseEventsModuleDelete {
	aemd.mutation.Where(ps...)
	return aemd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aemd *AdverseEventsModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aemd.hooks) == 0 {
		affected, err = aemd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdverseEventsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aemd.mutation = mutation
			affected, err = aemd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aemd.hooks) - 1; i >= 0; i-- {
			if aemd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = aemd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aemd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aemd *AdverseEventsModuleDelete) ExecX(ctx context.Context) int {
	n, err := aemd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aemd *AdverseEventsModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: adverseeventsmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adverseeventsmodule.FieldID,
			},
		},
	}
	if ps := aemd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aemd.driver, _spec)
}

// AdverseEventsModuleDeleteOne is the builder for deleting a single AdverseEventsModule entity.
type AdverseEventsModuleDeleteOne struct {
	aemd *AdverseEventsModuleDelete
}

// Exec executes the deletion query.
func (aemdo *AdverseEventsModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := aemdo.aemd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adverseeventsmodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aemdo *AdverseEventsModuleDeleteOne) ExecX(ctx context.Context) {
	aemdo.aemd.ExecX(ctx)
}
