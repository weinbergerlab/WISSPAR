// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// DoseDescriptionUpdate is the builder for updating DoseDescription entities.
type DoseDescriptionUpdate struct {
	config
	hooks    []Hook
	mutation *DoseDescriptionMutation
}

// Where appends a list predicates to the DoseDescriptionUpdate builder.
func (ddu *DoseDescriptionUpdate) Where(ps ...predicate.DoseDescription) *DoseDescriptionUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetName sets the "name" field.
func (ddu *DoseDescriptionUpdate) SetName(s string) *DoseDescriptionUpdate {
	ddu.mutation.SetName(s)
	return ddu
}

// Mutation returns the DoseDescriptionMutation object of the builder.
func (ddu *DoseDescriptionUpdate) Mutation() *DoseDescriptionMutation {
	return ddu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DoseDescriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ddu.hooks) == 0 {
		if err = ddu.check(); err != nil {
			return 0, err
		}
		affected, err = ddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoseDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddu.check(); err != nil {
				return 0, err
			}
			ddu.mutation = mutation
			affected, err = ddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ddu.hooks) - 1; i >= 0; i-- {
			if ddu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DoseDescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DoseDescriptionUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DoseDescriptionUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddu *DoseDescriptionUpdate) check() error {
	if v, ok := ddu.mutation.Name(); ok {
		if err := dosedescription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "DoseDescription.name": %w`, err)}
		}
	}
	return nil
}

func (ddu *DoseDescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dosedescription.Table,
			Columns: dosedescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dosedescription.FieldID,
			},
		},
	}
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dosedescription.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dosedescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DoseDescriptionUpdateOne is the builder for updating a single DoseDescription entity.
type DoseDescriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DoseDescriptionMutation
}

// SetName sets the "name" field.
func (dduo *DoseDescriptionUpdateOne) SetName(s string) *DoseDescriptionUpdateOne {
	dduo.mutation.SetName(s)
	return dduo
}

// Mutation returns the DoseDescriptionMutation object of the builder.
func (dduo *DoseDescriptionUpdateOne) Mutation() *DoseDescriptionMutation {
	return dduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DoseDescriptionUpdateOne) Select(field string, fields ...string) *DoseDescriptionUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DoseDescription entity.
func (dduo *DoseDescriptionUpdateOne) Save(ctx context.Context) (*DoseDescription, error) {
	var (
		err  error
		node *DoseDescription
	)
	if len(dduo.hooks) == 0 {
		if err = dduo.check(); err != nil {
			return nil, err
		}
		node, err = dduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoseDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dduo.check(); err != nil {
				return nil, err
			}
			dduo.mutation = mutation
			node, err = dduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dduo.hooks) - 1; i >= 0; i-- {
			if dduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = dduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DoseDescriptionUpdateOne) SaveX(ctx context.Context) *DoseDescription {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DoseDescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DoseDescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dduo *DoseDescriptionUpdateOne) check() error {
	if v, ok := dduo.mutation.Name(); ok {
		if err := dosedescription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "DoseDescription.name": %w`, err)}
		}
	}
	return nil
}

func (dduo *DoseDescriptionUpdateOne) sqlSave(ctx context.Context) (_node *DoseDescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dosedescription.Table,
			Columns: dosedescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dosedescription.FieldID,
			},
		},
	}
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "DoseDescription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dosedescription.FieldID)
		for _, f := range fields {
			if !dosedescription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != dosedescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dosedescription.FieldName,
		})
	}
	_node = &DoseDescription{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dosedescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
