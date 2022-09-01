// Code generated (@generated) by entc, DO NOT EDIT.

package seriouseventstats

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
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
func IDGT(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SeriousEventStatsGroupID applies equality check predicate on the "serious_event_stats_group_id" field. It's identical to SeriousEventStatsGroupIDEQ.
func SeriousEventStatsGroupID(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsNumEvents applies equality check predicate on the "serious_event_stats_num_events" field. It's identical to SeriousEventStatsNumEventsEQ.
func SeriousEventStatsNumEvents(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumAffected applies equality check predicate on the "serious_event_stats_num_affected" field. It's identical to SeriousEventStatsNumAffectedEQ.
func SeriousEventStatsNumAffected(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAtRisk applies equality check predicate on the "serious_event_stats_num_at_risk" field. It's identical to SeriousEventStatsNumAtRiskEQ.
func SeriousEventStatsNumAtRisk(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsGroupIDEQ applies the EQ predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDNEQ applies the NEQ predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDNEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDIn applies the In predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventStatsGroupID), v...))
	})
}

// SeriousEventStatsGroupIDNotIn applies the NotIn predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDNotIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventStatsGroupID), v...))
	})
}

// SeriousEventStatsGroupIDGT applies the GT predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDGT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDGTE applies the GTE predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDGTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDLT applies the LT predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDLT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDLTE applies the LTE predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDLTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDContains applies the Contains predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDContains(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDHasPrefix applies the HasPrefix predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDHasPrefix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDHasSuffix applies the HasSuffix predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDHasSuffix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDEqualFold applies the EqualFold predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDEqualFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsGroupIDContainsFold applies the ContainsFold predicate on the "serious_event_stats_group_id" field.
func SeriousEventStatsGroupIDContainsFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventStatsGroupID), v))
	})
}

// SeriousEventStatsNumEventsEQ applies the EQ predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsNEQ applies the NEQ predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsNEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsIn applies the In predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventStatsNumEvents), v...))
	})
}

// SeriousEventStatsNumEventsNotIn applies the NotIn predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsNotIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventStatsNumEvents), v...))
	})
}

// SeriousEventStatsNumEventsGT applies the GT predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsGT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsGTE applies the GTE predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsGTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsLT applies the LT predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsLT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsLTE applies the LTE predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsLTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsContains applies the Contains predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsContains(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsHasPrefix applies the HasPrefix predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsHasPrefix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsHasSuffix applies the HasSuffix predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsHasSuffix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsEqualFold applies the EqualFold predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsEqualFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumEventsContainsFold applies the ContainsFold predicate on the "serious_event_stats_num_events" field.
func SeriousEventStatsNumEventsContainsFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventStatsNumEvents), v))
	})
}

// SeriousEventStatsNumAffectedEQ applies the EQ predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedNEQ applies the NEQ predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedNEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedIn applies the In predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventStatsNumAffected), v...))
	})
}

// SeriousEventStatsNumAffectedNotIn applies the NotIn predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedNotIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventStatsNumAffected), v...))
	})
}

// SeriousEventStatsNumAffectedGT applies the GT predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedGT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedGTE applies the GTE predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedGTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedLT applies the LT predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedLT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedLTE applies the LTE predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedLTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedContains applies the Contains predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedContains(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedHasPrefix applies the HasPrefix predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedHasPrefix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedHasSuffix applies the HasSuffix predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedHasSuffix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedEqualFold applies the EqualFold predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedEqualFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAffectedContainsFold applies the ContainsFold predicate on the "serious_event_stats_num_affected" field.
func SeriousEventStatsNumAffectedContainsFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventStatsNumAffected), v))
	})
}

// SeriousEventStatsNumAtRiskEQ applies the EQ predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskNEQ applies the NEQ predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskNEQ(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskIn applies the In predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventStatsNumAtRisk), v...))
	})
}

// SeriousEventStatsNumAtRiskNotIn applies the NotIn predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskNotIn(vs ...string) predicate.SeriousEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventStatsNumAtRisk), v...))
	})
}

// SeriousEventStatsNumAtRiskGT applies the GT predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskGT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskGTE applies the GTE predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskGTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskLT applies the LT predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskLT(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskLTE applies the LTE predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskLTE(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskContains applies the Contains predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskContains(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskHasPrefix applies the HasPrefix predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskHasPrefix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskHasSuffix applies the HasSuffix predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskHasSuffix(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskEqualFold applies the EqualFold predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskEqualFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// SeriousEventStatsNumAtRiskContainsFold applies the ContainsFold predicate on the "serious_event_stats_num_at_risk" field.
func SeriousEventStatsNumAtRiskContainsFold(v string) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventStatsNumAtRisk), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.SeriousEvent) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
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
func And(predicates ...predicate.SeriousEventStats) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SeriousEventStats) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
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
func Not(p predicate.SeriousEventStats) predicate.SeriousEventStats {
	return predicate.SeriousEventStats(func(s *sql.Selector) {
		p(s.Not())
	})
}
