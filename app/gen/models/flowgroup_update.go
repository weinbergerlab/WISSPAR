// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowGroupUpdate is the builder for updating FlowGroup entities.
type FlowGroupUpdate struct {
	config
	hooks    []Hook
	mutation *FlowGroupMutation
}

// Where appends a list predicates to the FlowGroupUpdate builder.
func (fgu *FlowGroupUpdate) Where(ps ...predicate.FlowGroup) *FlowGroupUpdate {
	fgu.mutation.Where(ps...)
	return fgu
}

// SetFlowGroupID sets the "flow_group_id" field.
func (fgu *FlowGroupUpdate) SetFlowGroupID(s string) *FlowGroupUpdate {
	fgu.mutation.SetFlowGroupID(s)
	return fgu
}

// SetFlowGroupTitle sets the "flow_group_title" field.
func (fgu *FlowGroupUpdate) SetFlowGroupTitle(s string) *FlowGroupUpdate {
	fgu.mutation.SetFlowGroupTitle(s)
	return fgu
}

// SetFlowGroupDescription sets the "flow_group_description" field.
func (fgu *FlowGroupUpdate) SetFlowGroupDescription(s string) *FlowGroupUpdate {
	fgu.mutation.SetFlowGroupDescription(s)
	return fgu
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fgu *FlowGroupUpdate) SetParentID(id int) *FlowGroupUpdate {
	fgu.mutation.SetParentID(id)
	return fgu
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fgu *FlowGroupUpdate) SetNillableParentID(id *int) *FlowGroupUpdate {
	if id != nil {
		fgu = fgu.SetParentID(*id)
	}
	return fgu
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fgu *FlowGroupUpdate) SetParent(p *ParticipantFlowModule) *FlowGroupUpdate {
	return fgu.SetParentID(p.ID)
}

// Mutation returns the FlowGroupMutation object of the builder.
func (fgu *FlowGroupUpdate) Mutation() *FlowGroupMutation {
	return fgu.mutation
}

// ClearParent clears the "parent" edge to the ParticipantFlowModule entity.
func (fgu *FlowGroupUpdate) ClearParent() *FlowGroupUpdate {
	fgu.mutation.ClearParent()
	return fgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fgu *FlowGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fgu.hooks) == 0 {
		affected, err = fgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fgu.mutation = mutation
			affected, err = fgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fgu.hooks) - 1; i >= 0; i-- {
			if fgu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fgu *FlowGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := fgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fgu *FlowGroupUpdate) Exec(ctx context.Context) error {
	_, err := fgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgu *FlowGroupUpdate) ExecX(ctx context.Context) {
	if err := fgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fgu *FlowGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowgroup.Table,
			Columns: flowgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowgroup.FieldID,
			},
		},
	}
	if ps := fgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fgu.mutation.FlowGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupID,
		})
	}
	if value, ok := fgu.mutation.FlowGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupTitle,
		})
	}
	if value, ok := fgu.mutation.FlowGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupDescription,
		})
	}
	if fgu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowgroup.ParentTable,
			Columns: []string{flowgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fgu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowgroup.ParentTable,
			Columns: []string{flowgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowGroupUpdateOne is the builder for updating a single FlowGroup entity.
type FlowGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowGroupMutation
}

// SetFlowGroupID sets the "flow_group_id" field.
func (fguo *FlowGroupUpdateOne) SetFlowGroupID(s string) *FlowGroupUpdateOne {
	fguo.mutation.SetFlowGroupID(s)
	return fguo
}

// SetFlowGroupTitle sets the "flow_group_title" field.
func (fguo *FlowGroupUpdateOne) SetFlowGroupTitle(s string) *FlowGroupUpdateOne {
	fguo.mutation.SetFlowGroupTitle(s)
	return fguo
}

// SetFlowGroupDescription sets the "flow_group_description" field.
func (fguo *FlowGroupUpdateOne) SetFlowGroupDescription(s string) *FlowGroupUpdateOne {
	fguo.mutation.SetFlowGroupDescription(s)
	return fguo
}

// SetParentID sets the "parent" edge to the ParticipantFlowModule entity by ID.
func (fguo *FlowGroupUpdateOne) SetParentID(id int) *FlowGroupUpdateOne {
	fguo.mutation.SetParentID(id)
	return fguo
}

// SetNillableParentID sets the "parent" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (fguo *FlowGroupUpdateOne) SetNillableParentID(id *int) *FlowGroupUpdateOne {
	if id != nil {
		fguo = fguo.SetParentID(*id)
	}
	return fguo
}

// SetParent sets the "parent" edge to the ParticipantFlowModule entity.
func (fguo *FlowGroupUpdateOne) SetParent(p *ParticipantFlowModule) *FlowGroupUpdateOne {
	return fguo.SetParentID(p.ID)
}

// Mutation returns the FlowGroupMutation object of the builder.
func (fguo *FlowGroupUpdateOne) Mutation() *FlowGroupMutation {
	return fguo.mutation
}

// ClearParent clears the "parent" edge to the ParticipantFlowModule entity.
func (fguo *FlowGroupUpdateOne) ClearParent() *FlowGroupUpdateOne {
	fguo.mutation.ClearParent()
	return fguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fguo *FlowGroupUpdateOne) Select(field string, fields ...string) *FlowGroupUpdateOne {
	fguo.fields = append([]string{field}, fields...)
	return fguo
}

// Save executes the query and returns the updated FlowGroup entity.
func (fguo *FlowGroupUpdateOne) Save(ctx context.Context) (*FlowGroup, error) {
	var (
		err  error
		node *FlowGroup
	)
	if len(fguo.hooks) == 0 {
		node, err = fguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fguo.mutation = mutation
			node, err = fguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fguo.hooks) - 1; i >= 0; i-- {
			if fguo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fguo *FlowGroupUpdateOne) SaveX(ctx context.Context) *FlowGroup {
	node, err := fguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fguo *FlowGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := fguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fguo *FlowGroupUpdateOne) ExecX(ctx context.Context) {
	if err := fguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fguo *FlowGroupUpdateOne) sqlSave(ctx context.Context) (_node *FlowGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowgroup.Table,
			Columns: flowgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowgroup.FieldID,
			},
		},
	}
	id, ok := fguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowgroup.FieldID)
		for _, f := range fields {
			if !flowgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fguo.mutation.FlowGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupID,
		})
	}
	if value, ok := fguo.mutation.FlowGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupTitle,
		})
	}
	if value, ok := fguo.mutation.FlowGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowgroup.FieldFlowGroupDescription,
		})
	}
	if fguo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowgroup.ParentTable,
			Columns: []string{flowgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fguo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowgroup.ParentTable,
			Columns: []string{flowgroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowGroup{config: fguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
