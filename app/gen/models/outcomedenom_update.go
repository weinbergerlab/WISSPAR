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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomUpdate is the builder for updating OutcomeDenom entities.
type OutcomeDenomUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeDenomMutation
}

// Where appends a list predicates to the OutcomeDenomUpdate builder.
func (odu *OutcomeDenomUpdate) Where(ps ...predicate.OutcomeDenom) *OutcomeDenomUpdate {
	odu.mutation.Where(ps...)
	return odu
}

// SetOutcomeDenomUnits sets the "outcome_denom_units" field.
func (odu *OutcomeDenomUpdate) SetOutcomeDenomUnits(s string) *OutcomeDenomUpdate {
	odu.mutation.SetOutcomeDenomUnits(s)
	return odu
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (odu *OutcomeDenomUpdate) SetParentID(id int) *OutcomeDenomUpdate {
	odu.mutation.SetParentID(id)
	return odu
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (odu *OutcomeDenomUpdate) SetNillableParentID(id *int) *OutcomeDenomUpdate {
	if id != nil {
		odu = odu.SetParentID(*id)
	}
	return odu
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (odu *OutcomeDenomUpdate) SetParent(o *OutcomeMeasure) *OutcomeDenomUpdate {
	return odu.SetParentID(o.ID)
}

// AddOutcomeDenomCountListIDs adds the "outcome_denom_count_list" edge to the OutcomeDenomCount entity by IDs.
func (odu *OutcomeDenomUpdate) AddOutcomeDenomCountListIDs(ids ...int) *OutcomeDenomUpdate {
	odu.mutation.AddOutcomeDenomCountListIDs(ids...)
	return odu
}

// AddOutcomeDenomCountList adds the "outcome_denom_count_list" edges to the OutcomeDenomCount entity.
func (odu *OutcomeDenomUpdate) AddOutcomeDenomCountList(o ...*OutcomeDenomCount) *OutcomeDenomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return odu.AddOutcomeDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeDenomMutation object of the builder.
func (odu *OutcomeDenomUpdate) Mutation() *OutcomeDenomMutation {
	return odu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (odu *OutcomeDenomUpdate) ClearParent() *OutcomeDenomUpdate {
	odu.mutation.ClearParent()
	return odu
}

// ClearOutcomeDenomCountList clears all "outcome_denom_count_list" edges to the OutcomeDenomCount entity.
func (odu *OutcomeDenomUpdate) ClearOutcomeDenomCountList() *OutcomeDenomUpdate {
	odu.mutation.ClearOutcomeDenomCountList()
	return odu
}

// RemoveOutcomeDenomCountListIDs removes the "outcome_denom_count_list" edge to OutcomeDenomCount entities by IDs.
func (odu *OutcomeDenomUpdate) RemoveOutcomeDenomCountListIDs(ids ...int) *OutcomeDenomUpdate {
	odu.mutation.RemoveOutcomeDenomCountListIDs(ids...)
	return odu
}

// RemoveOutcomeDenomCountList removes "outcome_denom_count_list" edges to OutcomeDenomCount entities.
func (odu *OutcomeDenomUpdate) RemoveOutcomeDenomCountList(o ...*OutcomeDenomCount) *OutcomeDenomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return odu.RemoveOutcomeDenomCountListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odu *OutcomeDenomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odu.hooks) == 0 {
		affected, err = odu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odu.mutation = mutation
			affected, err = odu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odu.hooks) - 1; i >= 0; i-- {
			if odu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = odu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (odu *OutcomeDenomUpdate) SaveX(ctx context.Context) int {
	affected, err := odu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odu *OutcomeDenomUpdate) Exec(ctx context.Context) error {
	_, err := odu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odu *OutcomeDenomUpdate) ExecX(ctx context.Context) {
	if err := odu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (odu *OutcomeDenomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenom.Table,
			Columns: outcomedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenom.FieldID,
			},
		},
	}
	if ps := odu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odu.mutation.OutcomeDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenom.FieldOutcomeDenomUnits,
		})
	}
	if odu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenom.ParentTable,
			Columns: []string{outcomedenom.ParentColumn},
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
	if nodes := odu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenom.ParentTable,
			Columns: []string{outcomedenom.ParentColumn},
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
	if odu.mutation.OutcomeDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.RemovedOutcomeDenomCountListIDs(); len(nodes) > 0 && !odu.mutation.OutcomeDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.OutcomeDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, odu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomedenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeDenomUpdateOne is the builder for updating a single OutcomeDenom entity.
type OutcomeDenomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeDenomMutation
}

