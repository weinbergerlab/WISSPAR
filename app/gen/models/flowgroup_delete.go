// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowGroupDelete is the builder for deleting a FlowGroup entity.
type FlowGroupDelete struct {
	config
	hooks    []Hook
	mutation *FlowGroupMutation
}

// Where appends a list predicates to the FlowGroupDelete builder.
func (fgd *FlowGroupDelete) Where(ps ...predicate.FlowGroup) *FlowGroupDelete {
	fgd.mutation.Where(ps...)
	return fgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fgd *FlowGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fgd.hooks) == 0 {
		affected, err = fgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fgd.mutation = mutation
			affected, err = fgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fgd.hooks) - 1; i >= 0; i-- {
			if fgd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgd *FlowGroupDelete) ExecX(ctx context.Context) int {
	n, err := fgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fgd *FlowGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowgroup.FieldID,
			},
		},
	}
	if ps := fgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fgd.driver, _spec)
}

// FlowGroupDeleteOne is the builder for deleting a single FlowGroup entity.
type FlowGroupDeleteOne struct {
	fgd *FlowGroupDelete
}

// Exec executes the deletion query.
func (fgdo *FlowGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := fgdo.fgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fgdo *FlowGroupDeleteOne) ExecX(ctx context.Context) {
	fgdo.fgd.ExecX(ctx)
}
