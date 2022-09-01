// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeGroupUpdate is the builder for updating OutcomeGroup entities.
type OutcomeGroupUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeGroupMutation
}

// Where appends a list predicates to the OutcomeGroupUpdate builder.
func (ogu *OutcomeGroupUpdate) Where(ps ...predicate.OutcomeGroup) *OutcomeGroupUpdate {
	ogu.mutation.Where(ps...)
	return ogu
}

// SetOutcomeGroupID sets the "outcome_group_id" field.
func (ogu *OutcomeGroupUpdate) SetOutcomeGroupID(s string) *OutcomeGroupUpdate {
	ogu.mutation.SetOutcomeGroupID(s)
	return ogu
}

// SetOutcomeGroupTitle sets the "outcome_group_title" field.
func (ogu *OutcomeGroupUpdate) SetOutcomeGroupTitle(s string) *OutcomeGroupUpdate {
	ogu.mutation.SetOutcomeGroupTitle(s)
	return ogu
}

// SetOutcomeGroupDescription sets the "outcome_group_description" field.
func (ogu *OutcomeGroupUpdate) SetOutcomeGroupDescription(s string) *OutcomeGroupUpdate {
	ogu.mutation.SetOutcomeGroupDescription(s)
	return ogu
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (ogu *OutcomeGroupUpdate) SetParentID(id int) *OutcomeGroupUpdate {
	ogu.mutation.SetParentID(id)
	return ogu
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (ogu *OutcomeGroupUpdate) SetNillableParentID(id *int) *OutcomeGroupUpdate {
	if id != nil {
		ogu = ogu.SetParentID(*id)
	}
	return ogu
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (ogu *OutcomeGroupUpdate) SetParent(o *OutcomeMeasure) *OutcomeGroupUpdate {
	return ogu.SetParentID(o.ID)
}

// Mutation returns the OutcomeGroupMutation object of the builder.
func (ogu *OutcomeGroupUpdate) Mutation() *OutcomeGroupMutation {
	return ogu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (ogu *OutcomeGroupUpdate) ClearParent() *OutcomeGroupUpdate {
	ogu.mutation.ClearParent()
	return ogu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ogu *OutcomeGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ogu.hooks) == 0 {
		affected, err = ogu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ogu.mutation = mutation
			affected, err = ogu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ogu.hooks) - 1; i >= 0; i-- {
			if ogu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ogu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ogu *OutcomeGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ogu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ogu *OutcomeGroupUpdate) Exec(ctx context.Context) error {
	_, err := ogu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogu *OutcomeGroupUpdate) ExecX(ctx context.Context) {
	if err := ogu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ogu *OutcomeGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomegroup.Table,
			Columns: outcomegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomegroup.FieldID,
			},
		},
	}
	if ps := ogu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ogu.mutation.OutcomeGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupID,
		})
	}
	if value, ok := ogu.mutation.OutcomeGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupTitle,
		})
	}
	if value, ok := ogu.mutation.OutcomeGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupDescription,
		})
	}
	if ogu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomegroup.ParentTable,
			Columns: []string{outcomegroup.ParentColumn},
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
	if nodes := ogu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomegroup.ParentTable,
			Columns: []string{outcomegroup.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ogu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeGroupUpdateOne is the builder for updating a single OutcomeGroup entity.
type OutcomeGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeGroupMutation
}

// SetOutcomeGroupID sets the "outcome_group_id" field.
func (oguo *OutcomeGroupUpdateOne) SetOutcomeGroupID(s string) *OutcomeGroupUpdateOne {
	oguo.mutation.SetOutcomeGroupID(s)
	return oguo
}

// SetOutcomeGroupTitle sets the "outcome_group_title" field.
func (oguo *OutcomeGroupUpdateOne) SetOutcomeGroupTitle(s string) *OutcomeGroupUpdateOne {
	oguo.mutation.SetOutcomeGroupTitle(s)
	return oguo
}

// SetOutcomeGroupDescription sets the "outcome_group_description" field.
func (oguo *OutcomeGroupUpdateOne) SetOutcomeGroupDescription(s string) *OutcomeGroupUpdateOne {
	oguo.mutation.SetOutcomeGroupDescription(s)
	return oguo
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oguo *OutcomeGroupUpdateOne) SetParentID(id int) *OutcomeGroupUpdateOne {
	oguo.mutation.SetParentID(id)
	return oguo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oguo *OutcomeGroupUpdateOne) SetNillableParentID(id *int) *OutcomeGroupUpdateOne {
	if id != nil {
		oguo = oguo.SetParentID(*id)
	}
	return oguo
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oguo *OutcomeGroupUpdateOne) SetParent(o *OutcomeMeasure) *OutcomeGroupUpdateOne {
	return oguo.SetParentID(o.ID)
}

// Mutation returns the OutcomeGroupMutation object of the builder.
func (oguo *OutcomeGroupUpdateOne) Mutation() *OutcomeGroupMutation {
	return oguo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oguo *OutcomeGroupUpdateOne) ClearParent() *OutcomeGroupUpdateOne {
	oguo.mutation.ClearParent()
	return oguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oguo *OutcomeGroupUpdateOne) Select(field string, fields ...string) *OutcomeGroupUpdateOne {
	oguo.fields = append([]string{field}, fields...)
	return oguo
}

// Save executes the query and returns the updated OutcomeGroup entity.
func (oguo *OutcomeGroupUpdateOne) Save(ctx context.Context) (*OutcomeGroup, error) {
	var (
		err  error
		node *OutcomeGroup
	)
	if len(oguo.hooks) == 0 {
		node, err = oguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oguo.mutation = mutation
			node, err = oguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oguo.hooks) - 1; i >= 0; i-- {
			if oguo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oguo *OutcomeGroupUpdateOne) SaveX(ctx context.Context) *OutcomeGroup {
	node, err := oguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oguo *OutcomeGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := oguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oguo *OutcomeGroupUpdateOne) ExecX(ctx context.Context) {
	if err := oguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oguo *OutcomeGroupUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomegroup.Table,
			Columns: outcomegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomegroup.FieldID,
			},
		},
	}
	id, ok := oguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomegroup.FieldID)
		for _, f := range fields {
			if !outcomegroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oguo.mutation.OutcomeGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupID,
		})
	}
	if value, ok := oguo.mutation.OutcomeGroupTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupTitle,
		})
	}
	if value, ok := oguo.mutation.OutcomeGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomegroup.FieldOutcomeGroupDescription,
		})
	}
	if oguo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomegroup.ParentTable,
			Columns: []string{outcomegroup.ParentColumn},
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
	if nodes := oguo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomegroup.ParentTable,
			Columns: []string{outcomegroup.ParentColumn},
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
	_node = &OutcomeGroup{config: oguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
