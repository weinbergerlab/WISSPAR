// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventStatsCreate is the builder for creating a SeriousEventStats entity.
type SeriousEventStatsCreate struct {
	config
	mutation *SeriousEventStatsMutation
	hooks    []Hook
}

// SetSeriousEventStatsGroupID sets the "serious_event_stats_group_id" field.
func (sesc *SeriousEventStatsCreate) SetSeriousEventStatsGroupID(s string) *SeriousEventStatsCreate {
	sesc.mutation.SetSeriousEventStatsGroupID(s)
	return sesc
}

// SetSeriousEventStatsNumEvents sets the "serious_event_stats_num_events" field.
func (sesc *SeriousEventStatsCreate) SetSeriousEventStatsNumEvents(s string) *SeriousEventStatsCreate {
	sesc.mutation.SetSeriousEventStatsNumEvents(s)
	return sesc
}

// SetSeriousEventStatsNumAffected sets the "serious_event_stats_num_affected" field.
func (sesc *SeriousEventStatsCreate) SetSeriousEventStatsNumAffected(s string) *SeriousEventStatsCreate {
	sesc.mutation.SetSeriousEventStatsNumAffected(s)
	return sesc
}

// SetSeriousEventStatsNumAtRisk sets the "serious_event_stats_num_at_risk" field.
func (sesc *SeriousEventStatsCreate) SetSeriousEventStatsNumAtRisk(s string) *SeriousEventStatsCreate {
	sesc.mutation.SetSeriousEventStatsNumAtRisk(s)
	return sesc
}

// SetParentID sets the "parent" edge to the SeriousEvent entity by ID.
func (sesc *SeriousEventStatsCreate) SetParentID(id int) *SeriousEventStatsCreate {
	sesc.mutation.SetParentID(id)
	return sesc
}

// SetNillableParentID sets the "parent" edge to the SeriousEvent entity by ID if the given value is not nil.
func (sesc *SeriousEventStatsCreate) SetNillableParentID(id *int) *SeriousEventStatsCreate {
	if id != nil {
		sesc = sesc.SetParentID(*id)
	}
	return sesc
}

// SetParent sets the "parent" edge to the SeriousEvent entity.
func (sesc *SeriousEventStatsCreate) SetParent(s *SeriousEvent) *SeriousEventStatsCreate {
	return sesc.SetParentID(s.ID)
}

// Mutation returns the SeriousEventStatsMutation object of the builder.
func (sesc *SeriousEventStatsCreate) Mutation() *SeriousEventStatsMutation {
	return sesc.mutation
}

// Save creates the SeriousEventStats in the database.
func (sesc *SeriousEventStatsCreate) Save(ctx context.Context) (*SeriousEventStats, error) {
	var (
		err  error
		node *SeriousEventStats
	)
	if len(sesc.hooks) == 0 {
		if err = sesc.check(); err != nil {
			return nil, err
		}
		node, err = sesc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sesc.check(); err != nil {
				return nil, err
			}
			sesc.mutation = mutation
			if node, err = sesc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sesc.hooks) - 1; i >= 0; i-- {
			if sesc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sesc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sesc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sesc *SeriousEventStatsCreate) SaveX(ctx context.Context) *SeriousEventStats {
	v, err := sesc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sesc *SeriousEventStatsCreate) Exec(ctx context.Context) error {
	_, err := sesc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sesc *SeriousEventStatsCreate) ExecX(ctx context.Context) {
	if err := sesc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sesc *SeriousEventStatsCreate) check() error {
	if _, ok := sesc.mutation.SeriousEventStatsGroupID(); !ok {
		return &ValidationError{Name: "serious_event_stats_group_id", err: errors.New(`models: missing required field "SeriousEventStats.serious_event_stats_group_id"`)}
	}
	if _, ok := sesc.mutation.SeriousEventStatsNumEvents(); !ok {
		return &ValidationError{Name: "serious_event_stats_num_events", err: errors.New(`models: missing required field "SeriousEventStats.serious_event_stats_num_events"`)}
	}
	if _, ok := sesc.mutation.SeriousEventStatsNumAffected(); !ok {
		return &ValidationError{Name: "serious_event_stats_num_affected", err: errors.New(`models: missing required field "SeriousEventStats.serious_event_stats_num_affected"`)}
	}
	if _, ok := sesc.mutation.SeriousEventStatsNumAtRisk(); !ok {
		return &ValidationError{Name: "serious_event_stats_num_at_risk", err: errors.New(`models: missing required field "SeriousEventStats.serious_event_stats_num_at_risk"`)}
	}
	return nil
}

func (sesc *SeriousEventStatsCreate) sqlSave(ctx context.Context) (*SeriousEventStats, error) {
	_node, _spec := sesc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sesc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sesc *SeriousEventStatsCreate) createSpec() (*SeriousEventStats, *sqlgraph.CreateSpec) {
	var (
		_node = &SeriousEventStats{config: sesc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: seriouseventstats.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriouseventstats.FieldID,
			},
		}
	)
	if value, ok := sesc.mutation.SeriousEventStatsGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsGroupID,
		})
		_node.SeriousEventStatsGroupID = value
	}
	if value, ok := sesc.mutation.SeriousEventStatsNumEvents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumEvents,
		})
		_node.SeriousEventStatsNumEvents = value
	}
	if value, ok := sesc.mutation.SeriousEventStatsNumAffected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAffected,
		})
		_node.SeriousEventStatsNumAffected = value
	}
	if value, ok := sesc.mutation.SeriousEventStatsNumAtRisk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAtRisk,
		})
		_node.SeriousEventStatsNumAtRisk = value
	}
	if nodes := sesc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seriouseventstats.ParentTable,
			Columns: []string{seriouseventstats.ParentColumn},
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
		_node.serious_event_serious_event_stats_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SeriousEventStatsCreateBulk is the builder for creating many SeriousEventStats entities in bulk.
type SeriousEventStatsCreateBulk struct {
	config
	builders []*SeriousEventStatsCreate
}

// Save creates the SeriousEventStats entities in the database.
func (sescb *SeriousEventStatsCreateBulk) Save(ctx context.Context) ([]*SeriousEventStats, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sescb.builders))
	nodes := make([]*SeriousEventStats, len(sescb.builders))
	mutators := make([]Mutator, len(sescb.builders))
	for i := range sescb.builders {
		func(i int, root context.Context) {
			builder := sescb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SeriousEventStatsMutation)
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
					_, err = mutators[i+1].Mutate(root, sescb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sescb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sescb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sescb *SeriousEventStatsCreateBulk) SaveX(ctx context.Context) []*SeriousEventStats {
	v, err := sescb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sescb *SeriousEventStatsCreateBulk) Exec(ctx context.Context) error {
	_, err := sescb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sescb *SeriousEventStatsCreateBulk) ExecX(ctx context.Context) {
	if err := sescb.Exec(ctx); err != nil {
		panic(err)
	}
}
