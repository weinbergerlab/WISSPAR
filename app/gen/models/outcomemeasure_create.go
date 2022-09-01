// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
)

// OutcomeMeasureCreate is the builder for creating a OutcomeMeasure entity.
type OutcomeMeasureCreate struct {
	config
	mutation *OutcomeMeasureMutation
	hooks    []Hook
}

// SetOutcomeMeasureType sets the "outcome_measure_type" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureType(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureType(s)
	return omc
}

// SetOutcomeMeasureTitle sets the "outcome_measure_title" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureTitle(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureTitle(s)
	return omc
}

// SetOutcomeMeasureDescription sets the "outcome_measure_description" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureDescription(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureDescription(s)
	return omc
}

// SetOutcomeMeasurePopulationDescription sets the "outcome_measure_population_description" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasurePopulationDescription(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasurePopulationDescription(s)
	return omc
}

// SetOutcomeMeasureReportingStatus sets the "outcome_measure_reporting_status" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureReportingStatus(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureReportingStatus(s)
	return omc
}

// SetOutcomeMeasureAnticipatedPostingDate sets the "outcome_measure_anticipated_posting_date" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureAnticipatedPostingDate(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureAnticipatedPostingDate(s)
	return omc
}

// SetOutcomeMeasureParamType sets the "outcome_measure_param_type" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureParamType(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureParamType(s)
	return omc
}

// SetOutcomeMeasureDispersionType sets the "outcome_measure_dispersion_type" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureDispersionType(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureDispersionType(s)
	return omc
}

// SetOutcomeMeasureUnitOfMeasure sets the "outcome_measure_unit_of_measure" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureUnitOfMeasure(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureUnitOfMeasure(s)
	return omc
}

// SetOutcomeMeasureCalculatePct sets the "outcome_measure_calculate_pct" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureCalculatePct(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureCalculatePct(s)
	return omc
}

// SetOutcomeMeasureTimeFrame sets the "outcome_measure_time_frame" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureTimeFrame(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureTimeFrame(s)
	return omc
}

// SetOutcomeMeasureTypeUnitsAnalyzed sets the "outcome_measure_type_units_analyzed" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureTypeUnitsAnalyzed(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureTypeUnitsAnalyzed(s)
	return omc
}

// SetOutcomeMeasureDenomUnitsSelected sets the "outcome_measure_denom_units_selected" field.
func (omc *OutcomeMeasureCreate) SetOutcomeMeasureDenomUnitsSelected(s string) *OutcomeMeasureCreate {
	omc.mutation.SetOutcomeMeasureDenomUnitsSelected(s)
	return omc
}

// SetParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID.
func (omc *OutcomeMeasureCreate) SetParentID(id int) *OutcomeMeasureCreate {
	omc.mutation.SetParentID(id)
	return omc
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (omc *OutcomeMeasureCreate) SetNillableParentID(id *int) *OutcomeMeasureCreate {
	if id != nil {
		omc = omc.SetParentID(*id)
	}
	return omc
}

// SetParent sets the "parent" edge to the OutcomeMeasuresModule entity.
func (omc *OutcomeMeasureCreate) SetParent(o *OutcomeMeasuresModule) *OutcomeMeasureCreate {
	return omc.SetParentID(o.ID)
}

// AddOutcomeGroupListIDs adds the "outcome_group_list" edge to the OutcomeGroup entity by IDs.
func (omc *OutcomeMeasureCreate) AddOutcomeGroupListIDs(ids ...int) *OutcomeMeasureCreate {
	omc.mutation.AddOutcomeGroupListIDs(ids...)
	return omc
}

// AddOutcomeGroupList adds the "outcome_group_list" edges to the OutcomeGroup entity.
func (omc *OutcomeMeasureCreate) AddOutcomeGroupList(o ...*OutcomeGroup) *OutcomeMeasureCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omc.AddOutcomeGroupListIDs(ids...)
}

// AddOutcomeOverviewListIDs adds the "outcome_overview_list" edge to the OutcomeOverview entity by IDs.
func (omc *OutcomeMeasureCreate) AddOutcomeOverviewListIDs(ids ...int) *OutcomeMeasureCreate {
	omc.mutation.AddOutcomeOverviewListIDs(ids...)
	return omc
}

// AddOutcomeOverviewList adds the "outcome_overview_list" edges to the OutcomeOverview entity.
func (omc *OutcomeMeasureCreate) AddOutcomeOverviewList(o ...*OutcomeOverview) *OutcomeMeasureCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omc.AddOutcomeOverviewListIDs(ids...)
}

