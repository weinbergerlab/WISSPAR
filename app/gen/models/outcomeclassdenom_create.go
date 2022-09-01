// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
)

// OutcomeClassDenomCreate is the builder for creating a OutcomeClassDenom entity.
type OutcomeClassDenomCreate struct {
	config
	mutation *OutcomeClassDenomMutation
	hooks    []Hook
}

// SetOutcomeClassDenomUnits sets the "outcome_class_denom_units" field.
func (ocdc *OutcomeClassDenomCreate) SetOutcomeClassDenomUnits(s string) *OutcomeClassDenomCreate {
	ocdc.mutation.SetOutcomeClassDenomUnits(s)
	return ocdc
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (ocdc *OutcomeClassDenomCreate) SetParentID(id int) *OutcomeClassDenomCreate {
	ocdc.mutation.SetParentID(id)
	return ocdc
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (ocdc *OutcomeClassDenomCreate) SetNillableParentID(id *int) *OutcomeClassDenomCreate {
	if id != nil {
		ocdc = ocdc.SetParentID(*id)
	}
	return ocdc
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (ocdc *OutcomeClassDenomCreate) SetParent(o *OutcomeClass) *OutcomeClassDenomCreate {
	return ocdc.SetParentID(o.ID)
}

// AddOutcomeClassDenomCountListIDs adds the "outcome_class_denom_count_list" edge to the OutcomeClassDenomCount entity by IDs.
func (ocdc *OutcomeClassDenomCreate) AddOutcomeClassDenomCountListIDs(ids ...int) *OutcomeClassDenomCreate {
	ocdc.mutation.AddOutcomeClassDenomCountListIDs(ids...)
	return ocdc
}

// AddOutcomeClassDenomCountList adds the "outcome_class_denom_count_list" edges to the OutcomeClassDenomCount entity.
func (ocdc *OutcomeClassDenomCreate) AddOutcomeClassDenomCountList(o ...*OutcomeClassDenomCount) *OutcomeClassDenomCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocdc.AddOutcomeClassDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeClassDenomMutation object of the builder.
func (ocdc *OutcomeClassDenomCreate) Mutation() *OutcomeClassDenomMutation {
	return ocdc.mutation
}

// Save creates the OutcomeClassDenom in the database.
func (ocdc *OutcomeClassDenomCreate) Save(ctx context.Context) (*OutcomeClassDenom, error) {
	var (
		err  error
		node *OutcomeClassDenom
	)
	if len(ocdc.hooks) == 0 {
		if err = ocdc.check(); err != nil {
			return nil, err
		}
		node, err = ocdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ocdc.check(); err != nil {
				return nil, err
			}
			ocdc.mutation = mutation
			if node, err = ocdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ocdc.hooks) - 1; i >= 0; i-- {
			if ocdc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ocdc *OutcomeClassDenomCreate) SaveX(ctx context.Context) *OutcomeClassDenom {
	v, err := ocdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocdc *OutcomeClassDenomCreate) Exec(ctx context.Context) error {
	_, err := ocdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdc *OutcomeClassDenomCreate) ExecX(ctx context.Context) {
	if err := ocdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocdc *OutcomeClassDenomCreate) check() error {
	if _, ok := ocdc.mutation.OutcomeClassDenomUnits(); !ok {
		return &ValidationError{Name: "outcome_class_denom_units", err: errors.New(`models: missing required field "OutcomeClassDenom.outcome_class_denom_units"`)}
	}
	return nil
}

func (ocdc *OutcomeClassDenomCreate) sqlSave(ctx context.Context) (*OutcomeClassDenom, error) {
	_node, _spec := ocdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ocdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ocdc *OutcomeClassDenomCreate) createSpec() (*OutcomeClassDenom, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeClassDenom{config: ocdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeclassdenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenom.FieldID,
			},
		}
	)
	if value, ok := ocdc.mutation.OutcomeClassDenomUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenom.FieldOutcomeClassDenomUnits,
		})
		_node.OutcomeClassDenomUnits = value
	}
	if nodes := ocdc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenom.ParentTable,
			Columns: []string{outcomeclassdenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_class_outcome_class_denom_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ocdc.mutation.OutcomeClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeClassDenomCreateBulk is the builder for creating many OutcomeClassDenom entities in bulk.
type OutcomeClassDenomCreateBulk struct {
	config
	builders []*OutcomeClassDenomCreate
}

// Save creates the OutcomeClassDenom entities in the database.
func (ocdcb *OutcomeClassDenomCreateBulk) Save(ctx context.Context) ([]*OutcomeClassDenom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocdcb.builders))
	nodes := make([]*OutcomeClassDenom, len(ocdcb.builders))
	mutators := make([]Mutator, len(ocdcb.builders))
	for i := range ocdcb.builders {
		func(i int, root context.Context) {
			builder := ocdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeClassDenomMutation)
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
					_, err = mutators[i+1].Mutate(root, ocdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocdcb *OutcomeClassDenomCreateBulk) SaveX(ctx context.Context) []*OutcomeClassDenom {
	v, err := ocdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocdcb *OutcomeClassDenomCreateBulk) Exec(ctx context.Context) error {
	_, err := ocdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcb *OutcomeClassDenomCreateBulk) ExecX(ctx context.Context) {
	if err := ocdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
