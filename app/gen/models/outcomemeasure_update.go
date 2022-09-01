// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasureUpdate is the builder for updating OutcomeMeasure entities.
type OutcomeMeasureUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeMeasureMutation
}

// Where appends a list predicates to the OutcomeMeasureUpdate builder.
func (omu *OutcomeMeasureUpdate) Where(ps ...predicate.OutcomeMeasure) *OutcomeMeasureUpdate {
	omu.mutation.Where(ps...)
	return omu
}

// SetOutcomeMeasureType sets the "outcome_measure_type" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureType(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureType(s)
	return omu
}

// SetOutcomeMeasureTitle sets the "outcome_measure_title" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureTitle(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureTitle(s)
	return omu
}

// SetOutcomeMeasureDescription sets the "outcome_measure_description" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureDescription(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureDescription(s)
	return omu
}

// SetOutcomeMeasurePopulationDescription sets the "outcome_measure_population_description" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasurePopulationDescription(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasurePopulationDescription(s)
	return omu
}

// SetOutcomeMeasureReportingStatus sets the "outcome_measure_reporting_status" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureReportingStatus(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureReportingStatus(s)
	return omu
}

// SetOutcomeMeasureAnticipatedPostingDate sets the "outcome_measure_anticipated_posting_date" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureAnticipatedPostingDate(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureAnticipatedPostingDate(s)
	return omu
}

// SetOutcomeMeasureParamType sets the "outcome_measure_param_type" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureParamType(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureParamType(s)
	return omu
}

// SetOutcomeMeasureDispersionType sets the "outcome_measure_dispersion_type" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureDispersionType(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureDispersionType(s)
	return omu
}

// SetOutcomeMeasureUnitOfMeasure sets the "outcome_measure_unit_of_measure" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureUnitOfMeasure(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureUnitOfMeasure(s)
	return omu
}

// SetOutcomeMeasureCalculatePct sets the "outcome_measure_calculate_pct" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureCalculatePct(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureCalculatePct(s)
	return omu
}

// SetOutcomeMeasureTimeFrame sets the "outcome_measure_time_frame" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureTimeFrame(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureTimeFrame(s)
	return omu
}

// SetOutcomeMeasureTypeUnitsAnalyzed sets the "outcome_measure_type_units_analyzed" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureTypeUnitsAnalyzed(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureTypeUnitsAnalyzed(s)
	return omu
}

// SetOutcomeMeasureDenomUnitsSelected sets the "outcome_measure_denom_units_selected" field.
func (omu *OutcomeMeasureUpdate) SetOutcomeMeasureDenomUnitsSelected(s string) *OutcomeMeasureUpdate {
	omu.mutation.SetOutcomeMeasureDenomUnitsSelected(s)
	return omu
}

// SetParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID.
func (omu *OutcomeMeasureUpdate) SetParentID(id int) *OutcomeMeasureUpdate {
	omu.mutation.SetParentID(id)
	return omu
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (omu *OutcomeMeasureUpdate) SetNillableParentID(id *int) *OutcomeMeasureUpdate {
	if id != nil {
		omu = omu.SetParentID(*id)
	}
	return omu
}

// SetParent sets the "parent" edge to the OutcomeMeasuresModule entity.
func (omu *OutcomeMeasureUpdate) SetParent(o *OutcomeMeasuresModule) *OutcomeMeasureUpdate {
	return omu.SetParentID(o.ID)
}

// AddOutcomeGroupListIDs adds the "outcome_group_list" edge to the OutcomeGroup entity by IDs.
func (omu *OutcomeMeasureUpdate) AddOutcomeGroupListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.AddOutcomeGroupListIDs(ids...)
	return omu
}

// AddOutcomeGroupList adds the "outcome_group_list" edges to the OutcomeGroup entity.
func (omu *OutcomeMeasureUpdate) AddOutcomeGroupList(o ...*OutcomeGroup) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.AddOutcomeGroupListIDs(ids...)
}

// AddOutcomeOverviewListIDs adds the "outcome_overview_list" edge to the OutcomeOverview entity by IDs.
func (omu *OutcomeMeasureUpdate) AddOutcomeOverviewListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.AddOutcomeOverviewListIDs(ids...)
	return omu
}

