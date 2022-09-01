// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
)

// ClinicalTrialMetadataCreate is the builder for creating a ClinicalTrialMetadata entity.
type ClinicalTrialMetadataCreate struct {
	config
	mutation *ClinicalTrialMetadataMutation
	hooks    []Hook
}

// SetTagName sets the "tag_name" field.
func (ctmc *ClinicalTrialMetadataCreate) SetTagName(s string) *ClinicalTrialMetadataCreate {
	ctmc.mutation.SetTagName(s)
	return ctmc
}

// SetTagValue sets the "tag_value" field.
func (ctmc *ClinicalTrialMetadataCreate) SetTagValue(s string) *ClinicalTrialMetadataCreate {
	ctmc.mutation.SetTagValue(s)
	return ctmc
}

// SetUseCaseCode sets the "use_case_code" field.
func (ctmc *ClinicalTrialMetadataCreate) SetUseCaseCode(s string) *ClinicalTrialMetadataCreate {
	ctmc.mutation.SetUseCaseCode(s)
	return ctmc
}

// SetNillableUseCaseCode sets the "use_case_code" field if the given value is not nil.
func (ctmc *ClinicalTrialMetadataCreate) SetNillableUseCaseCode(s *string) *ClinicalTrialMetadataCreate {
	if s != nil {
		ctmc.SetUseCaseCode(*s)
	}
	return ctmc
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (ctmc *ClinicalTrialMetadataCreate) SetParentID(id int) *ClinicalTrialMetadataCreate {
	ctmc.mutation.SetParentID(id)
	return ctmc
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (ctmc *ClinicalTrialMetadataCreate) SetNillableParentID(id *int) *ClinicalTrialMetadataCreate {
	if id != nil {
		ctmc = ctmc.SetParentID(*id)
	}
	return ctmc
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (ctmc *ClinicalTrialMetadataCreate) SetParent(c *ClinicalTrial) *ClinicalTrialMetadataCreate {
	return ctmc.SetParentID(c.ID)
}

// Mutation returns the ClinicalTrialMetadataMutation object of the builder.
func (ctmc *ClinicalTrialMetadataCreate) Mutation() *ClinicalTrialMetadataMutation {
	return ctmc.mutation
}

// Save creates the ClinicalTrialMetadata in the database.
func (ctmc *ClinicalTrialMetadataCreate) Save(ctx context.Context) (*ClinicalTrialMetadata, error) {
	var (
		err  error
		node *ClinicalTrialMetadata
	)
	ctmc.defaults()
	if len(ctmc.hooks) == 0 {
		if err = ctmc.check(); err != nil {
			return nil, err
		}
		node, err = ctmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctmc.check(); err != nil {
				return nil, err
			}
			ctmc.mutation = mutation
			if node, err = ctmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctmc.hooks) - 1; i >= 0; i-- {
			if ctmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctmc *ClinicalTrialMetadataCreate) SaveX(ctx context.Context) *ClinicalTrialMetadata {
	v, err := ctmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctmc *ClinicalTrialMetadataCreate) Exec(ctx context.Context) error {
	_, err := ctmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmc *ClinicalTrialMetadataCreate) ExecX(ctx context.Context) {
	if err := ctmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctmc *ClinicalTrialMetadataCreate) defaults() {
	if _, ok := ctmc.mutation.UseCaseCode(); !ok {
		v := clinicaltrialmetadata.DefaultUseCaseCode
		ctmc.mutation.SetUseCaseCode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctmc *ClinicalTrialMetadataCreate) check() error {
	if _, ok := ctmc.mutation.TagName(); !ok {
		return &ValidationError{Name: "tag_name", err: errors.New(`models: missing required field "ClinicalTrialMetadata.tag_name"`)}
	}
	if _, ok := ctmc.mutation.TagValue(); !ok {
		return &ValidationError{Name: "tag_value", err: errors.New(`models: missing required field "ClinicalTrialMetadata.tag_value"`)}
	}
	if _, ok := ctmc.mutation.UseCaseCode(); !ok {
		return &ValidationError{Name: "use_case_code", err: errors.New(`models: missing required field "ClinicalTrialMetadata.use_case_code"`)}
	}
	return nil
}

func (ctmc *ClinicalTrialMetadataCreate) sqlSave(ctx context.Context) (*ClinicalTrialMetadata, error) {
	_node, _spec := ctmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctmc *ClinicalTrialMetadataCreate) createSpec() (*ClinicalTrialMetadata, *sqlgraph.CreateSpec) {
	var (
		_node = &ClinicalTrialMetadata{config: ctmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clinicaltrialmetadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrialmetadata.FieldID,
			},
		}
	)
	if value, ok := ctmc.mutation.TagName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagName,
		})
		_node.TagName = value
	}
	if value, ok := ctmc.mutation.TagValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagValue,
		})
		_node.TagValue = value
	}
	if value, ok := ctmc.mutation.UseCaseCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldUseCaseCode,
		})
		_node.UseCaseCode = value
	}
	if nodes := ctmc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicaltrialmetadata.ParentTable,
			Columns: []string{clinicaltrialmetadata.ParentColumn},
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
		_node.clinical_trial_metadata = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClinicalTrialMetadataCreateBulk is the builder for creating many ClinicalTrialMetadata entities in bulk.
type ClinicalTrialMetadataCreateBulk struct {
	config
	builders []*ClinicalTrialMetadataCreate
}

// Save creates the ClinicalTrialMetadata entities in the database.
func (ctmcb *ClinicalTrialMetadataCreateBulk) Save(ctx context.Context) ([]*ClinicalTrialMetadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctmcb.builders))
	nodes := make([]*ClinicalTrialMetadata, len(ctmcb.builders))
	mutators := make([]Mutator, len(ctmcb.builders))
	for i := range ctmcb.builders {
		func(i int, root context.Context) {
			builder := ctmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClinicalTrialMetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, ctmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctmcb *ClinicalTrialMetadataCreateBulk) SaveX(ctx context.Context) []*ClinicalTrialMetadata {
	v, err := ctmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctmcb *ClinicalTrialMetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := ctmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmcb *ClinicalTrialMetadataCreateBulk) ExecX(ctx context.Context) {
	if err := ctmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
