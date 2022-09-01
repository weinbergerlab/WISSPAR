// Code generated (@generated) by entc, DO NOT EDIT.

package resultsdefinition

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ClinicalTrial) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
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

// HasParticipantFlowModule applies the HasEdge predicate on the "participant_flow_module" edge.
func HasParticipantFlowModule() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantFlowModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ParticipantFlowModuleTable, ParticipantFlowModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantFlowModuleWith applies the HasEdge predicate on the "participant_flow_module" edge with a given conditions (other predicates).
func HasParticipantFlowModuleWith(preds ...predicate.ParticipantFlowModule) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantFlowModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ParticipantFlowModuleTable, ParticipantFlowModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineCharacteristicsModule applies the HasEdge predicate on the "baseline_characteristics_module" edge.
func HasBaselineCharacteristicsModule() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineCharacteristicsModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BaselineCharacteristicsModuleTable, BaselineCharacteristicsModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineCharacteristicsModuleWith applies the HasEdge predicate on the "baseline_characteristics_module" edge with a given conditions (other predicates).
func HasBaselineCharacteristicsModuleWith(preds ...predicate.BaselineCharacteristicsModule) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineCharacteristicsModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BaselineCharacteristicsModuleTable, BaselineCharacteristicsModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeMeasuresModule applies the HasEdge predicate on the "outcome_measures_module" edge.
func HasOutcomeMeasuresModule() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasuresModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OutcomeMeasuresModuleTable, OutcomeMeasuresModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeMeasuresModuleWith applies the HasEdge predicate on the "outcome_measures_module" edge with a given conditions (other predicates).
func HasOutcomeMeasuresModuleWith(preds ...predicate.OutcomeMeasuresModule) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeMeasuresModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OutcomeMeasuresModuleTable, OutcomeMeasuresModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdverseEventsModule applies the HasEdge predicate on the "adverse_events_module" edge.
func HasAdverseEventsModule() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdverseEventsModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AdverseEventsModuleTable, AdverseEventsModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdverseEventsModuleWith applies the HasEdge predicate on the "adverse_events_module" edge with a given conditions (other predicates).
func HasAdverseEventsModuleWith(preds ...predicate.AdverseEventsModule) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdverseEventsModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AdverseEventsModuleTable, AdverseEventsModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMoreInfoModule applies the HasEdge predicate on the "more_info_module" edge.
func HasMoreInfoModule() predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoreInfoModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MoreInfoModuleTable, MoreInfoModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMoreInfoModuleWith applies the HasEdge predicate on the "more_info_module" edge with a given conditions (other predicates).
func HasMoreInfoModuleWith(preds ...predicate.MoreInfoModule) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoreInfoModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MoreInfoModuleTable, MoreInfoModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ResultsDefinition) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ResultsDefinition) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
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
func Not(p predicate.ResultsDefinition) predicate.ResultsDefinition {
	return predicate.ResultsDefinition(func(s *sql.Selector) {
		p(s.Not())
	})
}
