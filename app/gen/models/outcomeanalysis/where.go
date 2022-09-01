// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeanalysis

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeAnalysisGroupDescription applies equality check predicate on the "outcome_analysis_group_description" field. It's identical to OutcomeAnalysisGroupDescriptionEQ.
func OutcomeAnalysisGroupDescription(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisTestedNonInferiority applies equality check predicate on the "outcome_analysis_tested_non_inferiority" field. It's identical to OutcomeAnalysisTestedNonInferiorityEQ.
func OutcomeAnalysisTestedNonInferiority(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisNonInferiorityType applies equality check predicate on the "outcome_analysis_non_inferiority_type" field. It's identical to OutcomeAnalysisNonInferiorityTypeEQ.
func OutcomeAnalysisNonInferiorityType(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityComment applies equality check predicate on the "outcome_analysis_non_inferiority_comment" field. It's identical to OutcomeAnalysisNonInferiorityCommentEQ.
func OutcomeAnalysisNonInferiorityComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisPValue applies equality check predicate on the "outcome_analysis_p_value" field. It's identical to OutcomeAnalysisPValueEQ.
func OutcomeAnalysisPValue(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueComment applies equality check predicate on the "outcome_analysis_p_value_comment" field. It's identical to OutcomeAnalysisPValueCommentEQ.
func OutcomeAnalysisPValueComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisStatisticalMethod applies equality check predicate on the "outcome_analysis_statistical_method" field. It's identical to OutcomeAnalysisStatisticalMethodEQ.
func OutcomeAnalysisStatisticalMethod(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalComment applies equality check predicate on the "outcome_analysis_statistical_comment" field. It's identical to OutcomeAnalysisStatisticalCommentEQ.
func OutcomeAnalysisStatisticalComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisParamType applies equality check predicate on the "outcome_analysis_param_type" field. It's identical to OutcomeAnalysisParamTypeEQ.
func OutcomeAnalysisParamType(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamValue applies equality check predicate on the "outcome_analysis_param_value" field. It's identical to OutcomeAnalysisParamValueEQ.
func OutcomeAnalysisParamValue(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisCiPctValue applies equality check predicate on the "outcome_analysis_ci_pct_value" field. It's identical to OutcomeAnalysisCiPctValueEQ.
func OutcomeAnalysisCiPctValue(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiNumSides applies equality check predicate on the "outcome_analysis_ci_num_sides" field. It's identical to OutcomeAnalysisCiNumSidesEQ.
func OutcomeAnalysisCiNumSides(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiLowerLimit applies equality check predicate on the "outcome_analysis_ci_lower_limit" field. It's identical to OutcomeAnalysisCiLowerLimitEQ.
func OutcomeAnalysisCiLowerLimit(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimit applies equality check predicate on the "outcome_analysis_ci_upper_limit" field. It's identical to OutcomeAnalysisCiUpperLimitEQ.
func OutcomeAnalysisCiUpperLimit(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitComment applies equality check predicate on the "outcome_analysis_ci_lower_limit_comment" field. It's identical to OutcomeAnalysisCiLowerLimitCommentEQ.
func OutcomeAnalysisCiLowerLimitComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitComment applies equality check predicate on the "outcome_analysis_ci_upper_limit_comment" field. It's identical to OutcomeAnalysisCiUpperLimitCommentEQ.
func OutcomeAnalysisCiUpperLimitComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisDispersionType applies equality check predicate on the "outcome_analysis_dispersion_type" field. It's identical to OutcomeAnalysisDispersionTypeEQ.
func OutcomeAnalysisDispersionType(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionValue applies equality check predicate on the "outcome_analysis_dispersion_value" field. It's identical to OutcomeAnalysisDispersionValueEQ.
func OutcomeAnalysisDispersionValue(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisEstimateComment applies equality check predicate on the "outcome_analysis_estimate_comment" field. It's identical to OutcomeAnalysisEstimateCommentEQ.
func OutcomeAnalysisEstimateComment(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescription applies equality check predicate on the "outcome_analysis_other_analysis_description" field. It's identical to OutcomeAnalysisOtherAnalysisDescriptionEQ.
func OutcomeAnalysisOtherAnalysisDescription(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionEQ applies the EQ predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionNEQ applies the NEQ predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionIn applies the In predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisGroupDescription), v...))
	})
}

// OutcomeAnalysisGroupDescriptionNotIn applies the NotIn predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisGroupDescription), v...))
	})
}

// OutcomeAnalysisGroupDescriptionGT applies the GT predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionGTE applies the GTE predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionLT applies the LT predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionLTE applies the LTE predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionContains applies the Contains predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionEqualFold applies the EqualFold predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisGroupDescriptionContainsFold applies the ContainsFold predicate on the "outcome_analysis_group_description" field.
func OutcomeAnalysisGroupDescriptionContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisGroupDescription), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityEQ applies the EQ predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityNEQ applies the NEQ predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityIn applies the In predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisTestedNonInferiority), v...))
	})
}

