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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
)

// AdverseEventsModuleUpdate is the builder for updating AdverseEventsModule entities.
type AdverseEventsModuleUpdate struct {
	config
	hooks    []Hook
	mutation *AdverseEventsModuleMutation
}

// Where appends a list predicates to the AdverseEventsModuleUpdate builder.
func (aemu *AdverseEventsModuleUpdate) Where(ps ...predicate.AdverseEventsModule) *AdverseEventsModuleUpdate {
	aemu.mutation.Where(ps...)
	return aemu
}

// SetEventsFrequencyThreshold sets the "events_frequency_threshold" field.
func (aemu *AdverseEventsModuleUpdate) SetEventsFrequencyThreshold(s string) *AdverseEventsModuleUpdate {
	aemu.mutation.SetEventsFrequencyThreshold(s)
	return aemu
}

// SetEventsTimeFrame sets the "events_time_frame" field.
func (aemu *AdverseEventsModuleUpdate) SetEventsTimeFrame(s string) *AdverseEventsModuleUpdate {
	aemu.mutation.SetEventsTimeFrame(s)
	return aemu
}

// SetEventsDescription sets the "events_description" field.
func (aemu *AdverseEventsModuleUpdate) SetEventsDescription(s string) *AdverseEventsModuleUpdate {
	aemu.mutation.SetEventsDescription(s)
	return aemu
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (aemu *AdverseEventsModuleUpdate) SetParentID(id int) *AdverseEventsModuleUpdate {
	aemu.mutation.SetParentID(id)
	return aemu
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (aemu *AdverseEventsModuleUpdate) SetNillableParentID(id *int) *AdverseEventsModuleUpdate {
	if id != nil {
		aemu = aemu.SetParentID(*id)
	}
	return aemu
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (aemu *AdverseEventsModuleUpdate) SetParent(r *ResultsDefinition) *AdverseEventsModuleUpdate {
	return aemu.SetParentID(r.ID)
}

// AddEventGroupListIDs adds the "event_group_list" edge to the EventGroup entity by IDs.
func (aemu *AdverseEventsModuleUpdate) AddEventGroupListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.AddEventGroupListIDs(ids...)
	return aemu
}

// AddEventGroupList adds the "event_group_list" edges to the EventGroup entity.
func (aemu *AdverseEventsModuleUpdate) AddEventGroupList(e ...*EventGroup) *AdverseEventsModuleUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return aemu.AddEventGroupListIDs(ids...)
}

// AddSeriousEventListIDs adds the "serious_event_list" edge to the SeriousEvent entity by IDs.
func (aemu *AdverseEventsModuleUpdate) AddSeriousEventListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.AddSeriousEventListIDs(ids...)
	return aemu
}

// AddSeriousEventList adds the "serious_event_list" edges to the SeriousEvent entity.
func (aemu *AdverseEventsModuleUpdate) AddSeriousEventList(s ...*SeriousEvent) *AdverseEventsModuleUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aemu.AddSeriousEventListIDs(ids...)
}

// AddOtherEventListIDs adds the "other_event_list" edge to the OtherEvent entity by IDs.
func (aemu *AdverseEventsModuleUpdate) AddOtherEventListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.AddOtherEventListIDs(ids...)
	return aemu
}

// AddOtherEventList adds the "other_event_list" edges to the OtherEvent entity.
func (aemu *AdverseEventsModuleUpdate) AddOtherEventList(o ...*OtherEvent) *AdverseEventsModuleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return aemu.AddOtherEventListIDs(ids...)
}

