// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasure

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeMeasureType applies equality check predicate on the "outcome_measure_type" field. It's identical to OutcomeMeasureTypeEQ.
func OutcomeMeasureType(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTitle applies equality check predicate on the "outcome_measure_title" field. It's identical to OutcomeMeasureTitleEQ.
func OutcomeMeasureTitle(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureDescription applies equality check predicate on the "outcome_measure_description" field. It's identical to OutcomeMeasureDescriptionEQ.
func OutcomeMeasureDescription(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasurePopulationDescription applies equality check predicate on the "outcome_measure_population_description" field. It's identical to OutcomeMeasurePopulationDescriptionEQ.
func OutcomeMeasurePopulationDescription(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasureReportingStatus applies equality check predicate on the "outcome_measure_reporting_status" field. It's identical to OutcomeMeasureReportingStatusEQ.
func OutcomeMeasureReportingStatus(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureAnticipatedPostingDate applies equality check predicate on the "outcome_measure_anticipated_posting_date" field. It's identical to OutcomeMeasureAnticipatedPostingDateEQ.
func OutcomeMeasureAnticipatedPostingDate(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureParamType applies equality check predicate on the "outcome_measure_param_type" field. It's identical to OutcomeMeasureParamTypeEQ.
func OutcomeMeasureParamType(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureDispersionType applies equality check predicate on the "outcome_measure_dispersion_type" field. It's identical to OutcomeMeasureDispersionTypeEQ.
func OutcomeMeasureDispersionType(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureUnitOfMeasure applies equality check predicate on the "outcome_measure_unit_of_measure" field. It's identical to OutcomeMeasureUnitOfMeasureEQ.
func OutcomeMeasureUnitOfMeasure(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureCalculatePct applies equality check predicate on the "outcome_measure_calculate_pct" field. It's identical to OutcomeMeasureCalculatePctEQ.
func OutcomeMeasureCalculatePct(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureTimeFrame applies equality check predicate on the "outcome_measure_time_frame" field. It's identical to OutcomeMeasureTimeFrameEQ.
func OutcomeMeasureTimeFrame(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzed applies equality check predicate on the "outcome_measure_type_units_analyzed" field. It's identical to OutcomeMeasureTypeUnitsAnalyzedEQ.
func OutcomeMeasureTypeUnitsAnalyzed(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureDenomUnitsSelected applies equality check predicate on the "outcome_measure_denom_units_selected" field. It's identical to OutcomeMeasureDenomUnitsSelectedEQ.
func OutcomeMeasureDenomUnitsSelected(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureTypeEQ applies the EQ predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeNEQ applies the NEQ predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeIn applies the In predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureType), v...))
	})
}

// OutcomeMeasureTypeNotIn applies the NotIn predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureType), v...))
	})
}

// OutcomeMeasureTypeGT applies the GT predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeGTE applies the GTE predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeLT applies the LT predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeLTE applies the LTE predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeContains applies the Contains predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeHasPrefix applies the HasPrefix predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeHasSuffix applies the HasSuffix predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeEqualFold applies the EqualFold predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTypeContainsFold applies the ContainsFold predicate on the "outcome_measure_type" field.
func OutcomeMeasureTypeContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureType), v))
	})
}

// OutcomeMeasureTitleEQ applies the EQ predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleNEQ applies the NEQ predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleIn applies the In predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureTitle), v...))
	})
}

// OutcomeMeasureTitleNotIn applies the NotIn predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureTitle), v...))
	})
}

// OutcomeMeasureTitleGT applies the GT predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleGTE applies the GTE predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleLT applies the LT predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleLTE applies the LTE predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleContains applies the Contains predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleHasPrefix applies the HasPrefix predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleHasSuffix applies the HasSuffix predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleEqualFold applies the EqualFold predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureTitleContainsFold applies the ContainsFold predicate on the "outcome_measure_title" field.
func OutcomeMeasureTitleContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureTitle), v))
	})
}

// OutcomeMeasureDescriptionEQ applies the EQ predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionNEQ applies the NEQ predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionIn applies the In predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureDescription), v...))
	})
}

// OutcomeMeasureDescriptionNotIn applies the NotIn predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureDescription), v...))
	})
}

// OutcomeMeasureDescriptionGT applies the GT predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionGTE applies the GTE predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionLT applies the LT predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionLTE applies the LTE predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionContains applies the Contains predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionEqualFold applies the EqualFold predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasureDescriptionContainsFold applies the ContainsFold predicate on the "outcome_measure_description" field.
func OutcomeMeasureDescriptionContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionEQ applies the EQ predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionNEQ applies the NEQ predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionIn applies the In predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurePopulationDescription), v...))
	})
}

// OutcomeMeasurePopulationDescriptionNotIn applies the NotIn predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurePopulationDescription), v...))
	})
}