// OutcomeAnalysisTestedNonInferiorityNotIn applies the NotIn predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisTestedNonInferiority), v...))
	})
}

// OutcomeAnalysisTestedNonInferiorityGT applies the GT predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityGTE applies the GTE predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityLT applies the LT predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityLTE applies the LTE predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityContains applies the Contains predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityHasPrefix applies the HasPrefix predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityHasSuffix applies the HasSuffix predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityEqualFold applies the EqualFold predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisTestedNonInferiorityContainsFold applies the ContainsFold predicate on the "outcome_analysis_tested_non_inferiority" field.
func OutcomeAnalysisTestedNonInferiorityContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisTestedNonInferiority), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeEQ applies the EQ predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeNEQ applies the NEQ predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeIn applies the In predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisNonInferiorityType), v...))
	})
}

// OutcomeAnalysisNonInferiorityTypeNotIn applies the NotIn predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisNonInferiorityType), v...))
	})
}

// OutcomeAnalysisNonInferiorityTypeGT applies the GT predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeGTE applies the GTE predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeLT applies the LT predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeLTE applies the LTE predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeContains applies the Contains predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeHasPrefix applies the HasPrefix predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeHasSuffix applies the HasSuffix predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeEqualFold applies the EqualFold predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityTypeContainsFold applies the ContainsFold predicate on the "outcome_analysis_non_inferiority_type" field.
func OutcomeAnalysisNonInferiorityTypeContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisNonInferiorityType), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentEQ applies the EQ predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentNEQ applies the NEQ predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentIn applies the In predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisNonInferiorityComment), v...))
	})
}

// OutcomeAnalysisNonInferiorityCommentNotIn applies the NotIn predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisNonInferiorityComment), v...))
	})
}

// OutcomeAnalysisNonInferiorityCommentGT applies the GT predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentGTE applies the GTE predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentLT applies the LT predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentLTE applies the LTE predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentContains applies the Contains predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisNonInferiorityCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_non_inferiority_comment" field.
func OutcomeAnalysisNonInferiorityCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisNonInferiorityComment), v))
	})
}

// OutcomeAnalysisPValueEQ applies the EQ predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueNEQ applies the NEQ predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueIn applies the In predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisPValue), v...))
	})
}

// OutcomeAnalysisPValueNotIn applies the NotIn predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisPValue), v...))
	})
}

// OutcomeAnalysisPValueGT applies the GT predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueGTE applies the GTE predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueLT applies the LT predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueLTE applies the LTE predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueContains applies the Contains predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueHasPrefix applies the HasPrefix predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueHasSuffix applies the HasSuffix predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueEqualFold applies the EqualFold predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueContainsFold applies the ContainsFold predicate on the "outcome_analysis_p_value" field.
func OutcomeAnalysisPValueContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisPValue), v))
	})
}

// OutcomeAnalysisPValueCommentEQ applies the EQ predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentNEQ applies the NEQ predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentIn applies the In predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisPValueComment), v...))
	})
}

// OutcomeAnalysisPValueCommentNotIn applies the NotIn predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisPValueComment), v...))
	})
}

// OutcomeAnalysisPValueCommentGT applies the GT predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentGTE applies the GTE predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentLT applies the LT predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentLTE applies the LTE predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentContains applies the Contains predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisPValueCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_p_value_comment" field.
func OutcomeAnalysisPValueCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisPValueComment), v))
	})
}

// OutcomeAnalysisStatisticalMethodEQ applies the EQ predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodNEQ applies the NEQ predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodIn applies the In predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisStatisticalMethod), v...))
	})
}

// OutcomeAnalysisStatisticalMethodNotIn applies the NotIn predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisStatisticalMethod), v...))
	})
}

// OutcomeAnalysisStatisticalMethodGT applies the GT predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodGTE applies the GTE predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodLT applies the LT predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodLTE applies the LTE predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodContains applies the Contains predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodHasPrefix applies the HasPrefix predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodHasSuffix applies the HasSuffix predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodEqualFold applies the EqualFold predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalMethodContainsFold applies the ContainsFold predicate on the "outcome_analysis_statistical_method" field.
func OutcomeAnalysisStatisticalMethodContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisStatisticalMethod), v))
	})
}

// OutcomeAnalysisStatisticalCommentEQ applies the EQ predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentNEQ applies the NEQ predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentIn applies the In predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisStatisticalComment), v...))
	})
}

// OutcomeAnalysisStatisticalCommentNotIn applies the NotIn predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisStatisticalComment), v...))
	})
}

