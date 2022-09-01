// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineGroupUpdate is the builder for updating BaselineGroup entities.
type BaselineGroupUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineGroupMutation
}

// Where appends a list predicates to the BaselineGroupUpdate builder.
func (bgu *BaselineGroupUpdate) Where(ps ...predicate.BaselineGroup) *BaselineGroupUpdate {
	bgu.mutation.Where(ps...)
	return bgu
}

// SetBaselineGroupID sets the "baseline_group_id" field.
func (bgu *BaselineGroupUpdate) SetBaselineGroupID(s string) *BaselineGroupUpdate {
	bgu.mutation.SetBaselineGroupID(s)
	return bgu
}

// SetBaselineGroupTitle sets the "baseline_group_title" field.
func (bgu *BaselineGroupUpdate) SetBaselineGroupTitle(s string) *BaselineGroupUpdate {
	bgu.mutation.SetBaselineGroupTitle(s)
	return bgu
}

// SetBaselineGroupDescription sets the "baseline_group_description" field.
func (bgu *BaselineGroupUpdate) SetBaselineGroupDescription(s string) *BaselineGroupUpdate {
	bgu.mutation.SetBaselineGroupDescription(s)
	return bgu
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bgu *BaselineGroupUpdate) SetParentID(id int) *BaselineGroupUpdate {
	bgu.mutation.SetParentID(id)
	return bgu
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bgu *BaselineGroupUpdate) SetNillableParentID(id *int) *BaselineGroupUpdate {
	if id != nil {
		bgu = bgu.SetParentID(*id)
	}
	return bgu
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bgu *BaselineGroupUpdate) SetParent(b *BaselineCharacteristicsModule) *BaselineGroupUpdate {
	return bgu.SetParentID(b.ID)
}

// Mutation returns the BaselineGroupMutation object of the builder.
func (bgu *BaselineGroupUpdate) Mutation() *BaselineGroupMutation {
	return bgu.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bgu *BaselineGroupUpdate) ClearParent() *BaselineGroupUpdate {
	bgu.mutation.ClearParent()
	return bgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bgu *BaselineGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bgu.hooks) == 0 {
		affected, err = bgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bgu.mutation = mutation
			affected, err = bgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bgu.hooks) - 1; i >= 0; i-- {
			if bgu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bgu *BaselineGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := bgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bgu *BaselineGroupUpdate) Exec(ctx context.Context) error {
	_, err := bgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgu *BaselineGroupUpdate) ExecX(ctx context.Context) {
	if err := bgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bgu *BaselineGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinegroup.Table,
			Columns: baselinegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinegroup.FieldID,
			},
		},
	}
	if ps := bgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bgu.mutation.BaselineGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupID,
		})
	}
	if value, ok := bgu.mutation.BaselineGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupTitle,
		})
	}
	if value, ok := bgu.mutation.BaselineGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupDescription,
		})
	}
	if bgu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinegroup.ParentTable,
			Columns: []string{baselinegroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bgu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinegroup.ParentTable,
			Columns: []string{baselinegroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineGroupUpdateOne is the builder for updating a single BaselineGroup entity.
type BaselineGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineGroupMutation
}

// SetBaselineGroupID sets the "baseline_group_id" field.
func (bguo *BaselineGroupUpdateOne) SetBaselineGroupID(s string) *BaselineGroupUpdateOne {
	bguo.mutation.SetBaselineGroupID(s)
	return bguo
}

// SetBaselineGroupTitle sets the "baseline_group_title" field.
func (bguo *BaselineGroupUpdateOne) SetBaselineGroupTitle(s string) *BaselineGroupUpdateOne {
	bguo.mutation.SetBaselineGroupTitle(s)
	return bguo
}

// SetBaselineGroupDescription sets the "baseline_group_description" field.
func (bguo *BaselineGroupUpdateOne) SetBaselineGroupDescription(s string) *BaselineGroupUpdateOne {
	bguo.mutation.SetBaselineGroupDescription(s)
	return bguo
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bguo *BaselineGroupUpdateOne) SetParentID(id int) *BaselineGroupUpdateOne {
	bguo.mutation.SetParentID(id)
	return bguo
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bguo *BaselineGroupUpdateOne) SetNillableParentID(id *int) *BaselineGroupUpdateOne {
	if id != nil {
		bguo = bguo.SetParentID(*id)
	}
	return bguo
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bguo *BaselineGroupUpdateOne) SetParent(b *BaselineCharacteristicsModule) *BaselineGroupUpdateOne {
	return bguo.SetParentID(b.ID)
}

// Mutation returns the BaselineGroupMutation object of the builder.
func (bguo *BaselineGroupUpdateOne) Mutation() *BaselineGroupMutation {
	return bguo.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bguo *BaselineGroupUpdateOne) ClearParent() *BaselineGroupUpdateOne {
	bguo.mutation.ClearParent()
	return bguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bguo *BaselineGroupUpdateOne) Select(field string, fields ...string) *BaselineGroupUpdateOne {
	bguo.fields = append([]string{field}, fields...)
	return bguo
}

// Save executes the query and returns the updated BaselineGroup entity.
func (bguo *BaselineGroupUpdateOne) Save(ctx context.Context) (*BaselineGroup, error) {
	var (
		err  error
		node *BaselineGroup
	)
	if len(bguo.hooks) == 0 {
		node, err = bguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bguo.mutation = mutation
			node, err = bguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bguo.hooks) - 1; i >= 0; i-- {
			if bguo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bguo *BaselineGroupUpdateOne) SaveX(ctx context.Context) *BaselineGroup {
	node, err := bguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bguo *BaselineGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := bguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bguo *BaselineGroupUpdateOne) ExecX(ctx context.Context) {
	if err := bguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bguo *BaselineGroupUpdateOne) sqlSave(ctx context.Context) (_node *BaselineGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinegroup.Table,
			Columns: baselinegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinegroup.FieldID,
			},
		},
	}
	id, ok := bguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinegroup.FieldID)
		for _, f := range fields {
			if !baselinegroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bguo.mutation.BaselineGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupID,
		})
	}
	if value, ok := bguo.mutation.BaselineGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupTitle,
		})
	}
	if value, ok := bguo.mutation.BaselineGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinegroup.FieldBaselineGroupDescription,
		})
	}
	if bguo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinegroup.ParentTable,
			Columns: []string{baselinegroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bguo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinegroup.ParentTable,
			Columns: []string{baselinegroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineGroup{config: bguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
