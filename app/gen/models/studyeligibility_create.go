// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
)

// StudyEligibilityCreate is the builder for creating a StudyEligibility entity.
type StudyEligibilityCreate struct {
	config
	mutation *StudyEligibilityMutation
	hooks    []Hook
}

// SetEligibilityCriteria sets the "EligibilityCriteria" field.
func (sec *StudyEligibilityCreate) SetEligibilityCriteria(s string) *StudyEligibilityCreate {
	sec.mutation.SetEligibilityCriteria(s)
	return sec
}

// SetHealthyVolunteers sets the "HealthyVolunteers" field.
func (sec *StudyEligibilityCreate) SetHealthyVolunteers(s string) *StudyEligibilityCreate {
	sec.mutation.SetHealthyVolunteers(s)
	return sec
}

// SetGender sets the "Gender" field.
func (sec *StudyEligibilityCreate) SetGender(s string) *StudyEligibilityCreate {
	sec.mutation.SetGender(s)
	return sec
}

// SetMinimumAge sets the "MinimumAge" field.
func (sec *StudyEligibilityCreate) SetMinimumAge(s string) *StudyEligibilityCreate {
	sec.mutation.SetMinimumAge(s)
	return sec
}

// SetMaximumAge sets the "MaximumAge" field.
func (sec *StudyEligibilityCreate) SetMaximumAge(s string) *StudyEligibilityCreate {
	sec.mutation.SetMaximumAge(s)
	return sec
}

// SetStdAgeList sets the "StdAgeList" field.
func (sec *StudyEligibilityCreate) SetStdAgeList(s string) *StudyEligibilityCreate {
	sec.mutation.SetStdAgeList(s)
	return sec
}

// SetEthnicity sets the "Ethnicity" field.
func (sec *StudyEligibilityCreate) SetEthnicity(s string) *StudyEligibilityCreate {
	sec.mutation.SetEthnicity(s)
	return sec
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (sec *StudyEligibilityCreate) SetParentID(id int) *StudyEligibilityCreate {
	sec.mutation.SetParentID(id)
	return sec
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (sec *StudyEligibilityCreate) SetNillableParentID(id *int) *StudyEligibilityCreate {
	if id != nil {
		sec = sec.SetParentID(*id)
	}
	return sec
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (sec *StudyEligibilityCreate) SetParent(c *ClinicalTrial) *StudyEligibilityCreate {
	return sec.SetParentID(c.ID)
}

// Mutation returns the StudyEligibilityMutation object of the builder.
func (sec *StudyEligibilityCreate) Mutation() *StudyEligibilityMutation {
	return sec.mutation
}

// Save creates the StudyEligibility in the database.
func (sec *StudyEligibilityCreate) Save(ctx context.Context) (*StudyEligibility, error) {
	var (
		err  error
		node *StudyEligibility
	)
	if len(sec.hooks) == 0 {
		if err = sec.check(); err != nil {
			return nil, err
		}
		node, err = sec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyEligibilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sec.check(); err != nil {
				return nil, err
			}
			sec.mutation = mutation
			if node, err = sec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sec.hooks) - 1; i >= 0; i-- {
			if sec.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sec *StudyEligibilityCreate) SaveX(ctx context.Context) *StudyEligibility {
	v, err := sec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sec *StudyEligibilityCreate) Exec(ctx context.Context) error {
	_, err := sec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sec *StudyEligibilityCreate) ExecX(ctx context.Context) {
	if err := sec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sec *StudyEligibilityCreate) check() error {
	if _, ok := sec.mutation.EligibilityCriteria(); !ok {
		return &ValidationError{Name: "EligibilityCriteria", err: errors.New(`models: missing required field "StudyEligibility.EligibilityCriteria"`)}
	}
	if _, ok := sec.mutation.HealthyVolunteers(); !ok {
		return &ValidationError{Name: "HealthyVolunteers", err: errors.New(`models: missing required field "StudyEligibility.HealthyVolunteers"`)}
	}
	if _, ok := sec.mutation.Gender(); !ok {
		return &ValidationError{Name: "Gender", err: errors.New(`models: missing required field "StudyEligibility.Gender"`)}
	}
	if _, ok := sec.mutation.MinimumAge(); !ok {
		return &ValidationError{Name: "MinimumAge", err: errors.New(`models: missing required field "StudyEligibility.MinimumAge"`)}
	}
	if _, ok := sec.mutation.MaximumAge(); !ok {
		return &ValidationError{Name: "MaximumAge", err: errors.New(`models: missing required field "StudyEligibility.MaximumAge"`)}
	}
	if _, ok := sec.mutation.StdAgeList(); !ok {
		return &ValidationError{Name: "StdAgeList", err: errors.New(`models: missing required field "StudyEligibility.StdAgeList"`)}
	}
	if _, ok := sec.mutation.Ethnicity(); !ok {
		return &ValidationError{Name: "Ethnicity", err: errors.New(`models: missing required field "StudyEligibility.Ethnicity"`)}
	}
	return nil
}

func (sec *StudyEligibilityCreate) sqlSave(ctx context.Context) (*StudyEligibility, error) {
	_node, _spec := sec.createSpec()
	if err := sqlgraph.CreateNode(ctx, sec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sec *StudyEligibilityCreate) createSpec() (*StudyEligibility, *sqlgraph.CreateSpec) {
	var (
		_node = &StudyEligibility{config: sec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: studyeligibility.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyeligibility.FieldID,
			},
		}
	)
	if value, ok := sec.mutation.EligibilityCriteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEligibilityCriteria,
		})
		_node.EligibilityCriteria = value
	}
	if value, ok := sec.mutation.HealthyVolunteers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldHealthyVolunteers,
		})
		_node.HealthyVolunteers = value
	}
	if value, ok := sec.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := sec.mutation.MinimumAge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMinimumAge,
		})
		_node.MinimumAge = value
	}
	if value, ok := sec.mutation.MaximumAge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldMaximumAge,
		})
		_node.MaximumAge = value
	}
	if value, ok := sec.mutation.StdAgeList(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldStdAgeList,
		})
		_node.StdAgeList = value
	}
	if value, ok := sec.mutation.Ethnicity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studyeligibility.FieldEthnicity,
		})
		_node.Ethnicity = value
	}
	if nodes := sec.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.clinical_trial_study_eligibility = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudyEligibilityCreateBulk is the builder for creating many StudyEligibility entities in bulk.
type StudyEligibilityCreateBulk struct {
	config
	builders []*StudyEligibilityCreate
}

// Save creates the StudyEligibility entities in the database.
func (secb *StudyEligibilityCreateBulk) Save(ctx context.Context) ([]*StudyEligibility, error) {
	specs := make([]*sqlgraph.CreateSpec, len(secb.builders))
	nodes := make([]*StudyEligibility, len(secb.builders))
	mutators := make([]Mutator, len(secb.builders))
	for i := range secb.builders {
		func(i int, root context.Context) {
			builder := secb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyEligibilityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, secb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, secb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, secb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (secb *StudyEligibilityCreateBulk) SaveX(ctx context.Context) []*StudyEligibility {
	v, err := secb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (secb *StudyEligibilityCreateBulk) Exec(ctx context.Context) error {
	_, err := secb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (secb *StudyEligibilityCreateBulk) ExecX(ctx context.Context) {
	if err := secb.Exec(ctx); err != nil {
		panic(err)
	}
}
