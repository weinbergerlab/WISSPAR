// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
)

// EventGroupCreate is the builder for creating a EventGroup entity.
type EventGroupCreate struct {
	config
	mutation *EventGroupMutation
	hooks    []Hook
}

// SetEventGroupID sets the "event_group_id" field.
func (egc *EventGroupCreate) SetEventGroupID(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupID(s)
	return egc
}

// SetEventGroupTitle sets the "event_group_title" field.
func (egc *EventGroupCreate) SetEventGroupTitle(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupTitle(s)
	return egc
}

// SetEventGroupDescription sets the "event_group_description" field.
func (egc *EventGroupCreate) SetEventGroupDescription(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupDescription(s)
	return egc
}

// SetEventGroupDeathsNumAffected sets the "event_group_deaths_num_affected" field.
func (egc *EventGroupCreate) SetEventGroupDeathsNumAffected(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupDeathsNumAffected(s)
	return egc
}

// SetEventGroupDeathsNumAtRisk sets the "event_group_deaths_num_at_risk" field.
func (egc *EventGroupCreate) SetEventGroupDeathsNumAtRisk(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupDeathsNumAtRisk(s)
	return egc
}

// SetEventGroupSeriousNumAffected sets the "event_group_serious_num_affected" field.
func (egc *EventGroupCreate) SetEventGroupSeriousNumAffected(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupSeriousNumAffected(s)
	return egc
}

// SetEventGroupSeriousNumAtRisk sets the "event_group_serious_num_at_risk" field.
func (egc *EventGroupCreate) SetEventGroupSeriousNumAtRisk(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupSeriousNumAtRisk(s)
	return egc
}

// SetEventGroupOtherNumAffected sets the "event_group_other_num_affected" field.
func (egc *EventGroupCreate) SetEventGroupOtherNumAffected(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupOtherNumAffected(s)
	return egc
}

