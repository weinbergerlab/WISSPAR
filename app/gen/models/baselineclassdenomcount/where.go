// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclassdenomcount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineClassDenomCountGroupID applies equality check predicate on the "baseline_class_denom_count_group_id" field. It's identical to BaselineClassDenomCountGroupIDEQ.
func BaselineClassDenomCountGroupID(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountValue applies equality check predicate on the "baseline_class_denom_count_value" field. It's identical to BaselineClassDenomCountValueEQ.
func BaselineClassDenomCountValue(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountGroupIDEQ applies the EQ predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDEQ(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDNEQ applies the NEQ predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDNEQ(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDIn applies the In predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDIn(vs ...string) predicate.BaselineClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineClassDenomCountGroupID), v...))
	})
}

// BaselineClassDenomCountGroupIDNotIn applies the NotIn predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDNotIn(vs ...string) predicate.BaselineClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineClassDenomCountGroupID), v...))
	})
}

// BaselineClassDenomCountGroupIDGT applies the GT predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDGT(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDGTE applies the GTE predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDGTE(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDLT applies the LT predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDLT(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDLTE applies the LTE predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDLTE(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDContains applies the Contains predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDContains(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDHasPrefix applies the HasPrefix predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDHasPrefix(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDHasSuffix applies the HasSuffix predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDHasSuffix(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDEqualFold applies the EqualFold predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDEqualFold(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountGroupIDContainsFold applies the ContainsFold predicate on the "baseline_class_denom_count_group_id" field.
func BaselineClassDenomCountGroupIDContainsFold(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineClassDenomCountGroupID), v))
	})
}

// BaselineClassDenomCountValueEQ applies the EQ predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueEQ(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueNEQ applies the NEQ predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueNEQ(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueIn applies the In predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueIn(vs ...string) predicate.BaselineClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineClassDenomCountValue), v...))
	})
}

// BaselineClassDenomCountValueNotIn applies the NotIn predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueNotIn(vs ...string) predicate.BaselineClassDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineClassDenomCountValue), v...))
	})
}

// BaselineClassDenomCountValueGT applies the GT predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueGT(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueGTE applies the GTE predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueGTE(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueLT applies the LT predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueLT(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueLTE applies the LTE predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueLTE(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueContains applies the Contains predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueContains(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueHasPrefix applies the HasPrefix predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueHasPrefix(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueHasSuffix applies the HasSuffix predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueHasSuffix(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueEqualFold applies the EqualFold predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueEqualFold(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// BaselineClassDenomCountValueContainsFold applies the ContainsFold predicate on the "baseline_class_denom_count_value" field.
func BaselineClassDenomCountValueContainsFold(v string) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineClassDenomCountValue), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineClassDenom) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
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
func And(predicates ...predicate.BaselineClassDenomCount) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineClassDenomCount) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
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
func Not(p predicate.BaselineClassDenomCount) predicate.BaselineClassDenomCount {
	return predicate.BaselineClassDenomCount(func(s *sql.Selector) {
		p(s.Not())
	})
}
