// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// OutcomeMeasuresModuleUpdate is the builder for updating OutcomeMeasuresModule entities.
type OutcomeMeasuresModuleUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasuresModuleMutation
}

// Where appends a list predicates to the OutcomeMeasuresModuleUpdate builder.
func (ommu *OutcomeMeasuresModuleUpdate) Where(ps ...predicate.OutcomeMeasuresModule) *OutcomeMeasuresModuleUpdate {
	ommu.mutation.Where(ps...)
	return ommu
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (ommu *OutcomeMeasuresModuleUpdate) SetParentID(id int) *OutcomeMeasuresModuleUpdate {
	ommu.mutation.SetParentID(id)
	return ommu
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ommu *OutcomeMeasuresModuleUpdate) SetNillableParentID(id *int) *OutcomeMeasuresModuleUpdate {
	if id != nil {
		ommu = ommu.SetParentID(*id)
	}
	return ommu
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (ommu *OutcomeMeasuresModuleUpdate) SetParent(r *ResultsDefinition) *OutcomeMeasuresModuleUpdate {
	return ommu.SetParentID(r.ID)
}

// AddOutcomeMeasureListIDs adds the "outcome_measure_list" edge to the OutcomeMeasure entity by IDs.
func (ommu *OutcomeMeasuresModuleUpdate) AddOutcomeMeasureListIDs(ids ...int) *OutcomeMeasuresModuleUpdate {
	ommu.mutation.AddOutcomeMeasureListIDs(ids...)
	return ommu
}

// AddOutcomeMeasureList adds the "outcome_measure_list" edges to the OutcomeMeasure entity.
func (ommu *OutcomeMeasuresModuleUpdate) AddOutcomeMeasureList(o ...*OutcomeMeasure) *OutcomeMeasuresModuleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ommu.AddOutcomeMeasureListIDs(ids...)
}

// Mutation returns the OutcomeMeasuresModuleMutation object of the builder.
func (ommu *OutcomeMeasuresModuleUpdate) Mutation() *OutcomeMeasuresModuleMutation {
	return ommu.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (ommu *OutcomeMeasuresModuleUpdate) ClearParent() *OutcomeMeasuresModuleUpdate {
	ommu.mutation.ClearParent()
	return ommu
}

// ClearOutcomeMeasureList clears all "outcome_measure_list" edges to the OutcomeMeasure entity.
func (ommu *OutcomeMeasuresModuleUpdate) ClearOutcomeMeasureList() *OutcomeMeasuresModuleUpdate {
	ommu.mutation.ClearOutcomeMeasureList()
	return ommu
}

// RemoveOutcomeMeasureListIDs removes the "outcome_measure_list" edge to OutcomeMeasure entities by IDs.
func (ommu *OutcomeMeasuresModuleUpdate) RemoveOutcomeMeasureListIDs(ids ...int) *OutcomeMeasuresModuleUpdate {
	ommu.mutation.RemoveOutcomeMeasureListIDs(ids...)
	return ommu
}

// RemoveOutcomeMeasureList removes "outcome_measure_list" edges to OutcomeMeasure entities.
func (ommu *OutcomeMeasuresModuleUpdate) RemoveOutcomeMeasureList(o ...*OutcomeMeasure) *OutcomeMeasuresModuleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ommu.RemoveOutcomeMeasureListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ommu *OutcomeMeasuresModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ommu.hooks) == 0 {
		affected, err = ommu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasuresModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ommu.mutation = mutation
			affected, err = ommu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ommu.hooks) - 1; i >= 0; i-- {
			if ommu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ommu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ommu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ommu *OutcomeMeasuresModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := ommu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ommu *OutcomeMeasuresModuleUpdate) Exec(ctx context.Context) error {
	_, err := ommu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ommu *OutcomeMeasuresModuleUpdate) ExecX(ctx context.Context) {
	if err := ommu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ommu *OutcomeMeasuresModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasuresmodule.Table,
			Columns: outcomemeasuresmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasuresmodule.FieldID,
			},
		},
	}
	if ps := ommu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ommu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   outcomemeasuresmodule.ParentTable,
			Columns: []string{outcomemeasuresmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ommu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   outcomemeasuresmodule.ParentTable,
			Columns: []string{outcomemeasuresmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ommu.mutation.OutcomeMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
	if nodes := ommu.mutation.RemovedOutcomeMeasureListIDs(); len(nodes) > 0 && !ommu.mutation.OutcomeMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ommu.mutation.OutcomeMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ommu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasuresmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeMeasuresModuleUpdateOne is the builder for updating a single OutcomeMeasuresModule entity.
type OutcomeMeasuresModuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeMeasuresModuleMutation
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (ommuo *OutcomeMeasuresModuleUpdateOne) SetParentID(id int) *OutcomeMeasuresModuleUpdateOne {
	ommuo.mutation.SetParentID(id)
	return ommuo
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ommuo *OutcomeMeasuresModuleUpdateOne) SetNillableParentID(id *int) *OutcomeMeasuresModuleUpdateOne {
	if id != nil {
		ommuo = ommuo.SetParentID(*id)
	}
	return ommuo
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) SetParent(r *ResultsDefinition) *OutcomeMeasuresModuleUpdateOne {
	return ommuo.SetParentID(r.ID)
}

// AddOutcomeMeasureListIDs adds the "outcome_measure_list" edge to the OutcomeMeasure entity by IDs.
func (ommuo *OutcomeMeasuresModuleUpdateOne) AddOutcomeMeasureListIDs(ids ...int) *OutcomeMeasuresModuleUpdateOne {
	ommuo.mutation.AddOutcomeMeasureListIDs(ids...)
	return ommuo
}

// AddOutcomeMeasureList adds the "outcome_measure_list" edges to the OutcomeMeasure entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) AddOutcomeMeasureList(o ...*OutcomeMeasure) *OutcomeMeasuresModuleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ommuo.AddOutcomeMeasureListIDs(ids...)
}

// Mutation returns the OutcomeMeasuresModuleMutation object of the builder.
func (ommuo *OutcomeMeasuresModuleUpdateOne) Mutation() *OutcomeMeasuresModuleMutation {
	return ommuo.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) ClearParent() *OutcomeMeasuresModuleUpdateOne {
	ommuo.mutation.ClearParent()
	return ommuo
}

// ClearOutcomeMeasureList clears all "outcome_measure_list" edges to the OutcomeMeasure entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) ClearOutcomeMeasureList() *OutcomeMeasuresModuleUpdateOne {
	ommuo.mutation.ClearOutcomeMeasureList()
	return ommuo
}

