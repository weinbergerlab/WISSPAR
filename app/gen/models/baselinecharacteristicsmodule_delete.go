// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineCharacteristicsModuleDelete is the builder for deleting a BaselineCharacteristicsModule entity.
type BaselineCharacteristicsModuleDelete struct {
	config
	hooks    []Hook
	mutation *BaselineCharacteristicsModuleMutation
}

// Where appends a list predicates to the BaselineCharacteristicsModuleDelete builder.
func (bcmd *BaselineCharacteristicsModuleDelete) Where(ps ...predicate.BaselineCharacteristicsModule) *BaselineCharacteristicsModuleDelete {
	bcmd.mutation.Where(ps...)
	return bcmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcmd *BaselineCharacteristicsModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcmd.hooks) == 0 {
		affected, err = bcmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCharacteristicsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcmd.mutation = mutation
			affected, err = bcmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcmd.hooks) - 1; i >= 0; i-- {
			if bcmd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmd *BaselineCharacteristicsModuleDelete) ExecX(ctx context.Context) int {
	n, err := bcmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcmd *BaselineCharacteristicsModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baselinecharacteristicsmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecharacteristicsmodule.FieldID,
			},
		},
	}
	if ps := bcmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bcmd.driver, _spec)
}

// BaselineCharacteristicsModuleDeleteOne is the builder for deleting a single BaselineCharacteristicsModule entity.
type BaselineCharacteristicsModuleDeleteOne struct {
	bcmd *BaselineCharacteristicsModuleDelete
}

// Exec executes the deletion query.
func (bcmdo *BaselineCharacteristicsModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := bcmdo.bcmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baselinecharacteristicsmodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmdo *BaselineCharacteristicsModuleDeleteOne) ExecX(ctx context.Context) {
	bcmdo.bcmd.ExecX(ctx)
}
