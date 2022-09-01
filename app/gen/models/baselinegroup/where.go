// Code generated (@generated) by entc, DO NOT EDIT.

package baselinegroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineGroupID applies equality check predicate on the "baseline_group_id" field. It's identical to BaselineGroupIDEQ.
func BaselineGroupID(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupTitle applies equality check predicate on the "baseline_group_title" field. It's identical to BaselineGroupTitleEQ.
func BaselineGroupTitle(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupDescription applies equality check predicate on the "baseline_group_description" field. It's identical to BaselineGroupDescriptionEQ.
func BaselineGroupDescription(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupIDEQ applies the EQ predicate on the "baseline_group_id" field.
func BaselineGroupIDEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDNEQ applies the NEQ predicate on the "baseline_group_id" field.
func BaselineGroupIDNEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDIn applies the In predicate on the "baseline_group_id" field.
func BaselineGroupIDIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineGroupID), v...))
	})
}

// BaselineGroupIDNotIn applies the NotIn predicate on the "baseline_group_id" field.
func BaselineGroupIDNotIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineGroupID), v...))
	})
}

// BaselineGroupIDGT applies the GT predicate on the "baseline_group_id" field.
func BaselineGroupIDGT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDGTE applies the GTE predicate on the "baseline_group_id" field.
func BaselineGroupIDGTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDLT applies the LT predicate on the "baseline_group_id" field.
func BaselineGroupIDLT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDLTE applies the LTE predicate on the "baseline_group_id" field.
func BaselineGroupIDLTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDContains applies the Contains predicate on the "baseline_group_id" field.
func BaselineGroupIDContains(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDHasPrefix applies the HasPrefix predicate on the "baseline_group_id" field.
func BaselineGroupIDHasPrefix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDHasSuffix applies the HasSuffix predicate on the "baseline_group_id" field.
func BaselineGroupIDHasSuffix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDEqualFold applies the EqualFold predicate on the "baseline_group_id" field.
func BaselineGroupIDEqualFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupIDContainsFold applies the ContainsFold predicate on the "baseline_group_id" field.
func BaselineGroupIDContainsFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineGroupID), v))
	})
}

// BaselineGroupTitleEQ applies the EQ predicate on the "baseline_group_title" field.
func BaselineGroupTitleEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleNEQ applies the NEQ predicate on the "baseline_group_title" field.
func BaselineGroupTitleNEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleIn applies the In predicate on the "baseline_group_title" field.
func BaselineGroupTitleIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineGroupTitle), v...))
	})
}

// BaselineGroupTitleNotIn applies the NotIn predicate on the "baseline_group_title" field.
func BaselineGroupTitleNotIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineGroupTitle), v...))
	})
}

// BaselineGroupTitleGT applies the GT predicate on the "baseline_group_title" field.
func BaselineGroupTitleGT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleGTE applies the GTE predicate on the "baseline_group_title" field.
func BaselineGroupTitleGTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleLT applies the LT predicate on the "baseline_group_title" field.
func BaselineGroupTitleLT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleLTE applies the LTE predicate on the "baseline_group_title" field.
func BaselineGroupTitleLTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleContains applies the Contains predicate on the "baseline_group_title" field.
func BaselineGroupTitleContains(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleHasPrefix applies the HasPrefix predicate on the "baseline_group_title" field.
func BaselineGroupTitleHasPrefix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleHasSuffix applies the HasSuffix predicate on the "baseline_group_title" field.
func BaselineGroupTitleHasSuffix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleEqualFold applies the EqualFold predicate on the "baseline_group_title" field.
func BaselineGroupTitleEqualFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupTitleContainsFold applies the ContainsFold predicate on the "baseline_group_title" field.
func BaselineGroupTitleContainsFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineGroupTitle), v))
	})
}

// BaselineGroupDescriptionEQ applies the EQ predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionNEQ applies the NEQ predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionNEQ(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionIn applies the In predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineGroupDescription), v...))
	})
}

// BaselineGroupDescriptionNotIn applies the NotIn predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionNotIn(vs ...string) predicate.BaselineGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineGroupDescription), v...))
	})
}

// BaselineGroupDescriptionGT applies the GT predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionGT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionGTE applies the GTE predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionGTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionLT applies the LT predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionLT(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionLTE applies the LTE predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionLTE(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionContains applies the Contains predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionContains(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionHasPrefix applies the HasPrefix predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionHasPrefix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionHasSuffix applies the HasSuffix predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionHasSuffix(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionEqualFold applies the EqualFold predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionEqualFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineGroupDescription), v))
	})
}

// BaselineGroupDescriptionContainsFold applies the ContainsFold predicate on the "baseline_group_description" field.
func BaselineGroupDescriptionContainsFold(v string) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineGroupDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineCharacteristicsModule) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
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
func And(predicates ...predicate.BaselineGroup) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineGroup) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
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
func Not(p predicate.BaselineGroup) predicate.BaselineGroup {
	return predicate.BaselineGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