// AddOutcomeOverviewList adds the "outcome_overview_list" edges to the OutcomeOverview entity.
func (omu *OutcomeMeasureUpdate) AddOutcomeOverviewList(o ...*OutcomeOverview) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.AddOutcomeOverviewListIDs(ids...)
}

// AddOutcomeDenomListIDs adds the "outcome_denom_list" edge to the OutcomeDenom entity by IDs.
func (omu *OutcomeMeasureUpdate) AddOutcomeDenomListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.AddOutcomeDenomListIDs(ids...)
	return omu
}

// AddOutcomeDenomList adds the "outcome_denom_list" edges to the OutcomeDenom entity.
func (omu *OutcomeMeasureUpdate) AddOutcomeDenomList(o ...*OutcomeDenom) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.AddOutcomeDenomListIDs(ids...)
}

// AddOutcomeClassListIDs adds the "outcome_class_list" edge to the OutcomeClass entity by IDs.
func (omu *OutcomeMeasureUpdate) AddOutcomeClassListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.AddOutcomeClassListIDs(ids...)
	return omu
}

// AddOutcomeClassList adds the "outcome_class_list" edges to the OutcomeClass entity.
func (omu *OutcomeMeasureUpdate) AddOutcomeClassList(o ...*OutcomeClass) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.AddOutcomeClassListIDs(ids...)
}

// AddOutcomeAnalysisListIDs adds the "outcome_analysis_list" edge to the OutcomeAnalysis entity by IDs.
func (omu *OutcomeMeasureUpdate) AddOutcomeAnalysisListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.AddOutcomeAnalysisListIDs(ids...)
	return omu
}

// AddOutcomeAnalysisList adds the "outcome_analysis_list" edges to the OutcomeAnalysis entity.
func (omu *OutcomeMeasureUpdate) AddOutcomeAnalysisList(o ...*OutcomeAnalysis) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.AddOutcomeAnalysisListIDs(ids...)
}

// Mutation returns the OutcomeMeasureMutation object of the builder.
func (omu *OutcomeMeasureUpdate) Mutation() *OutcomeMeasureMutation {
	return omu.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasuresModule entity.
func (omu *OutcomeMeasureUpdate) ClearParent() *OutcomeMeasureUpdate {
	omu.mutation.ClearParent()
	return omu
}

// ClearOutcomeGroupList clears all "outcome_group_list" edges to the OutcomeGroup entity.
func (omu *OutcomeMeasureUpdate) ClearOutcomeGroupList() *OutcomeMeasureUpdate {
	omu.mutation.ClearOutcomeGroupList()
	return omu
}

// RemoveOutcomeGroupListIDs removes the "outcome_group_list" edge to OutcomeGroup entities by IDs.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeGroupListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.RemoveOutcomeGroupListIDs(ids...)
	return omu
}

// RemoveOutcomeGroupList removes "outcome_group_list" edges to OutcomeGroup entities.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeGroupList(o ...*OutcomeGroup) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.RemoveOutcomeGroupListIDs(ids...)
}

// ClearOutcomeOverviewList clears all "outcome_overview_list" edges to the OutcomeOverview entity.
func (omu *OutcomeMeasureUpdate) ClearOutcomeOverviewList() *OutcomeMeasureUpdate {
	omu.mutation.ClearOutcomeOverviewList()
	return omu
}

// RemoveOutcomeOverviewListIDs removes the "outcome_overview_list" edge to OutcomeOverview entities by IDs.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeOverviewListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.RemoveOutcomeOverviewListIDs(ids...)
	return omu
}

// RemoveOutcomeOverviewList removes "outcome_overview_list" edges to OutcomeOverview entities.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeOverviewList(o ...*OutcomeOverview) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.RemoveOutcomeOverviewListIDs(ids...)
}

// ClearOutcomeDenomList clears all "outcome_denom_list" edges to the OutcomeDenom entity.
func (omu *OutcomeMeasureUpdate) ClearOutcomeDenomList() *OutcomeMeasureUpdate {
	omu.mutation.ClearOutcomeDenomList()
	return omu
}

