// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// OutcomeMeasuresModuleCreate is the builder for creating a OutcomeMeasuresModule entity.
type OutcomeMeasuresModuleCreate struct {
	config
	mutation *OutcomeMeasuresModuleMutation
	hooks    []Hook
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (ommc *OutcomeMeasuresModuleCreate) SetParentID(id int) *OutcomeMeasuresModuleCreate {
	ommc.mutation.SetParentID(id)
	return ommc
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ommc *OutcomeMeasuresModuleCreate) SetNillableParentID(id *int) *OutcomeMeasuresModuleCreate {
	if id != nil {
		ommc = ommc.SetParentID(*id)
	}
	return ommc
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (ommc *OutcomeMeasuresModuleCreate) SetParent(r *ResultsDefinition) *OutcomeMeasuresModuleCreate {
	return ommc.SetParentID(r.ID)
}

// AddOutcomeMeasureListIDs adds the "outcome_measure_list" edge to the OutcomeMeasure entity by IDs.
func (ommc *OutcomeMeasuresModuleCreate) AddOutcomeMeasureListIDs(ids ...int) *OutcomeMeasuresModuleCreate {
	ommc.mutation.AddOutcomeMeasureListIDs(ids...)
	return ommc
}

// AddOutcomeMeasureList adds the "outcome_measure_list" edges to the OutcomeMeasure entity.
func (ommc *OutcomeMeasuresModuleCreate) AddOutcomeMeasureList(o ...*OutcomeMeasure) *OutcomeMeasuresModuleCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ommc.AddOutcomeMeasureListIDs(ids...)
}

// Mutation returns the OutcomeMeasuresModuleMutation object of the builder.
func (ommc *OutcomeMeasuresModuleCreate) Mutation() *OutcomeMeasuresModuleMutation {
	return ommc.mutation
}

// Save creates the OutcomeMeasuresModule in the database.
func (ommc *OutcomeMeasuresModuleCreate) Save(ctx context.Context) (*OutcomeMeasuresModule, error) {
	var (
		err  error
		node *OutcomeMeasuresModule
	)
	if len(ommc.hooks) == 0 {
		if err = ommc.check(); err != nil {
			return nil, err
		}
		node, err = ommc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasuresModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ommc.check(); err != nil {
				return nil, err
			}
			ommc.mutation = mutation
			if node, err = ommc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ommc.hooks) - 1; i >= 0; i-- {
			if ommc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ommc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ommc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ommc *OutcomeMeasuresModuleCreate) SaveX(ctx context.Context) *OutcomeMeasuresModule {
	v, err := ommc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ommc *OutcomeMeasuresModuleCreate) Exec(ctx context.Context) error {
	_, err := ommc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ommc *OutcomeMeasuresModuleCreate) ExecX(ctx context.Context) {
	if err := ommc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ommc *OutcomeMeasuresModuleCreate) check() error {
	return nil
}

func (ommc *OutcomeMeasuresModuleCreate) sqlSave(ctx context.Context) (*OutcomeMeasuresModule, error) {
	_node, _spec := ommc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ommc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ommc *OutcomeMeasuresModuleCreate) createSpec() (*OutcomeMeasuresModule, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeMeasuresModule{config: ommc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomemeasuresmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasuresmodule.FieldID,
			},
		}
	)
	if nodes := ommc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   outcomemeasuresmodule.ParentTable,
			Columns: []string{outcomemeasuresmodule.ParentColumn},
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
		_node.results_definition_outcome_measures_module = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ommc.mutation.OutcomeMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasuresmodule.OutcomeMeasureListTable,
			Columns: []string{outcomemeasuresmodule.OutcomeMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeMeasuresModuleCreateBulk is the builder for creating many OutcomeMeasuresModule entities in bulk.
type OutcomeMeasuresModuleCreateBulk struct {
	config
	builders []*OutcomeMeasuresModuleCreate
}

// Save creates the OutcomeMeasuresModule entities in the database.
func (ommcb *OutcomeMeasuresModuleCreateBulk) Save(ctx context.Context) ([]*OutcomeMeasuresModule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ommcb.builders))
	nodes := make([]*OutcomeMeasuresModule, len(ommcb.builders))
	mutators := make([]Mutator, len(ommcb.builders))
	for i := range ommcb.builders {
		func(i int, root context.Context) {
			builder := ommcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeMeasuresModuleMutation)
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
					_, err = mutators[i+1].Mutate(root, ommcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ommcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ommcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ommcb *OutcomeMeasuresModuleCreateBulk) SaveX(ctx context.Context) []*OutcomeMeasuresModule {
	v, err := ommcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ommcb *OutcomeMeasuresModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := ommcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ommcb *OutcomeMeasuresModuleCreateBulk) ExecX(ctx context.Context) {
	if err := ommcb.Exec(ctx); err != nil {
		panic(err)
	}
}
