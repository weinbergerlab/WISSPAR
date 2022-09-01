// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
)

// FlowGroupCreate is the builder for creating a FlowGroup entity.
type FlowGroupCreate struct {
	config
	mutation *FlowGroupMutation
	hooks    []Hook
}

// SetFlowGroupID sets the "flow_group_id" field.
func (fgc *FlowGroupCreate) SetFlowGroupID(s string) *FlowGroupCreate {
	fgc.mutation.SetFlowGroupID(s)
	return fgc
}

// SetFlowGroupTitle sets the "flow_group_title" field.
func (fgc *FlowGroupCreate) SetFlowGroupTitle(s string) *FlowGroupCreate {
	fgc.mutation.SetFlowGroupTitle(s)
	return fgc
}

// SetFlowGroupDescription sets the "flow_group_description" field.
func (fgc *FlowGroupCreate) SetFlowGroupDescription(s string) *FlowGroupCreate {
	fgc.mutation.SetFlowGroupDescription(s)
	return fgc
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fgc *FlowGroupCreate) SetParentID(id int) *FlowGroupCreate {
	fgc.mutation.SetParentID(id)
	return fgc
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fgc *FlowGroupCreate) SetNillableParentID(id *int) *FlowGroupCreate {
	if id != nil {
		fgc = fgc.SetParentID(*id)
	}
	return fgc
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fgc *FlowGroupCreate) SetParent(p *ParticipantFlowModule) *FlowGroupCreate {
	return fgc.SetParentID(p.ID)
}

// Mutation returns the FlowGroupMutation object of the builder.
func (fgc *FlowGroupCreate) Mutation() *FlowGroupMutation {
	return fgc.mutation
}

// Save creates the FlowGroup in the database.
func (fgc *FlowGroupCreate) Save(ctx context.Context) (*FlowGroup, error) {
	var (
		err  error
		node *FlowGroup
	)
	if len(fgc.hooks) == 0 {
		if err = fgc.check(); err != nil {
			return nil, err
		}
		node, err = fgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fgc.check(); err != nil {
				return nil, err
			}
			fgc.mutation = mutation
			if node, err = fgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fgc.hooks) - 1; i >= 0; i-- {
			if fgc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fgc *FlowGroupCreate) SaveX(ctx context.Context) *FlowGroup {
	v, err := fgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fgc *FlowGroupCreate) Exec(ctx context.Context) error {
	_, err := fgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgc *FlowGroupCreate) ExecX(ctx context.Context) {
	if err := fgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fgc *FlowGroupCreate) check() error {
	if _, ok := fgc.mutation.FlowGroupID(); !ok {
		return &ValidationError{Name: "flow_group_id", err: errors.New(`models: missing required field "FlowGroup.flow_group_id"`)}
	}
	if _, ok := fgc.mutation.FlowGroupTitle(); !ok {
		return &ValidationError{Name: "flow_group_title", err: errors.New(`models: missing required field "FlowGroup.flow_group_title"`)}
	}
	if _, ok := fgc.mutation.FlowGroupDescription(); !ok {
		return &ValidationError{Name: "flow_group_description", err: errors.New(`models: missing required field "FlowGroup.flow_group_description"`)}
	}
	return nil
}

func (fgc *FlowGroupCreate) sqlSave(ctx context.Context) (*FlowGroup, error) {
	_node, _spec := fgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fgc *FlowGroupCreate) createSpec() (*FlowGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowGroup{config: fgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowgroup.FieldID,
			},
		}
	)
	if value, ok := fgc.mutation.FlowGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupID,
		})
		_node.FlowGroupID = value
	}
	if value, ok := fgc.mutation.FlowGroupTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupTitle,
		})
		_node.FlowGroupTitle = value
	}
	if value, ok := fgc.mutation.FlowGroupDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupDescription,
		})
		_node.FlowGroupDescription = value
	}
	if nodes := fgc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowgroup.ParentTable,
			Columns: []string{flowgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.participant_flow_module_flow_group_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlowGroupCreateBulk is the builder for creating many FlowGroup entities in bulk.
type FlowGroupCreateBulk struct {
	config
	builders []*FlowGroupCreate
}

// Save creates the FlowGroup entities in the database.
func (fgcb *FlowGroupCreateBulk) Save(ctx context.Context) ([]*FlowGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fgcb.builders))
	nodes := make([]*FlowGroup, len(fgcb.builders))
	mutators := make([]Mutator, len(fgcb.builders))
	for i := range fgcb.builders {
		func(i int, root context.Context) {
			builder := fgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, fgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fgcb *FlowGroupCreateBulk) SaveX(ctx context.Context) []*FlowGroup {
	v, err := fgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fgcb *FlowGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := fgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgcb *FlowGroupCreateBulk) ExecX(ctx context.Context) {
	if err := fgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
