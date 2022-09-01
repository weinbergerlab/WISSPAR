// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
)

// OutcomeAnalysisCreate is the builder for creating a OutcomeAnalysis entity.
type OutcomeAnalysisCreate struct {
	config
	mutation *OutcomeAnalysisMutation
	hooks    []Hook
}

// SetOutcomeAnalysisGroupDescription sets the "outcome_analysis_group_description" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisGroupDescription(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisGroupDescription(s)
	return oac
}

// SetOutcomeAnalysisTestedNonInferiority sets the "outcome_analysis_tested_non_inferiority" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisTestedNonInferiority(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisTestedNonInferiority(s)
	return oac
}

// SetOutcomeAnalysisNonInferiorityType sets the "outcome_analysis_non_inferiority_type" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisNonInferiorityType(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisNonInferiorityType(s)
	return oac
}

// SetOutcomeAnalysisNonInferiorityComment sets the "outcome_analysis_non_inferiority_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisNonInferiorityComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisNonInferiorityComment(s)
	return oac
}

// SetOutcomeAnalysisPValue sets the "outcome_analysis_p_value" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisPValue(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisPValue(s)
	return oac
}

// SetOutcomeAnalysisPValueComment sets the "outcome_analysis_p_value_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisPValueComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisPValueComment(s)
	return oac
}

// SetOutcomeAnalysisStatisticalMethod sets the "outcome_analysis_statistical_method" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisStatisticalMethod(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisStatisticalMethod(s)
	return oac
}

// SetOutcomeAnalysisStatisticalComment sets the "outcome_analysis_statistical_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisStatisticalComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisStatisticalComment(s)
	return oac
}

// SetOutcomeAnalysisParamType sets the "outcome_analysis_param_type" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisParamType(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisParamType(s)
	return oac
}

// SetOutcomeAnalysisParamValue sets the "outcome_analysis_param_value" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisParamValue(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisParamValue(s)
	return oac
}

// SetOutcomeAnalysisCiPctValue sets the "outcome_analysis_ci_pct_value" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiPctValue(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiPctValue(s)
	return oac
}

// SetOutcomeAnalysisCiNumSides sets the "outcome_analysis_ci_num_sides" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiNumSides(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiNumSides(s)
	return oac
}

// SetOutcomeAnalysisCiLowerLimit sets the "outcome_analysis_ci_lower_limit" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiLowerLimit(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiLowerLimit(s)
	return oac
}

// SetOutcomeAnalysisCiUpperLimit sets the "outcome_analysis_ci_upper_limit" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiUpperLimit(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiUpperLimit(s)
	return oac
}

// SetOutcomeAnalysisCiLowerLimitComment sets the "outcome_analysis_ci_lower_limit_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiLowerLimitComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiLowerLimitComment(s)
	return oac
}

// SetOutcomeAnalysisCiUpperLimitComment sets the "outcome_analysis_ci_upper_limit_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisCiUpperLimitComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisCiUpperLimitComment(s)
	return oac
}

// SetOutcomeAnalysisDispersionType sets the "outcome_analysis_dispersion_type" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisDispersionType(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisDispersionType(s)
	return oac
}

// SetOutcomeAnalysisDispersionValue sets the "outcome_analysis_dispersion_value" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisDispersionValue(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisDispersionValue(s)
	return oac
}

// SetOutcomeAnalysisEstimateComment sets the "outcome_analysis_estimate_comment" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisEstimateComment(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisEstimateComment(s)
	return oac
}

// SetOutcomeAnalysisOtherAnalysisDescription sets the "outcome_analysis_other_analysis_description" field.
func (oac *OutcomeAnalysisCreate) SetOutcomeAnalysisOtherAnalysisDescription(s string) *OutcomeAnalysisCreate {
	oac.mutation.SetOutcomeAnalysisOtherAnalysisDescription(s)
	return oac
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oac *OutcomeAnalysisCreate) SetParentID(id int) *OutcomeAnalysisCreate {
	oac.mutation.SetParentID(id)
	return oac
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oac *OutcomeAnalysisCreate) SetNillableParentID(id *int) *OutcomeAnalysisCreate {
	if id != nil {
		oac = oac.SetParentID(*id)
	}
	return oac
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oac *OutcomeAnalysisCreate) SetParent(o *OutcomeMeasure) *OutcomeAnalysisCreate {
	return oac.SetParentID(o.ID)
}

