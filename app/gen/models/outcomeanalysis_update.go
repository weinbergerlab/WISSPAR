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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisUpdate is the builder for updating OutcomeAnalysis entities.
type OutcomeAnalysisUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeAnalysisMutation
}

// Where appends a list predicates to the OutcomeAnalysisUpdate builder.
func (oau *OutcomeAnalysisUpdate) Where(ps ...predicate.OutcomeAnalysis) *OutcomeAnalysisUpdate {
	oau.mutation.Where(ps...)
	return oau
}

// SetOutcomeAnalysisGroupDescription sets the "outcome_analysis_group_description" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisGroupDescription(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisGroupDescription(s)
	return oau
}

// SetOutcomeAnalysisTestedNonInferiority sets the "outcome_analysis_tested_non_inferiority" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisTestedNonInferiority(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisTestedNonInferiority(s)
	return oau
}

// SetOutcomeAnalysisNonInferiorityType sets the "outcome_analysis_non_inferiority_type" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisNonInferiorityType(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisNonInferiorityType(s)
	return oau
}

// SetOutcomeAnalysisNonInferiorityComment sets the "outcome_analysis_non_inferiority_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisNonInferiorityComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisNonInferiorityComment(s)
	return oau
}

// SetOutcomeAnalysisPValue sets the "outcome_analysis_p_value" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisPValue(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisPValue(s)
	return oau
}

// SetOutcomeAnalysisPValueComment sets the "outcome_analysis_p_value_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisPValueComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisPValueComment(s)
	return oau
}

// SetOutcomeAnalysisStatisticalMethod sets the "outcome_analysis_statistical_method" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisStatisticalMethod(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisStatisticalMethod(s)
	return oau
}

// SetOutcomeAnalysisStatisticalComment sets the "outcome_analysis_statistical_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisStatisticalComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisStatisticalComment(s)
	return oau
}

// SetOutcomeAnalysisParamType sets the "outcome_analysis_param_type" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisParamType(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisParamType(s)
	return oau
}

// SetOutcomeAnalysisParamValue sets the "outcome_analysis_param_value" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisParamValue(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisParamValue(s)
	return oau
}

// SetOutcomeAnalysisCiPctValue sets the "outcome_analysis_ci_pct_value" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiPctValue(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiPctValue(s)
	return oau
}

// SetOutcomeAnalysisCiNumSides sets the "outcome_analysis_ci_num_sides" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiNumSides(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiNumSides(s)
	return oau
}

// SetOutcomeAnalysisCiLowerLimit sets the "outcome_analysis_ci_lower_limit" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiLowerLimit(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiLowerLimit(s)
	return oau
}

// SetOutcomeAnalysisCiUpperLimit sets the "outcome_analysis_ci_upper_limit" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiUpperLimit(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiUpperLimit(s)
	return oau
}

// SetOutcomeAnalysisCiLowerLimitComment sets the "outcome_analysis_ci_lower_limit_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiLowerLimitComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiLowerLimitComment(s)
	return oau
}

// SetOutcomeAnalysisCiUpperLimitComment sets the "outcome_analysis_ci_upper_limit_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisCiUpperLimitComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisCiUpperLimitComment(s)
	return oau
}

// SetOutcomeAnalysisDispersionType sets the "outcome_analysis_dispersion_type" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisDispersionType(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisDispersionType(s)
	return oau
}

// SetOutcomeAnalysisDispersionValue sets the "outcome_analysis_dispersion_value" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisDispersionValue(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisDispersionValue(s)
	return oau
}

// SetOutcomeAnalysisEstimateComment sets the "outcome_analysis_estimate_comment" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisEstimateComment(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisEstimateComment(s)
	return oau
}

// SetOutcomeAnalysisOtherAnalysisDescription sets the "outcome_analysis_other_analysis_description" field.
func (oau *OutcomeAnalysisUpdate) SetOutcomeAnalysisOtherAnalysisDescription(s string) *OutcomeAnalysisUpdate {
	oau.mutation.SetOutcomeAnalysisOtherAnalysisDescription(s)
	return oau
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oau *OutcomeAnalysisUpdate) SetParentID(id int) *OutcomeAnalysisUpdate {
	oau.mutation.SetParentID(id)
	return oau
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oau *OutcomeAnalysisUpdate) SetNillableParentID(id *int) *OutcomeAnalysisUpdate {
	if id != nil {
		oau = oau.SetParentID(*id)
	}
	return oau
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oau *OutcomeAnalysisUpdate) SetParent(o *OutcomeMeasure) *OutcomeAnalysisUpdate {
	return oau.SetParentID(o.ID)
}

