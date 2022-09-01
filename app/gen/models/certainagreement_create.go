// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
)

// CertainAgreementCreate is the builder for creating a CertainAgreement entity.
type CertainAgreementCreate struct {
	config
	mutation *CertainAgreementMutation
	hooks    []Hook
}

// SetAgreementPiSponsorEmployee sets the "agreement_pi_sponsor_employee" field.
func (cac *CertainAgreementCreate) SetAgreementPiSponsorEmployee(s string) *CertainAgreementCreate {
	cac.mutation.SetAgreementPiSponsorEmployee(s)
	return cac
}

// SetAgreementRestrictionType sets the "agreement_restriction_type" field.
func (cac *CertainAgreementCreate) SetAgreementRestrictionType(s string) *CertainAgreementCreate {
	cac.mutation.SetAgreementRestrictionType(s)
	return cac
}

// SetAgreementRestrictiveAgreement sets the "agreement_restrictive_agreement" field.
func (cac *CertainAgreementCreate) SetAgreementRestrictiveAgreement(s string) *CertainAgreementCreate {
	cac.mutation.SetAgreementRestrictiveAgreement(s)
	return cac
}

// SetAgreementOtherDetails sets the "agreement_other_details" field.
func (cac *CertainAgreementCreate) SetAgreementOtherDetails(s string) *CertainAgreementCreate {
	cac.mutation.SetAgreementOtherDetails(s)
	return cac
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (cac *CertainAgreementCreate) SetParentID(id int) *CertainAgreementCreate {
	cac.mutation.SetParentID(id)
	return cac
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (cac *CertainAgreementCreate) SetParent(m *MoreInfoModule) *CertainAgreementCreate {
	return cac.SetParentID(m.ID)
}

// Mutation returns the CertainAgreementMutation object of the builder.
func (cac *CertainAgreementCreate) Mutation() *CertainAgreementMutation {
	return cac.mutation
}

// Save creates the CertainAgreement in the database.
func (cac *CertainAgreementCreate) Save(ctx context.Context) (*CertainAgreement, error) {
	var (
		err  error
		node *CertainAgreement
	)
	if len(cac.hooks) == 0 {
		if err = cac.check(); err != nil {
			return nil, err
		}
		node, err = cac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertainAgreementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cac.check(); err != nil {
				return nil, err
			}
			cac.mutation = mutation
			if node, err = cac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cac.hooks) - 1; i >= 0; i-- {
			if cac.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = cac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cac *CertainAgreementCreate) SaveX(ctx context.Context) *CertainAgreement {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cac *CertainAgreementCreate) Exec(ctx context.Context) error {
	_, err := cac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cac *CertainAgreementCreate) ExecX(ctx context.Context) {
	if err := cac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *CertainAgreementCreate) check() error {
	if _, ok := cac.mutation.AgreementPiSponsorEmployee(); !ok {
		return &ValidationError{Name: "agreement_pi_sponsor_employee", err: errors.New(`models: missing required field "CertainAgreement.agreement_pi_sponsor_employee"`)}
	}
	if _, ok := cac.mutation.AgreementRestrictionType(); !ok {
		return &ValidationError{Name: "agreement_restriction_type", err: errors.New(`models: missing required field "CertainAgreement.agreement_restriction_type"`)}
	}
	if _, ok := cac.mutation.AgreementRestrictiveAgreement(); !ok {
		return &ValidationError{Name: "agreement_restrictive_agreement", err: errors.New(`models: missing required field "CertainAgreement.agreement_restrictive_agreement"`)}
	}
	if _, ok := cac.mutation.AgreementOtherDetails(); !ok {
		return &ValidationError{Name: "agreement_other_details", err: errors.New(`models: missing required field "CertainAgreement.agreement_other_details"`)}
	}
	if _, ok := cac.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`models: missing required edge "CertainAgreement.parent"`)}
	}
	return nil
}

func (cac *CertainAgreementCreate) sqlSave(ctx context.Context) (*CertainAgreement, error) {
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cac *CertainAgreementCreate) createSpec() (*CertainAgreement, *sqlgraph.CreateSpec) {
	var (
		_node = &CertainAgreement{config: cac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: certainagreement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certainagreement.FieldID,
			},
		}
	)
	if value, ok := cac.mutation.AgreementPiSponsorEmployee(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementPiSponsorEmployee,
		})
		_node.AgreementPiSponsorEmployee = value
	}
	if value, ok := cac.mutation.AgreementRestrictionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictionType,
		})
		_node.AgreementRestrictionType = value
	}
	if value, ok := cac.mutation.AgreementRestrictiveAgreement(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictiveAgreement,
		})
		_node.AgreementRestrictiveAgreement = value
	}
	if value, ok := cac.mutation.AgreementOtherDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementOtherDetails,
		})
		_node.AgreementOtherDetails = value
	}
	if nodes := cac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certainagreement.ParentTable,
			Columns: []string{certainagreement.ParentColumn},
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
		_node.certain_agreement_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CertainAgreementCreateBulk is the builder for creating many CertainAgreement entities in bulk.
type CertainAgreementCreateBulk struct {
	config
	builders []*CertainAgreementCreate
}

// Save creates the CertainAgreement entities in the database.
func (cacb *CertainAgreementCreateBulk) Save(ctx context.Context) ([]*CertainAgreement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*CertainAgreement, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CertainAgreementMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *CertainAgreementCreateBulk) SaveX(ctx context.Context) []*CertainAgreement {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cacb *CertainAgreementCreateBulk) Exec(ctx context.Context) error {
	_, err := cacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cacb *CertainAgreementCreateBulk) ExecX(ctx context.Context) {
	if err := cacb.Exec(ctx); err != nil {
		panic(err)
	}
}
