// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
)

// UseCaseUpdate is the builder for updating UseCase entities.
type UseCaseUpdate struct {
	config
	hooks    []Hook
	mutation *UseCaseMutation
}

// Where appends a list predicates to the UseCaseUpdate builder.
func (ucu *UseCaseUpdate) Where(ps ...predicate.UseCase) *UseCaseUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUseCaseName sets the "use_case_name" field.
func (ucu *UseCaseUpdate) SetUseCaseName(s string) *UseCaseUpdate {
	ucu.mutation.SetUseCaseName(s)
	return ucu
}

// SetUseCaseDescription sets the "use_case_description" field.
func (ucu *UseCaseUpdate) SetUseCaseDescription(s string) *UseCaseUpdate {
	ucu.mutation.SetUseCaseDescription(s)
	return ucu
}

// SetUseCaseCode sets the "use_case_code" field.
func (ucu *UseCaseUpdate) SetUseCaseCode(s string) *UseCaseUpdate {
	ucu.mutation.SetUseCaseCode(s)
	return ucu
}

// Mutation returns the UseCaseMutation object of the builder.
func (ucu *UseCaseUpdate) Mutation() *UseCaseMutation {
	return ucu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UseCaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucu.hooks) == 0 {
		affected, err = ucu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UseCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucu.mutation = mutation
			affected, err = ucu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucu.hooks) - 1; i >= 0; i-- {
			if ucu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ucu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UseCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UseCaseUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UseCaseUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucu *UseCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usecase.Table,
			Columns: usecase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usecase.FieldID,
			},
		},
	}
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.UseCaseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseName,
		})
	}
	if value, ok := ucu.mutation.UseCaseDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseDescription,
		})
	}
	if value, ok := ucu.mutation.UseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseCode,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usecase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UseCaseUpdateOne is the builder for updating a single UseCase entity.
type UseCaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UseCaseMutation
}

// SetUseCaseName sets the "use_case_name" field.
func (ucuo *UseCaseUpdateOne) SetUseCaseName(s string) *UseCaseUpdateOne {
	ucuo.mutation.SetUseCaseName(s)
	return ucuo
}

// SetUseCaseDescription sets the "use_case_description" field.
func (ucuo *UseCaseUpdateOne) SetUseCaseDescription(s string) *UseCaseUpdateOne {
	ucuo.mutation.SetUseCaseDescription(s)
	return ucuo
}

// SetUseCaseCode sets the "use_case_code" field.
func (ucuo *UseCaseUpdateOne) SetUseCaseCode(s string) *UseCaseUpdateOne {
	ucuo.mutation.SetUseCaseCode(s)
	return ucuo
}

// Mutation returns the UseCaseMutation object of the builder.
func (ucuo *UseCaseUpdateOne) Mutation() *UseCaseMutation {
	return ucuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UseCaseUpdateOne) Select(field string, fields ...string) *UseCaseUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UseCase entity.
func (ucuo *UseCaseUpdateOne) Save(ctx context.Context) (*UseCase, error) {
	var (
		err  error
		node *UseCase
	)
	if len(ucuo.hooks) == 0 {
		node, err = ucuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UseCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucuo.mutation = mutation
			node, err = ucuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucuo.hooks) - 1; i >= 0; i-- {
			if ucuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ucuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UseCaseUpdateOne) SaveX(ctx context.Context) *UseCase {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UseCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UseCaseUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucuo *UseCaseUpdateOne) sqlSave(ctx context.Context) (_node *UseCase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usecase.Table,
			Columns: usecase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usecase.FieldID,
			},
		},
	}
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "UseCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usecase.FieldID)
		for _, f := range fields {
			if !usecase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != usecase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.UseCaseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseName,
		})
	}
	if value, ok := ucuo.mutation.UseCaseDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseDescription,
		})
	}
	if value, ok := ucuo.mutation.UseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usecase.FieldUseCaseCode,
		})
	}
	_node = &UseCase{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usecase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
