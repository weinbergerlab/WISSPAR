// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomCountUpdate is the builder for updating BaselineClassDenomCount entities.
type BaselineClassDenomCountUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineClassDenomCountMutation
}

// Where appends a list predicates to the BaselineClassDenomCountUpdate builder.
func (bcdcu *BaselineClassDenomCountUpdate) Where(ps ...predicate.BaselineClassDenomCount) *BaselineClassDenomCountUpdate {
	bcdcu.mutation.Where(ps...)
	return bcdcu
}

// SetBaselineClassDenomCountGroupID sets the "baseline_class_denom_count_group_id" field.
func (bcdcu *BaselineClassDenomCountUpdate) SetBaselineClassDenomCountGroupID(s string) *BaselineClassDenomCountUpdate {
	bcdcu.mutation.SetBaselineClassDenomCountGroupID(s)
	return bcdcu
}

// SetBaselineClassDenomCountValue sets the "baseline_class_denom_count_value" field.
func (bcdcu *BaselineClassDenomCountUpdate) SetBaselineClassDenomCountValue(s string) *BaselineClassDenomCountUpdate {
	bcdcu.mutation.SetBaselineClassDenomCountValue(s)
	return bcdcu
}

// SetParentID sets the "parent" edge to the BaselineClassDenom entity by ID.
func (bcdcu *BaselineClassDenomCountUpdate) SetParentID(id int) *BaselineClassDenomCountUpdate {
	bcdcu.mutation.SetParentID(id)
	return bcdcu
}

// SetNillableParentID sets the "parent" edge to the BaselineClassDenom entity by ID if the given value is not nil.
func (bcdcu *BaselineClassDenomCountUpdate) SetNillableParentID(id *int) *BaselineClassDenomCountUpdate {
	if id != nil {
		bcdcu = bcdcu.SetParentID(*id)
	}
	return bcdcu
}

// SetParent sets the "parent" edge to the BaselineClassDenom entity.
func (bcdcu *BaselineClassDenomCountUpdate) SetParent(b *BaselineClassDenom) *BaselineClassDenomCountUpdate {
	return bcdcu.SetParentID(b.ID)
}

// Mutation returns the BaselineClassDenomCountMutation object of the builder.
func (bcdcu *BaselineClassDenomCountUpdate) Mutation() *BaselineClassDenomCountMutation {
	return bcdcu.mutation
}

// ClearParent clears the "parent" edge to the BaselineClassDenom entity.
func (bcdcu *BaselineClassDenomCountUpdate) ClearParent() *BaselineClassDenomCountUpdate {
	bcdcu.mutation.ClearParent()
	return bcdcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcdcu *BaselineClassDenomCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcdcu.hooks) == 0 {
		affected, err = bcdcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcdcu.mutation = mutation
			affected, err = bcdcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcdcu.hooks) - 1; i >= 0; i-- {
			if bcdcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcdcu *BaselineClassDenomCountUpdate) SaveX(ctx context.Context) int {
	affected, err := bcdcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcdcu *BaselineClassDenomCountUpdate) Exec(ctx context.Context) error {
	_, err := bcdcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcu *BaselineClassDenomCountUpdate) ExecX(ctx context.Context) {
	if err := bcdcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcdcu *BaselineClassDenomCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenomcount.Table,
			Columns: baselineclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenomcount.FieldID,
			},
		},
	}
	if ps := bcdcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcdcu.mutation.BaselineClassDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountGroupID,
		})
	}
	if value, ok := bcdcu.mutation.BaselineClassDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountValue,
		})
	}
	if bcdcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenomcount.ParentTable,
			Columns: []string{baselineclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcdcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenomcount.ParentTable,
			Columns: []string{baselineclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcdcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclassdenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineClassDenomCountUpdateOne is the builder for updating a single BaselineClassDenomCount entity.
type BaselineClassDenomCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineClassDenomCountMutation
}

// SetBaselineClassDenomCountGroupID sets the "baseline_class_denom_count_group_id" field.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SetBaselineClassDenomCountGroupID(s string) *BaselineClassDenomCountUpdateOne {
	bcdcuo.mutation.SetBaselineClassDenomCountGroupID(s)
	return bcdcuo
}

// SetBaselineClassDenomCountValue sets the "baseline_class_denom_count_value" field.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SetBaselineClassDenomCountValue(s string) *BaselineClassDenomCountUpdateOne {
	bcdcuo.mutation.SetBaselineClassDenomCountValue(s)
	return bcdcuo
}

// SetParentID sets the "parent" edge to the BaselineClassDenom entity by ID.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SetParentID(id int) *BaselineClassDenomCountUpdateOne {
	bcdcuo.mutation.SetParentID(id)
	return bcdcuo
}

// SetNillableParentID sets the "parent" edge to the BaselineClassDenom entity by ID if the given value is not nil.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SetNillableParentID(id *int) *BaselineClassDenomCountUpdateOne {
	if id != nil {
		bcdcuo = bcdcuo.SetParentID(*id)
	}
	return bcdcuo
}

// SetParent sets the "parent" edge to the BaselineClassDenom entity.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SetParent(b *BaselineClassDenom) *BaselineClassDenomCountUpdateOne {
	return bcdcuo.SetParentID(b.ID)
}

// Mutation returns the BaselineClassDenomCountMutation object of the builder.
func (bcdcuo *BaselineClassDenomCountUpdateOne) Mutation() *BaselineClassDenomCountMutation {
	return bcdcuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineClassDenom entity.
func (bcdcuo *BaselineClassDenomCountUpdateOne) ClearParent() *BaselineClassDenomCountUpdateOne {
	bcdcuo.mutation.ClearParent()
	return bcdcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcdcuo *BaselineClassDenomCountUpdateOne) Select(field string, fields ...string) *BaselineClassDenomCountUpdateOne {
	bcdcuo.fields = append([]string{field}, fields...)
	return bcdcuo
}

// Save executes the query and returns the updated BaselineClassDenomCount entity.
func (bcdcuo *BaselineClassDenomCountUpdateOne) Save(ctx context.Context) (*BaselineClassDenomCount, error) {
	var (
		err  error
		node *BaselineClassDenomCount
	)
	if len(bcdcuo.hooks) == 0 {
		node, err = bcdcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcdcuo.mutation = mutation
			node, err = bcdcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcdcuo.hooks) - 1; i >= 0; i-- {
			if bcdcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcdcuo *BaselineClassDenomCountUpdateOne) SaveX(ctx context.Context) *BaselineClassDenomCount {
	node, err := bcdcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcdcuo *BaselineClassDenomCountUpdateOne) Exec(ctx context.Context) error {
	_, err := bcdcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcuo *BaselineClassDenomCountUpdateOne) ExecX(ctx context.Context) {
	if err := bcdcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcdcuo *BaselineClassDenomCountUpdateOne) sqlSave(ctx context.Context) (_node *BaselineClassDenomCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenomcount.Table,
			Columns: baselineclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenomcount.FieldID,
			},
		},
	}
	id, ok := bcdcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineClassDenomCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcdcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenomcount.FieldID)
		for _, f := range fields {
			if !baselineclassdenomcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselineclassdenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcdcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcdcuo.mutation.BaselineClassDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountGroupID,
		})
	}
	if value, ok := bcdcuo.mutation.BaselineClassDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenomcount.FieldBaselineClassDenomCountValue,
		})
	}
	if bcdcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenomcount.ParentTable,
			Columns: []string{baselineclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcdcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenomcount.ParentTable,
			Columns: []string{baselineclassdenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineClassDenomCount{config: bcdcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcdcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclassdenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
