// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowAchievementUpdate is the builder for updating FlowAchievement entities.
type FlowAchievementUpdate struct {
	config
	hooks    []Hook
	mutation *FlowAchievementMutation
}

// Where appends a list predicates to the FlowAchievementUpdate builder.
func (fau *FlowAchievementUpdate) Where(ps ...predicate.FlowAchievement) *FlowAchievementUpdate {
	fau.mutation.Where(ps...)
	return fau
}

// SetFlowAchievementGroupID sets the "flow_achievement_group_id" field.
func (fau *FlowAchievementUpdate) SetFlowAchievementGroupID(s string) *FlowAchievementUpdate {
	fau.mutation.SetFlowAchievementGroupID(s)
	return fau
}

// SetFlowAchievementComment sets the "flow_achievement_comment" field.
func (fau *FlowAchievementUpdate) SetFlowAchievementComment(s string) *FlowAchievementUpdate {
	fau.mutation.SetFlowAchievementComment(s)
	return fau
}

// SetFlowAchievementNumSubjects sets the "flow_achievement_num_subjects" field.
func (fau *FlowAchievementUpdate) SetFlowAchievementNumSubjects(s string) *FlowAchievementUpdate {
	fau.mutation.SetFlowAchievementNumSubjects(s)
	return fau
}

// SetFlowAchievementNumUnits sets the "flow_achievement_num_units" field.
func (fau *FlowAchievementUpdate) SetFlowAchievementNumUnits(s string) *FlowAchievementUpdate {
	fau.mutation.SetFlowAchievementNumUnits(s)
	return fau
}

// SetParentID sets the "parent" edge to the FlowMilestone entity by ID.
func (fau *FlowAchievementUpdate) SetParentID(id int) *FlowAchievementUpdate {
	fau.mutation.SetParentID(id)
	return fau
}

// SetNillableParentID sets the "parent" edge to the FlowMilestone entity by ID if the given value is not nil.
func (fau *FlowAchievementUpdate) SetNillableParentID(id *int) *FlowAchievementUpdate {
	if id != nil {
		fau = fau.SetParentID(*id)
	}
	return fau
}

// SetParent sets the "parent" edge to the FlowMilestone entity.
func (fau *FlowAchievementUpdate) SetParent(f *FlowMilestone) *FlowAchievementUpdate {
	return fau.SetParentID(f.ID)
}

// Mutation returns the FlowAchievementMutation object of the builder.
func (fau *FlowAchievementUpdate) Mutation() *FlowAchievementMutation {
	return fau.mutation
}