// AddOutcomeAnalysisGroupIDListIDs adds the "outcome_analysis_group_id_list" edge to the OutcomeAnalysisGroupID entity by IDs.
func (oau *OutcomeAnalysisUpdate) AddOutcomeAnalysisGroupIDListIDs(ids ...int) *OutcomeAnalysisUpdate {
	oau.mutation.AddOutcomeAnalysisGroupIDListIDs(ids...)
	return oau
}

// AddOutcomeAnalysisGroupIDList adds the "outcome_analysis_group_id_list" edges to the OutcomeAnalysisGroupID entity.
func (oau *OutcomeAnalysisUpdate) AddOutcomeAnalysisGroupIDList(o ...*OutcomeAnalysisGroupID) *OutcomeAnalysisUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oau.AddOutcomeAnalysisGroupIDListIDs(ids...)
}

// Mutation returns the OutcomeAnalysisMutation object of the builder.
func (oau *OutcomeAnalysisUpdate) Mutation() *OutcomeAnalysisMutation {
	return oau.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oau *OutcomeAnalysisUpdate) ClearParent() *OutcomeAnalysisUpdate {
	oau.mutation.ClearParent()
	return oau
}

// ClearOutcomeAnalysisGroupIDList clears all "outcome_analysis_group_id_list" edges to the OutcomeAnalysisGroupID entity.
func (oau *OutcomeAnalysisUpdate) ClearOutcomeAnalysisGroupIDList() *OutcomeAnalysisUpdate {
	oau.mutation.ClearOutcomeAnalysisGroupIDList()
	return oau
}

// RemoveOutcomeAnalysisGroupIDListIDs removes the "outcome_analysis_group_id_list" edge to OutcomeAnalysisGroupID entities by IDs.
func (oau *OutcomeAnalysisUpdate) RemoveOutcomeAnalysisGroupIDListIDs(ids ...int) *OutcomeAnalysisUpdate {
	oau.mutation.RemoveOutcomeAnalysisGroupIDListIDs(ids...)
	return oau
}

// RemoveOutcomeAnalysisGroupIDList removes "outcome_analysis_group_id_list" edges to OutcomeAnalysisGroupID entities.
func (oau *OutcomeAnalysisUpdate) RemoveOutcomeAnalysisGroupIDList(o ...*OutcomeAnalysisGroupID) *OutcomeAnalysisUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oau.RemoveOutcomeAnalysisGroupIDListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oau *OutcomeAnalysisUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oau.hooks) == 0 {
		affected, err = oau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oau.mutation = mutation
			affected, err = oau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oau.hooks) - 1; i >= 0; i-- {
			if oau.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oau *OutcomeAnalysisUpdate) SaveX(ctx context.Context) int {
	affected, err := oau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oau *OutcomeAnalysisUpdate) Exec(ctx context.Context) error {
	_, err := oau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oau *OutcomeAnalysisUpdate) ExecX(ctx context.Context) {
	if err := oau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oau *OutcomeAnalysisUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysis.Table,
			Columns: outcomeanalysis.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysis.FieldID,
			},
		},
	}
	if ps := oau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oau.mutation.OutcomeAnalysisGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisGroupDescription,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisTestedNonInferiority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisTestedNonInferiority,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisNonInferiorityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityType,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisNonInferiorityComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisPValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValue,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisPValueComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValueComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisStatisticalMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalMethod,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisStatisticalComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamType,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisParamValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamValue,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiPctValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiPctValue,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiNumSides(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiNumSides,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimit,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimit,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiLowerLimitComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimitComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisCiUpperLimitComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimitComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionType,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisDispersionValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionValue,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisEstimateComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisEstimateComment,
		})
	}
	if value, ok := oau.mutation.OutcomeAnalysisOtherAnalysisDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisOtherAnalysisDescription,
		})
	}
	if oau.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysis.ParentTable,
			Columns: []string{outcomeanalysis.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oau.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysis.ParentTable,
			Columns: []string{outcomeanalysis.ParentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oau.mutation.OutcomeAnalysisGroupIDListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oau.mutation.RemovedOutcomeAnalysisGroupIDListIDs(); len(nodes) > 0 && !oau.mutation.OutcomeAnalysisGroupIDListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oau.mutation.OutcomeAnalysisGroupIDListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeanalysis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeAnalysisUpdateOne is the builder for updating a single OutcomeAnalysis entity.
type OutcomeAnalysisUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeAnalysisMutation
}

// SetOutcomeAnalysisGroupDescription sets the "outcome_analysis_group_description" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisGroupDescription(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisGroupDescription(s)
	return oauo
}

