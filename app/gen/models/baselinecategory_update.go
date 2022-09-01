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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineCategoryUpdate is the builder for updating BaselineCategory entities.
type BaselineCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineCategoryMutation
}

// Where appends a list predicates to the BaselineCategoryUpdate builder.
func (bcu *BaselineCategoryUpdate) Where(ps ...predicate.BaselineCategory) *BaselineCategoryUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetBaselineCategoryTitle sets the "baseline_category_title" field.
func (bcu *BaselineCategoryUpdate) SetBaselineCategoryTitle(s string) *BaselineCategoryUpdate {
	bcu.mutation.SetBaselineCategoryTitle(s)
	return bcu
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcu *BaselineCategoryUpdate) SetParentID(id int) *BaselineCategoryUpdate {
	bcu.mutation.SetParentID(id)
	return bcu
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcu *BaselineCategoryUpdate) SetNillableParentID(id *int) *BaselineCategoryUpdate {
	if id != nil {
		bcu = bcu.SetParentID(*id)
	}
	return bcu
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcu *BaselineCategoryUpdate) SetParent(b *BaselineClass) *BaselineCategoryUpdate {
	return bcu.SetParentID(b.ID)
}

// AddBaselineMeasurementListIDs adds the "baseline_measurement_list" edge to the BaselineMeasurement entity by IDs.
func (bcu *BaselineCategoryUpdate) AddBaselineMeasurementListIDs(ids ...int) *BaselineCategoryUpdate {
	bcu.mutation.AddBaselineMeasurementListIDs(ids...)
	return bcu
}

// AddBaselineMeasurementList adds the "baseline_measurement_list" edges to the BaselineMeasurement entity.
func (bcu *BaselineCategoryUpdate) AddBaselineMeasurementList(b ...*BaselineMeasurement) *BaselineCategoryUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.AddBaselineMeasurementListIDs(ids...)
}

// Mutation returns the BaselineCategoryMutation object of the builder.
func (bcu *BaselineCategoryUpdate) Mutation() *BaselineCategoryMutation {
	return bcu.mutation
}

// ClearParent clears the "parent" edge to the BaselineClass entity.
func (bcu *BaselineCategoryUpdate) ClearParent() *BaselineCategoryUpdate {
	bcu.mutation.ClearParent()
	return bcu
}

// ClearBaselineMeasurementList clears all "baseline_measurement_list" edges to the BaselineMeasurement entity.
func (bcu *BaselineCategoryUpdate) ClearBaselineMeasurementList() *BaselineCategoryUpdate {
	bcu.mutation.ClearBaselineMeasurementList()
	return bcu
}

// RemoveBaselineMeasurementListIDs removes the "baseline_measurement_list" edge to BaselineMeasurement entities by IDs.
func (bcu *BaselineCategoryUpdate) RemoveBaselineMeasurementListIDs(ids ...int) *BaselineCategoryUpdate {
	bcu.mutation.RemoveBaselineMeasurementListIDs(ids...)
	return bcu
}

