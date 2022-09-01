// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
)

// StudyEligibilityUpdate is the builder for updating StudyEligibility entities.
type StudyEligibilityUpdate struct {
	config
	hooks    []Hook
	mutation *StudyEligibilityMutation
}

// Where appends a list predicates to the StudyEligibilityUpdate builder.
func (seu *StudyEligibilityUpdate) Where(ps ...predicate.StudyEligibility) *StudyEligibilityUpdate {
	seu.mutation.Where(ps...)
	return seu
}

// SetEligibilityCriteria sets the "EligibilityCriteria" field.
func (seu *StudyEligibilityUpdate) SetEligibilityCriteria(s string) *StudyEligibilityUpdate {
	seu.mutation.SetEligibilityCriteria(s)
	return seu
}

// SetHealthyVolunteers sets the "HealthyVolunteers" field.
func (seu *StudyEligibilityUpdate) SetHealthyVolunteers(s string) *StudyEligibilityUpdate {
	seu.mutation.SetHealthyVolunteers(s)
	return seu
}

// SetGender sets the "Gender" field.
func (seu *StudyEligibilityUpdate) SetGender(s string) *StudyEligibilityUpdate {
	seu.mutation.SetGender(s)
	return seu
}

// SetMinimumAge sets the "MinimumAge" field.
func (seu *StudyEligibilityUpdate) SetMinimumAge(s string) *StudyEligibilityUpdate {
	seu.mutation.SetMinimumAge(s)
	return seu
}

// SetMaximumAge sets the "MaximumAge" field.
func (seu *StudyEligibilityUpdate) SetMaximumAge(s string) *StudyEligibilityUpdate {
	seu.mutation.SetMaximumAge(s)
	return seu
}

// SetStdAgeList sets the "StdAgeList" field.
func (seu *StudyEligibilityUpdate) SetStdAgeList(s string) *StudyEligibilityUpdate {
	seu.mutation.SetStdAgeList(s)
	return seu
}

