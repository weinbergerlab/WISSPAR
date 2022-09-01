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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeDenomCreate is the builder for creating a OutcomeDenom entity.
type OutcomeDenomCreate struct {
	config
	mutation *OutcomeDenomMutation
	hooks    []Hook
}

// SetOutcomeDenomUnits sets the "outcome_denom_units" field.
func (odc *OutcomeDenomCreate) SetOutcomeDenomUnits(s string) *OutcomeDenomCreate {
	odc.mutation.SetOutcomeDenomUnits(s)
	return odc
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (odc *OutcomeDenomCreate) SetParentID(id int) *OutcomeDenomCreate {
	odc.mutation.SetParentID(id)
	return odc
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (odc *OutcomeDenomCreate) SetNillableParentID(id *int) *OutcomeDenomCreate {
	if id != nil {
		odc = odc.SetParentID(*id)
	}
	return odc
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (odc *OutcomeDenomCreate) SetParent(o *OutcomeMeasure) *OutcomeDenomCreate {
	return odc.SetParentID(o.ID)
}

// AddOutcomeDenomCountListIDs adds the "outcome_denom_count_list" edge to the OutcomeDenomCount entity by IDs.
func (odc *OutcomeDenomCreate) AddOutcomeDenomCountListIDs(ids ...int) *OutcomeDenomCreate {
	odc.mutation.AddOutcomeDenomCountListIDs(ids...)
	return odc
}

// AddOutcomeDenomCountList adds the "outcome_denom_count_list" edges to the OutcomeDenomCount entity.
func (odc *OutcomeDenomCreate) AddOutcomeDenomCountList(o ...*OutcomeDenomCount) *OutcomeDenomCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return odc.AddOutcomeDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeDenomMutation object of the builder.
func (odc *OutcomeDenomCreate) Mutation() *OutcomeDenomMutation {
	return odc.mutation
}

// Save creates the OutcomeDenom in the database.
func (odc *OutcomeDenomCreate) Save(ctx context.Context) (*OutcomeDenom, error) {
	var (
		err  error
		node *OutcomeDenom
	)
	if len(odc.hooks) == 0 {
		if err = odc.check(); err != nil {
			return nil, err
		}
		node, err = odc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = odc.check(); err != nil {
				return nil, err
			}
			odc.mutation = mutation
			if node, err = odc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(odc.hooks) - 1; i >= 0; i-- {
			if odc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (odc *OutcomeDenomCreate) SaveX(ctx context.Context) *OutcomeDenom {
	v, err := odc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odc *OutcomeDenomCreate) Exec(ctx context.Context) error {
	_, err := odc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odc *OutcomeDenomCreate) ExecX(ctx context.Context) {
	if err := odc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odc *OutcomeDenomCreate) check() error {
	if _, ok := odc.mutation.OutcomeDenomUnits(); !ok {
		return &ValidationError{Name: "outcome_denom_units", err: errors.New(`models: missing required field "OutcomeDenom.outcome_denom_units"`)}
	}
	return nil
}

func (odc *OutcomeDenomCreate) sqlSave(ctx context.Context) (*OutcomeDenom, error) {
	_node, _spec := odc.createSpec()
	if err := sqlgraph.CreateNode(ctx, odc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (odc *OutcomeDenomCreate) createSpec() (*OutcomeDenom, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeDenom{config: odc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomedenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenom.FieldID,
			},
		}
	)
	if value, ok := odc.mutation.OutcomeDenomUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenom.FieldOutcomeDenomUnits,
		})
		_node.OutcomeDenomUnits = value
	}
	if nodes := odc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenom.ParentTable,
			Columns: []string{outcomedenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_measure_outcome_denom_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := odc.mutation.OutcomeDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
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

// OutcomeDenomCreateBulk is the builder for creating many OutcomeDenom entities in bulk.
type OutcomeDenomCreateBulk struct {
	config
	builders []*OutcomeDenomCreate
}

// Save creates the OutcomeDenom entities in the database.
func (odcb *OutcomeDenomCreateBulk) Save(ctx context.Context) ([]*OutcomeDenom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(odcb.builders))
	nodes := make([]*OutcomeDenom, len(odcb.builders))
	mutators := make([]Mutator, len(odcb.builders))
	for i := range odcb.builders {
		func(i int, root context.Context) {
			builder := odcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeDenomMutation)
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
					_, err = mutators[i+1].Mutate(root, odcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, odcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, odcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (odcb *OutcomeDenomCreateBulk) SaveX(ctx context.Context) []*OutcomeDenom {
	v, err := odcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odcb *OutcomeDenomCreateBulk) Exec(ctx context.Context) error {
	_, err := odcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcb *OutcomeDenomCreateBulk) ExecX(ctx context.Context) {
	if err := odcb.Exec(ctx); err != nil {
		panic(err)
	}
}
