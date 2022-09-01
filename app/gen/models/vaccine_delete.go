// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
)

// VaccineDelete is the builder for deleting a Vaccine entity.
type VaccineDelete struct {
	config
	hooks    []Hook
	mutation *VaccineMutation
}

// Where appends a list predicates to the VaccineDelete builder.
func (vd *VaccineDelete) Where(ps ...predicate.Vaccine) *VaccineDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VaccineDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VaccineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			if vd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VaccineDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VaccineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vaccine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vaccine.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VaccineDeleteOne is the builder for deleting a single Vaccine entity.
type VaccineDeleteOne struct {
	vd *VaccineDelete
}

// Exec executes the deletion query.
func (vdo *VaccineDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vaccine.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VaccineDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
