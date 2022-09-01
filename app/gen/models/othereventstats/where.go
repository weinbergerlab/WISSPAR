// Code generated (@generated) by entc, DO NOT EDIT.

package othereventstats

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OtherEventStatsGroupID applies equality check predicate on the "other_event_stats_group_id" field. It's identical to OtherEventStatsGroupIDEQ.
func OtherEventStatsGroupID(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsNumEvents applies equality check predicate on the "other_event_stats_num_events" field. It's identical to OtherEventStatsNumEventsEQ.
func OtherEventStatsNumEvents(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumAffected applies equality check predicate on the "other_event_stats_num_affected" field. It's identical to OtherEventStatsNumAffectedEQ.
func OtherEventStatsNumAffected(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAtRisk applies equality check predicate on the "other_event_stats_num_at_risk" field. It's identical to OtherEventStatsNumAtRiskEQ.
func OtherEventStatsNumAtRisk(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsGroupIDEQ applies the EQ predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDNEQ applies the NEQ predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDNEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDIn applies the In predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventStatsGroupID), v...))
	})
}

// OtherEventStatsGroupIDNotIn applies the NotIn predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDNotIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventStatsGroupID), v...))
	})
}

// OtherEventStatsGroupIDGT applies the GT predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDGT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDGTE applies the GTE predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDGTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDLT applies the LT predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDLT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDLTE applies the LTE predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDLTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDContains applies the Contains predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDContains(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDHasPrefix applies the HasPrefix predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDHasPrefix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDHasSuffix applies the HasSuffix predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDHasSuffix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDEqualFold applies the EqualFold predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDEqualFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsGroupIDContainsFold applies the ContainsFold predicate on the "other_event_stats_group_id" field.
func OtherEventStatsGroupIDContainsFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventStatsGroupID), v))
	})
}

// OtherEventStatsNumEventsEQ applies the EQ predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsNEQ applies the NEQ predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsNEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsIn applies the In predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventStatsNumEvents), v...))
	})
}

// OtherEventStatsNumEventsNotIn applies the NotIn predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsNotIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventStatsNumEvents), v...))
	})
}

// OtherEventStatsNumEventsGT applies the GT predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsGT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsGTE applies the GTE predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsGTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsLT applies the LT predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsLT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsLTE applies the LTE predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsLTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsContains applies the Contains predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsContains(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsHasPrefix applies the HasPrefix predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsHasPrefix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsHasSuffix applies the HasSuffix predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsHasSuffix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsEqualFold applies the EqualFold predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsEqualFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumEventsContainsFold applies the ContainsFold predicate on the "other_event_stats_num_events" field.
func OtherEventStatsNumEventsContainsFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventStatsNumEvents), v))
	})
}

// OtherEventStatsNumAffectedEQ applies the EQ predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedNEQ applies the NEQ predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedNEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedIn applies the In predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventStatsNumAffected), v...))
	})
}

// OtherEventStatsNumAffectedNotIn applies the NotIn predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedNotIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventStatsNumAffected), v...))
	})
}

// OtherEventStatsNumAffectedGT applies the GT predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedGT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedGTE applies the GTE predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedGTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedLT applies the LT predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedLT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedLTE applies the LTE predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedLTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedContains applies the Contains predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedContains(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedHasPrefix applies the HasPrefix predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedHasPrefix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedHasSuffix applies the HasSuffix predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedHasSuffix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedEqualFold applies the EqualFold predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedEqualFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAffectedContainsFold applies the ContainsFold predicate on the "other_event_stats_num_affected" field.
func OtherEventStatsNumAffectedContainsFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventStatsNumAffected), v))
	})
}

// OtherEventStatsNumAtRiskEQ applies the EQ predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskNEQ applies the NEQ predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskNEQ(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskIn applies the In predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventStatsNumAtRisk), v...))
	})
}

// OtherEventStatsNumAtRiskNotIn applies the NotIn predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskNotIn(vs ...string) predicate.OtherEventStats {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEventStats(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventStatsNumAtRisk), v...))
	})
}

// OtherEventStatsNumAtRiskGT applies the GT predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskGT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskGTE applies the GTE predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskGTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskLT applies the LT predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskLT(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskLTE applies the LTE predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskLTE(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskContains applies the Contains predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskContains(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskHasPrefix applies the HasPrefix predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskHasPrefix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskHasSuffix applies the HasSuffix predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskHasSuffix(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskEqualFold applies the EqualFold predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskEqualFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// OtherEventStatsNumAtRiskContainsFold applies the ContainsFold predicate on the "other_event_stats_num_at_risk" field.
func OtherEventStatsNumAtRiskContainsFold(v string) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventStatsNumAtRisk), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OtherEvent) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
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
func And(predicates ...predicate.OtherEventStats) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OtherEventStats) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
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
func Not(p predicate.OtherEventStats) predicate.OtherEventStats {
	return predicate.OtherEventStats(func(s *sql.Selector) {
		p(s.Not())
	})
}
