// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventStatsUpdate is the builder for updating OtherEventStats entities.
type OtherEventStatsUpdate struct {
	config
	hooks    []Hook
	mutation *OtherEventStatsMutation
}

// Where appends a list predicates to the OtherEventStatsUpdate builder.
func (oesu *OtherEventStatsUpdate) Where(ps ...predicate.OtherEventStats) *OtherEventStatsUpdate {
	oesu.mutation.Where(ps...)
	return oesu
}

// SetOtherEventStatsGroupID sets the "other_event_stats_group_id" field.
func (oesu *OtherEventStatsUpdate) SetOtherEventStatsGroupID(s string) *OtherEventStatsUpdate {
	oesu.mutation.SetOtherEventStatsGroupID(s)
	return oesu
}

// SetOtherEventStatsNumEvents sets the "other_event_stats_num_events" field.
func (oesu *OtherEventStatsUpdate) SetOtherEventStatsNumEvents(s string) *OtherEventStatsUpdate {
	oesu.mutation.SetOtherEventStatsNumEvents(s)
	return oesu
}

// SetOtherEventStatsNumAffected sets the "other_event_stats_num_affected" field.
func (oesu *OtherEventStatsUpdate) SetOtherEventStatsNumAffected(s string) *OtherEventStatsUpdate {
	oesu.mutation.SetOtherEventStatsNumAffected(s)
	return oesu
}

// SetOtherEventStatsNumAtRisk sets the "other_event_stats_num_at_risk" field.
func (oesu *OtherEventStatsUpdate) SetOtherEventStatsNumAtRisk(s string) *OtherEventStatsUpdate {
	oesu.mutation.SetOtherEventStatsNumAtRisk(s)
	return oesu
}

// SetParentID sets the "parent" edge to the OtherEvent entity by ID.
func (oesu *OtherEventStatsUpdate) SetParentID(id int) *OtherEventStatsUpdate {
	oesu.mutation.SetParentID(id)
	return oesu
}

// SetNillableParentID sets the "parent" edge to the OtherEvent entity by ID if the given value is not nil.
func (oesu *OtherEventStatsUpdate) SetNillableParentID(id *int) *OtherEventStatsUpdate {
	if id != nil {
		oesu = oesu.SetParentID(*id)
	}
	return oesu
}

// SetParent sets the "parent" edge to the OtherEvent entity.
func (oesu *OtherEventStatsUpdate) SetParent(o *OtherEvent) *OtherEventStatsUpdate {
	return oesu.SetParentID(o.ID)
}

// Mutation returns the OtherEventStatsMutation object of the builder.
func (oesu *OtherEventStatsUpdate) Mutation() *OtherEventStatsMutation {
	return oesu.mutation
}

// ClearParent clears the "parent" edge to the OtherEvent entity.
func (oesu *OtherEventStatsUpdate) ClearParent() *OtherEventStatsUpdate {
	oesu.mutation.ClearParent()
	return oesu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oesu *OtherEventStatsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oesu.hooks) == 0 {
		affected, err = oesu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oesu.mutation = mutation
			affected, err = oesu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oesu.hooks) - 1; i >= 0; i-- {
			if oesu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oesu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oesu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oesu *OtherEventStatsUpdate) SaveX(ctx context.Context) int {
	affected, err := oesu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oesu *OtherEventStatsUpdate) Exec(ctx context.Context) error {
	_, err := oesu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oesu *OtherEventStatsUpdate) ExecX(ctx context.Context) {
	if err := oesu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oesu *OtherEventStatsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   othereventstats.Table,
			Columns: othereventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: othereventstats.FieldID,
			},
		},
	}
	if ps := oesu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oesu.mutation.OtherEventStatsGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsGroupID,
		})
	}
	if value, ok := oesu.mutation.OtherEventStatsNumEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumEvents,
		})
	}
	if value, ok := oesu.mutation.OtherEventStatsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAffected,
		})
	}
	if value, ok := oesu.mutation.OtherEventStatsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAtRisk,
		})
	}
	if oesu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oesu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oesu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{othereventstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OtherEventStatsUpdateOne is the builder for updating a single OtherEventStats entity.
type OtherEventStatsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OtherEventStatsMutation
}

// SetOtherEventStatsGroupID sets the "other_event_stats_group_id" field.
func (oesuo *OtherEventStatsUpdateOne) SetOtherEventStatsGroupID(s string) *OtherEventStatsUpdateOne {
	oesuo.mutation.SetOtherEventStatsGroupID(s)
	return oesuo
}

