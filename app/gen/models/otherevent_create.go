// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
)

// OtherEventCreate is the builder for creating a OtherEvent entity.
type OtherEventCreate struct {
	config
	mutation *OtherEventMutation
	hooks    []Hook
}

// SetOtherEventTerm sets the "other_event_term" field.
func (oec *OtherEventCreate) SetOtherEventTerm(s string) *OtherEventCreate {
	oec.mutation.SetOtherEventTerm(s)
	return oec
}

// SetOtherEventOrganSystem sets the "other_event_organ_system" field.
func (oec *OtherEventCreate) SetOtherEventOrganSystem(s string) *OtherEventCreate {
	oec.mutation.SetOtherEventOrganSystem(s)
	return oec
}

// SetOtherEventSourceVocabulary sets the "other_event_source_vocabulary" field.
func (oec *OtherEventCreate) SetOtherEventSourceVocabulary(s string) *OtherEventCreate {
	oec.mutation.SetOtherEventSourceVocabulary(s)
	return oec
}

// SetOtherEventAssessmentType sets the "other_event_assessment_type" field.
func (oec *OtherEventCreate) SetOtherEventAssessmentType(s string) *OtherEventCreate {
	oec.mutation.SetOtherEventAssessmentType(s)
	return oec
}

// SetOtherEventNotes sets the "other_event_notes" field.
func (oec *OtherEventCreate) SetOtherEventNotes(s string) *OtherEventCreate {
	oec.mutation.SetOtherEventNotes(s)
	return oec
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (oec *OtherEventCreate) SetParentID(id int) *OtherEventCreate {
	oec.mutation.SetParentID(id)
	return oec
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (oec *OtherEventCreate) SetNillableParentID(id *int) *OtherEventCreate {
	if id != nil {
		oec = oec.SetParentID(*id)
	}
	return oec
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (oec *OtherEventCreate) SetParent(a *AdverseEventsModule) *OtherEventCreate {
	return oec.SetParentID(a.ID)
}

// AddOtherEventStatsListIDs adds the "other_event_stats_list" edge to the OtherEventStats entity by IDs.
func (oec *OtherEventCreate) AddOtherEventStatsListIDs(ids ...int) *OtherEventCreate {
	oec.mutation.AddOtherEventStatsListIDs(ids...)
	return oec
}

// AddOtherEventStatsList adds the "other_event_stats_list" edges to the OtherEventStats entity.
func (oec *OtherEventCreate) AddOtherEventStatsList(o ...*OtherEventStats) *OtherEventCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oec.AddOtherEventStatsListIDs(ids...)
}

// Mutation returns the OtherEventMutation object of the builder.
func (oec *OtherEventCreate) Mutation() *OtherEventMutation {
	return oec.mutation
}

// Save creates the OtherEvent in the database.
func (oec *OtherEventCreate) Save(ctx context.Context) (*OtherEvent, error) {
	var (
		err  error
		node *OtherEvent
	)
	if len(oec.hooks) == 0 {
		if err = oec.check(); err != nil {
			return nil, err
		}
		node, err = oec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oec.check(); err != nil {
				return nil, err
			}
			oec.mutation = mutation
			if node, err = oec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oec.hooks) - 1; i >= 0; i-- {
			if oec.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oec *OtherEventCreate) SaveX(ctx context.Context) *OtherEvent {
	v, err := oec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oec *OtherEventCreate) Exec(ctx context.Context) error {
	_, err := oec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oec *OtherEventCreate) ExecX(ctx context.Context) {
	if err := oec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oec *OtherEventCreate) check() error {
	if _, ok := oec.mutation.OtherEventTerm(); !ok {
		return &ValidationError{Name: "other_event_term", err: errors.New(`models: missing required field "OtherEvent.other_event_term"`)}
	}
	if _, ok := oec.mutation.OtherEventOrganSystem(); !ok {
		return &ValidationError{Name: "other_event_organ_system", err: errors.New(`models: missing required field "OtherEvent.other_event_organ_system"`)}
	}
	if _, ok := oec.mutation.OtherEventSourceVocabulary(); !ok {
		return &ValidationError{Name: "other_event_source_vocabulary", err: errors.New(`models: missing required field "OtherEvent.other_event_source_vocabulary"`)}
	}
	if _, ok := oec.mutation.OtherEventAssessmentType(); !ok {
		return &ValidationError{Name: "other_event_assessment_type", err: errors.New(`models: missing required field "OtherEvent.other_event_assessment_type"`)}
	}
	if _, ok := oec.mutation.OtherEventNotes(); !ok {
		return &ValidationError{Name: "other_event_notes", err: errors.New(`models: missing required field "OtherEvent.other_event_notes"`)}
	}
	return nil
}

func (oec *OtherEventCreate) sqlSave(ctx context.Context) (*OtherEvent, error) {
	_node, _spec := oec.createSpec()
	if err := sqlgraph.CreateNode(ctx, oec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oec *OtherEventCreate) createSpec() (*OtherEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &OtherEvent{config: oec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: otherevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: otherevent.FieldID,
			},
		}
	)
	if value, ok := oec.mutation.OtherEventTerm(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventTerm,
		})
		_node.OtherEventTerm = value
	}
	if value, ok := oec.mutation.OtherEventOrganSystem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventOrganSystem,
		})
		_node.OtherEventOrganSystem = value
	}
	if value, ok := oec.mutation.OtherEventSourceVocabulary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventSourceVocabulary,
		})
		_node.OtherEventSourceVocabulary = value
	}
	if value, ok := oec.mutation.OtherEventAssessmentType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventAssessmentType,
		})
		_node.OtherEventAssessmentType = value
	}
	if value, ok := oec.mutation.OtherEventNotes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventNotes,
		})
		_node.OtherEventNotes = value
	}
	if nodes := oec.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   otherevent.ParentTable,
			Columns: []string{otherevent.ParentColumn},
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
		_node.adverse_events_module_other_event_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oec.mutation.OtherEventStatsListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   otherevent.OtherEventStatsListTable,
			Columns: []string{otherevent.OtherEventStatsListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: othereventstats.FieldID,
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

// OtherEventCreateBulk is the builder for creating many OtherEvent entities in bulk.
type OtherEventCreateBulk struct {
	config
	builders []*OtherEventCreate
}

// Save creates the OtherEvent entities in the database.
func (oecb *OtherEventCreateBulk) Save(ctx context.Context) ([]*OtherEvent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oecb.builders))
	nodes := make([]*OtherEvent, len(oecb.builders))
	mutators := make([]Mutator, len(oecb.builders))
	for i := range oecb.builders {
		func(i int, root context.Context) {
			builder := oecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OtherEventMutation)
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
					_, err = mutators[i+1].Mutate(root, oecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oecb *OtherEventCreateBulk) SaveX(ctx context.Context) []*OtherEvent {
	v, err := oecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oecb *OtherEventCreateBulk) Exec(ctx context.Context) error {
	_, err := oecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oecb *OtherEventCreateBulk) ExecX(ctx context.Context) {
	if err := oecb.Exec(ctx); err != nil {
		panic(err)
	}
}
