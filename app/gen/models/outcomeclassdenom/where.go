// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclassdenom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeClassDenomUnits applies equality check predicate on the "outcome_class_denom_units" field. It's identical to OutcomeClassDenomUnitsEQ.
func OutcomeClassDenomUnits(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsEQ applies the EQ predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsEQ(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsNEQ applies the NEQ predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsNEQ(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsIn applies the In predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsIn(vs ...string) predicate.OutcomeClassDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeClassDenomUnits), v...))
	})
}

// OutcomeClassDenomUnitsNotIn applies the NotIn predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsNotIn(vs ...string) predicate.OutcomeClassDenom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeClassDenomUnits), v...))
	})
}

// OutcomeClassDenomUnitsGT applies the GT predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsGT(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsGTE applies the GTE predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsGTE(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsLT applies the LT predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsLT(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsLTE applies the LTE predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsLTE(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsContains applies the Contains predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsContains(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsHasPrefix applies the HasPrefix predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsHasPrefix(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsHasSuffix applies the HasSuffix predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsHasSuffix(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsEqualFold applies the EqualFold predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsEqualFold(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// OutcomeClassDenomUnitsContainsFold applies the ContainsFold predicate on the "outcome_class_denom_units" field.
func OutcomeClassDenomUnitsContainsFold(v string) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeClassDenomUnits), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeClass) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
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

// HasOutcomeClassDenomCountList applies the HasEdge predicate on the "outcome_class_denom_count_list" edge.
func HasOutcomeClassDenomCountList() predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassDenomCountListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassDenomCountListTable, OutcomeClassDenomCountListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeClassDenomCountListWith applies the HasEdge predicate on the "outcome_class_denom_count_list" edge with a given conditions (other predicates).
func HasOutcomeClassDenomCountListWith(preds ...predicate.OutcomeClassDenomCount) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassDenomCountListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassDenomCountListTable, OutcomeClassDenomCountListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeClassDenom) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeClassDenom) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeClassDenom) predicate.OutcomeClassDenom {
	return predicate.OutcomeClassDenom(func(s *sql.Selector) {
		p(s.Not())
	})
}
