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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomUpdate is the builder for updating BaselineDenom entities.
type BaselineDenomUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineDenomMutation
}

// Where appends a list predicates to the BaselineDenomUpdate builder.
func (bdu *BaselineDenomUpdate) Where(ps ...predicate.BaselineDenom) *BaselineDenomUpdate {
	bdu.mutation.Where(ps...)
	return bdu
}

// SetBaselineDenomUnits sets the "baseline_denom_units" field.
func (bdu *BaselineDenomUpdate) SetBaselineDenomUnits(s string) *BaselineDenomUpdate {
	bdu.mutation.SetBaselineDenomUnits(s)
	return bdu
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bdu *BaselineDenomUpdate) SetParentID(id int) *BaselineDenomUpdate {
	bdu.mutation.SetParentID(id)
	return bdu
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bdu *BaselineDenomUpdate) SetNillableParentID(id *int) *BaselineDenomUpdate {
	if id != nil {
		bdu = bdu.SetParentID(*id)
	}
	return bdu
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bdu *BaselineDenomUpdate) SetParent(b *BaselineCharacteristicsModule) *BaselineDenomUpdate {
	return bdu.SetParentID(b.ID)
}

// AddBaselineDenomCountListIDs adds the "baseline_denom_count_list" edge to the BaselineDenomCount entity by IDs.
func (bdu *BaselineDenomUpdate) AddBaselineDenomCountListIDs(ids ...int) *BaselineDenomUpdate {
	bdu.mutation.AddBaselineDenomCountListIDs(ids...)
	return bdu
}

// AddBaselineDenomCountList adds the "baseline_denom_count_list" edges to the BaselineDenomCount entity.
func (bdu *BaselineDenomUpdate) AddBaselineDenomCountList(b ...*BaselineDenomCount) *BaselineDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bdu.AddBaselineDenomCountListIDs(ids...)
}

// Mutation returns the BaselineDenomMutation object of the builder.
func (bdu *BaselineDenomUpdate) Mutation() *BaselineDenomMutation {
	return bdu.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bdu *BaselineDenomUpdate) ClearParent() *BaselineDenomUpdate {
	bdu.mutation.ClearParent()
	return bdu
}

// ClearBaselineDenomCountList clears all "baseline_denom_count_list" edges to the BaselineDenomCount entity.
func (bdu *BaselineDenomUpdate) ClearBaselineDenomCountList() *BaselineDenomUpdate {
	bdu.mutation.ClearBaselineDenomCountList()
	return bdu
}

// RemoveBaselineDenomCountListIDs removes the "baseline_denom_count_list" edge to BaselineDenomCount entities by IDs.
func (bdu *BaselineDenomUpdate) RemoveBaselineDenomCountListIDs(ids ...int) *BaselineDenomUpdate {
	bdu.mutation.RemoveBaselineDenomCountListIDs(ids...)
	return bdu
}

// RemoveBaselineDenomCountList removes "baseline_denom_count_list" edges to BaselineDenomCount entities.
func (bdu *BaselineDenomUpdate) RemoveBaselineDenomCountList(b ...*BaselineDenomCount) *BaselineDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bdu.RemoveBaselineDenomCountListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bdu *BaselineDenomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bdu.hooks) == 0 {
		affected, err = bdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdu.mutation = mutation
			affected, err = bdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bdu.hooks) - 1; i >= 0; i-- {
			if bdu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bdu *BaselineDenomUpdate) SaveX(ctx context.Context) int {
	affected, err := bdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bdu *BaselineDenomUpdate) Exec(ctx context.Context) error {
	_, err := bdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdu *BaselineDenomUpdate) ExecX(ctx context.Context) {
	if err := bdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bdu *BaselineDenomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenom.Table,
			Columns: baselinedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenom.FieldID,
			},
		},
	}
	if ps := bdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bdu.mutation.BaselineDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenom.FieldBaselineDenomUnits,
		})
	}
	if bdu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenom.ParentTable,
			Columns: []string{baselinedenom.ParentColumn},
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
	if nodes := bdu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenom.ParentTable,
			Columns: []string{baselinedenom.ParentColumn},
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
	if bdu.mutation.BaselineDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdu.mutation.RemovedBaselineDenomCountListIDs(); len(nodes) > 0 && !bdu.mutation.BaselineDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdu.mutation.BaselineDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinedenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineDenomUpdateOne is the builder for updating a single BaselineDenom entity.
type BaselineDenomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineDenomMutation
}

