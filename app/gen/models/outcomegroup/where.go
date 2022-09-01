// Code generated (@generated) by entc, DO NOT EDIT.

package outcomegroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeGroupID applies equality check predicate on the "outcome_group_id" field. It's identical to OutcomeGroupIDEQ.
func OutcomeGroupID(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupTitle applies equality check predicate on the "outcome_group_title" field. It's identical to OutcomeGroupTitleEQ.
func OutcomeGroupTitle(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupDescription applies equality check predicate on the "outcome_group_description" field. It's identical to OutcomeGroupDescriptionEQ.
func OutcomeGroupDescription(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupIDEQ applies the EQ predicate on the "outcome_group_id" field.
func OutcomeGroupIDEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDNEQ applies the NEQ predicate on the "outcome_group_id" field.
func OutcomeGroupIDNEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDIn applies the In predicate on the "outcome_group_id" field.
func OutcomeGroupIDIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeGroupID), v...))
	})
}

// OutcomeGroupIDNotIn applies the NotIn predicate on the "outcome_group_id" field.
func OutcomeGroupIDNotIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeGroupID), v...))
	})
}

// OutcomeGroupIDGT applies the GT predicate on the "outcome_group_id" field.
func OutcomeGroupIDGT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDGTE applies the GTE predicate on the "outcome_group_id" field.
func OutcomeGroupIDGTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDLT applies the LT predicate on the "outcome_group_id" field.
func OutcomeGroupIDLT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDLTE applies the LTE predicate on the "outcome_group_id" field.
func OutcomeGroupIDLTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDContains applies the Contains predicate on the "outcome_group_id" field.
func OutcomeGroupIDContains(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_group_id" field.
func OutcomeGroupIDHasPrefix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_group_id" field.
func OutcomeGroupIDHasSuffix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDEqualFold applies the EqualFold predicate on the "outcome_group_id" field.
func OutcomeGroupIDEqualFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupIDContainsFold applies the ContainsFold predicate on the "outcome_group_id" field.
func OutcomeGroupIDContainsFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeGroupID), v))
	})
}

// OutcomeGroupTitleEQ applies the EQ predicate on the "outcome_group_title" field.
func OutcomeGroupTitleEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleNEQ applies the NEQ predicate on the "outcome_group_title" field.
func OutcomeGroupTitleNEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleIn applies the In predicate on the "outcome_group_title" field.
func OutcomeGroupTitleIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeGroupTitle), v...))
	})
}

// OutcomeGroupTitleNotIn applies the NotIn predicate on the "outcome_group_title" field.
func OutcomeGroupTitleNotIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeGroupTitle), v...))
	})
}

// OutcomeGroupTitleGT applies the GT predicate on the "outcome_group_title" field.
func OutcomeGroupTitleGT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleGTE applies the GTE predicate on the "outcome_group_title" field.
func OutcomeGroupTitleGTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleLT applies the LT predicate on the "outcome_group_title" field.
func OutcomeGroupTitleLT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleLTE applies the LTE predicate on the "outcome_group_title" field.
func OutcomeGroupTitleLTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleContains applies the Contains predicate on the "outcome_group_title" field.
func OutcomeGroupTitleContains(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleHasPrefix applies the HasPrefix predicate on the "outcome_group_title" field.
func OutcomeGroupTitleHasPrefix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleHasSuffix applies the HasSuffix predicate on the "outcome_group_title" field.
func OutcomeGroupTitleHasSuffix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleEqualFold applies the EqualFold predicate on the "outcome_group_title" field.
func OutcomeGroupTitleEqualFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupTitleContainsFold applies the ContainsFold predicate on the "outcome_group_title" field.
func OutcomeGroupTitleContainsFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeGroupTitle), v))
	})
}

// OutcomeGroupDescriptionEQ applies the EQ predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionNEQ applies the NEQ predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionNEQ(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionIn applies the In predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeGroupDescription), v...))
	})
}

// OutcomeGroupDescriptionNotIn applies the NotIn predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionNotIn(vs ...string) predicate.OutcomeGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeGroupDescription), v...))
	})
}

// OutcomeGroupDescriptionGT applies the GT predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionGT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionGTE applies the GTE predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionGTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionLT applies the LT predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionLT(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionLTE applies the LTE predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionLTE(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionContains applies the Contains predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionContains(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionHasPrefix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionHasSuffix(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionEqualFold applies the EqualFold predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionEqualFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeGroupDescription), v))
	})
}

// OutcomeGroupDescriptionContainsFold applies the ContainsFold predicate on the "outcome_group_description" field.
func OutcomeGroupDescriptionContainsFold(v string) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeGroupDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeGroup) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeGroup) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeGroup) predicate.OutcomeGroup {
	return predicate.OutcomeGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