// RemoveOutcomeDenomListIDs removes the "outcome_denom_list" edge to OutcomeDenom entities by IDs.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeDenomListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.RemoveOutcomeDenomListIDs(ids...)
	return omu
}

// RemoveOutcomeDenomList removes "outcome_denom_list" edges to OutcomeDenom entities.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeDenomList(o ...*OutcomeDenom) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.RemoveOutcomeDenomListIDs(ids...)
}

// ClearOutcomeClassList clears all "outcome_class_list" edges to the OutcomeClass entity.
func (omu *OutcomeMeasureUpdate) ClearOutcomeClassList() *OutcomeMeasureUpdate {
	omu.mutation.ClearOutcomeClassList()
	return omu
}

// RemoveOutcomeClassListIDs removes the "outcome_class_list" edge to OutcomeClass entities by IDs.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeClassListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.RemoveOutcomeClassListIDs(ids...)
	return omu
}

// RemoveOutcomeClassList removes "outcome_class_list" edges to OutcomeClass entities.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeClassList(o ...*OutcomeClass) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.RemoveOutcomeClassListIDs(ids...)
}

// ClearOutcomeAnalysisList clears all "outcome_analysis_list" edges to the OutcomeAnalysis entity.
func (omu *OutcomeMeasureUpdate) ClearOutcomeAnalysisList() *OutcomeMeasureUpdate {
	omu.mutation.ClearOutcomeAnalysisList()
	return omu
}

// RemoveOutcomeAnalysisListIDs removes the "outcome_analysis_list" edge to OutcomeAnalysis entities by IDs.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeAnalysisListIDs(ids ...int) *OutcomeMeasureUpdate {
	omu.mutation.RemoveOutcomeAnalysisListIDs(ids...)
	return omu
}

// RemoveOutcomeAnalysisList removes "outcome_analysis_list" edges to OutcomeAnalysis entities.
func (omu *OutcomeMeasureUpdate) RemoveOutcomeAnalysisList(o ...*OutcomeAnalysis) *OutcomeMeasureUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omu.RemoveOutcomeAnalysisListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (omu *OutcomeMeasureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(omu.hooks) == 0 {
		affected, err = omu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omu.mutation = mutation
			affected, err = omu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(omu.hooks) - 1; i >= 0; i-- {
			if omu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (omu *OutcomeMeasureUpdate) SaveX(ctx context.Context) int {
	affected, err := omu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (omu *OutcomeMeasureUpdate) Exec(ctx context.Context) error {
	_, err := omu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omu *OutcomeMeasureUpdate) ExecX(ctx context.Context) {
	if err := omu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (omu *OutcomeMeasureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasure.Table,
			Columns: outcomemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasure.FieldID,
			},
		},
	}
	if ps := omu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omu.mutation.OutcomeMeasureType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureType,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTitle,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDescription,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasurePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasurePopulationDescription,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureReportingStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureReportingStatus,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureAnticipatedPostingDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureAnticipatedPostingDate,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureParamType,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDispersionType,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureUnitOfMeasure(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureUnitOfMeasure,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureCalculatePct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureCalculatePct,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTimeFrame,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTypeUnitsAnalyzed,
		})
	}
	if value, ok := omu.mutation.OutcomeMeasureDenomUnitsSelected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDenomUnitsSelected,
		})
	}
	if omu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omu.mutation.OutcomeGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.RemovedOutcomeGroupListIDs(); len(nodes) > 0 && !omu.mutation.OutcomeGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.OutcomeGroupListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omu.mutation.OutcomeOverviewListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.RemovedOutcomeOverviewListIDs(); len(nodes) > 0 && !omu.mutation.OutcomeOverviewListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.OutcomeOverviewListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omu.mutation.OutcomeDenomListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.RemovedOutcomeDenomListIDs(); len(nodes) > 0 && !omu.mutation.OutcomeDenomListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.OutcomeDenomListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omu.mutation.OutcomeClassListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.RemovedOutcomeClassListIDs(); len(nodes) > 0 && !omu.mutation.OutcomeClassListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.OutcomeClassListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omu.mutation.OutcomeAnalysisListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.RemovedOutcomeAnalysisListIDs(); len(nodes) > 0 && !omu.mutation.OutcomeAnalysisListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omu.mutation.OutcomeAnalysisListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, omu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeMeasureUpdateOne is the builder for updating a single OutcomeMeasure entity.
type OutcomeMeasureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeMeasureMutation
}

// SetOutcomeMeasureType sets the "outcome_measure_type" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureType(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureType(s)
	return omuo
}

// SetOutcomeMeasureTitle sets the "outcome_measure_title" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureTitle(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureTitle(s)
	return omuo
}

// SetOutcomeMeasureDescription sets the "outcome_measure_description" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureDescription(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureDescription(s)
	return omuo
}

// SetOutcomeMeasurePopulationDescription sets the "outcome_measure_population_description" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasurePopulationDescription(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasurePopulationDescription(s)
	return omuo
}

// SetOutcomeMeasureReportingStatus sets the "outcome_measure_reporting_status" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureReportingStatus(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureReportingStatus(s)
	return omuo
}

// SetOutcomeMeasureAnticipatedPostingDate sets the "outcome_measure_anticipated_posting_date" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureAnticipatedPostingDate(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureAnticipatedPostingDate(s)
	return omuo
}

// SetOutcomeMeasureParamType sets the "outcome_measure_param_type" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureParamType(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureParamType(s)
	return omuo
}

// SetOutcomeMeasureDispersionType sets the "outcome_measure_dispersion_type" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureDispersionType(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureDispersionType(s)
	return omuo
}

// SetOutcomeMeasureUnitOfMeasure sets the "outcome_measure_unit_of_measure" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureUnitOfMeasure(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureUnitOfMeasure(s)
	return omuo
}

// SetOutcomeMeasureCalculatePct sets the "outcome_measure_calculate_pct" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureCalculatePct(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureCalculatePct(s)
	return omuo
}

// SetOutcomeMeasureTimeFrame sets the "outcome_measure_time_frame" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureTimeFrame(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureTimeFrame(s)
	return omuo
}

// SetOutcomeMeasureTypeUnitsAnalyzed sets the "outcome_measure_type_units_analyzed" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureTypeUnitsAnalyzed(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureTypeUnitsAnalyzed(s)
	return omuo
}

// SetOutcomeMeasureDenomUnitsSelected sets the "outcome_measure_denom_units_selected" field.
func (omuo *OutcomeMeasureUpdateOne) SetOutcomeMeasureDenomUnitsSelected(s string) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetOutcomeMeasureDenomUnitsSelected(s)
	return omuo
}

// SetParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID.
func (omuo *OutcomeMeasureUpdateOne) SetParentID(id int) *OutcomeMeasureUpdateOne {
	omuo.mutation.SetParentID(id)
	return omuo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (omuo *OutcomeMeasureUpdateOne) SetNillableParentID(id *int) *OutcomeMeasureUpdateOne {
	if id != nil {
		omuo = omuo.SetParentID(*id)
	}
	return omuo
}

// SetParent sets the "parent" edge to the OutcomeMeasuresModule entity.
func (omuo *OutcomeMeasureUpdateOne) SetParent(o *OutcomeMeasuresModule) *OutcomeMeasureUpdateOne {
	return omuo.SetParentID(o.ID)
}

// AddOutcomeGroupListIDs adds the "outcome_group_list" edge to the OutcomeGroup entity by IDs.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeGroupListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.AddOutcomeGroupListIDs(ids...)
	return omuo
}

// AddOutcomeGroupList adds the "outcome_group_list" edges to the OutcomeGroup entity.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeGroupList(o ...*OutcomeGroup) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.AddOutcomeGroupListIDs(ids...)
}

// AddOutcomeOverviewListIDs adds the "outcome_overview_list" edge to the OutcomeOverview entity by IDs.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeOverviewListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.AddOutcomeOverviewListIDs(ids...)
	return omuo
}

// AddOutcomeOverviewList adds the "outcome_overview_list" edges to the OutcomeOverview entity.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeOverviewList(o ...*OutcomeOverview) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.AddOutcomeOverviewListIDs(ids...)
}

