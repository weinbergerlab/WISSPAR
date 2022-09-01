// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomUpdate is the builder for updating BaselineClassDenom entities.
type BaselineClassDenomUpdate struct {
	config
	hooks    []Hook
	mutation *BaselineClassDenomMutation
}

// Where appends a list predicates to the BaselineClassDenomUpdate builder.
func (bcdu *BaselineClassDenomUpdate) Where(ps ...predicate.BaselineClassDenom) *BaselineClassDenomUpdate {
	bcdu.mutation.Where(ps...)
	return bcdu
}

// SetBaselineClassDenomUnits sets the "baseline_class_denom_units" field.
func (bcdu *BaselineClassDenomUpdate) SetBaselineClassDenomUnits(s string) *BaselineClassDenomUpdate {
	bcdu.mutation.SetBaselineClassDenomUnits(s)
	return bcdu
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcdu *BaselineClassDenomUpdate) SetParentID(id int) *BaselineClassDenomUpdate {
	bcdu.mutation.SetParentID(id)
	return bcdu
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcdu *BaselineClassDenomUpdate) SetNillableParentID(id *int) *BaselineClassDenomUpdate {
	if id != nil {
		bcdu = bcdu.SetParentID(*id)
	}
	return bcdu
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcdu *BaselineClassDenomUpdate) SetParent(b *BaselineClass) *BaselineClassDenomUpdate {
	return bcdu.SetParentID(b.ID)
}

// AddBaselineClassDenomCountListIDs adds the "baseline_class_denom_count_list" edge to the BaselineClassDenomCount entity by IDs.
func (bcdu *BaselineClassDenomUpdate) AddBaselineClassDenomCountListIDs(ids ...int) *BaselineClassDenomUpdate {
	bcdu.mutation.AddBaselineClassDenomCountListIDs(ids...)
	return bcdu
}

// AddBaselineClassDenomCountList adds the "baseline_class_denom_count_list" edges to the BaselineClassDenomCount entity.
func (bcdu *BaselineClassDenomUpdate) AddBaselineClassDenomCountList(b ...*BaselineClassDenomCount) *BaselineClassDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcdu.AddBaselineClassDenomCountListIDs(ids...)
}

// Mutation returns the BaselineClassDenomMutation object of the builder.
func (bcdu *BaselineClassDenomUpdate) Mutation() *BaselineClassDenomMutation {
	return bcdu.mutation
}

// ClearParent clears the "parent" edge to the BaselineClass entity.
func (bcdu *BaselineClassDenomUpdate) ClearParent() *BaselineClassDenomUpdate {
	bcdu.mutation.ClearParent()
	return bcdu
}

// ClearBaselineClassDenomCountList clears all "baseline_class_denom_count_list" edges to the BaselineClassDenomCount entity.
func (bcdu *BaselineClassDenomUpdate) ClearBaselineClassDenomCountList() *BaselineClassDenomUpdate {
	bcdu.mutation.ClearBaselineClassDenomCountList()
	return bcdu
}

// RemoveBaselineClassDenomCountListIDs removes the "baseline_class_denom_count_list" edge to BaselineClassDenomCount entities by IDs.
func (bcdu *BaselineClassDenomUpdate) RemoveBaselineClassDenomCountListIDs(ids ...int) *BaselineClassDenomUpdate {
	bcdu.mutation.RemoveBaselineClassDenomCountListIDs(ids...)
	return bcdu
}

// RemoveBaselineClassDenomCountList removes "baseline_class_denom_count_list" edges to BaselineClassDenomCount entities.
func (bcdu *BaselineClassDenomUpdate) RemoveBaselineClassDenomCountList(b ...*BaselineClassDenomCount) *BaselineClassDenomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcdu.RemoveBaselineClassDenomCountListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcdu *BaselineClassDenomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcdu.hooks) == 0 {
		affected, err = bcdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcdu.mutation = mutation
			affected, err = bcdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcdu.hooks) - 1; i >= 0; i-- {
			if bcdu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcdu *BaselineClassDenomUpdate) SaveX(ctx context.Context) int {
	affected, err := bcdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcdu *BaselineClassDenomUpdate) Exec(ctx context.Context) error {
	_, err := bcdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdu *BaselineClassDenomUpdate) ExecX(ctx context.Context) {
	if err := bcdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcdu *BaselineClassDenomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenom.Table,
			Columns: baselineclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenom.FieldID,
			},
		},
	}
	if ps := bcdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcdu.mutation.BaselineClassDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenom.FieldBaselineClassDenomUnits,
		})
	}
	if bcdu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenom.ParentTable,
			Columns: []string{baselineclassdenom.ParentColumn},
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
	if nodes := bcdu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenom.ParentTable,
			Columns: []string{baselineclassdenom.ParentColumn},
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
	if bcdu.mutation.BaselineClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcdu.mutation.RemovedBaselineClassDenomCountListIDs(); len(nodes) > 0 && !bcdu.mutation.BaselineClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcdu.mutation.BaselineClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclassdenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BaselineClassDenomUpdateOne is the builder for updating a single BaselineClassDenom entity.
type BaselineClassDenomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaselineClassDenomMutation
}