// SetOutcomeDenomUnits sets the "outcome_denom_units" field.
func (oduo *OutcomeDenomUpdateOne) SetOutcomeDenomUnits(s string) *OutcomeDenomUpdateOne {
	oduo.mutation.SetOutcomeDenomUnits(s)
	return oduo
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oduo *OutcomeDenomUpdateOne) SetParentID(id int) *OutcomeDenomUpdateOne {
	oduo.mutation.SetParentID(id)
	return oduo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oduo *OutcomeDenomUpdateOne) SetNillableParentID(id *int) *OutcomeDenomUpdateOne {
	if id != nil {
		oduo = oduo.SetParentID(*id)
	}
	return oduo
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oduo *OutcomeDenomUpdateOne) SetParent(o *OutcomeMeasure) *OutcomeDenomUpdateOne {
	return oduo.SetParentID(o.ID)
}

// AddOutcomeDenomCountListIDs adds the "outcome_denom_count_list" edge to the OutcomeDenomCount entity by IDs.
func (oduo *OutcomeDenomUpdateOne) AddOutcomeDenomCountListIDs(ids ...int) *OutcomeDenomUpdateOne {
	oduo.mutation.AddOutcomeDenomCountListIDs(ids...)
	return oduo
}

// AddOutcomeDenomCountList adds the "outcome_denom_count_list" edges to the OutcomeDenomCount entity.
func (oduo *OutcomeDenomUpdateOne) AddOutcomeDenomCountList(o ...*OutcomeDenomCount) *OutcomeDenomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oduo.AddOutcomeDenomCountListIDs(ids...)
}

// Mutation returns the OutcomeDenomMutation object of the builder.
func (oduo *OutcomeDenomUpdateOne) Mutation() *OutcomeDenomMutation {
	return oduo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oduo *OutcomeDenomUpdateOne) ClearParent() *OutcomeDenomUpdateOne {
	oduo.mutation.ClearParent()
	return oduo
}

// ClearOutcomeDenomCountList clears all "outcome_denom_count_list" edges to the OutcomeDenomCount entity.
func (oduo *OutcomeDenomUpdateOne) ClearOutcomeDenomCountList() *OutcomeDenomUpdateOne {
	oduo.mutation.ClearOutcomeDenomCountList()
	return oduo
}

// RemoveOutcomeDenomCountListIDs removes the "outcome_denom_count_list" edge to OutcomeDenomCount entities by IDs.
func (oduo *OutcomeDenomUpdateOne) RemoveOutcomeDenomCountListIDs(ids ...int) *OutcomeDenomUpdateOne {
	oduo.mutation.RemoveOutcomeDenomCountListIDs(ids...)
	return oduo
}

// RemoveOutcomeDenomCountList removes "outcome_denom_count_list" edges to OutcomeDenomCount entities.
func (oduo *OutcomeDenomUpdateOne) RemoveOutcomeDenomCountList(o ...*OutcomeDenomCount) *OutcomeDenomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oduo.RemoveOutcomeDenomCountListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oduo *OutcomeDenomUpdateOne) Select(field string, fields ...string) *OutcomeDenomUpdateOne {
	oduo.fields = append([]string{field}, fields...)
	return oduo
}

// Save executes the query and returns the updated OutcomeDenom entity.
func (oduo *OutcomeDenomUpdateOne) Save(ctx context.Context) (*OutcomeDenom, error) {
	var (
		err  error
		node *OutcomeDenom
	)
	if len(oduo.hooks) == 0 {
		node, err = oduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oduo.mutation = mutation
			node, err = oduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oduo.hooks) - 1; i >= 0; i-- {
			if oduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oduo *OutcomeDenomUpdateOne) SaveX(ctx context.Context) *OutcomeDenom {
	node, err := oduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oduo *OutcomeDenomUpdateOne) Exec(ctx context.Context) error {
	_, err := oduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oduo *OutcomeDenomUpdateOne) ExecX(ctx context.Context) {
	if err := oduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oduo *OutcomeDenomUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeDenom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenom.Table,
			Columns: outcomedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenom.FieldID,
			},
		},
	}
	id, ok := oduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeDenom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenom.FieldID)
		for _, f := range fields {
			if !outcomedenom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomedenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oduo.mutation.OutcomeDenomUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomedenom.FieldOutcomeDenomUnits,
		})
	}
	if oduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenom.ParentTable,
			Columns: []string{outcomedenom.ParentColumn},
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
	if nodes := oduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomedenom.ParentTable,
			Columns: []string{outcomedenom.ParentColumn},
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
	if oduo.mutation.OutcomeDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.RemovedOutcomeDenomCountListIDs(); len(nodes) > 0 && !oduo.mutation.OutcomeDenomCountListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.OutcomeDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomedenom.OutcomeDenomCountListTable,
			Columns: []string{outcomedenom.OutcomeDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenomcount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeDenom{config: oduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomedenom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
