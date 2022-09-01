// Code generated (@generated) by entc, DO NOT EDIT.

package flowachievement

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowAchievementGroupID applies equality check predicate on the "flow_achievement_group_id" field. It's identical to FlowAchievementGroupIDEQ.
func FlowAchievementGroupID(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementComment applies equality check predicate on the "flow_achievement_comment" field. It's identical to FlowAchievementCommentEQ.
func FlowAchievementComment(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementNumSubjects applies equality check predicate on the "flow_achievement_num_subjects" field. It's identical to FlowAchievementNumSubjectsEQ.
func FlowAchievementNumSubjects(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumUnits applies equality check predicate on the "flow_achievement_num_units" field. It's identical to FlowAchievementNumUnitsEQ.
func FlowAchievementNumUnits(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementGroupIDEQ applies the EQ predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDNEQ applies the NEQ predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDNEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDIn applies the In predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowAchievementGroupID), v...))
	})
}

// FlowAchievementGroupIDNotIn applies the NotIn predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDNotIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowAchievementGroupID), v...))
	})
}

// FlowAchievementGroupIDGT applies the GT predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDGT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDGTE applies the GTE predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDGTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDLT applies the LT predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDLT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDLTE applies the LTE predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDLTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDContains applies the Contains predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDContains(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDHasPrefix applies the HasPrefix predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDHasPrefix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDHasSuffix applies the HasSuffix predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDHasSuffix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDEqualFold applies the EqualFold predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDEqualFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementGroupIDContainsFold applies the ContainsFold predicate on the "flow_achievement_group_id" field.
func FlowAchievementGroupIDContainsFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowAchievementGroupID), v))
	})
}

// FlowAchievementCommentEQ applies the EQ predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentNEQ applies the NEQ predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentNEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentIn applies the In predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowAchievementComment), v...))
	})
}

// FlowAchievementCommentNotIn applies the NotIn predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentNotIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowAchievementComment), v...))
	})
}

// FlowAchievementCommentGT applies the GT predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentGT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentGTE applies the GTE predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentGTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentLT applies the LT predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentLT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentLTE applies the LTE predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentLTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentContains applies the Contains predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentContains(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentHasPrefix applies the HasPrefix predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentHasPrefix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentHasSuffix applies the HasSuffix predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentHasSuffix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentEqualFold applies the EqualFold predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentEqualFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementCommentContainsFold applies the ContainsFold predicate on the "flow_achievement_comment" field.
func FlowAchievementCommentContainsFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowAchievementComment), v))
	})
}

// FlowAchievementNumSubjectsEQ applies the EQ predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsNEQ applies the NEQ predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsNEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsIn applies the In predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowAchievementNumSubjects), v...))
	})
}

// FlowAchievementNumSubjectsNotIn applies the NotIn predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsNotIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowAchievementNumSubjects), v...))
	})
}

// FlowAchievementNumSubjectsGT applies the GT predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsGT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsGTE applies the GTE predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsGTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsLT applies the LT predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsLT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsLTE applies the LTE predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsLTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsContains applies the Contains predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsContains(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsHasPrefix applies the HasPrefix predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsHasPrefix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsHasSuffix applies the HasSuffix predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsHasSuffix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsEqualFold applies the EqualFold predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsEqualFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumSubjectsContainsFold applies the ContainsFold predicate on the "flow_achievement_num_subjects" field.
func FlowAchievementNumSubjectsContainsFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowAchievementNumSubjects), v))
	})
}

// FlowAchievementNumUnitsEQ applies the EQ predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsNEQ applies the NEQ predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsNEQ(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsIn applies the In predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowAchievementNumUnits), v...))
	})
}

// FlowAchievementNumUnitsNotIn applies the NotIn predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsNotIn(vs ...string) predicate.FlowAchievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowAchievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowAchievementNumUnits), v...))
	})
}

// FlowAchievementNumUnitsGT applies the GT predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsGT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsGTE applies the GTE predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsGTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsLT applies the LT predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsLT(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsLTE applies the LTE predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsLTE(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsContains applies the Contains predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsContains(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsHasPrefix applies the HasPrefix predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsHasPrefix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsHasSuffix applies the HasSuffix predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsHasSuffix(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsEqualFold applies the EqualFold predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsEqualFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// FlowAchievementNumUnitsContainsFold applies the ContainsFold predicate on the "flow_achievement_num_units" field.
func FlowAchievementNumUnitsContainsFold(v string) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowAchievementNumUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.FlowMilestone) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FlowAchievement) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowAchievement) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
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
func Not(p predicate.FlowAchievement) predicate.FlowAchievement {
	return predicate.FlowAchievement(func(s *sql.Selector) {
		p(s.Not())
	})
}
