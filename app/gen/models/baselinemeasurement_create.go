// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
)

// BaselineMeasurementCreate is the builder for creating a BaselineMeasurement entity.
type BaselineMeasurementCreate struct {
	config
	mutation *BaselineMeasurementMutation
	hooks    []Hook
}

// SetBaselineMeasurementGroupID sets the "baseline_measurement_group_id" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementGroupID(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementGroupID(s)
	return bmc
}

// SetBaselineMeasurementValue sets the "baseline_measurement_value" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementValue(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementValue(s)
	return bmc
}

// SetBaselineMeasurementSpread sets the "baseline_measurement_spread" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementSpread(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementSpread(s)
	return bmc
}

// SetBaselineMeasurementLowerLimit sets the "baseline_measurement_lower_limit" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementLowerLimit(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementLowerLimit(s)
	return bmc
}

// SetBaselineMeasurementUpperLimit sets the "baseline_measurement_upper_limit" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementUpperLimit(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementUpperLimit(s)
	return bmc
}

// SetBaselineMeasurementComment sets the "baseline_measurement_comment" field.
func (bmc *BaselineMeasurementCreate) SetBaselineMeasurementComment(s string) *BaselineMeasurementCreate {
	bmc.mutation.SetBaselineMeasurementComment(s)
	return bmc
}

// SetParentID sets the "parent" edge to the BaselineCategory entity by ID.
func (bmc *BaselineMeasurementCreate) SetParentID(id int) *BaselineMeasurementCreate {
	bmc.mutation.SetParentID(id)
	return bmc
}

// SetNillableParentID sets the "parent" edge to the BaselineCategory entity by ID if the given value is not nil.
func (bmc *BaselineMeasurementCreate) SetNillableParentID(id *int) *BaselineMeasurementCreate {
	if id != nil {
		bmc = bmc.SetParentID(*id)
	}
	return bmc
}

// SetParent sets the "parent" edge to the BaselineCategory entity.
func (bmc *BaselineMeasurementCreate) SetParent(b *BaselineCategory) *BaselineMeasurementCreate {
	return bmc.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasurementMutation object of the builder.
func (bmc *BaselineMeasurementCreate) Mutation() *BaselineMeasurementMutation {
	return bmc.mutation
}

// Save creates the BaselineMeasurement in the database.
func (bmc *BaselineMeasurementCreate) Save(ctx context.Context) (*BaselineMeasurement, error) {
	var (
		err  error
		node *BaselineMeasurement
	)
	if len(bmc.hooks) == 0 {
		if err = bmc.check(); err != nil {
			return nil, err
		}
		node, err = bmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bmc.check(); err != nil {
				return nil, err
			}
			bmc.mutation = mutation
			if node, err = bmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bmc.hooks) - 1; i >= 0; i-- {
			if bmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bmc *BaselineMeasurementCreate) SaveX(ctx context.Context) *BaselineMeasurement {
	v, err := bmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmc *BaselineMeasurementCreate) Exec(ctx context.Context) error {
	_, err := bmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmc *BaselineMeasurementCreate) ExecX(ctx context.Context) {
	if err := bmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bmc *BaselineMeasurementCreate) check() error {
	if _, ok := bmc.mutation.BaselineMeasurementGroupID(); !ok {
		return &ValidationError{Name: "baseline_measurement_group_id", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_group_id"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurementValue(); !ok {
		return &ValidationError{Name: "baseline_measurement_value", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_value"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurementSpread(); !ok {
		return &ValidationError{Name: "baseline_measurement_spread", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_spread"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurementLowerLimit(); !ok {
		return &ValidationError{Name: "baseline_measurement_lower_limit", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_lower_limit"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurementUpperLimit(); !ok {
		return &ValidationError{Name: "baseline_measurement_upper_limit", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_upper_limit"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurementComment(); !ok {
		return &ValidationError{Name: "baseline_measurement_comment", err: errors.New(`models: missing required field "BaselineMeasurement.baseline_measurement_comment"`)}
	}
	return nil
}

func (bmc *BaselineMeasurementCreate) sqlSave(ctx context.Context) (*BaselineMeasurement, error) {
	_node, _spec := bmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bmc *BaselineMeasurementCreate) createSpec() (*BaselineMeasurement, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineMeasurement{config: bmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinemeasurement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasurement.FieldID,
			},
		}
	)
	if value, ok := bmc.mutation.BaselineMeasurementGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementGroupID,
		})
		_node.BaselineMeasurementGroupID = value
	}
	if value, ok := bmc.mutation.BaselineMeasurementValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementValue,
		})
		_node.BaselineMeasurementValue = value
	}
	if value, ok := bmc.mutation.BaselineMeasurementSpread(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementSpread,
		})
		_node.BaselineMeasurementSpread = value
	}
	if value, ok := bmc.mutation.BaselineMeasurementLowerLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementLowerLimit,
		})
		_node.BaselineMeasurementLowerLimit = value
	}
	if value, ok := bmc.mutation.BaselineMeasurementUpperLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementUpperLimit,
		})
		_node.BaselineMeasurementUpperLimit = value
	}
	if value, ok := bmc.mutation.BaselineMeasurementComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementComment,
		})
		_node.BaselineMeasurementComment = value
	}
	if nodes := bmc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasurement.ParentTable,
			Columns: []string{baselinemeasurement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_category_baseline_measurement_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BaselineMeasurementCreateBulk is the builder for creating many BaselineMeasurement entities in bulk.
type BaselineMeasurementCreateBulk struct {
	config
	builders []*BaselineMeasurementCreate
}

// Save creates the BaselineMeasurement entities in the database.
func (bmcb *BaselineMeasurementCreateBulk) Save(ctx context.Context) ([]*BaselineMeasurement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bmcb.builders))
	nodes := make([]*BaselineMeasurement, len(bmcb.builders))
	mutators := make([]Mutator, len(bmcb.builders))
	for i := range bmcb.builders {
		func(i int, root context.Context) {
			builder := bmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineMeasurementMutation)
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
					_, err = mutators[i+1].Mutate(root, bmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bmcb *BaselineMeasurementCreateBulk) SaveX(ctx context.Context) []*BaselineMeasurement {
	v, err := bmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmcb *BaselineMeasurementCreateBulk) Exec(ctx context.Context) error {
	_, err := bmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmcb *BaselineMeasurementCreateBulk) ExecX(ctx context.Context) {
	if err := bmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
