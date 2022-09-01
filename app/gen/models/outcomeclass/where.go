// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeclass

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeClassTitle applies equality check predicate on the "outcome_class_title" field. It's identical to OutcomeClassTitleEQ.
func OutcomeClassTitle(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleEQ applies the EQ predicate on the "outcome_class_title" field.
func OutcomeClassTitleEQ(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleNEQ applies the NEQ predicate on the "outcome_class_title" field.
func OutcomeClassTitleNEQ(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleIn applies the In predicate on the "outcome_class_title" field.
func OutcomeClassTitleIn(vs ...string) predicate.OutcomeClass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeClassTitle), v...))
	})
}

// OutcomeClassTitleNotIn applies the NotIn predicate on the "outcome_class_title" field.
func OutcomeClassTitleNotIn(vs ...string) predicate.OutcomeClass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeClass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeClassTitle), v...))
	})
}

// OutcomeClassTitleGT applies the GT predicate on the "outcome_class_title" field.
func OutcomeClassTitleGT(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleGTE applies the GTE predicate on the "outcome_class_title" field.
func OutcomeClassTitleGTE(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleLT applies the LT predicate on the "outcome_class_title" field.
func OutcomeClassTitleLT(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleLTE applies the LTE predicate on the "outcome_class_title" field.
func OutcomeClassTitleLTE(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleContains applies the Contains predicate on the "outcome_class_title" field.
func OutcomeClassTitleContains(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleHasPrefix applies the HasPrefix predicate on the "outcome_class_title" field.
func OutcomeClassTitleHasPrefix(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleHasSuffix applies the HasSuffix predicate on the "outcome_class_title" field.
func OutcomeClassTitleHasSuffix(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleEqualFold applies the EqualFold predicate on the "outcome_class_title" field.
func OutcomeClassTitleEqualFold(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeClassTitle), v))
	})
}

// OutcomeClassTitleContainsFold applies the ContainsFold predicate on the "outcome_class_title" field.
func OutcomeClassTitleContainsFold(v string) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeClassTitle), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
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

// HasOutcomeClassDenomList applies the HasEdge predicate on the "outcome_class_denom_list" edge.
func HasOutcomeClassDenomList() predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassDenomListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassDenomListTable, OutcomeClassDenomListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeClassDenomListWith applies the HasEdge predicate on the "outcome_class_denom_list" edge with a given conditions (other predicates).
func HasOutcomeClassDenomListWith(preds ...predicate.OutcomeClassDenom) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassDenomListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassDenomListTable, OutcomeClassDenomListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeCategoryList applies the HasEdge predicate on the "outcome_category_list" edge.
func HasOutcomeCategoryList() predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeCategoryListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeCategoryListTable, OutcomeCategoryListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeCategoryListWith applies the HasEdge predicate on the "outcome_category_list" edge with a given conditions (other predicates).
func HasOutcomeCategoryListWith(preds ...predicate.OutcomeCategory) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeCategoryListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeCategoryListTable, OutcomeCategoryListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeClass) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeClass) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeClass) predicate.OutcomeClass {
	return predicate.OutcomeClass(func(s *sql.Selector) {
		p(s.Not())
	})
}