// SetBaselineDenomUnits sets the "baseline_denom_units" field.
func (bduo *BaselineDenomUpdateOne) SetBaselineDenomUnits(s string) *BaselineDenomUpdateOne {
	bduo.mutation.SetBaselineDenomUnits(s)
	return bduo
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bduo *BaselineDenomUpdateOne) SetParentID(id int) *BaselineDenomUpdateOne {
	bduo.mutation.SetParentID(id)
	return bduo
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bduo *BaselineDenomUpdateOne) SetNillableParentID(id *int) *BaselineDenomUpdateOne {
	if id != nil {
		bduo = bduo.SetParentID(*id)
	}
	return bduo
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bduo *BaselineDenomUpdateOne) SetParent(b *BaselineCharacteristicsModule) *BaselineDenomUpdateOne {
	return bduo.SetParentID(b.ID)
}

// AddBaselineDenomCountListIDs adds the "baseline_denom_count_list" edge to the BaselineDenomCount entity by IDs.
func (bduo *BaselineDenomUpdateOne) AddBaselineDenomCountListIDs(ids ...int) *BaselineDenomUpdateOne {
	bduo.mutation.AddBaselineDenomCountListIDs(ids...)
	return bduo
}

// AddBaselineDenomCountList adds the "baseline_denom_count_list" edges to the BaselineDenomCount entity.
func (bduo *BaselineDenomUpdateOne) AddBaselineDenomCountList(b ...*BaselineDenomCount) *BaselineDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bduo.AddBaselineDenomCountListIDs(ids...)
}

// Mutation returns the BaselineDenomMutation object of the builder.
func (bduo *BaselineDenomUpdateOne) Mutation() *BaselineDenomMutation {
	return bduo.mutation
}

// ClearParent clears the "parent" edge to the BaselineCharacteristicsModule entity.
func (bduo *BaselineDenomUpdateOne) ClearParent() *BaselineDenomUpdateOne {
	bduo.mutation.ClearParent()
	return bduo
}

// ClearBaselineDenomCountList clears all "baseline_denom_count_list" edges to the BaselineDenomCount entity.
func (bduo *BaselineDenomUpdateOne) ClearBaselineDenomCountList() *BaselineDenomUpdateOne {
	bduo.mutation.ClearBaselineDenomCountList()
	return bduo
}

// RemoveBaselineDenomCountListIDs removes the "baseline_denom_count_list" edge to BaselineDenomCount entities by IDs.
func (bduo *BaselineDenomUpdateOne) RemoveBaselineDenomCountListIDs(ids ...int) *BaselineDenomUpdateOne {
	bduo.mutation.RemoveBaselineDenomCountListIDs(ids...)
	return bduo
}

// RemoveBaselineDenomCountList removes "baseline_denom_count_list" edges to BaselineDenomCount entities.
func (bduo *BaselineDenomUpdateOne) RemoveBaselineDenomCountList(b ...*BaselineDenomCount) *BaselineDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bduo.RemoveBaselineDenomCountListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bduo *BaselineDenomUpdateOne) Select(field string, fields ...string) *BaselineDenomUpdateOne {
	bduo.fields = append([]string{field}, fields...)
	return bduo
}

// Save executes the query and returns the updated BaselineDenom entity.
func (bduo *BaselineDenomUpdateOne) Save(ctx context.Context) (*BaselineDenom, error) {
	var (
		err  error
		node *BaselineDenom
	)
	if len(bduo.hooks) == 0 {
		node, err = bduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bduo.mutation = mutation
			node, err = bduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bduo.hooks) - 1; i >= 0; i-- {
			if bduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bduo *BaselineDenomUpdateOne) SaveX(ctx context.Context) *BaselineDenom {
	node, err := bduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bduo *BaselineDenomUpdateOne) Exec(ctx context.Context) error {
	_, err := bduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bduo *BaselineDenomUpdateOne) ExecX(ctx context.Context) {
	if err := bduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bduo *BaselineDenomUpdateOne) sqlSave(ctx context.Context) (_node *BaselineDenom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenom.Table,
			Columns: baselinedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenom.FieldID,
			},
		},
	}
	id, ok := bduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineDenom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenom.FieldID)
		for _, f := range fields {
			if !baselinedenom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinedenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bduo.mutation.BaselineDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenom.FieldBaselineDenomUnits,
		})
	}
	if bduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenom.ParentTable,
			Columns: []string{baselinedenom.ParentColumn},
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
	if nodes := bduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenom.ParentTable,
			Columns: []string{baselinedenom.ParentColumn},
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
	if bduo.mutation.BaselineDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bduo.mutation.RemovedBaselineDenomCountListIDs(); len(nodes) > 0 && !bduo.mutation.BaselineDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bduo.mutation.BaselineDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinedenom.BaselineDenomCountListTable,
			Columns: []string{baselinedenom.BaselineDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineDenom{config: bduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinedenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
