// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomCountUpdate is the builder for updating OutcomeClassDenomCount entities.
type OutcomeClassDenomCountUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeClassDenomCountMutation
}

// Where appends a list predicates to the OutcomeClassDenomCountUpdate builder.
func (ocdcu *OutcomeClassDenomCountUpdate) Where(ps ...predicate.OutcomeClassDenomCount) *OutcomeClassDenomCountUpdate {
	ocdcu.mutation.Where(ps...)
	return ocdcu
}

// SetOutcomeClassDenomCountGroupID sets the "outcome_class_denom_count_group_id" field.
func (ocdcu *OutcomeClassDenomCountUpdate) SetOutcomeClassDenomCountGroupID(s string) *OutcomeClassDenomCountUpdate {
	ocdcu.mutation.SetOutcomeClassDenomCountGroupID(s)
	return ocdcu
}

// SetOutcomeClassDenomCountValue sets the "outcome_class_denom_count_value" field.
func (ocdcu *OutcomeClassDenomCountUpdate) SetOutcomeClassDenomCountValue(s string) *OutcomeClassDenomCountUpdate {
	ocdcu.mutation.SetOutcomeClassDenomCountValue(s)
	return ocdcu
}

// SetParentID sets the "parent" edge to the OutcomeClassDenom entity by ID.
func (ocdcu *OutcomeClassDenomCountUpdate) SetParentID(id int) *OutcomeClassDenomCountUpdate {
	ocdcu.mutation.SetParentID(id)
	return ocdcu
}

// SetNillableParentID sets the "parent" edge to the OutcomeClassDenom entity by ID if the given value is not nil.
func (ocdcu *OutcomeClassDenomCountUpdate) SetNillableParentID(id *int) *OutcomeClassDenomCountUpdate {
	if id != nil {
		ocdcu = ocdcu.SetParentID(*id)
	}
	return ocdcu
}

// SetParent sets the "parent" edge to the OutcomeClassDenom entity.
func (ocdcu *OutcomeClassDenomCountUpdate) SetParent(o *OutcomeClassDenom) *OutcomeClassDenomCountUpdate {
	return ocdcu.SetParentID(o.ID)
}

// Mutation returns the OutcomeClassDenomCountMutation object of the builder.
func (ocdcu *OutcomeClassDenomCountUpdate) Mutation() *OutcomeClassDenomCountMutation {
	return ocdcu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClassDenom entity.
func (ocdcu *OutcomeClassDenomCountUpdate) ClearParent() *OutcomeClassDenomCountUpdate {
	ocdcu.mutation.ClearParent()
	return ocdcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocdcu *OutcomeClassDenomCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ocdcu.hooks) == 0 {
		affected, err = ocdcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocdcu.mutation = mutation
			affected, err = ocdcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocdcu.hooks) - 1; i >= 0; i-- {
			if ocdcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocdcu *OutcomeClassDenomCountUpdate) SaveX(ctx context.Context) int {
	affected, err := ocdcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocdcu *OutcomeClassDenomCountUpdate) Exec(ctx context.Context) error {
	_, err := ocdcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcu *OutcomeClassDenomCountUpdate) ExecX(ctx context.Context) {
	if err := ocdcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocdcu *OutcomeClassDenomCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenomcount.Table,
			Columns: outcomeclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenomcount.FieldID,
			},
		},
	}
	if ps := ocdcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocdcu.mutation.OutcomeClassDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID,
		})
	}
	if value, ok := ocdcu.mutation.OutcomeClassDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountValue,
		})
	}
	if ocdcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenomcount.ParentTable,
			Columns: []string{outcomeclassdenomcount.ParentColumn},
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
	if nodes := ocdcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenomcount.ParentTable,
			Columns: []string{outcomeclassdenomcount.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ocdcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeclassdenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeClassDenomCountUpdateOne is the builder for updating a single OutcomeClassDenomCount entity.
type OutcomeClassDenomCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeClassDenomCountMutation
}

// SetOutcomeClassDenomCountGroupID sets the "outcome_class_denom_count_group_id" field.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SetOutcomeClassDenomCountGroupID(s string) *OutcomeClassDenomCountUpdateOne {
	ocdcuo.mutation.SetOutcomeClassDenomCountGroupID(s)
	return ocdcuo
}

// SetOutcomeClassDenomCountValue sets the "outcome_class_denom_count_value" field.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SetOutcomeClassDenomCountValue(s string) *OutcomeClassDenomCountUpdateOne {
	ocdcuo.mutation.SetOutcomeClassDenomCountValue(s)
	return ocdcuo
}

// SetParentID sets the "parent" edge to the OutcomeClassDenom entity by ID.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SetParentID(id int) *OutcomeClassDenomCountUpdateOne {
	ocdcuo.mutation.SetParentID(id)
	return ocdcuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeClassDenom entity by ID if the given value is not nil.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SetNillableParentID(id *int) *OutcomeClassDenomCountUpdateOne {
	if id != nil {
		ocdcuo = ocdcuo.SetParentID(*id)
	}
	return ocdcuo
}

// SetParent sets the "parent" edge to the OutcomeClassDenom entity.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SetParent(o *OutcomeClassDenom) *OutcomeClassDenomCountUpdateOne {
	return ocdcuo.SetParentID(o.ID)
}

// Mutation returns the OutcomeClassDenomCountMutation object of the builder.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) Mutation() *OutcomeClassDenomCountMutation {
	return ocdcuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeClassDenom entity.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) ClearParent() *OutcomeClassDenomCountUpdateOne {
	ocdcuo.mutation.ClearParent()
	return ocdcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) Select(field string, fields ...string) *OutcomeClassDenomCountUpdateOne {
	ocdcuo.fields = append([]string{field}, fields...)
	return ocdcuo
}

// Save executes the query and returns the updated OutcomeClassDenomCount entity.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) Save(ctx context.Context) (*OutcomeClassDenomCount, error) {
	var (
		err  error
		node *OutcomeClassDenomCount
	)
	if len(ocdcuo.hooks) == 0 {
		node, err = ocdcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeClassDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocdcuo.mutation = mutation
			node, err = ocdcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ocdcuo.hooks) - 1; i >= 0; i-- {
			if ocdcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ocdcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocdcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) SaveX(ctx context.Context) *OutcomeClassDenomCount {
	node, err := ocdcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) Exec(ctx context.Context) error {
	_, err := ocdcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdcuo *OutcomeClassDenomCountUpdateOne) ExecX(ctx context.Context) {
	if err := ocdcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocdcuo *OutcomeClassDenomCountUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeClassDenomCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenomcount.Table,
			Columns: outcomeclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenomcount.FieldID,
			},
		},
	}
	id, ok := ocdcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeClassDenomCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocdcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenomcount.FieldID)
		for _, f := range fields {
			if !outcomeclassdenomcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeclassdenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocdcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocdcuo.mutation.OutcomeClassDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID,
		})
	}
	if value, ok := ocdcuo.mutation.OutcomeClassDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeclassdenomcount.FieldOutcomeClassDenomCountValue,
		})
	}
	if ocdcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenomcount.ParentTable,
			Columns: []string{outcomeclassdenomcount.ParentColumn},
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
	if nodes := ocdcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeclassdenomcount.ParentTable,
			Columns: []string{outcomeclassdenomcount.ParentColumn},
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
	_node = &OutcomeClassDenomCount{config: ocdcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocdcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeclassdenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
