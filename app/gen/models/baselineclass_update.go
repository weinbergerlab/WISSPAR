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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassUpdate is the builder for updating BaselineClass entities.
type BaselineClassUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineClassMutation
}

// Where appends a list predicates to the BaselineClassUpdate builder.
func (bcu *BaselineClassUpdate) Where(ps ...predicate.BaselineClass) *BaselineClassUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetBaselineClassTitle sets the "baseline_class_title" field.
func (bcu *BaselineClassUpdate) SetBaselineClassTitle(s string) *BaselineClassUpdate {
	bcu.mutation.SetBaselineClassTitle(s)
	return bcu
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bcu *BaselineClassUpdate) SetParentID(id int) *BaselineClassUpdate {
	bcu.mutation.SetParentID(id)
	return bcu
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bcu *BaselineClassUpdate) SetNillableParentID(id *int) *BaselineClassUpdate {
	if id != nil {
		bcu = bcu.SetParentID(*id)
	}
	return bcu
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bcu *BaselineClassUpdate) SetParent(b *BaselineMeasure) *BaselineClassUpdate {
	return bcu.SetParentID(b.ID)
}

// AddBaselineClassDenomListIDs adds the "baseline_class_denom_list" edge to the BaselineClassDenom entity by IDs.
func (bcu *BaselineClassUpdate) AddBaselineClassDenomListIDs(ids ...int) *BaselineClassUpdate {
	bcu.mutation.AddBaselineClassDenomListIDs(ids...)
	return bcu
}

// AddBaselineClassDenomList adds the "baseline_class_denom_list" edges to the BaselineClassDenom entity.
func (bcu *BaselineClassUpdate) AddBaselineClassDenomList(b ...*BaselineClassDenom) *BaselineClassUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.AddBaselineClassDenomListIDs(ids...)
}

// AddBaselineCategoryListIDs adds the "baseline_category_list" edge to the BaselineCategory entity by IDs.
func (bcu *BaselineClassUpdate) AddBaselineCategoryListIDs(ids ...int) *BaselineClassUpdate {
	bcu.mutation.AddBaselineCategoryListIDs(ids...)
	return bcu
}

// AddBaselineCategoryList adds the "baseline_category_list" edges to the BaselineCategory entity.
func (bcu *BaselineClassUpdate) AddBaselineCategoryList(b ...*BaselineCategory) *BaselineClassUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.AddBaselineCategoryListIDs(ids...)
}

// Mutation returns the BaselineClassMutation object of the builder.
func (bcu *BaselineClassUpdate) Mutation() *BaselineClassMutation {
	return bcu.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasure entity.
func (bcu *BaselineClassUpdate) ClearParent() *BaselineClassUpdate {
	bcu.mutation.ClearParent()
	return bcu
}

// ClearBaselineClassDenomList clears all "baseline_class_denom_list" edges to the BaselineClassDenom entity.
func (bcu *BaselineClassUpdate) ClearBaselineClassDenomList() *BaselineClassUpdate {
	bcu.mutation.ClearBaselineClassDenomList()
	return bcu
}

// RemoveBaselineClassDenomListIDs removes the "baseline_class_denom_list" edge to BaselineClassDenom entities by IDs.
func (bcu *BaselineClassUpdate) RemoveBaselineClassDenomListIDs(ids ...int) *BaselineClassUpdate {
	bcu.mutation.RemoveBaselineClassDenomListIDs(ids...)
	return bcu
}

// RemoveBaselineClassDenomList removes "baseline_class_denom_list" edges to BaselineClassDenom entities.
func (bcu *BaselineClassUpdate) RemoveBaselineClassDenomList(b ...*BaselineClassDenom) *BaselineClassUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.RemoveBaselineClassDenomListIDs(ids...)
}

// ClearBaselineCategoryList clears all "baseline_category_list" edges to the BaselineCategory entity.
func (bcu *BaselineClassUpdate) ClearBaselineCategoryList() *BaselineClassUpdate {
	bcu.mutation.ClearBaselineCategoryList()
	return bcu
}

