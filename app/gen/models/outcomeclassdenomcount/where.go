// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclassdenomcount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeClassDenomCountGroupID applies equality check predicate on the "outcome_class_denom_count_group_id" field. It's identical to OutcomeClassDenomCountGroupIDEQ.
func OutcomeClassDenomCountGroupID(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountValue applies equality check predicate on the "outcome_class_denom_count_value" field. It's identical to OutcomeClassDenomCountValueEQ.
func OutcomeClassDenomCountValue(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountGroupIDEQ applies the EQ predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDEQ(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDNEQ applies the NEQ predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDNEQ(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDIn applies the In predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDIn(vs ...string) predicate.OutcomeClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeClassDenomCountGroupID), v...))
	})
}

// OutcomeClassDenomCountGroupIDNotIn applies the NotIn predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDNotIn(vs ...string) predicate.OutcomeClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeClassDenomCountGroupID), v...))
	})
}

// OutcomeClassDenomCountGroupIDGT applies the GT predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDGT(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDGTE applies the GTE predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDGTE(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDLT applies the LT predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDLT(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDLTE applies the LTE predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDLTE(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDContains applies the Contains predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDContains(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDHasPrefix(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDHasSuffix(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDEqualFold applies the EqualFold predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDEqualFold(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountGroupIDContainsFold applies the ContainsFold predicate on the "outcome_class_denom_count_group_id" field.
func OutcomeClassDenomCountGroupIDContainsFold(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeClassDenomCountGroupID), v))
	})
}

// OutcomeClassDenomCountValueEQ applies the EQ predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueEQ(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueNEQ applies the NEQ predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueNEQ(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueIn applies the In predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueIn(vs ...string) predicate.OutcomeClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeClassDenomCountValue), v...))
	})
}

// OutcomeClassDenomCountValueNotIn applies the NotIn predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueNotIn(vs ...string) predicate.OutcomeClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeClassDenomCountValue), v...))
	})
}

// OutcomeClassDenomCountValueGT applies the GT predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueGT(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueGTE applies the GTE predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueGTE(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueLT applies the LT predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueLT(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueLTE applies the LTE predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueLTE(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueContains applies the Contains predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueContains(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueHasPrefix applies the HasPrefix predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueHasPrefix(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueHasSuffix applies the HasSuffix predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueHasSuffix(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueEqualFold applies the EqualFold predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueEqualFold(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// OutcomeClassDenomCountValueContainsFold applies the ContainsFold predicate on the "outcome_class_denom_count_value" field.
func OutcomeClassDenomCountValueContainsFold(v string) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeClassDenomCountValue), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeClassDenom) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
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
func And(predicates ...predicate.OutcomeClassDenomCount) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeClassDenomCount) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeClassDenomCount) predicate.OutcomeClassDenomCount {
	return predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
		p(s.Not())
	})
}
