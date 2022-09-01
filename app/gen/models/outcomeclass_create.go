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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeClassCreate is the builder for creating a OutcomeClass entity.
type OutcomeClassCreate struct {
	config
	mutation *OutcomeClassMutation
	hooks    []Hook
}

// SetOutcomeClassTitle sets the "outcome_class_title" field.
func (occ *OutcomeClassCreate) SetOutcomeClassTitle(s string) *OutcomeClassCreate {
	occ.mutation.SetOutcomeClassTitle(s)
	return occ
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (occ *OutcomeClassCreate) SetParentID(id int) *OutcomeClassCreate {
	occ.mutation.SetParentID(id)
	return occ
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (occ *OutcomeClassCreate) SetNillableParentID(id *int) *OutcomeClassCreate {
	if id != nil {
		occ = occ.SetParentID(*id)
	}
	return occ
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (occ *OutcomeClassCreate) SetParent(o *OutcomeMeasure) *OutcomeClassCreate {
	return occ.SetParentID(o.ID)
}

// AddOutcomeClassDenomListIDs adds the "outcome_class_denom_list" edge to the OutcomeClassDenom entity by IDs.
func (occ *OutcomeClassCreate) AddOutcomeClassDenomListIDs(ids ...int) *OutcomeClassCreate {
	occ.mutation.AddOutcomeClassDenomListIDs(ids...)
	return occ
}

// AddOutcomeClassDenomList adds the "outcome_class_denom_list" edges to the OutcomeClassDenom entity.
func (occ *OutcomeClassCreate) AddOutcomeClassDenomList(o ...*OutcomeClassDenom) *OutcomeClassCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return occ.AddOutcomeClassDenomListIDs(ids...)
}

// AddOutcomeCategoryListIDs adds the "outcome_category_list" edge to the OutcomeCategory entity by IDs.
func (occ *OutcomeClassCreate) AddOutcomeCategoryListIDs(ids ...int) *OutcomeClassCreate {
	occ.mutation.AddOutcomeCategoryListIDs(ids...)
	return occ
}

// AddOutcomeCategoryList adds the "outcome_category_list" edges to the OutcomeCategory entity.
func (occ *OutcomeClassCreate) AddOutcomeCategoryList(o ...*OutcomeCategory) *OutcomeClassCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return occ.AddOutcomeCategoryListIDs(ids...)
}

// Mutation returns the OutcomeClassMutation object of the builder.
func (occ *OutcomeClassCreate) Mutation() *OutcomeClassMutation {
	return occ.mutation
}

// Save creates the OutcomeClass in the database.
func (occ *OutcomeClassCreate) Save(ctx context.Context) (*OutcomeClass, error) {
	var (
		err  error
		node *OutcomeClass
	)
	if len(occ.hooks) == 0 {
		if err = occ.check(); err != nil {
			return nil, err
		}
		node, err = occ.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassMutation)
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
func (occ *OutcomeClassCreate) SaveX(ctx context.Context) *OutcomeClass {
	v, err := occ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occ *OutcomeClassCreate) Exec(ctx context.Context) error {
	_, err := occ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occ *OutcomeClassCreate) ExecX(ctx context.Context) {
	if err := occ.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (occ *OutcomeClassCreate) check() error {
	if _, ok := occ.mutation.OutcomeClassTitle(); !ok {
		return &ValidationError{Name: "outcome_class_title", err: errors.New(`models: missing required field "OutcomeClass.outcome_class_title"`)}
	}
	return nil
}

func (occ *OutcomeClassCreate) sqlSave(ctx context.Context) (*OutcomeClass, error) {
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

func (occ *OutcomeClassCreate) createSpec() (*OutcomeClass, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeClass{config: occ.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclass.FieldID,
			},
		}
	)
	if value, ok := occ.mutation.OutcomeClassTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclass.FieldOutcomeClassTitle,
		})
		_node.OutcomeClassTitle = value
	}
	if nodes := occ.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclass.ParentTable,
			Columns: []string{outcomeclass.ParentColumn},
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
		_node.outcome_measure_outcome_class_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := occ.mutation.OutcomeClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := occ.mutation.OutcomeCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
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

// OutcomeClassCreateBulk is the builder for creating many OutcomeClass entities in bulk.
type OutcomeClassCreateBulk struct {
	config
	builders []*OutcomeClassCreate
}

// Save creates the OutcomeClass entities in the database.
func (occb *OutcomeClassCreateBulk) Save(ctx context.Context) ([]*OutcomeClass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(occb.builders))
	nodes := make([]*OutcomeClass, len(occb.builders))
	mutators := make([]Mutator, len(occb.builders))
	for i := range occb.builders {
		func(i int, root context.Context) {
			builder := occb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeClassMutation)
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
func (occb *OutcomeClassCreateBulk) SaveX(ctx context.Context) []*OutcomeClass {
	v, err := occb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occb *OutcomeClassCreateBulk) Exec(ctx context.Context) error {
	_, err := occb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occb *OutcomeClassCreateBulk) ExecX(ctx context.Context) {
	if err := occb.Exec(ctx); err != nil {
		panic(err)
	}
}
