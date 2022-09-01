// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
)

// OutcomeClassDenomCountCreate is the builder for creating a OutcomeClassDenomCount entity.
type OutcomeClassDenomCountCreate struct {
	config
	mutation *OutcomeClassDenomCountMutation
	hooks    []Hook
}

// SetOutcomeClassDenomCountGroupID sets the "outcome_class_denom_count_group_id" field.
func (ocdcc *OutcomeClassDenomCountCreate) SetOutcomeClassDenomCountGroupID(s string) *OutcomeClassDenomCountCreate {
	ocdcc.mutation.SetOutcomeClassDenomCountGroupID(s)
	return ocdcc
}

// SetOutcomeClassDenomCountValue sets the "outcome_class_denom_count_value" field.
func (ocdcc *OutcomeClassDenomCountCreate) SetOutcomeClassDenomCountValue(s string) *OutcomeClassDenomCountCreate {
	ocdcc.mutation.SetOutcomeClassDenomCountValue(s)
	return ocdcc
}

// SetParentID sets the "parent" edge to the OutcomeClassDenom entity by ID.
func (ocdcc *OutcomeClassDenomCountCreate) SetParentID(id int) *OutcomeClassDenomCountCreate {
	ocdcc.mutation.SetParentID(id)
	return ocdcc
}

// SetNillableParentID sets the "parent" edge to the OutcomeClassDenom entity by ID if the given value is not nil.
func (ocdcc *OutcomeClassDenomCountCreate) SetNillableParentID(id *int) *OutcomeClassDenomCountCreate {
	if id != nil {
		ocdcc = ocdcc.SetParentID(*id)
	}
	return ocdcc
}

// SetParent sets the "parent" edge to the OutcomeClassDenom entity.
func (ocdcc *OutcomeClassDenomCountCreate) SetParent(o *OutcomeClassDenom) *OutcomeClassDenomCountCreate {
	return ocdcc.SetParentID(o.ID)
}

// Mutation returns the OutcomeClassDenomCountMutation object of the builder.
func (ocdcc *OutcomeClassDenomCountCreate) Mutation() *OutcomeClassDenomCountMutation {
	return ocdcc.mutation
}

// Save creates the OutcomeClassDenomCount in the database.
func (ocdcc *OutcomeClassDenomCountCreate) Save(ctx context.Context) (*OutcomeClassDenomCount, error) {
	var (
		err  error
		node *OutcomeClassDenomCount
	)
	if len(ocdcc.hooks) == 0 {
		if err = ocdcc.check(); err != nil {
			return nil, err
		}
		node, err = ocdcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ocdcc.check(); err != nil {
				return nil, err
			}
			ocdcc.mutation = mutation
			if node, err = ocdcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ocdcc.hooks) - 1; i >= 0; i-- {
			if ocdcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ocdcc *OutcomeClassDenomCountCreate) SaveX(ctx context.Context) *OutcomeClassDenomCount {
	v, err := ocdcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocdcc *OutcomeClassDenomCountCreate) Exec(ctx context.Context) error {
	_, err := ocdcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcc *OutcomeClassDenomCountCreate) ExecX(ctx context.Context) {
	if err := ocdcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocdcc *OutcomeClassDenomCountCreate) check() error {
	if _, ok := ocdcc.mutation.OutcomeClassDenomCountGroupID(); !ok {
		return &ValidationError{Name: "outcome_class_denom_count_group_id", err: errors.New(`models: missing required field "OutcomeClassDenomCount.outcome_class_denom_count_group_id"`)}
	}
	if _, ok := ocdcc.mutation.OutcomeClassDenomCountValue(); !ok {
		return &ValidationError{Name: "outcome_class_denom_count_value", err: errors.New(`models: missing required field "OutcomeClassDenomCount.outcome_class_denom_count_value"`)}
	}
	return nil
}

func (ocdcc *OutcomeClassDenomCountCreate) sqlSave(ctx context.Context) (*OutcomeClassDenomCount, error) {
	_node, _spec := ocdcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ocdcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ocdcc *OutcomeClassDenomCountCreate) createSpec() (*OutcomeClassDenomCount, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeClassDenomCount{config: ocdcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeclassdenomcount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenomcount.FieldID,
			},
		}
	)
	if value, ok := ocdcc.mutation.OutcomeClassDenomCountGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID,
		})
		_node.OutcomeClassDenomCountGroupID = value
	}
	if value, ok := ocdcc.mutation.OutcomeClassDenomCountValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountValue,
		})
		_node.OutcomeClassDenomCountValue = value
	}
	if nodes := ocdcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenomcount.ParentTable,
			Columns: []string{outcomeclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_class_denom_outcome_class_denom_count_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeClassDenomCountCreateBulk is the builder for creating many OutcomeClassDenomCount entities in bulk.
type OutcomeClassDenomCountCreateBulk struct {
	config
	builders []*OutcomeClassDenomCountCreate
}

// Save creates the OutcomeClassDenomCount entities in the database.
func (ocdccb *OutcomeClassDenomCountCreateBulk) Save(ctx context.Context) ([]*OutcomeClassDenomCount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocdccb.builders))
	nodes := make([]*OutcomeClassDenomCount, len(ocdccb.builders))
	mutators := make([]Mutator, len(ocdccb.builders))
	for i := range ocdccb.builders {
		func(i int, root context.Context) {
			builder := ocdccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeClassDenomCountMutation)
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
					_, err = mutators[i+1].Mutate(root, ocdccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocdccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocdccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocdccb *OutcomeClassDenomCountCreateBulk) SaveX(ctx context.Context) []*OutcomeClassDenomCount {
	v, err := ocdccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocdccb *OutcomeClassDenomCountCreateBulk) Exec(ctx context.Context) error {
	_, err := ocdccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdccb *OutcomeClassDenomCountCreateBulk) ExecX(ctx context.Context) {
	if err := ocdccb.Exec(ctx); err != nil {
		panic(err)
	}
}
