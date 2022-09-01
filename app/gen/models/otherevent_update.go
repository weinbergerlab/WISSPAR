// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventUpdate is the builder for updating OtherEvent entities.
type OtherEventUpdate struct {
	config
	hooks    []Hook
	mutation *OtherEventMutation
}

// Where appends a list predicates to the OtherEventUpdate builder.
func (oeu *OtherEventUpdate) Where(ps ...predicate.OtherEvent) *OtherEventUpdate {
	oeu.mutation.Where(ps...)
	return oeu
}

// SetOtherEventTerm sets the "other_event_term" field.
func (oeu *OtherEventUpdate) SetOtherEventTerm(s string) *OtherEventUpdate {
	oeu.mutation.SetOtherEventTerm(s)
	return oeu
}

// SetOtherEventOrganSystem sets the "other_event_organ_system" field.
func (oeu *OtherEventUpdate) SetOtherEventOrganSystem(s string) *OtherEventUpdate {
	oeu.mutation.SetOtherEventOrganSystem(s)
	return oeu
}

// SetOtherEventSourceVocabulary sets the "other_event_source_vocabulary" field.
func (oeu *OtherEventUpdate) SetOtherEventSourceVocabulary(s string) *OtherEventUpdate {
	oeu.mutation.SetOtherEventSourceVocabulary(s)
	return oeu
}

// SetOtherEventAssessmentType sets the "other_event_assessment_type" field.
func (oeu *OtherEventUpdate) SetOtherEventAssessmentType(s string) *OtherEventUpdate {
	oeu.mutation.SetOtherEventAssessmentType(s)
	return oeu
}

// SetOtherEventNotes sets the "other_event_notes" field.
func (oeu *OtherEventUpdate) SetOtherEventNotes(s string) *OtherEventUpdate {
	oeu.mutation.SetOtherEventNotes(s)
	return oeu
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (oeu *OtherEventUpdate) SetParentID(id int) *OtherEventUpdate {
	oeu.mutation.SetParentID(id)
	return oeu
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (oeu *OtherEventUpdate) SetNillableParentID(id *int) *OtherEventUpdate {
	if id != nil {
		oeu = oeu.SetParentID(*id)
	}
	return oeu
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (oeu *OtherEventUpdate) SetParent(a *AdverseEventsModule) *OtherEventUpdate {
	return oeu.SetParentID(a.ID)
}

// AddOtherEventStatsListIDs adds the "other_event_stats_list" edge to the OtherEventStats entity by IDs.
func (oeu *OtherEventUpdate) AddOtherEventStatsListIDs(ids ...int) *OtherEventUpdate {
	oeu.mutation.AddOtherEventStatsListIDs(ids...)
	return oeu
}

// AddOtherEventStatsList adds the "other_event_stats_list" edges to the OtherEventStats entity.
func (oeu *OtherEventUpdate) AddOtherEventStatsList(o ...*OtherEventStats) *OtherEventUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oeu.AddOtherEventStatsListIDs(ids...)
}

// Mutation returns the OtherEventMutation object of the builder.
func (oeu *OtherEventUpdate) Mutation() *OtherEventMutation {
	return oeu.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (oeu *OtherEventUpdate) ClearParent() *OtherEventUpdate {
	oeu.mutation.ClearParent()
	return oeu
}

// ClearOtherEventStatsList clears all "other_event_stats_list" edges to the OtherEventStats entity.
func (oeu *OtherEventUpdate) ClearOtherEventStatsList() *OtherEventUpdate {
	oeu.mutation.ClearOtherEventStatsList()
	return oeu
}

// RemoveOtherEventStatsListIDs removes the "other_event_stats_list" edge to OtherEventStats entities by IDs.
func (oeu *OtherEventUpdate) RemoveOtherEventStatsListIDs(ids ...int) *OtherEventUpdate {
	oeu.mutation.RemoveOtherEventStatsListIDs(ids...)
	return oeu
}

// RemoveOtherEventStatsList removes "other_event_stats_list" edges to OtherEventStats entities.
func (oeu *OtherEventUpdate) RemoveOtherEventStatsList(o ...*OtherEventStats) *OtherEventUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oeu.RemoveOtherEventStatsListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oeu *OtherEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oeu.hooks) == 0 {
		affected, err = oeu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oeu.mutation = mutation
			affected, err = oeu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oeu.hooks) - 1; i >= 0; i-- {
			if oeu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oeu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oeu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oeu *OtherEventUpdate) SaveX(ctx context.Context) int {
	affected, err := oeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oeu *OtherEventUpdate) Exec(ctx context.Context) error {
	_, err := oeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeu *OtherEventUpdate) ExecX(ctx context.Context) {
	if err := oeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oeu *OtherEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   otherevent.Table,
			Columns: otherevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: otherevent.FieldID,
			},
		},
	}
	if ps := oeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeu.mutation.OtherEventTerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventTerm,
		})
	}
	if value, ok := oeu.mutation.OtherEventOrganSystem(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventOrganSystem,
		})
	}
	if value, ok := oeu.mutation.OtherEventSourceVocabulary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventSourceVocabulary,
		})
	}
	if value, ok := oeu.mutation.OtherEventAssessmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventAssessmentType,
		})
	}
	if value, ok := oeu.mutation.OtherEventNotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventNotes,
		})
	}
	if oeu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oeu.mutation.OtherEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeu.mutation.RemovedOtherEventStatsListIDs(); len(nodes) > 0 && !oeu.mutation.OtherEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeu.mutation.OtherEventStatsListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{otherevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OtherEventUpdateOne is the builder for updating a single OtherEvent entity.
type OtherEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OtherEventMutation
}

