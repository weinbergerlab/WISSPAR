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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// MoreInfoModuleCreate is the builder for creating a MoreInfoModule entity.
type MoreInfoModuleCreate struct {
	config
	mutation *MoreInfoModuleMutation
	hooks    []Hook
}

// SetLimitationsAndCaveatsDescription sets the "limitations_and_caveats_description" field.
func (mimc *MoreInfoModuleCreate) SetLimitationsAndCaveatsDescription(s string) *MoreInfoModuleCreate {
	mimc.mutation.SetLimitationsAndCaveatsDescription(s)
	return mimc
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (mimc *MoreInfoModuleCreate) SetParentID(id int) *MoreInfoModuleCreate {
	mimc.mutation.SetParentID(id)
	return mimc
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (mimc *MoreInfoModuleCreate) SetNillableParentID(id *int) *MoreInfoModuleCreate {
	if id != nil {
		mimc = mimc.SetParentID(*id)
	}
	return mimc
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (mimc *MoreInfoModuleCreate) SetParent(r *ResultsDefinition) *MoreInfoModuleCreate {
	return mimc.SetParentID(r.ID)
}

// SetCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID.
func (mimc *MoreInfoModuleCreate) SetCertainAgreementID(id int) *MoreInfoModuleCreate {
	mimc.mutation.SetCertainAgreementID(id)
	return mimc
}

// SetNillableCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID if the given value is not nil.
func (mimc *MoreInfoModuleCreate) SetNillableCertainAgreementID(id *int) *MoreInfoModuleCreate {
	if id != nil {
		mimc = mimc.SetCertainAgreementID(*id)
	}
	return mimc
}

// SetCertainAgreement sets the "certain_agreement" edge to the CertainAgreement entity.
func (mimc *MoreInfoModuleCreate) SetCertainAgreement(c *CertainAgreement) *MoreInfoModuleCreate {
	return mimc.SetCertainAgreementID(c.ID)
}

// SetPointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID.
func (mimc *MoreInfoModuleCreate) SetPointOfContactID(id int) *MoreInfoModuleCreate {
	mimc.mutation.SetPointOfContactID(id)
	return mimc
}

// SetNillablePointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID if the given value is not nil.
func (mimc *MoreInfoModuleCreate) SetNillablePointOfContactID(id *int) *MoreInfoModuleCreate {
	if id != nil {
		mimc = mimc.SetPointOfContactID(*id)
	}
	return mimc
}

// SetPointOfContact sets the "point_of_contact" edge to the PointOfContact entity.
func (mimc *MoreInfoModuleCreate) SetPointOfContact(p *PointOfContact) *MoreInfoModuleCreate {
	return mimc.SetPointOfContactID(p.ID)
}

// Mutation returns the MoreInfoModuleMutation object of the builder.
func (mimc *MoreInfoModuleCreate) Mutation() *MoreInfoModuleMutation {
	return mimc.mutation
}

// Save creates the MoreInfoModule in the database.
func (mimc *MoreInfoModuleCreate) Save(ctx context.Context) (*MoreInfoModule, error) {
	var (
		err  error
		node *MoreInfoModule
	)
	if len(mimc.hooks) == 0 {
		if err = mimc.check(); err != nil {
			return nil, err
		}
		node, err = mimc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoreInfoModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mimc.check(); err != nil {
				return nil, err
			}
			mimc.mutation = mutation
			if node, err = mimc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mimc.hooks) - 1; i >= 0; i-- {
			if mimc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = mimc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mimc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mimc *MoreInfoModuleCreate) SaveX(ctx context.Context) *MoreInfoModule {
	v, err := mimc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mimc *MoreInfoModuleCreate) Exec(ctx context.Context) error {
	_, err := mimc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mimc *MoreInfoModuleCreate) ExecX(ctx context.Context) {
	if err := mimc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mimc *MoreInfoModuleCreate) check() error {
	if _, ok := mimc.mutation.LimitationsAndCaveatsDescription(); !ok {
		return &ValidationError{Name: "limitations_and_caveats_description", err: errors.New(`models: missing required field "MoreInfoModule.limitations_and_caveats_description"`)}
	}
	return nil
}

func (mimc *MoreInfoModuleCreate) sqlSave(ctx context.Context) (*MoreInfoModule, error) {
	_node, _spec := mimc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mimc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mimc *MoreInfoModuleCreate) createSpec() (*MoreInfoModule, *sqlgraph.CreateSpec) {
	var (
		_node = &MoreInfoModule{config: mimc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moreinfomodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moreinfomodule.FieldID,
			},
		}
	)
	if value, ok := mimc.mutation.LimitationsAndCaveatsDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moreinfomodule.FieldLimitationsAndCaveatsDescription,
		})
		_node.LimitationsAndCaveatsDescription = value
	}
	if nodes := mimc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   moreinfomodule.ParentTable,
			Columns: []string{moreinfomodule.ParentColumn},
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
		_node.results_definition_more_info_module = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mimc.mutation.CertainAgreementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.CertainAgreementTable,
			Columns: []string{moreinfomodule.CertainAgreementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: certainagreement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.more_info_module_certain_agreement = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mimc.mutation.PointOfContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.PointOfContactTable,
			Columns: []string{moreinfomodule.PointOfContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pointofcontact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.more_info_module_point_of_contact = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MoreInfoModuleCreateBulk is the builder for creating many MoreInfoModule entities in bulk.
type MoreInfoModuleCreateBulk struct {
	config
	builders []*MoreInfoModuleCreate
}

// Save creates the MoreInfoModule entities in the database.
func (mimcb *MoreInfoModuleCreateBulk) Save(ctx context.Context) ([]*MoreInfoModule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mimcb.builders))
	nodes := make([]*MoreInfoModule, len(mimcb.builders))
	mutators := make([]Mutator, len(mimcb.builders))
	for i := range mimcb.builders {
		func(i int, root context.Context) {
			builder := mimcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MoreInfoModuleMutation)
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
					_, err = mutators[i+1].Mutate(root, mimcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mimcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mimcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mimcb *MoreInfoModuleCreateBulk) SaveX(ctx context.Context) []*MoreInfoModule {
	v, err := mimcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mimcb *MoreInfoModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := mimcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mimcb *MoreInfoModuleCreateBulk) ExecX(ctx context.Context) {
	if err := mimcb.Exec(ctx); err != nil {
		panic(err)
	}
}
