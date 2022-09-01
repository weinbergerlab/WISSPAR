// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomUpdate is the builder for updating OutcomeClassDenom entities.
type OutcomeClassDenomUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeClassDenomMutation
}

// Where appends a list predicates to the OutcomeClassDenomUpdate builder.
func (ocdu *OutcomeClassDenomUpdate) Where(ps ...predicate.OutcomeClassDenom) *OutcomeClassDenomUpdate {
	ocdu.mutation.Where(ps...)
	return ocdu
}

// SetOutcomeClassDenomUnits sets the "outcome_class_denom_units" field.
func (ocdu *OutcomeClassDenomUpdate) SetOutcomeClassDenomUnits(s string) *OutcomeClassDenomUpdate {
	ocdu.mutation.SetOutcomeClassDenomUnits(s)
	return ocdu
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (ocdu *OutcomeClassDenomUpdate) SetParentID(id int) *OutcomeClassDenomUpdate {
	ocdu.mutation.SetParentID(id)
	return ocdu
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (ocdu *OutcomeClassDenomUpdate) SetNillableParentID(id *int) *OutcomeClassDenomUpdate {
	if id != nil {
		ocdu = ocdu.SetParentID(*id)
	}
	return ocdu
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (ocdu *OutcomeClassDenomUpdate) SetParent(o *OutcomeClass) *OutcomeClassDenomUpdate {
	return ocdu.SetParentID(o.ID)
}

// AddOutcomeClassDenomCountListIDs adds the "outcome_class_denom_count_list" edge to the OutcomeClassDenomCount entity by IDs.
func (ocdu *OutcomeClassDenomUpdate) AddOutcomeClassDenomCountListIDs(ids ...int) *OutcomeClassDenomUpdate {
	ocdu.mutation.AddOutcomeClassDenomCountListIDs(ids...)
	return ocdu
}

// AddOutcomeClassDenomCountList adds the "outcome_class_denom_count_list" edges to the OutcomeClassDenomCount entity.
func (ocdu *OutcomeClassDenomUpdate) AddOutcomeClassDenomCountList(o ...*OutcomeClassDenomCount) *OutcomeClassDenomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocdu.AddOutcomeClassDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeClassDenomMutation object of the builder.
func (ocdu *OutcomeClassDenomUpdate) Mutation() *OutcomeClassDenomMutation {
	return ocdu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClass entity.
func (ocdu *OutcomeClassDenomUpdate) ClearParent() *OutcomeClassDenomUpdate {
	ocdu.mutation.ClearParent()
	return ocdu
}

// ClearOutcomeClassDenomCountList clears all "outcome_class_denom_count_list" edges to the OutcomeClassDenomCount entity.
func (ocdu *OutcomeClassDenomUpdate) ClearOutcomeClassDenomCountList() *OutcomeClassDenomUpdate {
	ocdu.mutation.ClearOutcomeClassDenomCountList()
	return ocdu
}

// RemoveOutcomeClassDenomCountListIDs removes the "outcome_class_denom_count_list" edge to OutcomeClassDenomCount entities by IDs.
func (ocdu *OutcomeClassDenomUpdate) RemoveOutcomeClassDenomCountListIDs(ids ...int) *OutcomeClassDenomUpdate {
	ocdu.mutation.RemoveOutcomeClassDenomCountListIDs(ids...)
	return ocdu
}

// RemoveOutcomeClassDenomCountList removes "outcome_class_denom_count_list" edges to OutcomeClassDenomCount entities.
func (ocdu *OutcomeClassDenomUpdate) RemoveOutcomeClassDenomCountList(o ...*OutcomeClassDenomCount) *OutcomeClassDenomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocdu.RemoveOutcomeClassDenomCountListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocdu *OutcomeClassDenomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocdu.hooks) == 0 {
		affected, err = ocdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocdu.mutation = mutation
			affected, err = ocdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocdu.hooks) - 1; i >= 0; i-- {
			if ocdu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocdu *OutcomeClassDenomUpdate) SaveX(ctx context.Context) int {
	affected, err := ocdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocdu *OutcomeClassDenomUpdate) Exec(ctx context.Context) error {
	_, err := ocdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdu *OutcomeClassDenomUpdate) ExecX(ctx context.Context) {
	if err := ocdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocdu *OutcomeClassDenomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenom.Table,
			Columns: outcomeclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenom.FieldID,
			},
		},
	}
	if ps := ocdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocdu.mutation.OutcomeClassDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenom.FieldOutcomeClassDenomUnits,
		})
	}
	if ocdu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenom.ParentTable,
			Columns: []string{outcomeclassdenom.ParentColumn},
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
	if nodes := ocdu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenom.ParentTable,
			Columns: []string{outcomeclassdenom.ParentColumn},
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
	if ocdu.mutation.OutcomeClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocdu.mutation.RemovedOutcomeClassDenomCountListIDs(); len(nodes) > 0 && !ocdu.mutation.OutcomeClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocdu.mutation.OutcomeClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ocdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeclassdenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeClassDenomUpdateOne is the builder for updating a single OutcomeClassDenom entity.
type OutcomeClassDenomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeClassDenomMutation
}

// SetOutcomeClassDenomUnits sets the "outcome_class_denom_units" field.
func (ocduo *OutcomeClassDenomUpdateOne) SetOutcomeClassDenomUnits(s string) *OutcomeClassDenomUpdateOne {
	ocduo.mutation.SetOutcomeClassDenomUnits(s)
	return ocduo
}