// SetOtherEventTerm sets the "other_event_term" field.
func (oeuo *OtherEventUpdateOne) SetOtherEventTerm(s string) *OtherEventUpdateOne {
	oeuo.mutation.SetOtherEventTerm(s)
	return oeuo
}

// SetOtherEventOrganSystem sets the "other_event_organ_system" field.
func (oeuo *OtherEventUpdateOne) SetOtherEventOrganSystem(s string) *OtherEventUpdateOne {
	oeuo.mutation.SetOtherEventOrganSystem(s)
	return oeuo
}

// SetOtherEventSourceVocabulary sets the "other_event_source_vocabulary" field.
func (oeuo *OtherEventUpdateOne) SetOtherEventSourceVocabulary(s string) *OtherEventUpdateOne {
	oeuo.mutation.SetOtherEventSourceVocabulary(s)
	return oeuo
}

// SetOtherEventAssessmentType sets the "other_event_assessment_type" field.
func (oeuo *OtherEventUpdateOne) SetOtherEventAssessmentType(s string) *OtherEventUpdateOne {
	oeuo.mutation.SetOtherEventAssessmentType(s)
	return oeuo
}

// SetOtherEventNotes sets the "other_event_notes" field.
func (oeuo *OtherEventUpdateOne) SetOtherEventNotes(s string) *OtherEventUpdateOne {
	oeuo.mutation.SetOtherEventNotes(s)
	return oeuo
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (oeuo *OtherEventUpdateOne) SetParentID(id int) *OtherEventUpdateOne {
	oeuo.mutation.SetParentID(id)
	return oeuo
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (oeuo *OtherEventUpdateOne) SetNillableParentID(id *int) *OtherEventUpdateOne {
	if id != nil {
		oeuo = oeuo.SetParentID(*id)
	}
	return oeuo
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (oeuo *OtherEventUpdateOne) SetParent(a *AdverseEventsModule) *OtherEventUpdateOne {
	return oeuo.SetParentID(a.ID)
}

// AddOtherEventStatsListIDs adds the "other_event_stats_list" edge to the OtherEventStats entity by IDs.
func (oeuo *OtherEventUpdateOne) AddOtherEventStatsListIDs(ids ...int) *OtherEventUpdateOne {
	oeuo.mutation.AddOtherEventStatsListIDs(ids...)
	return oeuo
}

// AddOtherEventStatsList adds the "other_event_stats_list" edges to the OtherEventStats entity.
func (oeuo *OtherEventUpdateOne) AddOtherEventStatsList(o ...*OtherEventStats) *OtherEventUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oeuo.AddOtherEventStatsListIDs(ids...)
}

// Mutation returns the OtherEventMutation object of the builder.
func (oeuo *OtherEventUpdateOne) Mutation() *OtherEventMutation {
	return oeuo.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (oeuo *OtherEventUpdateOne) ClearParent() *OtherEventUpdateOne {
	oeuo.mutation.ClearParent()
	return oeuo
}

// ClearOtherEventStatsList clears all "other_event_stats_list" edges to the OtherEventStats entity.
func (oeuo *OtherEventUpdateOne) ClearOtherEventStatsList() *OtherEventUpdateOne {
	oeuo.mutation.ClearOtherEventStatsList()
	return oeuo
}

// RemoveOtherEventStatsListIDs removes the "other_event_stats_list" edge to OtherEventStats entities by IDs.
func (oeuo *OtherEventUpdateOne) RemoveOtherEventStatsListIDs(ids ...int) *OtherEventUpdateOne {
	oeuo.mutation.RemoveOtherEventStatsListIDs(ids...)
	return oeuo
}

// RemoveOtherEventStatsList removes "other_event_stats_list" edges to OtherEventStats entities.
func (oeuo *OtherEventUpdateOne) RemoveOtherEventStatsList(o ...*OtherEventStats) *OtherEventUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oeuo.RemoveOtherEventStatsListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oeuo *OtherEventUpdateOne) Select(field string, fields ...string) *OtherEventUpdateOne {
	oeuo.fields = append([]string{field}, fields...)
	return oeuo
}

// Save executes the query and returns the updated OtherEvent entity.
func (oeuo *OtherEventUpdateOne) Save(ctx context.Context) (*OtherEvent, error) {
	var (
		err  error
		node *OtherEvent
	)
	if len(oeuo.hooks) == 0 {
		node, err = oeuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OtherEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oeuo.mutation = mutation
			node, err = oeuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oeuo.hooks) - 1; i >= 0; i-- {
			if oeuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oeuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oeuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oeuo *OtherEventUpdateOne) SaveX(ctx context.Context) *OtherEvent {
	node, err := oeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oeuo *OtherEventUpdateOne) Exec(ctx context.Context) error {
	_, err := oeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeuo *OtherEventUpdateOne) ExecX(ctx context.Context) {
	if err := oeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oeuo *OtherEventUpdateOne) sqlSave(ctx context.Context) (_node *OtherEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   otherevent.Table,
			Columns: otherevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: otherevent.FieldID,
			},
		},
	}
	id, ok := oeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OtherEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, otherevent.FieldID)
		for _, f := range fields {
			if !otherevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != otherevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeuo.mutation.OtherEventTerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventTerm,
		})
	}
	if value, ok := oeuo.mutation.OtherEventOrganSystem(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventOrganSystem,
		})
	}
	if value, ok := oeuo.mutation.OtherEventSourceVocabulary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventSourceVocabulary,
		})
	}
	if value, ok := oeuo.mutation.OtherEventAssessmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventAssessmentType,
		})
	}
	if value, ok := oeuo.mutation.OtherEventNotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: otherevent.FieldOtherEventNotes,
		})
	}
	if oeuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oeuo.mutation.OtherEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeuo.mutation.RemovedOtherEventStatsListIDs(); len(nodes) > 0 && !oeuo.mutation.OtherEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeuo.mutation.OtherEventStatsListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OtherEvent{config: oeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{otherevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
