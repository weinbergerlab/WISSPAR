// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
)

// PointOfContactCreate is the builder for creating a PointOfContact entity.
type PointOfContactCreate struct {
	config
	mutation *PointOfContactMutation
	hooks    []Hook
}

// SetPointOfContactTitle sets the "point_of_contact_title" field.
func (pocc *PointOfContactCreate) SetPointOfContactTitle(s string) *PointOfContactCreate {
	pocc.mutation.SetPointOfContactTitle(s)
	return pocc
}

// SetPointOfContactOrganization sets the "point_of_contact_organization" field.
func (pocc *PointOfContactCreate) SetPointOfContactOrganization(s string) *PointOfContactCreate {
	pocc.mutation.SetPointOfContactOrganization(s)
	return pocc
}

// SetPointOfContactEmail sets the "point_of_contact_email" field.
func (pocc *PointOfContactCreate) SetPointOfContactEmail(s string) *PointOfContactCreate {
	pocc.mutation.SetPointOfContactEmail(s)
	return pocc
}

// SetPointOfContactPhone sets the "point_of_contact_phone" field.
func (pocc *PointOfContactCreate) SetPointOfContactPhone(s string) *PointOfContactCreate {
	pocc.mutation.SetPointOfContactPhone(s)
	return pocc
}

// SetPointOfContactPhoneExt sets the "point_of_contact_phone_ext" field.
func (pocc *PointOfContactCreate) SetPointOfContactPhoneExt(s string) *PointOfContactCreate {
	pocc.mutation.SetPointOfContactPhoneExt(s)
	return pocc
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (pocc *PointOfContactCreate) SetParentID(id int) *PointOfContactCreate {
	pocc.mutation.SetParentID(id)
	return pocc
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (pocc *PointOfContactCreate) SetParent(m *MoreInfoModule) *PointOfContactCreate {
	return pocc.SetParentID(m.ID)
}

// Mutation returns the PointOfContactMutation object of the builder.
func (pocc *PointOfContactCreate) Mutation() *PointOfContactMutation {
	return pocc.mutation
}

// Save creates the PointOfContact in the database.
func (pocc *PointOfContactCreate) Save(ctx context.Context) (*PointOfContact, error) {
	var (
		err  error
		node *PointOfContact
	)
	if len(pocc.hooks) == 0 {
		if err = pocc.check(); err != nil {
			return nil, err
		}
		node, err = pocc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PointOfContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pocc.check(); err != nil {
				return nil, err
			}
			pocc.mutation = mutation
			if node, err = pocc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pocc.hooks) - 1; i >= 0; i-- {
			if pocc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pocc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pocc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pocc *PointOfContactCreate) SaveX(ctx context.Context) *PointOfContact {
	v, err := pocc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pocc *PointOfContactCreate) Exec(ctx context.Context) error {
	_, err := pocc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocc *PointOfContactCreate) ExecX(ctx context.Context) {
	if err := pocc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pocc *PointOfContactCreate) check() error {
	if _, ok := pocc.mutation.PointOfContactTitle(); !ok {
		return &ValidationError{Name: "point_of_contact_title", err: errors.New(`models: missing required field "PointOfContact.point_of_contact_title"`)}
	}
	if _, ok := pocc.mutation.PointOfContactOrganization(); !ok {
		return &ValidationError{Name: "point_of_contact_organization", err: errors.New(`models: missing required field "PointOfContact.point_of_contact_organization"`)}
	}
	if _, ok := pocc.mutation.PointOfContactEmail(); !ok {
		return &ValidationError{Name: "point_of_contact_email", err: errors.New(`models: missing required field "PointOfContact.point_of_contact_email"`)}
	}
	if _, ok := pocc.mutation.PointOfContactPhone(); !ok {
		return &ValidationError{Name: "point_of_contact_phone", err: errors.New(`models: missing required field "PointOfContact.point_of_contact_phone"`)}
	}
	if _, ok := pocc.mutation.PointOfContactPhoneExt(); !ok {
		return &ValidationError{Name: "point_of_contact_phone_ext", err: errors.New(`models: missing required field "PointOfContact.point_of_contact_phone_ext"`)}
	}
	if _, ok := pocc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`models: missing required edge "PointOfContact.parent"`)}
	}
	return nil
}

func (pocc *PointOfContactCreate) sqlSave(ctx context.Context) (*PointOfContact, error) {
	_node, _spec := pocc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pocc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pocc *PointOfContactCreate) createSpec() (*PointOfContact, *sqlgraph.CreateSpec) {
	var (
		_node = &PointOfContact{config: pocc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pointofcontact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pointofcontact.FieldID,
			},
		}
	)
	if value, ok := pocc.mutation.PointOfContactTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactTitle,
		})
		_node.PointOfContactTitle = value
	}
	if value, ok := pocc.mutation.PointOfContactOrganization(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactOrganization,
		})
		_node.PointOfContactOrganization = value
	}
	if value, ok := pocc.mutation.PointOfContactEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactEmail,
		})
		_node.PointOfContactEmail = value
	}
	if value, ok := pocc.mutation.PointOfContactPhone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhone,
		})
		_node.PointOfContactPhone = value
	}
	if value, ok := pocc.mutation.PointOfContactPhoneExt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhoneExt,
		})
		_node.PointOfContactPhoneExt = value
	}
	if nodes := pocc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pointofcontact.ParentTable,
			Columns: []string{pointofcontact.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.point_of_contact_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PointOfContactCreateBulk is the builder for creating many PointOfContact entities in bulk.
type PointOfContactCreateBulk struct {
	config
	builders []*PointOfContactCreate
}

// Save creates the PointOfContact entities in the database.
func (poccb *PointOfContactCreateBulk) Save(ctx context.Context) ([]*PointOfContact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(poccb.builders))
	nodes := make([]*PointOfContact, len(poccb.builders))
	mutators := make([]Mutator, len(poccb.builders))
	for i := range poccb.builders {
		func(i int, root context.Context) {
			builder := poccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PointOfContactMutation)
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
					_, err = mutators[i+1].Mutate(root, poccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, poccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, poccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (poccb *PointOfContactCreateBulk) SaveX(ctx context.Context) []*PointOfContact {
	v, err := poccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poccb *PointOfContactCreateBulk) Exec(ctx context.Context) error {
	_, err := poccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poccb *PointOfContactCreateBulk) ExecX(ctx context.Context) {
	if err := poccb.Exec(ctx); err != nil {
		panic(err)
	}
}
