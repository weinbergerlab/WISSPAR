// Code generated (@generated) by entc, DO NOT EDIT.

package studyeligibility

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
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
func IDGT(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EligibilityCriteria applies equality check predicate on the "EligibilityCriteria" field. It's identical to EligibilityCriteriaEQ.
func EligibilityCriteria(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEligibilityCriteria), v))
	})
}

// HealthyVolunteers applies equality check predicate on the "HealthyVolunteers" field. It's identical to HealthyVolunteersEQ.
func HealthyVolunteers(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHealthyVolunteers), v))
	})
}

// Gender applies equality check predicate on the "Gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// MinimumAge applies equality check predicate on the "MinimumAge" field. It's identical to MinimumAgeEQ.
func MinimumAge(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinimumAge), v))
	})
}

// MaximumAge applies equality check predicate on the "MaximumAge" field. It's identical to MaximumAgeEQ.
func MaximumAge(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaximumAge), v))
	})
}

// StdAgeList applies equality check predicate on the "StdAgeList" field. It's identical to StdAgeListEQ.
func StdAgeList(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStdAgeList), v))
	})
}

// Ethnicity applies equality check predicate on the "Ethnicity" field. It's identical to EthnicityEQ.
func Ethnicity(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEthnicity), v))
	})
}

// EligibilityCriteriaEQ applies the EQ predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaNEQ applies the NEQ predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaIn applies the In predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEligibilityCriteria), v...))
	})
}

// EligibilityCriteriaNotIn applies the NotIn predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEligibilityCriteria), v...))
	})
}

// EligibilityCriteriaGT applies the GT predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaGTE applies the GTE predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaLT applies the LT predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaLTE applies the LTE predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaContains applies the Contains predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaHasPrefix applies the HasPrefix predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaHasSuffix applies the HasSuffix predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaEqualFold applies the EqualFold predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEligibilityCriteria), v))
	})
}

// EligibilityCriteriaContainsFold applies the ContainsFold predicate on the "EligibilityCriteria" field.
func EligibilityCriteriaContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEligibilityCriteria), v))
	})
}

// HealthyVolunteersEQ applies the EQ predicate on the "HealthyVolunteers" field.
func HealthyVolunteersEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersNEQ applies the NEQ predicate on the "HealthyVolunteers" field.
func HealthyVolunteersNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersIn applies the In predicate on the "HealthyVolunteers" field.
func HealthyVolunteersIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHealthyVolunteers), v...))
	})
}

// HealthyVolunteersNotIn applies the NotIn predicate on the "HealthyVolunteers" field.
func HealthyVolunteersNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHealthyVolunteers), v...))
	})
}

// HealthyVolunteersGT applies the GT predicate on the "HealthyVolunteers" field.
func HealthyVolunteersGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersGTE applies the GTE predicate on the "HealthyVolunteers" field.
func HealthyVolunteersGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersLT applies the LT predicate on the "HealthyVolunteers" field.
func HealthyVolunteersLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersLTE applies the LTE predicate on the "HealthyVolunteers" field.
func HealthyVolunteersLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersContains applies the Contains predicate on the "HealthyVolunteers" field.
func HealthyVolunteersContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersHasPrefix applies the HasPrefix predicate on the "HealthyVolunteers" field.
func HealthyVolunteersHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersHasSuffix applies the HasSuffix predicate on the "HealthyVolunteers" field.
func HealthyVolunteersHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersEqualFold applies the EqualFold predicate on the "HealthyVolunteers" field.
func HealthyVolunteersEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHealthyVolunteers), v))
	})
}

// HealthyVolunteersContainsFold applies the ContainsFold predicate on the "HealthyVolunteers" field.
func HealthyVolunteersContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHealthyVolunteers), v))
	})
}

// GenderEQ applies the EQ predicate on the "Gender" field.
func GenderEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// GenderNEQ applies the NEQ predicate on the "Gender" field.
func GenderNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGender), v))
	})
}