// OutcomeMeasurePopulationDescriptionGT applies the GT predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionGTE applies the GTE predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionLT applies the LT predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionLTE applies the LTE predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionContains applies the Contains predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionHasPrefix applies the HasPrefix predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionHasSuffix applies the HasSuffix predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionEqualFold applies the EqualFold predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasurePopulationDescriptionContainsFold applies the ContainsFold predicate on the "outcome_measure_population_description" field.
func OutcomeMeasurePopulationDescriptionContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurePopulationDescription), v))
	})
}

// OutcomeMeasureReportingStatusEQ applies the EQ predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusNEQ applies the NEQ predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusIn applies the In predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureReportingStatus), v...))
	})
}

// OutcomeMeasureReportingStatusNotIn applies the NotIn predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureReportingStatus), v...))
	})
}

// OutcomeMeasureReportingStatusGT applies the GT predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusGTE applies the GTE predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusLT applies the LT predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusLTE applies the LTE predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusContains applies the Contains predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusHasPrefix applies the HasPrefix predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusHasSuffix applies the HasSuffix predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusEqualFold applies the EqualFold predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureReportingStatusContainsFold applies the ContainsFold predicate on the "outcome_measure_reporting_status" field.
func OutcomeMeasureReportingStatusContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureReportingStatus), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateEQ applies the EQ predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateNEQ applies the NEQ predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateIn applies the In predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v...))
	})
}

// OutcomeMeasureAnticipatedPostingDateNotIn applies the NotIn predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v...))
	})
}

// OutcomeMeasureAnticipatedPostingDateGT applies the GT predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateGTE applies the GTE predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateLT applies the LT predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateLTE applies the LTE predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateContains applies the Contains predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateHasPrefix applies the HasPrefix predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateHasSuffix applies the HasSuffix predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateEqualFold applies the EqualFold predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureAnticipatedPostingDateContainsFold applies the ContainsFold predicate on the "outcome_measure_anticipated_posting_date" field.
func OutcomeMeasureAnticipatedPostingDateContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureAnticipatedPostingDate), v))
	})
}

// OutcomeMeasureParamTypeEQ applies the EQ predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeNEQ applies the NEQ predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeIn applies the In predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureParamType), v...))
	})
}

// OutcomeMeasureParamTypeNotIn applies the NotIn predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureParamType), v...))
	})
}

// OutcomeMeasureParamTypeGT applies the GT predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeGTE applies the GTE predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeLT applies the LT predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeLTE applies the LTE predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeContains applies the Contains predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeHasPrefix applies the HasPrefix predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeHasSuffix applies the HasSuffix predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeEqualFold applies the EqualFold predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureParamTypeContainsFold applies the ContainsFold predicate on the "outcome_measure_param_type" field.
func OutcomeMeasureParamTypeContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureParamType), v))
	})
}

// OutcomeMeasureDispersionTypeEQ applies the EQ predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeNEQ applies the NEQ predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeIn applies the In predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureDispersionType), v...))
	})
}

// OutcomeMeasureDispersionTypeNotIn applies the NotIn predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureDispersionType), v...))
	})
}

// OutcomeMeasureDispersionTypeGT applies the GT predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeGTE applies the GTE predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeLT applies the LT predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeLTE applies the LTE predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeContains applies the Contains predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeHasPrefix applies the HasPrefix predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeHasSuffix applies the HasSuffix predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeEqualFold applies the EqualFold predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureDispersionTypeContainsFold applies the ContainsFold predicate on the "outcome_measure_dispersion_type" field.
func OutcomeMeasureDispersionTypeContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureDispersionType), v))
	})
}

// OutcomeMeasureUnitOfMeasureEQ applies the EQ predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureNEQ applies the NEQ predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureIn applies the In predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureUnitOfMeasure), v...))
	})
}

// OutcomeMeasureUnitOfMeasureNotIn applies the NotIn predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureUnitOfMeasure), v...))
	})
}

// OutcomeMeasureUnitOfMeasureGT applies the GT predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureGTE applies the GTE predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureLT applies the LT predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureLTE applies the LTE predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureContains applies the Contains predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureHasPrefix applies the HasPrefix predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureHasSuffix applies the HasSuffix predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureEqualFold applies the EqualFold predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureUnitOfMeasureContainsFold applies the ContainsFold predicate on the "outcome_measure_unit_of_measure" field.
func OutcomeMeasureUnitOfMeasureContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureUnitOfMeasure), v))
	})
}

// OutcomeMeasureCalculatePctEQ applies the EQ predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctNEQ applies the NEQ predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctIn applies the In predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureCalculatePct), v...))
	})
}

// OutcomeMeasureCalculatePctNotIn applies the NotIn predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureCalculatePct), v...))
	})
}

// OutcomeMeasureCalculatePctGT applies the GT predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctGTE applies the GTE predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctLT applies the LT predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctLTE applies the LTE predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctContains applies the Contains predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctHasPrefix applies the HasPrefix predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctHasSuffix applies the HasSuffix predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctEqualFold applies the EqualFold predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureCalculatePctContainsFold applies the ContainsFold predicate on the "outcome_measure_calculate_pct" field.
func OutcomeMeasureCalculatePctContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureCalculatePct), v))
	})
}