// ClearParent clears the "parent" edge to the FlowMilestone entity.
func (fau *FlowAchievementUpdate) ClearParent() *FlowAchievementUpdate {
	fau.mutation.ClearParent()
	return fau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fau *FlowAchievementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fau.hooks) == 0 {
		affected, err = fau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowAchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fau.mutation = mutation
			affected, err = fau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fau.hooks) - 1; i >= 0; i-- {
			if fau.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fau *FlowAchievementUpdate) SaveX(ctx context.Context) int {
	affected, err := fau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fau *FlowAchievementUpdate) Exec(ctx context.Context) error {
	_, err := fau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fau *FlowAchievementUpdate) ExecX(ctx context.Context) {
	if err := fau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fau *FlowAchievementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowachievement.Table,
			Columns: flowachievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowachievement.FieldID,
			},
		},
	}
	if ps := fau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fau.mutation.FlowAchievementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementGroupID,
		})
	}
	if value, ok := fau.mutation.FlowAchievementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementComment,
		})
	}
	if value, ok := fau.mutation.FlowAchievementNumSubjects(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumSubjects,
		})
	}
	if value, ok := fau.mutation.FlowAchievementNumUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumUnits,
		})
	}
	if fau.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowachievement.ParentTable,
			Columns: []string{flowachievement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fau.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowachievement.ParentTable,
			Columns: []string{flowachievement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowachievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowAchievementUpdateOne is the builder for updating a single FlowAchievement entity.
type FlowAchievementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowAchievementMutation
}

// SetFlowAchievementGroupID sets the "flow_achievement_group_id" field.
func (fauo *FlowAchievementUpdateOne) SetFlowAchievementGroupID(s string) *FlowAchievementUpdateOne {
	fauo.mutation.SetFlowAchievementGroupID(s)
	return fauo
}

// SetFlowAchievementComment sets the "flow_achievement_comment" field.
func (fauo *FlowAchievementUpdateOne) SetFlowAchievementComment(s string) *FlowAchievementUpdateOne {
	fauo.mutation.SetFlowAchievementComment(s)
	return fauo
}

// SetFlowAchievementNumSubjects sets the "flow_achievement_num_subjects" field.
func (fauo *FlowAchievementUpdateOne) SetFlowAchievementNumSubjects(s string) *FlowAchievementUpdateOne {
	fauo.mutation.SetFlowAchievementNumSubjects(s)
	return fauo
}

// SetFlowAchievementNumUnits sets the "flow_achievement_num_units" field.
func (fauo *FlowAchievementUpdateOne) SetFlowAchievementNumUnits(s string) *FlowAchievementUpdateOne {
	fauo.mutation.SetFlowAchievementNumUnits(s)
	return fauo
}

// SetParentID sets the "parent" edge to the FlowMilestone entity by ID.
func (fauo *FlowAchievementUpdateOne) SetParentID(id int) *FlowAchievementUpdateOne {
	fauo.mutation.SetParentID(id)
	return fauo
}

// SetNillableParentID sets the "parent" edge to the FlowMilestone entity by ID if the given value is not nil.
func (fauo *FlowAchievementUpdateOne) SetNillableParentID(id *int) *FlowAchievementUpdateOne {
	if id != nil {
		fauo = fauo.SetParentID(*id)
	}
	return fauo
}

// SetParent sets the "parent" edge to the FlowMilestone entity.
func (fauo *FlowAchievementUpdateOne) SetParent(f *FlowMilestone) *FlowAchievementUpdateOne {
	return fauo.SetParentID(f.ID)
}

// Mutation returns the FlowAchievementMutation object of the builder.
func (fauo *FlowAchievementUpdateOne) Mutation() *FlowAchievementMutation {
	return fauo.mutation
}

// ClearParent clears the "parent" edge to the FlowMilestone entity.
func (fauo *FlowAchievementUpdateOne) ClearParent() *FlowAchievementUpdateOne {
	fauo.mutation.ClearParent()
	return fauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fauo *FlowAchievementUpdateOne) Select(field string, fields ...string) *FlowAchievementUpdateOne {
	fauo.fields = append([]string{field}, fields...)
	return fauo
}

// Save executes the query and returns the updated FlowAchievement entity.
func (fauo *FlowAchievementUpdateOne) Save(ctx context.Context) (*FlowAchievement, error) {
	var (
		err  error
		node *FlowAchievement
	)
	if len(fauo.hooks) == 0 {
		node, err = fauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowAchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fauo.mutation = mutation
			node, err = fauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fauo.hooks) - 1; i >= 0; i-- {
			if fauo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = fauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fauo *FlowAchievementUpdateOne) SaveX(ctx context.Context) *FlowAchievement {
	node, err := fauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fauo *FlowAchievementUpdateOne) Exec(ctx context.Context) error {
	_, err := fauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fauo *FlowAchievementUpdateOne) ExecX(ctx context.Context) {
	if err := fauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fauo *FlowAchievementUpdateOne) sqlSave(ctx context.Context) (_node *FlowAchievement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowachievement.Table,
			Columns: flowachievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowachievement.FieldID,
			},
		},
	}
	id, ok := fauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "FlowAchievement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowachievement.FieldID)
		for _, f := range fields {
			if !flowachievement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != flowachievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fauo.mutation.FlowAchievementGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementGroupID,
		})
	}
	if value, ok := fauo.mutation.FlowAchievementComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementComment,
		})
	}
	if value, ok := fauo.mutation.FlowAchievementNumSubjects(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumSubjects,
		})
	}
	if value, ok := fauo.mutation.FlowAchievementNumUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowachievement.FieldFlowAchievementNumUnits,
		})
	}
	if fauo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowachievement.ParentTable,
			Columns: []string{flowachievement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fauo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flowachievement.ParentTable,
			Columns: []string{flowachievement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowmilestone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowAchievement{config: fauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowachievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
