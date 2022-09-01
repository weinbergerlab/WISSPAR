// Code generated (@generated) by entc, DO NOT EDIT.

package baselinedenom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineDenomUnits applies equality check predicate on the "baseline_denom_units" field. It's identical to BaselineDenomUnitsEQ.
func BaselineDenomUnits(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsEQ applies the EQ predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsEQ(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsNEQ applies the NEQ predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsNEQ(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsIn applies the In predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsIn(vs ...string) predicate.BaselineDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineDenomUnits), v...))
	})
}

// BaselineDenomUnitsNotIn applies the NotIn predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsNotIn(vs ...string) predicate.BaselineDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineDenomUnits), v...))
	})
}

// BaselineDenomUnitsGT applies the GT predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsGT(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsGTE applies the GTE predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsGTE(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsLT applies the LT predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsLT(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsLTE applies the LTE predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsLTE(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsContains applies the Contains predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsContains(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsHasPrefix applies the HasPrefix predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsHasPrefix(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsHasSuffix applies the HasSuffix predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsHasSuffix(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsEqualFold applies the EqualFold predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsEqualFold(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineDenomUnits), v))
	})
}

// BaselineDenomUnitsContainsFold applies the ContainsFold predicate on the "baseline_denom_units" field.
func BaselineDenomUnitsContainsFold(v string) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineDenomUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineCharacteristicsModule) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
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

// HasBaselineDenomCountList applies the HasEdge predicate on the "baseline_denom_count_list" edge.
func HasBaselineDenomCountList() predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineDenomCountListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineDenomCountListTable, BaselineDenomCountListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineDenomCountListWith applies the HasEdge predicate on the "baseline_denom_count_list" edge with a given conditions (other predicates).
func HasBaselineDenomCountListWith(preds ...predicate.BaselineDenomCount) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineDenomCountListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineDenomCountListTable, BaselineDenomCountListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineDenom) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineDenom) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
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
func Not(p predicate.BaselineDenom) predicate.BaselineDenom {
	return predicate.BaselineDenom(func(s *sql.Selector) {
		p(s.Not())
	})
}
