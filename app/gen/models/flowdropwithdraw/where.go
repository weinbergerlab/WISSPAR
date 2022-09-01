// Code generated (@generated) by entc, DO NOT EDIT.

package flowdropwithdraw

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowDropWithdrawType applies equality check predicate on the "flow_drop_withdraw_type" field. It's identical to FlowDropWithdrawTypeEQ.
func FlowDropWithdrawType(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawComment applies equality check predicate on the "flow_drop_withdraw_comment" field. It's identical to FlowDropWithdrawCommentEQ.
func FlowDropWithdrawComment(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawTypeEQ applies the EQ predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeEQ(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeNEQ applies the NEQ predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeNEQ(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeIn applies the In predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeIn(vs ...string) predicate.FlowDropWithdraw {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowDropWithdrawType), v...))
	})
}

// FlowDropWithdrawTypeNotIn applies the NotIn predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeNotIn(vs ...string) predicate.FlowDropWithdraw {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowDropWithdrawType), v...))
	})
}

// FlowDropWithdrawTypeGT applies the GT predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeGT(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeGTE applies the GTE predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeGTE(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeLT applies the LT predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeLT(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeLTE applies the LTE predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeLTE(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeContains applies the Contains predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeContains(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeHasPrefix applies the HasPrefix predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeHasPrefix(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeHasSuffix applies the HasSuffix predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeHasSuffix(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeEqualFold applies the EqualFold predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeEqualFold(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawTypeContainsFold applies the ContainsFold predicate on the "flow_drop_withdraw_type" field.
func FlowDropWithdrawTypeContainsFold(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowDropWithdrawType), v))
	})
}

// FlowDropWithdrawCommentEQ applies the EQ predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentEQ(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentNEQ applies the NEQ predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentNEQ(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentIn applies the In predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentIn(vs ...string) predicate.FlowDropWithdraw {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowDropWithdrawComment), v...))
	})
}

// FlowDropWithdrawCommentNotIn applies the NotIn predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentNotIn(vs ...string) predicate.FlowDropWithdraw {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowDropWithdrawComment), v...))
	})
}

// FlowDropWithdrawCommentGT applies the GT predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentGT(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentGTE applies the GTE predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentGTE(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentLT applies the LT predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentLT(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentLTE applies the LTE predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentLTE(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentContains applies the Contains predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentContains(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentHasPrefix applies the HasPrefix predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentHasPrefix(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentHasSuffix applies the HasSuffix predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentHasSuffix(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentEqualFold applies the EqualFold predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentEqualFold(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// FlowDropWithdrawCommentContainsFold applies the ContainsFold predicate on the "flow_drop_withdraw_comment" field.
func FlowDropWithdrawCommentContainsFold(v string) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowDropWithdrawComment), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.FlowPeriod) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
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

// HasFlowReasonList applies the HasEdge predicate on the "flow_reason_list" edge.
func HasFlowReasonList() predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowReasonListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowReasonListTable, FlowReasonListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowReasonListWith applies the HasEdge predicate on the "flow_reason_list" edge with a given conditions (other predicates).
func HasFlowReasonListWith(preds ...predicate.FlowReason) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowReasonListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowReasonListTable, FlowReasonListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FlowDropWithdraw) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowDropWithdraw) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
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
func Not(p predicate.FlowDropWithdraw) predicate.FlowDropWithdraw {
	return predicate.FlowDropWithdraw(func(s *sql.Selector) {
		p(s.Not())
	})
}