// GenderIn applies the In predicate on the "Gender" field.
func GenderIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGender), v...))
	})
}

// GenderNotIn applies the NotIn predicate on the "Gender" field.
func GenderNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGender), v...))
	})
}

// GenderGT applies the GT predicate on the "Gender" field.
func GenderGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGender), v))
	})
}

// GenderGTE applies the GTE predicate on the "Gender" field.
func GenderGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGender), v))
	})
}

// GenderLT applies the LT predicate on the "Gender" field.
func GenderLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGender), v))
	})
}

// GenderLTE applies the LTE predicate on the "Gender" field.
func GenderLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGender), v))
	})
}

// GenderContains applies the Contains predicate on the "Gender" field.
func GenderContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGender), v))
	})
}

// GenderHasPrefix applies the HasPrefix predicate on the "Gender" field.
func GenderHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGender), v))
	})
}

// GenderHasSuffix applies the HasSuffix predicate on the "Gender" field.
func GenderHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGender), v))
	})
}

// GenderEqualFold applies the EqualFold predicate on the "Gender" field.
func GenderEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGender), v))
	})
}

// GenderContainsFold applies the ContainsFold predicate on the "Gender" field.
func GenderContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGender), v))
	})
}

// MinimumAgeEQ applies the EQ predicate on the "MinimumAge" field.
func MinimumAgeEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeNEQ applies the NEQ predicate on the "MinimumAge" field.
func MinimumAgeNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeIn applies the In predicate on the "MinimumAge" field.
func MinimumAgeIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMinimumAge), v...))
	})
}

// MinimumAgeNotIn applies the NotIn predicate on the "MinimumAge" field.
func MinimumAgeNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMinimumAge), v...))
	})
}

// MinimumAgeGT applies the GT predicate on the "MinimumAge" field.
func MinimumAgeGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeGTE applies the GTE predicate on the "MinimumAge" field.
func MinimumAgeGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeLT applies the LT predicate on the "MinimumAge" field.
func MinimumAgeLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeLTE applies the LTE predicate on the "MinimumAge" field.
func MinimumAgeLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeContains applies the Contains predicate on the "MinimumAge" field.
func MinimumAgeContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeHasPrefix applies the HasPrefix predicate on the "MinimumAge" field.
func MinimumAgeHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeHasSuffix applies the HasSuffix predicate on the "MinimumAge" field.
func MinimumAgeHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeEqualFold applies the EqualFold predicate on the "MinimumAge" field.
func MinimumAgeEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMinimumAge), v))
	})
}

// MinimumAgeContainsFold applies the ContainsFold predicate on the "MinimumAge" field.
func MinimumAgeContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMinimumAge), v))
	})
}

// MaximumAgeEQ applies the EQ predicate on the "MaximumAge" field.
func MaximumAgeEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeNEQ applies the NEQ predicate on the "MaximumAge" field.
func MaximumAgeNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeIn applies the In predicate on the "MaximumAge" field.
func MaximumAgeIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaximumAge), v...))
	})
}

// MaximumAgeNotIn applies the NotIn predicate on the "MaximumAge" field.
func MaximumAgeNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaximumAge), v...))
	})
}

// MaximumAgeGT applies the GT predicate on the "MaximumAge" field.
func MaximumAgeGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeGTE applies the GTE predicate on the "MaximumAge" field.
func MaximumAgeGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeLT applies the LT predicate on the "MaximumAge" field.
func MaximumAgeLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeLTE applies the LTE predicate on the "MaximumAge" field.
func MaximumAgeLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeContains applies the Contains predicate on the "MaximumAge" field.
func MaximumAgeContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeHasPrefix applies the HasPrefix predicate on the "MaximumAge" field.
func MaximumAgeHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeHasSuffix applies the HasSuffix predicate on the "MaximumAge" field.
func MaximumAgeHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeEqualFold applies the EqualFold predicate on the "MaximumAge" field.
func MaximumAgeEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMaximumAge), v))
	})
}