// AddOutcomeDenomListIDs adds the "outcome_denom_list" edge to the OutcomeDenom entity by IDs.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeDenomListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.AddOutcomeDenomListIDs(ids...)
	return omuo
}

// AddOutcomeDenomList adds the "outcome_denom_list" edges to the OutcomeDenom entity.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeDenomList(o ...*OutcomeDenom) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.AddOutcomeDenomListIDs(ids...)
}

// AddOutcomeClassListIDs adds the "outcome_class_list" edge to the OutcomeClass entity by IDs.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeClassListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.AddOutcomeClassListIDs(ids...)
	return omuo
}

// AddOutcomeClassList adds the "outcome_class_list" edges to the OutcomeClass entity.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeClassList(o ...*OutcomeClass) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.AddOutcomeClassListIDs(ids...)
}

// AddOutcomeAnalysisListIDs adds the "outcome_analysis_list" edge to the OutcomeAnalysis entity by IDs.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeAnalysisListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.AddOutcomeAnalysisListIDs(ids...)
	return omuo
}

// AddOutcomeAnalysisList adds the "outcome_analysis_list" edges to the OutcomeAnalysis entity.
func (omuo *OutcomeMeasureUpdateOne) AddOutcomeAnalysisList(o ...*OutcomeAnalysis) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.AddOutcomeAnalysisListIDs(ids...)
}

// Mutation returns the OutcomeMeasureMutation object of the builder.
func (omuo *OutcomeMeasureUpdateOne) Mutation() *OutcomeMeasureMutation {
	return omuo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasuresModule entity.
func (omuo *OutcomeMeasureUpdateOne) ClearParent() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearParent()
	return omuo
}

// ClearOutcomeGroupList clears all "outcome_group_list" edges to the OutcomeGroup entity.
func (omuo *OutcomeMeasureUpdateOne) ClearOutcomeGroupList() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearOutcomeGroupList()
	return omuo
}

// RemoveOutcomeGroupListIDs removes the "outcome_group_list" edge to OutcomeGroup entities by IDs.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeGroupListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.RemoveOutcomeGroupListIDs(ids...)
	return omuo
}

// RemoveOutcomeGroupList removes "outcome_group_list" edges to OutcomeGroup entities.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeGroupList(o ...*OutcomeGroup) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.RemoveOutcomeGroupListIDs(ids...)
}

// ClearOutcomeOverviewList clears all "outcome_overview_list" edges to the OutcomeOverview entity.
func (omuo *OutcomeMeasureUpdateOne) ClearOutcomeOverviewList() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearOutcomeOverviewList()
	return omuo
}

// RemoveOutcomeOverviewListIDs removes the "outcome_overview_list" edge to OutcomeOverview entities by IDs.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeOverviewListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.RemoveOutcomeOverviewListIDs(ids...)
	return omuo
}

// RemoveOutcomeOverviewList removes "outcome_overview_list" edges to OutcomeOverview entities.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeOverviewList(o ...*OutcomeOverview) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.RemoveOutcomeOverviewListIDs(ids...)
}

// ClearOutcomeDenomList clears all "outcome_denom_list" edges to the OutcomeDenom entity.
func (omuo *OutcomeMeasureUpdateOne) ClearOutcomeDenomList() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearOutcomeDenomList()
	return omuo
}

// RemoveOutcomeDenomListIDs removes the "outcome_denom_list" edge to OutcomeDenom entities by IDs.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeDenomListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.RemoveOutcomeDenomListIDs(ids...)
	return omuo
}

// RemoveOutcomeDenomList removes "outcome_denom_list" edges to OutcomeDenom entities.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeDenomList(o ...*OutcomeDenom) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.RemoveOutcomeDenomListIDs(ids...)
}

// ClearOutcomeClassList clears all "outcome_class_list" edges to the OutcomeClass entity.
func (omuo *OutcomeMeasureUpdateOne) ClearOutcomeClassList() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearOutcomeClassList()
	return omuo
}

// RemoveOutcomeClassListIDs removes the "outcome_class_list" edge to OutcomeClass entities by IDs.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeClassListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.RemoveOutcomeClassListIDs(ids...)
	return omuo
}

