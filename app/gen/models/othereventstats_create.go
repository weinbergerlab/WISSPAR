// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
)

// OtherEventStatsCreate is the builder for creating a OtherEventStats entity.
type OtherEventStatsCreate struct {
	config
	mutation *OtherEventStatsMutation
	hooks    []Hook
}

// SetOtherEventStatsGroupID sets the "other_event_stats_group_id" field.
func (oesc *OtherEventStatsCreate) SetOtherEventStatsGroupID(s string) *OtherEventStatsCreate {
	oesc.mutation.SetOtherEventStatsGroupID(s)
	return oesc
}

// SetOtherEventStatsNumEvents sets the "other_event_stats_num_events" field.
func (oesc *OtherEventStatsCreate) SetOtherEventStatsNumEvents(s string) *OtherEventStatsCreate {
	oesc.mutation.SetOtherEventStatsNumEvents(s)
	return oesc
}

// SetOtherEventStatsNumAffected sets the "other_event_stats_num_affected" field.
func (oesc *OtherEventStatsCreate) SetOtherEventStatsNumAffected(s string) *OtherEventStatsCreate {
	oesc.mutation.SetOtherEventStatsNumAffected(s)
	return oesc
}

// SetOtherEventStatsNumAtRisk sets the "other_event_stats_num_at_risk" field.
func (oesc *OtherEventStatsCreate) SetOtherEventStatsNumAtRisk(s string) *OtherEventStatsCreate {
	oesc.mutation.SetOtherEventStatsNumAtRisk(s)
	return oesc
}

// SetParentID sets the "parent" edge to the OtherEvent entity by ID.
func (oesc *OtherEventStatsCreate) SetParentID(id int) *OtherEventStatsCreate {
	oesc.mutation.SetParentID(id)
	return oesc
}

// SetNillableParentID sets the "parent" edge to the OtherEvent entity by ID if the given value is not nil.
func (oesc *OtherEventStatsCreate) SetNillableParentID(id *int) *OtherEventStatsCreate {
	if id != nil {
		oesc = oesc.SetParentID(*id)
	}
	return oesc
}

// SetParent sets the "parent" edge to the OtherEvent entity.
func (oesc *OtherEventStatsCreate) SetParent(o *OtherEvent) *OtherEventStatsCreate {
	return oesc.SetParentID(o.ID)
}

// Mutation returns the OtherEventStatsMutation object of the builder.
func (oesc *OtherEventStatsCreate) Mutation() *OtherEventStatsMutation {
	return oesc.mutation
}

// Save creates the OtherEventStats in the database.
func (oesc *OtherEventStatsCreate) Save(ctx context.Context) (*OtherEventStats, error) {
	var (
		err  error
		node *OtherEventStats
	)
	if len(oesc.hooks) == 0 {
		if err = oesc.check(); err != nil {
			return nil, err
		}
		node, err = oesc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oesc.check(); err != nil {
				return nil, err
			}
			oesc.mutation = mutation
			if node, err = oesc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oesc.hooks) - 1; i >= 0; i-- {
			if oesc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oesc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oesc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oesc *OtherEventStatsCreate) SaveX(ctx context.Context) *OtherEventStats {
	v, err := oesc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oesc *OtherEventStatsCreate) Exec(ctx context.Context) error {
	_, err := oesc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oesc *OtherEventStatsCreate) ExecX(ctx context.Context) {
	if err := oesc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oesc *OtherEventStatsCreate) check() error {
	if _, ok := oesc.mutation.OtherEventStatsGroupID(); !ok {
		return &ValidationError{Name: "other_event_stats_group_id", err: errors.New(`models: missing required field "OtherEventStats.other_event_stats_group_id"`)}
	}
	if _, ok := oesc.mutation.OtherEventStatsNumEvents(); !ok {
		return &ValidationError{Name: "other_event_stats_num_events", err: errors.New(`models: missing required field "OtherEventStats.other_event_stats_num_events"`)}
	}
	if _, ok := oesc.mutation.OtherEventStatsNumAffected(); !ok {
		return &ValidationError{Name: "other_event_stats_num_affected", err: errors.New(`models: missing required field "OtherEventStats.other_event_stats_num_affected"`)}
	}
	if _, ok := oesc.mutation.OtherEventStatsNumAtRisk(); !ok {
		return &ValidationError{Name: "other_event_stats_num_at_risk", err: errors.New(`models: missing required field "OtherEventStats.other_event_stats_num_at_risk"`)}
	}
	return nil
}

func (oesc *OtherEventStatsCreate) sqlSave(ctx context.Context) (*OtherEventStats, error) {
	_node, _spec := oesc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oesc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oesc *OtherEventStatsCreate) createSpec() (*OtherEventStats, *sqlgraph.CreateSpec) {
	var (
		_node = &OtherEventStats{config: oesc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: othereventstats.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: othereventstats.FieldID,
			},
		}
	)
	if value, ok := oesc.mutation.OtherEventStatsGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsGroupID,
		})
		_node.OtherEventStatsGroupID = value
	}
	if value, ok := oesc.mutation.OtherEventStatsNumEvents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumEvents,
		})
		_node.OtherEventStatsNumEvents = value
	}
	if value, ok := oesc.mutation.OtherEventStatsNumAffected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAffected,
		})
		_node.OtherEventStatsNumAffected = value
	}
	if value, ok := oesc.mutation.OtherEventStatsNumAtRisk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAtRisk,
		})
		_node.OtherEventStatsNumAtRisk = value
	}
	if nodes := oesc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   othereventstats.ParentTable,
			Columns: []string{othereventstats.ParentColumn},
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
		_node.other_event_other_event_stats_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OtherEventStatsCreateBulk is the builder for creating many OtherEventStats entities in bulk.
type OtherEventStatsCreateBulk struct {
	config
	builders []*OtherEventStatsCreate
}

// Save creates the OtherEventStats entities in the database.
func (oescb *OtherEventStatsCreateBulk) Save(ctx context.Context) ([]*OtherEventStats, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oescb.builders))
	nodes := make([]*OtherEventStats, len(oescb.builders))
	mutators := make([]Mutator, len(oescb.builders))
	for i := range oescb.builders {
		func(i int, root context.Context) {
			builder := oescb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OtherEventStatsMutation)
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
					_, err = mutators[i+1].Mutate(root, oescb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oescb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oescb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oescb *OtherEventStatsCreateBulk) SaveX(ctx context.Context) []*OtherEventStats {
	v, err := oescb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oescb *OtherEventStatsCreateBulk) Exec(ctx context.Context) error {
	_, err := oescb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oescb *OtherEventStatsCreateBulk) ExecX(ctx context.Context) {
	if err := oescb.Exec(ctx); err != nil {
		panic(err)
	}
}
