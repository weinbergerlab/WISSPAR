// Code generated (@generated) by entc, DO NOT EDIT.

package baselinecategory

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineCategoryTitle applies equality check predicate on the "baseline_category_title" field. It's identical to BaselineCategoryTitleEQ.
func BaselineCategoryTitle(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleEQ applies the EQ predicate on the "baseline_category_title" field.
func BaselineCategoryTitleEQ(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleNEQ applies the NEQ predicate on the "baseline_category_title" field.
func BaselineCategoryTitleNEQ(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleIn applies the In predicate on the "baseline_category_title" field.
func BaselineCategoryTitleIn(vs ...string) predicate.BaselineCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCategory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineCategoryTitle), v...))
	})
}

// BaselineCategoryTitleNotIn applies the NotIn predicate on the "baseline_category_title" field.
func BaselineCategoryTitleNotIn(vs ...string) predicate.BaselineCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCategory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineCategoryTitle), v...))
	})
}

// BaselineCategoryTitleGT applies the GT predicate on the "baseline_category_title" field.
func BaselineCategoryTitleGT(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleGTE applies the GTE predicate on the "baseline_category_title" field.
func BaselineCategoryTitleGTE(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleLT applies the LT predicate on the "baseline_category_title" field.
func BaselineCategoryTitleLT(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleLTE applies the LTE predicate on the "baseline_category_title" field.
func BaselineCategoryTitleLTE(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleContains applies the Contains predicate on the "baseline_category_title" field.
func BaselineCategoryTitleContains(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleHasPrefix applies the HasPrefix predicate on the "baseline_category_title" field.
func BaselineCategoryTitleHasPrefix(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleHasSuffix applies the HasSuffix predicate on the "baseline_category_title" field.
func BaselineCategoryTitleHasSuffix(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleEqualFold applies the EqualFold predicate on the "baseline_category_title" field.
func BaselineCategoryTitleEqualFold(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineCategoryTitle), v))
	})
}

// BaselineCategoryTitleContainsFold applies the ContainsFold predicate on the "baseline_category_title" field.
func BaselineCategoryTitleContainsFold(v string) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineCategoryTitle), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineClass) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
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

// HasBaselineMeasurementList applies the HasEdge predicate on the "baseline_measurement_list" edge.
func HasBaselineMeasurementList() predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasurementListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasurementListTable, BaselineMeasurementListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineMeasurementListWith applies the HasEdge predicate on the "baseline_measurement_list" edge with a given conditions (other predicates).
func HasBaselineMeasurementListWith(preds ...predicate.BaselineMeasurement) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasurementListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasurementListTable, BaselineMeasurementListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineCategory) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineCategory) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
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
func Not(p predicate.BaselineCategory) predicate.BaselineCategory {
	return predicate.BaselineCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}
