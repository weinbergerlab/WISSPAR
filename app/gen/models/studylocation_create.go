// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// StudyLocationCreate is the builder for creating a StudyLocation entity.
type StudyLocationCreate struct {
	config
	mutation *StudyLocationMutation
	hooks    []Hook
}

// SetLocationFacility sets the "LocationFacility" field.
func (slc *StudyLocationCreate) SetLocationFacility(s string) *StudyLocationCreate {
	slc.mutation.SetLocationFacility(s)
	return slc
}

// SetLocationCity sets the "LocationCity" field.
func (slc *StudyLocationCreate) SetLocationCity(s string) *StudyLocationCreate {
	slc.mutation.SetLocationCity(s)
	return slc
}

// SetLocationCountry sets the "LocationCountry" field.
func (slc *StudyLocationCreate) SetLocationCountry(s string) *StudyLocationCreate {
	slc.mutation.SetLocationCountry(s)
	return slc
}

// SetLocationCountryCode sets the "LocationCountryCode" field.
func (slc *StudyLocationCreate) SetLocationCountryCode(s string) *StudyLocationCreate {
	slc.mutation.SetLocationCountryCode(s)
	return slc
}

// SetLocationContinentName sets the "LocationContinentName" field.
func (slc *StudyLocationCreate) SetLocationContinentName(s string) *StudyLocationCreate {
	slc.mutation.SetLocationContinentName(s)
	return slc
}

// SetLocationContinentCode sets the "LocationContinentCode" field.
func (slc *StudyLocationCreate) SetLocationContinentCode(s string) *StudyLocationCreate {
	slc.mutation.SetLocationContinentCode(s)
	return slc
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (slc *StudyLocationCreate) SetParentID(id int) *StudyLocationCreate {
	slc.mutation.SetParentID(id)
	return slc
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (slc *StudyLocationCreate) SetNillableParentID(id *int) *StudyLocationCreate {
	if id != nil {
		slc = slc.SetParentID(*id)
	}
	return slc
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (slc *StudyLocationCreate) SetParent(c *ClinicalTrial) *StudyLocationCreate {
	return slc.SetParentID(c.ID)
}

// Mutation returns the StudyLocationMutation object of the builder.
func (slc *StudyLocationCreate) Mutation() *StudyLocationMutation {
	return slc.mutation
}

// Save creates the StudyLocation in the database.
func (slc *StudyLocationCreate) Save(ctx context.Context) (*StudyLocation, error) {
	var (
		err  error
		node *StudyLocation
	)
	if len(slc.hooks) == 0 {
		if err = slc.check(); err != nil {
			return nil, err
		}
		node, err = slc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = slc.check(); err != nil {
				return nil, err
			}
			slc.mutation = mutation
			if node, err = slc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(slc.hooks) - 1; i >= 0; i-- {
			if slc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = slc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, slc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (slc *StudyLocationCreate) SaveX(ctx context.Context) *StudyLocation {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *StudyLocationCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *StudyLocationCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *StudyLocationCreate) check() error {
	if _, ok := slc.mutation.LocationFacility(); !ok {
		return &ValidationError{Name: "LocationFacility", err: errors.New(`models: missing required field "StudyLocation.LocationFacility"`)}
	}
	if _, ok := slc.mutation.LocationCity(); !ok {
		return &ValidationError{Name: "LocationCity", err: errors.New(`models: missing required field "StudyLocation.LocationCity"`)}
	}
	if _, ok := slc.mutation.LocationCountry(); !ok {
		return &ValidationError{Name: "LocationCountry", err: errors.New(`models: missing required field "StudyLocation.LocationCountry"`)}
	}
	if _, ok := slc.mutation.LocationCountryCode(); !ok {
		return &ValidationError{Name: "LocationCountryCode", err: errors.New(`models: missing required field "StudyLocation.LocationCountryCode"`)}
	}
	if _, ok := slc.mutation.LocationContinentName(); !ok {
		return &ValidationError{Name: "LocationContinentName", err: errors.New(`models: missing required field "StudyLocation.LocationContinentName"`)}
	}
	if _, ok := slc.mutation.LocationContinentCode(); !ok {
		return &ValidationError{Name: "LocationContinentCode", err: errors.New(`models: missing required field "StudyLocation.LocationContinentCode"`)}
	}
	return nil
}

func (slc *StudyLocationCreate) sqlSave(ctx context.Context) (*StudyLocation, error) {
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (slc *StudyLocationCreate) createSpec() (*StudyLocation, *sqlgraph.CreateSpec) {
	var (
		_node = &StudyLocation{config: slc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: studylocation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studylocation.FieldID,
			},
		}
	)
	if value, ok := slc.mutation.LocationFacility(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationFacility,
		})
		_node.LocationFacility = value
	}
	if value, ok := slc.mutation.LocationCity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCity,
		})
		_node.LocationCity = value
	}
	if value, ok := slc.mutation.LocationCountry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountry,
		})
		_node.LocationCountry = value
	}
	if value, ok := slc.mutation.LocationCountryCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountryCode,
		})
		_node.LocationCountryCode = value
	}
	if value, ok := slc.mutation.LocationContinentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentName,
		})
		_node.LocationContinentName = value
	}
	if value, ok := slc.mutation.LocationContinentCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentCode,
		})
		_node.LocationContinentCode = value
	}
	if nodes := slc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylocation.ParentTable,
			Columns: []string{studylocation.ParentColumn},
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
		_node.clinical_trial_study_locations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudyLocationCreateBulk is the builder for creating many StudyLocation entities in bulk.
type StudyLocationCreateBulk struct {
	config
	builders []*StudyLocationCreate
}

// Save creates the StudyLocation entities in the database.
func (slcb *StudyLocationCreateBulk) Save(ctx context.Context) ([]*StudyLocation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*StudyLocation, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyLocationMutation)
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
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *StudyLocationCreateBulk) SaveX(ctx context.Context) []*StudyLocation {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *StudyLocationCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *StudyLocationCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}
