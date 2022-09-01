// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclassdenom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineClassDenomUnits applies equality check predicate on the "baseline_class_denom_units" field. It's identical to BaselineClassDenomUnitsEQ.
func BaselineClassDenomUnits(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsEQ applies the EQ predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsEQ(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsNEQ applies the NEQ predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsNEQ(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsIn applies the In predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsIn(vs ...string) predicate.BaselineClassDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineClassDenomUnits), v...))
	})
}

// BaselineClassDenomUnitsNotIn applies the NotIn predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsNotIn(vs ...string) predicate.BaselineClassDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineClassDenomUnits), v...))
	})
}

// BaselineClassDenomUnitsGT applies the GT predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsGT(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsGTE applies the GTE predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsGTE(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsLT applies the LT predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsLT(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsLTE applies the LTE predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsLTE(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsContains applies the Contains predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsContains(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsHasPrefix applies the HasPrefix predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsHasPrefix(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsHasSuffix applies the HasSuffix predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsHasSuffix(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsEqualFold applies the EqualFold predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsEqualFold(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// BaselineClassDenomUnitsContainsFold applies the ContainsFold predicate on the "baseline_class_denom_units" field.
func BaselineClassDenomUnitsContainsFold(v string) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineClassDenomUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineClass) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
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

// HasBaselineClassDenomCountList applies the HasEdge predicate on the "baseline_class_denom_count_list" edge.
func HasBaselineClassDenomCountList() predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassDenomCountListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassDenomCountListTable, BaselineClassDenomCountListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineClassDenomCountListWith applies the HasEdge predicate on the "baseline_class_denom_count_list" edge with a given conditions (other predicates).
func HasBaselineClassDenomCountListWith(preds ...predicate.BaselineClassDenomCount) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassDenomCountListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassDenomCountListTable, BaselineClassDenomCountListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineClassDenom) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineClassDenom) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
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
func Not(p predicate.BaselineClassDenom) predicate.BaselineClassDenom {
	return predicate.BaselineClassDenom(func(s *sql.Selector) {
		p(s.Not())
	})
}
