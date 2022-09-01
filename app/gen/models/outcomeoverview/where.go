// Code generated (@generated) by entc, DO NOT EDIT.

package outcomeoverview

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeOverviewID applies equality check predicate on the "outcome_overview_id" field. It's identical to OutcomeOverviewIDEQ.
func OutcomeOverviewID(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewTitle applies equality check predicate on the "outcome_overview_title" field. It's identical to OutcomeOverviewTitleEQ.
func OutcomeOverviewTitle(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewDescription applies equality check predicate on the "outcome_overview_description" field. It's identical to OutcomeOverviewDescriptionEQ.
func OutcomeOverviewDescription(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewParticipants applies equality check predicate on the "outcome_overview_participants" field. It's identical to OutcomeOverviewParticipantsEQ.
func OutcomeOverviewParticipants(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewTimeFrame applies equality check predicate on the "outcome_overview_time_frame" field. It's identical to OutcomeOverviewTimeFrameEQ.
func OutcomeOverviewTimeFrame(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewSerotype applies equality check predicate on the "outcome_overview_serotype" field. It's identical to OutcomeOverviewSerotypeEQ.
func OutcomeOverviewSerotype(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewAssay applies equality check predicate on the "outcome_overview_assay" field. It's identical to OutcomeOverviewAssayEQ.
func OutcomeOverviewAssay(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewDoseNumber applies equality check predicate on the "outcome_overview_dose_number" field. It's identical to OutcomeOverviewDoseNumberEQ.
func OutcomeOverviewDoseNumber(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewValue applies equality check predicate on the "outcome_overview_value" field. It's identical to OutcomeOverviewValueEQ.
func OutcomeOverviewValue(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewUpperLimit applies equality check predicate on the "outcome_overview_upper_limit" field. It's identical to OutcomeOverviewUpperLimitEQ.
func OutcomeOverviewUpperLimit(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewLowerLimit applies equality check predicate on the "outcome_overview_lower_limit" field. It's identical to OutcomeOverviewLowerLimitEQ.
func OutcomeOverviewLowerLimit(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewGroupID applies equality check predicate on the "outcome_overview_group_id" field. It's identical to OutcomeOverviewGroupIDEQ.
func OutcomeOverviewGroupID(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewRatio applies equality check predicate on the "outcome_overview_ratio" field. It's identical to OutcomeOverviewRatioEQ.
func OutcomeOverviewRatio(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewMeasureTitle applies equality check predicate on the "outcome_overview_measure_title" field. It's identical to OutcomeOverviewMeasureTitleEQ.
func OutcomeOverviewMeasureTitle(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewVaccine applies equality check predicate on the "outcome_overview_vaccine" field. It's identical to OutcomeOverviewVaccineEQ.
func OutcomeOverviewVaccine(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulation applies equality check predicate on the "outcome_overview_immunocompromised_population" field. It's identical to OutcomeOverviewImmunocompromisedPopulationEQ.
func OutcomeOverviewImmunocompromisedPopulation(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewManufacturer applies equality check predicate on the "outcome_overview_manufacturer" field. It's identical to OutcomeOverviewManufacturerEQ.
func OutcomeOverviewManufacturer(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewConfidenceInterval applies equality check predicate on the "outcome_overview_confidence_interval" field. It's identical to OutcomeOverviewConfidenceIntervalEQ.
func OutcomeOverviewConfidenceInterval(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewPercentResponders applies equality check predicate on the "outcome_overview_percent_responders" field. It's identical to OutcomeOverviewPercentRespondersEQ.
func OutcomeOverviewPercentResponders(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewTimeFrameWeeks applies equality check predicate on the "outcome_overview_time_frame_weeks" field. It's identical to OutcomeOverviewTimeFrameWeeksEQ.
func OutcomeOverviewTimeFrameWeeks(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewDoseDescription applies equality check predicate on the "outcome_overview_dose_description" field. It's identical to OutcomeOverviewDoseDescriptionEQ.
func OutcomeOverviewDoseDescription(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewSchedule applies equality check predicate on the "outcome_overview_schedule" field. It's identical to OutcomeOverviewScheduleEQ.
func OutcomeOverviewSchedule(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewUseCaseCode applies equality check predicate on the "outcome_overview_use_case_code" field. It's identical to OutcomeOverviewUseCaseCodeEQ.
func OutcomeOverviewUseCaseCode(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewIDEQ applies the EQ predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDNEQ applies the NEQ predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDIn applies the In predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewID), v...))
	})
}

// OutcomeOverviewIDNotIn applies the NotIn predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewID), v...))
	})
}

// OutcomeOverviewIDGT applies the GT predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDGTE applies the GTE predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDLT applies the LT predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDLTE applies the LTE predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDContains applies the Contains predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDHasPrefix applies the HasPrefix predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDHasSuffix applies the HasSuffix predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDEqualFold applies the EqualFold predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewIDContainsFold applies the ContainsFold predicate on the "outcome_overview_id" field.
func OutcomeOverviewIDContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewID), v))
	})
}

// OutcomeOverviewTitleEQ applies the EQ predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleNEQ applies the NEQ predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleIn applies the In predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewTitle), v...))
	})
}

// OutcomeOverviewTitleNotIn applies the NotIn predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewTitle), v...))
	})
}

// OutcomeOverviewTitleGT applies the GT predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleGTE applies the GTE predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleLT applies the LT predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleLTE applies the LTE predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleContains applies the Contains predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleHasPrefix applies the HasPrefix predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleHasSuffix applies the HasSuffix predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleEqualFold applies the EqualFold predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewTitleContainsFold applies the ContainsFold predicate on the "outcome_overview_title" field.
func OutcomeOverviewTitleContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewTitle), v))
	})
}

