// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasure

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineMeasureTitle applies equality check predicate on the "baseline_measure_title" field. It's identical to BaselineMeasureTitleEQ.
func BaselineMeasureTitle(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureDescription applies equality check predicate on the "baseline_measure_description" field. It's identical to BaselineMeasureDescriptionEQ.
func BaselineMeasureDescription(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasurePopulationDescription applies equality check predicate on the "baseline_measure_population_description" field. It's identical to BaselineMeasurePopulationDescriptionEQ.
func BaselineMeasurePopulationDescription(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasureParamType applies equality check predicate on the "baseline_measure_param_type" field. It's identical to BaselineMeasureParamTypeEQ.
func BaselineMeasureParamType(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureDispersionType applies equality check predicate on the "baseline_measure_dispersion_type" field. It's identical to BaselineMeasureDispersionTypeEQ.
func BaselineMeasureDispersionType(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureUnitOfMeasure applies equality check predicate on the "baseline_measure_unit_of_measure" field. It's identical to BaselineMeasureUnitOfMeasureEQ.
func BaselineMeasureUnitOfMeasure(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureCalculatePct applies equality check predicate on the "baseline_measure_calculate_pct" field. It's identical to BaselineMeasureCalculatePctEQ.
func BaselineMeasureCalculatePct(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureDenomUnitsSelected applies equality check predicate on the "baseline_measure_denom_units_selected" field. It's identical to BaselineMeasureDenomUnitsSelectedEQ.
func BaselineMeasureDenomUnitsSelected(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureTitleEQ applies the EQ predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleNEQ applies the NEQ predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleIn applies the In predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureTitle), v...))
	})
}

// BaselineMeasureTitleNotIn applies the NotIn predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureTitle), v...))
	})
}

// BaselineMeasureTitleGT applies the GT predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleGTE applies the GTE predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleLT applies the LT predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleLTE applies the LTE predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleContains applies the Contains predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleHasPrefix applies the HasPrefix predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleHasSuffix applies the HasSuffix predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleEqualFold applies the EqualFold predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureTitleContainsFold applies the ContainsFold predicate on the "baseline_measure_title" field.
func BaselineMeasureTitleContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureTitle), v))
	})
}

// BaselineMeasureDescriptionEQ applies the EQ predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionNEQ applies the NEQ predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionIn applies the In predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDescription), v...))
	})
}

// BaselineMeasureDescriptionNotIn applies the NotIn predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDescription), v...))
	})
}

// BaselineMeasureDescriptionGT applies the GT predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionGTE applies the GTE predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionLT applies the LT predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionLTE applies the LTE predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionContains applies the Contains predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionHasPrefix applies the HasPrefix predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionHasSuffix applies the HasSuffix predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionEqualFold applies the EqualFold predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasureDescriptionContainsFold applies the ContainsFold predicate on the "baseline_measure_description" field.
func BaselineMeasureDescriptionContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionEQ applies the EQ predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionNEQ applies the NEQ predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionIn applies the In predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurePopulationDescription), v...))
	})
}

// BaselineMeasurePopulationDescriptionNotIn applies the NotIn predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurePopulationDescription), v...))
	})
}

// BaselineMeasurePopulationDescriptionGT applies the GT predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionGTE applies the GTE predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionLT applies the LT predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionLTE applies the LTE predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionContains applies the Contains predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionHasPrefix applies the HasPrefix predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionHasSuffix applies the HasSuffix predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionEqualFold applies the EqualFold predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasurePopulationDescriptionContainsFold applies the ContainsFold predicate on the "baseline_measure_population_description" field.
func BaselineMeasurePopulationDescriptionContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurePopulationDescription), v))
	})
}

// BaselineMeasureParamTypeEQ applies the EQ predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeNEQ applies the NEQ predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeIn applies the In predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureParamType), v...))
	})
}

// BaselineMeasureParamTypeNotIn applies the NotIn predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureParamType), v...))
	})
}

// BaselineMeasureParamTypeGT applies the GT predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeGTE applies the GTE predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeLT applies the LT predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeLTE applies the LTE predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeContains applies the Contains predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeHasPrefix applies the HasPrefix predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeHasSuffix applies the HasSuffix predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeEqualFold applies the EqualFold predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureParamTypeContainsFold applies the ContainsFold predicate on the "baseline_measure_param_type" field.
func BaselineMeasureParamTypeContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureParamType), v))
	})
}

// BaselineMeasureDispersionTypeEQ applies the EQ predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeNEQ applies the NEQ predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeIn applies the In predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDispersionType), v...))
	})
}

