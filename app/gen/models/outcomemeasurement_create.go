// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
)

// OutcomeMeasurementCreate is the builder for creating a OutcomeMeasurement entity.
type OutcomeMeasurementCreate struct {
	config
	mutation *OutcomeMeasurementMutation
	hooks    []Hook
}

// SetOutcomeMeasurementGroupID sets the "outcome_measurement_group_id" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementGroupID(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementGroupID(s)
	return omc
}

// SetOutcomeMeasurementValue sets the "outcome_measurement_value" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementValue(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementValue(s)
	return omc
}

// SetOutcomeMeasurementSpread sets the "outcome_measurement_spread" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementSpread(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementSpread(s)
	return omc
}

// SetOutcomeMeasurementLowerLimit sets the "outcome_measurement_lower_limit" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementLowerLimit(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementLowerLimit(s)
	return omc
}

// SetOutcomeMeasurementUpperLimit sets the "outcome_measurement_upper_limit" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementUpperLimit(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementUpperLimit(s)
	return omc
}

// SetOutcomeMeasurementComment sets the "outcome_measurement_comment" field.
func (omc *OutcomeMeasurementCreate) SetOutcomeMeasurementComment(s string) *OutcomeMeasurementCreate {
	omc.mutation.SetOutcomeMeasurementComment(s)
	return omc
}

// SetParentID sets the "parent" edge to the OutcomeCategory entity by ID.
func (omc *OutcomeMeasurementCreate) SetParentID(id int) *OutcomeMeasurementCreate {
	omc.mutation.SetParentID(id)
	return omc
}

// SetNillableParentID sets the "parent" edge to the OutcomeCategory entity by ID if the given value is not nil.
func (omc *OutcomeMeasurementCreate) SetNillableParentID(id *int) *OutcomeMeasurementCreate {
	if id != nil {
		omc = omc.SetParentID(*id)
	}
	return omc
}

// SetParent sets the "parent" edge to the OutcomeCategory entity.
func (omc *OutcomeMeasurementCreate) SetParent(o *OutcomeCategory) *OutcomeMeasurementCreate {
	return omc.SetParentID(o.ID)
}

// Mutation returns the OutcomeMeasurementMutation object of the builder.
func (omc *OutcomeMeasurementCreate) Mutation() *OutcomeMeasurementMutation {
	return omc.mutation
}

// Save creates the OutcomeMeasurement in the database.
func (omc *OutcomeMeasurementCreate) Save(ctx context.Context) (*OutcomeMeasurement, error) {
	var (
		err  error
		node *OutcomeMeasurement
	)
	if len(omc.hooks) == 0 {
		if err = omc.check(); err != nil {
			return nil, err
		}
		node, err = omc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = omc.check(); err != nil {
				return nil, err
			}
			omc.mutation = mutation
			if node, err = omc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(omc.hooks) - 1; i >= 0; i-- {
			if omc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (omc *OutcomeMeasurementCreate) SaveX(ctx context.Context) *OutcomeMeasurement {
	v, err := omc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omc *OutcomeMeasurementCreate) Exec(ctx context.Context) error {
	_, err := omc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omc *OutcomeMeasurementCreate) ExecX(ctx context.Context) {
	if err := omc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (omc *OutcomeMeasurementCreate) check() error {
	if _, ok := omc.mutation.OutcomeMeasurementGroupID(); !ok {
		return &ValidationError{Name: "outcome_measurement_group_id", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_group_id"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurementValue(); !ok {
		return &ValidationError{Name: "outcome_measurement_value", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_value"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurementSpread(); !ok {
		return &ValidationError{Name: "outcome_measurement_spread", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_spread"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurementLowerLimit(); !ok {
		return &ValidationError{Name: "outcome_measurement_lower_limit", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_lower_limit"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurementUpperLimit(); !ok {
		return &ValidationError{Name: "outcome_measurement_upper_limit", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_upper_limit"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurementComment(); !ok {
		return &ValidationError{Name: "outcome_measurement_comment", err: errors.New(`models: missing required field "OutcomeMeasurement.outcome_measurement_comment"`)}
	}
	return nil
}

func (omc *OutcomeMeasurementCreate) sqlSave(ctx context.Context) (*OutcomeMeasurement, error) {
	_node, _spec := omc.createSpec()
	if err := sqlgraph.CreateNode(ctx, omc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (omc *OutcomeMeasurementCreate) createSpec() (*OutcomeMeasurement, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeMeasurement{config: omc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomemeasurement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasurement.FieldID,
			},
		}
	)
	if value, ok := omc.mutation.OutcomeMeasurementGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementGroupID,
		})
		_node.OutcomeMeasurementGroupID = value
	}
	if value, ok := omc.mutation.OutcomeMeasurementValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementValue,
		})
		_node.OutcomeMeasurementValue = value
	}
	if value, ok := omc.mutation.OutcomeMeasurementSpread(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementSpread,
		})
		_node.OutcomeMeasurementSpread = value
	}
	if value, ok := omc.mutation.OutcomeMeasurementLowerLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementLowerLimit,
		})
		_node.OutcomeMeasurementLowerLimit = value
	}
	if value, ok := omc.mutation.OutcomeMeasurementUpperLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementUpperLimit,
		})
		_node.OutcomeMeasurementUpperLimit = value
	}
	if value, ok := omc.mutation.OutcomeMeasurementComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementComment,
		})
		_node.OutcomeMeasurementComment = value
	}
	if nodes := omc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasurement.ParentTable,
			Columns: []string{outcomemeasurement.ParentColumn},
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
		_node.outcome_category_outcome_measurement_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeMeasurementCreateBulk is the builder for creating many OutcomeMeasurement entities in bulk.
type OutcomeMeasurementCreateBulk struct {
	config
	builders []*OutcomeMeasurementCreate
}

// Save creates the OutcomeMeasurement entities in the database.
func (omcb *OutcomeMeasurementCreateBulk) Save(ctx context.Context) ([]*OutcomeMeasurement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(omcb.builders))
	nodes := make([]*OutcomeMeasurement, len(omcb.builders))
	mutators := make([]Mutator, len(omcb.builders))
	for i := range omcb.builders {
		func(i int, root context.Context) {
			builder := omcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeMeasurementMutation)
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
					_, err = mutators[i+1].Mutate(root, omcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, omcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, omcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (omcb *OutcomeMeasurementCreateBulk) SaveX(ctx context.Context) []*OutcomeMeasurement {
	v, err := omcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omcb *OutcomeMeasurementCreateBulk) Exec(ctx context.Context) error {
	_, err := omcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omcb *OutcomeMeasurementCreateBulk) ExecX(ctx context.Context) {
	if err := omcb.Exec(ctx); err != nil {
		panic(err)
	}
}