// RemoveBaselineMeasurementList removes "baseline_measurement_list" edges to BaselineMeasurement entities.
func (bcu *BaselineCategoryUpdate) RemoveBaselineMeasurementList(b ...*BaselineMeasurement) *BaselineCategoryUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.RemoveBaselineMeasurementListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BaselineCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcu.hooks) == 0 {
		affected, err = bcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcu.mutation = mutation
			affected, err = bcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcu.hooks) - 1; i >= 0; i-- {
			if bcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BaselineCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BaselineCategoryUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BaselineCategoryUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcu *BaselineCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecategory.Table,
			Columns: baselinecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecategory.FieldID,
			},
		},
	}
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.BaselineCategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecategory.FieldBaselineCategoryTitle,
		})
	}
	if bcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinecategory.ParentTable,
			Columns: []string{baselinecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinecategory.ParentTable,
			Columns: []string{baselinecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcu.mutation.BaselineMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.RemovedBaselineMeasurementListIDs(); len(nodes) > 0 && !bcu.mutation.BaselineMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BaselineMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineCategoryUpdateOne is the builder for updating a single BaselineCategory entity.
type BaselineCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineCategoryMutation
}

// SetBaselineCategoryTitle sets the "baseline_category_title" field.
func (bcuo *BaselineCategoryUpdateOne) SetBaselineCategoryTitle(s string) *BaselineCategoryUpdateOne {
	bcuo.mutation.SetBaselineCategoryTitle(s)
	return bcuo
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcuo *BaselineCategoryUpdateOne) SetParentID(id int) *BaselineCategoryUpdateOne {
	bcuo.mutation.SetParentID(id)
	return bcuo
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcuo *BaselineCategoryUpdateOne) SetNillableParentID(id *int) *BaselineCategoryUpdateOne {
	if id != nil {
		bcuo = bcuo.SetParentID(*id)
	}
	return bcuo
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcuo *BaselineCategoryUpdateOne) SetParent(b *BaselineClass) *BaselineCategoryUpdateOne {
	return bcuo.SetParentID(b.ID)
}

// AddBaselineMeasurementListIDs adds the "baseline_measurement_list" edge to the BaselineMeasurement entity by IDs.
func (bcuo *BaselineCategoryUpdateOne) AddBaselineMeasurementListIDs(ids ...int) *BaselineCategoryUpdateOne {
	bcuo.mutation.AddBaselineMeasurementListIDs(ids...)
	return bcuo
}

// AddBaselineMeasurementList adds the "baseline_measurement_list" edges to the BaselineMeasurement entity.
func (bcuo *BaselineCategoryUpdateOne) AddBaselineMeasurementList(b ...*BaselineMeasurement) *BaselineCategoryUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.AddBaselineMeasurementListIDs(ids...)
}

// Mutation returns the BaselineCategoryMutation object of the builder.
func (bcuo *BaselineCategoryUpdateOne) Mutation() *BaselineCategoryMutation {
	return bcuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineClass entity.
func (bcuo *BaselineCategoryUpdateOne) ClearParent() *BaselineCategoryUpdateOne {
	bcuo.mutation.ClearParent()
	return bcuo
}

// ClearBaselineMeasurementList clears all "baseline_measurement_list" edges to the BaselineMeasurement entity.
func (bcuo *BaselineCategoryUpdateOne) ClearBaselineMeasurementList() *BaselineCategoryUpdateOne {
	bcuo.mutation.ClearBaselineMeasurementList()
	return bcuo
}

// RemoveBaselineMeasurementListIDs removes the "baseline_measurement_list" edge to BaselineMeasurement entities by IDs.
func (bcuo *BaselineCategoryUpdateOne) RemoveBaselineMeasurementListIDs(ids ...int) *BaselineCategoryUpdateOne {
	bcuo.mutation.RemoveBaselineMeasurementListIDs(ids...)
	return bcuo
}

// RemoveBaselineMeasurementList removes "baseline_measurement_list" edges to BaselineMeasurement entities.
func (bcuo *BaselineCategoryUpdateOne) RemoveBaselineMeasurementList(b ...*BaselineMeasurement) *BaselineCategoryUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.RemoveBaselineMeasurementListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BaselineCategoryUpdateOne) Select(field string, fields ...string) *BaselineCategoryUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BaselineCategory entity.
func (bcuo *BaselineCategoryUpdateOne) Save(ctx context.Context) (*BaselineCategory, error) {
	var (
		err  error
		node *BaselineCategory
	)
	if len(bcuo.hooks) == 0 {
		node, err = bcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcuo.mutation = mutation
			node, err = bcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcuo.hooks) - 1; i >= 0; i-- {
			if bcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BaselineCategoryUpdateOne) SaveX(ctx context.Context) *BaselineCategory {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BaselineCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BaselineCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcuo *BaselineCategoryUpdateOne) sqlSave(ctx context.Context) (_node *BaselineCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecategory.Table,
			Columns: baselinecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecategory.FieldID,
			},
		},
	}
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecategory.FieldID)
		for _, f := range fields {
			if !baselinecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselinecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.BaselineCategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecategory.FieldBaselineCategoryTitle,
		})
	}
	if bcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinecategory.ParentTable,
			Columns: []string{baselinecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinecategory.ParentTable,
			Columns: []string{baselinecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcuo.mutation.BaselineMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.RemovedBaselineMeasurementListIDs(); len(nodes) > 0 && !bcuo.mutation.BaselineMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BaselineMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineCategory{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselinecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
