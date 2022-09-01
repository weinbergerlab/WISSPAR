// Code generated (@generated) by entc, DO NOT EDIT.

package usecase

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UseCaseName applies equality check predicate on the "use_case_name" field. It's identical to UseCaseNameEQ.
func UseCaseName(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseName), v))
	})
}

// UseCaseDescription applies equality check predicate on the "use_case_description" field. It's identical to UseCaseDescriptionEQ.
func UseCaseDescription(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseCode applies equality check predicate on the "use_case_code" field. It's identical to UseCaseCodeEQ.
func UseCaseCode(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseNameEQ applies the EQ predicate on the "use_case_name" field.
func UseCaseNameEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameNEQ applies the NEQ predicate on the "use_case_name" field.
func UseCaseNameNEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameIn applies the In predicate on the "use_case_name" field.
func UseCaseNameIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUseCaseName), v...))
	})
}

// UseCaseNameNotIn applies the NotIn predicate on the "use_case_name" field.
func UseCaseNameNotIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUseCaseName), v...))
	})
}

// UseCaseNameGT applies the GT predicate on the "use_case_name" field.
func UseCaseNameGT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameGTE applies the GTE predicate on the "use_case_name" field.
func UseCaseNameGTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameLT applies the LT predicate on the "use_case_name" field.
func UseCaseNameLT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameLTE applies the LTE predicate on the "use_case_name" field.
func UseCaseNameLTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameContains applies the Contains predicate on the "use_case_name" field.
func UseCaseNameContains(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameHasPrefix applies the HasPrefix predicate on the "use_case_name" field.
func UseCaseNameHasPrefix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameHasSuffix applies the HasSuffix predicate on the "use_case_name" field.
func UseCaseNameHasSuffix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameEqualFold applies the EqualFold predicate on the "use_case_name" field.
func UseCaseNameEqualFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUseCaseName), v))
	})
}

// UseCaseNameContainsFold applies the ContainsFold predicate on the "use_case_name" field.
func UseCaseNameContainsFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUseCaseName), v))
	})
}

// UseCaseDescriptionEQ applies the EQ predicate on the "use_case_description" field.
func UseCaseDescriptionEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionNEQ applies the NEQ predicate on the "use_case_description" field.
func UseCaseDescriptionNEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionIn applies the In predicate on the "use_case_description" field.
func UseCaseDescriptionIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUseCaseDescription), v...))
	})
}

// UseCaseDescriptionNotIn applies the NotIn predicate on the "use_case_description" field.
func UseCaseDescriptionNotIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUseCaseDescription), v...))
	})
}

// UseCaseDescriptionGT applies the GT predicate on the "use_case_description" field.
func UseCaseDescriptionGT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionGTE applies the GTE predicate on the "use_case_description" field.
func UseCaseDescriptionGTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionLT applies the LT predicate on the "use_case_description" field.
func UseCaseDescriptionLT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionLTE applies the LTE predicate on the "use_case_description" field.
func UseCaseDescriptionLTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionContains applies the Contains predicate on the "use_case_description" field.
func UseCaseDescriptionContains(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionHasPrefix applies the HasPrefix predicate on the "use_case_description" field.
func UseCaseDescriptionHasPrefix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionHasSuffix applies the HasSuffix predicate on the "use_case_description" field.
func UseCaseDescriptionHasSuffix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionEqualFold applies the EqualFold predicate on the "use_case_description" field.
func UseCaseDescriptionEqualFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseDescriptionContainsFold applies the ContainsFold predicate on the "use_case_description" field.
func UseCaseDescriptionContainsFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUseCaseDescription), v))
	})
}

// UseCaseCodeEQ applies the EQ predicate on the "use_case_code" field.
func UseCaseCodeEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeNEQ applies the NEQ predicate on the "use_case_code" field.
func UseCaseCodeNEQ(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeIn applies the In predicate on the "use_case_code" field.
func UseCaseCodeIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUseCaseCode), v...))
	})
}

// UseCaseCodeNotIn applies the NotIn predicate on the "use_case_code" field.
func UseCaseCodeNotIn(vs ...string) predicate.UseCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UseCase(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUseCaseCode), v...))
	})
}

// UseCaseCodeGT applies the GT predicate on the "use_case_code" field.
func UseCaseCodeGT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeGTE applies the GTE predicate on the "use_case_code" field.
func UseCaseCodeGTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeLT applies the LT predicate on the "use_case_code" field.
func UseCaseCodeLT(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeLTE applies the LTE predicate on the "use_case_code" field.
func UseCaseCodeLTE(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeContains applies the Contains predicate on the "use_case_code" field.
func UseCaseCodeContains(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeHasPrefix applies the HasPrefix predicate on the "use_case_code" field.
func UseCaseCodeHasPrefix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeHasSuffix applies the HasSuffix predicate on the "use_case_code" field.
func UseCaseCodeHasSuffix(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeEqualFold applies the EqualFold predicate on the "use_case_code" field.
func UseCaseCodeEqualFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeContainsFold applies the ContainsFold predicate on the "use_case_code" field.
func UseCaseCodeContainsFold(v string) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUseCaseCode), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UseCase) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UseCase) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
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
func Not(p predicate.UseCase) predicate.UseCase {
	return predicate.UseCase(func(s *sql.Selector) {
		p(s.Not())
	})
}
