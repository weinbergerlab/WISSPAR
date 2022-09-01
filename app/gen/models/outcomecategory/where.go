// Code generated (@generated) by entc, DO NOT EDIT.

package outcomecategory

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeCategoryTitle applies equality check predicate on the "outcome_category_title" field. It's identical to OutcomeCategoryTitleEQ.
func OutcomeCategoryTitle(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleEQ applies the EQ predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleEQ(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleNEQ applies the NEQ predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleNEQ(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleIn applies the In predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleIn(vs ...string) predicate.OutcomeCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeCategoryTitle), v...))
	})
}

// OutcomeCategoryTitleNotIn applies the NotIn predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleNotIn(vs ...string) predicate.OutcomeCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeCategoryTitle), v...))
	})
}

// OutcomeCategoryTitleGT applies the GT predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleGT(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleGTE applies the GTE predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleGTE(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleLT applies the LT predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleLT(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleLTE applies the LTE predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleLTE(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleContains applies the Contains predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleContains(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleHasPrefix applies the HasPrefix predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleHasPrefix(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleHasSuffix applies the HasSuffix predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleHasSuffix(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleEqualFold applies the EqualFold predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleEqualFold(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// OutcomeCategoryTitleContainsFold applies the ContainsFold predicate on the "outcome_category_title" field.
func OutcomeCategoryTitleContainsFold(v string) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeCategoryTitle), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeClass) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
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

// HasOutcomeMeasurementList applies the HasEdge predicate on the "outcome_measurement_list" edge.
func HasOutcomeMeasurementList() predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasurementListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeMeasurementListTable, OutcomeMeasurementListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeMeasurementListWith applies the HasEdge predicate on the "outcome_measurement_list" edge with a given conditions (other predicates).
func HasOutcomeMeasurementListWith(preds ...predicate.OutcomeMeasurement) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasurementListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeMeasurementListTable, OutcomeMeasurementListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeCategory) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeCategory) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeCategory) predicate.OutcomeCategory {
	return predicate.OutcomeCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}
