// Code generated (@generated) by entc, DO NOT EDIT.

package participantflowmodule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlowPreAssignmentDetails applies equality check predicate on the "flow_pre_assignment_details" field. It's identical to FlowPreAssignmentDetailsEQ.
func FlowPreAssignmentDetails(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowRecruitmentDetails applies equality check predicate on the "flow_recruitment_details" field. It's identical to FlowRecruitmentDetailsEQ.
func FlowRecruitmentDetails(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowTypeUnitsAnalyzed applies equality check predicate on the "flow_type_units_analyzed" field. It's identical to FlowTypeUnitsAnalyzedEQ.
func FlowTypeUnitsAnalyzed(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowPreAssignmentDetailsEQ applies the EQ predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsNEQ applies the NEQ predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsNEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsIn applies the In predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowPreAssignmentDetails), v...))
	})
}

// FlowPreAssignmentDetailsNotIn applies the NotIn predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsNotIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowPreAssignmentDetails), v...))
	})
}

// FlowPreAssignmentDetailsGT applies the GT predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsGT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsGTE applies the GTE predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsGTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsLT applies the LT predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsLT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsLTE applies the LTE predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsLTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsContains applies the Contains predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsContains(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsHasPrefix applies the HasPrefix predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsHasPrefix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsHasSuffix applies the HasSuffix predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsHasSuffix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsEqualFold applies the EqualFold predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsEqualFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowPreAssignmentDetailsContainsFold applies the ContainsFold predicate on the "flow_pre_assignment_details" field.
func FlowPreAssignmentDetailsContainsFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowPreAssignmentDetails), v))
	})
}

// FlowRecruitmentDetailsEQ applies the EQ predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsNEQ applies the NEQ predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsNEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsIn applies the In predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowRecruitmentDetails), v...))
	})
}

// FlowRecruitmentDetailsNotIn applies the NotIn predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsNotIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowRecruitmentDetails), v...))
	})
}

// FlowRecruitmentDetailsGT applies the GT predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsGT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsGTE applies the GTE predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsGTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsLT applies the LT predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsLT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsLTE applies the LTE predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsLTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsContains applies the Contains predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsContains(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsHasPrefix applies the HasPrefix predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsHasPrefix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsHasSuffix applies the HasSuffix predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsHasSuffix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsEqualFold applies the EqualFold predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsEqualFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowRecruitmentDetailsContainsFold applies the ContainsFold predicate on the "flow_recruitment_details" field.
func FlowRecruitmentDetailsContainsFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowRecruitmentDetails), v))
	})
}

// FlowTypeUnitsAnalyzedEQ applies the EQ predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedNEQ applies the NEQ predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedNEQ(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedIn applies the In predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFlowTypeUnitsAnalyzed), v...))
	})
}

// FlowTypeUnitsAnalyzedNotIn applies the NotIn predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedNotIn(vs ...string) predicate.ParticipantFlowModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFlowTypeUnitsAnalyzed), v...))
	})
}

// FlowTypeUnitsAnalyzedGT applies the GT predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedGT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedGTE applies the GTE predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedGTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedLT applies the LT predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedLT(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedLTE applies the LTE predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedLTE(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedContains applies the Contains predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedContains(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedHasPrefix applies the HasPrefix predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedHasPrefix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedHasSuffix applies the HasSuffix predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedHasSuffix(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedEqualFold applies the EqualFold predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedEqualFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// FlowTypeUnitsAnalyzedContainsFold applies the ContainsFold predicate on the "flow_type_units_analyzed" field.
func FlowTypeUnitsAnalyzedContainsFold(v string) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlowTypeUnitsAnalyzed), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ResultsDefinition) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
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

// HasFlowGroupList applies the HasEdge predicate on the "flow_group_list" edge.
func HasFlowGroupList() predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowGroupListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowGroupListTable, FlowGroupListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowGroupListWith applies the HasEdge predicate on the "flow_group_list" edge with a given conditions (other predicates).
func HasFlowGroupListWith(preds ...predicate.FlowGroup) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowGroupListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowGroupListTable, FlowGroupListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFlowPeriodList applies the HasEdge predicate on the "flow_period_list" edge.
func HasFlowPeriodList() predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowPeriodListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowPeriodListTable, FlowPeriodListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlowPeriodListWith applies the HasEdge predicate on the "flow_period_list" edge with a given conditions (other predicates).
func HasFlowPeriodListWith(preds ...predicate.FlowPeriod) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlowPeriodListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlowPeriodListTable, FlowPeriodListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ParticipantFlowModule) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ParticipantFlowModule) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
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
func Not(p predicate.ParticipantFlowModule) predicate.ParticipantFlowModule {
	return predicate.ParticipantFlowModule(func(s *sql.Selector) {
		p(s.Not())
	})
}
