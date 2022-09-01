// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
)

// BaselineMeasureCreate is the builder for creating a BaselineMeasure entity.
type BaselineMeasureCreate struct {
	config
	mutation *BaselineMeasureMutation
	hooks    []Hook
}

// SetBaselineMeasureTitle sets the "baseline_measure_title" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureTitle(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureTitle(s)
	return bmc
}

// SetBaselineMeasureDescription sets the "baseline_measure_description" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureDescription(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureDescription(s)
	return bmc
}

// SetBaselineMeasurePopulationDescription sets the "baseline_measure_population_description" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasurePopulationDescription(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasurePopulationDescription(s)
	return bmc
}

// SetBaselineMeasureParamType sets the "baseline_measure_param_type" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureParamType(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureParamType(s)
	return bmc
}

// SetBaselineMeasureDispersionType sets the "baseline_measure_dispersion_type" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureDispersionType(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureDispersionType(s)
	return bmc
}

// SetBaselineMeasureUnitOfMeasure sets the "baseline_measure_unit_of_measure" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureUnitOfMeasure(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureUnitOfMeasure(s)
	return bmc
}

// SetBaselineMeasureCalculatePct sets the "baseline_measure_calculate_pct" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureCalculatePct(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureCalculatePct(s)
	return bmc
}

// SetBaselineMeasureDenomUnitsSelected sets the "baseline_measure_denom_units_selected" field.
func (bmc *BaselineMeasureCreate) SetBaselineMeasureDenomUnitsSelected(s string) *BaselineMeasureCreate {
	bmc.mutation.SetBaselineMeasureDenomUnitsSelected(s)
	return bmc
}

// SetParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID.
func (bmc *BaselineMeasureCreate) SetParentID(id int) *BaselineMeasureCreate {
	bmc.mutation.SetParentID(id)
	return bmc
}

// SetNillableParentID sets the "parent" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (bmc *BaselineMeasureCreate) SetNillableParentID(id *int) *BaselineMeasureCreate {
	if id != nil {
		bmc = bmc.SetParentID(*id)
	}
	return bmc
}

// SetParent sets the "parent" edge to the BaselineCharacteristicsModule entity.
func (bmc *BaselineMeasureCreate) SetParent(b *BaselineCharacteristicsModule) *BaselineMeasureCreate {
	return bmc.SetParentID(b.ID)
}

// AddBaselineMeasureDenomListIDs adds the "baseline_measure_denom_list" edge to the BaselineMeasureDenom entity by IDs.
func (bmc *BaselineMeasureCreate) AddBaselineMeasureDenomListIDs(ids ...int) *BaselineMeasureCreate {
	bmc.mutation.AddBaselineMeasureDenomListIDs(ids...)
	return bmc
}

// AddBaselineMeasureDenomList adds the "baseline_measure_denom_list" edges to the BaselineMeasureDenom entity.
func (bmc *BaselineMeasureCreate) AddBaselineMeasureDenomList(b ...*BaselineMeasureDenom) *BaselineMeasureCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmc.AddBaselineMeasureDenomListIDs(ids...)
}

// AddBaselineClassListIDs adds the "baseline_class_list" edge to the BaselineClass entity by IDs.
func (bmc *BaselineMeasureCreate) AddBaselineClassListIDs(ids ...int) *BaselineMeasureCreate {
	bmc.mutation.AddBaselineClassListIDs(ids...)
	return bmc
}

// AddBaselineClassList adds the "baseline_class_list" edges to the BaselineClass entity.
func (bmc *BaselineMeasureCreate) AddBaselineClassList(b ...*BaselineClass) *BaselineMeasureCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bmc.AddBaselineClassListIDs(ids...)
}

// Mutation returns the BaselineMeasureMutation object of the builder.
func (bmc *BaselineMeasureCreate) Mutation() *BaselineMeasureMutation {
	return bmc.mutation
}

