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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// StudyLocationUpdate is the builder for updating StudyLocation entities.
type StudyLocationUpdate struct {
	config
	hooks    []Hook
	mutation *StudyLocationMutation
}

// Where appends a list predicates to the StudyLocationUpdate builder.
func (slu *StudyLocationUpdate) Where(ps ...predicate.StudyLocation) *StudyLocationUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetLocationFacility sets the "LocationFacility" field.
func (slu *StudyLocationUpdate) SetLocationFacility(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationFacility(s)
	return slu
}

// SetLocationCity sets the "LocationCity" field.
func (slu *StudyLocationUpdate) SetLocationCity(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationCity(s)
	return slu
}

// SetLocationCountry sets the "LocationCountry" field.
func (slu *StudyLocationUpdate) SetLocationCountry(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationCountry(s)
	return slu
}

// SetLocationCountryCode sets the "LocationCountryCode" field.
func (slu *StudyLocationUpdate) SetLocationCountryCode(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationCountryCode(s)
	return slu
}

// SetLocationContinentName sets the "LocationContinentName" field.
func (slu *StudyLocationUpdate) SetLocationContinentName(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationContinentName(s)
	return slu
}

// SetLocationContinentCode sets the "LocationContinentCode" field.
func (slu *StudyLocationUpdate) SetLocationContinentCode(s string) *StudyLocationUpdate {
	slu.mutation.SetLocationContinentCode(s)
	return slu
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (slu *StudyLocationUpdate) SetParentID(id int) *StudyLocationUpdate {
	slu.mutation.SetParentID(id)
	return slu
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (slu *StudyLocationUpdate) SetNillableParentID(id *int) *StudyLocationUpdate {
	if id != nil {
		slu = slu.SetParentID(*id)
	}
	return slu
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (slu *StudyLocationUpdate) SetParent(c *ClinicalTrial) *StudyLocationUpdate {
	return slu.SetParentID(c.ID)
}

// Mutation returns the StudyLocationMutation object of the builder.
func (slu *StudyLocationUpdate) Mutation() *StudyLocationMutation {
	return slu.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (slu *StudyLocationUpdate) ClearParent() *StudyLocationUpdate {
	slu.mutation.ClearParent()
	return slu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *StudyLocationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(slu.hooks) == 0 {
		affected, err = slu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			slu.mutation = mutation
			affected, err = slu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(slu.hooks) - 1; i >= 0; i-- {
			if slu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = slu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, slu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (slu *StudyLocationUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *StudyLocationUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *StudyLocationUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (slu *StudyLocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studylocation.Table,
			Columns: studylocation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studylocation.FieldID,
			},
		},
	}
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.LocationFacility(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationFacility,
		})
	}
	if value, ok := slu.mutation.LocationCity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCity,
		})
	}
	if value, ok := slu.mutation.LocationCountry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountry,
		})
	}
	if value, ok := slu.mutation.LocationCountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountryCode,
		})
	}
	if value, ok := slu.mutation.LocationContinentName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentName,
		})
	}
	if value, ok := slu.mutation.LocationContinentCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentCode,
		})
	}
	if slu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := slu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studylocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StudyLocationUpdateOne is the builder for updating a single StudyLocation entity.
type StudyLocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudyLocationMutation
}

// SetLocationFacility sets the "LocationFacility" field.
func (sluo *StudyLocationUpdateOne) SetLocationFacility(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationFacility(s)
	return sluo
}

// SetLocationCity sets the "LocationCity" field.
func (sluo *StudyLocationUpdateOne) SetLocationCity(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationCity(s)
	return sluo
}

// SetLocationCountry sets the "LocationCountry" field.
func (sluo *StudyLocationUpdateOne) SetLocationCountry(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationCountry(s)
	return sluo
}

// SetLocationCountryCode sets the "LocationCountryCode" field.
func (sluo *StudyLocationUpdateOne) SetLocationCountryCode(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationCountryCode(s)
	return sluo
}

// SetLocationContinentName sets the "LocationContinentName" field.
func (sluo *StudyLocationUpdateOne) SetLocationContinentName(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationContinentName(s)
	return sluo
}

// SetLocationContinentCode sets the "LocationContinentCode" field.
func (sluo *StudyLocationUpdateOne) SetLocationContinentCode(s string) *StudyLocationUpdateOne {
	sluo.mutation.SetLocationContinentCode(s)
	return sluo
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (sluo *StudyLocationUpdateOne) SetParentID(id int) *StudyLocationUpdateOne {
	sluo.mutation.SetParentID(id)
	return sluo
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (sluo *StudyLocationUpdateOne) SetNillableParentID(id *int) *StudyLocationUpdateOne {
	if id != nil {
		sluo = sluo.SetParentID(*id)
	}
	return sluo
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (sluo *StudyLocationUpdateOne) SetParent(c *ClinicalTrial) *StudyLocationUpdateOne {
	return sluo.SetParentID(c.ID)
}

// Mutation returns the StudyLocationMutation object of the builder.
func (sluo *StudyLocationUpdateOne) Mutation() *StudyLocationMutation {
	return sluo.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (sluo *StudyLocationUpdateOne) ClearParent() *StudyLocationUpdateOne {
	sluo.mutation.ClearParent()
	return sluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *StudyLocationUpdateOne) Select(field string, fields ...string) *StudyLocationUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated StudyLocation entity.
func (sluo *StudyLocationUpdateOne) Save(ctx context.Context) (*StudyLocation, error) {
	var (
		err  error
		node *StudyLocation
	)
	if len(sluo.hooks) == 0 {
		node, err = sluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sluo.mutation = mutation
			node, err = sluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sluo.hooks) - 1; i >= 0; i-- {
			if sluo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = sluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *StudyLocationUpdateOne) SaveX(ctx context.Context) *StudyLocation {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *StudyLocationUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *StudyLocationUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sluo *StudyLocationUpdateOne) sqlSave(ctx context.Context) (_node *StudyLocation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studylocation.Table,
			Columns: studylocation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studylocation.FieldID,
			},
		},
	}
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "StudyLocation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studylocation.FieldID)
		for _, f := range fields {
			if !studylocation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != studylocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.LocationFacility(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationFacility,
		})
	}
	if value, ok := sluo.mutation.LocationCity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCity,
		})
	}
	if value, ok := sluo.mutation.LocationCountry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountry,
		})
	}
	if value, ok := sluo.mutation.LocationCountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationCountryCode,
		})
	}
	if value, ok := sluo.mutation.LocationContinentName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentName,
		})
	}
	if value, ok := sluo.mutation.LocationContinentCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studylocation.FieldLocationContinentCode,
		})
	}
	if sluo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sluo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudyLocation{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studylocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
