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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventUpdate is the builder for updating SeriousEvent entities.
type SeriousEventUpdate struct {
	config
	hooks    []Hook
	mutation *SeriousEventMutation
}

// Where appends a list predicates to the SeriousEventUpdate builder.
func (seu *SeriousEventUpdate) Where(ps ...predicate.SeriousEvent) *SeriousEventUpdate {
	seu.mutation.Where(ps...)
	return seu
}

// SetSeriousEventTerm sets the "serious_event_term" field.
func (seu *SeriousEventUpdate) SetSeriousEventTerm(s string) *SeriousEventUpdate {
	seu.mutation.SetSeriousEventTerm(s)
	return seu
}

// SetSeriousEventOrganSystem sets the "serious_event_organ_system" field.
func (seu *SeriousEventUpdate) SetSeriousEventOrganSystem(s string) *SeriousEventUpdate {
	seu.mutation.SetSeriousEventOrganSystem(s)
	return seu
}

// SetSeriousEventSourceVocabulary sets the "serious_event_source_vocabulary" field.
func (seu *SeriousEventUpdate) SetSeriousEventSourceVocabulary(s string) *SeriousEventUpdate {
	seu.mutation.SetSeriousEventSourceVocabulary(s)
	return seu
}

// SetSeriousEventAssessmentType sets the "serious_event_assessment_type" field.
func (seu *SeriousEventUpdate) SetSeriousEventAssessmentType(s string) *SeriousEventUpdate {
	seu.mutation.SetSeriousEventAssessmentType(s)
	return seu
}

// SetSeriousEventNotes sets the "serious_event_notes" field.
func (seu *SeriousEventUpdate) SetSeriousEventNotes(s string) *SeriousEventUpdate {
	seu.mutation.SetSeriousEventNotes(s)
	return seu
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (seu *SeriousEventUpdate) SetParentID(id int) *SeriousEventUpdate {
	seu.mutation.SetParentID(id)
	return seu
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (seu *SeriousEventUpdate) SetNillableParentID(id *int) *SeriousEventUpdate {
	if id != nil {
		seu = seu.SetParentID(*id)
	}
	return seu
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (seu *SeriousEventUpdate) SetParent(a *AdverseEventsModule) *SeriousEventUpdate {
	return seu.SetParentID(a.ID)
}

// AddSeriousEventStatsListIDs adds the "serious_event_stats_list" edge to the SeriousEventStats entity by IDs.
func (seu *SeriousEventUpdate) AddSeriousEventStatsListIDs(ids ...int) *SeriousEventUpdate {
	seu.mutation.AddSeriousEventStatsListIDs(ids...)
	return seu
}

// AddSeriousEventStatsList adds the "serious_event_stats_list" edges to the SeriousEventStats entity.
func (seu *SeriousEventUpdate) AddSeriousEventStatsList(s ...*SeriousEventStats) *SeriousEventUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return seu.AddSeriousEventStatsListIDs(ids...)
}

// Mutation returns the SeriousEventMutation object of the builder.
func (seu *SeriousEventUpdate) Mutation() *SeriousEventMutation {
	return seu.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (seu *SeriousEventUpdate) ClearParent() *SeriousEventUpdate {
	seu.mutation.ClearParent()
	return seu
}

// ClearSeriousEventStatsList clears all "serious_event_stats_list" edges to the SeriousEventStats entity.
func (seu *SeriousEventUpdate) ClearSeriousEventStatsList() *SeriousEventUpdate {
	seu.mutation.ClearSeriousEventStatsList()
	return seu
}

// RemoveSeriousEventStatsListIDs removes the "serious_event_stats_list" edge to SeriousEventStats entities by IDs.
func (seu *SeriousEventUpdate) RemoveSeriousEventStatsListIDs(ids ...int) *SeriousEventUpdate {
	seu.mutation.RemoveSeriousEventStatsListIDs(ids...)
	return seu
}

// RemoveSeriousEventStatsList removes "serious_event_stats_list" edges to SeriousEventStats entities.
func (seu *SeriousEventUpdate) RemoveSeriousEventStatsList(s ...*SeriousEventStats) *SeriousEventUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return seu.RemoveSeriousEventStatsListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seu *SeriousEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(seu.hooks) == 0 {
		affected, err = seu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seu.mutation = mutation
			affected, err = seu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(seu.hooks) - 1; i >= 0; i-- {
			if seu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = seu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (seu *SeriousEventUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *SeriousEventUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *SeriousEventUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seu *SeriousEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriousevent.Table,
			Columns: seriousevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriousevent.FieldID,
			},
		},
	}
	if ps := seu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seu.mutation.SeriousEventTerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventTerm,
		})
	}
	if value, ok := seu.mutation.SeriousEventOrganSystem(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventOrganSystem,
		})
	}
	if value, ok := seu.mutation.SeriousEventSourceVocabulary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventSourceVocabulary,
		})
	}
	if value, ok := seu.mutation.SeriousEventAssessmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventAssessmentType,
		})
	}
	if value, ok := seu.mutation.SeriousEventNotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventNotes,
		})
	}
	if seu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seu.mutation.SeriousEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.RemovedSeriousEventStatsListIDs(); len(nodes) > 0 && !seu.mutation.SeriousEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.SeriousEventStatsListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seriousevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SeriousEventUpdateOne is the builder for updating a single SeriousEvent entity.
type SeriousEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeriousEventMutation
}

// SetSeriousEventTerm sets the "serious_event_term" field.
func (seuo *SeriousEventUpdateOne) SetSeriousEventTerm(s string) *SeriousEventUpdateOne {
	seuo.mutation.SetSeriousEventTerm(s)
	return seuo
}

