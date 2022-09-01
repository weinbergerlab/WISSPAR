// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ImmunocompromisedGroupsDelete is the builder for deleting a ImmunocompromisedGroups entity.
type ImmunocompromisedGroupsDelete struct {
	config
	hooks    []Hook
	mutation *ImmunocompromisedGroupsMutation
}

// Where appends a list predicates to the ImmunocompromisedGroupsDelete builder.
func (igd *ImmunocompromisedGroupsDelete) Where(ps ...predicate.ImmunocompromisedGroups) *ImmunocompromisedGroupsDelete {
	igd.mutation.Where(ps...)
	return igd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (igd *ImmunocompromisedGroupsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(igd.hooks) == 0 {
		affected, err = igd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImmunocompromisedGroupsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			igd.mutation = mutation
			affected, err = igd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(igd.hooks) - 1; i >= 0; i-- {
			if igd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = igd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, igd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (igd *ImmunocompromisedGroupsDelete) ExecX(ctx context.Context) int {
	n, err := igd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (igd *ImmunocompromisedGroupsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: immunocompromisedgroups.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: immunocompromisedgroups.FieldID,
			},
		},
	}
	if ps := igd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, igd.driver, _spec)
}

// ImmunocompromisedGroupsDeleteOne is the builder for deleting a single ImmunocompromisedGroups entity.
type ImmunocompromisedGroupsDeleteOne struct {
	igd *ImmunocompromisedGroupsDelete
}

// Exec executes the deletion query.
func (igdo *ImmunocompromisedGroupsDeleteOne) Exec(ctx context.Context) error {
	n, err := igdo.igd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{immunocompromisedgroups.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (igdo *ImmunocompromisedGroupsDeleteOne) ExecX(ctx context.Context) {
	igdo.igd.ExecX(ctx)
}