// SetOtherEventStatsNumEvents sets the "other_event_stats_num_events" field.
func (oesuo *OtherEventStatsUpdateOne) SetOtherEventStatsNumEvents(s string) *OtherEventStatsUpdateOne {
	oesuo.mutation.SetOtherEventStatsNumEvents(s)
	return oesuo
}

// SetOtherEventStatsNumAffected sets the "other_event_stats_num_affected" field.
func (oesuo *OtherEventStatsUpdateOne) SetOtherEventStatsNumAffected(s string) *OtherEventStatsUpdateOne {
	oesuo.mutation.SetOtherEventStatsNumAffected(s)
	return oesuo
}

// SetOtherEventStatsNumAtRisk sets the "other_event_stats_num_at_risk" field.
func (oesuo *OtherEventStatsUpdateOne) SetOtherEventStatsNumAtRisk(s string) *OtherEventStatsUpdateOne {
	oesuo.mutation.SetOtherEventStatsNumAtRisk(s)
	return oesuo
}

// SetParentID sets the "parent" edge to the OtherEvent entity by ID.
func (oesuo *OtherEventStatsUpdateOne) SetParentID(id int) *OtherEventStatsUpdateOne {
	oesuo.mutation.SetParentID(id)
	return oesuo
}

// SetNillableParentID sets the "parent" edge to the OtherEvent entity by ID if the given value is not nil.
func (oesuo *OtherEventStatsUpdateOne) SetNillableParentID(id *int) *OtherEventStatsUpdateOne {
	if id != nil {
		oesuo = oesuo.SetParentID(*id)
	}
	return oesuo
}

// SetParent sets the "parent" edge to the OtherEvent entity.
func (oesuo *OtherEventStatsUpdateOne) SetParent(o *OtherEvent) *OtherEventStatsUpdateOne {
	return oesuo.SetParentID(o.ID)
}

// Mutation returns the OtherEventStatsMutation object of the builder.
func (oesuo *OtherEventStatsUpdateOne) Mutation() *OtherEventStatsMutation {
	return oesuo.mutation
}

// ClearParent clears the "parent" edge to the OtherEvent entity.
func (oesuo *OtherEventStatsUpdateOne) ClearParent() *OtherEventStatsUpdateOne {
	oesuo.mutation.ClearParent()
	return oesuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oesuo *OtherEventStatsUpdateOne) Select(field string, fields ...string) *OtherEventStatsUpdateOne {
	oesuo.fields = append([]string{field}, fields...)
	return oesuo
}

// Save executes the query and returns the updated OtherEventStats entity.
func (oesuo *OtherEventStatsUpdateOne) Save(ctx context.Context) (*OtherEventStats, error) {
	var (
		err  error
		node *OtherEventStats
	)
	if len(oesuo.hooks) == 0 {
		node, err = oesuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventStatsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oesuo.mutation = mutation
			node, err = oesuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oesuo.hooks) - 1; i >= 0; i-- {
			if oesuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oesuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oesuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oesuo *OtherEventStatsUpdateOne) SaveX(ctx context.Context) *OtherEventStats {
	node, err := oesuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oesuo *OtherEventStatsUpdateOne) Exec(ctx context.Context) error {
	_, err := oesuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oesuo *OtherEventStatsUpdateOne) ExecX(ctx context.Context) {
	if err := oesuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oesuo *OtherEventStatsUpdateOne) sqlSave(ctx context.Context) (_node *OtherEventStats, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   othereventstats.Table,
			Columns: othereventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: othereventstats.FieldID,
			},
		},
	}
	id, ok := oesuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OtherEventStats.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oesuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, othereventstats.FieldID)
		for _, f := range fields {
			if !othereventstats.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != othereventstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oesuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oesuo.mutation.OtherEventStatsGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsGroupID,
		})
	}
	if value, ok := oesuo.mutation.OtherEventStatsNumEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumEvents,
		})
	}
	if value, ok := oesuo.mutation.OtherEventStatsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAffected,
		})
	}
	if value, ok := oesuo.mutation.OtherEventStatsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: othereventstats.FieldOtherEventStatsNumAtRisk,
		})
	}
	if oesuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oesuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OtherEventStats{config: oesuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oesuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{othereventstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