// OutcomeOverviewDescriptionEQ applies the EQ predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionNEQ applies the NEQ predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionIn applies the In predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewDescription), v...))
	})
}

// OutcomeOverviewDescriptionNotIn applies the NotIn predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewDescription), v...))
	})
}

// OutcomeOverviewDescriptionGT applies the GT predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionGTE applies the GTE predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionLT applies the LT predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionLTE applies the LTE predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionContains applies the Contains predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionEqualFold applies the EqualFold predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewDescriptionContainsFold applies the ContainsFold predicate on the "outcome_overview_description" field.
func OutcomeOverviewDescriptionContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewDescription), v))
	})
}

// OutcomeOverviewParticipantsEQ applies the EQ predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsNEQ applies the NEQ predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsIn applies the In predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewParticipants), v...))
	})
}

// OutcomeOverviewParticipantsNotIn applies the NotIn predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewParticipants), v...))
	})
}

// OutcomeOverviewParticipantsGT applies the GT predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsGTE applies the GTE predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsLT applies the LT predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsLTE applies the LTE predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsContains applies the Contains predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsHasPrefix applies the HasPrefix predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsHasSuffix applies the HasSuffix predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsEqualFold applies the EqualFold predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewParticipantsContainsFold applies the ContainsFold predicate on the "outcome_overview_participants" field.
func OutcomeOverviewParticipantsContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewParticipants), v))
	})
}

// OutcomeOverviewTimeFrameEQ applies the EQ predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameNEQ applies the NEQ predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameIn applies the In predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewTimeFrame), v...))
	})
}

// OutcomeOverviewTimeFrameNotIn applies the NotIn predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewTimeFrame), v...))
	})
}

// OutcomeOverviewTimeFrameGT applies the GT predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameGTE applies the GTE predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameLT applies the LT predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameLTE applies the LTE predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameContains applies the Contains predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameHasPrefix applies the HasPrefix predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameHasSuffix applies the HasSuffix predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameEqualFold applies the EqualFold predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewTimeFrameContainsFold applies the ContainsFold predicate on the "outcome_overview_time_frame" field.
func OutcomeOverviewTimeFrameContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewTimeFrame), v))
	})
}

// OutcomeOverviewSerotypeEQ applies the EQ predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeNEQ applies the NEQ predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeIn applies the In predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewSerotype), v...))
	})
}

// OutcomeOverviewSerotypeNotIn applies the NotIn predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewSerotype), v...))
	})
}

// OutcomeOverviewSerotypeGT applies the GT predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeGTE applies the GTE predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeLT applies the LT predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeLTE applies the LTE predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeContains applies the Contains predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeHasPrefix applies the HasPrefix predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeHasSuffix applies the HasSuffix predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeEqualFold applies the EqualFold predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewSerotypeContainsFold applies the ContainsFold predicate on the "outcome_overview_serotype" field.
func OutcomeOverviewSerotypeContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewSerotype), v))
	})
}

// OutcomeOverviewAssayEQ applies the EQ predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayNEQ applies the NEQ predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayIn applies the In predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewAssay), v...))
	})
}

// OutcomeOverviewAssayNotIn applies the NotIn predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewAssay), v...))
	})
}