// Mutation returns the AdverseEventsModuleMutation object of the builder.
func (aemu *AdverseEventsModuleUpdate) Mutation() *AdverseEventsModuleMutation {
	return aemu.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (aemu *AdverseEventsModuleUpdate) ClearParent() *AdverseEventsModuleUpdate {
	aemu.mutation.ClearParent()
	return aemu
}

// ClearEventGroupList clears all "event_group_list" edges to the EventGroup entity.
func (aemu *AdverseEventsModuleUpdate) ClearEventGroupList() *AdverseEventsModuleUpdate {
	aemu.mutation.ClearEventGroupList()
	return aemu
}

// RemoveEventGroupListIDs removes the "event_group_list" edge to EventGroup entities by IDs.
func (aemu *AdverseEventsModuleUpdate) RemoveEventGroupListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.RemoveEventGroupListIDs(ids...)
	return aemu
}

// RemoveEventGroupList removes "event_group_list" edges to EventGroup entities.
func (aemu *AdverseEventsModuleUpdate) RemoveEventGroupList(e ...*EventGroup) *AdverseEventsModuleUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return aemu.RemoveEventGroupListIDs(ids...)
}

// ClearSeriousEventList clears all "serious_event_list" edges to the SeriousEvent entity.
func (aemu *AdverseEventsModuleUpdate) ClearSeriousEventList() *AdverseEventsModuleUpdate {
	aemu.mutation.ClearSeriousEventList()
	return aemu
}

// RemoveSeriousEventListIDs removes the "serious_event_list" edge to SeriousEvent entities by IDs.
func (aemu *AdverseEventsModuleUpdate) RemoveSeriousEventListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.RemoveSeriousEventListIDs(ids...)
	return aemu
}

// RemoveSeriousEventList removes "serious_event_list" edges to SeriousEvent entities.
func (aemu *AdverseEventsModuleUpdate) RemoveSeriousEventList(s ...*SeriousEvent) *AdverseEventsModuleUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aemu.RemoveSeriousEventListIDs(ids...)
}

// ClearOtherEventList clears all "other_event_list" edges to the OtherEvent entity.
func (aemu *AdverseEventsModuleUpdate) ClearOtherEventList() *AdverseEventsModuleUpdate {
	aemu.mutation.ClearOtherEventList()
	return aemu
}

// RemoveOtherEventListIDs removes the "other_event_list" edge to OtherEvent entities by IDs.
func (aemu *AdverseEventsModuleUpdate) RemoveOtherEventListIDs(ids ...int) *AdverseEventsModuleUpdate {
	aemu.mutation.RemoveOtherEventListIDs(ids...)
	return aemu
}

// RemoveOtherEventList removes "other_event_list" edges to OtherEvent entities.
func (aemu *AdverseEventsModuleUpdate) RemoveOtherEventList(o ...*OtherEvent) *AdverseEventsModuleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return aemu.RemoveOtherEventListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aemu *AdverseEventsModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aemu.hooks) == 0 {
		affected, err = aemu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdverseEventsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aemu.mutation = mutation
			affected, err = aemu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aemu.hooks) - 1; i >= 0; i-- {
			if aemu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = aemu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aemu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aemu *AdverseEventsModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := aemu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aemu *AdverseEventsModuleUpdate) Exec(ctx context.Context) error {
	_, err := aemu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aemu *AdverseEventsModuleUpdate) ExecX(ctx context.Context) {
	if err := aemu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aemu *AdverseEventsModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adverseeventsmodule.Table,
			Columns: adverseeventsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adverseeventsmodule.FieldID,
			},
		},
	}
	if ps := aemu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aemu.mutation.EventsFrequencyThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsFrequencyThreshold,
		})
	}
	if value, ok := aemu.mutation.EventsTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsTimeFrame,
		})
	}
	if value, ok := aemu.mutation.EventsDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsDescription,
		})
	}
	if aemu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemu.mutation.EventGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.RemovedEventGroupListIDs(); len(nodes) > 0 && !aemu.mutation.EventGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.EventGroupListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemu.mutation.SeriousEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.RemovedSeriousEventListIDs(); len(nodes) > 0 && !aemu.mutation.SeriousEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.SeriousEventListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemu.mutation.OtherEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.RemovedOtherEventListIDs(); len(nodes) > 0 && !aemu.mutation.OtherEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemu.mutation.OtherEventListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aemu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adverseeventsmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdverseEventsModuleUpdateOne is the builder for updating a single AdverseEventsModule entity.
type AdverseEventsModuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdverseEventsModuleMutation
}

// SetEventsFrequencyThreshold sets the "events_frequency_threshold" field.
func (aemuo *AdverseEventsModuleUpdateOne) SetEventsFrequencyThreshold(s string) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.SetEventsFrequencyThreshold(s)
	return aemuo
}

// SetEventsTimeFrame sets the "events_time_frame" field.
func (aemuo *AdverseEventsModuleUpdateOne) SetEventsTimeFrame(s string) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.SetEventsTimeFrame(s)
	return aemuo
}

// SetEventsDescription sets the "events_description" field.
func (aemuo *AdverseEventsModuleUpdateOne) SetEventsDescription(s string) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.SetEventsDescription(s)
	return aemuo
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (aemuo *AdverseEventsModuleUpdateOne) SetParentID(id int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.SetParentID(id)
	return aemuo
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (aemuo *AdverseEventsModuleUpdateOne) SetNillableParentID(id *int) *AdverseEventsModuleUpdateOne {
	if id != nil {
		aemuo = aemuo.SetParentID(*id)
	}
	return aemuo
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (aemuo *AdverseEventsModuleUpdateOne) SetParent(r *ResultsDefinition) *AdverseEventsModuleUpdateOne {
	return aemuo.SetParentID(r.ID)
}

// AddEventGroupListIDs adds the "event_group_list" edge to the EventGroup entity by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) AddEventGroupListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.AddEventGroupListIDs(ids...)
	return aemuo
}

// AddEventGroupList adds the "event_group_list" edges to the EventGroup entity.
func (aemuo *AdverseEventsModuleUpdateOne) AddEventGroupList(e ...*EventGroup) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return aemuo.AddEventGroupListIDs(ids...)
}

// AddSeriousEventListIDs adds the "serious_event_list" edge to the SeriousEvent entity by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) AddSeriousEventListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.AddSeriousEventListIDs(ids...)
	return aemuo
}

// AddSeriousEventList adds the "serious_event_list" edges to the SeriousEvent entity.
func (aemuo *AdverseEventsModuleUpdateOne) AddSeriousEventList(s ...*SeriousEvent) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aemuo.AddSeriousEventListIDs(ids...)
}

// AddOtherEventListIDs adds the "other_event_list" edge to the OtherEvent entity by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) AddOtherEventListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.AddOtherEventListIDs(ids...)
	return aemuo
}

// AddOtherEventList adds the "other_event_list" edges to the OtherEvent entity.
func (aemuo *AdverseEventsModuleUpdateOne) AddOtherEventList(o ...*OtherEvent) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return aemuo.AddOtherEventListIDs(ids...)
}

// Mutation returns the AdverseEventsModuleMutation object of the builder.
func (aemuo *AdverseEventsModuleUpdateOne) Mutation() *AdverseEventsModuleMutation {
	return aemuo.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (aemuo *AdverseEventsModuleUpdateOne) ClearParent() *AdverseEventsModuleUpdateOne {
	aemuo.mutation.ClearParent()
	return aemuo
}

// ClearEventGroupList clears all "event_group_list" edges to the EventGroup entity.
func (aemuo *AdverseEventsModuleUpdateOne) ClearEventGroupList() *AdverseEventsModuleUpdateOne {
	aemuo.mutation.ClearEventGroupList()
	return aemuo
}

// RemoveEventGroupListIDs removes the "event_group_list" edge to EventGroup entities by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveEventGroupListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.RemoveEventGroupListIDs(ids...)
	return aemuo
}

// RemoveEventGroupList removes "event_group_list" edges to EventGroup entities.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveEventGroupList(e ...*EventGroup) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return aemuo.RemoveEventGroupListIDs(ids...)
}

