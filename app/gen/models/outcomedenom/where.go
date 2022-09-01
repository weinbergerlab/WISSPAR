// Code generated (@generated) by entc, DO NOT EDIT.

package outcomedenom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeDenomUnits applies equality check predicate on the "outcome_denom_units" field. It's identical to OutcomeDenomUnitsEQ.
func OutcomeDenomUnits(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsEQ applies the EQ predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsEQ(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsNEQ applies the NEQ predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsNEQ(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsIn applies the In predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsIn(vs ...string) predicate.OutcomeDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeDenomUnits), v...))
	})
}

// OutcomeDenomUnitsNotIn applies the NotIn predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsNotIn(vs ...string) predicate.OutcomeDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeDenomUnits), v...))
	})
}

// OutcomeDenomUnitsGT applies the GT predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsGT(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsGTE applies the GTE predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsGTE(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsLT applies the LT predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsLT(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsLTE applies the LTE predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsLTE(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsContains applies the Contains predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsContains(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsHasPrefix applies the HasPrefix predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsHasPrefix(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsHasSuffix applies the HasSuffix predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsHasSuffix(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsEqualFold applies the EqualFold predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsEqualFold(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeDenomUnits), v))
	})
}

// OutcomeDenomUnitsContainsFold applies the ContainsFold predicate on the "outcome_denom_units" field.
func OutcomeDenomUnitsContainsFold(v string) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeDenomUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
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

// HasOutcomeDenomCountList applies the HasEdge predicate on the "outcome_denom_count_list" edge.
func HasOutcomeDenomCountList() predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeDenomCountListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeDenomCountListTable, OutcomeDenomCountListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeDenomCountListWith applies the HasEdge predicate on the "outcome_denom_count_list" edge with a given conditions (other predicates).
func HasOutcomeDenomCountListWith(preds ...predicate.OutcomeDenomCount) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeDenomCountListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeDenomCountListTable, OutcomeDenomCountListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeDenom) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeDenom) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeDenom) predicate.OutcomeDenom {
	return predicate.OutcomeDenom(func(s *sql.Selector) {
		p(s.Not())
	})
}
