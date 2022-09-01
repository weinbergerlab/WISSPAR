// Code generated (@generated) by entc, DO NOT EDIT.

package adverseeventsmodule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EventsFrequencyThreshold applies equality check predicate on the "events_frequency_threshold" field. It's identical to EventsFrequencyThresholdEQ.
func EventsFrequencyThreshold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsTimeFrame applies equality check predicate on the "events_time_frame" field. It's identical to EventsTimeFrameEQ.
func EventsTimeFrame(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsDescription applies equality check predicate on the "events_description" field. It's identical to EventsDescriptionEQ.
func EventsDescription(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsDescription), v))
	})
}

// EventsFrequencyThresholdEQ applies the EQ predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdNEQ applies the NEQ predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdNEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdIn applies the In predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventsFrequencyThreshold), v...))
	})
}

// EventsFrequencyThresholdNotIn applies the NotIn predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdNotIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventsFrequencyThreshold), v...))
	})
}

// EventsFrequencyThresholdGT applies the GT predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdGT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdGTE applies the GTE predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdGTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdLT applies the LT predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdLT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdLTE applies the LTE predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdLTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdContains applies the Contains predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdContains(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdHasPrefix applies the HasPrefix predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdHasPrefix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdHasSuffix applies the HasSuffix predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdHasSuffix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdEqualFold applies the EqualFold predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdEqualFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsFrequencyThresholdContainsFold applies the ContainsFold predicate on the "events_frequency_threshold" field.
func EventsFrequencyThresholdContainsFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventsFrequencyThreshold), v))
	})
}

// EventsTimeFrameEQ applies the EQ predicate on the "events_time_frame" field.
func EventsTimeFrameEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameNEQ applies the NEQ predicate on the "events_time_frame" field.
func EventsTimeFrameNEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameIn applies the In predicate on the "events_time_frame" field.
func EventsTimeFrameIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventsTimeFrame), v...))
	})
}

// EventsTimeFrameNotIn applies the NotIn predicate on the "events_time_frame" field.
func EventsTimeFrameNotIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventsTimeFrame), v...))
	})
}

// EventsTimeFrameGT applies the GT predicate on the "events_time_frame" field.
func EventsTimeFrameGT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameGTE applies the GTE predicate on the "events_time_frame" field.
func EventsTimeFrameGTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameLT applies the LT predicate on the "events_time_frame" field.
func EventsTimeFrameLT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameLTE applies the LTE predicate on the "events_time_frame" field.
func EventsTimeFrameLTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameContains applies the Contains predicate on the "events_time_frame" field.
func EventsTimeFrameContains(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameHasPrefix applies the HasPrefix predicate on the "events_time_frame" field.
func EventsTimeFrameHasPrefix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameHasSuffix applies the HasSuffix predicate on the "events_time_frame" field.
func EventsTimeFrameHasSuffix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameEqualFold applies the EqualFold predicate on the "events_time_frame" field.
func EventsTimeFrameEqualFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsTimeFrameContainsFold applies the ContainsFold predicate on the "events_time_frame" field.
func EventsTimeFrameContainsFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventsTimeFrame), v))
	})
}

// EventsDescriptionEQ applies the EQ predicate on the "events_description" field.
func EventsDescriptionEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionNEQ applies the NEQ predicate on the "events_description" field.
func EventsDescriptionNEQ(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionIn applies the In predicate on the "events_description" field.
func EventsDescriptionIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventsDescription), v...))
	})
}

// EventsDescriptionNotIn applies the NotIn predicate on the "events_description" field.
func EventsDescriptionNotIn(vs ...string) predicate.AdverseEventsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventsDescription), v...))
	})
}

// EventsDescriptionGT applies the GT predicate on the "events_description" field.
func EventsDescriptionGT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionGTE applies the GTE predicate on the "events_description" field.
func EventsDescriptionGTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionLT applies the LT predicate on the "events_description" field.
func EventsDescriptionLT(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionLTE applies the LTE predicate on the "events_description" field.
func EventsDescriptionLTE(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionContains applies the Contains predicate on the "events_description" field.
func EventsDescriptionContains(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionHasPrefix applies the HasPrefix predicate on the "events_description" field.
func EventsDescriptionHasPrefix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionHasSuffix applies the HasSuffix predicate on the "events_description" field.
func EventsDescriptionHasSuffix(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionEqualFold applies the EqualFold predicate on the "events_description" field.
func EventsDescriptionEqualFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventsDescription), v))
	})
}

// EventsDescriptionContainsFold applies the ContainsFold predicate on the "events_description" field.
func EventsDescriptionContainsFold(v string) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventsDescription), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ResultsDefinition) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
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

// HasEventGroupList applies the HasEdge predicate on the "event_group_list" edge.
func HasEventGroupList() predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventGroupListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventGroupListTable, EventGroupListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventGroupListWith applies the HasEdge predicate on the "event_group_list" edge with a given conditions (other predicates).
func HasEventGroupListWith(preds ...predicate.EventGroup) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventGroupListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventGroupListTable, EventGroupListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSeriousEventList applies the HasEdge predicate on the "serious_event_list" edge.
func HasSeriousEventList() predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriousEventListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeriousEventListTable, SeriousEventListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeriousEventListWith applies the HasEdge predicate on the "serious_event_list" edge with a given conditions (other predicates).
func HasSeriousEventListWith(preds ...predicate.SeriousEvent) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriousEventListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeriousEventListTable, SeriousEventListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOtherEventList applies the HasEdge predicate on the "other_event_list" edge.
func HasOtherEventList() predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OtherEventListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OtherEventListTable, OtherEventListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOtherEventListWith applies the HasEdge predicate on the "other_event_list" edge with a given conditions (other predicates).
func HasOtherEventListWith(preds ...predicate.OtherEvent) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OtherEventListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OtherEventListTable, OtherEventListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdverseEventsModule) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdverseEventsModule) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
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
func Not(p predicate.AdverseEventsModule) predicate.AdverseEventsModule {
	return predicate.AdverseEventsModule(func(s *sql.Selector) {
		p(s.Not())
	})
}
