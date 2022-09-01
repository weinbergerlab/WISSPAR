// Code generated (@generated) by entc, DO NOT EDIT.

package baselineclass

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineClassTitle applies equality check predicate on the "baseline_class_title" field. It's identical to BaselineClassTitleEQ.
func BaselineClassTitle(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleEQ applies the EQ predicate on the "baseline_class_title" field.
func BaselineClassTitleEQ(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleNEQ applies the NEQ predicate on the "baseline_class_title" field.
func BaselineClassTitleNEQ(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleIn applies the In predicate on the "baseline_class_title" field.
func BaselineClassTitleIn(vs ...string) predicate.BaselineClass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineClassTitle), v...))
	})
}

// BaselineClassTitleNotIn applies the NotIn predicate on the "baseline_class_title" field.
func BaselineClassTitleNotIn(vs ...string) predicate.BaselineClass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineClass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineClassTitle), v...))
	})
}

// BaselineClassTitleGT applies the GT predicate on the "baseline_class_title" field.
func BaselineClassTitleGT(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleGTE applies the GTE predicate on the "baseline_class_title" field.
func BaselineClassTitleGTE(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleLT applies the LT predicate on the "baseline_class_title" field.
func BaselineClassTitleLT(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleLTE applies the LTE predicate on the "baseline_class_title" field.
func BaselineClassTitleLTE(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleContains applies the Contains predicate on the "baseline_class_title" field.
func BaselineClassTitleContains(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleHasPrefix applies the HasPrefix predicate on the "baseline_class_title" field.
func BaselineClassTitleHasPrefix(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleHasSuffix applies the HasSuffix predicate on the "baseline_class_title" field.
func BaselineClassTitleHasSuffix(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleEqualFold applies the EqualFold predicate on the "baseline_class_title" field.
func BaselineClassTitleEqualFold(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineClassTitle), v))
	})
}

// BaselineClassTitleContainsFold applies the ContainsFold predicate on the "baseline_class_title" field.
func BaselineClassTitleContainsFold(v string) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineClassTitle), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineMeasure) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
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

// HasBaselineClassDenomList applies the HasEdge predicate on the "baseline_class_denom_list" edge.
func HasBaselineClassDenomList() predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassDenomListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassDenomListTable, BaselineClassDenomListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineClassDenomListWith applies the HasEdge predicate on the "baseline_class_denom_list" edge with a given conditions (other predicates).
func HasBaselineClassDenomListWith(preds ...predicate.BaselineClassDenom) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassDenomListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassDenomListTable, BaselineClassDenomListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineCategoryList applies the HasEdge predicate on the "baseline_category_list" edge.
func HasBaselineCategoryList() predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineCategoryListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineCategoryListTable, BaselineCategoryListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineCategoryListWith applies the HasEdge predicate on the "baseline_category_list" edge with a given conditions (other predicates).
func HasBaselineCategoryListWith(preds ...predicate.BaselineCategory) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineCategoryListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineCategoryListTable, BaselineCategoryListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineClass) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineClass) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
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
func Not(p predicate.BaselineClass) predicate.BaselineClass {
	return predicate.BaselineClass(func(s *sql.Selector) {
		p(s.Not())
	})
}
