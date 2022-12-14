// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/schedule"
)

// ScheduleDelete is the builder for deleting a Schedule entity.
type ScheduleDelete struct {
	config
	hooks    []Hook
	mutation *ScheduleMutation
}

// Where appends a list predicates to the ScheduleDelete builder.
func (sd *ScheduleDelete) Where(ps ...predicate.Schedule) *ScheduleDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ScheduleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sd.hooks) - 1; i >= 0; i-- {
			if sd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ScheduleDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ScheduleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: schedule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		},
	}
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// ScheduleDeleteOne is the builder for deleting a single Schedule entity.
type ScheduleDeleteOne struct {
	sd *ScheduleDelete
}

// Exec executes the deletion query.
func (sdo *ScheduleDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{schedule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ScheduleDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
