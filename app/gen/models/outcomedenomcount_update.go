// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomCountUpdate is the builder for updating OutcomeDenomCount entities.
type OutcomeDenomCountUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeDenomCountMutation
}

// Where appends a list predicates to the OutcomeDenomCountUpdate builder.
func (odcu *OutcomeDenomCountUpdate) Where(ps ...predicate.OutcomeDenomCount) *OutcomeDenomCountUpdate {
	odcu.mutation.Where(ps...)
	return odcu
}

// SetOutcomeDenomCountGroupID sets the "outcome_denom_count_group_id" field.
func (odcu *OutcomeDenomCountUpdate) SetOutcomeDenomCountGroupID(s string) *OutcomeDenomCountUpdate {
	odcu.mutation.SetOutcomeDenomCountGroupID(s)
	return odcu
}

// SetOutcomeDenomCountValue sets the "outcome_denom_count_value" field.
func (odcu *OutcomeDenomCountUpdate) SetOutcomeDenomCountValue(s string) *OutcomeDenomCountUpdate {
	odcu.mutation.SetOutcomeDenomCountValue(s)
	return odcu
}

// SetParentID sets the "parent" edge to the OutcomeDenom entity by ID.
func (odcu *OutcomeDenomCountUpdate) SetParentID(id int) *OutcomeDenomCountUpdate {
	odcu.mutation.SetParentID(id)
	return odcu
}

// SetNillableParentID sets the "parent" edge to the OutcomeDenom entity by ID if the given value is not nil.
func (odcu *OutcomeDenomCountUpdate) SetNillableParentID(id *int) *OutcomeDenomCountUpdate {
	if id != nil {
		odcu = odcu.SetParentID(*id)
	}
	return odcu
}

// SetParent sets the "parent" edge to the OutcomeDenom entity.
func (odcu *OutcomeDenomCountUpdate) SetParent(o *OutcomeDenom) *OutcomeDenomCountUpdate {
	return odcu.SetParentID(o.ID)
}

// Mutation returns the OutcomeDenomCountMutation object of the builder.
func (odcu *OutcomeDenomCountUpdate) Mutation() *OutcomeDenomCountMutation {
	return odcu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeDenom entity.
func (odcu *OutcomeDenomCountUpdate) ClearParent() *OutcomeDenomCountUpdate {
	odcu.mutation.ClearParent()
	return odcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odcu *OutcomeDenomCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odcu.hooks) == 0 {
		affected, err = odcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odcu.mutation = mutation
			affected, err = odcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odcu.hooks) - 1; i >= 0; i-- {
			if odcu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (odcu *OutcomeDenomCountUpdate) SaveX(ctx context.Context) int {
	affected, err := odcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odcu *OutcomeDenomCountUpdate) Exec(ctx context.Context) error {
	_, err := odcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcu *OutcomeDenomCountUpdate) ExecX(ctx context.Context) {
	if err := odcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (odcu *OutcomeDenomCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenomcount.Table,
			Columns: outcomedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenomcount.FieldID,
			},
		},
	}
	if ps := odcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odcu.mutation.OutcomeDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountGroupID,
		})
	}
	if value, ok := odcu.mutation.OutcomeDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountValue,
		})
	}
	if odcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenomcount.ParentTable,
			Columns: []string{outcomedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenomcount.ParentTable,
			Columns: []string{outcomedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, odcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomedenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeDenomCountUpdateOne is the builder for updating a single OutcomeDenomCount entity.
type OutcomeDenomCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeDenomCountMutation
}

// SetOutcomeDenomCountGroupID sets the "outcome_denom_count_group_id" field.
func (odcuo *OutcomeDenomCountUpdateOne) SetOutcomeDenomCountGroupID(s string) *OutcomeDenomCountUpdateOne {
	odcuo.mutation.SetOutcomeDenomCountGroupID(s)
	return odcuo
}

// SetOutcomeDenomCountValue sets the "outcome_denom_count_value" field.
func (odcuo *OutcomeDenomCountUpdateOne) SetOutcomeDenomCountValue(s string) *OutcomeDenomCountUpdateOne {
	odcuo.mutation.SetOutcomeDenomCountValue(s)
	return odcuo
}

// SetParentID sets the "parent" edge to the OutcomeDenom entity by ID.
func (odcuo *OutcomeDenomCountUpdateOne) SetParentID(id int) *OutcomeDenomCountUpdateOne {
	odcuo.mutation.SetParentID(id)
	return odcuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeDenom entity by ID if the given value is not nil.
func (odcuo *OutcomeDenomCountUpdateOne) SetNillableParentID(id *int) *OutcomeDenomCountUpdateOne {
	if id != nil {
		odcuo = odcuo.SetParentID(*id)
	}
	return odcuo
}

// SetParent sets the "parent" edge to the OutcomeDenom entity.
func (odcuo *OutcomeDenomCountUpdateOne) SetParent(o *OutcomeDenom) *OutcomeDenomCountUpdateOne {
	return odcuo.SetParentID(o.ID)
}

// Mutation returns the OutcomeDenomCountMutation object of the builder.
func (odcuo *OutcomeDenomCountUpdateOne) Mutation() *OutcomeDenomCountMutation {
	return odcuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeDenom entity.
func (odcuo *OutcomeDenomCountUpdateOne) ClearParent() *OutcomeDenomCountUpdateOne {
	odcuo.mutation.ClearParent()
	return odcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (odcuo *OutcomeDenomCountUpdateOne) Select(field string, fields ...string) *OutcomeDenomCountUpdateOne {
	odcuo.fields = append([]string{field}, fields...)
	return odcuo
}

// Save executes the query and returns the updated OutcomeDenomCount entity.
func (odcuo *OutcomeDenomCountUpdateOne) Save(ctx context.Context) (*OutcomeDenomCount, error) {
	var (
		err  error
		node *OutcomeDenomCount
	)
	if len(odcuo.hooks) == 0 {
		node, err = odcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odcuo.mutation = mutation
			node, err = odcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(odcuo.hooks) - 1; i >= 0; i-- {
			if odcuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (odcuo *OutcomeDenomCountUpdateOne) SaveX(ctx context.Context) *OutcomeDenomCount {
	node, err := odcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (odcuo *OutcomeDenomCountUpdateOne) Exec(ctx context.Context) error {
	_, err := odcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcuo *OutcomeDenomCountUpdateOne) ExecX(ctx context.Context) {
	if err := odcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (odcuo *OutcomeDenomCountUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeDenomCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenomcount.Table,
			Columns: outcomedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenomcount.FieldID,
			},
		},
	}
	id, ok := odcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeDenomCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := odcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenomcount.FieldID)
		for _, f := range fields {
			if !outcomedenomcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomedenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := odcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odcuo.mutation.OutcomeDenomCountGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountGroupID,
		})
	}
	if value, ok := odcuo.mutation.OutcomeDenomCountValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenomcount.FieldOutcomeDenomCountValue,
		})
	}
	if odcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenomcount.ParentTable,
			Columns: []string{outcomedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenomcount.ParentTable,
			Columns: []string{outcomedenomcount.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeDenomCount{config: odcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, odcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomedenomcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