// AddOutcomeAnalysisGroupIDListIDs adds the "outcome_analysis_group_id_list" edge to the OutcomeAnalysisGroupID entity by IDs.
func (oac *OutcomeAnalysisCreate) AddOutcomeAnalysisGroupIDListIDs(ids ...int) *OutcomeAnalysisCreate {
	oac.mutation.AddOutcomeAnalysisGroupIDListIDs(ids...)
	return oac
}

// AddOutcomeAnalysisGroupIDList adds the "outcome_analysis_group_id_list" edges to the OutcomeAnalysisGroupID entity.
func (oac *OutcomeAnalysisCreate) AddOutcomeAnalysisGroupIDList(o ...*OutcomeAnalysisGroupID) *OutcomeAnalysisCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oac.AddOutcomeAnalysisGroupIDListIDs(ids...)
}

// Mutation returns the OutcomeAnalysisMutation object of the builder.
func (oac *OutcomeAnalysisCreate) Mutation() *OutcomeAnalysisMutation {
	return oac.mutation
}

// Save creates the OutcomeAnalysis in the database.
func (oac *OutcomeAnalysisCreate) Save(ctx context.Context) (*OutcomeAnalysis, error) {
	var (
		err  error
		node *OutcomeAnalysis
	)
	if len(oac.hooks) == 0 {
		if err = oac.check(); err != nil {
			return nil, err
		}
		node, err = oac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oac.check(); err != nil {
				return nil, err
			}
			oac.mutation = mutation
			if node, err = oac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oac.hooks) - 1; i >= 0; i-- {
			if oac.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OutcomeAnalysisCreate) SaveX(ctx context.Context) *OutcomeAnalysis {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oac *OutcomeAnalysisCreate) Exec(ctx context.Context) error {
	_, err := oac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oac *OutcomeAnalysisCreate) ExecX(ctx context.Context) {
	if err := oac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oac *OutcomeAnalysisCreate) check() error {
	if _, ok := oac.mutation.OutcomeAnalysisGroupDescription(); !ok {
		return &ValidationError{Name: "outcome_analysis_group_description", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_group_description"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisTestedNonInferiority(); !ok {
		return &ValidationError{Name: "outcome_analysis_tested_non_inferiority", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_tested_non_inferiority"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisNonInferiorityType(); !ok {
		return &ValidationError{Name: "outcome_analysis_non_inferiority_type", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_non_inferiority_type"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisNonInferiorityComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_non_inferiority_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_non_inferiority_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisPValue(); !ok {
		return &ValidationError{Name: "outcome_analysis_p_value", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_p_value"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisPValueComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_p_value_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_p_value_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisStatisticalMethod(); !ok {
		return &ValidationError{Name: "outcome_analysis_statistical_method", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_statistical_method"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisStatisticalComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_statistical_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_statistical_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisParamType(); !ok {
		return &ValidationError{Name: "outcome_analysis_param_type", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_param_type"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisParamValue(); !ok {
		return &ValidationError{Name: "outcome_analysis_param_value", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_param_value"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiPctValue(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_pct_value", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_pct_value"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiNumSides(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_num_sides", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_num_sides"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiLowerLimit(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_lower_limit", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_lower_limit"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiUpperLimit(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_upper_limit", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_upper_limit"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiLowerLimitComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_lower_limit_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_lower_limit_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisCiUpperLimitComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_ci_upper_limit_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_ci_upper_limit_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisDispersionType(); !ok {
		return &ValidationError{Name: "outcome_analysis_dispersion_type", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_dispersion_type"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisDispersionValue(); !ok {
		return &ValidationError{Name: "outcome_analysis_dispersion_value", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_dispersion_value"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisEstimateComment(); !ok {
		return &ValidationError{Name: "outcome_analysis_estimate_comment", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_estimate_comment"`)}
	}
	if _, ok := oac.mutation.OutcomeAnalysisOtherAnalysisDescription(); !ok {
		return &ValidationError{Name: "outcome_analysis_other_analysis_description", err: errors.New(`models: missing required field "OutcomeAnalysis.outcome_analysis_other_analysis_description"`)}
	}
	return nil
}

func (oac *OutcomeAnalysisCreate) sqlSave(ctx context.Context) (*OutcomeAnalysis, error) {
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oac *OutcomeAnalysisCreate) createSpec() (*OutcomeAnalysis, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeAnalysis{config: oac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeanalysis.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysis.FieldID,
			},
		}
	)
	if value, ok := oac.mutation.OutcomeAnalysisGroupDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisGroupDescription,
		})
		_node.OutcomeAnalysisGroupDescription = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisTestedNonInferiority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisTestedNonInferiority,
		})
		_node.OutcomeAnalysisTestedNonInferiority = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisNonInferiorityType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityType,
		})
		_node.OutcomeAnalysisNonInferiorityType = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisNonInferiorityComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisNonInferiorityComment,
		})
		_node.OutcomeAnalysisNonInferiorityComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisPValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValue,
		})
		_node.OutcomeAnalysisPValue = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisPValueComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisPValueComment,
		})
		_node.OutcomeAnalysisPValueComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisStatisticalMethod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalMethod,
		})
		_node.OutcomeAnalysisStatisticalMethod = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisStatisticalComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisStatisticalComment,
		})
		_node.OutcomeAnalysisStatisticalComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisParamType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamType,
		})
		_node.OutcomeAnalysisParamType = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisParamValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisParamValue,
		})
		_node.OutcomeAnalysisParamValue = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiPctValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiPctValue,
		})
		_node.OutcomeAnalysisCiPctValue = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiNumSides(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiNumSides,
		})
		_node.OutcomeAnalysisCiNumSides = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiLowerLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimit,
		})
		_node.OutcomeAnalysisCiLowerLimit = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiUpperLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimit,
		})
		_node.OutcomeAnalysisCiUpperLimit = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiLowerLimitComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiLowerLimitComment,
		})
		_node.OutcomeAnalysisCiLowerLimitComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisCiUpperLimitComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisCiUpperLimitComment,
		})
		_node.OutcomeAnalysisCiUpperLimitComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisDispersionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionType,
		})
		_node.OutcomeAnalysisDispersionType = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisDispersionValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisDispersionValue,
		})
		_node.OutcomeAnalysisDispersionValue = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisEstimateComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisEstimateComment,
		})
		_node.OutcomeAnalysisEstimateComment = value
	}
	if value, ok := oac.mutation.OutcomeAnalysisOtherAnalysisDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysis.FieldOutcomeAnalysisOtherAnalysisDescription,
		})
		_node.OutcomeAnalysisOtherAnalysisDescription = value
	}
	if nodes := oac.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.outcome_measure_outcome_analysis_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.OutcomeAnalysisGroupIDListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeAnalysisCreateBulk is the builder for creating many OutcomeAnalysis entities in bulk.
type OutcomeAnalysisCreateBulk struct {
	config
	builders []*OutcomeAnalysisCreate
}

// Save creates the OutcomeAnalysis entities in the database.
func (oacb *OutcomeAnalysisCreateBulk) Save(ctx context.Context) ([]*OutcomeAnalysis, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OutcomeAnalysis, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeAnalysisMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OutcomeAnalysisCreateBulk) SaveX(ctx context.Context) []*OutcomeAnalysis {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oacb *OutcomeAnalysisCreateBulk) Exec(ctx context.Context) error {
	_, err := oacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacb *OutcomeAnalysisCreateBulk) ExecX(ctx context.Context) {
	if err := oacb.Exec(ctx); err != nil {
		panic(err)
	}
}
