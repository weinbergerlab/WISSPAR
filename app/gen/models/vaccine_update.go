// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
)

// VaccineUpdate is the builder for updating Vaccine entities.
type VaccineUpdate struct {
	config
	hooks    []Hook
	mutation *VaccineMutation
}

// Where appends a list predicates to the VaccineUpdate builder.
func (vu *VaccineUpdate) Where(ps ...predicate.Vaccine) *VaccineUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetVaccineName sets the "vaccine_name" field.
func (vu *VaccineUpdate) SetVaccineName(s string) *VaccineUpdate {
	vu.mutation.SetVaccineName(s)
	return vu
}

// Mutation returns the VaccineMutation object of the builder.
func (vu *VaccineUpdate) Mutation() *VaccineMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VaccineUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VaccineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			if vu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VaccineUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VaccineUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VaccineUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VaccineUpdate) check() error {
	if v, ok := vu.mutation.VaccineName(); ok {
		if err := vaccine.VaccineNameValidator(v); err != nil {
			return &ValidationError{Name: "vaccine_name", err: fmt.Errorf(`models: validator failed for field "Vaccine.vaccine_name": %w`, err)}
		}
	}
	return nil
}

func (vu *VaccineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vaccine.Table,
			Columns: vaccine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vaccine.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.VaccineName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vaccine.FieldVaccineName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vaccine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VaccineUpdateOne is the builder for updating a single Vaccine entity.
type VaccineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VaccineMutation
}

// SetVaccineName sets the "vaccine_name" field.
func (vuo *VaccineUpdateOne) SetVaccineName(s string) *VaccineUpdateOne {
	vuo.mutation.SetVaccineName(s)
	return vuo
}

// Mutation returns the VaccineMutation object of the builder.
func (vuo *VaccineUpdateOne) Mutation() *VaccineMutation {
	return vuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VaccineUpdateOne) Select(field string, fields ...string) *VaccineUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vaccine entity.
func (vuo *VaccineUpdateOne) Save(ctx context.Context) (*Vaccine, error) {
	var (
		err  error
		node *Vaccine
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VaccineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			if vuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VaccineUpdateOne) SaveX(ctx context.Context) *Vaccine {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VaccineUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VaccineUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VaccineUpdateOne) check() error {
	if v, ok := vuo.mutation.VaccineName(); ok {
		if err := vaccine.VaccineNameValidator(v); err != nil {
			return &ValidationError{Name: "vaccine_name", err: fmt.Errorf(`models: validator failed for field "Vaccine.vaccine_name": %w`, err)}
		}
	}
	return nil
}

func (vuo *VaccineUpdateOne) sqlSave(ctx context.Context) (_node *Vaccine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vaccine.Table,
			Columns: vaccine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vaccine.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Vaccine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vaccine.FieldID)
		for _, f := range fields {
			if !vaccine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != vaccine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.VaccineName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vaccine.FieldVaccineName,
		})
	}
	_node = &Vaccine{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vaccine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
