// Code generated (@generated) by entc, DO NOT EDIT.

package immunocompromisedgroups

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GroupName applies equality check predicate on the "group_name" field. It's identical to GroupNameEQ.
func GroupName(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// GroupNameEQ applies the EQ predicate on the "group_name" field.
func GroupNameEQ(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// GroupNameNEQ applies the NEQ predicate on the "group_name" field.
func GroupNameNEQ(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupName), v))
	})
}

// GroupNameIn applies the In predicate on the "group_name" field.
func GroupNameIn(vs ...string) predicate.ImmunocompromisedGroups {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupName), v...))
	})
}

// GroupNameNotIn applies the NotIn predicate on the "group_name" field.
func GroupNameNotIn(vs ...string) predicate.ImmunocompromisedGroups {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupName), v...))
	})
}

// GroupNameGT applies the GT predicate on the "group_name" field.
func GroupNameGT(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupName), v))
	})
}

// GroupNameGTE applies the GTE predicate on the "group_name" field.
func GroupNameGTE(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupName), v))
	})
}

// GroupNameLT applies the LT predicate on the "group_name" field.
func GroupNameLT(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupName), v))
	})
}

// GroupNameLTE applies the LTE predicate on the "group_name" field.
func GroupNameLTE(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupName), v))
	})
}

// GroupNameContains applies the Contains predicate on the "group_name" field.
func GroupNameContains(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupName), v))
	})
}

// GroupNameHasPrefix applies the HasPrefix predicate on the "group_name" field.
func GroupNameHasPrefix(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupName), v))
	})
}

// GroupNameHasSuffix applies the HasSuffix predicate on the "group_name" field.
func GroupNameHasSuffix(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupName), v))
	})
}

// GroupNameEqualFold applies the EqualFold predicate on the "group_name" field.
func GroupNameEqualFold(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupName), v))
	})
}

// GroupNameContainsFold applies the ContainsFold predicate on the "group_name" field.
func GroupNameContainsFold(v string) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ImmunocompromisedGroups) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ImmunocompromisedGroups) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
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
func Not(p predicate.ImmunocompromisedGroups) predicate.ImmunocompromisedGroups {
	return predicate.ImmunocompromisedGroups(func(s *sql.Selector) {
		p(s.Not())
	})
}
