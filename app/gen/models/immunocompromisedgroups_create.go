// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
)

// ImmunocompromisedGroupsCreate is the builder for creating a ImmunocompromisedGroups entity.
type ImmunocompromisedGroupsCreate struct {
	config
	mutation *ImmunocompromisedGroupsMutation
	hooks    []Hook
}

// SetGroupName sets the "group_name" field.
func (igc *ImmunocompromisedGroupsCreate) SetGroupName(s string) *ImmunocompromisedGroupsCreate {
	igc.mutation.SetGroupName(s)
	return igc
}

// Mutation returns the ImmunocompromisedGroupsMutation object of the builder.
func (igc *ImmunocompromisedGroupsCreate) Mutation() *ImmunocompromisedGroupsMutation {
	return igc.mutation
}

// Save creates the ImmunocompromisedGroups in the database.
func (igc *ImmunocompromisedGroupsCreate) Save(ctx context.Context) (*ImmunocompromisedGroups, error) {
	var (
		err  error
		node *ImmunocompromisedGroups
	)
	if len(igc.hooks) == 0 {
		if err = igc.check(); err != nil {
			return nil, err
		}
		node, err = igc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImmunocompromisedGroupsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = igc.check(); err != nil {
				return nil, err
			}
			igc.mutation = mutation
			if node, err = igc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(igc.hooks) - 1; i >= 0; i-- {
			if igc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = igc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, igc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (igc *ImmunocompromisedGroupsCreate) SaveX(ctx context.Context) *ImmunocompromisedGroups {
	v, err := igc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (igc *ImmunocompromisedGroupsCreate) Exec(ctx context.Context) error {
	_, err := igc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (igc *ImmunocompromisedGroupsCreate) ExecX(ctx context.Context) {
	if err := igc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (igc *ImmunocompromisedGroupsCreate) check() error {
	if _, ok := igc.mutation.GroupName(); !ok {
		return &ValidationError{Name: "group_name", err: errors.New(`models: missing required field "ImmunocompromisedGroups.group_name"`)}
	}
	if v, ok := igc.mutation.GroupName(); ok {
		if err := immunocompromisedgroups.GroupNameValidator(v); err != nil {
			return &ValidationError{Name: "group_name", err: fmt.Errorf(`models: validator failed for field "ImmunocompromisedGroups.group_name": %w`, err)}
		}
	}
	return nil
}

func (igc *ImmunocompromisedGroupsCreate) sqlSave(ctx context.Context) (*ImmunocompromisedGroups, error) {
	_node, _spec := igc.createSpec()
	if err := sqlgraph.CreateNode(ctx, igc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (igc *ImmunocompromisedGroupsCreate) createSpec() (*ImmunocompromisedGroups, *sqlgraph.CreateSpec) {
	var (
		_node = &ImmunocompromisedGroups{config: igc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: immunocompromisedgroups.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: immunocompromisedgroups.FieldID,
			},
		}
	)
	if value, ok := igc.mutation.GroupName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: immunocompromisedgroups.FieldGroupName,
		})
		_node.GroupName = value
	}
	return _node, _spec
}

// ImmunocompromisedGroupsCreateBulk is the builder for creating many ImmunocompromisedGroups entities in bulk.
type ImmunocompromisedGroupsCreateBulk struct {
	config
	builders []*ImmunocompromisedGroupsCreate
}

// Save creates the ImmunocompromisedGroups entities in the database.
func (igcb *ImmunocompromisedGroupsCreateBulk) Save(ctx context.Context) ([]*ImmunocompromisedGroups, error) {
	specs := make([]*sqlgraph.CreateSpec, len(igcb.builders))
	nodes := make([]*ImmunocompromisedGroups, len(igcb.builders))
	mutators := make([]Mutator, len(igcb.builders))
	for i := range igcb.builders {
		func(i int, root context.Context) {
			builder := igcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImmunocompromisedGroupsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, igcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, igcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, igcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (igcb *ImmunocompromisedGroupsCreateBulk) SaveX(ctx context.Context) []*ImmunocompromisedGroups {
	v, err := igcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (igcb *ImmunocompromisedGroupsCreateBulk) Exec(ctx context.Context) error {
	_, err := igcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (igcb *ImmunocompromisedGroupsCreateBulk) ExecX(ctx context.Context) {
	if err := igcb.Exec(ctx); err != nil {
		panic(err)
	}
}
