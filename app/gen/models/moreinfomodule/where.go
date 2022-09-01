// Code generated (@generated) by entc, DO NOT EDIT.

package moreinfomodule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LimitationsAndCaveatsDescription applies equality check predicate on the "limitations_and_caveats_description" field. It's identical to LimitationsAndCaveatsDescriptionEQ.
func LimitationsAndCaveatsDescription(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionEQ applies the EQ predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionEQ(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionNEQ applies the NEQ predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionNEQ(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionIn applies the In predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionIn(vs ...string) predicate.MoreInfoModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLimitationsAndCaveatsDescription), v...))
	})
}

// LimitationsAndCaveatsDescriptionNotIn applies the NotIn predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionNotIn(vs ...string) predicate.MoreInfoModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLimitationsAndCaveatsDescription), v...))
	})
}

// LimitationsAndCaveatsDescriptionGT applies the GT predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionGT(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionGTE applies the GTE predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionGTE(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionLT applies the LT predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionLT(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionLTE applies the LTE predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionLTE(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionContains applies the Contains predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionContains(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionHasPrefix applies the HasPrefix predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionHasPrefix(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionHasSuffix applies the HasSuffix predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionHasSuffix(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionEqualFold applies the EqualFold predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionEqualFold(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// LimitationsAndCaveatsDescriptionContainsFold applies the ContainsFold predicate on the "limitations_and_caveats_description" field.
func LimitationsAndCaveatsDescriptionContainsFold(v string) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLimitationsAndCaveatsDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ResultsDefinition) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
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

// HasCertainAgreement applies the HasEdge predicate on the "certain_agreement" edge.
func HasCertainAgreement() predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertainAgreementTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CertainAgreementTable, CertainAgreementColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertainAgreementWith applies the HasEdge predicate on the "certain_agreement" edge with a given conditions (other predicates).
func HasCertainAgreementWith(preds ...predicate.CertainAgreement) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertainAgreementInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CertainAgreementTable, CertainAgreementColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPointOfContact applies the HasEdge predicate on the "point_of_contact" edge.
func HasPointOfContact() predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PointOfContactTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PointOfContactTable, PointOfContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPointOfContactWith applies the HasEdge predicate on the "point_of_contact" edge with a given conditions (other predicates).
func HasPointOfContactWith(preds ...predicate.PointOfContact) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PointOfContactInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PointOfContactTable, PointOfContactColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MoreInfoModule) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MoreInfoModule) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
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
func Not(p predicate.MoreInfoModule) predicate.MoreInfoModule {
	return predicate.MoreInfoModule(func(s *sql.Selector) {
		p(s.Not())
	})
}