// BaselineMeasureDispersionTypeNotIn applies the NotIn predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDispersionType), v...))
	})
}

// BaselineMeasureDispersionTypeGT applies the GT predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeGTE applies the GTE predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeLT applies the LT predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeLTE applies the LTE predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeContains applies the Contains predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeHasPrefix applies the HasPrefix predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeHasSuffix applies the HasSuffix predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeEqualFold applies the EqualFold predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureDispersionTypeContainsFold applies the ContainsFold predicate on the "baseline_measure_dispersion_type" field.
func BaselineMeasureDispersionTypeContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDispersionType), v))
	})
}

// BaselineMeasureUnitOfMeasureEQ applies the EQ predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureNEQ applies the NEQ predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureIn applies the In predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureUnitOfMeasure), v...))
	})
}

// BaselineMeasureUnitOfMeasureNotIn applies the NotIn predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureUnitOfMeasure), v...))
	})
}

// BaselineMeasureUnitOfMeasureGT applies the GT predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureGTE applies the GTE predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureLT applies the LT predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureLTE applies the LTE predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureContains applies the Contains predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureHasPrefix applies the HasPrefix predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureHasSuffix applies the HasSuffix predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureEqualFold applies the EqualFold predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureUnitOfMeasureContainsFold applies the ContainsFold predicate on the "baseline_measure_unit_of_measure" field.
func BaselineMeasureUnitOfMeasureContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureUnitOfMeasure), v))
	})
}

// BaselineMeasureCalculatePctEQ applies the EQ predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctNEQ applies the NEQ predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctIn applies the In predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureCalculatePct), v...))
	})
}

// BaselineMeasureCalculatePctNotIn applies the NotIn predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureCalculatePct), v...))
	})
}

// BaselineMeasureCalculatePctGT applies the GT predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctGTE applies the GTE predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctLT applies the LT predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctLTE applies the LTE predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctContains applies the Contains predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctHasPrefix applies the HasPrefix predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctHasSuffix applies the HasSuffix predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctEqualFold applies the EqualFold predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureCalculatePctContainsFold applies the ContainsFold predicate on the "baseline_measure_calculate_pct" field.
func BaselineMeasureCalculatePctContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureCalculatePct), v))
	})
}

// BaselineMeasureDenomUnitsSelectedEQ applies the EQ predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedNEQ applies the NEQ predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedNEQ(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedIn applies the In predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasureDenomUnitsSelected), v...))
	})
}

// BaselineMeasureDenomUnitsSelectedNotIn applies the NotIn predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedNotIn(vs ...string) predicate.BaselineMeasure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasureDenomUnitsSelected), v...))
	})
}

// BaselineMeasureDenomUnitsSelectedGT applies the GT predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedGT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedGTE applies the GTE predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedGTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedLT applies the LT predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedLT(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedLTE applies the LTE predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedLTE(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedContains applies the Contains predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedContains(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedHasPrefix applies the HasPrefix predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedHasPrefix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedHasSuffix applies the HasSuffix predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedHasSuffix(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedEqualFold applies the EqualFold predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedEqualFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// BaselineMeasureDenomUnitsSelectedContainsFold applies the ContainsFold predicate on the "baseline_measure_denom_units_selected" field.
func BaselineMeasureDenomUnitsSelectedContainsFold(v string) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasureDenomUnitsSelected), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineCharacteristicsModule) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
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

// HasBaselineMeasureDenomList applies the HasEdge predicate on the "baseline_measure_denom_list" edge.
func HasBaselineMeasureDenomList() predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureDenomListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureDenomListTable, BaselineMeasureDenomListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineMeasureDenomListWith applies the HasEdge predicate on the "baseline_measure_denom_list" edge with a given conditions (other predicates).
func HasBaselineMeasureDenomListWith(preds ...predicate.BaselineMeasureDenom) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureDenomListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureDenomListTable, BaselineMeasureDenomListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineClassList applies the HasEdge predicate on the "baseline_class_list" edge.
func HasBaselineClassList() predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassListTable, BaselineClassListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineClassListWith applies the HasEdge predicate on the "baseline_class_list" edge with a given conditions (other predicates).
func HasBaselineClassListWith(preds ...predicate.BaselineClass) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineClassListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineClassListTable, BaselineClassListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineMeasure) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineMeasure) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
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
func Not(p predicate.BaselineMeasure) predicate.BaselineMeasure {
	return predicate.BaselineMeasure(func(s *sql.Selector) {
		p(s.Not())
	})
}
