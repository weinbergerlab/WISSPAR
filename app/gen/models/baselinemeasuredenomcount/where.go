// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasuredenomcount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineMeasureDenomCountGroupID applies equality check predicate on the "baseline_measure_denom_count_group_id" field. It's identical to BaselineMeasureDenomCountGroupIDEQ.
func BaselineMeasureDenomCountGroupID(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountValue applies equality check predicate on the "baseline_measure_denom_count_value" field. It's identical to BaselineMeasureDenomCountValueEQ.
func BaselineMeasureDenomCountValue(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountGroupIDEQ applies the EQ predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDEQ(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDNEQ applies the NEQ predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDNEQ(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDIn applies the In predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDIn(vs ...string) predicate.BaselineMeasureDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDenomCountGroupID), v...))
	})
}

// BaselineMeasureDenomCountGroupIDNotIn applies the NotIn predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDNotIn(vs ...string) predicate.BaselineMeasureDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDenomCountGroupID), v...))
	})
}

// BaselineMeasureDenomCountGroupIDGT applies the GT predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDGT(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDGTE applies the GTE predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDGTE(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDLT applies the LT predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDLT(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDLTE applies the LTE predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDLTE(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDContains applies the Contains predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDContains(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDHasPrefix applies the HasPrefix predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDHasPrefix(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDHasSuffix applies the HasSuffix predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDHasSuffix(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDEqualFold applies the EqualFold predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDEqualFold(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountGroupIDContainsFold applies the ContainsFold predicate on the "baseline_measure_denom_count_group_id" field.
func BaselineMeasureDenomCountGroupIDContainsFold(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDenomCountGroupID), v))
	})
}

// BaselineMeasureDenomCountValueEQ applies the EQ predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueEQ(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueNEQ applies the NEQ predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueNEQ(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueIn applies the In predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueIn(vs ...string) predicate.BaselineMeasureDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDenomCountValue), v...))
	})
}

// BaselineMeasureDenomCountValueNotIn applies the NotIn predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueNotIn(vs ...string) predicate.BaselineMeasureDenomCount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDenomCountValue), v...))
	})
}

// BaselineMeasureDenomCountValueGT applies the GT predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueGT(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueGTE applies the GTE predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueGTE(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueLT applies the LT predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueLT(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueLTE applies the LTE predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueLTE(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueContains applies the Contains predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueContains(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueHasPrefix applies the HasPrefix predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueHasPrefix(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueHasSuffix applies the HasSuffix predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueHasSuffix(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueEqualFold applies the EqualFold predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueEqualFold(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// BaselineMeasureDenomCountValueContainsFold applies the ContainsFold predicate on the "baseline_measure_denom_count_value" field.
func BaselineMeasureDenomCountValueContainsFold(v string) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDenomCountValue), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineMeasureDenom) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
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
func And(predicates ...predicate.BaselineMeasureDenomCount) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineMeasureDenomCount) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
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
func Not(p predicate.BaselineMeasureDenomCount) predicate.BaselineMeasureDenomCount {
	return predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
		p(s.Not())
	})
}
