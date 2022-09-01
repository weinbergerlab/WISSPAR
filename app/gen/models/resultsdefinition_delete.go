// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ResultsDefinitionDelete is the builder for deleting a ResultsDefinition entity.
type ResultsDefinitionDelete struct {
	config
	hooks    []Hook
	mutation *ResultsDefinitionMutation
}

// Where appends a list predicates to the ResultsDefinitionDelete builder.
func (rdd *ResultsDefinitionDelete) Where(ps ...predicate.ResultsDefinition) *ResultsDefinitionDelete {
	rdd.mutation.Where(ps...)
	return rdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rdd *ResultsDefinitionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rdd.hooks) == 0 {
		affected, err = rdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rdd.mutation = mutation
			affected, err = rdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rdd.hooks) - 1; i >= 0; i-- {
			if rdd.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = rdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdd *ResultsDefinitionDelete) ExecX(ctx context.Context) int {
	n, err := rdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rdd *ResultsDefinitionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: resultsdefinition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultsdefinition.FieldID,
			},
		},
	}
	if ps := rdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rdd.driver, _spec)
}

// ResultsDefinitionDeleteOne is the builder for deleting a single ResultsDefinition entity.
type ResultsDefinitionDeleteOne struct {
	rdd *ResultsDefinitionDelete
}

// Exec executes the deletion query.
func (rddo *ResultsDefinitionDeleteOne) Exec(ctx context.Context) error {
	n, err := rddo.rdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{resultsdefinition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rddo *ResultsDefinitionDeleteOne) ExecX(ctx context.Context) {
	rddo.rdd.ExecX(ctx)
}
