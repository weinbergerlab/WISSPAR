// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasuredenom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineMeasureDenomUnits applies equality check predicate on the "baseline_measure_denom_units" field. It's identical to BaselineMeasureDenomUnitsEQ.
func BaselineMeasureDenomUnits(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsEQ applies the EQ predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsEQ(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsNEQ applies the NEQ predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsNEQ(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsIn applies the In predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsIn(vs ...string) predicate.BaselineMeasureDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDenomUnits), v...))
	})
}

// BaselineMeasureDenomUnitsNotIn applies the NotIn predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsNotIn(vs ...string) predicate.BaselineMeasureDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDenomUnits), v...))
	})
}

// BaselineMeasureDenomUnitsGT applies the GT predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsGT(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsGTE applies the GTE predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsGTE(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsLT applies the LT predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsLT(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsLTE applies the LTE predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsLTE(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsContains applies the Contains predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsContains(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsHasPrefix applies the HasPrefix predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsHasPrefix(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsHasSuffix applies the HasSuffix predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsHasSuffix(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsEqualFold applies the EqualFold predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsEqualFold(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// BaselineMeasureDenomUnitsContainsFold applies the ContainsFold predicate on the "baseline_measure_denom_units" field.
func BaselineMeasureDenomUnitsContainsFold(v string) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDenomUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineMeasure) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
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

// HasBaselineMeasureDenomCountList applies the HasEdge predicate on the "baseline_measure_denom_count_list" edge.
func HasBaselineMeasureDenomCountList() predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureDenomCountListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureDenomCountListTable, BaselineMeasureDenomCountListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineMeasureDenomCountListWith applies the HasEdge predicate on the "baseline_measure_denom_count_list" edge with a given conditions (other predicates).
func HasBaselineMeasureDenomCountListWith(preds ...predicate.BaselineMeasureDenomCount) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureDenomCountListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureDenomCountListTable, BaselineMeasureDenomCountListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineMeasureDenom) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineMeasureDenom) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
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
func Not(p predicate.BaselineMeasureDenom) predicate.BaselineMeasureDenom {
	return predicate.BaselineMeasureDenom(func(s *sql.Selector) {
		p(s.Not())
	})
}