// RemoveOutcomeMeasureListIDs removes the "outcome_measure_list" edge to OutcomeMeasure entities by IDs.
func (ommuo *OutcomeMeasuresModuleUpdateOne) RemoveOutcomeMeasureListIDs(ids ...int) *OutcomeMeasuresModuleUpdateOne {
	ommuo.mutation.RemoveOutcomeMeasureListIDs(ids...)
	return ommuo
}

// RemoveOutcomeMeasureList removes "outcome_measure_list" edges to OutcomeMeasure entities.
func (ommuo *OutcomeMeasuresModuleUpdateOne) RemoveOutcomeMeasureList(o ...*OutcomeMeasure) *OutcomeMeasuresModuleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ommuo.RemoveOutcomeMeasureListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ommuo *OutcomeMeasuresModuleUpdateOne) Select(field string, fields ...string) *OutcomeMeasuresModuleUpdateOne {
	ommuo.fields = append([]string{field}, fields...)
	return ommuo
}

// Save executes the query and returns the updated OutcomeMeasuresModule entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) Save(ctx context.Context) (*OutcomeMeasuresModule, error) {
	var (
		err  error
		node *OutcomeMeasuresModule
	)
	if len(ommuo.hooks) == 0 {
		node, err = ommuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasuresModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ommuo.mutation = mutation
			node, err = ommuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ommuo.hooks) - 1; i >= 0; i-- {
			if ommuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ommuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ommuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ommuo *OutcomeMeasuresModuleUpdateOne) SaveX(ctx context.Context) *OutcomeMeasuresModule {
	node, err := ommuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ommuo *OutcomeMeasuresModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := ommuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ommuo *OutcomeMeasuresModuleUpdateOne) ExecX(ctx context.Context) {
	if err := ommuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ommuo *OutcomeMeasuresModuleUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeMeasuresModule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasuresmodule.Table,
			Columns: outcomemeasuresmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasuresmodule.FieldID,
			},
		},
	}
	id, ok := ommuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeMeasuresModule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ommuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasuresmodule.FieldID)
		for _, f := range fields {
			if !outcomemeasuresmodule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomemeasuresmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ommuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ommuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   outcomemeasuresmodule.ParentTable,
			Columns: []string{outcomemeasuresmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ommuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   outcomemeasuresmodule.ParentTable,
			Columns: []string{outcomemeasuresmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ommuo.mutation.OutcomeMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
	if nodes := ommuo.mutation.RemovedOutcomeMeasureListIDs(); len(nodes) > 0 && !ommuo.mutation.OutcomeMeasureListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ommuo.mutation.OutcomeMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
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
	_node = &OutcomeMeasuresModule{config: ommuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ommuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasuresmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