// SetEthnicity sets the "Ethnicity" field.
func (seu *StudyEligibilityUpdate) SetEthnicity(s string) *StudyEligibilityUpdate {
	seu.mutation.SetEthnicity(s)
	return seu
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (seu *StudyEligibilityUpdate) SetParentID(id int) *StudyEligibilityUpdate {
	seu.mutation.SetParentID(id)
	return seu
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (seu *StudyEligibilityUpdate) SetNillableParentID(id *int) *StudyEligibilityUpdate {
	if id != nil {
		seu = seu.SetParentID(*id)
	}
	return seu
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (seu *StudyEligibilityUpdate) SetParent(c *ClinicalTrial) *StudyEligibilityUpdate {
	return seu.SetParentID(c.ID)
}

// Mutation returns the StudyEligibilityMutation object of the builder.
func (seu *StudyEligibilityUpdate) Mutation() *StudyEligibilityMutation {
	return seu.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (seu *StudyEligibilityUpdate) ClearParent() *StudyEligibilityUpdate {
	seu.mutation.ClearParent()
	return seu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seu *StudyEligibilityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(seu.hooks) == 0 {
		affected, err = seu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyEligibilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seu.mutation = mutation
			affected, err = seu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(seu.hooks) - 1; i >= 0; i-- {
			if seu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = seu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (seu *StudyEligibilityUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *StudyEligibilityUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *StudyEligibilityUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seu *StudyEligibilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyeligibility.Table,
			Columns: studyeligibility.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyeligibility.FieldID,
			},
		},
	}
	if ps := seu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seu.mutation.EligibilityCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEligibilityCriteria,
		})
	}
	if value, ok := seu.mutation.HealthyVolunteers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldHealthyVolunteers,
		})
	}
	if value, ok := seu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldGender,
		})
	}
	if value, ok := seu.mutation.MinimumAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMinimumAge,
		})
	}
	if value, ok := seu.mutation.MaximumAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMaximumAge,
		})
	}
	if value, ok := seu.mutation.StdAgeList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldStdAgeList,
		})
	}
	if value, ok := seu.mutation.Ethnicity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEthnicity,
		})
	}
	if seu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studyeligibility.ParentTable,
			Columns: []string{studyeligibility.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studyeligibility.ParentTable,
			Columns: []string{studyeligibility.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studyeligibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StudyEligibilityUpdateOne is the builder for updating a single StudyEligibility entity.
type StudyEligibilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudyEligibilityMutation
}

// SetEligibilityCriteria sets the "EligibilityCriteria" field.
func (seuo *StudyEligibilityUpdateOne) SetEligibilityCriteria(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetEligibilityCriteria(s)
	return seuo
}

// SetHealthyVolunteers sets the "HealthyVolunteers" field.
func (seuo *StudyEligibilityUpdateOne) SetHealthyVolunteers(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetHealthyVolunteers(s)
	return seuo
}

// SetGender sets the "Gender" field.
func (seuo *StudyEligibilityUpdateOne) SetGender(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetGender(s)
	return seuo
}

// SetMinimumAge sets the "MinimumAge" field.
func (seuo *StudyEligibilityUpdateOne) SetMinimumAge(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetMinimumAge(s)
	return seuo
}

// SetMaximumAge sets the "MaximumAge" field.
func (seuo *StudyEligibilityUpdateOne) SetMaximumAge(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetMaximumAge(s)
	return seuo
}

// SetStdAgeList sets the "StdAgeList" field.
func (seuo *StudyEligibilityUpdateOne) SetStdAgeList(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetStdAgeList(s)
	return seuo
}

// SetEthnicity sets the "Ethnicity" field.
func (seuo *StudyEligibilityUpdateOne) SetEthnicity(s string) *StudyEligibilityUpdateOne {
	seuo.mutation.SetEthnicity(s)
	return seuo
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (seuo *StudyEligibilityUpdateOne) SetParentID(id int) *StudyEligibilityUpdateOne {
	seuo.mutation.SetParentID(id)
	return seuo
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (seuo *StudyEligibilityUpdateOne) SetNillableParentID(id *int) *StudyEligibilityUpdateOne {
	if id != nil {
		seuo = seuo.SetParentID(*id)
	}
	return seuo
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (seuo *StudyEligibilityUpdateOne) SetParent(c *ClinicalTrial) *StudyEligibilityUpdateOne {
	return seuo.SetParentID(c.ID)
}

// Mutation returns the StudyEligibilityMutation object of the builder.
func (seuo *StudyEligibilityUpdateOne) Mutation() *StudyEligibilityMutation {
	return seuo.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (seuo *StudyEligibilityUpdateOne) ClearParent() *StudyEligibilityUpdateOne {
	seuo.mutation.ClearParent()
	return seuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seuo *StudyEligibilityUpdateOne) Select(field string, fields ...string) *StudyEligibilityUpdateOne {
	seuo.fields = append([]string{field}, fields...)
	return seuo
}

// Save executes the query and returns the updated StudyEligibility entity.
func (seuo *StudyEligibilityUpdateOne) Save(ctx context.Context) (*StudyEligibility, error) {
	var (
		err  error
		node *StudyEligibility
	)
	if len(seuo.hooks) == 0 {
		node, err = seuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyEligibilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seuo.mutation = mutation
			node, err = seuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(seuo.hooks) - 1; i >= 0; i-- {
			if seuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = seuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *StudyEligibilityUpdateOne) SaveX(ctx context.Context) *StudyEligibility {
	node, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seuo *StudyEligibilityUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *StudyEligibilityUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seuo *StudyEligibilityUpdateOne) sqlSave(ctx context.Context) (_node *StudyEligibility, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyeligibility.Table,
			Columns: studyeligibility.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyeligibility.FieldID,
			},
		},
	}
	id, ok := seuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "StudyEligibility.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studyeligibility.FieldID)
		for _, f := range fields {
			if !studyeligibility.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != studyeligibility.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seuo.mutation.EligibilityCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEligibilityCriteria,
		})
	}
	if value, ok := seuo.mutation.HealthyVolunteers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldHealthyVolunteers,
		})
	}
	if value, ok := seuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldGender,
		})
	}
	if value, ok := seuo.mutation.MinimumAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMinimumAge,
		})
	}
	if value, ok := seuo.mutation.MaximumAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMaximumAge,
		})
	}
	if value, ok := seuo.mutation.StdAgeList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldStdAgeList,
		})
	}
	if value, ok := seuo.mutation.Ethnicity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEthnicity,
		})
	}
	if seuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studyeligibility.ParentTable,
			Columns: []string{studyeligibility.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studyeligibility.ParentTable,
			Columns: []string{studyeligibility.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudyEligibility{config: seuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studyeligibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
