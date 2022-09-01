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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// EventGroupUpdate is the builder for updating EventGroup entities.
type EventGroupUpdate struct {
	config
	hooks    []Hook
	mutation *EventGroupMutation
}

// Where appends a list predicates to the EventGroupUpdate builder.
func (egu *EventGroupUpdate) Where(ps ...predicate.EventGroup) *EventGroupUpdate {
	egu.mutation.Where(ps...)
	return egu
}

// SetEventGroupID sets the "event_group_id" field.
func (egu *EventGroupUpdate) SetEventGroupID(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupID(s)
	return egu
}

// SetEventGroupTitle sets the "event_group_title" field.
func (egu *EventGroupUpdate) SetEventGroupTitle(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupTitle(s)
	return egu
}

// SetEventGroupDescription sets the "event_group_description" field.
func (egu *EventGroupUpdate) SetEventGroupDescription(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupDescription(s)
	return egu
}

// SetEventGroupDeathsNumAffected sets the "event_group_deaths_num_affected" field.
func (egu *EventGroupUpdate) SetEventGroupDeathsNumAffected(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupDeathsNumAffected(s)
	return egu
}

// SetEventGroupDeathsNumAtRisk sets the "event_group_deaths_num_at_risk" field.
func (egu *EventGroupUpdate) SetEventGroupDeathsNumAtRisk(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupDeathsNumAtRisk(s)
	return egu
}

// SetEventGroupSeriousNumAffected sets the "event_group_serious_num_affected" field.
func (egu *EventGroupUpdate) SetEventGroupSeriousNumAffected(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupSeriousNumAffected(s)
	return egu
}

// SetEventGroupSeriousNumAtRisk sets the "event_group_serious_num_at_risk" field.
func (egu *EventGroupUpdate) SetEventGroupSeriousNumAtRisk(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupSeriousNumAtRisk(s)
	return egu
}

// SetEventGroupOtherNumAffected sets the "event_group_other_num_affected" field.
func (egu *EventGroupUpdate) SetEventGroupOtherNumAffected(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupOtherNumAffected(s)
	return egu
}

// SetEventGroupOtherNumAtRisk sets the "event_group_other_num_at_risk" field.
func (egu *EventGroupUpdate) SetEventGroupOtherNumAtRisk(s string) *EventGroupUpdate {
	egu.mutation.SetEventGroupOtherNumAtRisk(s)
	return egu
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (egu *EventGroupUpdate) SetParentID(id int) *EventGroupUpdate {
	egu.mutation.SetParentID(id)
	return egu
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (egu *EventGroupUpdate) SetNillableParentID(id *int) *EventGroupUpdate {
	if id != nil {
		egu = egu.SetParentID(*id)
	}
	return egu
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (egu *EventGroupUpdate) SetParent(a *AdverseEventsModule) *EventGroupUpdate {
	return egu.SetParentID(a.ID)
}

// Mutation returns the EventGroupMutation object of the builder.
func (egu *EventGroupUpdate) Mutation() *EventGroupMutation {
	return egu.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (egu *EventGroupUpdate) ClearParent() *EventGroupUpdate {
	egu.mutation.ClearParent()
	return egu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (egu *EventGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(egu.hooks) == 0 {
		affected, err = egu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			egu.mutation = mutation
			affected, err = egu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(egu.hooks) - 1; i >= 0; i-- {
			if egu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = egu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, egu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (egu *EventGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := egu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (egu *EventGroupUpdate) Exec(ctx context.Context) error {
	_, err := egu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (egu *EventGroupUpdate) ExecX(ctx context.Context) {
	if err := egu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (egu *EventGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventgroup.Table,
			Columns: eventgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventgroup.FieldID,
			},
		},
	}
	if ps := egu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := egu.mutation.EventGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupID,
		})
	}
	if value, ok := egu.mutation.EventGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupTitle,
		})
	}
	if value, ok := egu.mutation.EventGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDescription,
		})
	}
	if value, ok := egu.mutation.EventGroupDeathsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAffected,
		})
	}
	if value, ok := egu.mutation.EventGroupDeathsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAtRisk,
		})
	}
	if value, ok := egu.mutation.EventGroupSeriousNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAffected,
		})
	}
	if value, ok := egu.mutation.EventGroupSeriousNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAtRisk,
		})
	}
	if value, ok := egu.mutation.EventGroupOtherNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAffected,
		})
	}
	if value, ok := egu.mutation.EventGroupOtherNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAtRisk,
		})
	}
	if egu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventgroup.ParentTable,
			Columns: []string{eventgroup.ParentColumn},
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
	if nodes := egu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventgroup.ParentTable,
			Columns: []string{eventgroup.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, egu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventGroupUpdateOne is the builder for updating a single EventGroup entity.
type EventGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventGroupMutation
}

// SetEventGroupID sets the "event_group_id" field.
func (eguo *EventGroupUpdateOne) SetEventGroupID(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupID(s)
	return eguo
}

// SetEventGroupTitle sets the "event_group_title" field.
func (eguo *EventGroupUpdateOne) SetEventGroupTitle(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupTitle(s)
	return eguo
}

