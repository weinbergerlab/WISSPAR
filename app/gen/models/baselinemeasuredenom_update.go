// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomUpdate is the builder for updating BaselineMeasureDenom entities.
type BaselineMeasureDenomUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureDenomMutation
}

// Where appends a list predicates to the BaselineMeasureDenomUpdate builder.
func (bmdu *BaselineMeasureDenomUpdate) Where(ps ...predicate.BaselineMeasureDenom) *BaselineMeasureDenomUpdate {
	bmdu.mutation.Where(ps...)
	return bmdu
}

// SetBaselineMeasureDenomUnits sets the "baseline_measure_denom_units" field.
func (bmdu *BaselineMeasureDenomUpdate) SetBaselineMeasureDenomUnits(s string) *BaselineMeasureDenomUpdate {
	bmdu.mutation.SetBaselineMeasureDenomUnits(s)
	return bmdu
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bmdu *BaselineMeasureDenomUpdate) SetParentID(id int) *BaselineMeasureDenomUpdate {
	bmdu.mutation.SetParentID(id)
	return bmdu
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bmdu *BaselineMeasureDenomUpdate) SetNillableParentID(id *int) *BaselineMeasureDenomUpdate {
	if id != nil {
		bmdu = bmdu.SetParentID(*id)
	}
	return bmdu
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bmdu *BaselineMeasureDenomUpdate) SetParent(b *BaselineMeasure) *BaselineMeasureDenomUpdate {
	return bmdu.SetParentID(b.ID)
}

// AddBaselineMeasureDenomCountListIDs adds the "baseline_measure_denom_count_list" edge to the BaselineMeasureDenomCount entity by IDs.
func (bmdu *BaselineMeasureDenomUpdate) AddBaselineMeasureDenomCountListIDs(ids ...int) *BaselineMeasureDenomUpdate {
	bmdu.mutation.AddBaselineMeasureDenomCountListIDs(ids...)
	return bmdu
}

// AddBaselineMeasureDenomCountList adds the "baseline_measure_denom_count_list" edges to the BaselineMeasureDenomCount entity.
func (bmdu *BaselineMeasureDenomUpdate) AddBaselineMeasureDenomCountList(b ...*BaselineMeasureDenomCount) *BaselineMeasureDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmdu.AddBaselineMeasureDenomCountListIDs(ids...)
}

// Mutation returns the BaselineMeasureDenomMutation object of the builder.
func (bmdu *BaselineMeasureDenomUpdate) Mutation() *BaselineMeasureDenomMutation {
	return bmdu.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasure entity.
func (bmdu *BaselineMeasureDenomUpdate) ClearParent() *BaselineMeasureDenomUpdate {
	bmdu.mutation.ClearParent()
	return bmdu
}

// ClearBaselineMeasureDenomCountList clears all "baseline_measure_denom_count_list" edges to the BaselineMeasureDenomCount entity.
func (bmdu *BaselineMeasureDenomUpdate) ClearBaselineMeasureDenomCountList() *BaselineMeasureDenomUpdate {
	bmdu.mutation.ClearBaselineMeasureDenomCountList()
	return bmdu
}

// RemoveBaselineMeasureDenomCountListIDs removes the "baseline_measure_denom_count_list" edge to BaselineMeasureDenomCount entities by IDs.
func (bmdu *BaselineMeasureDenomUpdate) RemoveBaselineMeasureDenomCountListIDs(ids ...int) *BaselineMeasureDenomUpdate {
	bmdu.mutation.RemoveBaselineMeasureDenomCountListIDs(ids...)
	return bmdu
}

// RemoveBaselineMeasureDenomCountList removes "baseline_measure_denom_count_list" edges to BaselineMeasureDenomCount entities.
func (bmdu *BaselineMeasureDenomUpdate) RemoveBaselineMeasureDenomCountList(b ...*BaselineMeasureDenomCount) *BaselineMeasureDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmdu.RemoveBaselineMeasureDenomCountListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bmdu *BaselineMeasureDenomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmdu.hooks) == 0 {
		affected, err = bmdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmdu.mutation = mutation
			affected, err = bmdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmdu.hooks) - 1; i >= 0; i-- {
			if bmdu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmdu *BaselineMeasureDenomUpdate) SaveX(ctx context.Context) int {
	affected, err := bmdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bmdu *BaselineMeasureDenomUpdate) Exec(ctx context.Context) error {
	_, err := bmdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdu *BaselineMeasureDenomUpdate) ExecX(ctx context.Context) {
	if err := bmdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmdu *BaselineMeasureDenomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenom.Table,
			Columns: baselinemeasuredenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenom.FieldID,
			},
		},
	}
	if ps := bmdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmdu.mutation.BaselineMeasureDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenom.FieldBaselineMeasureDenomUnits,
		})
	}
	if bmdu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenom.ParentTable,
			Columns: []string{baselinemeasuredenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmdu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenom.ParentTable,
			Columns: []string{baselinemeasuredenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmdu.mutation.BaselineMeasureDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmdu.mutation.RemovedBaselineMeasureDenomCountListIDs(); len(nodes) > 0 && !bmdu.mutation.BaselineMeasureDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmdu.mutation.BaselineMeasureDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bmdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasuredenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineMeasureDenomUpdateOne is the builder for updating a single BaselineMeasureDenom entity.
type BaselineMeasureDenomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineMeasureDenomMutation
}

// SetBaselineMeasureDenomUnits sets the "baseline_measure_denom_units" field.
func (bmduo *BaselineMeasureDenomUpdateOne) SetBaselineMeasureDenomUnits(s string) *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.SetBaselineMeasureDenomUnits(s)
	return bmduo
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bmduo *BaselineMeasureDenomUpdateOne) SetParentID(id int) *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.SetParentID(id)
	return bmduo
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bmduo *BaselineMeasureDenomUpdateOne) SetNillableParentID(id *int) *BaselineMeasureDenomUpdateOne {
	if id != nil {
		bmduo = bmduo.SetParentID(*id)
	}
	return bmduo
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bmduo *BaselineMeasureDenomUpdateOne) SetParent(b *BaselineMeasure) *BaselineMeasureDenomUpdateOne {
	return bmduo.SetParentID(b.ID)
}