// OutcomeAnalysisStatisticalCommentGT applies the GT predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentGTE applies the GTE predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentLT applies the LT predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentLTE applies the LTE predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentContains applies the Contains predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisStatisticalCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_statistical_comment" field.
func OutcomeAnalysisStatisticalCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisStatisticalComment), v))
	})
}

// OutcomeAnalysisParamTypeEQ applies the EQ predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeNEQ applies the NEQ predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeIn applies the In predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisParamType), v...))
	})
}

// OutcomeAnalysisParamTypeNotIn applies the NotIn predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisParamType), v...))
	})
}

// OutcomeAnalysisParamTypeGT applies the GT predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeGTE applies the GTE predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeLT applies the LT predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeLTE applies the LTE predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeContains applies the Contains predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeHasPrefix applies the HasPrefix predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeHasSuffix applies the HasSuffix predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeEqualFold applies the EqualFold predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamTypeContainsFold applies the ContainsFold predicate on the "outcome_analysis_param_type" field.
func OutcomeAnalysisParamTypeContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisParamType), v))
	})
}

// OutcomeAnalysisParamValueEQ applies the EQ predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueNEQ applies the NEQ predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueIn applies the In predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisParamValue), v...))
	})
}

// OutcomeAnalysisParamValueNotIn applies the NotIn predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisParamValue), v...))
	})
}

// OutcomeAnalysisParamValueGT applies the GT predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueGTE applies the GTE predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueLT applies the LT predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueLTE applies the LTE predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueContains applies the Contains predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueHasPrefix applies the HasPrefix predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueHasSuffix applies the HasSuffix predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueEqualFold applies the EqualFold predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisParamValueContainsFold applies the ContainsFold predicate on the "outcome_analysis_param_value" field.
func OutcomeAnalysisParamValueContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisParamValue), v))
	})
}

// OutcomeAnalysisCiPctValueEQ applies the EQ predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueNEQ applies the NEQ predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueIn applies the In predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiPctValue), v...))
	})
}

// OutcomeAnalysisCiPctValueNotIn applies the NotIn predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiPctValue), v...))
	})
}

// OutcomeAnalysisCiPctValueGT applies the GT predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueGTE applies the GTE predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueLT applies the LT predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueLTE applies the LTE predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueContains applies the Contains predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiPctValueContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_pct_value" field.
func OutcomeAnalysisCiPctValueContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiPctValue), v))
	})
}

// OutcomeAnalysisCiNumSidesEQ applies the EQ predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesNEQ applies the NEQ predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesIn applies the In predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiNumSides), v...))
	})
}

// OutcomeAnalysisCiNumSidesNotIn applies the NotIn predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiNumSides), v...))
	})
}

// OutcomeAnalysisCiNumSidesGT applies the GT predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesGTE applies the GTE predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesLT applies the LT predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesLTE applies the LTE predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesContains applies the Contains predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiNumSidesContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_num_sides" field.
func OutcomeAnalysisCiNumSidesContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiNumSides), v))
	})
}

// OutcomeAnalysisCiLowerLimitEQ applies the EQ predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitNEQ applies the NEQ predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitIn applies the In predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiLowerLimit), v...))
	})
}

// OutcomeAnalysisCiLowerLimitNotIn applies the NotIn predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiLowerLimit), v...))
	})
}

// OutcomeAnalysisCiLowerLimitGT applies the GT predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitGTE applies the GTE predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitLT applies the LT predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitLTE applies the LTE predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitContains applies the Contains predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_lower_limit" field.
func OutcomeAnalysisCiLowerLimitContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiLowerLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitEQ applies the EQ predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitNEQ applies the NEQ predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitIn applies the In predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiUpperLimit), v...))
	})
}

// OutcomeAnalysisCiUpperLimitNotIn applies the NotIn predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiUpperLimit), v...))
	})
}

// OutcomeAnalysisCiUpperLimitGT applies the GT predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitGTE applies the GTE predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitLT applies the LT predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitLTE applies the LTE predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitContains applies the Contains predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiUpperLimitContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_upper_limit" field.
func OutcomeAnalysisCiUpperLimitContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiUpperLimit), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentEQ applies the EQ predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentNEQ applies the NEQ predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentIn applies the In predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v...))
	})
}

// OutcomeAnalysisCiLowerLimitCommentNotIn applies the NotIn predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v...))
	})
}

// OutcomeAnalysisCiLowerLimitCommentGT applies the GT predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentGTE applies the GTE predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentLT applies the LT predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentLTE applies the LTE predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentContains applies the Contains predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiLowerLimitCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_lower_limit_comment" field.
func OutcomeAnalysisCiLowerLimitCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiLowerLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentEQ applies the EQ predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentNEQ applies the NEQ predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentIn applies the In predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v...))
	})
}