// OutcomeOverviewAssayGT applies the GT predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayGTE applies the GTE predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayLT applies the LT predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayLTE applies the LTE predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayContains applies the Contains predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayHasPrefix applies the HasPrefix predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayHasSuffix applies the HasSuffix predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayEqualFold applies the EqualFold predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewAssayContainsFold applies the ContainsFold predicate on the "outcome_overview_assay" field.
func OutcomeOverviewAssayContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewAssay), v))
	})
}

// OutcomeOverviewDoseNumberEQ applies the EQ predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberEQ(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewDoseNumberNEQ applies the NEQ predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberNEQ(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewDoseNumberIn applies the In predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberIn(vs ...int64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewDoseNumber), v...))
	})
}

// OutcomeOverviewDoseNumberNotIn applies the NotIn predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberNotIn(vs ...int64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewDoseNumber), v...))
	})
}

// OutcomeOverviewDoseNumberGT applies the GT predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberGT(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewDoseNumberGTE applies the GTE predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberGTE(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewDoseNumberLT applies the LT predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberLT(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewDoseNumberLTE applies the LTE predicate on the "outcome_overview_dose_number" field.
func OutcomeOverviewDoseNumberLTE(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewDoseNumber), v))
	})
}

// OutcomeOverviewValueEQ applies the EQ predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueEQ(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewValueNEQ applies the NEQ predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueNEQ(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewValueIn applies the In predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueIn(vs ...float64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewValue), v...))
	})
}

// OutcomeOverviewValueNotIn applies the NotIn predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueNotIn(vs ...float64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewValue), v...))
	})
}

// OutcomeOverviewValueGT applies the GT predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueGT(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewValueGTE applies the GTE predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueGTE(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewValueLT applies the LT predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueLT(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewValueLTE applies the LTE predicate on the "outcome_overview_value" field.
func OutcomeOverviewValueLTE(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewValue), v))
	})
}

// OutcomeOverviewUpperLimitEQ applies the EQ predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitNEQ applies the NEQ predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitIn applies the In predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewUpperLimit), v...))
	})
}

// OutcomeOverviewUpperLimitNotIn applies the NotIn predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewUpperLimit), v...))
	})
}

// OutcomeOverviewUpperLimitGT applies the GT predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitGTE applies the GTE predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitLT applies the LT predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitLTE applies the LTE predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitContains applies the Contains predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitHasPrefix applies the HasPrefix predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitHasSuffix applies the HasSuffix predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitEqualFold applies the EqualFold predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewUpperLimitContainsFold applies the ContainsFold predicate on the "outcome_overview_upper_limit" field.
func OutcomeOverviewUpperLimitContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewUpperLimit), v))
	})
}

// OutcomeOverviewLowerLimitEQ applies the EQ predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitNEQ applies the NEQ predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitIn applies the In predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewLowerLimit), v...))
	})
}

// OutcomeOverviewLowerLimitNotIn applies the NotIn predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewLowerLimit), v...))
	})
}

// OutcomeOverviewLowerLimitGT applies the GT predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitGTE applies the GTE predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitLT applies the LT predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitLTE applies the LTE predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitContains applies the Contains predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitHasPrefix applies the HasPrefix predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitHasSuffix applies the HasSuffix predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitEqualFold applies the EqualFold predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewLowerLimitContainsFold applies the ContainsFold predicate on the "outcome_overview_lower_limit" field.
func OutcomeOverviewLowerLimitContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewLowerLimit), v))
	})
}

// OutcomeOverviewGroupIDEQ applies the EQ predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDNEQ applies the NEQ predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDIn applies the In predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewGroupID), v...))
	})
}

// OutcomeOverviewGroupIDNotIn applies the NotIn predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewGroupID), v...))
	})
}

// OutcomeOverviewGroupIDGT applies the GT predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDGTE applies the GTE predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDLT applies the LT predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDLTE applies the LTE predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDContains applies the Contains predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDEqualFold applies the EqualFold predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewGroupIDContainsFold applies the ContainsFold predicate on the "outcome_overview_group_id" field.
func OutcomeOverviewGroupIDContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewGroupID), v))
	})
}

// OutcomeOverviewRatioEQ applies the EQ predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioNEQ applies the NEQ predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioIn applies the In predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewRatio), v...))
	})
}

// OutcomeOverviewRatioNotIn applies the NotIn predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewRatio), v...))
	})
}

