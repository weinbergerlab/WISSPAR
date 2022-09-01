// Code generated (@generated) by entc, DO NOT EDIT.

package baselinemeasurement

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselineMeasurementGroupID applies equality check predicate on the "baseline_measurement_group_id" field. It's identical to BaselineMeasurementGroupIDEQ.
func BaselineMeasurementGroupID(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementValue applies equality check predicate on the "baseline_measurement_value" field. It's identical to BaselineMeasurementValueEQ.
func BaselineMeasurementValue(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementSpread applies equality check predicate on the "baseline_measurement_spread" field. It's identical to BaselineMeasurementSpreadEQ.
func BaselineMeasurementSpread(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementLowerLimit applies equality check predicate on the "baseline_measurement_lower_limit" field. It's identical to BaselineMeasurementLowerLimitEQ.
func BaselineMeasurementLowerLimit(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementUpperLimit applies equality check predicate on the "baseline_measurement_upper_limit" field. It's identical to BaselineMeasurementUpperLimitEQ.
func BaselineMeasurementUpperLimit(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementComment applies equality check predicate on the "baseline_measurement_comment" field. It's identical to BaselineMeasurementCommentEQ.
func BaselineMeasurementComment(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementGroupIDEQ applies the EQ predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDNEQ applies the NEQ predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDIn applies the In predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementGroupID), v...))
	})
}

// BaselineMeasurementGroupIDNotIn applies the NotIn predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementGroupID), v...))
	})
}

// BaselineMeasurementGroupIDGT applies the GT predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDGTE applies the GTE predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDLT applies the LT predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDLTE applies the LTE predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDContains applies the Contains predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDHasPrefix applies the HasPrefix predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDHasSuffix applies the HasSuffix predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDEqualFold applies the EqualFold predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementGroupIDContainsFold applies the ContainsFold predicate on the "baseline_measurement_group_id" field.
func BaselineMeasurementGroupIDContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementGroupID), v))
	})
}

// BaselineMeasurementValueEQ applies the EQ predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueNEQ applies the NEQ predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueIn applies the In predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementValue), v...))
	})
}

// BaselineMeasurementValueNotIn applies the NotIn predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementValue), v...))
	})
}

// BaselineMeasurementValueGT applies the GT predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueGTE applies the GTE predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueLT applies the LT predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueLTE applies the LTE predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueContains applies the Contains predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueHasPrefix applies the HasPrefix predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueHasSuffix applies the HasSuffix predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueEqualFold applies the EqualFold predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementValueContainsFold applies the ContainsFold predicate on the "baseline_measurement_value" field.
func BaselineMeasurementValueContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementValue), v))
	})
}

// BaselineMeasurementSpreadEQ applies the EQ predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadNEQ applies the NEQ predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadIn applies the In predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementSpread), v...))
	})
}

// BaselineMeasurementSpreadNotIn applies the NotIn predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementSpread), v...))
	})
}

// BaselineMeasurementSpreadGT applies the GT predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadGTE applies the GTE predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadLT applies the LT predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadLTE applies the LTE predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadContains applies the Contains predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadHasPrefix applies the HasPrefix predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadHasSuffix applies the HasSuffix predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadEqualFold applies the EqualFold predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementSpreadContainsFold applies the ContainsFold predicate on the "baseline_measurement_spread" field.
func BaselineMeasurementSpreadContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementSpread), v))
	})
}

// BaselineMeasurementLowerLimitEQ applies the EQ predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitNEQ applies the NEQ predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitIn applies the In predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementLowerLimit), v...))
	})
}

// BaselineMeasurementLowerLimitNotIn applies the NotIn predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementLowerLimit), v...))
	})
}

// BaselineMeasurementLowerLimitGT applies the GT predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitGTE applies the GTE predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitLT applies the LT predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitLTE applies the LTE predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitContains applies the Contains predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitHasPrefix applies the HasPrefix predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitHasSuffix applies the HasSuffix predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitEqualFold applies the EqualFold predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementLowerLimitContainsFold applies the ContainsFold predicate on the "baseline_measurement_lower_limit" field.
func BaselineMeasurementLowerLimitContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementLowerLimit), v))
	})
}

// BaselineMeasurementUpperLimitEQ applies the EQ predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitNEQ applies the NEQ predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitIn applies the In predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementUpperLimit), v...))
	})
}

// BaselineMeasurementUpperLimitNotIn applies the NotIn predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementUpperLimit), v...))
	})
}

// BaselineMeasurementUpperLimitGT applies the GT predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitGTE applies the GTE predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitLT applies the LT predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitLTE applies the LTE predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitContains applies the Contains predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitHasPrefix applies the HasPrefix predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitHasSuffix applies the HasSuffix predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitEqualFold applies the EqualFold predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementUpperLimitContainsFold applies the ContainsFold predicate on the "baseline_measurement_upper_limit" field.
func BaselineMeasurementUpperLimitContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementUpperLimit), v))
	})
}

// BaselineMeasurementCommentEQ applies the EQ predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentNEQ applies the NEQ predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentNEQ(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentIn applies the In predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineMeasurementComment), v...))
	})
}

// BaselineMeasurementCommentNotIn applies the NotIn predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentNotIn(vs ...string) predicate.BaselineMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineMeasurementComment), v...))
	})
}

// BaselineMeasurementCommentGT applies the GT predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentGT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentGTE applies the GTE predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentGTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentLT applies the LT predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentLT(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentLTE applies the LTE predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentLTE(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentContains applies the Contains predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentContains(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentHasPrefix applies the HasPrefix predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentHasPrefix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentHasSuffix applies the HasSuffix predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentHasSuffix(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentEqualFold applies the EqualFold predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentEqualFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineMeasurementComment), v))
	})
}

// BaselineMeasurementCommentContainsFold applies the ContainsFold predicate on the "baseline_measurement_comment" field.
func BaselineMeasurementCommentContainsFold(v string) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineMeasurementComment), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.BaselineCategory) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
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
func And(predicates ...predicate.BaselineMeasurement) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineMeasurement) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
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
func Not(p predicate.BaselineMeasurement) predicate.BaselineMeasurement {
	return predicate.BaselineMeasurement(func(s *sql.Selector) {
		p(s.Not())
	})
}
