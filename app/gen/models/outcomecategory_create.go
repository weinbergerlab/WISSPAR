// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
)

// OutcomeCategoryCreate is the builder for creating a OutcomeCategory entity.
type OutcomeCategoryCreate struct {
	config
	mutation *OutcomeCategoryMutation
	hooks    []Hook
}

// SetOutcomeCategoryTitle sets the "outcome_category_title" field.
func (occ *OutcomeCategoryCreate) SetOutcomeCategoryTitle(s string) *OutcomeCategoryCreate {
	occ.mutation.SetOutcomeCategoryTitle(s)
	return occ
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (occ *OutcomeCategoryCreate) SetParentID(id int) *OutcomeCategoryCreate {
	occ.mutation.SetParentID(id)
	return occ
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (occ *OutcomeCategoryCreate) SetNillableParentID(id *int) *OutcomeCategoryCreate {
	if id != nil {
		occ = occ.SetParentID(*id)
	}
	return occ
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (occ *OutcomeCategoryCreate) SetParent(o *OutcomeClass) *OutcomeCategoryCreate {
	return occ.SetParentID(o.ID)
}

// AddOutcomeMeasurementListIDs adds the "outcome_measurement_list" edge to the OutcomeMeasurement entity by IDs.
func (occ *OutcomeCategoryCreate) AddOutcomeMeasurementListIDs(ids ...int) *OutcomeCategoryCreate {
	occ.mutation.AddOutcomeMeasurementListIDs(ids...)
	return occ
}

// AddOutcomeMeasurementList adds the "outcome_measurement_list" edges to the OutcomeMeasurement entity.
func (occ *OutcomeCategoryCreate) AddOutcomeMeasurementList(o ...*OutcomeMeasurement) *OutcomeCategoryCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return occ.AddOutcomeMeasurementListIDs(ids...)
}

// Mutation returns the OutcomeCategoryMutation object of the builder.
func (occ *OutcomeCategoryCreate) Mutation() *OutcomeCategoryMutation {
	return occ.mutation
}

// Save creates the OutcomeCategory in the database.
func (occ *OutcomeCategoryCreate) Save(ctx context.Context) (*OutcomeCategory, error) {
	var (
		err  error
		node *OutcomeCategory
	)
	if len(occ.hooks) == 0 {
		if err = occ.check(); err != nil {
			return nil, err
		}
		node, err = occ.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = occ.check(); err != nil {
				return nil, err
			}
			occ.mutation = mutation
			if node, err = occ.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(occ.hooks) - 1; i >= 0; i-- {
			if occ.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = occ.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, occ.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (occ *OutcomeCategoryCreate) SaveX(ctx context.Context) *OutcomeCategory {
	v, err := occ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occ *OutcomeCategoryCreate) Exec(ctx context.Context) error {
	_, err := occ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occ *OutcomeCategoryCreate) ExecX(ctx context.Context) {
	if err := occ.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (occ *OutcomeCategoryCreate) check() error {
	if _, ok := occ.mutation.OutcomeCategoryTitle(); !ok {
		return &ValidationError{Name: "outcome_category_title", err: errors.New(`models: missing required field "OutcomeCategory.outcome_category_title"`)}
	}
	return nil
}

func (occ *OutcomeCategoryCreate) sqlSave(ctx context.Context) (*OutcomeCategory, error) {
	_node, _spec := occ.createSpec()
	if err := sqlgraph.CreateNode(ctx, occ.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (occ *OutcomeCategoryCreate) createSpec() (*OutcomeCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeCategory{config: occ.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomecategory.FieldID,
			},
		}
	)
	if value, ok := occ.mutation.OutcomeCategoryTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomecategory.FieldOutcomeCategoryTitle,
		})
		_node.OutcomeCategoryTitle = value
	}
	if nodes := occ.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomecategory.ParentTable,
			Columns: []string{outcomecategory.ParentColumn},
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
		_node.outcome_class_outcome_category_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := occ.mutation.OutcomeMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
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

// OutcomeCategoryCreateBulk is the builder for creating many OutcomeCategory entities in bulk.
type OutcomeCategoryCreateBulk struct {
	config
	builders []*OutcomeCategoryCreate
}

// Save creates the OutcomeCategory entities in the database.
func (occb *OutcomeCategoryCreateBulk) Save(ctx context.Context) ([]*OutcomeCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(occb.builders))
	nodes := make([]*OutcomeCategory, len(occb.builders))
	mutators := make([]Mutator, len(occb.builders))
	for i := range occb.builders {
		func(i int, root context.Context) {
			builder := occb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, occb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, occb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, occb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (occb *OutcomeCategoryCreateBulk) SaveX(ctx context.Context) []*OutcomeCategory {
	v, err := occb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occb *OutcomeCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := occb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occb *OutcomeCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := occb.Exec(ctx); err != nil {
		panic(err)
	}
}