// SetOutcomeAnalysisTestedNonInferiority sets the "outcome_analysis_tested_non_inferiority" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisTestedNonInferiority(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisTestedNonInferiority(s)
	return oauo
}

// SetOutcomeAnalysisNonInferiorityType sets the "outcome_analysis_non_inferiority_type" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisNonInferiorityType(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisNonInferiorityType(s)
	return oauo
}

// SetOutcomeAnalysisNonInferiorityComment sets the "outcome_analysis_non_inferiority_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisNonInferiorityComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisNonInferiorityComment(s)
	return oauo
}

// SetOutcomeAnalysisPValue sets the "outcome_analysis_p_value" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisPValue(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisPValue(s)
	return oauo
}

// SetOutcomeAnalysisPValueComment sets the "outcome_analysis_p_value_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisPValueComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisPValueComment(s)
	return oauo
}

// SetOutcomeAnalysisStatisticalMethod sets the "outcome_analysis_statistical_method" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisStatisticalMethod(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisStatisticalMethod(s)
	return oauo
}

// SetOutcomeAnalysisStatisticalComment sets the "outcome_analysis_statistical_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisStatisticalComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisStatisticalComment(s)
	return oauo
}

// SetOutcomeAnalysisParamType sets the "outcome_analysis_param_type" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisParamType(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisParamType(s)
	return oauo
}

// SetOutcomeAnalysisParamValue sets the "outcome_analysis_param_value" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisParamValue(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisParamValue(s)
	return oauo
}

// SetOutcomeAnalysisCiPctValue sets the "outcome_analysis_ci_pct_value" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiPctValue(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiPctValue(s)
	return oauo
}

// SetOutcomeAnalysisCiNumSides sets the "outcome_analysis_ci_num_sides" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiNumSides(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiNumSides(s)
	return oauo
}

// SetOutcomeAnalysisCiLowerLimit sets the "outcome_analysis_ci_lower_limit" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiLowerLimit(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiLowerLimit(s)
	return oauo
}

// SetOutcomeAnalysisCiUpperLimit sets the "outcome_analysis_ci_upper_limit" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiUpperLimit(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiUpperLimit(s)
	return oauo
}

// SetOutcomeAnalysisCiLowerLimitComment sets the "outcome_analysis_ci_lower_limit_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiLowerLimitComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiLowerLimitComment(s)
	return oauo
}

// SetOutcomeAnalysisCiUpperLimitComment sets the "outcome_analysis_ci_upper_limit_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisCiUpperLimitComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisCiUpperLimitComment(s)
	return oauo
}

// SetOutcomeAnalysisDispersionType sets the "outcome_analysis_dispersion_type" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisDispersionType(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisDispersionType(s)
	return oauo
}

// SetOutcomeAnalysisDispersionValue sets the "outcome_analysis_dispersion_value" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisDispersionValue(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisDispersionValue(s)
	return oauo
}

// SetOutcomeAnalysisEstimateComment sets the "outcome_analysis_estimate_comment" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisEstimateComment(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisEstimateComment(s)
	return oauo
}

// SetOutcomeAnalysisOtherAnalysisDescription sets the "outcome_analysis_other_analysis_description" field.
func (oauo *OutcomeAnalysisUpdateOne) SetOutcomeAnalysisOtherAnalysisDescription(s string) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetOutcomeAnalysisOtherAnalysisDescription(s)
	return oauo
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oauo *OutcomeAnalysisUpdateOne) SetParentID(id int) *OutcomeAnalysisUpdateOne {
	oauo.mutation.SetParentID(id)
	return oauo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oauo *OutcomeAnalysisUpdateOne) SetNillableParentID(id *int) *OutcomeAnalysisUpdateOne {
	if id != nil {
		oauo = oauo.SetParentID(*id)
	}
	return oauo
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oauo *OutcomeAnalysisUpdateOne) SetParent(o *OutcomeMeasure) *OutcomeAnalysisUpdateOne {
	return oauo.SetParentID(o.ID)
}

// AddOutcomeAnalysisGroupIDListIDs adds the "outcome_analysis_group_id_list" edge to the OutcomeAnalysisGroupID entity by IDs.
func (oauo *OutcomeAnalysisUpdateOne) AddOutcomeAnalysisGroupIDListIDs(ids ...int) *OutcomeAnalysisUpdateOne {
	oauo.mutation.AddOutcomeAnalysisGroupIDListIDs(ids...)
	return oauo
}

// AddOutcomeAnalysisGroupIDList adds the "outcome_analysis_group_id_list" edges to the OutcomeAnalysisGroupID entity.
func (oauo *OutcomeAnalysisUpdateOne) AddOutcomeAnalysisGroupIDList(o ...*OutcomeAnalysisGroupID) *OutcomeAnalysisUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oauo.AddOutcomeAnalysisGroupIDListIDs(ids...)
}