// OutcomeMeasureTimeFrameEQ applies the EQ predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameNEQ applies the NEQ predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameIn applies the In predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureTimeFrame), v...))
	})
}

// OutcomeMeasureTimeFrameNotIn applies the NotIn predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureTimeFrame), v...))
	})
}

// OutcomeMeasureTimeFrameGT applies the GT predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameGTE applies the GTE predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameLT applies the LT predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameLTE applies the LTE predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameContains applies the Contains predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameHasPrefix applies the HasPrefix predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameHasSuffix applies the HasSuffix predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameEqualFold applies the EqualFold predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTimeFrameContainsFold applies the ContainsFold predicate on the "outcome_measure_time_frame" field.
func OutcomeMeasureTimeFrameContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureTimeFrame), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedEQ applies the EQ predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedNEQ applies the NEQ predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedIn applies the In predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v...))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedNotIn applies the NotIn predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v...))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedGT applies the GT predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedGTE applies the GTE predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedLT applies the LT predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedLTE applies the LTE predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedContains applies the Contains predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedHasPrefix applies the HasPrefix predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedHasSuffix applies the HasSuffix predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedEqualFold applies the EqualFold predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureTypeUnitsAnalyzedContainsFold applies the ContainsFold predicate on the "outcome_measure_type_units_analyzed" field.
func OutcomeMeasureTypeUnitsAnalyzedContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureTypeUnitsAnalyzed), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedEQ applies the EQ predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedNEQ applies the NEQ predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedNEQ(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedIn applies the In predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasureDenomUnitsSelected), v...))
	})
}

// OutcomeMeasureDenomUnitsSelectedNotIn applies the NotIn predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedNotIn(vs ...string) predicate.OutcomeMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasureDenomUnitsSelected), v...))
	})
}

// OutcomeMeasureDenomUnitsSelectedGT applies the GT predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedGT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedGTE applies the GTE predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedGTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedLT applies the LT predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedLT(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedLTE applies the LTE predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedLTE(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedContains applies the Contains predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedContains(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedHasPrefix applies the HasPrefix predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedHasPrefix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedHasSuffix applies the HasSuffix predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedHasSuffix(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedEqualFold applies the EqualFold predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedEqualFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// OutcomeMeasureDenomUnitsSelectedContainsFold applies the ContainsFold predicate on the "outcome_measure_denom_units_selected" field.
func OutcomeMeasureDenomUnitsSelectedContainsFold(v string) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasureDenomUnitsSelected), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeMeasuresModule) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
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

// HasOutcomeGroupList applies the HasEdge predicate on the "outcome_group_list" edge.
func HasOutcomeGroupList() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeGroupListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeGroupListTable, OutcomeGroupListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeGroupListWith applies the HasEdge predicate on the "outcome_group_list" edge with a given conditions (other predicates).
func HasOutcomeGroupListWith(preds ...predicate.OutcomeGroup) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeGroupListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeGroupListTable, OutcomeGroupListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeOverviewList applies the HasEdge predicate on the "outcome_overview_list" edge.
func HasOutcomeOverviewList() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeOverviewListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeOverviewListTable, OutcomeOverviewListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeOverviewListWith applies the HasEdge predicate on the "outcome_overview_list" edge with a given conditions (other predicates).
func HasOutcomeOverviewListWith(preds ...predicate.OutcomeOverview) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeOverviewListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeOverviewListTable, OutcomeOverviewListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeDenomList applies the HasEdge predicate on the "outcome_denom_list" edge.
func HasOutcomeDenomList() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeDenomListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeDenomListTable, OutcomeDenomListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeDenomListWith applies the HasEdge predicate on the "outcome_denom_list" edge with a given conditions (other predicates).
func HasOutcomeDenomListWith(preds ...predicate.OutcomeDenom) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeDenomListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeDenomListTable, OutcomeDenomListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeClassList applies the HasEdge predicate on the "outcome_class_list" edge.
func HasOutcomeClassList() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassListTable, OutcomeClassListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeClassListWith applies the HasEdge predicate on the "outcome_class_list" edge with a given conditions (other predicates).
func HasOutcomeClassListWith(preds ...predicate.OutcomeClass) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeClassListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeClassListTable, OutcomeClassListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutcomeAnalysisList applies the HasEdge predicate on the "outcome_analysis_list" edge.
func HasOutcomeAnalysisList() predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeAnalysisListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeAnalysisListTable, OutcomeAnalysisListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutcomeAnalysisListWith applies the HasEdge predicate on the "outcome_analysis_list" edge with a given conditions (other predicates).
func HasOutcomeAnalysisListWith(preds ...predicate.OutcomeAnalysis) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutcomeAnalysisListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutcomeAnalysisListTable, OutcomeAnalysisListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OutcomeMeasure) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeMeasure) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeMeasure) predicate.OutcomeMeasure {
	return predicate.OutcomeMeasure(func(s *sql.Selector) {
		p(s.Not())
	})
}