// SetSeriousEventOrganSystem sets the "serious_event_organ_system" field.
func (seuo *SeriousEventUpdateOne) SetSeriousEventOrganSystem(s string) *SeriousEventUpdateOne {
	seuo.mutation.SetSeriousEventOrganSystem(s)
	return seuo
}

// SetSeriousEventSourceVocabulary sets the "serious_event_source_vocabulary" field.
func (seuo *SeriousEventUpdateOne) SetSeriousEventSourceVocabulary(s string) *SeriousEventUpdateOne {
	seuo.mutation.SetSeriousEventSourceVocabulary(s)
	return seuo
}

// SetSeriousEventAssessmentType sets the "serious_event_assessment_type" field.
func (seuo *SeriousEventUpdateOne) SetSeriousEventAssessmentType(s string) *SeriousEventUpdateOne {
	seuo.mutation.SetSeriousEventAssessmentType(s)
	return seuo
}

// SetSeriousEventNotes sets the "serious_event_notes" field.
func (seuo *SeriousEventUpdateOne) SetSeriousEventNotes(s string) *SeriousEventUpdateOne {
	seuo.mutation.SetSeriousEventNotes(s)
	return seuo
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (seuo *SeriousEventUpdateOne) SetParentID(id int) *SeriousEventUpdateOne {
	seuo.mutation.SetParentID(id)
	return seuo
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (seuo *SeriousEventUpdateOne) SetNillableParentID(id *int) *SeriousEventUpdateOne {
	if id != nil {
		seuo = seuo.SetParentID(*id)
	}
	return seuo
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (seuo *SeriousEventUpdateOne) SetParent(a *AdverseEventsModule) *SeriousEventUpdateOne {
	return seuo.SetParentID(a.ID)
}

// AddSeriousEventStatsListIDs adds the "serious_event_stats_list" edge to the SeriousEventStats entity by IDs.
func (seuo *SeriousEventUpdateOne) AddSeriousEventStatsListIDs(ids ...int) *SeriousEventUpdateOne {
	seuo.mutation.AddSeriousEventStatsListIDs(ids...)
	return seuo
}

// AddSeriousEventStatsList adds the "serious_event_stats_list" edges to the SeriousEventStats entity.
func (seuo *SeriousEventUpdateOne) AddSeriousEventStatsList(s ...*SeriousEventStats) *SeriousEventUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return seuo.AddSeriousEventStatsListIDs(ids...)
}

// Mutation returns the SeriousEventMutation object of the builder.
func (seuo *SeriousEventUpdateOne) Mutation() *SeriousEventMutation {
	return seuo.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (seuo *SeriousEventUpdateOne) ClearParent() *SeriousEventUpdateOne {
	seuo.mutation.ClearParent()
	return seuo
}

// ClearSeriousEventStatsList clears all "serious_event_stats_list" edges to the SeriousEventStats entity.
func (seuo *SeriousEventUpdateOne) ClearSeriousEventStatsList() *SeriousEventUpdateOne {
	seuo.mutation.ClearSeriousEventStatsList()
	return seuo
}

// RemoveSeriousEventStatsListIDs removes the "serious_event_stats_list" edge to SeriousEventStats entities by IDs.
func (seuo *SeriousEventUpdateOne) RemoveSeriousEventStatsListIDs(ids ...int) *SeriousEventUpdateOne {
	seuo.mutation.RemoveSeriousEventStatsListIDs(ids...)
	return seuo
}

// RemoveSeriousEventStatsList removes "serious_event_stats_list" edges to SeriousEventStats entities.
func (seuo *SeriousEventUpdateOne) RemoveSeriousEventStatsList(s ...*SeriousEventStats) *SeriousEventUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return seuo.RemoveSeriousEventStatsListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seuo *SeriousEventUpdateOne) Select(field string, fields ...string) *SeriousEventUpdateOne {
	seuo.fields = append([]string{field}, fields...)
	return seuo
}

// Save executes the query and returns the updated SeriousEvent entity.
func (seuo *SeriousEventUpdateOne) Save(ctx context.Context) (*SeriousEvent, error) {
	var (
		err  error
		node *SeriousEvent
	)
	if len(seuo.hooks) == 0 {
		node, err = seuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriousEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seuo.mutation = mutation
			node, err = seuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(seuo.hooks) - 1; i >= 0; i-- {
			if seuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = seuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *SeriousEventUpdateOne) SaveX(ctx context.Context) *SeriousEvent {
	node, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seuo *SeriousEventUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *SeriousEventUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seuo *SeriousEventUpdateOne) sqlSave(ctx context.Context) (_node *SeriousEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriousevent.Table,
			Columns: seriousevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriousevent.FieldID,
			},
		},
	}
	id, ok := seuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "SeriousEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seriousevent.FieldID)
		for _, f := range fields {
			if !seriousevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != seriousevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seuo.mutation.SeriousEventTerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventTerm,
		})
	}
	if value, ok := seuo.mutation.SeriousEventOrganSystem(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventOrganSystem,
		})
	}
	if value, ok := seuo.mutation.SeriousEventSourceVocabulary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventSourceVocabulary,
		})
	}
	if value, ok := seuo.mutation.SeriousEventAssessmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventAssessmentType,
		})
	}
	if value, ok := seuo.mutation.SeriousEventNotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seriousevent.FieldSeriousEventNotes,
		})
	}
	if seuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seuo.mutation.SeriousEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.RemovedSeriousEventStatsListIDs(); len(nodes) > 0 && !seuo.mutation.SeriousEventStatsListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.SeriousEventStatsListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SeriousEvent{config: seuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seriousevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