// Mutation returns the OutcomeAnalysisMutation object of the builder.
func (oauo *OutcomeAnalysisUpdateOne) Mutation() *OutcomeAnalysisMutation {
	return oauo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oauo *OutcomeAnalysisUpdateOne) ClearParent() *OutcomeAnalysisUpdateOne {
	oauo.mutation.ClearParent()
	return oauo
}

// ClearOutcomeAnalysisGroupIDList clears all "outcome_analysis_group_id_list" edges to the OutcomeAnalysisGroupID entity.
func (oauo *OutcomeAnalysisUpdateOne) ClearOutcomeAnalysisGroupIDList() *OutcomeAnalysisUpdateOne {
	oauo.mutation.ClearOutcomeAnalysisGroupIDList()
	return oauo
}

// RemoveOutcomeAnalysisGroupIDListIDs removes the "outcome_analysis_group_id_list" edge to OutcomeAnalysisGroupID entities by IDs.
func (oauo *OutcomeAnalysisUpdateOne) RemoveOutcomeAnalysisGroupIDListIDs(ids ...int) *OutcomeAnalysisUpdateOne {
	oauo.mutation.RemoveOutcomeAnalysisGroupIDListIDs(ids...)
	return oauo
}

// RemoveOutcomeAnalysisGroupIDList removes "outcome_analysis_group_id_list" edges to OutcomeAnalysisGroupID entities.
func (oauo *OutcomeAnalysisUpdateOne) RemoveOutcomeAnalysisGroupIDList(o ...*OutcomeAnalysisGroupID) *OutcomeAnalysisUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oauo.RemoveOutcomeAnalysisGroupIDListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oauo *OutcomeAnalysisUpdateOne) Select(field string, fields ...string) *OutcomeAnalysisUpdateOne {
	oauo.fields = append([]string{field}, fields...)
	return oauo
}

// Save executes the query and returns the updated OutcomeAnalysis entity.
func (oauo *OutcomeAnalysisUpdateOne) Save(ctx context.Context) (*OutcomeAnalysis, error) {
	var (
		err  error
		node *OutcomeAnalysis
	)
	if len(oauo.hooks) == 0 {
		node, err = oauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oauo.mutation = mutation
			node, err = oauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oauo.hooks) - 1; i >= 0; i-- {
			if oauo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oauo *OutcomeAnalysisUpdateOne) SaveX(ctx context.Context) *OutcomeAnalysis {
	node, err := oauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oauo *OutcomeAnalysisUpdateOne) Exec(ctx context.Context) error {
	_, err := oauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oauo *OutcomeAnalysisUpdateOne) ExecX(ctx context.Context) {
	if err := oauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oauo *OutcomeAnalysisUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeAnalysis, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysis.Table,
			Columns: outcomeanalysis.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysis.FieldID,
			},
		},
	}
	id, ok := oauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeAnalysis.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysis.FieldID)
		for _, f := range fields {
			if !outcomeanalysis.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeanalysis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oauo.mutation.OutcomeAnalysisGroupDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisGroupDescription,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisTestedNonInferiority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisTestedNonInferiority,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisNonInferiorityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityType,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisNonInferiorityComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisPValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValue,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisPValueComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValueComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisStatisticalMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalMethod,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisStatisticalComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisParamType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamType,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisParamValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamValue,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiPctValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiPctValue,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiNumSides(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiNumSides,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimit,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimit,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiLowerLimitComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimitComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisCiUpperLimitComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimitComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisDispersionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionType,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisDispersionValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionValue,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisEstimateComment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisEstimateComment,
		})
	}
	if value, ok := oauo.mutation.OutcomeAnalysisOtherAnalysisDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisOtherAnalysisDescription,
		})
	}
	if oauo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysis.ParentTable,
			Columns: []string{outcomeanalysis.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oauo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysis.ParentTable,
			Columns: []string{outcomeanalysis.ParentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oauo.mutation.OutcomeAnalysisGroupIDListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oauo.mutation.RemovedOutcomeAnalysisGroupIDListIDs(); len(nodes) > 0 && !oauo.mutation.OutcomeAnalysisGroupIDListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oauo.mutation.OutcomeAnalysisGroupIDListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   outcomeanalysis.OutcomeAnalysisGroupIDListTable,
			Columns: []string{outcomeanalysis.OutcomeAnalysisGroupIDListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysisgroupid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeAnalysis{config: oauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeanalysis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