// OutcomeAnalysisCiUpperLimitCommentNotIn applies the NotIn predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v...))
	})
}

// OutcomeAnalysisCiUpperLimitCommentGT applies the GT predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentGTE applies the GTE predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentLT applies the LT predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentLTE applies the LTE predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentContains applies the Contains predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisCiUpperLimitCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_ci_upper_limit_comment" field.
func OutcomeAnalysisCiUpperLimitCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisCiUpperLimitComment), v))
	})
}

// OutcomeAnalysisDispersionTypeEQ applies the EQ predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeNEQ applies the NEQ predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeIn applies the In predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisDispersionType), v...))
	})
}

// OutcomeAnalysisDispersionTypeNotIn applies the NotIn predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisDispersionType), v...))
	})
}

// OutcomeAnalysisDispersionTypeGT applies the GT predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeGTE applies the GTE predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeLT applies the LT predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeLTE applies the LTE predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeContains applies the Contains predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeHasPrefix applies the HasPrefix predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeHasSuffix applies the HasSuffix predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeEqualFold applies the EqualFold predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionTypeContainsFold applies the ContainsFold predicate on the "outcome_analysis_dispersion_type" field.
func OutcomeAnalysisDispersionTypeContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisDispersionType), v))
	})
}

// OutcomeAnalysisDispersionValueEQ applies the EQ predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueNEQ applies the NEQ predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueIn applies the In predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisDispersionValue), v...))
	})
}

// OutcomeAnalysisDispersionValueNotIn applies the NotIn predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisDispersionValue), v...))
	})
}

// OutcomeAnalysisDispersionValueGT applies the GT predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueGTE applies the GTE predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueLT applies the LT predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueLTE applies the LTE predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueContains applies the Contains predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueHasPrefix applies the HasPrefix predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueHasSuffix applies the HasSuffix predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueEqualFold applies the EqualFold predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisDispersionValueContainsFold applies the ContainsFold predicate on the "outcome_analysis_dispersion_value" field.
func OutcomeAnalysisDispersionValueContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisDispersionValue), v))
	})
}

// OutcomeAnalysisEstimateCommentEQ applies the EQ predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentNEQ applies the NEQ predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentIn applies the In predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisEstimateComment), v...))
	})
}

// OutcomeAnalysisEstimateCommentNotIn applies the NotIn predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisEstimateComment), v...))
	})
}

// OutcomeAnalysisEstimateCommentGT applies the GT predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentGTE applies the GTE predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentLT applies the LT predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentLTE applies the LTE predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentContains applies the Contains predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentHasPrefix applies the HasPrefix predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentHasSuffix applies the HasSuffix predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentEqualFold applies the EqualFold predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisEstimateCommentContainsFold applies the ContainsFold predicate on the "outcome_analysis_estimate_comment" field.
func OutcomeAnalysisEstimateCommentContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisEstimateComment), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionEQ applies the EQ predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionNEQ applies the NEQ predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionNEQ(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionIn applies the In predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v...))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionNotIn applies the NotIn predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionNotIn(vs ...string) predicate.OutcomeAnalysis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v...))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionGT applies the GT predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionGT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionGTE applies the GTE predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionGTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionLT applies the LT predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionLT(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionLTE applies the LTE predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionLTE(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionContains applies the Contains predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionContains(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionHasPrefix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionHasSuffix(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionEqualFold applies the EqualFold predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionEqualFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// OutcomeAnalysisOtherAnalysisDescriptionContainsFold applies the ContainsFold predicate on the "outcome_analysis_other_analysis_description" field.
func OutcomeAnalysisOtherAnalysisDescriptionContainsFold(v string) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisOtherAnalysisDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeAnalysisGroupIDList applies the HasEdge predicate on the "outcome_analysis_group_id_list" edge.
func HasOutcomeAnalysisGroupIDList() predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeAnalysisGroupIDListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeAnalysisGroupIDListTable, OutcomeAnalysisGroupIDListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeAnalysisGroupIDListWith applies the HasEdge predicate on the "outcome_analysis_group_id_list" edge with a given conditions (other predicates).
func HasOutcomeAnalysisGroupIDListWith(preds ...predicate.OutcomeAnalysisGroupID) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeAnalysisGroupIDListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeAnalysisGroupIDListTable, OutcomeAnalysisGroupIDListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeAnalysis) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeAnalysis) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OutcomeAnalysis) predicate.OutcomeAnalysis {
	return predicate.OutcomeAnalysis(func(s *sql.Selector) {
		p(s.Not())
	})
}
