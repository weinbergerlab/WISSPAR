// Code generated (@generated) by entc, DO NOT EDIT.

package outcomedenomcount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeDenomCountGroupID applies equality check predicate on the "outcome_denom_count_group_id" field. It's identical to OutcomeDenomCountGroupIDEQ.
func OutcomeDenomCountGroupID(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountValue applies equality check predicate on the "outcome_denom_count_value" field. It's identical to OutcomeDenomCountValueEQ.
func OutcomeDenomCountValue(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountGroupIDEQ applies the EQ predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDEQ(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDNEQ applies the NEQ predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDNEQ(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDIn applies the In predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDIn(vs ...string) predicate.OutcomeDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeDenomCountGroupID), v...))
	})
}

// OutcomeDenomCountGroupIDNotIn applies the NotIn predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDNotIn(vs ...string) predicate.OutcomeDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeDenomCountGroupID), v...))
	})
}

// OutcomeDenomCountGroupIDGT applies the GT predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDGT(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDGTE applies the GTE predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDGTE(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDLT applies the LT predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDLT(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDLTE applies the LTE predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDLTE(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDContains applies the Contains predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDContains(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDHasPrefix(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDHasSuffix(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDEqualFold applies the EqualFold predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDEqualFold(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountGroupIDContainsFold applies the ContainsFold predicate on the "outcome_denom_count_group_id" field.
func OutcomeDenomCountGroupIDContainsFold(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeDenomCountGroupID), v))
	})
}

// OutcomeDenomCountValueEQ applies the EQ predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueEQ(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueNEQ applies the NEQ predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueNEQ(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueIn applies the In predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueIn(vs ...string) predicate.OutcomeDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeDenomCountValue), v...))
	})
}

// OutcomeDenomCountValueNotIn applies the NotIn predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueNotIn(vs ...string) predicate.OutcomeDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeDenomCountValue), v...))
	})
}

// OutcomeDenomCountValueGT applies the GT predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueGT(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueGTE applies the GTE predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueGTE(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueLT applies the LT predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueLT(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueLTE applies the LTE predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueLTE(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueContains applies the Contains predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueContains(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueHasPrefix applies the HasPrefix predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueHasPrefix(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueHasSuffix applies the HasSuffix predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueHasSuffix(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueEqualFold applies the EqualFold predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueEqualFold(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// OutcomeDenomCountValueContainsFold applies the ContainsFold predicate on the "outcome_denom_count_value" field.
func OutcomeDenomCountValueContainsFold(v string) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeDenomCountValue), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeDenom) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
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
func And(predicates ...predicate.OutcomeDenomCount) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeDenomCount) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeDenomCount) predicate.OutcomeDenomCount {
	return predicate.OutcomeDenomCount(func(s *sql.Selector) {
		p(s.Not())
	})
}
