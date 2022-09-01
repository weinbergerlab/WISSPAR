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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasurementUpdate is the builder for updating OutcomeMeasurement entities.
type OutcomeMeasurementUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasurementMutation
}

// Where appends a list predicates to the OutcomeMeasurementUpdate builder.
func (omu *OutcomeMeasurementUpdate) Where(ps ...predicate.OutcomeMeasurement) *OutcomeMeasurementUpdate {
	omu.mutation.Where(ps...)
	return omu
}

// SetOutcomeMeasurementGroupID sets the "outcome_measurement_group_id" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementGroupID(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementGroupID(s)
	return omu
}

// SetOutcomeMeasurementValue sets the "outcome_measurement_value" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementValue(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementValue(s)
	return omu
}

// SetOutcomeMeasurementSpread sets the "outcome_measurement_spread" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementSpread(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementSpread(s)
	return omu
}

// SetOutcomeMeasurementLowerLimit sets the "outcome_measurement_lower_limit" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementLowerLimit(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementLowerLimit(s)
	return omu
}

// SetOutcomeMeasurementUpperLimit sets the "outcome_measurement_upper_limit" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementUpperLimit(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementUpperLimit(s)
	return omu
}

// SetOutcomeMeasurementComment sets the "outcome_measurement_comment" field.
func (omu *OutcomeMeasurementUpdate) SetOutcomeMeasurementComment(s string) *OutcomeMeasurementUpdate {
	omu.mutation.SetOutcomeMeasurementComment(s)
	return omu
}

// SetParentID sets the "parent" edge to the OutcomeCategory entity by ID.
func (omu *OutcomeMeasurementUpdate) SetParentID(id int) *OutcomeMeasurementUpdate {
	omu.mutation.SetParentID(id)
	return omu
}

// SetNillableParentID sets the "parent" edge to the OutcomeCategory entity by ID if the given value is not nil.
func (omu *OutcomeMeasurementUpdate) SetNillableParentID(id *int) *OutcomeMeasurementUpdate {
	if id != nil {
		omu = omu.SetParentID(*id)
	}
	return omu
}

// SetParent sets the "parent" edge to the OutcomeCategory entity.
func (omu *OutcomeMeasurementUpdate) SetParent(o *OutcomeCategory) *OutcomeMeasurementUpdate {
	return omu.SetParentID(o.ID)
}

// Mutation returns the OutcomeMeasurementMutation object of the builder.
func (omu *OutcomeMeasurementUpdate) Mutation() *OutcomeMeasurementMutation {
	return omu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeCategory entity.
func (omu *OutcomeMeasurementUpdate) ClearParent() *OutcomeMeasurementUpdate {
	omu.mutation.ClearParent()
	return omu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (omu *OutcomeMeasurementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(omu.hooks) == 0 {
		affected, err = omu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omu.mutation = mutation
			affected, err = omu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(omu.hooks) - 1; i >= 0; i-- {
			if omu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (omu *OutcomeMeasurementUpdate) SaveX(ctx context.Context) int {
	affected, err := omu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (omu *OutcomeMeasurementUpdate) Exec(ctx context.Context) error {
	_, err := omu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omu *OutcomeMeasurementUpdate) ExecX(ctx context.Context) {
	if err := omu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (omu *OutcomeMeasurementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasurement.Table,
			Columns: outcomemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasurement.FieldID,
			},
		},
	}
	if ps := omu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omu.mutation.OutcomeMeasurementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementGroupID,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurementValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementValue,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurementSpread(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementSpread,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurementLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementLowerLimit,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurementUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementUpperLimit,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementComment,
		})
	}
	if omu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasurement.ParentTable,
			Columns: []string{outcomemeasurement.ParentColumn},
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
	if nodes := omu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasurement.ParentTable,
			Columns: []string{outcomemeasurement.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, omu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasurement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeMeasurementUpdateOne is the builder for updating a single OutcomeMeasurement entity.
type OutcomeMeasurementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeMeasurementMutation
}

// SetOutcomeMeasurementGroupID sets the "outcome_measurement_group_id" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementGroupID(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementGroupID(s)
	return omuo
}

// SetOutcomeMeasurementValue sets the "outcome_measurement_value" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementValue(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementValue(s)
	return omuo
}

// SetOutcomeMeasurementSpread sets the "outcome_measurement_spread" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementSpread(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementSpread(s)
	return omuo
}