// SetParentID sets the "parent" edge to the OutcomeClass entity by ID.
func (ocduo *OutcomeClassDenomUpdateOne) SetParentID(id int) *OutcomeClassDenomUpdateOne {
	ocduo.mutation.SetParentID(id)
	return ocduo
}

// SetNillableParentID sets the "parent" edge to the OutcomeClass entity by ID if the given value is not nil.
func (ocduo *OutcomeClassDenomUpdateOne) SetNillableParentID(id *int) *OutcomeClassDenomUpdateOne {
	if id != nil {
		ocduo = ocduo.SetParentID(*id)
	}
	return ocduo
}

// SetParent sets the "parent" edge to the OutcomeClass entity.
func (ocduo *OutcomeClassDenomUpdateOne) SetParent(o *OutcomeClass) *OutcomeClassDenomUpdateOne {
	return ocduo.SetParentID(o.ID)
}

// AddOutcomeClassDenomCountListIDs adds the "outcome_class_denom_count_list" edge to the OutcomeClassDenomCount entity by IDs.
func (ocduo *OutcomeClassDenomUpdateOne) AddOutcomeClassDenomCountListIDs(ids ...int) *OutcomeClassDenomUpdateOne {
	ocduo.mutation.AddOutcomeClassDenomCountListIDs(ids...)
	return ocduo
}

// AddOutcomeClassDenomCountList adds the "outcome_class_denom_count_list" edges to the OutcomeClassDenomCount entity.
func (ocduo *OutcomeClassDenomUpdateOne) AddOutcomeClassDenomCountList(o ...*OutcomeClassDenomCount) *OutcomeClassDenomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocduo.AddOutcomeClassDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeClassDenomMutation object of the builder.
func (ocduo *OutcomeClassDenomUpdateOne) Mutation() *OutcomeClassDenomMutation {
	return ocduo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClass entity.
func (ocduo *OutcomeClassDenomUpdateOne) ClearParent() *OutcomeClassDenomUpdateOne {
	ocduo.mutation.ClearParent()
	return ocduo
}

// ClearOutcomeClassDenomCountList clears all "outcome_class_denom_count_list" edges to the OutcomeClassDenomCount entity.
func (ocduo *OutcomeClassDenomUpdateOne) ClearOutcomeClassDenomCountList() *OutcomeClassDenomUpdateOne {
	ocduo.mutation.ClearOutcomeClassDenomCountList()
	return ocduo
}

// RemoveOutcomeClassDenomCountListIDs removes the "outcome_class_denom_count_list" edge to OutcomeClassDenomCount entities by IDs.
func (ocduo *OutcomeClassDenomUpdateOne) RemoveOutcomeClassDenomCountListIDs(ids ...int) *OutcomeClassDenomUpdateOne {
	ocduo.mutation.RemoveOutcomeClassDenomCountListIDs(ids...)
	return ocduo
}

// RemoveOutcomeClassDenomCountList removes "outcome_class_denom_count_list" edges to OutcomeClassDenomCount entities.
func (ocduo *OutcomeClassDenomUpdateOne) RemoveOutcomeClassDenomCountList(o ...*OutcomeClassDenomCount) *OutcomeClassDenomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ocduo.RemoveOutcomeClassDenomCountListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocduo *OutcomeClassDenomUpdateOne) Select(field string, fields ...string) *OutcomeClassDenomUpdateOne {
	ocduo.fields = append([]string{field}, fields...)
	return ocduo
}

// Save executes the query and returns the updated OutcomeClassDenom entity.
func (ocduo *OutcomeClassDenomUpdateOne) Save(ctx context.Context) (*OutcomeClassDenom, error) {
	var (
		err  error
		node *OutcomeClassDenom
	)
	if len(ocduo.hooks) == 0 {
		node, err = ocduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocduo.mutation = mutation
			node, err = ocduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ocduo.hooks) - 1; i >= 0; i-- {
			if ocduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocduo *OutcomeClassDenomUpdateOne) SaveX(ctx context.Context) *OutcomeClassDenom {
	node, err := ocduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocduo *OutcomeClassDenomUpdateOne) Exec(ctx context.Context) error {
	_, err := ocduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocduo *OutcomeClassDenomUpdateOne) ExecX(ctx context.Context) {
	if err := ocduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocduo *OutcomeClassDenomUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeClassDenom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenom.Table,
			Columns: outcomeclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenom.FieldID,
			},
		},
	}
	id, ok := ocduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeClassDenom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenom.FieldID)
		for _, f := range fields {
			if !outcomeclassdenom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeclassdenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocduo.mutation.OutcomeClassDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenom.FieldOutcomeClassDenomUnits,
		})
	}
	if ocduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenom.ParentTable,
			Columns: []string{outcomeclassdenom.ParentColumn},
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
	if nodes := ocduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenom.ParentTable,
			Columns: []string{outcomeclassdenom.ParentColumn},
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
	if ocduo.mutation.OutcomeClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocduo.mutation.RemovedOutcomeClassDenomCountListIDs(); len(nodes) > 0 && !ocduo.mutation.OutcomeClassDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocduo.mutation.OutcomeClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeclassdenom.OutcomeClassDenomCountListTable,
			Columns: []string{outcomeclassdenom.OutcomeClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclassdenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeClassDenom{config: ocduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeclassdenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