// Save creates the BaselineMeasure in the database.
func (bmc *BaselineMeasureCreate) Save(ctx context.Context) (*BaselineMeasure, error) {
	var (
		err  error
		node *BaselineMeasure
	)
	if len(bmc.hooks) == 0 {
		if err = bmc.check(); err != nil {
			return nil, err
		}
		node, err = bmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bmc.check(); err != nil {
				return nil, err
			}
			bmc.mutation = mutation
			if node, err = bmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bmc.hooks) - 1; i >= 0; i-- {
			if bmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bmc *BaselineMeasureCreate) SaveX(ctx context.Context) *BaselineMeasure {
	v, err := bmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmc *BaselineMeasureCreate) Exec(ctx context.Context) error {
	_, err := bmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmc *BaselineMeasureCreate) ExecX(ctx context.Context) {
	if err := bmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bmc *BaselineMeasureCreate) check() error {
	if _, ok := bmc.mutation.BaselineMeasureTitle(); !ok {
		return &ValidationError{Name: "baseline_measure_title", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_title"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureDescription(); !ok {
		return &ValidationError{Name: "baseline_measure_description", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_description"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasurePopulationDescription(); !ok {
		return &ValidationError{Name: "baseline_measure_population_description", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_population_description"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureParamType(); !ok {
		return &ValidationError{Name: "baseline_measure_param_type", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_param_type"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureDispersionType(); !ok {
		return &ValidationError{Name: "baseline_measure_dispersion_type", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_dispersion_type"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureUnitOfMeasure(); !ok {
		return &ValidationError{Name: "baseline_measure_unit_of_measure", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_unit_of_measure"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureCalculatePct(); !ok {
		return &ValidationError{Name: "baseline_measure_calculate_pct", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_calculate_pct"`)}
	}
	if _, ok := bmc.mutation.BaselineMeasureDenomUnitsSelected(); !ok {
		return &ValidationError{Name: "baseline_measure_denom_units_selected", err: errors.New(`models: missing required field "BaselineMeasure.baseline_measure_denom_units_selected"`)}
	}
	return nil
}

func (bmc *BaselineMeasureCreate) sqlSave(ctx context.Context) (*BaselineMeasure, error) {
	_node, _spec := bmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bmc *BaselineMeasureCreate) createSpec() (*BaselineMeasure, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineMeasure{config: bmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinemeasure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasure.FieldID,
			},
		}
	)
	if value, ok := bmc.mutation.BaselineMeasureTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureTitle,
		})
		_node.BaselineMeasureTitle = value
	}
	if value, ok := bmc.mutation.BaselineMeasureDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDescription,
		})
		_node.BaselineMeasureDescription = value
	}
	if value, ok := bmc.mutation.BaselineMeasurePopulationDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasurePopulationDescription,
		})
		_node.BaselineMeasurePopulationDescription = value
	}
	if value, ok := bmc.mutation.BaselineMeasureParamType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureParamType,
		})
		_node.BaselineMeasureParamType = value
	}
	if value, ok := bmc.mutation.BaselineMeasureDispersionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDispersionType,
		})
		_node.BaselineMeasureDispersionType = value
	}
	if value, ok := bmc.mutation.BaselineMeasureUnitOfMeasure(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureUnitOfMeasure,
		})
		_node.BaselineMeasureUnitOfMeasure = value
	}
	if value, ok := bmc.mutation.BaselineMeasureCalculatePct(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureCalculatePct,
		})
		_node.BaselineMeasureCalculatePct = value
	}
	if value, ok := bmc.mutation.BaselineMeasureDenomUnitsSelected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinemeasure.FieldBaselineMeasureDenomUnitsSelected,
		})
		_node.BaselineMeasureDenomUnitsSelected = value
	}
	if nodes := bmc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinemeasure.ParentTable,
			Columns: []string{baselinemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_characteristics_module_baseline_measure_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bmc.mutation.BaselineMeasureDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineMeasureDenomListTable,
			Columns: []string{baselinemeasure.BaselineMeasureDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasuredenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bmc.mutation.BaselineClassListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinemeasure.BaselineClassListTable,
			Columns: []string{baselinemeasure.BaselineClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
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

// BaselineMeasureCreateBulk is the builder for creating many BaselineMeasure entities in bulk.
type BaselineMeasureCreateBulk struct {
	config
	builders []*BaselineMeasureCreate
}

// Save creates the BaselineMeasure entities in the database.
func (bmcb *BaselineMeasureCreateBulk) Save(ctx context.Context) ([]*BaselineMeasure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bmcb.builders))
	nodes := make([]*BaselineMeasure, len(bmcb.builders))
	mutators := make([]Mutator, len(bmcb.builders))
	for i := range bmcb.builders {
		func(i int, root context.Context) {
			builder := bmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineMeasureMutation)
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
					_, err = mutators[i+1].Mutate(root, bmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bmcb *BaselineMeasureCreateBulk) SaveX(ctx context.Context) []*BaselineMeasure {
	v, err := bmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmcb *BaselineMeasureCreateBulk) Exec(ctx context.Context) error {
	_, err := bmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmcb *BaselineMeasureCreateBulk) ExecX(ctx context.Context) {
	if err := bmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