// SetEventGroupDescription sets the "event_group_description" field.
func (eguo *EventGroupUpdateOne) SetEventGroupDescription(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupDescription(s)
	return eguo
}

// SetEventGroupDeathsNumAffected sets the "event_group_deaths_num_affected" field.
func (eguo *EventGroupUpdateOne) SetEventGroupDeathsNumAffected(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupDeathsNumAffected(s)
	return eguo
}

// SetEventGroupDeathsNumAtRisk sets the "event_group_deaths_num_at_risk" field.
func (eguo *EventGroupUpdateOne) SetEventGroupDeathsNumAtRisk(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupDeathsNumAtRisk(s)
	return eguo
}

// SetEventGroupSeriousNumAffected sets the "event_group_serious_num_affected" field.
func (eguo *EventGroupUpdateOne) SetEventGroupSeriousNumAffected(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupSeriousNumAffected(s)
	return eguo
}

// SetEventGroupSeriousNumAtRisk sets the "event_group_serious_num_at_risk" field.
func (eguo *EventGroupUpdateOne) SetEventGroupSeriousNumAtRisk(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupSeriousNumAtRisk(s)
	return eguo
}

// SetEventGroupOtherNumAffected sets the "event_group_other_num_affected" field.
func (eguo *EventGroupUpdateOne) SetEventGroupOtherNumAffected(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupOtherNumAffected(s)
	return eguo
}

// SetEventGroupOtherNumAtRisk sets the "event_group_other_num_at_risk" field.
func (eguo *EventGroupUpdateOne) SetEventGroupOtherNumAtRisk(s string) *EventGroupUpdateOne {
	eguo.mutation.SetEventGroupOtherNumAtRisk(s)
	return eguo
}

// SetParentID sets the "parent" edge to the AdverseEventsModule entity by ID.
func (eguo *EventGroupUpdateOne) SetParentID(id int) *EventGroupUpdateOne {
	eguo.mutation.SetParentID(id)
	return eguo
}

// SetNillableParentID sets the "parent" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (eguo *EventGroupUpdateOne) SetNillableParentID(id *int) *EventGroupUpdateOne {
	if id != nil {
		eguo = eguo.SetParentID(*id)
	}
	return eguo
}

// SetParent sets the "parent" edge to the AdverseEventsModule entity.
func (eguo *EventGroupUpdateOne) SetParent(a *AdverseEventsModule) *EventGroupUpdateOne {
	return eguo.SetParentID(a.ID)
}

// Mutation returns the EventGroupMutation object of the builder.
func (eguo *EventGroupUpdateOne) Mutation() *EventGroupMutation {
	return eguo.mutation
}

// ClearParent clears the "parent" edge to the AdverseEventsModule entity.
func (eguo *EventGroupUpdateOne) ClearParent() *EventGroupUpdateOne {
	eguo.mutation.ClearParent()
	return eguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eguo *EventGroupUpdateOne) Select(field string, fields ...string) *EventGroupUpdateOne {
	eguo.fields = append([]string{field}, fields...)
	return eguo
}

// Save executes the query and returns the updated EventGroup entity.
func (eguo *EventGroupUpdateOne) Save(ctx context.Context) (*EventGroup, error) {
	var (
		err  error
		node *EventGroup
	)
	if len(eguo.hooks) == 0 {
		node, err = eguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eguo.mutation = mutation
			node, err = eguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eguo.hooks) - 1; i >= 0; i-- {
			if eguo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = eguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eguo *EventGroupUpdateOne) SaveX(ctx context.Context) *EventGroup {
	node, err := eguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eguo *EventGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := eguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eguo *EventGroupUpdateOne) ExecX(ctx context.Context) {
	if err := eguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eguo *EventGroupUpdateOne) sqlSave(ctx context.Context) (_node *EventGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventgroup.Table,
			Columns: eventgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventgroup.FieldID,
			},
		},
	}
	id, ok := eguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "EventGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventgroup.FieldID)
		for _, f := range fields {
			if !eventgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != eventgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eguo.mutation.EventGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupID,
		})
	}
	if value, ok := eguo.mutation.EventGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupTitle,
		})
	}
	if value, ok := eguo.mutation.EventGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDescription,
		})
	}
	if value, ok := eguo.mutation.EventGroupDeathsNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAffected,
		})
	}
	if value, ok := eguo.mutation.EventGroupDeathsNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupDeathsNumAtRisk,
		})
	}
	if value, ok := eguo.mutation.EventGroupSeriousNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAffected,
		})
	}
	if value, ok := eguo.mutation.EventGroupSeriousNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupSeriousNumAtRisk,
		})
	}
	if value, ok := eguo.mutation.EventGroupOtherNumAffected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAffected,
		})
	}
	if value, ok := eguo.mutation.EventGroupOtherNumAtRisk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventgroup.FieldEventGroupOtherNumAtRisk,
		})
	}
	if eguo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventgroup.ParentTable,
			Columns: []string{eventgroup.ParentColumn},
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
	if nodes := eguo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventgroup.ParentTable,
			Columns: []string{eventgroup.ParentColumn},
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
	_node = &EventGroup{config: eguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
