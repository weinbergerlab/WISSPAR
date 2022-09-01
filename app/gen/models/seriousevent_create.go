// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventCreate is the builder for creating a SeriousEvent entity.
type SeriousEventCreate struct {
	config
	mutation *SeriousEventMutation
	hooks    []Hook
}

// SetSeriousEventTerm sets the "serious_event_term" field.
func (sec *SeriousEventCreate) SetSeriousEventTerm(s string) *SeriousEventCreate {
	sec.mutation.SetSeriousEventTerm(s)
	return sec
}

// SetSeriousEventOrganSystem sets the "serious_event_organ_system" field.
func (sec *SeriousEventCreate) SetSeriousEventOrganSystem(s string) *SeriousEventCreate {
	sec.mutation.SetSeriousEventOrganSystem(s)
	return sec
}

// SetSeriousEventSourceVocabulary sets the "serious_event_source_vocabulary" field.
func (sec *SeriousEventCreate) SetSeriousEventSourceVocabulary(s string) *SeriousEventCreate {
	sec.mutation.SetSeriousEventSourceVocabulary(s)
	return sec
}

// SetSeriousEventAssessmentType sets the "serious_event_assessment_type" field.
func (sec *SeriousEventCreate) SetSeriousEventAssessmentType(s string) *SeriousEventCreate {
	sec.mutation.SetSeriousEventAssessmentType(s)
	return sec
}

// SetSeriousEventNotes sets the "serious_event_notes" field.
func (sec *SeriousEventCreate) SetSeriousEventNotes(s string) *SeriousEventCreate {
	sec.mutation.SetSeriousEventNotes(s)
	return sec
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (sec *SeriousEventCreate) SetParentID(id int) *SeriousEventCreate {
	sec.mutation.SetParentID(id)
	return sec
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (sec *SeriousEventCreate) SetNillableParentID(id *int) *SeriousEventCreate {
	if id != nil {
		sec = sec.SetParentID(*id)
	}
	return sec
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (sec *SeriousEventCreate) SetParent(a *AdverseEventsModule) *SeriousEventCreate {
	return sec.SetParentID(a.ID)
}

// AddSeriousEventStatsListIDs adds the "serious_event_stats_list" edge to the SeriousEventStats entity by IDs.
func (sec *SeriousEventCreate) AddSeriousEventStatsListIDs(ids ...int) *SeriousEventCreate {
	sec.mutation.AddSeriousEventStatsListIDs(ids...)
	return sec
}

// AddSeriousEventStatsList adds the "serious_event_stats_list" edges to the SeriousEventStats entity.
func (sec *SeriousEventCreate) AddSeriousEventStatsList(s ...*SeriousEventStats) *SeriousEventCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sec.AddSeriousEventStatsListIDs(ids...)
}

// Mutation returns the SeriousEventMutation object of the builder.
func (sec *SeriousEventCreate) Mutation() *SeriousEventMutation {
	return sec.mutation
}

// Save creates the SeriousEvent in the database.
func (sec *SeriousEventCreate) Save(ctx context.Context) (*SeriousEvent, error) {
	var (
		err  error
		node *SeriousEvent
	)
	if len(sec.hooks) == 0 {
		if err = sec.check(); err != nil {
			return nil, err
		}
		node, err = sec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sec.check(); err != nil {
				return nil, err
			}
			sec.mutation = mutation
			if node, err = sec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sec.hooks) - 1; i >= 0; i-- {
			if sec.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sec *SeriousEventCreate) SaveX(ctx context.Context) *SeriousEvent {
	v, err := sec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sec *SeriousEventCreate) Exec(ctx context.Context) error {
	_, err := sec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sec *SeriousEventCreate) ExecX(ctx context.Context) {
	if err := sec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sec *SeriousEventCreate) check() error {
	if _, ok := sec.mutation.SeriousEventTerm(); !ok {
		return &ValidationError{Name: "serious_event_term", err: errors.New(`models: missing required field "SeriousEvent.serious_event_term"`)}
	}
	if _, ok := sec.mutation.SeriousEventOrganSystem(); !ok {
		return &ValidationError{Name: "serious_event_organ_system", err: errors.New(`models: missing required field "SeriousEvent.serious_event_organ_system"`)}
	}
	if _, ok := sec.mutation.SeriousEventSourceVocabulary(); !ok {
		return &ValidationError{Name: "serious_event_source_vocabulary", err: errors.New(`models: missing required field "SeriousEvent.serious_event_source_vocabulary"`)}
	}
	if _, ok := sec.mutation.SeriousEventAssessmentType(); !ok {
		return &ValidationError{Name: "serious_event_assessment_type", err: errors.New(`models: missing required field "SeriousEvent.serious_event_assessment_type"`)}
	}
	if _, ok := sec.mutation.SeriousEventNotes(); !ok {
		return &ValidationError{Name: "serious_event_notes", err: errors.New(`models: missing required field "SeriousEvent.serious_event_notes"`)}
	}
	return nil
}

func (sec *SeriousEventCreate) sqlSave(ctx context.Context) (*SeriousEvent, error) {
	_node, _spec := sec.createSpec()
	if err := sqlgraph.CreateNode(ctx, sec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sec *SeriousEventCreate) createSpec() (*SeriousEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &SeriousEvent{config: sec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: seriousevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriousevent.FieldID,
			},
		}
	)
	if value, ok := sec.mutation.SeriousEventTerm(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventTerm,
		})
		_node.SeriousEventTerm = value
	}
	if value, ok := sec.mutation.SeriousEventOrganSystem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventOrganSystem,
		})
		_node.SeriousEventOrganSystem = value
	}
	if value, ok := sec.mutation.SeriousEventSourceVocabulary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventSourceVocabulary,
		})
		_node.SeriousEventSourceVocabulary = value
	}
	if value, ok := sec.mutation.SeriousEventAssessmentType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventAssessmentType,
		})
		_node.SeriousEventAssessmentType = value
	}
	if value, ok := sec.mutation.SeriousEventNotes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventNotes,
		})
		_node.SeriousEventNotes = value
	}
	if nodes := sec.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seriousevent.ParentTable,
			Columns: []string{seriousevent.ParentColumn},
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
		_node.adverse_events_module_serious_event_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sec.mutation.SeriousEventStatsListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seriousevent.SeriousEventStatsListTable,
			Columns: []string{seriousevent.SeriousEventStatsListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: seriouseventstats.FieldID,
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

// SeriousEventCreateBulk is the builder for creating many SeriousEvent entities in bulk.
type SeriousEventCreateBulk struct {
	config
	builders []*SeriousEventCreate
}

// Save creates the SeriousEvent entities in the database.
func (secb *SeriousEventCreateBulk) Save(ctx context.Context) ([]*SeriousEvent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(secb.builders))
	nodes := make([]*SeriousEvent, len(secb.builders))
	mutators := make([]Mutator, len(secb.builders))
	for i := range secb.builders {
		func(i int, root context.Context) {
			builder := secb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SeriousEventMutation)
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
					_, err = mutators[i+1].Mutate(root, secb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, secb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, secb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (secb *SeriousEventCreateBulk) SaveX(ctx context.Context) []*SeriousEvent {
	v, err := secb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (secb *SeriousEventCreateBulk) Exec(ctx context.Context) error {
	_, err := secb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (secb *SeriousEventCreateBulk) ExecX(ctx context.Context) {
	if err := secb.Exec(ctx); err != nil {
		panic(err)
	}
}