// RemoveBaselineCategoryListIDs removes the "baseline_category_list" edge to BaselineCategory entities by IDs.
func (bcu *BaselineClassUpdate) RemoveBaselineCategoryListIDs(ids ...int) *BaselineClassUpdate {
	bcu.mutation.RemoveBaselineCategoryListIDs(ids...)
	return bcu
}

// RemoveBaselineCategoryList removes "baseline_category_list" edges to BaselineCategory entities.
func (bcu *BaselineClassUpdate) RemoveBaselineCategoryList(b ...*BaselineCategory) *BaselineClassUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.RemoveBaselineCategoryListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BaselineClassUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcu.hooks) == 0 {
		affected, err = bcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassMutation)
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
func (bcu *BaselineClassUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BaselineClassUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BaselineClassUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcu *BaselineClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclass.Table,
			Columns: baselineclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclass.FieldID,
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
	if value, ok := bcu.mutation.BaselineClassTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclass.FieldBaselineClassTitle,
		})
	}
	if bcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclass.ParentTable,
			Columns: []string{baselineclass.ParentColumn},
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
	if nodes := bcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclass.ParentTable,
			Columns: []string{baselineclass.ParentColumn},
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
	if bcu.mutation.BaselineClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
	if nodes := bcu.mutation.RemovedBaselineClassDenomListIDs(); len(nodes) > 0 && !bcu.mutation.BaselineClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BaselineClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
	if bcu.mutation.BaselineCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
	if nodes := bcu.mutation.RemovedBaselineCategoryListIDs(); len(nodes) > 0 && !bcu.mutation.BaselineCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BaselineCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineClassUpdateOne is the builder for updating a single BaselineClass entity.
type BaselineClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineClassMutation
}

// SetBaselineClassTitle sets the "baseline_class_title" field.
func (bcuo *BaselineClassUpdateOne) SetBaselineClassTitle(s string) *BaselineClassUpdateOne {
	bcuo.mutation.SetBaselineClassTitle(s)
	return bcuo
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bcuo *BaselineClassUpdateOne) SetParentID(id int) *BaselineClassUpdateOne {
	bcuo.mutation.SetParentID(id)
	return bcuo
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bcuo *BaselineClassUpdateOne) SetNillableParentID(id *int) *BaselineClassUpdateOne {
	if id != nil {
		bcuo = bcuo.SetParentID(*id)
	}
	return bcuo
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bcuo *BaselineClassUpdateOne) SetParent(b *BaselineMeasure) *BaselineClassUpdateOne {
	return bcuo.SetParentID(b.ID)
}

// AddBaselineClassDenomListIDs adds the "baseline_class_denom_list" edge to the BaselineClassDenom entity by IDs.
func (bcuo *BaselineClassUpdateOne) AddBaselineClassDenomListIDs(ids ...int) *BaselineClassUpdateOne {
	bcuo.mutation.AddBaselineClassDenomListIDs(ids...)
	return bcuo
}

// AddBaselineClassDenomList adds the "baseline_class_denom_list" edges to the BaselineClassDenom entity.
func (bcuo *BaselineClassUpdateOne) AddBaselineClassDenomList(b ...*BaselineClassDenom) *BaselineClassUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.AddBaselineClassDenomListIDs(ids...)
}

// AddBaselineCategoryListIDs adds the "baseline_category_list" edge to the BaselineCategory entity by IDs.
func (bcuo *BaselineClassUpdateOne) AddBaselineCategoryListIDs(ids ...int) *BaselineClassUpdateOne {
	bcuo.mutation.AddBaselineCategoryListIDs(ids...)
	return bcuo
}

// AddBaselineCategoryList adds the "baseline_category_list" edges to the BaselineCategory entity.
func (bcuo *BaselineClassUpdateOne) AddBaselineCategoryList(b ...*BaselineCategory) *BaselineClassUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.AddBaselineCategoryListIDs(ids...)
}

