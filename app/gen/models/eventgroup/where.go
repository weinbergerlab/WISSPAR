// Code generated (@generated) by entc, DO NOT EDIT.

package eventgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EventGroupID applies equality check predicate on the "event_group_id" field. It's identical to EventGroupIDEQ.
func EventGroupID(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupID), v))
	})
}

// EventGroupTitle applies equality check predicate on the "event_group_title" field. It's identical to EventGroupTitleEQ.
func EventGroupTitle(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupDescription applies equality check predicate on the "event_group_description" field. It's identical to EventGroupDescriptionEQ.
func EventGroupDescription(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDeathsNumAffected applies equality check predicate on the "event_group_deaths_num_affected" field. It's identical to EventGroupDeathsNumAffectedEQ.
func EventGroupDeathsNumAffected(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAtRisk applies equality check predicate on the "event_group_deaths_num_at_risk" field. It's identical to EventGroupDeathsNumAtRiskEQ.
func EventGroupDeathsNumAtRisk(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupSeriousNumAffected applies equality check predicate on the "event_group_serious_num_affected" field. It's identical to EventGroupSeriousNumAffectedEQ.
func EventGroupSeriousNumAffected(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAtRisk applies equality check predicate on the "event_group_serious_num_at_risk" field. It's identical to EventGroupSeriousNumAtRiskEQ.
func EventGroupSeriousNumAtRisk(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupOtherNumAffected applies equality check predicate on the "event_group_other_num_affected" field. It's identical to EventGroupOtherNumAffectedEQ.
func EventGroupOtherNumAffected(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAtRisk applies equality check predicate on the "event_group_other_num_at_risk" field. It's identical to EventGroupOtherNumAtRiskEQ.
func EventGroupOtherNumAtRisk(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupIDEQ applies the EQ predicate on the "event_group_id" field.
func EventGroupIDEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDNEQ applies the NEQ predicate on the "event_group_id" field.
func EventGroupIDNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDIn applies the In predicate on the "event_group_id" field.
func EventGroupIDIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupID), v...))
	})
}

// EventGroupIDNotIn applies the NotIn predicate on the "event_group_id" field.
func EventGroupIDNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupID), v...))
	})
}

// EventGroupIDGT applies the GT predicate on the "event_group_id" field.
func EventGroupIDGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDGTE applies the GTE predicate on the "event_group_id" field.
func EventGroupIDGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDLT applies the LT predicate on the "event_group_id" field.
func EventGroupIDLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDLTE applies the LTE predicate on the "event_group_id" field.
func EventGroupIDLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDContains applies the Contains predicate on the "event_group_id" field.
func EventGroupIDContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDHasPrefix applies the HasPrefix predicate on the "event_group_id" field.
func EventGroupIDHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDHasSuffix applies the HasSuffix predicate on the "event_group_id" field.
func EventGroupIDHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDEqualFold applies the EqualFold predicate on the "event_group_id" field.
func EventGroupIDEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupID), v))
	})
}

// EventGroupIDContainsFold applies the ContainsFold predicate on the "event_group_id" field.
func EventGroupIDContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupID), v))
	})
}

// EventGroupTitleEQ applies the EQ predicate on the "event_group_title" field.
func EventGroupTitleEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleNEQ applies the NEQ predicate on the "event_group_title" field.
func EventGroupTitleNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleIn applies the In predicate on the "event_group_title" field.
func EventGroupTitleIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupTitle), v...))
	})
}

// EventGroupTitleNotIn applies the NotIn predicate on the "event_group_title" field.
func EventGroupTitleNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupTitle), v...))
	})
}

// EventGroupTitleGT applies the GT predicate on the "event_group_title" field.
func EventGroupTitleGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleGTE applies the GTE predicate on the "event_group_title" field.
func EventGroupTitleGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleLT applies the LT predicate on the "event_group_title" field.
func EventGroupTitleLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleLTE applies the LTE predicate on the "event_group_title" field.
func EventGroupTitleLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleContains applies the Contains predicate on the "event_group_title" field.
func EventGroupTitleContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleHasPrefix applies the HasPrefix predicate on the "event_group_title" field.
func EventGroupTitleHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleHasSuffix applies the HasSuffix predicate on the "event_group_title" field.
func EventGroupTitleHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleEqualFold applies the EqualFold predicate on the "event_group_title" field.
func EventGroupTitleEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupTitleContainsFold applies the ContainsFold predicate on the "event_group_title" field.
func EventGroupTitleContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupTitle), v))
	})
}

// EventGroupDescriptionEQ applies the EQ predicate on the "event_group_description" field.
func EventGroupDescriptionEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionNEQ applies the NEQ predicate on the "event_group_description" field.
func EventGroupDescriptionNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionIn applies the In predicate on the "event_group_description" field.
func EventGroupDescriptionIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupDescription), v...))
	})
}

// EventGroupDescriptionNotIn applies the NotIn predicate on the "event_group_description" field.
func EventGroupDescriptionNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupDescription), v...))
	})
}

// EventGroupDescriptionGT applies the GT predicate on the "event_group_description" field.
func EventGroupDescriptionGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionGTE applies the GTE predicate on the "event_group_description" field.
func EventGroupDescriptionGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionLT applies the LT predicate on the "event_group_description" field.
func EventGroupDescriptionLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionLTE applies the LTE predicate on the "event_group_description" field.
func EventGroupDescriptionLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionContains applies the Contains predicate on the "event_group_description" field.
func EventGroupDescriptionContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionHasPrefix applies the HasPrefix predicate on the "event_group_description" field.
func EventGroupDescriptionHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionHasSuffix applies the HasSuffix predicate on the "event_group_description" field.
func EventGroupDescriptionHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionEqualFold applies the EqualFold predicate on the "event_group_description" field.
func EventGroupDescriptionEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDescriptionContainsFold applies the ContainsFold predicate on the "event_group_description" field.
func EventGroupDescriptionContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupDescription), v))
	})
}