// SetEventGroupOtherNumAtRisk sets the "event_group_other_num_at_risk" field.
func (egc *EventGroupCreate) SetEventGroupOtherNumAtRisk(s string) *EventGroupCreate {
	egc.mutation.SetEventGroupOtherNumAtRisk(s)
	return egc
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (egc *EventGroupCreate) SetParentID(id int) *EventGroupCreate {
	egc.mutation.SetParentID(id)
	return egc
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (egc *EventGroupCreate) SetNillableParentID(id *int) *EventGroupCreate {
	if id != nil {
		egc = egc.SetParentID(*id)
	}
	return egc
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (egc *EventGroupCreate) SetParent(a *AdverseEventsModule) *EventGroupCreate {
	return egc.SetParentID(a.ID)
}

// Mutation returns the EventGroupMutation object of the builder.
func (egc *EventGroupCreate) Mutation() *EventGroupMutation {
	return egc.mutation
}

// Save creates the EventGroup in the database.
func (egc *EventGroupCreate) Save(ctx context.Context) (*EventGroup, error) {
	var (
		err  error
		node *EventGroup
	)
	if len(egc.hooks) == 0 {
		if err = egc.check(); err != nil {
			return nil, err
		}
		node, err = egc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = egc.check(); err != nil {
				return nil, err
			}
			egc.mutation = mutation
			if node, err = egc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(egc.hooks) - 1; i >= 0; i-- {
			if egc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = egc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, egc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (egc *EventGroupCreate) SaveX(ctx context.Context) *EventGroup {
	v, err := egc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (egc *EventGroupCreate) Exec(ctx context.Context) error {
	_, err := egc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (egc *EventGroupCreate) ExecX(ctx context.Context) {
	if err := egc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (egc *EventGroupCreate) check() error {
	if _, ok := egc.mutation.EventGroupID(); !ok {
		return &ValidationError{Name: "event_group_id", err: errors.New(`models: missing required field "EventGroup.event_group_id"`)}
	}
	if _, ok := egc.mutation.EventGroupTitle(); !ok {
		return &ValidationError{Name: "event_group_title", err: errors.New(`models: missing required field "EventGroup.event_group_title"`)}
	}
	if _, ok := egc.mutation.EventGroupDescription(); !ok {
		return &ValidationError{Name: "event_group_description", err: errors.New(`models: missing required field "EventGroup.event_group_description"`)}
	}
	if _, ok := egc.mutation.EventGroupDeathsNumAffected(); !ok {
		return &ValidationError{Name: "event_group_deaths_num_affected", err: errors.New(`models: missing required field "EventGroup.event_group_deaths_num_affected"`)}
	}
	if _, ok := egc.mutation.EventGroupDeathsNumAtRisk(); !ok {
		return &ValidationError{Name: "event_group_deaths_num_at_risk", err: errors.New(`models: missing required field "EventGroup.event_group_deaths_num_at_risk"`)}
	}
	if _, ok := egc.mutation.EventGroupSeriousNumAffected(); !ok {
		return &ValidationError{Name: "event_group_serious_num_affected", err: errors.New(`models: missing required field "EventGroup.event_group_serious_num_affected"`)}
	}
	if _, ok := egc.mutation.EventGroupSeriousNumAtRisk(); !ok {
		return &ValidationError{Name: "event_group_serious_num_at_risk", err: errors.New(`models: missing required field "EventGroup.event_group_serious_num_at_risk"`)}
	}
	if _, ok := egc.mutation.EventGroupOtherNumAffected(); !ok {
		return &ValidationError{Name: "event_group_other_num_affected", err: errors.New(`models: missing required field "EventGroup.event_group_other_num_affected"`)}
	}
	if _, ok := egc.mutation.EventGroupOtherNumAtRisk(); !ok {
		return &ValidationError{Name: "event_group_other_num_at_risk", err: errors.New(`models: missing required field "EventGroup.event_group_other_num_at_risk"`)}
	}
	return nil
}

func (egc *EventGroupCreate) sqlSave(ctx context.Context) (*EventGroup, error) {
	_node, _spec := egc.createSpec()
	if err := sqlgraph.CreateNode(ctx, egc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (egc *EventGroupCreate) createSpec() (*EventGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &EventGroup{config: egc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: eventgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventgroup.FieldID,
			},
		}
	)
	if value, ok := egc.mutation.EventGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupID,
		})
		_node.EventGroupID = value
	}
	if value, ok := egc.mutation.EventGroupTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupTitle,
		})
		_node.EventGroupTitle = value
	}
	if value, ok := egc.mutation.EventGroupDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDescription,
		})
		_node.EventGroupDescription = value
	}
	if value, ok := egc.mutation.EventGroupDeathsNumAffected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAffected,
		})
		_node.EventGroupDeathsNumAffected = value
	}
	if value, ok := egc.mutation.EventGroupDeathsNumAtRisk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAtRisk,
		})
		_node.EventGroupDeathsNumAtRisk = value
	}
	if value, ok := egc.mutation.EventGroupSeriousNumAffected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAffected,
		})
		_node.EventGroupSeriousNumAffected = value
	}
	if value, ok := egc.mutation.EventGroupSeriousNumAtRisk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAtRisk,
		})
		_node.EventGroupSeriousNumAtRisk = value
	}
	if value, ok := egc.mutation.EventGroupOtherNumAffected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAffected,
		})
		_node.EventGroupOtherNumAffected = value
	}
	if value, ok := egc.mutation.EventGroupOtherNumAtRisk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAtRisk,
		})
		_node.EventGroupOtherNumAtRisk = value
	}
	if nodes := egc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventgroup.ParentTable,
			Columns: []string{eventgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adverseeventsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.adverse_events_module_event_group_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventGroupCreateBulk is the builder for creating many EventGroup entities in bulk.
type EventGroupCreateBulk struct {
	config
	builders []*EventGroupCreate
}

// Save creates the EventGroup entities in the database.
func (egcb *EventGroupCreateBulk) Save(ctx context.Context) ([]*EventGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(egcb.builders))
	nodes := make([]*EventGroup, len(egcb.builders))
	mutators := make([]Mutator, len(egcb.builders))
	for i := range egcb.builders {
		func(i int, root context.Context) {
			builder := egcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, egcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, egcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, egcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (egcb *EventGroupCreateBulk) SaveX(ctx context.Context) []*EventGroup {
	v, err := egcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (egcb *EventGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := egcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (egcb *EventGroupCreateBulk) ExecX(ctx context.Context) {
	if err := egcb.Exec(ctx); err != nil {
		panic(err)
	}
}
