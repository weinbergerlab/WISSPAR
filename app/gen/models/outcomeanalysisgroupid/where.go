// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeanalysisgroupid

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeAnalysisGroupID applies equality check predicate on the "outcome_analysis_group_id" field. It's identical to OutcomeAnalysisGroupIDEQ.
func OutcomeAnalysisGroupID(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDEQ applies the EQ predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDEQ(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDNEQ applies the NEQ predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDNEQ(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDIn applies the In predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDIn(vs ...string) predicate.OutcomeAnalysisGroupID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeAnalysisGroupID), v...))
	})
}

// OutcomeAnalysisGroupIDNotIn applies the NotIn predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDNotIn(vs ...string) predicate.OutcomeAnalysisGroupID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeAnalysisGroupID), v...))
	})
}

// OutcomeAnalysisGroupIDGT applies the GT predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDGT(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDGTE applies the GTE predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDGTE(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDLT applies the LT predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDLT(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDLTE applies the LTE predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDLTE(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDContains applies the Contains predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDContains(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDHasPrefix(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDHasSuffix(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDEqualFold applies the EqualFold predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDEqualFold(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// OutcomeAnalysisGroupIDContainsFold applies the ContainsFold predicate on the "outcome_analysis_group_id" field.
func OutcomeAnalysisGroupIDContainsFold(v string) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeAnalysisGroupID), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeAnalysis) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
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
func And(predicates ...predicate.OutcomeAnalysisGroupID) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeAnalysisGroupID) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeAnalysisGroupID) predicate.OutcomeAnalysisGroupID {
	return predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
		p(s.Not())
	})
}
