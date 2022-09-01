// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeCategoryUpdate is the builder for updating OutcomeCategory entities.
type OutcomeCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeCategoryMutation
}

// Where appends a list predicates to the OutcomeCategoryUpdate builder.
func (ocu *OutcomeCategoryUpdate) Where(ps ...predicate.OutcomeCategory) *OutcomeCategoryUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetOutcomeCategoryTitle sets the "outcome_category_title" field.
func (ocu *OutcomeCategoryUpdate) SetOutcomeCategoryTitle(s string) *OutcomeCategoryUpdate {
	ocu.mutation.SetOutcomeCategoryTitle(s)
	return ocu
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (ocu *OutcomeCategoryUpdate) SetParentID(id int) *OutcomeCategoryUpdate {
	ocu.mutation.SetParentID(id)
	return ocu
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (ocu *OutcomeCategoryUpdate) SetNillableParentID(id *int) *OutcomeCategoryUpdate {
	if id != nil {
		ocu = ocu.SetParentID(*id)
	}
	return ocu
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (ocu *OutcomeCategoryUpdate) SetParent(o *OutcomeClass) *OutcomeCategoryUpdate {
	return ocu.SetParentID(o.ID)
}

// AddOutcomeMeasurementListIDs adds the "outcome_measurement_list" edge to the OutcomeMeasurement entity by IDs.
func (ocu *OutcomeCategoryUpdate) AddOutcomeMeasurementListIDs(ids ...int) *OutcomeCategoryUpdate {
	ocu.mutation.AddOutcomeMeasurementListIDs(ids...)
	return ocu
}

// AddOutcomeMeasurementList adds the "outcome_measurement_list" edges to the OutcomeMeasurement entity.
func (ocu *OutcomeCategoryUpdate) AddOutcomeMeasurementList(o ...*OutcomeMeasurement) *OutcomeCategoryUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.AddOutcomeMeasurementListIDs(ids...)
}

// Mutation returns the OutcomeCategoryMutation object of the builder.
func (ocu *OutcomeCategoryUpdate) Mutation() *OutcomeCategoryMutation {
	return ocu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClass entity.
func (ocu *OutcomeCategoryUpdate) ClearParent() *OutcomeCategoryUpdate {
	ocu.mutation.ClearParent()
	return ocu
}

// ClearOutcomeMeasurementList clears all "outcome_measurement_list" edges to the OutcomeMeasurement entity.
func (ocu *OutcomeCategoryUpdate) ClearOutcomeMeasurementList() *OutcomeCategoryUpdate {
	ocu.mutation.ClearOutcomeMeasurementList()
	return ocu
}

// RemoveOutcomeMeasurementListIDs removes the "outcome_measurement_list" edge to OutcomeMeasurement entities by IDs.
func (ocu *OutcomeCategoryUpdate) RemoveOutcomeMeasurementListIDs(ids ...int) *OutcomeCategoryUpdate {
	ocu.mutation.RemoveOutcomeMeasurementListIDs(ids...)
	return ocu
}

// RemoveOutcomeMeasurementList removes "outcome_measurement_list" edges to OutcomeMeasurement entities.
func (ocu *OutcomeCategoryUpdate) RemoveOutcomeMeasurementList(o ...*OutcomeMeasurement) *OutcomeCategoryUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.RemoveOutcomeMeasurementListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OutcomeCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocu.hooks) == 0 {
		affected, err = ocu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocu.mutation = mutation
			affected, err = ocu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocu.hooks) - 1; i >= 0; i-- {
			if ocu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocu *OutcomeCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OutcomeCategoryUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OutcomeCategoryUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocu *OutcomeCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomecategory.Table,
			Columns: outcomecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomecategory.FieldID,
			},
		},
	}
	if ps := ocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocu.mutation.OutcomeCategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomecategory.FieldOutcomeCategoryTitle,
		})
	}
	if ocu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomecategory.ParentTable,
			Columns: []string{outcomecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomecategory.ParentTable,
			Columns: []string{outcomecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocu.mutation.OutcomeMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.RemovedOutcomeMeasurementListIDs(); len(nodes) > 0 && !ocu.mutation.OutcomeMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.OutcomeMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeCategoryUpdateOne is the builder for updating a single OutcomeCategory entity.
type OutcomeCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeCategoryMutation
}

// SetOutcomeCategoryTitle sets the "outcome_category_title" field.
func (ocuo *OutcomeCategoryUpdateOne) SetOutcomeCategoryTitle(s string) *OutcomeCategoryUpdateOne {
	ocuo.mutation.SetOutcomeCategoryTitle(s)
	return ocuo
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (ocuo *OutcomeCategoryUpdateOne) SetParentID(id int) *OutcomeCategoryUpdateOne {
	ocuo.mutation.SetParentID(id)
	return ocuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (ocuo *OutcomeCategoryUpdateOne) SetNillableParentID(id *int) *OutcomeCategoryUpdateOne {
	if id != nil {
		ocuo = ocuo.SetParentID(*id)
	}
	return ocuo
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (ocuo *OutcomeCategoryUpdateOne) SetParent(o *OutcomeClass) *OutcomeCategoryUpdateOne {
	return ocuo.SetParentID(o.ID)
}

// AddOutcomeMeasurementListIDs adds the "outcome_measurement_list" edge to the OutcomeMeasurement entity by IDs.
func (ocuo *OutcomeCategoryUpdateOne) AddOutcomeMeasurementListIDs(ids ...int) *OutcomeCategoryUpdateOne {
	ocuo.mutation.AddOutcomeMeasurementListIDs(ids...)
	return ocuo
}

// AddOutcomeMeasurementList adds the "outcome_measurement_list" edges to the OutcomeMeasurement entity.
func (ocuo *OutcomeCategoryUpdateOne) AddOutcomeMeasurementList(o ...*OutcomeMeasurement) *OutcomeCategoryUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.AddOutcomeMeasurementListIDs(ids...)
}

// Mutation returns the OutcomeCategoryMutation object of the builder.
func (ocuo *OutcomeCategoryUpdateOne) Mutation() *OutcomeCategoryMutation {
	return ocuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClass entity.
func (ocuo *OutcomeCategoryUpdateOne) ClearParent() *OutcomeCategoryUpdateOne {
	ocuo.mutation.ClearParent()
	return ocuo
}

// ClearOutcomeMeasurementList clears all "outcome_measurement_list" edges to the OutcomeMeasurement entity.
func (ocuo *OutcomeCategoryUpdateOne) ClearOutcomeMeasurementList() *OutcomeCategoryUpdateOne {
	ocuo.mutation.ClearOutcomeMeasurementList()
	return ocuo
}

// RemoveOutcomeMeasurementListIDs removes the "outcome_measurement_list" edge to OutcomeMeasurement entities by IDs.
func (ocuo *OutcomeCategoryUpdateOne) RemoveOutcomeMeasurementListIDs(ids ...int) *OutcomeCategoryUpdateOne {
	ocuo.mutation.RemoveOutcomeMeasurementListIDs(ids...)
	return ocuo
}

// RemoveOutcomeMeasurementList removes "outcome_measurement_list" edges to OutcomeMeasurement entities.
func (ocuo *OutcomeCategoryUpdateOne) RemoveOutcomeMeasurementList(o ...*OutcomeMeasurement) *OutcomeCategoryUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.RemoveOutcomeMeasurementListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OutcomeCategoryUpdateOne) Select(field string, fields ...string) *OutcomeCategoryUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OutcomeCategory entity.
func (ocuo *OutcomeCategoryUpdateOne) Save(ctx context.Context) (*OutcomeCategory, error) {
	var (
		err  error
		node *OutcomeCategory
	)
	if len(ocuo.hooks) == 0 {
		node, err = ocuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocuo.mutation = mutation
			node, err = ocuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ocuo.hooks) - 1; i >= 0; i-- {
			if ocuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocuo *OutcomeCategoryUpdateOne) SaveX(ctx context.Context) *OutcomeCategory {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OutcomeCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OutcomeCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocuo *OutcomeCategoryUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomecategory.Table,
			Columns: outcomecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomecategory.FieldID,
			},
		},
	}
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomecategory.FieldID)
		for _, f := range fields {
			if !outcomecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocuo.mutation.OutcomeCategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomecategory.FieldOutcomeCategoryTitle,
		})
	}
	if ocuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomecategory.ParentTable,
			Columns: []string{outcomecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomecategory.ParentTable,
			Columns: []string{outcomecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocuo.mutation.OutcomeMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.RemovedOutcomeMeasurementListIDs(); len(nodes) > 0 && !ocuo.mutation.OutcomeMeasurementListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.OutcomeMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomecategory.OutcomeMeasurementListTable,
			Columns: []string{outcomecategory.OutcomeMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeCategory{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
