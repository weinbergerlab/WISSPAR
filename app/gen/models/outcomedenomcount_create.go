// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
)

// OutcomeDenomCountCreate is the builder for creating a OutcomeDenomCount entity.
type OutcomeDenomCountCreate struct {
	config
	mutation *OutcomeDenomCountMutation
	hooks    []Hook
}

// SetOutcomeDenomCountGroupID sets the "outcome_denom_count_group_id" field.
func (odcc *OutcomeDenomCountCreate) SetOutcomeDenomCountGroupID(s string) *OutcomeDenomCountCreate {
	odcc.mutation.SetOutcomeDenomCountGroupID(s)
	return odcc
}

// SetOutcomeDenomCountValue sets the "outcome_denom_count_value" field.
func (odcc *OutcomeDenomCountCreate) SetOutcomeDenomCountValue(s string) *OutcomeDenomCountCreate {
	odcc.mutation.SetOutcomeDenomCountValue(s)
	return odcc
}

// SetParentID sets the "parent" edge to the OutcomeDenom entity by ID.
func (odcc *OutcomeDenomCountCreate) SetParentID(id int) *OutcomeDenomCountCreate {
	odcc.mutation.SetParentID(id)
	return odcc
}

// SetNillableParentID sets the "parent" edge to the OutcomeDenom entity by ID if the given value is not nil.
func (odcc *OutcomeDenomCountCreate) SetNillableParentID(id *int) *OutcomeDenomCountCreate {
	if id != nil {
		odcc = odcc.SetParentID(*id)
	}
	return odcc
}

// SetParent sets the "parent" edge to the OutcomeDenom entity.
func (odcc *OutcomeDenomCountCreate) SetParent(o *OutcomeDenom) *OutcomeDenomCountCreate {
	return odcc.SetParentID(o.ID)
}

// Mutation returns the OutcomeDenomCountMutation object of the builder.
func (odcc *OutcomeDenomCountCreate) Mutation() *OutcomeDenomCountMutation {
	return odcc.mutation
}

// Save creates the OutcomeDenomCount in the database.
func (odcc *OutcomeDenomCountCreate) Save(ctx context.Context) (*OutcomeDenomCount, error) {
	var (
		err  error
		node *OutcomeDenomCount
	)
	if len(odcc.hooks) == 0 {
		if err = odcc.check(); err != nil {
			return nil, err
		}
		node, err = odcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = odcc.check(); err != nil {
				return nil, err
			}
			odcc.mutation = mutation
			if node, err = odcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(odcc.hooks) - 1; i >= 0; i-- {
			if odcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (odcc *OutcomeDenomCountCreate) SaveX(ctx context.Context) *OutcomeDenomCount {
	v, err := odcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odcc *OutcomeDenomCountCreate) Exec(ctx context.Context) error {
	_, err := odcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcc *OutcomeDenomCountCreate) ExecX(ctx context.Context) {
	if err := odcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odcc *OutcomeDenomCountCreate) check() error {
	if _, ok := odcc.mutation.OutcomeDenomCountGroupID(); !ok {
		return &ValidationError{Name: "outcome_denom_count_group_id", err: errors.New(`models: missing required field "OutcomeDenomCount.outcome_denom_count_group_id"`)}
	}
	if _, ok := odcc.mutation.OutcomeDenomCountValue(); !ok {
		return &ValidationError{Name: "outcome_denom_count_value", err: errors.New(`models: missing required field "OutcomeDenomCount.outcome_denom_count_value"`)}
	}
	return nil
}

func (odcc *OutcomeDenomCountCreate) sqlSave(ctx context.Context) (*OutcomeDenomCount, error) {
	_node, _spec := odcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, odcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (odcc *OutcomeDenomCountCreate) createSpec() (*OutcomeDenomCount, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeDenomCount{config: odcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomedenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenomcount.FieldID,
			},
		}
	)
	if value, ok := odcc.mutation.OutcomeDenomCountGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountGroupID,
		})
		_node.OutcomeDenomCountGroupID = value
	}
	if value, ok := odcc.mutation.OutcomeDenomCountValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountValue,
		})
		_node.OutcomeDenomCountValue = value
	}
	if nodes := odcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenomcount.ParentTable,
			Columns: []string{outcomedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_denom_outcome_denom_count_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeDenomCountCreateBulk is the builder for creating many OutcomeDenomCount entities in bulk.
type OutcomeDenomCountCreateBulk struct {
	config
	builders []*OutcomeDenomCountCreate
}

// Save creates the OutcomeDenomCount entities in the database.
func (odccb *OutcomeDenomCountCreateBulk) Save(ctx context.Context) ([]*OutcomeDenomCount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(odccb.builders))
	nodes := make([]*OutcomeDenomCount, len(odccb.builders))
	mutators := make([]Mutator, len(odccb.builders))
	for i := range odccb.builders {
		func(i int, root context.Context) {
			builder := odccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeDenomCountMutation)
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
					_, err = mutators[i+1].Mutate(root, odccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, odccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, odccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (odccb *OutcomeDenomCountCreateBulk) SaveX(ctx context.Context) []*OutcomeDenomCount {
	v, err := odccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odccb *OutcomeDenomCountCreateBulk) Exec(ctx context.Context) error {
	_, err := odccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odccb *OutcomeDenomCountCreateBulk) ExecX(ctx context.Context) {
	if err := odccb.Exec(ctx); err != nil {
		panic(err)
	}
}
