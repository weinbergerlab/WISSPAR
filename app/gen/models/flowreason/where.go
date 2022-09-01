// Code generated (@generated) by entc, DO NOT EDIT.

package flowreason

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowReasonGroupID applies equality check predicate on the "flow_reason_group_id" field. It's identical to FlowReasonGroupIDEQ.
func FlowReasonGroupID(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonComment applies equality check predicate on the "flow_reason_comment" field. It's identical to FlowReasonCommentEQ.
func FlowReasonComment(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonNumSubjects applies equality check predicate on the "flow_reason_num_subjects" field. It's identical to FlowReasonNumSubjectsEQ.
func FlowReasonNumSubjects(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumUnits applies equality check predicate on the "flow_reason_num_units" field. It's identical to FlowReasonNumUnitsEQ.
func FlowReasonNumUnits(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonGroupIDEQ applies the EQ predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDNEQ applies the NEQ predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDNEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDIn applies the In predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowReasonGroupID), v...))
	})
}

// FlowReasonGroupIDNotIn applies the NotIn predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDNotIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowReasonGroupID), v...))
	})
}

// FlowReasonGroupIDGT applies the GT predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDGT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDGTE applies the GTE predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDGTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDLT applies the LT predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDLT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDLTE applies the LTE predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDLTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDContains applies the Contains predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDContains(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDHasPrefix applies the HasPrefix predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDHasPrefix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDHasSuffix applies the HasSuffix predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDHasSuffix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDEqualFold applies the EqualFold predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDEqualFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonGroupIDContainsFold applies the ContainsFold predicate on the "flow_reason_group_id" field.
func FlowReasonGroupIDContainsFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowReasonGroupID), v))
	})
}

// FlowReasonCommentEQ applies the EQ predicate on the "flow_reason_comment" field.
func FlowReasonCommentEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentNEQ applies the NEQ predicate on the "flow_reason_comment" field.
func FlowReasonCommentNEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentIn applies the In predicate on the "flow_reason_comment" field.
func FlowReasonCommentIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowReasonComment), v...))
	})
}

// FlowReasonCommentNotIn applies the NotIn predicate on the "flow_reason_comment" field.
func FlowReasonCommentNotIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowReasonComment), v...))
	})
}

// FlowReasonCommentGT applies the GT predicate on the "flow_reason_comment" field.
func FlowReasonCommentGT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentGTE applies the GTE predicate on the "flow_reason_comment" field.
func FlowReasonCommentGTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentLT applies the LT predicate on the "flow_reason_comment" field.
func FlowReasonCommentLT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentLTE applies the LTE predicate on the "flow_reason_comment" field.
func FlowReasonCommentLTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentContains applies the Contains predicate on the "flow_reason_comment" field.
func FlowReasonCommentContains(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentHasPrefix applies the HasPrefix predicate on the "flow_reason_comment" field.
func FlowReasonCommentHasPrefix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentHasSuffix applies the HasSuffix predicate on the "flow_reason_comment" field.
func FlowReasonCommentHasSuffix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentEqualFold applies the EqualFold predicate on the "flow_reason_comment" field.
func FlowReasonCommentEqualFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonCommentContainsFold applies the ContainsFold predicate on the "flow_reason_comment" field.
func FlowReasonCommentContainsFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowReasonComment), v))
	})
}

// FlowReasonNumSubjectsEQ applies the EQ predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsNEQ applies the NEQ predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsNEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsIn applies the In predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowReasonNumSubjects), v...))
	})
}

// FlowReasonNumSubjectsNotIn applies the NotIn predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsNotIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowReasonNumSubjects), v...))
	})
}

// FlowReasonNumSubjectsGT applies the GT predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsGT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsGTE applies the GTE predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsGTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsLT applies the LT predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsLT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsLTE applies the LTE predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsLTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsContains applies the Contains predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsContains(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsHasPrefix applies the HasPrefix predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsHasPrefix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsHasSuffix applies the HasSuffix predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsHasSuffix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsEqualFold applies the EqualFold predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsEqualFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumSubjectsContainsFold applies the ContainsFold predicate on the "flow_reason_num_subjects" field.
func FlowReasonNumSubjectsContainsFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowReasonNumSubjects), v))
	})
}

// FlowReasonNumUnitsEQ applies the EQ predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsNEQ applies the NEQ predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsNEQ(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsIn applies the In predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowReasonNumUnits), v...))
	})
}

// FlowReasonNumUnitsNotIn applies the NotIn predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsNotIn(vs ...string) predicate.FlowReason {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowReason(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowReasonNumUnits), v...))
	})
}

// FlowReasonNumUnitsGT applies the GT predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsGT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsGTE applies the GTE predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsGTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsLT applies the LT predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsLT(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsLTE applies the LTE predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsLTE(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsContains applies the Contains predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsContains(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsHasPrefix applies the HasPrefix predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsHasPrefix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsHasSuffix applies the HasSuffix predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsHasSuffix(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsEqualFold applies the EqualFold predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsEqualFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowReasonNumUnits), v))
	})
}

// FlowReasonNumUnitsContainsFold applies the ContainsFold predicate on the "flow_reason_num_units" field.
func FlowReasonNumUnitsContainsFold(v string) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowReasonNumUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.FlowDropWithdraw) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
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
func And(predicates ...predicate.FlowReason) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowReason) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
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
func Not(p predicate.FlowReason) predicate.FlowReason {
	return predicate.FlowReason(func(s *sql.Selector) {
		p(s.Not())
	})
}
