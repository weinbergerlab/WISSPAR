// Code generated (@generated) by entc, DO NOT EDIT.

package studylocation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
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
func IDGT(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LocationFacility applies equality check predicate on the "LocationFacility" field. It's identical to LocationFacilityEQ.
func LocationFacility(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationFacility), v))
	})
}

// LocationCity applies equality check predicate on the "LocationCity" field. It's identical to LocationCityEQ.
func LocationCity(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCity), v))
	})
}

// LocationCountry applies equality check predicate on the "LocationCountry" field. It's identical to LocationCountryEQ.
func LocationCountry(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryCode applies equality check predicate on the "LocationCountryCode" field. It's identical to LocationCountryCodeEQ.
func LocationCountryCode(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCountryCode), v))
	})
}

// LocationContinentName applies equality check predicate on the "LocationContinentName" field. It's identical to LocationContinentNameEQ.
func LocationContinentName(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentCode applies equality check predicate on the "LocationContinentCode" field. It's identical to LocationContinentCodeEQ.
func LocationContinentCode(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationContinentCode), v))
	})
}

// LocationFacilityEQ applies the EQ predicate on the "LocationFacility" field.
func LocationFacilityEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityNEQ applies the NEQ predicate on the "LocationFacility" field.
func LocationFacilityNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityIn applies the In predicate on the "LocationFacility" field.
func LocationFacilityIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationFacility), v...))
	})
}

// LocationFacilityNotIn applies the NotIn predicate on the "LocationFacility" field.
func LocationFacilityNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationFacility), v...))
	})
}

// LocationFacilityGT applies the GT predicate on the "LocationFacility" field.
func LocationFacilityGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityGTE applies the GTE predicate on the "LocationFacility" field.
func LocationFacilityGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityLT applies the LT predicate on the "LocationFacility" field.
func LocationFacilityLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityLTE applies the LTE predicate on the "LocationFacility" field.
func LocationFacilityLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityContains applies the Contains predicate on the "LocationFacility" field.
func LocationFacilityContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityHasPrefix applies the HasPrefix predicate on the "LocationFacility" field.
func LocationFacilityHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityHasSuffix applies the HasSuffix predicate on the "LocationFacility" field.
func LocationFacilityHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityEqualFold applies the EqualFold predicate on the "LocationFacility" field.
func LocationFacilityEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationFacility), v))
	})
}

// LocationFacilityContainsFold applies the ContainsFold predicate on the "LocationFacility" field.
func LocationFacilityContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationFacility), v))
	})
}

// LocationCityEQ applies the EQ predicate on the "LocationCity" field.
func LocationCityEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCity), v))
	})
}

// LocationCityNEQ applies the NEQ predicate on the "LocationCity" field.
func LocationCityNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationCity), v))
	})
}

// LocationCityIn applies the In predicate on the "LocationCity" field.
func LocationCityIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationCity), v...))
	})
}

// LocationCityNotIn applies the NotIn predicate on the "LocationCity" field.
func LocationCityNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationCity), v...))
	})
}

// LocationCityGT applies the GT predicate on the "LocationCity" field.
func LocationCityGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationCity), v))
	})
}

// LocationCityGTE applies the GTE predicate on the "LocationCity" field.
func LocationCityGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationCity), v))
	})
}

// LocationCityLT applies the LT predicate on the "LocationCity" field.
func LocationCityLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationCity), v))
	})
}

// LocationCityLTE applies the LTE predicate on the "LocationCity" field.
func LocationCityLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationCity), v))
	})
}

// LocationCityContains applies the Contains predicate on the "LocationCity" field.
func LocationCityContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationCity), v))
	})
}

// LocationCityHasPrefix applies the HasPrefix predicate on the "LocationCity" field.
func LocationCityHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationCity), v))
	})
}

// LocationCityHasSuffix applies the HasSuffix predicate on the "LocationCity" field.
func LocationCityHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationCity), v))
	})
}

// LocationCityEqualFold applies the EqualFold predicate on the "LocationCity" field.
func LocationCityEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationCity), v))
	})
}

// LocationCityContainsFold applies the ContainsFold predicate on the "LocationCity" field.
func LocationCityContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationCity), v))
	})
}

// LocationCountryEQ applies the EQ predicate on the "LocationCountry" field.
func LocationCountryEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryNEQ applies the NEQ predicate on the "LocationCountry" field.
func LocationCountryNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryIn applies the In predicate on the "LocationCountry" field.
func LocationCountryIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationCountry), v...))
	})
}

// LocationCountryNotIn applies the NotIn predicate on the "LocationCountry" field.
func LocationCountryNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationCountry), v...))
	})
}

// LocationCountryGT applies the GT predicate on the "LocationCountry" field.
func LocationCountryGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryGTE applies the GTE predicate on the "LocationCountry" field.
func LocationCountryGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryLT applies the LT predicate on the "LocationCountry" field.
func LocationCountryLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryLTE applies the LTE predicate on the "LocationCountry" field.
func LocationCountryLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryContains applies the Contains predicate on the "LocationCountry" field.
func LocationCountryContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryHasPrefix applies the HasPrefix predicate on the "LocationCountry" field.
func LocationCountryHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryHasSuffix applies the HasSuffix predicate on the "LocationCountry" field.
func LocationCountryHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryEqualFold applies the EqualFold predicate on the "LocationCountry" field.
func LocationCountryEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryContainsFold applies the ContainsFold predicate on the "LocationCountry" field.
func LocationCountryContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationCountry), v))
	})
}

