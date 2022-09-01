// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
)

// FlowReasonCreate is the builder for creating a FlowReason entity.
type FlowReasonCreate struct {
	config
	mutation *FlowReasonMutation
	hooks    []Hook
}

// SetFlowReasonGroupID sets the "flow_reason_group_id" field.
func (frc *FlowReasonCreate) SetFlowReasonGroupID(s string) *FlowReasonCreate {
	frc.mutation.SetFlowReasonGroupID(s)
	return frc
}

// SetFlowReasonComment sets the "flow_reason_comment" field.
func (frc *FlowReasonCreate) SetFlowReasonComment(s string) *FlowReasonCreate {
	frc.mutation.SetFlowReasonComment(s)
	return frc
}

// SetFlowReasonNumSubjects sets the "flow_reason_num_subjects" field.
func (frc *FlowReasonCreate) SetFlowReasonNumSubjects(s string) *FlowReasonCreate {
	frc.mutation.SetFlowReasonNumSubjects(s)
	return frc
}

// SetFlowReasonNumUnits sets the "flow_reason_num_units" field.
func (frc *FlowReasonCreate) SetFlowReasonNumUnits(s string) *FlowReasonCreate {
	frc.mutation.SetFlowReasonNumUnits(s)
	return frc
}

// SetParentID sets the "parent" edge to the FlowDropWithdraw entity by ID.
func (frc *FlowReasonCreate) SetParentID(id int) *FlowReasonCreate {
	frc.mutation.SetParentID(id)
	return frc
}

// SetNillableParentID sets the "parent" edge to the FlowDropWithdraw entity by ID if the given value is not nil.
func (frc *FlowReasonCreate) SetNillableParentID(id *int) *FlowReasonCreate {
	if id != nil {
		frc = frc.SetParentID(*id)
	}
	return frc
}

// SetParent sets the "parent" edge to the FlowDropWithdraw entity.
func (frc *FlowReasonCreate) SetParent(f *FlowDropWithdraw) *FlowReasonCreate {
	return frc.SetParentID(f.ID)
}

// Mutation returns the FlowReasonMutation object of the builder.
func (frc *FlowReasonCreate) Mutation() *FlowReasonMutation {
	return frc.mutation
}

// Save creates the FlowReason in the database.
func (frc *FlowReasonCreate) Save(ctx context.Context) (*FlowReason, error) {
	var (
		err  error
		node *FlowReason
	)
	if len(frc.hooks) == 0 {
		if err = frc.check(); err != nil {
			return nil, err
		}
		node, err = frc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = frc.check(); err != nil {
				return nil, err
			}
			frc.mutation = mutation
			if node, err = frc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(frc.hooks) - 1; i >= 0; i-- {
			if frc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = frc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, frc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (frc *FlowReasonCreate) SaveX(ctx context.Context) *FlowReason {
	v, err := frc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frc *FlowReasonCreate) Exec(ctx context.Context) error {
	_, err := frc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frc *FlowReasonCreate) ExecX(ctx context.Context) {
	if err := frc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (frc *FlowReasonCreate) check() error {
	if _, ok := frc.mutation.FlowReasonGroupID(); !ok {
		return &ValidationError{Name: "flow_reason_group_id", err: errors.New(`models: missing required field "FlowReason.flow_reason_group_id"`)}
	}
	if _, ok := frc.mutation.FlowReasonComment(); !ok {
		return &ValidationError{Name: "flow_reason_comment", err: errors.New(`models: missing required field "FlowReason.flow_reason_comment"`)}
	}
	if _, ok := frc.mutation.FlowReasonNumSubjects(); !ok {
		return &ValidationError{Name: "flow_reason_num_subjects", err: errors.New(`models: missing required field "FlowReason.flow_reason_num_subjects"`)}
	}
	if _, ok := frc.mutation.FlowReasonNumUnits(); !ok {
		return &ValidationError{Name: "flow_reason_num_units", err: errors.New(`models: missing required field "FlowReason.flow_reason_num_units"`)}
	}
	return nil
}

func (frc *FlowReasonCreate) sqlSave(ctx context.Context) (*FlowReason, error) {
	_node, _spec := frc.createSpec()
	if err := sqlgraph.CreateNode(ctx, frc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (frc *FlowReasonCreate) createSpec() (*FlowReason, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowReason{config: frc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowreason.FieldID,
			},
		}
	)
	if value, ok := frc.mutation.FlowReasonGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonGroupID,
		})
		_node.FlowReasonGroupID = value
	}
	if value, ok := frc.mutation.FlowReasonComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonComment,
		})
		_node.FlowReasonComment = value
	}
	if value, ok := frc.mutation.FlowReasonNumSubjects(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumSubjects,
		})
		_node.FlowReasonNumSubjects = value
	}
	if value, ok := frc.mutation.FlowReasonNumUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowreason.FieldFlowReasonNumUnits,
		})
		_node.FlowReasonNumUnits = value
	}
	if nodes := frc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowreason.ParentTable,
			Columns: []string{flowreason.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowdropwithdraw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.flow_drop_withdraw_flow_reason_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowReasonCreateBulk is the builder for creating many FlowReason entities in bulk.
type FlowReasonCreateBulk struct {
	config
	builders []*FlowReasonCreate
}

// Save creates the FlowReason entities in the database.
func (frcb *FlowReasonCreateBulk) Save(ctx context.Context) ([]*FlowReason, error) {
	specs := make([]*sqlgraph.CreateSpec, len(frcb.builders))
	nodes := make([]*FlowReason, len(frcb.builders))
	mutators := make([]Mutator, len(frcb.builders))
	for i := range frcb.builders {
		func(i int, root context.Context) {
			builder := frcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowReasonMutation)
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
					_, err = mutators[i+1].Mutate(root, frcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, frcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, frcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (frcb *FlowReasonCreateBulk) SaveX(ctx context.Context) []*FlowReason {
	v, err := frcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frcb *FlowReasonCreateBulk) Exec(ctx context.Context) error {
	_, err := frcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frcb *FlowReasonCreateBulk) ExecX(ctx context.Context) {
	if err := frcb.Exec(ctx); err != nil {
		panic(err)
	}
}
