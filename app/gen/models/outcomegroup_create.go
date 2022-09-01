// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeGroupCreate is the builder for creating a OutcomeGroup entity.
type OutcomeGroupCreate struct {
	config
	mutation *OutcomeGroupMutation
	hooks    []Hook
}

// SetOutcomeGroupID sets the "outcome_group_id" field.
func (ogc *OutcomeGroupCreate) SetOutcomeGroupID(s string) *OutcomeGroupCreate {
	ogc.mutation.SetOutcomeGroupID(s)
	return ogc
}

// SetOutcomeGroupTitle sets the "outcome_group_title" field.
func (ogc *OutcomeGroupCreate) SetOutcomeGroupTitle(s string) *OutcomeGroupCreate {
	ogc.mutation.SetOutcomeGroupTitle(s)
	return ogc
}

// SetOutcomeGroupDescription sets the "outcome_group_description" field.
func (ogc *OutcomeGroupCreate) SetOutcomeGroupDescription(s string) *OutcomeGroupCreate {
	ogc.mutation.SetOutcomeGroupDescription(s)
	return ogc
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (ogc *OutcomeGroupCreate) SetParentID(id int) *OutcomeGroupCreate {
	ogc.mutation.SetParentID(id)
	return ogc
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (ogc *OutcomeGroupCreate) SetNillableParentID(id *int) *OutcomeGroupCreate {
	if id != nil {
		ogc = ogc.SetParentID(*id)
	}
	return ogc
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (ogc *OutcomeGroupCreate) SetParent(o *OutcomeMeasure) *OutcomeGroupCreate {
	return ogc.SetParentID(o.ID)
}

// Mutation returns the OutcomeGroupMutation object of the builder.
func (ogc *OutcomeGroupCreate) Mutation() *OutcomeGroupMutation {
	return ogc.mutation
}

// Save creates the OutcomeGroup in the database.
func (ogc *OutcomeGroupCreate) Save(ctx context.Context) (*OutcomeGroup, error) {
	var (
		err  error
		node *OutcomeGroup
	)
	if len(ogc.hooks) == 0 {
		if err = ogc.check(); err != nil {
			return nil, err
		}
		node, err = ogc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ogc.check(); err != nil {
				return nil, err
			}
			ogc.mutation = mutation
			if node, err = ogc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ogc.hooks) - 1; i >= 0; i-- {
			if ogc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ogc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ogc *OutcomeGroupCreate) SaveX(ctx context.Context) *OutcomeGroup {
	v, err := ogc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ogc *OutcomeGroupCreate) Exec(ctx context.Context) error {
	_, err := ogc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogc *OutcomeGroupCreate) ExecX(ctx context.Context) {
	if err := ogc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ogc *OutcomeGroupCreate) check() error {
	if _, ok := ogc.mutation.OutcomeGroupID(); !ok {
		return &ValidationError{Name: "outcome_group_id", err: errors.New(`models: missing required field "OutcomeGroup.outcome_group_id"`)}
	}
	if _, ok := ogc.mutation.OutcomeGroupTitle(); !ok {
		return &ValidationError{Name: "outcome_group_title", err: errors.New(`models: missing required field "OutcomeGroup.outcome_group_title"`)}
	}
	if _, ok := ogc.mutation.OutcomeGroupDescription(); !ok {
		return &ValidationError{Name: "outcome_group_description", err: errors.New(`models: missing required field "OutcomeGroup.outcome_group_description"`)}
	}
	return nil
}

func (ogc *OutcomeGroupCreate) sqlSave(ctx context.Context) (*OutcomeGroup, error) {
	_node, _spec := ogc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ogc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ogc *OutcomeGroupCreate) createSpec() (*OutcomeGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeGroup{config: ogc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomegroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomegroup.FieldID,
			},
		}
	)
	if value, ok := ogc.mutation.OutcomeGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupID,
		})
		_node.OutcomeGroupID = value
	}
	if value, ok := ogc.mutation.OutcomeGroupTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupTitle,
		})
		_node.OutcomeGroupTitle = value
	}
	if value, ok := ogc.mutation.OutcomeGroupDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupDescription,
		})
		_node.OutcomeGroupDescription = value
	}
	if nodes := ogc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomegroup.ParentTable,
			Columns: []string{outcomegroup.ParentColumn},
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
		_node.outcome_measure_outcome_group_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeGroupCreateBulk is the builder for creating many OutcomeGroup entities in bulk.
type OutcomeGroupCreateBulk struct {
	config
	builders []*OutcomeGroupCreate
}

// Save creates the OutcomeGroup entities in the database.
func (ogcb *OutcomeGroupCreateBulk) Save(ctx context.Context) ([]*OutcomeGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ogcb.builders))
	nodes := make([]*OutcomeGroup, len(ogcb.builders))
	mutators := make([]Mutator, len(ogcb.builders))
	for i := range ogcb.builders {
		func(i int, root context.Context) {
			builder := ogcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, ogcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ogcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ogcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ogcb *OutcomeGroupCreateBulk) SaveX(ctx context.Context) []*OutcomeGroup {
	v, err := ogcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ogcb *OutcomeGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := ogcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogcb *OutcomeGroupCreateBulk) ExecX(ctx context.Context) {
	if err := ogcb.Exec(ctx); err != nil {
		panic(err)
	}
}