// RemoveOutcomeClassList removes "outcome_class_list" edges to OutcomeClass entities.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeClassList(o ...*OutcomeClass) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.RemoveOutcomeClassListIDs(ids...)
}

// ClearOutcomeAnalysisList clears all "outcome_analysis_list" edges to the OutcomeAnalysis entity.
func (omuo *OutcomeMeasureUpdateOne) ClearOutcomeAnalysisList() *OutcomeMeasureUpdateOne {
	omuo.mutation.ClearOutcomeAnalysisList()
	return omuo
}

// RemoveOutcomeAnalysisListIDs removes the "outcome_analysis_list" edge to OutcomeAnalysis entities by IDs.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeAnalysisListIDs(ids ...int) *OutcomeMeasureUpdateOne {
	omuo.mutation.RemoveOutcomeAnalysisListIDs(ids...)
	return omuo
}

// RemoveOutcomeAnalysisList removes "outcome_analysis_list" edges to OutcomeAnalysis entities.
func (omuo *OutcomeMeasureUpdateOne) RemoveOutcomeAnalysisList(o ...*OutcomeAnalysis) *OutcomeMeasureUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return omuo.RemoveOutcomeAnalysisListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (omuo *OutcomeMeasureUpdateOne) Select(field string, fields ...string) *OutcomeMeasureUpdateOne {
	omuo.fields = append([]string{field}, fields...)
	return omuo
}

// Save executes the query and returns the updated OutcomeMeasure entity.
func (omuo *OutcomeMeasureUpdateOne) Save(ctx context.Context) (*OutcomeMeasure, error) {
	var (
		err  error
		node *OutcomeMeasure
	)
	if len(omuo.hooks) == 0 {
		node, err = omuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeMeasureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omuo.mutation = mutation
			node, err = omuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(omuo.hooks) - 1; i >= 0; i-- {
			if omuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = omuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (omuo *OutcomeMeasureUpdateOne) SaveX(ctx context.Context) *OutcomeMeasure {
	node, err := omuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (omuo *OutcomeMeasureUpdateOne) Exec(ctx context.Context) error {
	_, err := omuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omuo *OutcomeMeasureUpdateOne) ExecX(ctx context.Context) {
	if err := omuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (omuo *OutcomeMeasureUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeMeasure, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasure.Table,
			Columns: outcomemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasure.FieldID,
			},
		},
	}
	id, ok := omuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeMeasure.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := omuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasure.FieldID)
		for _, f := range fields {
			if !outcomemeasure.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomemeasure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := omuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omuo.mutation.OutcomeMeasureType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureType,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTitle,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDescription,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasurePopulationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasurePopulationDescription,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureReportingStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureReportingStatus,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureAnticipatedPostingDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureAnticipatedPostingDate,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureParamType,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDispersionType,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureUnitOfMeasure(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureUnitOfMeasure,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureCalculatePct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureCalculatePct,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTimeFrame,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureTypeUnitsAnalyzed,
		})
	}
	if value, ok := omuo.mutation.OutcomeMeasureDenomUnitsSelected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomemeasure.FieldOutcomeMeasureDenomUnitsSelected,
		})
	}
	if omuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omuo.mutation.OutcomeGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.RemovedOutcomeGroupListIDs(); len(nodes) > 0 && !omuo.mutation.OutcomeGroupListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.OutcomeGroupListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omuo.mutation.OutcomeOverviewListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.RemovedOutcomeOverviewListIDs(); len(nodes) > 0 && !omuo.mutation.OutcomeOverviewListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.OutcomeOverviewListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omuo.mutation.OutcomeDenomListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.RemovedOutcomeDenomListIDs(); len(nodes) > 0 && !omuo.mutation.OutcomeDenomListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.OutcomeDenomListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omuo.mutation.OutcomeClassListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.RemovedOutcomeClassListIDs(); len(nodes) > 0 && !omuo.mutation.OutcomeClassListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.OutcomeClassListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if omuo.mutation.OutcomeAnalysisListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.RemovedOutcomeAnalysisListIDs(); len(nodes) > 0 && !omuo.mutation.OutcomeAnalysisListCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := omuo.mutation.OutcomeAnalysisListIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeMeasure{config: omuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, omuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomemeasure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
