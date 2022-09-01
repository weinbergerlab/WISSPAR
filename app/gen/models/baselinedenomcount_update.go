// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomCountUpdate is the builder for updating BaselineDenomCount entities.
type BaselineDenomCountUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineDenomCountMutation
}

// Where appends a list predicates to the BaselineDenomCountUpdate builder.
func (bdcu *BaselineDenomCountUpdate) Where(ps ...predicate.BaselineDenomCount) *BaselineDenomCountUpdate {
	bdcu.mutation.Where(ps...)
	return bdcu
}

// SetBaselineDenomCountGroupID sets the "baseline_denom_count_group_id" field.
func (bdcu *BaselineDenomCountUpdate) SetBaselineDenomCountGroupID(s string) *BaselineDenomCountUpdate {
	bdcu.mutation.SetBaselineDenomCountGroupID(s)
	return bdcu
}

// SetBaselineDenomCountValue sets the "baseline_denom_count_value" field.
func (bdcu *BaselineDenomCountUpdate) SetBaselineDenomCountValue(s string) *BaselineDenomCountUpdate {
	bdcu.mutation.SetBaselineDenomCountValue(s)
	return bdcu
}

// SetParentID sets the "parent" edge to the BaselineDenom entity by ID.
func (bdcu *BaselineDenomCountUpdate) SetParentID(id int) *BaselineDenomCountUpdate {
	bdcu.mutation.SetParentID(id)
	return bdcu
}

// SetNillableParentID sets the "parent" edge to the BaselineDenom entity by ID if the given value is not nil.
func (bdcu *BaselineDenomCountUpdate) SetNillableParentID(id *int) *BaselineDenomCountUpdate {
	if id != nil {
		bdcu = bdcu.SetParentID(*id)
	}
	return bdcu
}

// SetParent sets the "parent" edge to the BaselineDenom entity.
func (bdcu *BaselineDenomCountUpdate) SetParent(b *BaselineDenom) *BaselineDenomCountUpdate {
	return bdcu.SetParentID(b.ID)
}

// Mutation returns the BaselineDenomCountMutation object of the builder.
func (bdcu *BaselineDenomCountUpdate) Mutation() *BaselineDenomCountMutation {
	return bdcu.mutation
}

// ClearParent clears the "parent" edge to the BaselineDenom entity.
func (bdcu *BaselineDenomCountUpdate) ClearParent() *BaselineDenomCountUpdate {
	bdcu.mutation.ClearParent()
	return bdcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bdcu *BaselineDenomCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bdcu.hooks) == 0 {
		affected, err = bdcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdcu.mutation = mutation
			affected, err = bdcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bdcu.hooks) - 1; i >= 0; i-- {
			if bdcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bdcu *BaselineDenomCountUpdate) SaveX(ctx context.Context) int {
	affected, err := bdcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bdcu *BaselineDenomCountUpdate) Exec(ctx context.Context) error {
	_, err := bdcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcu *BaselineDenomCountUpdate) ExecX(ctx context.Context) {
	if err := bdcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bdcu *BaselineDenomCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenomcount.Table,
			Columns: baselinedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenomcount.FieldID,
			},
		},
	}
	if ps := bdcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bdcu.mutation.BaselineDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountGroupID,
		})
	}
	if value, ok := bdcu.mutation.BaselineDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountValue,
		})
	}
	if bdcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenomcount.ParentTable,
			Columns: []string{baselinedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenomcount.ParentTable,
			Columns: []string{baselinedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bdcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinedenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineDenomCountUpdateOne is the builder for updating a single BaselineDenomCount entity.
type BaselineDenomCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineDenomCountMutation
}

// SetBaselineDenomCountGroupID sets the "baseline_denom_count_group_id" field.
func (bdcuo *BaselineDenomCountUpdateOne) SetBaselineDenomCountGroupID(s string) *BaselineDenomCountUpdateOne {
	bdcuo.mutation.SetBaselineDenomCountGroupID(s)
	return bdcuo
}

// SetBaselineDenomCountValue sets the "baseline_denom_count_value" field.
func (bdcuo *BaselineDenomCountUpdateOne) SetBaselineDenomCountValue(s string) *BaselineDenomCountUpdateOne {
	bdcuo.mutation.SetBaselineDenomCountValue(s)
	return bdcuo
}

// SetParentID sets the "parent" edge to the BaselineDenom entity by ID.
func (bdcuo *BaselineDenomCountUpdateOne) SetParentID(id int) *BaselineDenomCountUpdateOne {
	bdcuo.mutation.SetParentID(id)
	return bdcuo
}

// SetNillableParentID sets the "parent" edge to the BaselineDenom entity by ID if the given value is not nil.
func (bdcuo *BaselineDenomCountUpdateOne) SetNillableParentID(id *int) *BaselineDenomCountUpdateOne {
	if id != nil {
		bdcuo = bdcuo.SetParentID(*id)
	}
	return bdcuo
}

// SetParent sets the "parent" edge to the BaselineDenom entity.
func (bdcuo *BaselineDenomCountUpdateOne) SetParent(b *BaselineDenom) *BaselineDenomCountUpdateOne {
	return bdcuo.SetParentID(b.ID)
}

// Mutation returns the BaselineDenomCountMutation object of the builder.
func (bdcuo *BaselineDenomCountUpdateOne) Mutation() *BaselineDenomCountMutation {
	return bdcuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineDenom entity.
func (bdcuo *BaselineDenomCountUpdateOne) ClearParent() *BaselineDenomCountUpdateOne {
	bdcuo.mutation.ClearParent()
	return bdcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bdcuo *BaselineDenomCountUpdateOne) Select(field string, fields ...string) *BaselineDenomCountUpdateOne {
	bdcuo.fields = append([]string{field}, fields...)
	return bdcuo
}

// Save executes the query and returns the updated BaselineDenomCount entity.
func (bdcuo *BaselineDenomCountUpdateOne) Save(ctx context.Context) (*BaselineDenomCount, error) {
	var (
		err  error
		node *BaselineDenomCount
	)
	if len(bdcuo.hooks) == 0 {
		node, err = bdcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdcuo.mutation = mutation
			node, err = bdcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bdcuo.hooks) - 1; i >= 0; i-- {
			if bdcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bdcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bdcuo *BaselineDenomCountUpdateOne) SaveX(ctx context.Context) *BaselineDenomCount {
	node, err := bdcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bdcuo *BaselineDenomCountUpdateOne) Exec(ctx context.Context) error {
	_, err := bdcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcuo *BaselineDenomCountUpdateOne) ExecX(ctx context.Context) {
	if err := bdcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bdcuo *BaselineDenomCountUpdateOne) sqlSave(ctx context.Context) (_node *BaselineDenomCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenomcount.Table,
			Columns: baselinedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenomcount.FieldID,
			},
		},
	}
	id, ok := bdcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineDenomCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bdcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenomcount.FieldID)
		for _, f := range fields {
			if !baselinedenomcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinedenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bdcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bdcuo.mutation.BaselineDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountGroupID,
		})
	}
	if value, ok := bdcuo.mutation.BaselineDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinedenomcount.FieldBaselineDenomCountValue,
		})
	}
	if bdcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenomcount.ParentTable,
			Columns: []string{baselinedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinedenomcount.ParentTable,
			Columns: []string{baselinedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineDenomCount{config: bdcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bdcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinedenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
