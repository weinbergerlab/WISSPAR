// Code generated (@generated) by entc, DO NOT EDIT.

package flowperiod

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowPeriodTitle applies equality check predicate on the "flow_period_title" field. It's identical to FlowPeriodTitleEQ.
func FlowPeriodTitle(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleEQ applies the EQ predicate on the "flow_period_title" field.
func FlowPeriodTitleEQ(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleNEQ applies the NEQ predicate on the "flow_period_title" field.
func FlowPeriodTitleNEQ(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleIn applies the In predicate on the "flow_period_title" field.
func FlowPeriodTitleIn(vs ...string) predicate.FlowPeriod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowPeriod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowPeriodTitle), v...))
	})
}

// FlowPeriodTitleNotIn applies the NotIn predicate on the "flow_period_title" field.
func FlowPeriodTitleNotIn(vs ...string) predicate.FlowPeriod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowPeriod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowPeriodTitle), v...))
	})
}

// FlowPeriodTitleGT applies the GT predicate on the "flow_period_title" field.
func FlowPeriodTitleGT(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleGTE applies the GTE predicate on the "flow_period_title" field.
func FlowPeriodTitleGTE(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleLT applies the LT predicate on the "flow_period_title" field.
func FlowPeriodTitleLT(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleLTE applies the LTE predicate on the "flow_period_title" field.
func FlowPeriodTitleLTE(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleContains applies the Contains predicate on the "flow_period_title" field.
func FlowPeriodTitleContains(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleHasPrefix applies the HasPrefix predicate on the "flow_period_title" field.
func FlowPeriodTitleHasPrefix(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleHasSuffix applies the HasSuffix predicate on the "flow_period_title" field.
func FlowPeriodTitleHasSuffix(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleEqualFold applies the EqualFold predicate on the "flow_period_title" field.
func FlowPeriodTitleEqualFold(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowPeriodTitle), v))
	})
}

// FlowPeriodTitleContainsFold applies the ContainsFold predicate on the "flow_period_title" field.
func FlowPeriodTitleContainsFold(v string) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowPeriodTitle), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ParticipantFlowModule) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
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

// HasFlowMilestoneList applies the HasEdge predicate on the "flow_milestone_list" edge.
func HasFlowMilestoneList() predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowMilestoneListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowMilestoneListTable, FlowMilestoneListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowMilestoneListWith applies the HasEdge predicate on the "flow_milestone_list" edge with a given conditions (other predicates).
func HasFlowMilestoneListWith(preds ...predicate.FlowMilestone) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowMilestoneListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowMilestoneListTable, FlowMilestoneListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFlowDropWithdrawList applies the HasEdge predicate on the "flow_drop_withdraw_list" edge.
func HasFlowDropWithdrawList() predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowDropWithdrawListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowDropWithdrawListTable, FlowDropWithdrawListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowDropWithdrawListWith applies the HasEdge predicate on the "flow_drop_withdraw_list" edge with a given conditions (other predicates).
func HasFlowDropWithdrawListWith(preds ...predicate.FlowDropWithdraw) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowDropWithdrawListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowDropWithdrawListTable, FlowDropWithdrawListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FlowPeriod) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowPeriod) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
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
func Not(p predicate.FlowPeriod) predicate.FlowPeriod {
	return predicate.FlowPeriod(func(s *sql.Selector) {
		p(s.Not())
	})
}
