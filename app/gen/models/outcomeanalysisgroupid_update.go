// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisGroupIDUpdate is the builder for updating OutcomeAnalysisGroupID entities.
type OutcomeAnalysisGroupIDUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeAnalysisGroupIDMutation
}

// Where appends a list predicates to the OutcomeAnalysisGroupIDUpdate builder.
func (oagiu *OutcomeAnalysisGroupIDUpdate) Where(ps ...predicate.OutcomeAnalysisGroupID) *OutcomeAnalysisGroupIDUpdate {
	oagiu.mutation.Where(ps...)
	return oagiu
}

// SetOutcomeAnalysisGroupID sets the "outcome_analysis_group_id" field.
func (oagiu *OutcomeAnalysisGroupIDUpdate) SetOutcomeAnalysisGroupID(s string) *OutcomeAnalysisGroupIDUpdate {
	oagiu.mutation.SetOutcomeAnalysisGroupID(s)
	return oagiu
}

// SetParentID sets the "parent" edge to the OutcomeAnalysis entity by ID.
func (oagiu *OutcomeAnalysisGroupIDUpdate) SetParentID(id int) *OutcomeAnalysisGroupIDUpdate {
	oagiu.mutation.SetParentID(id)
	return oagiu
}

// SetNillableParentID sets the "parent" edge to the OutcomeAnalysis entity by ID if the given value is not nil.
func (oagiu *OutcomeAnalysisGroupIDUpdate) SetNillableParentID(id *int) *OutcomeAnalysisGroupIDUpdate {
	if id != nil {
		oagiu = oagiu.SetParentID(*id)
	}
	return oagiu
}

// SetParent sets the "parent" edge to the OutcomeAnalysis entity.
func (oagiu *OutcomeAnalysisGroupIDUpdate) SetParent(o *OutcomeAnalysis) *OutcomeAnalysisGroupIDUpdate {
	return oagiu.SetParentID(o.ID)
}

// Mutation returns the OutcomeAnalysisGroupIDMutation object of the builder.
func (oagiu *OutcomeAnalysisGroupIDUpdate) Mutation() *OutcomeAnalysisGroupIDMutation {
	return oagiu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeAnalysis entity.
func (oagiu *OutcomeAnalysisGroupIDUpdate) ClearParent() *OutcomeAnalysisGroupIDUpdate {
	oagiu.mutation.ClearParent()
	return oagiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oagiu *OutcomeAnalysisGroupIDUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oagiu.hooks) == 0 {
		affected, err = oagiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisGroupIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oagiu.mutation = mutation
			affected, err = oagiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oagiu.hooks) - 1; i >= 0; i-- {
			if oagiu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oagiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oagiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oagiu *OutcomeAnalysisGroupIDUpdate) SaveX(ctx context.Context) int {
	affected, err := oagiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oagiu *OutcomeAnalysisGroupIDUpdate) Exec(ctx context.Context) error {
	_, err := oagiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oagiu *OutcomeAnalysisGroupIDUpdate) ExecX(ctx context.Context) {
	if err := oagiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oagiu *OutcomeAnalysisGroupIDUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysisgroupid.Table,
			Columns: outcomeanalysisgroupid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysisgroupid.FieldID,
			},
		},
	}
	if ps := oagiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oagiu.mutation.OutcomeAnalysisGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID,
		})
	}
	if oagiu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysisgroupid.ParentTable,
			Columns: []string{outcomeanalysisgroupid.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oagiu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysisgroupid.ParentTable,
			Columns: []string{outcomeanalysisgroupid.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oagiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeanalysisgroupid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeAnalysisGroupIDUpdateOne is the builder for updating a single OutcomeAnalysisGroupID entity.
type OutcomeAnalysisGroupIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeAnalysisGroupIDMutation
}

// SetOutcomeAnalysisGroupID sets the "outcome_analysis_group_id" field.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) SetOutcomeAnalysisGroupID(s string) *OutcomeAnalysisGroupIDUpdateOne {
	oagiuo.mutation.SetOutcomeAnalysisGroupID(s)
	return oagiuo
}

// SetParentID sets the "parent" edge to the OutcomeAnalysis entity by ID.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) SetParentID(id int) *OutcomeAnalysisGroupIDUpdateOne {
	oagiuo.mutation.SetParentID(id)
	return oagiuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeAnalysis entity by ID if the given value is not nil.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) SetNillableParentID(id *int) *OutcomeAnalysisGroupIDUpdateOne {
	if id != nil {
		oagiuo = oagiuo.SetParentID(*id)
	}
	return oagiuo
}

// SetParent sets the "parent" edge to the OutcomeAnalysis entity.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) SetParent(o *OutcomeAnalysis) *OutcomeAnalysisGroupIDUpdateOne {
	return oagiuo.SetParentID(o.ID)
}

// Mutation returns the OutcomeAnalysisGroupIDMutation object of the builder.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) Mutation() *OutcomeAnalysisGroupIDMutation {
	return oagiuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeAnalysis entity.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) ClearParent() *OutcomeAnalysisGroupIDUpdateOne {
	oagiuo.mutation.ClearParent()
	return oagiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) Select(field string, fields ...string) *OutcomeAnalysisGroupIDUpdateOne {
	oagiuo.fields = append([]string{field}, fields...)
	return oagiuo
}

// Save executes the query and returns the updated OutcomeAnalysisGroupID entity.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) Save(ctx context.Context) (*OutcomeAnalysisGroupID, error) {
	var (
		err  error
		node *OutcomeAnalysisGroupID
	)
	if len(oagiuo.hooks) == 0 {
		node, err = oagiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisGroupIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oagiuo.mutation = mutation
			node, err = oagiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oagiuo.hooks) - 1; i >= 0; i-- {
			if oagiuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oagiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oagiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) SaveX(ctx context.Context) *OutcomeAnalysisGroupID {
	node, err := oagiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) Exec(ctx context.Context) error {
	_, err := oagiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) ExecX(ctx context.Context) {
	if err := oagiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oagiuo *OutcomeAnalysisGroupIDUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeAnalysisGroupID, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysisgroupid.Table,
			Columns: outcomeanalysisgroupid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysisgroupid.FieldID,
			},
		},
	}
	id, ok := oagiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeAnalysisGroupID.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oagiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysisgroupid.FieldID)
		for _, f := range fields {
			if !outcomeanalysisgroupid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeanalysisgroupid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oagiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oagiuo.mutation.OutcomeAnalysisGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID,
		})
	}
	if oagiuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysisgroupid.ParentTable,
			Columns: []string{outcomeanalysisgroupid.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oagiuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysisgroupid.ParentTable,
			Columns: []string{outcomeanalysisgroupid.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeAnalysisGroupID{config: oagiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oagiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeanalysisgroupid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