// OutcomeOverviewRatioGT applies the GT predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioGTE applies the GTE predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioLT applies the LT predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioLTE applies the LTE predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioContains applies the Contains predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioHasPrefix applies the HasPrefix predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioHasSuffix applies the HasSuffix predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioEqualFold applies the EqualFold predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewRatioContainsFold applies the ContainsFold predicate on the "outcome_overview_ratio" field.
func OutcomeOverviewRatioContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewRatio), v))
	})
}

// OutcomeOverviewMeasureTitleEQ applies the EQ predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleNEQ applies the NEQ predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleIn applies the In predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewMeasureTitle), v...))
	})
}

// OutcomeOverviewMeasureTitleNotIn applies the NotIn predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewMeasureTitle), v...))
	})
}

// OutcomeOverviewMeasureTitleGT applies the GT predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleGTE applies the GTE predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleLT applies the LT predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleLTE applies the LTE predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleContains applies the Contains predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleHasPrefix applies the HasPrefix predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleHasSuffix applies the HasSuffix predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleEqualFold applies the EqualFold predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewMeasureTitleContainsFold applies the ContainsFold predicate on the "outcome_overview_measure_title" field.
func OutcomeOverviewMeasureTitleContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewMeasureTitle), v))
	})
}

// OutcomeOverviewVaccineEQ applies the EQ predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineNEQ applies the NEQ predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineIn applies the In predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewVaccine), v...))
	})
}

// OutcomeOverviewVaccineNotIn applies the NotIn predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewVaccine), v...))
	})
}

// OutcomeOverviewVaccineGT applies the GT predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineGTE applies the GTE predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineLT applies the LT predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineLTE applies the LTE predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineContains applies the Contains predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineHasPrefix applies the HasPrefix predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineHasSuffix applies the HasSuffix predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineEqualFold applies the EqualFold predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewVaccineContainsFold applies the ContainsFold predicate on the "outcome_overview_vaccine" field.
func OutcomeOverviewVaccineContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewVaccine), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationEQ applies the EQ predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationNEQ applies the NEQ predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationIn applies the In predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v...))
	})
}

// OutcomeOverviewImmunocompromisedPopulationNotIn applies the NotIn predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v...))
	})
}

// OutcomeOverviewImmunocompromisedPopulationGT applies the GT predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationGTE applies the GTE predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationLT applies the LT predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationLTE applies the LTE predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationContains applies the Contains predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationHasPrefix applies the HasPrefix predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationHasSuffix applies the HasSuffix predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationEqualFold applies the EqualFold predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewImmunocompromisedPopulationContainsFold applies the ContainsFold predicate on the "outcome_overview_immunocompromised_population" field.
func OutcomeOverviewImmunocompromisedPopulationContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewImmunocompromisedPopulation), v))
	})
}

// OutcomeOverviewManufacturerEQ applies the EQ predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerNEQ applies the NEQ predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerIn applies the In predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewManufacturer), v...))
	})
}

// OutcomeOverviewManufacturerNotIn applies the NotIn predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewManufacturer), v...))
	})
}

// OutcomeOverviewManufacturerGT applies the GT predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerGTE applies the GTE predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerLT applies the LT predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerLTE applies the LTE predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerContains applies the Contains predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerHasPrefix applies the HasPrefix predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerHasSuffix applies the HasSuffix predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerEqualFold applies the EqualFold predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewManufacturerContainsFold applies the ContainsFold predicate on the "outcome_overview_manufacturer" field.
func OutcomeOverviewManufacturerContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewManufacturer), v))
	})
}

// OutcomeOverviewConfidenceIntervalEQ applies the EQ predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalNEQ applies the NEQ predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalIn applies the In predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewConfidenceInterval), v...))
	})
}

// OutcomeOverviewConfidenceIntervalNotIn applies the NotIn predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewConfidenceInterval), v...))
	})
}

// OutcomeOverviewConfidenceIntervalGT applies the GT predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalGTE applies the GTE predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalLT applies the LT predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalLTE applies the LTE predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalContains applies the Contains predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalHasPrefix applies the HasPrefix predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalHasSuffix applies the HasSuffix predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalEqualFold applies the EqualFold predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewConfidenceIntervalContainsFold applies the ContainsFold predicate on the "outcome_overview_confidence_interval" field.
func OutcomeOverviewConfidenceIntervalContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewConfidenceInterval), v))
	})
}

