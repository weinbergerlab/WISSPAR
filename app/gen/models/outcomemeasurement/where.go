// Code generated (@generated) by entc, DO NOT EDIT.

package outcomemeasurement

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OutcomeMeasurementGroupID applies equality check predicate on the "outcome_measurement_group_id" field. It's identical to OutcomeMeasurementGroupIDEQ.
func OutcomeMeasurementGroupID(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementValue applies equality check predicate on the "outcome_measurement_value" field. It's identical to OutcomeMeasurementValueEQ.
func OutcomeMeasurementValue(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementSpread applies equality check predicate on the "outcome_measurement_spread" field. It's identical to OutcomeMeasurementSpreadEQ.
func OutcomeMeasurementSpread(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementLowerLimit applies equality check predicate on the "outcome_measurement_lower_limit" field. It's identical to OutcomeMeasurementLowerLimitEQ.
func OutcomeMeasurementLowerLimit(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementUpperLimit applies equality check predicate on the "outcome_measurement_upper_limit" field. It's identical to OutcomeMeasurementUpperLimitEQ.
func OutcomeMeasurementUpperLimit(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementComment applies equality check predicate on the "outcome_measurement_comment" field. It's identical to OutcomeMeasurementCommentEQ.
func OutcomeMeasurementComment(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementGroupIDEQ applies the EQ predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDNEQ applies the NEQ predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDIn applies the In predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementGroupID), v...))
	})
}

// OutcomeMeasurementGroupIDNotIn applies the NotIn predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementGroupID), v...))
	})
}

// OutcomeMeasurementGroupIDGT applies the GT predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDGTE applies the GTE predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDLT applies the LT predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDLTE applies the LTE predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDContains applies the Contains predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDHasPrefix applies the HasPrefix predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDHasSuffix applies the HasSuffix predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDEqualFold applies the EqualFold predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementGroupIDContainsFold applies the ContainsFold predicate on the "outcome_measurement_group_id" field.
func OutcomeMeasurementGroupIDContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementGroupID), v))
	})
}

// OutcomeMeasurementValueEQ applies the EQ predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueNEQ applies the NEQ predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueIn applies the In predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementValue), v...))
	})
}

// OutcomeMeasurementValueNotIn applies the NotIn predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementValue), v...))
	})
}

// OutcomeMeasurementValueGT applies the GT predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueGTE applies the GTE predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueLT applies the LT predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueLTE applies the LTE predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueContains applies the Contains predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueHasPrefix applies the HasPrefix predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueHasSuffix applies the HasSuffix predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueEqualFold applies the EqualFold predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementValueContainsFold applies the ContainsFold predicate on the "outcome_measurement_value" field.
func OutcomeMeasurementValueContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementValue), v))
	})
}

// OutcomeMeasurementSpreadEQ applies the EQ predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadNEQ applies the NEQ predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadIn applies the In predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementSpread), v...))
	})
}

// OutcomeMeasurementSpreadNotIn applies the NotIn predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementSpread), v...))
	})
}

// OutcomeMeasurementSpreadGT applies the GT predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadGTE applies the GTE predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadLT applies the LT predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadLTE applies the LTE predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadContains applies the Contains predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadHasPrefix applies the HasPrefix predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadHasSuffix applies the HasSuffix predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadEqualFold applies the EqualFold predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementSpreadContainsFold applies the ContainsFold predicate on the "outcome_measurement_spread" field.
func OutcomeMeasurementSpreadContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementSpread), v))
	})
}

// OutcomeMeasurementLowerLimitEQ applies the EQ predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitNEQ applies the NEQ predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitIn applies the In predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementLowerLimit), v...))
	})
}

// OutcomeMeasurementLowerLimitNotIn applies the NotIn predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementLowerLimit), v...))
	})
}

// OutcomeMeasurementLowerLimitGT applies the GT predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitGTE applies the GTE predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitLT applies the LT predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitLTE applies the LTE predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitContains applies the Contains predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitHasPrefix applies the HasPrefix predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitHasSuffix applies the HasSuffix predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitEqualFold applies the EqualFold predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementLowerLimitContainsFold applies the ContainsFold predicate on the "outcome_measurement_lower_limit" field.
func OutcomeMeasurementLowerLimitContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementLowerLimit), v))
	})
}

// OutcomeMeasurementUpperLimitEQ applies the EQ predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitNEQ applies the NEQ predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitIn applies the In predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementUpperLimit), v...))
	})
}

// OutcomeMeasurementUpperLimitNotIn applies the NotIn predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementUpperLimit), v...))
	})
}

// OutcomeMeasurementUpperLimitGT applies the GT predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitGTE applies the GTE predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitLT applies the LT predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitLTE applies the LTE predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitContains applies the Contains predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitHasPrefix applies the HasPrefix predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitHasSuffix applies the HasSuffix predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitEqualFold applies the EqualFold predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementUpperLimitContainsFold applies the ContainsFold predicate on the "outcome_measurement_upper_limit" field.
func OutcomeMeasurementUpperLimitContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementUpperLimit), v))
	})
}

// OutcomeMeasurementCommentEQ applies the EQ predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentNEQ applies the NEQ predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentNEQ(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentIn applies the In predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcomeMeasurementComment), v...))
	})
}

// OutcomeMeasurementCommentNotIn applies the NotIn predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentNotIn(vs ...string) predicate.OutcomeMeasurement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcomeMeasurementComment), v...))
	})
}

// OutcomeMeasurementCommentGT applies the GT predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentGT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentGTE applies the GTE predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentGTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentLT applies the LT predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentLT(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentLTE applies the LTE predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentLTE(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentContains applies the Contains predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentContains(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentHasPrefix applies the HasPrefix predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentHasPrefix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentHasSuffix applies the HasSuffix predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentHasSuffix(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentEqualFold applies the EqualFold predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentEqualFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// OutcomeMeasurementCommentContainsFold applies the ContainsFold predicate on the "outcome_measurement_comment" field.
func OutcomeMeasurementCommentContainsFold(v string) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutcomeMeasurementComment), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OutcomeCategory) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
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
func And(predicates ...predicate.OutcomeMeasurement) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OutcomeMeasurement) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
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
func Not(p predicate.OutcomeMeasurement) predicate.OutcomeMeasurement {
	return predicate.OutcomeMeasurement(func(s *sql.Selector) {
		p(s.Not())
	})
}