// AddOutcomeDenomListIDs adds the "outcome_denom_list" edge to the OutcomeDenom entity by IDs.
func (omc *OutcomeMeasureCreate) AddOutcomeDenomListIDs(ids ...int) *OutcomeMeasureCreate {
	omc.mutation.AddOutcomeDenomListIDs(ids...)
	return omc
}

// AddOutcomeDenomList adds the "outcome_denom_list" edges to the OutcomeDenom entity.
func (omc *OutcomeMeasureCreate) AddOutcomeDenomList(o ...*OutcomeDenom) *OutcomeMeasureCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omc.AddOutcomeDenomListIDs(ids...)
}

// AddOutcomeClassListIDs adds the "outcome_class_list" edge to the OutcomeClass entity by IDs.
func (omc *OutcomeMeasureCreate) AddOutcomeClassListIDs(ids ...int) *OutcomeMeasureCreate {
	omc.mutation.AddOutcomeClassListIDs(ids...)
	return omc
}

// AddOutcomeClassList adds the "outcome_class_list" edges to the OutcomeClass entity.
func (omc *OutcomeMeasureCreate) AddOutcomeClassList(o ...*OutcomeClass) *OutcomeMeasureCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omc.AddOutcomeClassListIDs(ids...)
}

// AddOutcomeAnalysisListIDs adds the "outcome_analysis_list" edge to the OutcomeAnalysis entity by IDs.
func (omc *OutcomeMeasureCreate) AddOutcomeAnalysisListIDs(ids ...int) *OutcomeMeasureCreate {
	omc.mutation.AddOutcomeAnalysisListIDs(ids...)
	return omc
}

// AddOutcomeAnalysisList adds the "outcome_analysis_list" edges to the OutcomeAnalysis entity.
func (omc *OutcomeMeasureCreate) AddOutcomeAnalysisList(o ...*OutcomeAnalysis) *OutcomeMeasureCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omc.AddOutcomeAnalysisListIDs(ids...)
}

// Mutation returns the OutcomeMeasureMutation object of the builder.
func (omc *OutcomeMeasureCreate) Mutation() *OutcomeMeasureMutation {
	return omc.mutation
}

