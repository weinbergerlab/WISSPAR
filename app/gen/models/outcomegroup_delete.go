// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeGroupDelete is the builder for deleting a OutcomeGroup entity.
type OutcomeGroupDelete struct {
	config
	hooks    []Hook
	mutation *OutcomeGroupMutation
}

// Where appends a list predicates to the OutcomeGroupDelete builder.
func (ogd *OutcomeGroupDelete) Where(ps ...predicate.OutcomeGroup) *OutcomeGroupDelete {
	ogd.mutation.Where(ps...)
	return ogd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ogd *OutcomeGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ogd.hooks) == 0 {
		affected, err = ogd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ogd.mutation = mutation
			affected, err = ogd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ogd.hooks) - 1; i >= 0; i-- {
			if ogd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ogd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogd *OutcomeGroupDelete) ExecX(ctx context.Context) int {
	n, err := ogd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ogd *OutcomeGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outcomegroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomegroup.FieldID,
			},
		},
	}
	if ps := ogd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ogd.driver, _spec)
}

// OutcomeGroupDeleteOne is the builder for deleting a single OutcomeGroup entity.
type OutcomeGroupDeleteOne struct {
	ogd *OutcomeGroupDelete
}

// Exec executes the deletion query.
func (ogdo *OutcomeGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := ogdo.ogd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outcomegroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ogdo *OutcomeGroupDeleteOne) ExecX(ctx context.Context) {
	ogdo.ogd.ExecX(ctx)
}