// EventGroupDeathsNumAffectedEQ applies the EQ predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedNEQ applies the NEQ predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedIn applies the In predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupDeathsNumAffected), v...))
	})
}

// EventGroupDeathsNumAffectedNotIn applies the NotIn predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupDeathsNumAffected), v...))
	})
}

// EventGroupDeathsNumAffectedGT applies the GT predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedGTE applies the GTE predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedLT applies the LT predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedLTE applies the LTE predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedContains applies the Contains predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedHasPrefix applies the HasPrefix predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedHasSuffix applies the HasSuffix predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedEqualFold applies the EqualFold predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAffectedContainsFold applies the ContainsFold predicate on the "event_group_deaths_num_affected" field.
func EventGroupDeathsNumAffectedContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupDeathsNumAffected), v))
	})
}

// EventGroupDeathsNumAtRiskEQ applies the EQ predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskNEQ applies the NEQ predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskIn applies the In predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupDeathsNumAtRisk), v...))
	})
}

// EventGroupDeathsNumAtRiskNotIn applies the NotIn predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupDeathsNumAtRisk), v...))
	})
}

// EventGroupDeathsNumAtRiskGT applies the GT predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskGTE applies the GTE predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskLT applies the LT predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskLTE applies the LTE predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskContains applies the Contains predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskHasPrefix applies the HasPrefix predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskHasSuffix applies the HasSuffix predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskEqualFold applies the EqualFold predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupDeathsNumAtRiskContainsFold applies the ContainsFold predicate on the "event_group_deaths_num_at_risk" field.
func EventGroupDeathsNumAtRiskContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupDeathsNumAtRisk), v))
	})
}

// EventGroupSeriousNumAffectedEQ applies the EQ predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedNEQ applies the NEQ predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedIn applies the In predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupSeriousNumAffected), v...))
	})
}

// EventGroupSeriousNumAffectedNotIn applies the NotIn predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupSeriousNumAffected), v...))
	})
}

// EventGroupSeriousNumAffectedGT applies the GT predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedGTE applies the GTE predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedLT applies the LT predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedLTE applies the LTE predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedContains applies the Contains predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedHasPrefix applies the HasPrefix predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedHasSuffix applies the HasSuffix predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedEqualFold applies the EqualFold predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAffectedContainsFold applies the ContainsFold predicate on the "event_group_serious_num_affected" field.
func EventGroupSeriousNumAffectedContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupSeriousNumAffected), v))
	})
}

// EventGroupSeriousNumAtRiskEQ applies the EQ predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskNEQ applies the NEQ predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskIn applies the In predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupSeriousNumAtRisk), v...))
	})
}

// EventGroupSeriousNumAtRiskNotIn applies the NotIn predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupSeriousNumAtRisk), v...))
	})
}

// EventGroupSeriousNumAtRiskGT applies the GT predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskGTE applies the GTE predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskLT applies the LT predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskLTE applies the LTE predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskContains applies the Contains predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskHasPrefix applies the HasPrefix predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskHasSuffix applies the HasSuffix predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskEqualFold applies the EqualFold predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupSeriousNumAtRiskContainsFold applies the ContainsFold predicate on the "event_group_serious_num_at_risk" field.
func EventGroupSeriousNumAtRiskContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupSeriousNumAtRisk), v))
	})
}

// EventGroupOtherNumAffectedEQ applies the EQ predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedNEQ applies the NEQ predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedIn applies the In predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupOtherNumAffected), v...))
	})
}

// EventGroupOtherNumAffectedNotIn applies the NotIn predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupOtherNumAffected), v...))
	})
}

// EventGroupOtherNumAffectedGT applies the GT predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedGTE applies the GTE predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedLT applies the LT predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedLTE applies the LTE predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedContains applies the Contains predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedHasPrefix applies the HasPrefix predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedHasSuffix applies the HasSuffix predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedEqualFold applies the EqualFold predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAffectedContainsFold applies the ContainsFold predicate on the "event_group_other_num_affected" field.
func EventGroupOtherNumAffectedContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupOtherNumAffected), v))
	})
}

// EventGroupOtherNumAtRiskEQ applies the EQ predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskNEQ applies the NEQ predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskNEQ(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskIn applies the In predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventGroupOtherNumAtRisk), v...))
	})
}

// EventGroupOtherNumAtRiskNotIn applies the NotIn predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskNotIn(vs ...string) predicate.EventGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventGroupOtherNumAtRisk), v...))
	})
}

// EventGroupOtherNumAtRiskGT applies the GT predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskGT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskGTE applies the GTE predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskGTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskLT applies the LT predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskLT(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskLTE applies the LTE predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskLTE(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskContains applies the Contains predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskContains(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskHasPrefix applies the HasPrefix predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskHasPrefix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskHasSuffix applies the HasSuffix predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskHasSuffix(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskEqualFold applies the EqualFold predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskEqualFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// EventGroupOtherNumAtRiskContainsFold applies the ContainsFold predicate on the "event_group_other_num_at_risk" field.
func EventGroupOtherNumAtRiskContainsFold(v string) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventGroupOtherNumAtRisk), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.AdverseEventsModule) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
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
func And(predicates ...predicate.EventGroup) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EventGroup) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
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
func Not(p predicate.EventGroup) predicate.EventGroup {
	return predicate.EventGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