// ClearSeriousEventList clears all "serious_event_list" edges to the SeriousEvent entity.
func (aemuo *AdverseEventsModuleUpdateOne) ClearSeriousEventList() *AdverseEventsModuleUpdateOne {
	aemuo.mutation.ClearSeriousEventList()
	return aemuo
}

// RemoveSeriousEventListIDs removes the "serious_event_list" edge to SeriousEvent entities by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveSeriousEventListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.RemoveSeriousEventListIDs(ids...)
	return aemuo
}

// RemoveSeriousEventList removes "serious_event_list" edges to SeriousEvent entities.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveSeriousEventList(s ...*SeriousEvent) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aemuo.RemoveSeriousEventListIDs(ids...)
}

// ClearOtherEventList clears all "other_event_list" edges to the OtherEvent entity.
func (aemuo *AdverseEventsModuleUpdateOne) ClearOtherEventList() *AdverseEventsModuleUpdateOne {
	aemuo.mutation.ClearOtherEventList()
	return aemuo
}

// RemoveOtherEventListIDs removes the "other_event_list" edge to OtherEvent entities by IDs.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveOtherEventListIDs(ids ...int) *AdverseEventsModuleUpdateOne {
	aemuo.mutation.RemoveOtherEventListIDs(ids...)
	return aemuo
}

// RemoveOtherEventList removes "other_event_list" edges to OtherEvent entities.
func (aemuo *AdverseEventsModuleUpdateOne) RemoveOtherEventList(o ...*OtherEvent) *AdverseEventsModuleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return aemuo.RemoveOtherEventListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aemuo *AdverseEventsModuleUpdateOne) Select(field string, fields ...string) *AdverseEventsModuleUpdateOne {
	aemuo.fields = append([]string{field}, fields...)
	return aemuo
}

// Save executes the query and returns the updated AdverseEventsModule entity.
func (aemuo *AdverseEventsModuleUpdateOne) Save(ctx context.Context) (*AdverseEventsModule, error) {
	var (
		err  error
		node *AdverseEventsModule
	)
	if len(aemuo.hooks) == 0 {
		node, err = aemuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdverseEventsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aemuo.mutation = mutation
			node, err = aemuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aemuo.hooks) - 1; i >= 0; i-- {
			if aemuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = aemuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aemuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aemuo *AdverseEventsModuleUpdateOne) SaveX(ctx context.Context) *AdverseEventsModule {
	node, err := aemuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aemuo *AdverseEventsModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := aemuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aemuo *AdverseEventsModuleUpdateOne) ExecX(ctx context.Context) {
	if err := aemuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aemuo *AdverseEventsModuleUpdateOne) sqlSave(ctx context.Context) (_node *AdverseEventsModule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adverseeventsmodule.Table,
			Columns: adverseeventsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adverseeventsmodule.FieldID,
			},
		},
	}
	id, ok := aemuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "AdverseEventsModule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aemuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adverseeventsmodule.FieldID)
		for _, f := range fields {
			if !adverseeventsmodule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != adverseeventsmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aemuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aemuo.mutation.EventsFrequencyThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsFrequencyThreshold,
		})
	}
	if value, ok := aemuo.mutation.EventsTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsTimeFrame,
		})
	}
	if value, ok := aemuo.mutation.EventsDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adverseeventsmodule.FieldEventsDescription,
		})
	}
	if aemuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemuo.mutation.EventGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.RemovedEventGroupListIDs(); len(nodes) > 0 && !aemuo.mutation.EventGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.EventGroupListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemuo.mutation.SeriousEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.RemovedSeriousEventListIDs(); len(nodes) > 0 && !aemuo.mutation.SeriousEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.SeriousEventListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aemuo.mutation.OtherEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.RemovedOtherEventListIDs(); len(nodes) > 0 && !aemuo.mutation.OtherEventListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aemuo.mutation.OtherEventListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdverseEventsModule{config: aemuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aemuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adverseeventsmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
