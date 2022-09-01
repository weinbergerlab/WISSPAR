// Code generated (@generated) by entc, DO NOT EDIT.

package vaccine

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// VaccineName applies equality check predicate on the "vaccine_name" field. It's identical to VaccineNameEQ.
func VaccineName(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVaccineName), v))
	})
}

// VaccineNameEQ applies the EQ predicate on the "vaccine_name" field.
func VaccineNameEQ(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVaccineName), v))
	})
}

// VaccineNameNEQ applies the NEQ predicate on the "vaccine_name" field.
func VaccineNameNEQ(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVaccineName), v))
	})
}

// VaccineNameIn applies the In predicate on the "vaccine_name" field.
func VaccineNameIn(vs ...string) predicate.Vaccine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vaccine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVaccineName), v...))
	})
}

// VaccineNameNotIn applies the NotIn predicate on the "vaccine_name" field.
func VaccineNameNotIn(vs ...string) predicate.Vaccine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vaccine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVaccineName), v...))
	})
}

// VaccineNameGT applies the GT predicate on the "vaccine_name" field.
func VaccineNameGT(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVaccineName), v))
	})
}

// VaccineNameGTE applies the GTE predicate on the "vaccine_name" field.
func VaccineNameGTE(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVaccineName), v))
	})
}

// VaccineNameLT applies the LT predicate on the "vaccine_name" field.
func VaccineNameLT(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVaccineName), v))
	})
}

// VaccineNameLTE applies the LTE predicate on the "vaccine_name" field.
func VaccineNameLTE(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVaccineName), v))
	})
}

// VaccineNameContains applies the Contains predicate on the "vaccine_name" field.
func VaccineNameContains(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVaccineName), v))
	})
}

// VaccineNameHasPrefix applies the HasPrefix predicate on the "vaccine_name" field.
func VaccineNameHasPrefix(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVaccineName), v))
	})
}

// VaccineNameHasSuffix applies the HasSuffix predicate on the "vaccine_name" field.
func VaccineNameHasSuffix(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVaccineName), v))
	})
}

// VaccineNameEqualFold applies the EqualFold predicate on the "vaccine_name" field.
func VaccineNameEqualFold(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVaccineName), v))
	})
}

// VaccineNameContainsFold applies the ContainsFold predicate on the "vaccine_name" field.
func VaccineNameContainsFold(v string) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVaccineName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vaccine) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vaccine) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
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
func Not(p predicate.Vaccine) predicate.Vaccine {
	return predicate.Vaccine(func(s *sql.Selector) {
		p(s.Not())
	})
}