// LocationCountryCodeEQ applies the EQ predicate on the "LocationCountryCode" field.
func LocationCountryCodeEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeNEQ applies the NEQ predicate on the "LocationCountryCode" field.
func LocationCountryCodeNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeIn applies the In predicate on the "LocationCountryCode" field.
func LocationCountryCodeIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationCountryCode), v...))
	})
}

// LocationCountryCodeNotIn applies the NotIn predicate on the "LocationCountryCode" field.
func LocationCountryCodeNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationCountryCode), v...))
	})
}

// LocationCountryCodeGT applies the GT predicate on the "LocationCountryCode" field.
func LocationCountryCodeGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeGTE applies the GTE predicate on the "LocationCountryCode" field.
func LocationCountryCodeGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeLT applies the LT predicate on the "LocationCountryCode" field.
func LocationCountryCodeLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeLTE applies the LTE predicate on the "LocationCountryCode" field.
func LocationCountryCodeLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeContains applies the Contains predicate on the "LocationCountryCode" field.
func LocationCountryCodeContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeHasPrefix applies the HasPrefix predicate on the "LocationCountryCode" field.
func LocationCountryCodeHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeHasSuffix applies the HasSuffix predicate on the "LocationCountryCode" field.
func LocationCountryCodeHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeEqualFold applies the EqualFold predicate on the "LocationCountryCode" field.
func LocationCountryCodeEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationCountryCode), v))
	})
}

// LocationCountryCodeContainsFold applies the ContainsFold predicate on the "LocationCountryCode" field.
func LocationCountryCodeContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationCountryCode), v))
	})
}

// LocationContinentNameEQ applies the EQ predicate on the "LocationContinentName" field.
func LocationContinentNameEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameNEQ applies the NEQ predicate on the "LocationContinentName" field.
func LocationContinentNameNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameIn applies the In predicate on the "LocationContinentName" field.
func LocationContinentNameIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationContinentName), v...))
	})
}

// LocationContinentNameNotIn applies the NotIn predicate on the "LocationContinentName" field.
func LocationContinentNameNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationContinentName), v...))
	})
}

// LocationContinentNameGT applies the GT predicate on the "LocationContinentName" field.
func LocationContinentNameGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameGTE applies the GTE predicate on the "LocationContinentName" field.
func LocationContinentNameGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameLT applies the LT predicate on the "LocationContinentName" field.
func LocationContinentNameLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameLTE applies the LTE predicate on the "LocationContinentName" field.
func LocationContinentNameLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameContains applies the Contains predicate on the "LocationContinentName" field.
func LocationContinentNameContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameHasPrefix applies the HasPrefix predicate on the "LocationContinentName" field.
func LocationContinentNameHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameHasSuffix applies the HasSuffix predicate on the "LocationContinentName" field.
func LocationContinentNameHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameEqualFold applies the EqualFold predicate on the "LocationContinentName" field.
func LocationContinentNameEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentNameContainsFold applies the ContainsFold predicate on the "LocationContinentName" field.
func LocationContinentNameContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationContinentName), v))
	})
}

// LocationContinentCodeEQ applies the EQ predicate on the "LocationContinentCode" field.
func LocationContinentCodeEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeNEQ applies the NEQ predicate on the "LocationContinentCode" field.
func LocationContinentCodeNEQ(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeIn applies the In predicate on the "LocationContinentCode" field.
func LocationContinentCodeIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationContinentCode), v...))
	})
}

// LocationContinentCodeNotIn applies the NotIn predicate on the "LocationContinentCode" field.
func LocationContinentCodeNotIn(vs ...string) predicate.StudyLocation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StudyLocation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationContinentCode), v...))
	})
}

// LocationContinentCodeGT applies the GT predicate on the "LocationContinentCode" field.
func LocationContinentCodeGT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeGTE applies the GTE predicate on the "LocationContinentCode" field.
func LocationContinentCodeGTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeLT applies the LT predicate on the "LocationContinentCode" field.
func LocationContinentCodeLT(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeLTE applies the LTE predicate on the "LocationContinentCode" field.
func LocationContinentCodeLTE(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeContains applies the Contains predicate on the "LocationContinentCode" field.
func LocationContinentCodeContains(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeHasPrefix applies the HasPrefix predicate on the "LocationContinentCode" field.
func LocationContinentCodeHasPrefix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeHasSuffix applies the HasSuffix predicate on the "LocationContinentCode" field.
func LocationContinentCodeHasSuffix(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeEqualFold applies the EqualFold predicate on the "LocationContinentCode" field.
func LocationContinentCodeEqualFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationContinentCode), v))
	})
}

// LocationContinentCodeContainsFold applies the ContainsFold predicate on the "LocationContinentCode" field.
func LocationContinentCodeContainsFold(v string) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationContinentCode), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ClinicalTrial) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
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
func And(predicates ...predicate.StudyLocation) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StudyLocation) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
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
func Not(p predicate.StudyLocation) predicate.StudyLocation {
	return predicate.StudyLocation(func(s *sql.Selector) {
		p(s.Not())
	})
}