// OutcomeOverviewPercentRespondersEQ applies the EQ predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersEQ(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewPercentRespondersNEQ applies the NEQ predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersNEQ(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewPercentRespondersIn applies the In predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersIn(vs ...float64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewPercentResponders), v...))
	})
}

// OutcomeOverviewPercentRespondersNotIn applies the NotIn predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersNotIn(vs ...float64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewPercentResponders), v...))
	})
}

// OutcomeOverviewPercentRespondersGT applies the GT predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersGT(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewPercentRespondersGTE applies the GTE predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersGTE(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewPercentRespondersLT applies the LT predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersLT(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewPercentRespondersLTE applies the LTE predicate on the "outcome_overview_percent_responders" field.
func OutcomeOverviewPercentRespondersLTE(v float64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewPercentResponders), v))
	})
}

// OutcomeOverviewTimeFrameWeeksEQ applies the EQ predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksEQ(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewTimeFrameWeeksNEQ applies the NEQ predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksNEQ(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewTimeFrameWeeksIn applies the In predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksIn(vs ...int64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewTimeFrameWeeks), v...))
	})
}

// OutcomeOverviewTimeFrameWeeksNotIn applies the NotIn predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksNotIn(vs ...int64) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewTimeFrameWeeks), v...))
	})
}

// OutcomeOverviewTimeFrameWeeksGT applies the GT predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksGT(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewTimeFrameWeeksGTE applies the GTE predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksGTE(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewTimeFrameWeeksLT applies the LT predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksLT(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewTimeFrameWeeksLTE applies the LTE predicate on the "outcome_overview_time_frame_weeks" field.
func OutcomeOverviewTimeFrameWeeksLTE(v int64) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewTimeFrameWeeks), v))
	})
}

// OutcomeOverviewDoseDescriptionEQ applies the EQ predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionNEQ applies the NEQ predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionIn applies the In predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewDoseDescription), v...))
	})
}

// OutcomeOverviewDoseDescriptionNotIn applies the NotIn predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewDoseDescription), v...))
	})
}

// OutcomeOverviewDoseDescriptionGT applies the GT predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionGTE applies the GTE predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionLT applies the LT predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionLTE applies the LTE predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionContains applies the Contains predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionEqualFold applies the EqualFold predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewDoseDescriptionContainsFold applies the ContainsFold predicate on the "outcome_overview_dose_description" field.
func OutcomeOverviewDoseDescriptionContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewDoseDescription), v))
	})
}

// OutcomeOverviewScheduleEQ applies the EQ predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleNEQ applies the NEQ predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleIn applies the In predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewSchedule), v...))
	})
}

// OutcomeOverviewScheduleNotIn applies the NotIn predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewSchedule), v...))
	})
}

// OutcomeOverviewScheduleGT applies the GT predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleGTE applies the GTE predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleLT applies the LT predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleLTE applies the LTE predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleContains applies the Contains predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleHasPrefix applies the HasPrefix predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleHasSuffix applies the HasSuffix predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleEqualFold applies the EqualFold predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewScheduleContainsFold applies the ContainsFold predicate on the "outcome_overview_schedule" field.
func OutcomeOverviewScheduleContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewSchedule), v))
	})
}

// OutcomeOverviewUseCaseCodeEQ applies the EQ predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeNEQ applies the NEQ predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeNEQ(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeIn applies the In predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeOverviewUseCaseCode), v...))
	})
}

// OutcomeOverviewUseCaseCodeNotIn applies the NotIn predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeNotIn(vs ...string) predicate.OutcomeOverview {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeOverviewUseCaseCode), v...))
	})
}

// OutcomeOverviewUseCaseCodeGT applies the GT predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeGT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeGTE applies the GTE predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeGTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeLT applies the LT predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeLT(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeLTE applies the LTE predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeLTE(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeContains applies the Contains predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeContains(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeHasPrefix applies the HasPrefix predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeHasPrefix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeHasSuffix applies the HasSuffix predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeHasSuffix(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeEqualFold applies the EqualFold predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeEqualFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// OutcomeOverviewUseCaseCodeContainsFold applies the ContainsFold predicate on the "outcome_overview_use_case_code" field.
func OutcomeOverviewUseCaseCodeContainsFold(v string) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeOverviewUseCaseCode), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasure) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
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
func And(predicates ...predicate.OutcomeOverview) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeOverview) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeOverview) predicate.OutcomeOverview {
	return predicate.OutcomeOverview(func(s *sql.Selector) {
		p(s.Not())
	})
}