// Save creates the OutcomeMeasure in the database.
func (omc *OutcomeMeasureCreate) Save(ctx context.Context) (*OutcomeMeasure, error) {
	var (
		err  error
		node *OutcomeMeasure
	)
	if len(omc.hooks) == 0 {
		if err = omc.check(); err != nil {
			return nil, err
		}
		node, err = omc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = omc.check(); err != nil {
				return nil, err
			}
			omc.mutation = mutation
			if node, err = omc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(omc.hooks) - 1; i >= 0; i-- {
			if omc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (omc *OutcomeMeasureCreate) SaveX(ctx context.Context) *OutcomeMeasure {
	v, err := omc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omc *OutcomeMeasureCreate) Exec(ctx context.Context) error {
	_, err := omc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omc *OutcomeMeasureCreate) ExecX(ctx context.Context) {
	if err := omc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (omc *OutcomeMeasureCreate) check() error {
	if _, ok := omc.mutation.OutcomeMeasureType(); !ok {
		return &ValidationError{Name: "outcome_measure_type", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_type"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureTitle(); !ok {
		return &ValidationError{Name: "outcome_measure_title", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_title"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureDescription(); !ok {
		return &ValidationError{Name: "outcome_measure_description", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_description"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasurePopulationDescription(); !ok {
		return &ValidationError{Name: "outcome_measure_population_description", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_population_description"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureReportingStatus(); !ok {
		return &ValidationError{Name: "outcome_measure_reporting_status", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_reporting_status"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureAnticipatedPostingDate(); !ok {
		return &ValidationError{Name: "outcome_measure_anticipated_posting_date", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_anticipated_posting_date"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureParamType(); !ok {
		return &ValidationError{Name: "outcome_measure_param_type", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_param_type"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureDispersionType(); !ok {
		return &ValidationError{Name: "outcome_measure_dispersion_type", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_dispersion_type"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureUnitOfMeasure(); !ok {
		return &ValidationError{Name: "outcome_measure_unit_of_measure", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_unit_of_measure"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureCalculatePct(); !ok {
		return &ValidationError{Name: "outcome_measure_calculate_pct", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_calculate_pct"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureTimeFrame(); !ok {
		return &ValidationError{Name: "outcome_measure_time_frame", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_time_frame"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureTypeUnitsAnalyzed(); !ok {
		return &ValidationError{Name: "outcome_measure_type_units_analyzed", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_type_units_analyzed"`)}
	}
	if _, ok := omc.mutation.OutcomeMeasureDenomUnitsSelected(); !ok {
		return &ValidationError{Name: "outcome_measure_denom_units_selected", err: errors.New(`models: missing required field "OutcomeMeasure.outcome_measure_denom_units_selected"`)}
	}
	return nil
}

func (omc *OutcomeMeasureCreate) sqlSave(ctx context.Context) (*OutcomeMeasure, error) {
	_node, _spec := omc.createSpec()
	if err := sqlgraph.CreateNode(ctx, omc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (omc *OutcomeMeasureCreate) createSpec() (*OutcomeMeasure, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeMeasure{config: omc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomemeasure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasure.FieldID,
			},
		}
	)
	if value, ok := omc.mutation.OutcomeMeasureType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureType,
		})
		_node.OutcomeMeasureType = value
	}
	if value, ok := omc.mutation.OutcomeMeasureTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTitle,
		})
		_node.OutcomeMeasureTitle = value
	}
	if value, ok := omc.mutation.OutcomeMeasureDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDescription,
		})
		_node.OutcomeMeasureDescription = value
	}
	if value, ok := omc.mutation.OutcomeMeasurePopulationDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasurePopulationDescription,
		})
		_node.OutcomeMeasurePopulationDescription = value
	}
	if value, ok := omc.mutation.OutcomeMeasureReportingStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureReportingStatus,
		})
		_node.OutcomeMeasureReportingStatus = value
	}
	if value, ok := omc.mutation.OutcomeMeasureAnticipatedPostingDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureAnticipatedPostingDate,
		})
		_node.OutcomeMeasureAnticipatedPostingDate = value
	}
	if value, ok := omc.mutation.OutcomeMeasureParamType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureParamType,
		})
		_node.OutcomeMeasureParamType = value
	}
	if value, ok := omc.mutation.OutcomeMeasureDispersionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDispersionType,
		})
		_node.OutcomeMeasureDispersionType = value
	}
	if value, ok := omc.mutation.OutcomeMeasureUnitOfMeasure(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureUnitOfMeasure,
		})
		_node.OutcomeMeasureUnitOfMeasure = value
	}
	if value, ok := omc.mutation.OutcomeMeasureCalculatePct(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureCalculatePct,
		})
		_node.OutcomeMeasureCalculatePct = value
	}
	if value, ok := omc.mutation.OutcomeMeasureTimeFrame(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTimeFrame,
		})
		_node.OutcomeMeasureTimeFrame = value
	}
	if value, ok := omc.mutation.OutcomeMeasureTypeUnitsAnalyzed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTypeUnitsAnalyzed,
		})
		_node.OutcomeMeasureTypeUnitsAnalyzed = value
	}
	if value, ok := omc.mutation.OutcomeMeasureDenomUnitsSelected(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDenomUnitsSelected,
		})
		_node.OutcomeMeasureDenomUnitsSelected = value
	}
	if nodes := omc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomemeasure.ParentTable,
			Columns: []string{outcomemeasure.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasuresmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_measures_module_outcome_measure_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := omc.mutation.OutcomeGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasure.OutcomeGroupListTable,
			Columns: []string{outcomemeasure.OutcomeGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomegroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := omc.mutation.OutcomeOverviewListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasure.OutcomeOverviewListTable,
			Columns: []string{outcomemeasure.OutcomeOverviewListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeoverview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := omc.mutation.OutcomeDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasure.OutcomeDenomListTable,
			Columns: []string{outcomemeasure.OutcomeDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomedenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := omc.mutation.OutcomeClassListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasure.OutcomeClassListTable,
			Columns: []string{outcomemeasure.OutcomeClassListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := omc.mutation.OutcomeAnalysisListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomemeasure.OutcomeAnalysisListTable,
			Columns: []string{outcomemeasure.OutcomeAnalysisListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
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

// OutcomeMeasureCreateBulk is the builder for creating many OutcomeMeasure entities in bulk.
type OutcomeMeasureCreateBulk struct {
	config
	builders []*OutcomeMeasureCreate
}

// Save creates the OutcomeMeasure entities in the database.
func (omcb *OutcomeMeasureCreateBulk) Save(ctx context.Context) ([]*OutcomeMeasure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(omcb.builders))
	nodes := make([]*OutcomeMeasure, len(omcb.builders))
	mutators := make([]Mutator, len(omcb.builders))
	for i := range omcb.builders {
		func(i int, root context.Context) {
			builder := omcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeMeasureMutation)
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
					_, err = mutators[i+1].Mutate(root, omcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, omcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, omcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (omcb *OutcomeMeasureCreateBulk) SaveX(ctx context.Context) []*OutcomeMeasure {
	v, err := omcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omcb *OutcomeMeasureCreateBulk) Exec(ctx context.Context) error {
	_, err := omcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omcb *OutcomeMeasureCreateBulk) ExecX(ctx context.Context) {
	if err := omcb.Exec(ctx); err != nil {
		panic(err)
	}
}
