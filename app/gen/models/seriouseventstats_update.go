// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventStatsUpdate is the builder for updating SeriousEventStats entities.
type SeriousEventStatsUpdate struct {
	config
	hooks    []Hook
	mutation *SeriousEventStatsMutation
}

// Where appends a list predicates to the SeriousEventStatsUpdate builder.
func (sesu *SeriousEventStatsUpdate) Where(ps ...predicate.SeriousEventStats) *SeriousEventStatsUpdate {
	sesu.mutation.Where(ps...)
	return sesu
}

// SetSeriousEventStatsGroupID sets the "serious_event_stats_group_id" field.
func (sesu *SeriousEventStatsUpdate) SetSeriousEventStatsGroupID(s string) *SeriousEventStatsUpdate {
	sesu.mutation.SetSeriousEventStatsGroupID(s)
	return sesu
}

// SetSeriousEventStatsNumEvents sets the "serious_event_stats_num_events" field.
func (sesu *SeriousEventStatsUpdate) SetSeriousEventStatsNumEvents(s string) *SeriousEventStatsUpdate {
	sesu.mutation.SetSeriousEventStatsNumEvents(s)
	return sesu
}

// SetSeriousEventStatsNumAffected sets the "serious_event_stats_num_affected" field.
func (sesu *SeriousEventStatsUpdate) SetSeriousEventStatsNumAffected(s string) *SeriousEventStatsUpdate {
	sesu.mutation.SetSeriousEventStatsNumAffected(s)
	return sesu
}

// SetSeriousEventStatsNumAtRisk sets the "serious_event_stats_num_at_risk" field.
func (sesu *SeriousEventStatsUpdate) SetSeriousEventStatsNumAtRisk(s string) *SeriousEventStatsUpdate {
	sesu.mutation.SetSeriousEventStatsNumAtRisk(s)
	return sesu
}

// SetParentID sets the "parent" edge to the SeriousEvent entity by ID.
func (sesu *SeriousEventStatsUpdate) SetParentID(id int) *SeriousEventStatsUpdate {
	sesu.mutation.SetParentID(id)
	return sesu
}

// SetNillableParentID sets the "parent" edge to the SeriousEvent entity by ID if the given value is not nil.
func (sesu *SeriousEventStatsUpdate) SetNillableParentID(id *int) *SeriousEventStatsUpdate {
	if id != nil {
		sesu = sesu.SetParentID(*id)
	}
	return sesu
}

// SetParent sets the "parent" edge to the SeriousEvent entity.
func (sesu *SeriousEventStatsUpdate) SetParent(s *SeriousEvent) *SeriousEventStatsUpdate {
	return sesu.SetParentID(s.ID)
}

// Mutation returns the SeriousEventStatsMutation object of the builder.
func (sesu *SeriousEventStatsUpdate) Mutation() *SeriousEventStatsMutation {
	return sesu.mutation
}

// ClearParent clears the "parent" edge to the SeriousEvent entity.
func (sesu *SeriousEventStatsUpdate) ClearParent() *SeriousEventStatsUpdate {
	sesu.mutation.ClearParent()
	return sesu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sesu *SeriousEventStatsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sesu.hooks) == 0 {
		affected, err = sesu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sesu.mutation = mutation
			affected, err = sesu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sesu.hooks) - 1; i >= 0; i-- {
			if sesu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sesu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sesu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sesu *SeriousEventStatsUpdate) SaveX(ctx context.Context) int {
	affected, err := sesu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sesu *SeriousEventStatsUpdate) Exec(ctx context.Context) error {
	_, err := sesu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sesu *SeriousEventStatsUpdate) ExecX(ctx context.Context) {
	if err := sesu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sesu *SeriousEventStatsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriouseventstats.Table,
			Columns: seriouseventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriouseventstats.FieldID,
			},
		},
	}
	if ps := sesu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sesu.mutation.SeriousEventStatsGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsGroupID,
		})
	}
	if value, ok := sesu.mutation.SeriousEventStatsNumEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumEvents,
		})
	}
	if value, ok := sesu.mutation.SeriousEventStatsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAffected,
		})
	}
	if value, ok := sesu.mutation.SeriousEventStatsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAtRisk,
		})
	}
	if sesu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sesu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sesu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seriouseventstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SeriousEventStatsUpdateOne is the builder for updating a single SeriousEventStats entity.
type SeriousEventStatsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeriousEventStatsMutation
}

