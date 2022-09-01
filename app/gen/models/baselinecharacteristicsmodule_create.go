// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// BaselineCharacteristicsModuleCreate is the builder for creating a BaselineCharacteristicsModule entity.
type BaselineCharacteristicsModuleCreate struct {
	config
	mutation *BaselineCharacteristicsModuleMutation
	hooks    []Hook
}

// SetBaselinePopulationDescription sets the "baseline_population_description" field.
func (bcmc *BaselineCharacteristicsModuleCreate) SetBaselinePopulationDescription(s string) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.SetBaselinePopulationDescription(s)
	return bcmc
}

// SetBaselineTypeUnitsAnalyzed sets the "baseline_type_units_analyzed" field.
func (bcmc *BaselineCharacteristicsModuleCreate) SetBaselineTypeUnitsAnalyzed(s string) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.SetBaselineTypeUnitsAnalyzed(s)
	return bcmc
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (bcmc *BaselineCharacteristicsModuleCreate) SetParentID(id int) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.SetParentID(id)
	return bcmc
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (bcmc *BaselineCharacteristicsModuleCreate) SetNillableParentID(id *int) *BaselineCharacteristicsModuleCreate {
	if id != nil {
		bcmc = bcmc.SetParentID(*id)
	}
	return bcmc
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (bcmc *BaselineCharacteristicsModuleCreate) SetParent(r *ResultsDefinition) *BaselineCharacteristicsModuleCreate {
	return bcmc.SetParentID(r.ID)
}

// AddBaselineGroupListIDs adds the "baseline_group_list" edge to the BaselineGroup entity by IDs.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineGroupListIDs(ids ...int) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.AddBaselineGroupListIDs(ids...)
	return bcmc
}

// AddBaselineGroupList adds the "baseline_group_list" edges to the BaselineGroup entity.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineGroupList(b ...*BaselineGroup) *BaselineCharacteristicsModuleCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmc.AddBaselineGroupListIDs(ids...)
}

// AddBaselineDenomListIDs adds the "baseline_denom_list" edge to the BaselineDenom entity by IDs.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineDenomListIDs(ids ...int) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.AddBaselineDenomListIDs(ids...)
	return bcmc
}

// AddBaselineDenomList adds the "baseline_denom_list" edges to the BaselineDenom entity.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineDenomList(b ...*BaselineDenom) *BaselineCharacteristicsModuleCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmc.AddBaselineDenomListIDs(ids...)
}

// AddBaselineMeasureListIDs adds the "baseline_measure_list" edge to the BaselineMeasure entity by IDs.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineMeasureListIDs(ids ...int) *BaselineCharacteristicsModuleCreate {
	bcmc.mutation.AddBaselineMeasureListIDs(ids...)
	return bcmc
}

// AddBaselineMeasureList adds the "baseline_measure_list" edges to the BaselineMeasure entity.
func (bcmc *BaselineCharacteristicsModuleCreate) AddBaselineMeasureList(b ...*BaselineMeasure) *BaselineCharacteristicsModuleCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcmc.AddBaselineMeasureListIDs(ids...)
}

// Mutation returns the BaselineCharacteristicsModuleMutation object of the builder.
func (bcmc *BaselineCharacteristicsModuleCreate) Mutation() *BaselineCharacteristicsModuleMutation {
	return bcmc.mutation
}

// Save creates the BaselineCharacteristicsModule in the database.
func (bcmc *BaselineCharacteristicsModuleCreate) Save(ctx context.Context) (*BaselineCharacteristicsModule, error) {
	var (
		err  error
		node *BaselineCharacteristicsModule
	)
	if len(bcmc.hooks) == 0 {
		if err = bcmc.check(); err != nil {
			return nil, err
		}
		node, err = bcmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCharacteristicsModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcmc.check(); err != nil {
				return nil, err
			}
			bcmc.mutation = mutation
			if node, err = bcmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcmc.hooks) - 1; i >= 0; i-- {
			if bcmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcmc *BaselineCharacteristicsModuleCreate) SaveX(ctx context.Context) *BaselineCharacteristicsModule {
	v, err := bcmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcmc *BaselineCharacteristicsModuleCreate) Exec(ctx context.Context) error {
	_, err := bcmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmc *BaselineCharacteristicsModuleCreate) ExecX(ctx context.Context) {
	if err := bcmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcmc *BaselineCharacteristicsModuleCreate) check() error {
	if _, ok := bcmc.mutation.BaselinePopulationDescription(); !ok {
		return &ValidationError{Name: "baseline_population_description", err: errors.New(`models: missing required field "BaselineCharacteristicsModule.baseline_population_description"`)}
	}
	if _, ok := bcmc.mutation.BaselineTypeUnitsAnalyzed(); !ok {
		return &ValidationError{Name: "baseline_type_units_analyzed", err: errors.New(`models: missing required field "BaselineCharacteristicsModule.baseline_type_units_analyzed"`)}
	}
	return nil
}

func (bcmc *BaselineCharacteristicsModuleCreate) sqlSave(ctx context.Context) (*BaselineCharacteristicsModule, error) {
	_node, _spec := bcmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bcmc *BaselineCharacteristicsModuleCreate) createSpec() (*BaselineCharacteristicsModule, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineCharacteristicsModule{config: bcmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinecharacteristicsmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecharacteristicsmodule.FieldID,
			},
		}
	)
	if value, ok := bcmc.mutation.BaselinePopulationDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselinePopulationDescription,
		})
		_node.BaselinePopulationDescription = value
	}
	if value, ok := bcmc.mutation.BaselineTypeUnitsAnalyzed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecharacteristicsmodule.FieldBaselineTypeUnitsAnalyzed,
		})
		_node.BaselineTypeUnitsAnalyzed = value
	}
	if nodes := bcmc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   baselinecharacteristicsmodule.ParentTable,
			Columns: []string{baselinecharacteristicsmodule.ParentColumn},
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
		_node.results_definition_baseline_characteristics_module = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcmc.mutation.BaselineGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineGroupListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcmc.mutation.BaselineDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineDenomListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcmc.mutation.BaselineMeasureListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecharacteristicsmodule.BaselineMeasureListTable,
			Columns: []string{baselinecharacteristicsmodule.BaselineMeasureListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
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

// BaselineCharacteristicsModuleCreateBulk is the builder for creating many BaselineCharacteristicsModule entities in bulk.
type BaselineCharacteristicsModuleCreateBulk struct {
	config
	builders []*BaselineCharacteristicsModuleCreate
}

// Save creates the BaselineCharacteristicsModule entities in the database.
func (bcmcb *BaselineCharacteristicsModuleCreateBulk) Save(ctx context.Context) ([]*BaselineCharacteristicsModule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcmcb.builders))
	nodes := make([]*BaselineCharacteristicsModule, len(bcmcb.builders))
	mutators := make([]Mutator, len(bcmcb.builders))
	for i := range bcmcb.builders {
		func(i int, root context.Context) {
			builder := bcmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineCharacteristicsModuleMutation)
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
					_, err = mutators[i+1].Mutate(root, bcmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcmcb *BaselineCharacteristicsModuleCreateBulk) SaveX(ctx context.Context) []*BaselineCharacteristicsModule {
	v, err := bcmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcmcb *BaselineCharacteristicsModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := bcmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcmcb *BaselineCharacteristicsModuleCreateBulk) ExecX(ctx context.Context) {
	if err := bcmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