// Mutation returns the BaselineClassMutation object of the builder.
func (bcuo *BaselineClassUpdateOne) Mutation() *BaselineClassMutation {
	return bcuo.mutation
}

// ClearParent clears the "parent" edge to the BaselineMeasure entity.
func (bcuo *BaselineClassUpdateOne) ClearParent() *BaselineClassUpdateOne {
	bcuo.mutation.ClearParent()
	return bcuo
}

// ClearBaselineClassDenomList clears all "baseline_class_denom_list" edges to the BaselineClassDenom entity.
func (bcuo *BaselineClassUpdateOne) ClearBaselineClassDenomList() *BaselineClassUpdateOne {
	bcuo.mutation.ClearBaselineClassDenomList()
	return bcuo
}

// RemoveBaselineClassDenomListIDs removes the "baseline_class_denom_list" edge to BaselineClassDenom entities by IDs.
func (bcuo *BaselineClassUpdateOne) RemoveBaselineClassDenomListIDs(ids ...int) *BaselineClassUpdateOne {
	bcuo.mutation.RemoveBaselineClassDenomListIDs(ids...)
	return bcuo
}

// RemoveBaselineClassDenomList removes "baseline_class_denom_list" edges to BaselineClassDenom entities.
func (bcuo *BaselineClassUpdateOne) RemoveBaselineClassDenomList(b ...*BaselineClassDenom) *BaselineClassUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.RemoveBaselineClassDenomListIDs(ids...)
}

// ClearBaselineCategoryList clears all "baseline_category_list" edges to the BaselineCategory entity.
func (bcuo *BaselineClassUpdateOne) ClearBaselineCategoryList() *BaselineClassUpdateOne {
	bcuo.mutation.ClearBaselineCategoryList()
	return bcuo
}

// RemoveBaselineCategoryListIDs removes the "baseline_category_list" edge to BaselineCategory entities by IDs.
func (bcuo *BaselineClassUpdateOne) RemoveBaselineCategoryListIDs(ids ...int) *BaselineClassUpdateOne {
	bcuo.mutation.RemoveBaselineCategoryListIDs(ids...)
	return bcuo
}

// RemoveBaselineCategoryList removes "baseline_category_list" edges to BaselineCategory entities.
func (bcuo *BaselineClassUpdateOne) RemoveBaselineCategoryList(b ...*BaselineCategory) *BaselineClassUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.RemoveBaselineCategoryListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BaselineClassUpdateOne) Select(field string, fields ...string) *BaselineClassUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BaselineClass entity.
func (bcuo *BaselineClassUpdateOne) Save(ctx context.Context) (*BaselineClass, error) {
	var (
		err  error
		node *BaselineClass
	)
	if len(bcuo.hooks) == 0 {
		node, err = bcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassMutation)
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
func (bcuo *BaselineClassUpdateOne) SaveX(ctx context.Context) *BaselineClass {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BaselineClassUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BaselineClassUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcuo *BaselineClassUpdateOne) sqlSave(ctx context.Context) (_node *BaselineClass, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclass.Table,
			Columns: baselineclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclass.FieldID,
			},
		},
	}
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineClass.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclass.FieldID)
		for _, f := range fields {
			if !baselineclass.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselineclass.FieldID {
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
	if value, ok := bcuo.mutation.BaselineClassTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclass.FieldBaselineClassTitle,
		})
	}
	if bcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclass.ParentTable,
			Columns: []string{baselineclass.ParentColumn},
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
	if nodes := bcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclass.ParentTable,
			Columns: []string{baselineclass.ParentColumn},
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
	if bcuo.mutation.BaselineClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
	if nodes := bcuo.mutation.RemovedBaselineClassDenomListIDs(); len(nodes) > 0 && !bcuo.mutation.BaselineClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BaselineClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
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
	if bcuo.mutation.BaselineCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
	if nodes := bcuo.mutation.RemovedBaselineCategoryListIDs(); len(nodes) > 0 && !bcuo.mutation.BaselineCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BaselineCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
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
	_node = &BaselineClass{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