// SetOutcomeMeasurementLowerLimit sets the "outcome_measurement_lower_limit" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementLowerLimit(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementLowerLimit(s)
	return omuo
}

// SetOutcomeMeasurementUpperLimit sets the "outcome_measurement_upper_limit" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementUpperLimit(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementUpperLimit(s)
	return omuo
}

// SetOutcomeMeasurementComment sets the "outcome_measurement_comment" field.
func (omuo *OutcomeMeasurementUpdateOne) SetOutcomeMeasurementComment(s string) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetOutcomeMeasurementComment(s)
	return omuo
}

// SetParentID sets the "parent" edge to the OutcomeCategory entity by ID.
func (omuo *OutcomeMeasurementUpdateOne) SetParentID(id int) *OutcomeMeasurementUpdateOne {
	omuo.mutation.SetParentID(id)
	return omuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeCategory entity by ID if the given value is not nil.
func (omuo *OutcomeMeasurementUpdateOne) SetNillableParentID(id *int) *OutcomeMeasurementUpdateOne {
	if id != nil {
		omuo = omuo.SetParentID(*id)
	}
	return omuo
}

// SetParent sets the "parent" edge to the OutcomeCategory entity.
func (omuo *OutcomeMeasurementUpdateOne) SetParent(o *OutcomeCategory) *OutcomeMeasurementUpdateOne {
	return omuo.SetParentID(o.ID)
}

// Mutation returns the OutcomeMeasurementMutation object of the builder.
func (omuo *OutcomeMeasurementUpdateOne) Mutation() *OutcomeMeasurementMutation {
	return omuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeCategory entity.
func (omuo *OutcomeMeasurementUpdateOne) ClearParent() *OutcomeMeasurementUpdateOne {
	omuo.mutation.ClearParent()
	return omuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (omuo *OutcomeMeasurementUpdateOne) Select(field string, fields ...string) *OutcomeMeasurementUpdateOne {
	omuo.fields = append([]string{field}, fields...)
	return omuo
}

// Save executes the query and returns the updated OutcomeMeasurement entity.
func (omuo *OutcomeMeasurementUpdateOne) Save(ctx context.Context) (*OutcomeMeasurement, error) {
	var (
		err  error
		node *OutcomeMeasurement
	)
	if len(omuo.hooks) == 0 {
		node, err = omuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omuo.mutation = mutation
			node, err = omuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(omuo.hooks) - 1; i >= 0; i-- {
			if omuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (omuo *OutcomeMeasurementUpdateOne) SaveX(ctx context.Context) *OutcomeMeasurement {
	node, err := omuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (omuo *OutcomeMeasurementUpdateOne) Exec(ctx context.Context) error {
	_, err := omuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omuo *OutcomeMeasurementUpdateOne) ExecX(ctx context.Context) {
	if err := omuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (omuo *OutcomeMeasurementUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeMeasurement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasurement.Table,
			Columns: outcomemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasurement.FieldID,
			},
		},
	}
	id, ok := omuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeMeasurement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := omuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasurement.FieldID)
		for _, f := range fields {
			if !outcomemeasurement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomemeasurement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := omuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omuo.mutation.OutcomeMeasurementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementGroupID,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurementValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementValue,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurementSpread(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementSpread,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurementLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementLowerLimit,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurementUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementUpperLimit,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasurement.FieldOutcomeMeasurementComment,
		})
	}
	if omuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasurement.ParentTable,
			Columns: []string{outcomemeasurement.ParentColumn},
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
	if nodes := omuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasurement.ParentTable,
			Columns: []string{outcomemeasurement.ParentColumn},
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
	_node = &OutcomeMeasurement{config: omuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, omuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasurement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