// MaximumAgeContainsFold applies the ContainsFold predicate on the "MaximumAge" field.
func MaximumAgeContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMaximumAge), v))
	})
}

// StdAgeListEQ applies the EQ predicate on the "StdAgeList" field.
func StdAgeListEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListNEQ applies the NEQ predicate on the "StdAgeList" field.
func StdAgeListNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListIn applies the In predicate on the "StdAgeList" field.
func StdAgeListIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStdAgeList), v...))
	})
}

// StdAgeListNotIn applies the NotIn predicate on the "StdAgeList" field.
func StdAgeListNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStdAgeList), v...))
	})
}

// StdAgeListGT applies the GT predicate on the "StdAgeList" field.
func StdAgeListGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListGTE applies the GTE predicate on the "StdAgeList" field.
func StdAgeListGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListLT applies the LT predicate on the "StdAgeList" field.
func StdAgeListLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListLTE applies the LTE predicate on the "StdAgeList" field.
func StdAgeListLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListContains applies the Contains predicate on the "StdAgeList" field.
func StdAgeListContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListHasPrefix applies the HasPrefix predicate on the "StdAgeList" field.
func StdAgeListHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListHasSuffix applies the HasSuffix predicate on the "StdAgeList" field.
func StdAgeListHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListEqualFold applies the EqualFold predicate on the "StdAgeList" field.
func StdAgeListEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStdAgeList), v))
	})
}

// StdAgeListContainsFold applies the ContainsFold predicate on the "StdAgeList" field.
func StdAgeListContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStdAgeList), v))
	})
}

// EthnicityEQ applies the EQ predicate on the "Ethnicity" field.
func EthnicityEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEthnicity), v))
	})
}

// EthnicityNEQ applies the NEQ predicate on the "Ethnicity" field.
func EthnicityNEQ(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEthnicity), v))
	})
}

// EthnicityIn applies the In predicate on the "Ethnicity" field.
func EthnicityIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEthnicity), v...))
	})
}

// EthnicityNotIn applies the NotIn predicate on the "Ethnicity" field.
func EthnicityNotIn(vs ...string) predicate.StudyEligibility {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyEligibility(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEthnicity), v...))
	})
}

// EthnicityGT applies the GT predicate on the "Ethnicity" field.
func EthnicityGT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEthnicity), v))
	})
}

// EthnicityGTE applies the GTE predicate on the "Ethnicity" field.
func EthnicityGTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEthnicity), v))
	})
}

// EthnicityLT applies the LT predicate on the "Ethnicity" field.
func EthnicityLT(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEthnicity), v))
	})
}

// EthnicityLTE applies the LTE predicate on the "Ethnicity" field.
func EthnicityLTE(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEthnicity), v))
	})
}

// EthnicityContains applies the Contains predicate on the "Ethnicity" field.
func EthnicityContains(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEthnicity), v))
	})
}

// EthnicityHasPrefix applies the HasPrefix predicate on the "Ethnicity" field.
func EthnicityHasPrefix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEthnicity), v))
	})
}

// EthnicityHasSuffix applies the HasSuffix predicate on the "Ethnicity" field.
func EthnicityHasSuffix(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEthnicity), v))
	})
}

// EthnicityEqualFold applies the EqualFold predicate on the "Ethnicity" field.
func EthnicityEqualFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEthnicity), v))
	})
}

// EthnicityContainsFold applies the ContainsFold predicate on the "Ethnicity" field.
func EthnicityContainsFold(v string) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEthnicity), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ClinicalTrial) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
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
func And(predicates ...predicate.StudyEligibility) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StudyEligibility) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
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
func Not(p predicate.StudyEligibility) predicate.StudyEligibility {
	return predicate.StudyEligibility(func(s *sql.Selector) {
		p(s.Not())
	})
}
