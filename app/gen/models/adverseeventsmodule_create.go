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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
)

// AdverseEventsModuleCreate is the builder for creating a AdverseEventsModule entity.
type AdverseEventsModuleCreate struct {
	config
	mutation *AdverseEventsModuleMutation
	hooks    []Hook
}

// SetEventsFrequencyThreshold sets the "events_frequency_threshold" field.
func (aemc *AdverseEventsModuleCreate) SetEventsFrequencyThreshold(s string) *AdverseEventsModuleCreate {
	aemc.mutation.SetEventsFrequencyThreshold(s)
	return aemc
}

// SetEventsTimeFrame sets the "events_time_frame" field.
func (aemc *AdverseEventsModuleCreate) SetEventsTimeFrame(s string) *AdverseEventsModuleCreate {
	aemc.mutation.SetEventsTimeFrame(s)
	return aemc
}

// SetEventsDescription sets the "events_description" field.
func (aemc *AdverseEventsModuleCreate) SetEventsDescription(s string) *AdverseEventsModuleCreate {
	aemc.mutation.SetEventsDescription(s)
	return aemc
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (aemc *AdverseEventsModuleCreate) SetParentID(id int) *AdverseEventsModuleCreate {
	aemc.mutation.SetParentID(id)
	return aemc
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (aemc *AdverseEventsModuleCreate) SetNillableParentID(id *int) *AdverseEventsModuleCreate {
	if id != nil {
		aemc = aemc.SetParentID(*id)
	}
	return aemc
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (aemc *AdverseEventsModuleCreate) SetParent(r *ResultsDefinition) *AdverseEventsModuleCreate {
	return aemc.SetParentID(r.ID)
}

// AddEventGroupListIDs adds the "event_group_list" edge to the EventGroup entity by IDs.
func (aemc *AdverseEventsModuleCreate) AddEventGroupListIDs(ids ...int) *AdverseEventsModuleCreate {
	aemc.mutation.AddEventGroupListIDs(ids...)
	return aemc
}

// AddEventGroupList adds the "event_group_list" edges to the EventGroup entity.
func (aemc *AdverseEventsModuleCreate) AddEventGroupList(e ...*EventGroup) *AdverseEventsModuleCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return aemc.AddEventGroupListIDs(ids...)
}

// AddSeriousEventListIDs adds the "serious_event_list" edge to the SeriousEvent entity by IDs.
func (aemc *AdverseEventsModuleCreate) AddSeriousEventListIDs(ids ...int) *AdverseEventsModuleCreate {
	aemc.mutation.AddSeriousEventListIDs(ids...)
	return aemc
}

// AddSeriousEventList adds the "serious_event_list" edges to the SeriousEvent entity.
func (aemc *AdverseEventsModuleCreate) AddSeriousEventList(s ...*SeriousEvent) *AdverseEventsModuleCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aemc.AddSeriousEventListIDs(ids...)
}

// AddOtherEventListIDs adds the "other_event_list" edge to the OtherEvent entity by IDs.
func (aemc *AdverseEventsModuleCreate) AddOtherEventListIDs(ids ...int) *AdverseEventsModuleCreate {
	aemc.mutation.AddOtherEventListIDs(ids...)
	return aemc
}

// AddOtherEventList adds the "other_event_list" edges to the OtherEvent entity.
func (aemc *AdverseEventsModuleCreate) AddOtherEventList(o ...*OtherEvent) *AdverseEventsModuleCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return aemc.AddOtherEventListIDs(ids...)
}

// Mutation returns the AdverseEventsModuleMutation object of the builder.
func (aemc *AdverseEventsModuleCreate) Mutation() *AdverseEventsModuleMutation {
	return aemc.mutation
}

