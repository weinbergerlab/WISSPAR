// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// StudyLocationDelete is the builder for deleting a StudyLocation entity.
type StudyLocationDelete struct {
	config
	hooks    []Hook
	mutation *StudyLocationMutation
}

// Where appends a list predicates to the StudyLocationDelete builder.
func (sld *StudyLocationDelete) Where(ps ...predicate.StudyLocation) *StudyLocationDelete {
	sld.mutation.Where(ps...)
	return sld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sld *StudyLocationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sld.hooks) == 0 {
		affected, err = sld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sld.mutation = mutation
			affected, err = sld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sld.hooks) - 1; i >= 0; i-- {
			if sld.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sld *StudyLocationDelete) ExecX(ctx context.Context) int {
	n, err := sld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sld *StudyLocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: studylocation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studylocation.FieldID,
			},
		},
	}
	if ps := sld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sld.driver, _spec)
}

// StudyLocationDeleteOne is the builder for deleting a single StudyLocation entity.
type StudyLocationDeleteOne struct {
	sld *StudyLocationDelete
}

// Exec executes the deletion query.
func (sldo *StudyLocationDeleteOne) Exec(ctx context.Context) error {
	n, err := sldo.sld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{studylocation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sldo *StudyLocationDeleteOne) ExecX(ctx context.Context) {
	sldo.sld.ExecX(ctx)
}
