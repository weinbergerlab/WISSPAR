// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasurementUpdate is the builder for updating BaselineMeasurement entities.
type BaselineMeasurementUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineMeasurementMutation
}

// Where appends a list predicates to the BaselineMeasurementUpdate builder.
func (bmu *BaselineMeasurementUpdate) Where(ps ...predicate.BaselineMeasurement) *BaselineMeasurementUpdate {
	bmu.mutation.Where(ps...)
	return bmu
}

// SetBaselineMeasurementGroupID sets the "baseline_measurement_group_id" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementGroupID(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementGroupID(s)
	return bmu
}

// SetBaselineMeasurementValue sets the "baseline_measurement_value" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementValue(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementValue(s)
	return bmu
}

// SetBaselineMeasurementSpread sets the "baseline_measurement_spread" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementSpread(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementSpread(s)
	return bmu
}

// SetBaselineMeasurementLowerLimit sets the "baseline_measurement_lower_limit" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementLowerLimit(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementLowerLimit(s)
	return bmu
}

// SetBaselineMeasurementUpperLimit sets the "baseline_measurement_upper_limit" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementUpperLimit(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementUpperLimit(s)
	return bmu
}

// SetBaselineMeasurementComment sets the "baseline_measurement_comment" field.
func (bmu *BaselineMeasurementUpdate) SetBaselineMeasurementComment(s string) *BaselineMeasurementUpdate {
	bmu.mutation.SetBaselineMeasurementComment(s)
	return bmu
}

// SetParentID sets the "parent" edge to the BaselineCategory entity by ID.
func (bmu *BaselineMeasurementUpdate) SetParentID(id int) *BaselineMeasurementUpdate {
	bmu.mutation.SetParentID(id)
	return bmu
}

// SetNillableParentID sets the "parent" edge to the BaselineCategory entity by ID if the given value is not nil.
func (bmu *BaselineMeasurementUpdate) SetNillableParentID(id *int) *BaselineMeasurementUpdate {
	if id != nil {
		bmu = bmu.SetParentID(*id)
	}
	return bmu
}

// SetParent sets the "parent" edge to the BaselineCategory entity.
func (bmu *BaselineMeasurementUpdate) SetParent(b *BaselineCategory) *BaselineMeasurementUpdate {
	return bmu.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasurementMutation object of the builder.
func (bmu *BaselineMeasurementUpdate) Mutation() *BaselineMeasurementMutation {
	return bmu.mutation
}

// ClearParent clears the "parent" edge to the BaselineCategory entity.
func (bmu *BaselineMeasurementUpdate) ClearParent() *BaselineMeasurementUpdate {
	bmu.mutation.ClearParent()
	return bmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bmu *BaselineMeasurementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmu.hooks) == 0 {
		affected, err = bmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmu.mutation = mutation
			affected, err = bmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmu.hooks) - 1; i >= 0; i-- {
			if bmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmu *BaselineMeasurementUpdate) SaveX(ctx context.Context) int {
	affected, err := bmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bmu *BaselineMeasurementUpdate) Exec(ctx context.Context) error {
	_, err := bmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmu *BaselineMeasurementUpdate) ExecX(ctx context.Context) {
	if err := bmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmu *BaselineMeasurementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasurement.Table,
			Columns: baselinemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasurement.FieldID,
			},
		},
	}
	if ps := bmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmu.mutation.BaselineMeasurementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementGroupID,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurementValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementValue,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurementSpread(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementSpread,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurementLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementLowerLimit,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurementUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementUpperLimit,
		})
	}
	if value, ok := bmu.mutation.BaselineMeasurementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementComment,
		})
	}
	if bmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasurement.ParentTable,
			Columns: []string{baselinemeasurement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasurement.ParentTable,
			Columns: []string{baselinemeasurement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasurement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineMeasurementUpdateOne is the builder for updating a single BaselineMeasurement entity.
type BaselineMeasurementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineMeasurementMutation
}

// SetBaselineMeasurementGroupID sets the "baseline_measurement_group_id" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementGroupID(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementGroupID(s)
	return bmuo
}

