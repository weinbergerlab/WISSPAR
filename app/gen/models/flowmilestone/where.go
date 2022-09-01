// Code generated (@generated) by entc, DO NOT EDIT.

package flowmilestone

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowMilestoneType applies equality check predicate on the "flow_milestone_type" field. It's identical to FlowMilestoneTypeEQ.
func FlowMilestoneType(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneComment applies equality check predicate on the "flow_milestone_comment" field. It's identical to FlowMilestoneCommentEQ.
func FlowMilestoneComment(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneTypeEQ applies the EQ predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeEQ(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeNEQ applies the NEQ predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeNEQ(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeIn applies the In predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeIn(vs ...string) predicate.FlowMilestone {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowMilestone(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowMilestoneType), v...))
	})
}

// FlowMilestoneTypeNotIn applies the NotIn predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeNotIn(vs ...string) predicate.FlowMilestone {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowMilestone(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowMilestoneType), v...))
	})
}

// FlowMilestoneTypeGT applies the GT predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeGT(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeGTE applies the GTE predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeGTE(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeLT applies the LT predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeLT(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeLTE applies the LTE predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeLTE(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeContains applies the Contains predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeContains(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeHasPrefix applies the HasPrefix predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeHasPrefix(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeHasSuffix applies the HasSuffix predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeHasSuffix(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeEqualFold applies the EqualFold predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeEqualFold(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneTypeContainsFold applies the ContainsFold predicate on the "flow_milestone_type" field.
func FlowMilestoneTypeContainsFold(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowMilestoneType), v))
	})
}

// FlowMilestoneCommentEQ applies the EQ predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentEQ(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentNEQ applies the NEQ predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentNEQ(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentIn applies the In predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentIn(vs ...string) predicate.FlowMilestone {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowMilestone(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowMilestoneComment), v...))
	})
}

// FlowMilestoneCommentNotIn applies the NotIn predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentNotIn(vs ...string) predicate.FlowMilestone {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowMilestone(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowMilestoneComment), v...))
	})
}

// FlowMilestoneCommentGT applies the GT predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentGT(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentGTE applies the GTE predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentGTE(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentLT applies the LT predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentLT(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentLTE applies the LTE predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentLTE(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentContains applies the Contains predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentContains(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentHasPrefix applies the HasPrefix predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentHasPrefix(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentHasSuffix applies the HasSuffix predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentHasSuffix(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentEqualFold applies the EqualFold predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentEqualFold(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowMilestoneComment), v))
	})
}

// FlowMilestoneCommentContainsFold applies the ContainsFold predicate on the "flow_milestone_comment" field.
func FlowMilestoneCommentContainsFold(v string) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowMilestoneComment), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.FlowPeriod) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
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

// HasFlowAchievementList applies the HasEdge predicate on the "flow_achievement_list" edge.
func HasFlowAchievementList() predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowAchievementListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowAchievementListTable, FlowAchievementListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowAchievementListWith applies the HasEdge predicate on the "flow_achievement_list" edge with a given conditions (other predicates).
func HasFlowAchievementListWith(preds ...predicate.FlowAchievement) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowAchievementListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowAchievementListTable, FlowAchievementListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FlowMilestone) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowMilestone) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
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
func Not(p predicate.FlowMilestone) predicate.FlowMilestone {
	return predicate.FlowMilestone(func(s *sql.Selector) {
		p(s.Not())
	})
}
