// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ImmunocompromisedGroupsUpdate is the builder for updating ImmunocompromisedGroups entities.
type ImmunocompromisedGroupsUpdate struct {
	config
	hooks    []Hook
	mutation *ImmunocompromisedGroupsMutation
}

// Where appends a list predicates to the ImmunocompromisedGroupsUpdate builder.
func (igu *ImmunocompromisedGroupsUpdate) Where(ps ...predicate.ImmunocompromisedGroups) *ImmunocompromisedGroupsUpdate {
	igu.mutation.Where(ps...)
	return igu
}

// SetGroupName sets the "group_name" field.
func (igu *ImmunocompromisedGroupsUpdate) SetGroupName(s string) *ImmunocompromisedGroupsUpdate {
	igu.mutation.SetGroupName(s)
	return igu
}

// Mutation returns the ImmunocompromisedGroupsMutation object of the builder.
func (igu *ImmunocompromisedGroupsUpdate) Mutation() *ImmunocompromisedGroupsMutation {
	return igu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (igu *ImmunocompromisedGroupsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(igu.hooks) == 0 {
		if err = igu.check(); err != nil {
			return 0, err
		}
		affected, err = igu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImmunocompromisedGroupsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = igu.check(); err != nil {
				return 0, err
			}
			igu.mutation = mutation
			affected, err = igu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(igu.hooks) - 1; i >= 0; i-- {
			if igu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = igu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, igu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (igu *ImmunocompromisedGroupsUpdate) SaveX(ctx context.Context) int {
	affected, err := igu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (igu *ImmunocompromisedGroupsUpdate) Exec(ctx context.Context) error {
	_, err := igu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (igu *ImmunocompromisedGroupsUpdate) ExecX(ctx context.Context) {
	if err := igu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (igu *ImmunocompromisedGroupsUpdate) check() error {
	if v, ok := igu.mutation.GroupName(); ok {
		if err := immunocompromisedgroups.GroupNameValidator(v); err != nil {
			return &ValidationError{Name: "group_name", err: fmt.Errorf(`models: validator failed for field "ImmunocompromisedGroups.group_name": %w`, err)}
		}
	}
	return nil
}

func (igu *ImmunocompromisedGroupsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   immunocompromisedgroups.Table,
			Columns: immunocompromisedgroups.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: immunocompromisedgroups.FieldID,
			},
		},
	}
	if ps := igu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := igu.mutation.GroupName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: immunocompromisedgroups.FieldGroupName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, igu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{immunocompromisedgroups.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ImmunocompromisedGroupsUpdateOne is the builder for updating a single ImmunocompromisedGroups entity.
type ImmunocompromisedGroupsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImmunocompromisedGroupsMutation
}

// SetGroupName sets the "group_name" field.
func (iguo *ImmunocompromisedGroupsUpdateOne) SetGroupName(s string) *ImmunocompromisedGroupsUpdateOne {
	iguo.mutation.SetGroupName(s)
	return iguo
}

// Mutation returns the ImmunocompromisedGroupsMutation object of the builder.
func (iguo *ImmunocompromisedGroupsUpdateOne) Mutation() *ImmunocompromisedGroupsMutation {
	return iguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iguo *ImmunocompromisedGroupsUpdateOne) Select(field string, fields ...string) *ImmunocompromisedGroupsUpdateOne {
	iguo.fields = append([]string{field}, fields...)
	return iguo
}

// Save executes the query and returns the updated ImmunocompromisedGroups entity.
func (iguo *ImmunocompromisedGroupsUpdateOne) Save(ctx context.Context) (*ImmunocompromisedGroups, error) {
	var (
		err  error
		node *ImmunocompromisedGroups
	)
	if len(iguo.hooks) == 0 {
		if err = iguo.check(); err != nil {
			return nil, err
		}
		node, err = iguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImmunocompromisedGroupsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iguo.check(); err != nil {
				return nil, err
			}
			iguo.mutation = mutation
			node, err = iguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iguo.hooks) - 1; i >= 0; i-- {
			if iguo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = iguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iguo *ImmunocompromisedGroupsUpdateOne) SaveX(ctx context.Context) *ImmunocompromisedGroups {
	node, err := iguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iguo *ImmunocompromisedGroupsUpdateOne) Exec(ctx context.Context) error {
	_, err := iguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iguo *ImmunocompromisedGroupsUpdateOne) ExecX(ctx context.Context) {
	if err := iguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iguo *ImmunocompromisedGroupsUpdateOne) check() error {
	if v, ok := iguo.mutation.GroupName(); ok {
		if err := immunocompromisedgroups.GroupNameValidator(v); err != nil {
			return &ValidationError{Name: "group_name", err: fmt.Errorf(`models: validator failed for field "ImmunocompromisedGroups.group_name": %w`, err)}
		}
	}
	return nil
}

func (iguo *ImmunocompromisedGroupsUpdateOne) sqlSave(ctx context.Context) (_node *ImmunocompromisedGroups, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   immunocompromisedgroups.Table,
			Columns: immunocompromisedgroups.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: immunocompromisedgroups.FieldID,
			},
		},
	}
	id, ok := iguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "ImmunocompromisedGroups.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, immunocompromisedgroups.FieldID)
		for _, f := range fields {
			if !immunocompromisedgroups.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != immunocompromisedgroups.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iguo.mutation.GroupName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: immunocompromisedgroups.FieldGroupName,
		})
	}
	_node = &ImmunocompromisedGroups{config: iguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{immunocompromisedgroups.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