// AddBaselineMeasureDenomCountListIDs adds the "baseline_measure_denom_count_list" edge to the BaselineMeasureDenomCount entity by IDs.
func (bmduo *BaselineMeasureDenomUpdateOne) AddBaselineMeasureDenomCountListIDs(ids ...int) *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.AddBaselineMeasureDenomCountListIDs(ids...)
	return bmduo
}

// AddBaselineMeasureDenomCountList adds the "baseline_measure_denom_count_list" edges to the BaselineMeasureDenomCount entity.
func (bmduo *BaselineMeasureDenomUpdateOne) AddBaselineMeasureDenomCountList(b ...*BaselineMeasureDenomCount) *BaselineMeasureDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmduo.AddBaselineMeasureDenomCountListIDs(ids...)
}

// Mutation returns the BaselineMeasureDenomMutation object of the builder.
func (bmduo *BaselineMeasureDenomUpdateOne) Mutation() *BaselineMeasureDenomMutation {
	return bmduo.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasure entity.
func (bmduo *BaselineMeasureDenomUpdateOne) ClearParent() *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.ClearParent()
	return bmduo
}

// ClearBaselineMeasureDenomCountList clears all "baseline_measure_denom_count_list" edges to the BaselineMeasureDenomCount entity.
func (bmduo *BaselineMeasureDenomUpdateOne) ClearBaselineMeasureDenomCountList() *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.ClearBaselineMeasureDenomCountList()
	return bmduo
}

// RemoveBaselineMeasureDenomCountListIDs removes the "baseline_measure_denom_count_list" edge to BaselineMeasureDenomCount entities by IDs.
func (bmduo *BaselineMeasureDenomUpdateOne) RemoveBaselineMeasureDenomCountListIDs(ids ...int) *BaselineMeasureDenomUpdateOne {
	bmduo.mutation.RemoveBaselineMeasureDenomCountListIDs(ids...)
	return bmduo
}

// RemoveBaselineMeasureDenomCountList removes "baseline_measure_denom_count_list" edges to BaselineMeasureDenomCount entities.
func (bmduo *BaselineMeasureDenomUpdateOne) RemoveBaselineMeasureDenomCountList(b ...*BaselineMeasureDenomCount) *BaselineMeasureDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmduo.RemoveBaselineMeasureDenomCountListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bmduo *BaselineMeasureDenomUpdateOne) Select(field string, fields ...string) *BaselineMeasureDenomUpdateOne {
	bmduo.fields = append([]string{field}, fields...)
	return bmduo
}

// Save executes the query and returns the updated BaselineMeasureDenom entity.
func (bmduo *BaselineMeasureDenomUpdateOne) Save(ctx context.Context) (*BaselineMeasureDenom, error) {
	var (
		err  error
		node *BaselineMeasureDenom
	)
	if len(bmduo.hooks) == 0 {
		node, err = bmduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmduo.mutation = mutation
			node, err = bmduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bmduo.hooks) - 1; i >= 0; i-- {
			if bmduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmduo *BaselineMeasureDenomUpdateOne) SaveX(ctx context.Context) *BaselineMeasureDenom {
	node, err := bmduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bmduo *BaselineMeasureDenomUpdateOne) Exec(ctx context.Context) error {
	_, err := bmduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmduo *BaselineMeasureDenomUpdateOne) ExecX(ctx context.Context) {
	if err := bmduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmduo *BaselineMeasureDenomUpdateOne) sqlSave(ctx context.Context) (_node *BaselineMeasureDenom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenom.Table,
			Columns: baselinemeasuredenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenom.FieldID,
			},
		},
	}
	id, ok := bmduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineMeasureDenom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bmduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenom.FieldID)
		for _, f := range fields {
			if !baselinemeasuredenom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinemeasuredenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bmduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmduo.mutation.BaselineMeasureDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenom.FieldBaselineMeasureDenomUnits,
		})
	}
	if bmduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenom.ParentTable,
			Columns: []string{baselinemeasuredenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenom.ParentTable,
			Columns: []string{baselinemeasuredenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bmduo.mutation.BaselineMeasureDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmduo.mutation.RemovedBaselineMeasureDenomCountListIDs(); len(nodes) > 0 && !bmduo.mutation.BaselineMeasureDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmduo.mutation.BaselineMeasureDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasuredenom.BaselineMeasureDenomCountListTable,
			Columns: []string{baselinemeasuredenom.BaselineMeasureDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineMeasureDenom{config: bmduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bmduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasuredenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