// SetBaselineMeasurementValue sets the "baseline_measurement_value" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementValue(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementValue(s)
	return bmuo
}

// SetBaselineMeasurementSpread sets the "baseline_measurement_spread" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementSpread(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementSpread(s)
	return bmuo
}

// SetBaselineMeasurementLowerLimit sets the "baseline_measurement_lower_limit" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementLowerLimit(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementLowerLimit(s)
	return bmuo
}

// SetBaselineMeasurementUpperLimit sets the "baseline_measurement_upper_limit" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementUpperLimit(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementUpperLimit(s)
	return bmuo
}

// SetBaselineMeasurementComment sets the "baseline_measurement_comment" field.
func (bmuo *BaselineMeasurementUpdateOne) SetBaselineMeasurementComment(s string) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetBaselineMeasurementComment(s)
	return bmuo
}

// SetParentID sets the "parent" edge to the BaselineCategory entity by ID.
func (bmuo *BaselineMeasurementUpdateOne) SetParentID(id int) *BaselineMeasurementUpdateOne {
	bmuo.mutation.SetParentID(id)
	return bmuo
}

// SetNillableParentID sets the "parent" edge to the BaselineCategory entity by ID if the given value is not nil.
func (bmuo *BaselineMeasurementUpdateOne) SetNillableParentID(id *int) *BaselineMeasurementUpdateOne {
	if id != nil {
		bmuo = bmuo.SetParentID(*id)
	}
	return bmuo
}

// SetParent sets the "parent" edge to the BaselineCategory entity.
func (bmuo *BaselineMeasurementUpdateOne) SetParent(b *BaselineCategory) *BaselineMeasurementUpdateOne {
	return bmuo.SetParentID(b.ID)
}

// Mutation returns the BaselineMeasurementMutation object of the builder.
func (bmuo *BaselineMeasurementUpdateOne) Mutation() *BaselineMeasurementMutation {
	return bmuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineCategory entity.
func (bmuo *BaselineMeasurementUpdateOne) ClearParent() *BaselineMeasurementUpdateOne {
	bmuo.mutation.ClearParent()
	return bmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bmuo *BaselineMeasurementUpdateOne) Select(field string, fields ...string) *BaselineMeasurementUpdateOne {
	bmuo.fields = append([]string{field}, fields...)
	return bmuo
}

// Save executes the query and returns the updated BaselineMeasurement entity.
func (bmuo *BaselineMeasurementUpdateOne) Save(ctx context.Context) (*BaselineMeasurement, error) {
	var (
		err  error
		node *BaselineMeasurement
	)
	if len(bmuo.hooks) == 0 {
		node, err = bmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmuo.mutation = mutation
			node, err = bmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bmuo.hooks) - 1; i >= 0; i-- {
			if bmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bmuo *BaselineMeasurementUpdateOne) SaveX(ctx context.Context) *BaselineMeasurement {
	node, err := bmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bmuo *BaselineMeasurementUpdateOne) Exec(ctx context.Context) error {
	_, err := bmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmuo *BaselineMeasurementUpdateOne) ExecX(ctx context.Context) {
	if err := bmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmuo *BaselineMeasurementUpdateOne) sqlSave(ctx context.Context) (_node *BaselineMeasurement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasurement.Table,
			Columns: baselinemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasurement.FieldID,
			},
		},
	}
	id, ok := bmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineMeasurement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasurement.FieldID)
		for _, f := range fields {
			if !baselinemeasurement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinemeasurement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmuo.mutation.BaselineMeasurementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementGroupID,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurementValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementValue,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurementSpread(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementSpread,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurementLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementLowerLimit,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurementUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementUpperLimit,
		})
	}
	if value, ok := bmuo.mutation.BaselineMeasurementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasurement.FieldBaselineMeasurementComment,
		})
	}
	if bmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasurement.ParentTable,
			Columns: []string{baselinemeasurement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasurement.ParentTable,
			Columns: []string{baselinemeasurement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineMeasurement{config: bmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinemeasurement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