// SetSeriousEventStatsGroupID sets the "serious_event_stats_group_id" field.
func (sesuo *SeriousEventStatsUpdateOne) SetSeriousEventStatsGroupID(s string) *SeriousEventStatsUpdateOne {
	sesuo.mutation.SetSeriousEventStatsGroupID(s)
	return sesuo
}

// SetSeriousEventStatsNumEvents sets the "serious_event_stats_num_events" field.
func (sesuo *SeriousEventStatsUpdateOne) SetSeriousEventStatsNumEvents(s string) *SeriousEventStatsUpdateOne {
	sesuo.mutation.SetSeriousEventStatsNumEvents(s)
	return sesuo
}

// SetSeriousEventStatsNumAffected sets the "serious_event_stats_num_affected" field.
func (sesuo *SeriousEventStatsUpdateOne) SetSeriousEventStatsNumAffected(s string) *SeriousEventStatsUpdateOne {
	sesuo.mutation.SetSeriousEventStatsNumAffected(s)
	return sesuo
}

// SetSeriousEventStatsNumAtRisk sets the "serious_event_stats_num_at_risk" field.
func (sesuo *SeriousEventStatsUpdateOne) SetSeriousEventStatsNumAtRisk(s string) *SeriousEventStatsUpdateOne {
	sesuo.mutation.SetSeriousEventStatsNumAtRisk(s)
	return sesuo
}

// SetParentID sets the "parent" edge to the SeriousEvent entity by ID.
func (sesuo *SeriousEventStatsUpdateOne) SetParentID(id int) *SeriousEventStatsUpdateOne {
	sesuo.mutation.SetParentID(id)
	return sesuo
}

// SetNillableParentID sets the "parent" edge to the SeriousEvent entity by ID if the given value is not nil.
func (sesuo *SeriousEventStatsUpdateOne) SetNillableParentID(id *int) *SeriousEventStatsUpdateOne {
	if id != nil {
		sesuo = sesuo.SetParentID(*id)
	}
	return sesuo
}

// SetParent sets the "parent" edge to the SeriousEvent entity.
func (sesuo *SeriousEventStatsUpdateOne) SetParent(s *SeriousEvent) *SeriousEventStatsUpdateOne {
	return sesuo.SetParentID(s.ID)
}

// Mutation returns the SeriousEventStatsMutation object of the builder.
func (sesuo *SeriousEventStatsUpdateOne) Mutation() *SeriousEventStatsMutation {
	return sesuo.mutation
}

// ClearParent clears the "parent" edge to the SeriousEvent entity.
func (sesuo *SeriousEventStatsUpdateOne) ClearParent() *SeriousEventStatsUpdateOne {
	sesuo.mutation.ClearParent()
	return sesuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sesuo *SeriousEventStatsUpdateOne) Select(field string, fields ...string) *SeriousEventStatsUpdateOne {
	sesuo.fields = append([]string{field}, fields...)
	return sesuo
}

// Save executes the query and returns the updated SeriousEventStats entity.
func (sesuo *SeriousEventStatsUpdateOne) Save(ctx context.Context) (*SeriousEventStats, error) {
	var (
		err  error
		node *SeriousEventStats
	)
	if len(sesuo.hooks) == 0 {
		node, err = sesuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sesuo.mutation = mutation
			node, err = sesuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sesuo.hooks) - 1; i >= 0; i-- {
			if sesuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sesuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sesuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sesuo *SeriousEventStatsUpdateOne) SaveX(ctx context.Context) *SeriousEventStats {
	node, err := sesuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sesuo *SeriousEventStatsUpdateOne) Exec(ctx context.Context) error {
	_, err := sesuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sesuo *SeriousEventStatsUpdateOne) ExecX(ctx context.Context) {
	if err := sesuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sesuo *SeriousEventStatsUpdateOne) sqlSave(ctx context.Context) (_node *SeriousEventStats, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriouseventstats.Table,
			Columns: seriouseventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriouseventstats.FieldID,
			},
		},
	}
	id, ok := sesuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "SeriousEventStats.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sesuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seriouseventstats.FieldID)
		for _, f := range fields {
			if !seriouseventstats.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != seriouseventstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sesuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sesuo.mutation.SeriousEventStatsGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsGroupID,
		})
	}
	if value, ok := sesuo.mutation.SeriousEventStatsNumEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumEvents,
		})
	}
	if value, ok := sesuo.mutation.SeriousEventStatsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAffected,
		})
	}
	if value, ok := sesuo.mutation.SeriousEventStatsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriouseventstats.FieldSeriousEventStatsNumAtRisk,
		})
	}
	if sesuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sesuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SeriousEventStats{config: sesuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sesuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seriouseventstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