// Save creates the AdverseEventsModule in the database.
func (aemc *AdverseEventsModuleCreate) Save(ctx context.Context) (*AdverseEventsModule, error) {
	var (
		err  error
		node *AdverseEventsModule
	)
	if len(aemc.hooks) == 0 {
		if err = aemc.check(); err != nil {
			return nil, err
		}
		node, err = aemc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdverseEventsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aemc.check(); err != nil {
				return nil, err
			}
			aemc.mutation = mutation
			if node, err = aemc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aemc.hooks) - 1; i >= 0; i-- {
			if aemc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = aemc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aemc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aemc *AdverseEventsModuleCreate) SaveX(ctx context.Context) *AdverseEventsModule {
	v, err := aemc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aemc *AdverseEventsModuleCreate) Exec(ctx context.Context) error {
	_, err := aemc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aemc *AdverseEventsModuleCreate) ExecX(ctx context.Context) {
	if err := aemc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aemc *AdverseEventsModuleCreate) check() error {
	if _, ok := aemc.mutation.EventsFrequencyThreshold(); !ok {
		return &ValidationError{Name: "events_frequency_threshold", err: errors.New(`models: missing required field "AdverseEventsModule.events_frequency_threshold"`)}
	}
	if _, ok := aemc.mutation.EventsTimeFrame(); !ok {
		return &ValidationError{Name: "events_time_frame", err: errors.New(`models: missing required field "AdverseEventsModule.events_time_frame"`)}
	}
	if _, ok := aemc.mutation.EventsDescription(); !ok {
		return &ValidationError{Name: "events_description", err: errors.New(`models: missing required field "AdverseEventsModule.events_description"`)}
	}
	return nil
}

func (aemc *AdverseEventsModuleCreate) sqlSave(ctx context.Context) (*AdverseEventsModule, error) {
	_node, _spec := aemc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aemc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (aemc *AdverseEventsModuleCreate) createSpec() (*AdverseEventsModule, *sqlgraph.CreateSpec) {
	var (
		_node = &AdverseEventsModule{config: aemc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: adverseeventsmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adverseeventsmodule.FieldID,
			},
		}
	)
	if value, ok := aemc.mutation.EventsFrequencyThreshold(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsFrequencyThreshold,
		})
		_node.EventsFrequencyThreshold = value
	}
	if value, ok := aemc.mutation.EventsTimeFrame(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsTimeFrame,
		})
		_node.EventsTimeFrame = value
	}
	if value, ok := aemc.mutation.EventsDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsDescription,
		})
		_node.EventsDescription = value
	}
	if nodes := aemc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   adverseeventsmodule.ParentTable,
			Columns: []string{adverseeventsmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.results_definition_adverse_events_module = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aemc.mutation.EventGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adverseeventsmodule.EventGroupListTable,
			Columns: []string{adverseeventsmodule.EventGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eventgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aemc.mutation.SeriousEventListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adverseeventsmodule.SeriousEventListTable,
			Columns: []string{adverseeventsmodule.SeriousEventListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: seriousevent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aemc.mutation.OtherEventListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adverseeventsmodule.OtherEventListTable,
			Columns: []string{adverseeventsmodule.OtherEventListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: otherevent.FieldID,
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

// AdverseEventsModuleCreateBulk is the builder for creating many AdverseEventsModule entities in bulk.
type AdverseEventsModuleCreateBulk struct {
	config
	builders []*AdverseEventsModuleCreate
}

// Save creates the AdverseEventsModule entities in the database.
func (aemcb *AdverseEventsModuleCreateBulk) Save(ctx context.Context) ([]*AdverseEventsModule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aemcb.builders))
	nodes := make([]*AdverseEventsModule, len(aemcb.builders))
	mutators := make([]Mutator, len(aemcb.builders))
	for i := range aemcb.builders {
		func(i int, root context.Context) {
			builder := aemcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdverseEventsModuleMutation)
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
					_, err = mutators[i+1].Mutate(root, aemcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aemcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aemcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aemcb *AdverseEventsModuleCreateBulk) SaveX(ctx context.Context) []*AdverseEventsModule {
	v, err := aemcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aemcb *AdverseEventsModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := aemcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aemcb *AdverseEventsModuleCreateBulk) ExecX(ctx context.Context) {
	if err := aemcb.Exec(ctx); err != nil {
		panic(err)
	}
}
