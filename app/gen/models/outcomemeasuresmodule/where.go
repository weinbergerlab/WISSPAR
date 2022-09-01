// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasuresmodule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ResultsDefinition) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeMeasureList applies the HasEdge predicate on the "outcome_measure_list" edge.
func HasOutcomeMeasureList() predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasureListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeMeasureListTable, OutcomeMeasureListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeMeasureListWith applies the HasEdge predicate on the "outcome_measure_list" edge with a given conditions (other predicates).
func HasOutcomeMeasureListWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasureListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeMeasureListTable, OutcomeMeasureListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeMeasuresModule) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeMeasuresModule) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeMeasuresModule) predicate.OutcomeMeasuresModule {
	return predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
		p(s.Not())
	})
}
