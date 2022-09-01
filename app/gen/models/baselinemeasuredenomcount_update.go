// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomCountUpdate is the builder for updating BaselineMeasureDenomCount entities.
type BaselineMeasureDenomCountUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineMeasureDenomCountMutation
}

// Where appends a list predicates to the BaselineMeasureDenomCountUpdate builder.
func (bmdcu *BaselineMeasureDenomCountUpdate) Where(ps ...predicate.BaselineMeasureDenomCount) *BaselineMeasureDenomCountUpdate {
	bmdcu.mutation.Where(ps...)
	return bmdcu
}

// SetBaselineMeasureDenomCountGroupID sets the "baseline_measure_denom_count_group_id" field.
func (bmdcu *BaselineMeasureDenomCountUpdate) SetBaselineMeasureDenomCountGroupID(s string) *BaselineMeasureDenomCountUpdate {
	bmdcu.mutation.SetBaselineMeasureDenomCountGroupID(s)
	return bmdcu
}

// SetBaselineMeasureDenomCountValue sets the "baseline_measure_denom_count_value" field.
func (bmdcu *BaselineMeasureDenomCountUpdate) SetBaselineMeasureDenomCountValue(s string) *BaselineMeasureDenomCountUpdate {
	bmdcu.mutation.SetBaselineMeasureDenomCountValue(s)
	return bmdcu
}

// SetParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID.
func (bmdcu *BaselineMeasureDenomCountUpdate) SetParentID(id int) *BaselineMeasureDenomCountUpdate {
	bmdcu.mutation.SetParentID(id)
	return bmdcu
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID if the given value is not nil.
func (bmdcu *BaselineMeasureDenomCountUpdate) SetNillableParentID(id *int) *BaselineMeasureDenomCountUpdate {
	if id != nil {
		bmdcu = bmdcu.SetParentID(*id)
	}
	return bmdcu
}

// SetParent sets the "parent" edge to the BaselineMeasureDenom entity.
func (bmdcu *BaselineMeasureDenomCountUpdate) SetParent(b *BaselineMeasureDenom) *BaselineMeasureDenomCountUpdate {
	return bmdcu.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasureDenomCountMutation object of the builder.
func (bmdcu *BaselineMeasureDenomCountUpdate) Mutation() *BaselineMeasureDenomCountMutation {
	return bmdcu.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasureDenom entity.
func (bmdcu *BaselineMeasureDenomCountUpdate) ClearParent() *BaselineMeasureDenomCountUpdate {
	bmdcu.mutation.ClearParent()
	return bmdcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bmdcu *BaselineMeasureDenomCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmdcu.hooks) == 0 {
		affected, err = bmdcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmdcu.mutation = mutation
			affected, err = bmdcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmdcu.hooks) - 1; i >= 0; i-- {
			if bmdcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmdcu *BaselineMeasureDenomCountUpdate) SaveX(ctx context.Context) int {
	affected, err := bmdcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bmdcu *BaselineMeasureDenomCountUpdate) Exec(ctx context.Context) error {
	_, err := bmdcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcu *BaselineMeasureDenomCountUpdate) ExecX(ctx context.Context) {
	if err := bmdcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmdcu *BaselineMeasureDenomCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenomcount.Table,
			Columns: baselinemeasuredenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenomcount.FieldID,
			},
		},
	}
	if ps := bmdcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmdcu.mutation.BaselineMeasureDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID,
		})
	}
	if value, ok := bmdcu.mutation.BaselineMeasureDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountValue,
		})
	}
	if bmdcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenomcount.ParentTable,
			Columns: []string{baselinemeasuredenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmdcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenomcount.ParentTable,
			Columns: []string{baselinemeasuredenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bmdcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasuredenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineMeasureDenomCountUpdateOne is the builder for updating a single BaselineMeasureDenomCount entity.
type BaselineMeasureDenomCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineMeasureDenomCountMutation
}

// SetBaselineMeasureDenomCountGroupID sets the "baseline_measure_denom_count_group_id" field.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SetBaselineMeasureDenomCountGroupID(s string) *BaselineMeasureDenomCountUpdateOne {
	bmdcuo.mutation.SetBaselineMeasureDenomCountGroupID(s)
	return bmdcuo
}

// SetBaselineMeasureDenomCountValue sets the "baseline_measure_denom_count_value" field.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SetBaselineMeasureDenomCountValue(s string) *BaselineMeasureDenomCountUpdateOne {
	bmdcuo.mutation.SetBaselineMeasureDenomCountValue(s)
	return bmdcuo
}

// SetParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SetParentID(id int) *BaselineMeasureDenomCountUpdateOne {
	bmdcuo.mutation.SetParentID(id)
	return bmdcuo
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasureDenom entity by ID if the given value is not nil.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SetNillableParentID(id *int) *BaselineMeasureDenomCountUpdateOne {
	if id != nil {
		bmdcuo = bmdcuo.SetParentID(*id)
	}
	return bmdcuo
}

// SetParent sets the "parent" edge to the BaselineMeasureDenom entity.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SetParent(b *BaselineMeasureDenom) *BaselineMeasureDenomCountUpdateOne {
	return bmdcuo.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasureDenomCountMutation object of the builder.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) Mutation() *BaselineMeasureDenomCountMutation {
	return bmdcuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasureDenom entity.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) ClearParent() *BaselineMeasureDenomCountUpdateOne {
	bmdcuo.mutation.ClearParent()
	return bmdcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) Select(field string, fields ...string) *BaselineMeasureDenomCountUpdateOne {
	bmdcuo.fields = append([]string{field}, fields...)
	return bmdcuo
}

// Save executes the query and returns the updated BaselineMeasureDenomCount entity.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) Save(ctx context.Context) (*BaselineMeasureDenomCount, error) {
	var (
		err  error
		node *BaselineMeasureDenomCount
	)
	if len(bmdcuo.hooks) == 0 {
		node, err = bmdcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmdcuo.mutation = mutation
			node, err = bmdcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bmdcuo.hooks) - 1; i >= 0; i-- {
			if bmdcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmdcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmdcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) SaveX(ctx context.Context) *BaselineMeasureDenomCount {
	node, err := bmdcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) Exec(ctx context.Context) error {
	_, err := bmdcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdcuo *BaselineMeasureDenomCountUpdateOne) ExecX(ctx context.Context) {
	if err := bmdcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmdcuo *BaselineMeasureDenomCountUpdateOne) sqlSave(ctx context.Context) (_node *BaselineMeasureDenomCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenomcount.Table,
			Columns: baselinemeasuredenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenomcount.FieldID,
			},
		},
	}
	id, ok := bmdcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineMeasureDenomCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bmdcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenomcount.FieldID)
		for _, f := range fields {
			if !baselinemeasuredenomcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinemeasuredenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bmdcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmdcuo.mutation.BaselineMeasureDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID,
		})
	}
	if value, ok := bmdcuo.mutation.BaselineMeasureDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasuredenomcount.FieldBaselineMeasureDenomCountValue,
		})
	}
	if bmdcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenomcount.ParentTable,
			Columns: []string{baselinemeasuredenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmdcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasuredenomcount.ParentTable,
			Columns: []string{baselinemeasuredenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineMeasureDenomCount{config: bmdcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bmdcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasuredenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
