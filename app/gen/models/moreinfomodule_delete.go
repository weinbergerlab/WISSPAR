// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// MoreInfoModuleDelete is the builder for deleting a MoreInfoModule entity.
type MoreInfoModuleDelete struct {
	config
	hooks    []Hook
	mutation *MoreInfoModuleMutation
}

// Where appends a list predicates to the MoreInfoModuleDelete builder.
func (mimd *MoreInfoModuleDelete) Where(ps ...predicate.MoreInfoModule) *MoreInfoModuleDelete {
	mimd.mutation.Where(ps...)
	return mimd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mimd *MoreInfoModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mimd.hooks) == 0 {
		affected, err = mimd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoreInfoModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mimd.mutation = mutation
			affected, err = mimd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mimd.hooks) - 1; i >= 0; i-- {
			if mimd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = mimd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mimd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mimd *MoreInfoModuleDelete) ExecX(ctx context.Context) int {
	n, err := mimd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mimd *MoreInfoModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: moreinfomodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moreinfomodule.FieldID,
			},
		},
	}
	if ps := mimd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mimd.driver, _spec)
}

// MoreInfoModuleDeleteOne is the builder for deleting a single MoreInfoModule entity.
type MoreInfoModuleDeleteOne struct {
	mimd *MoreInfoModuleDelete
}

// Exec executes the deletion query.
func (mimdo *MoreInfoModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := mimdo.mimd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{moreinfomodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mimdo *MoreInfoModuleDeleteOne) ExecX(ctx context.Context) {
	mimdo.mimd.ExecX(ctx)
}