// SetBaselineClassDenomUnits sets the "baseline_class_denom_units" field.
func (bcduo *BaselineClassDenomUpdateOne) SetBaselineClassDenomUnits(s string) *BaselineClassDenomUpdateOne {
	bcduo.mutation.SetBaselineClassDenomUnits(s)
	return bcduo
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcduo *BaselineClassDenomUpdateOne) SetParentID(id int) *BaselineClassDenomUpdateOne {
	bcduo.mutation.SetParentID(id)
	return bcduo
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcduo *BaselineClassDenomUpdateOne) SetNillableParentID(id *int) *BaselineClassDenomUpdateOne {
	if id != nil {
		bcduo = bcduo.SetParentID(*id)
	}
	return bcduo
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcduo *BaselineClassDenomUpdateOne) SetParent(b *BaselineClass) *BaselineClassDenomUpdateOne {
	return bcduo.SetParentID(b.ID)
}

// AddBaselineClassDenomCountListIDs adds the "baseline_class_denom_count_list" edge to the BaselineClassDenomCount entity by IDs.
func (bcduo *BaselineClassDenomUpdateOne) AddBaselineClassDenomCountListIDs(ids ...int) *BaselineClassDenomUpdateOne {
	bcduo.mutation.AddBaselineClassDenomCountListIDs(ids...)
	return bcduo
}

// AddBaselineClassDenomCountList adds the "baseline_class_denom_count_list" edges to the BaselineClassDenomCount entity.
func (bcduo *BaselineClassDenomUpdateOne) AddBaselineClassDenomCountList(b ...*BaselineClassDenomCount) *BaselineClassDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcduo.AddBaselineClassDenomCountListIDs(ids...)
}

// Mutation returns the BaselineClassDenomMutation object of the builder.
func (bcduo *BaselineClassDenomUpdateOne) Mutation() *BaselineClassDenomMutation {
	return bcduo.mutation
}

// ClearParent clears the "parent" edge to the BaselineClass entity.
func (bcduo *BaselineClassDenomUpdateOne) ClearParent() *BaselineClassDenomUpdateOne {
	bcduo.mutation.ClearParent()
	return bcduo
}

// ClearBaselineClassDenomCountList clears all "baseline_class_denom_count_list" edges to the BaselineClassDenomCount entity.
func (bcduo *BaselineClassDenomUpdateOne) ClearBaselineClassDenomCountList() *BaselineClassDenomUpdateOne {
	bcduo.mutation.ClearBaselineClassDenomCountList()
	return bcduo
}

// RemoveBaselineClassDenomCountListIDs removes the "baseline_class_denom_count_list" edge to BaselineClassDenomCount entities by IDs.
func (bcduo *BaselineClassDenomUpdateOne) RemoveBaselineClassDenomCountListIDs(ids ...int) *BaselineClassDenomUpdateOne {
	bcduo.mutation.RemoveBaselineClassDenomCountListIDs(ids...)
	return bcduo
}

// RemoveBaselineClassDenomCountList removes "baseline_class_denom_count_list" edges to BaselineClassDenomCount entities.
func (bcduo *BaselineClassDenomUpdateOne) RemoveBaselineClassDenomCountList(b ...*BaselineClassDenomCount) *BaselineClassDenomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcduo.RemoveBaselineClassDenomCountListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcduo *BaselineClassDenomUpdateOne) Select(field string, fields ...string) *BaselineClassDenomUpdateOne {
	bcduo.fields = append([]string{field}, fields...)
	return bcduo
}

// Save executes the query and returns the updated BaselineClassDenom entity.
func (bcduo *BaselineClassDenomUpdateOne) Save(ctx context.Context) (*BaselineClassDenom, error) {
	var (
		err  error
		node *BaselineClassDenom
	)
	if len(bcduo.hooks) == 0 {
		node, err = bcduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcduo.mutation = mutation
			node, err = bcduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcduo.hooks) - 1; i >= 0; i-- {
			if bcduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcduo *BaselineClassDenomUpdateOne) SaveX(ctx context.Context) *BaselineClassDenom {
	node, err := bcduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcduo *BaselineClassDenomUpdateOne) Exec(ctx context.Context) error {
	_, err := bcduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcduo *BaselineClassDenomUpdateOne) ExecX(ctx context.Context) {
	if err := bcduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bcduo *BaselineClassDenomUpdateOne) sqlSave(ctx context.Context) (_node *BaselineClassDenom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenom.Table,
			Columns: baselineclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenom.FieldID,
			},
		},
	}
	id, ok := bcduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "BaselineClassDenom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenom.FieldID)
		for _, f := range fields {
			if !baselineclassdenom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != baselineclassdenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcduo.mutation.BaselineClassDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenom.FieldBaselineClassDenomUnits,
		})
	}
	if bcduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenom.ParentTable,
			Columns: []string{baselineclassdenom.ParentColumn},
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
	if nodes := bcduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenom.ParentTable,
			Columns: []string{baselineclassdenom.ParentColumn},
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
	if bcduo.mutation.BaselineClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcduo.mutation.RemovedBaselineClassDenomCountListIDs(); len(nodes) > 0 && !bcduo.mutation.BaselineClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcduo.mutation.BaselineClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BaselineClassDenom{config: bcduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{baselineclassdenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
