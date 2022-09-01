// Code generated (@generated) by entc, DO NOT EDIT.

package baselinedenomcount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineDenomCountGroupID applies equality check predicate on the "baseline_denom_count_group_id" field. It's identical to BaselineDenomCountGroupIDEQ.
func BaselineDenomCountGroupID(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountValue applies equality check predicate on the "baseline_denom_count_value" field. It's identical to BaselineDenomCountValueEQ.
func BaselineDenomCountValue(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountGroupIDEQ applies the EQ predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDEQ(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDNEQ applies the NEQ predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDNEQ(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDIn applies the In predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDIn(vs ...string) predicate.BaselineDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineDenomCountGroupID), v...))
	})
}

// BaselineDenomCountGroupIDNotIn applies the NotIn predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDNotIn(vs ...string) predicate.BaselineDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineDenomCountGroupID), v...))
	})
}

// BaselineDenomCountGroupIDGT applies the GT predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDGT(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDGTE applies the GTE predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDGTE(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDLT applies the LT predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDLT(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDLTE applies the LTE predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDLTE(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDContains applies the Contains predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDContains(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDHasPrefix applies the HasPrefix predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDHasPrefix(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDHasSuffix applies the HasSuffix predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDHasSuffix(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDEqualFold applies the EqualFold predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDEqualFold(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountGroupIDContainsFold applies the ContainsFold predicate on the "baseline_denom_count_group_id" field.
func BaselineDenomCountGroupIDContainsFold(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineDenomCountGroupID), v))
	})
}

// BaselineDenomCountValueEQ applies the EQ predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueEQ(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueNEQ applies the NEQ predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueNEQ(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueIn applies the In predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueIn(vs ...string) predicate.BaselineDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineDenomCountValue), v...))
	})
}

// BaselineDenomCountValueNotIn applies the NotIn predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueNotIn(vs ...string) predicate.BaselineDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineDenomCountValue), v...))
	})
}

// BaselineDenomCountValueGT applies the GT predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueGT(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueGTE applies the GTE predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueGTE(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueLT applies the LT predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueLT(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueLTE applies the LTE predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueLTE(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueContains applies the Contains predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueContains(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueHasPrefix applies the HasPrefix predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueHasPrefix(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueHasSuffix applies the HasSuffix predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueHasSuffix(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueEqualFold applies the EqualFold predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueEqualFold(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineDenomCountValue), v))
	})
}

// BaselineDenomCountValueContainsFold applies the ContainsFold predicate on the "baseline_denom_count_value" field.
func BaselineDenomCountValueContainsFold(v string) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineDenomCountValue), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineDenom) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
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
func And(predicates ...predicate.BaselineDenomCount) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineDenomCount) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
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
func Not(p predicate.BaselineDenomCount) predicate.BaselineDenomCount {
	return predicate.BaselineDenomCount(func(s *sql.Selector) {
		p(s.Not())
	})
}
