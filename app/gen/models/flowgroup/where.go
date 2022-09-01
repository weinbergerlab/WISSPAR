// Code generated (@generated) by entc, DO NOT EDIT.

package flowgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowGroupID applies equality check predicate on the "flow_group_id" field. It's identical to FlowGroupIDEQ.
func FlowGroupID(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupTitle applies equality check predicate on the "flow_group_title" field. It's identical to FlowGroupTitleEQ.
func FlowGroupTitle(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupDescription applies equality check predicate on the "flow_group_description" field. It's identical to FlowGroupDescriptionEQ.
func FlowGroupDescription(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupIDEQ applies the EQ predicate on the "flow_group_id" field.
func FlowGroupIDEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDNEQ applies the NEQ predicate on the "flow_group_id" field.
func FlowGroupIDNEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDIn applies the In predicate on the "flow_group_id" field.
func FlowGroupIDIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowGroupID), v...))
	})
}

// FlowGroupIDNotIn applies the NotIn predicate on the "flow_group_id" field.
func FlowGroupIDNotIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowGroupID), v...))
	})
}

// FlowGroupIDGT applies the GT predicate on the "flow_group_id" field.
func FlowGroupIDGT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDGTE applies the GTE predicate on the "flow_group_id" field.
func FlowGroupIDGTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDLT applies the LT predicate on the "flow_group_id" field.
func FlowGroupIDLT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDLTE applies the LTE predicate on the "flow_group_id" field.
func FlowGroupIDLTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDContains applies the Contains predicate on the "flow_group_id" field.
func FlowGroupIDContains(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDHasPrefix applies the HasPrefix predicate on the "flow_group_id" field.
func FlowGroupIDHasPrefix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDHasSuffix applies the HasSuffix predicate on the "flow_group_id" field.
func FlowGroupIDHasSuffix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDEqualFold applies the EqualFold predicate on the "flow_group_id" field.
func FlowGroupIDEqualFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupIDContainsFold applies the ContainsFold predicate on the "flow_group_id" field.
func FlowGroupIDContainsFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowGroupID), v))
	})
}

// FlowGroupTitleEQ applies the EQ predicate on the "flow_group_title" field.
func FlowGroupTitleEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleNEQ applies the NEQ predicate on the "flow_group_title" field.
func FlowGroupTitleNEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleIn applies the In predicate on the "flow_group_title" field.
func FlowGroupTitleIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowGroupTitle), v...))
	})
}

// FlowGroupTitleNotIn applies the NotIn predicate on the "flow_group_title" field.
func FlowGroupTitleNotIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowGroupTitle), v...))
	})
}

// FlowGroupTitleGT applies the GT predicate on the "flow_group_title" field.
func FlowGroupTitleGT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleGTE applies the GTE predicate on the "flow_group_title" field.
func FlowGroupTitleGTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleLT applies the LT predicate on the "flow_group_title" field.
func FlowGroupTitleLT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleLTE applies the LTE predicate on the "flow_group_title" field.
func FlowGroupTitleLTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleContains applies the Contains predicate on the "flow_group_title" field.
func FlowGroupTitleContains(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleHasPrefix applies the HasPrefix predicate on the "flow_group_title" field.
func FlowGroupTitleHasPrefix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleHasSuffix applies the HasSuffix predicate on the "flow_group_title" field.
func FlowGroupTitleHasSuffix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleEqualFold applies the EqualFold predicate on the "flow_group_title" field.
func FlowGroupTitleEqualFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupTitleContainsFold applies the ContainsFold predicate on the "flow_group_title" field.
func FlowGroupTitleContainsFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowGroupTitle), v))
	})
}

// FlowGroupDescriptionEQ applies the EQ predicate on the "flow_group_description" field.
func FlowGroupDescriptionEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionNEQ applies the NEQ predicate on the "flow_group_description" field.
func FlowGroupDescriptionNEQ(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionIn applies the In predicate on the "flow_group_description" field.
func FlowGroupDescriptionIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowGroupDescription), v...))
	})
}

// FlowGroupDescriptionNotIn applies the NotIn predicate on the "flow_group_description" field.
func FlowGroupDescriptionNotIn(vs ...string) predicate.FlowGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlowGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowGroupDescription), v...))
	})
}

// FlowGroupDescriptionGT applies the GT predicate on the "flow_group_description" field.
func FlowGroupDescriptionGT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionGTE applies the GTE predicate on the "flow_group_description" field.
func FlowGroupDescriptionGTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionLT applies the LT predicate on the "flow_group_description" field.
func FlowGroupDescriptionLT(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionLTE applies the LTE predicate on the "flow_group_description" field.
func FlowGroupDescriptionLTE(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionContains applies the Contains predicate on the "flow_group_description" field.
func FlowGroupDescriptionContains(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionHasPrefix applies the HasPrefix predicate on the "flow_group_description" field.
func FlowGroupDescriptionHasPrefix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionHasSuffix applies the HasSuffix predicate on the "flow_group_description" field.
func FlowGroupDescriptionHasSuffix(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionEqualFold applies the EqualFold predicate on the "flow_group_description" field.
func FlowGroupDescriptionEqualFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowGroupDescription), v))
	})
}

// FlowGroupDescriptionContainsFold applies the ContainsFold predicate on the "flow_group_description" field.
func FlowGroupDescriptionContainsFold(v string) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowGroupDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ParticipantFlowModule) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
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
func And(predicates ...predicate.FlowGroup) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlowGroup) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
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
func Not(p predicate.FlowGroup) predicate.FlowGroup {
	return predicate.FlowGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
