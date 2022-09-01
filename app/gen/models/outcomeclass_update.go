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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassUpdate is the builder for updating OutcomeClass entities.
type OutcomeClassUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeClassMutation
}

// Where appends a list predicates to the OutcomeClassUpdate builder.
func (ocu *OutcomeClassUpdate) Where(ps ...predicate.OutcomeClass) *OutcomeClassUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetOutcomeClassTitle sets the "outcome_class_title" field.
func (ocu *OutcomeClassUpdate) SetOutcomeClassTitle(s string) *OutcomeClassUpdate {
	ocu.mutation.SetOutcomeClassTitle(s)
	return ocu
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (ocu *OutcomeClassUpdate) SetParentID(id int) *OutcomeClassUpdate {
	ocu.mutation.SetParentID(id)
	return ocu
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (ocu *OutcomeClassUpdate) SetNillableParentID(id *int) *OutcomeClassUpdate {
	if id != nil {
		ocu = ocu.SetParentID(*id)
	}
	return ocu
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (ocu *OutcomeClassUpdate) SetParent(o *OutcomeMeasure) *OutcomeClassUpdate {
	return ocu.SetParentID(o.ID)
}

// AddOutcomeClassDenomListIDs adds the "outcome_class_denom_list" edge to the OutcomeClassDenom entity by IDs.
func (ocu *OutcomeClassUpdate) AddOutcomeClassDenomListIDs(ids ...int) *OutcomeClassUpdate {
	ocu.mutation.AddOutcomeClassDenomListIDs(ids...)
	return ocu
}

// AddOutcomeClassDenomList adds the "outcome_class_denom_list" edges to the OutcomeClassDenom entity.
func (ocu *OutcomeClassUpdate) AddOutcomeClassDenomList(o ...*OutcomeClassDenom) *OutcomeClassUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.AddOutcomeClassDenomListIDs(ids...)
}

// AddOutcomeCategoryListIDs adds the "outcome_category_list" edge to the OutcomeCategory entity by IDs.
func (ocu *OutcomeClassUpdate) AddOutcomeCategoryListIDs(ids ...int) *OutcomeClassUpdate {
	ocu.mutation.AddOutcomeCategoryListIDs(ids...)
	return ocu
}

// AddOutcomeCategoryList adds the "outcome_category_list" edges to the OutcomeCategory entity.
func (ocu *OutcomeClassUpdate) AddOutcomeCategoryList(o ...*OutcomeCategory) *OutcomeClassUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.AddOutcomeCategoryListIDs(ids...)
}

// Mutation returns the OutcomeClassMutation object of the builder.
func (ocu *OutcomeClassUpdate) Mutation() *OutcomeClassMutation {
	return ocu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (ocu *OutcomeClassUpdate) ClearParent() *OutcomeClassUpdate {
	ocu.mutation.ClearParent()
	return ocu
}

// ClearOutcomeClassDenomList clears all "outcome_class_denom_list" edges to the OutcomeClassDenom entity.
func (ocu *OutcomeClassUpdate) ClearOutcomeClassDenomList() *OutcomeClassUpdate {
	ocu.mutation.ClearOutcomeClassDenomList()
	return ocu
}

// RemoveOutcomeClassDenomListIDs removes the "outcome_class_denom_list" edge to OutcomeClassDenom entities by IDs.
func (ocu *OutcomeClassUpdate) RemoveOutcomeClassDenomListIDs(ids ...int) *OutcomeClassUpdate {
	ocu.mutation.RemoveOutcomeClassDenomListIDs(ids...)
	return ocu
}

// RemoveOutcomeClassDenomList removes "outcome_class_denom_list" edges to OutcomeClassDenom entities.
func (ocu *OutcomeClassUpdate) RemoveOutcomeClassDenomList(o ...*OutcomeClassDenom) *OutcomeClassUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.RemoveOutcomeClassDenomListIDs(ids...)
}

// ClearOutcomeCategoryList clears all "outcome_category_list" edges to the OutcomeCategory entity.
func (ocu *OutcomeClassUpdate) ClearOutcomeCategoryList() *OutcomeClassUpdate {
	ocu.mutation.ClearOutcomeCategoryList()
	return ocu
}

// RemoveOutcomeCategoryListIDs removes the "outcome_category_list" edge to OutcomeCategory entities by IDs.
func (ocu *OutcomeClassUpdate) RemoveOutcomeCategoryListIDs(ids ...int) *OutcomeClassUpdate {
	ocu.mutation.RemoveOutcomeCategoryListIDs(ids...)
	return ocu
}

// RemoveOutcomeCategoryList removes "outcome_category_list" edges to OutcomeCategory entities.
func (ocu *OutcomeClassUpdate) RemoveOutcomeCategoryList(o ...*OutcomeCategory) *OutcomeClassUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocu.RemoveOutcomeCategoryListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OutcomeClassUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocu.hooks) == 0 {
		affected, err = ocu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassMutation)
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
func (ocu *OutcomeClassUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OutcomeClassUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OutcomeClassUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocu *OutcomeClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclass.Table,
			Columns: outcomeclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclass.FieldID,
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
	if value, ok := ocu.mutation.OutcomeClassTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclass.FieldOutcomeClassTitle,
		})
	}
	if ocu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclass.ParentTable,
			Columns: []string{outcomeclass.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclass.ParentTable,
			Columns: []string{outcomeclass.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocu.mutation.OutcomeClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.RemovedOutcomeClassDenomListIDs(); len(nodes) > 0 && !ocu.mutation.OutcomeClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.OutcomeClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocu.mutation.OutcomeCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.RemovedOutcomeCategoryListIDs(); len(nodes) > 0 && !ocu.mutation.OutcomeCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.OutcomeCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
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
			err = &NotFoundError{outcomeclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeClassUpdateOne is the builder for updating a single OutcomeClass entity.
type OutcomeClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeClassMutation
}

// SetOutcomeClassTitle sets the "outcome_class_title" field.
func (ocuo *OutcomeClassUpdateOne) SetOutcomeClassTitle(s string) *OutcomeClassUpdateOne {
	ocuo.mutation.SetOutcomeClassTitle(s)
	return ocuo
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (ocuo *OutcomeClassUpdateOne) SetParentID(id int) *OutcomeClassUpdateOne {
	ocuo.mutation.SetParentID(id)
	return ocuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (ocuo *OutcomeClassUpdateOne) SetNillableParentID(id *int) *OutcomeClassUpdateOne {
	if id != nil {
		ocuo = ocuo.SetParentID(*id)
	}
	return ocuo
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (ocuo *OutcomeClassUpdateOne) SetParent(o *OutcomeMeasure) *OutcomeClassUpdateOne {
	return ocuo.SetParentID(o.ID)
}

// AddOutcomeClassDenomListIDs adds the "outcome_class_denom_list" edge to the OutcomeClassDenom entity by IDs.
func (ocuo *OutcomeClassUpdateOne) AddOutcomeClassDenomListIDs(ids ...int) *OutcomeClassUpdateOne {
	ocuo.mutation.AddOutcomeClassDenomListIDs(ids...)
	return ocuo
}

// AddOutcomeClassDenomList adds the "outcome_class_denom_list" edges to the OutcomeClassDenom entity.
func (ocuo *OutcomeClassUpdateOne) AddOutcomeClassDenomList(o ...*OutcomeClassDenom) *OutcomeClassUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.AddOutcomeClassDenomListIDs(ids...)
}

// AddOutcomeCategoryListIDs adds the "outcome_category_list" edge to the OutcomeCategory entity by IDs.
func (ocuo *OutcomeClassUpdateOne) AddOutcomeCategoryListIDs(ids ...int) *OutcomeClassUpdateOne {
	ocuo.mutation.AddOutcomeCategoryListIDs(ids...)
	return ocuo
}

// AddOutcomeCategoryList adds the "outcome_category_list" edges to the OutcomeCategory entity.
func (ocuo *OutcomeClassUpdateOne) AddOutcomeCategoryList(o ...*OutcomeCategory) *OutcomeClassUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.AddOutcomeCategoryListIDs(ids...)
}

// Mutation returns the OutcomeClassMutation object of the builder.
func (ocuo *OutcomeClassUpdateOne) Mutation() *OutcomeClassMutation {
	return ocuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (ocuo *OutcomeClassUpdateOne) ClearParent() *OutcomeClassUpdateOne {
	ocuo.mutation.ClearParent()
	return ocuo
}

// ClearOutcomeClassDenomList clears all "outcome_class_denom_list" edges to the OutcomeClassDenom entity.
func (ocuo *OutcomeClassUpdateOne) ClearOutcomeClassDenomList() *OutcomeClassUpdateOne {
	ocuo.mutation.ClearOutcomeClassDenomList()
	return ocuo
}

// RemoveOutcomeClassDenomListIDs removes the "outcome_class_denom_list" edge to OutcomeClassDenom entities by IDs.
func (ocuo *OutcomeClassUpdateOne) RemoveOutcomeClassDenomListIDs(ids ...int) *OutcomeClassUpdateOne {
	ocuo.mutation.RemoveOutcomeClassDenomListIDs(ids...)
	return ocuo
}

// RemoveOutcomeClassDenomList removes "outcome_class_denom_list" edges to OutcomeClassDenom entities.
func (ocuo *OutcomeClassUpdateOne) RemoveOutcomeClassDenomList(o ...*OutcomeClassDenom) *OutcomeClassUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.RemoveOutcomeClassDenomListIDs(ids...)
}

// ClearOutcomeCategoryList clears all "outcome_category_list" edges to the OutcomeCategory entity.
func (ocuo *OutcomeClassUpdateOne) ClearOutcomeCategoryList() *OutcomeClassUpdateOne {
	ocuo.mutation.ClearOutcomeCategoryList()
	return ocuo
}

// RemoveOutcomeCategoryListIDs removes the "outcome_category_list" edge to OutcomeCategory entities by IDs.
func (ocuo *OutcomeClassUpdateOne) RemoveOutcomeCategoryListIDs(ids ...int) *OutcomeClassUpdateOne {
	ocuo.mutation.RemoveOutcomeCategoryListIDs(ids...)
	return ocuo
}

// RemoveOutcomeCategoryList removes "outcome_category_list" edges to OutcomeCategory entities.
func (ocuo *OutcomeClassUpdateOne) RemoveOutcomeCategoryList(o ...*OutcomeCategory) *OutcomeClassUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocuo.RemoveOutcomeCategoryListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OutcomeClassUpdateOne) Select(field string, fields ...string) *OutcomeClassUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OutcomeClass entity.
func (ocuo *OutcomeClassUpdateOne) Save(ctx context.Context) (*OutcomeClass, error) {
	var (
		err  error
		node *OutcomeClass
	)
	if len(ocuo.hooks) == 0 {
		node, err = ocuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassMutation)
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
func (ocuo *OutcomeClassUpdateOne) SaveX(ctx context.Context) *OutcomeClass {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OutcomeClassUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OutcomeClassUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocuo *OutcomeClassUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeClass, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclass.Table,
			Columns: outcomeclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclass.FieldID,
			},
		},
	}
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeClass.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclass.FieldID)
		for _, f := range fields {
			if !outcomeclass.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeclass.FieldID {
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
	if value, ok := ocuo.mutation.OutcomeClassTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclass.FieldOutcomeClassTitle,
		})
	}
	if ocuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclass.ParentTable,
			Columns: []string{outcomeclass.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclass.ParentTable,
			Columns: []string{outcomeclass.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocuo.mutation.OutcomeClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.RemovedOutcomeClassDenomListIDs(); len(nodes) > 0 && !ocuo.mutation.OutcomeClassDenomListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.OutcomeClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeClassDenomListTable,
			Columns: []string{outcomeclass.OutcomeClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ocuo.mutation.OutcomeCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.RemovedOutcomeCategoryListIDs(); len(nodes) > 0 && !ocuo.mutation.OutcomeCategoryListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.OutcomeCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclass.OutcomeCategoryListTable,
			Columns: []string{outcomeclass.OutcomeCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomecategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeClass{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
