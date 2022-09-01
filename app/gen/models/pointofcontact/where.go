// Code generated (@generated) by entc, DO NOT EDIT.

package pointofcontact

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
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
func IDGT(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PointOfContactTitle applies equality check predicate on the "point_of_contact_title" field. It's identical to PointOfContactTitleEQ.
func PointOfContactTitle(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactOrganization applies equality check predicate on the "point_of_contact_organization" field. It's identical to PointOfContactOrganizationEQ.
func PointOfContactOrganization(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactEmail applies equality check predicate on the "point_of_contact_email" field. It's identical to PointOfContactEmailEQ.
func PointOfContactEmail(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactPhone applies equality check predicate on the "point_of_contact_phone" field. It's identical to PointOfContactPhoneEQ.
func PointOfContactPhone(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneExt applies equality check predicate on the "point_of_contact_phone_ext" field. It's identical to PointOfContactPhoneExtEQ.
func PointOfContactPhoneExt(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactTitleEQ applies the EQ predicate on the "point_of_contact_title" field.
func PointOfContactTitleEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleNEQ applies the NEQ predicate on the "point_of_contact_title" field.
func PointOfContactTitleNEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleIn applies the In predicate on the "point_of_contact_title" field.
func PointOfContactTitleIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPointOfContactTitle), v...))
	})
}

// PointOfContactTitleNotIn applies the NotIn predicate on the "point_of_contact_title" field.
func PointOfContactTitleNotIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPointOfContactTitle), v...))
	})
}

// PointOfContactTitleGT applies the GT predicate on the "point_of_contact_title" field.
func PointOfContactTitleGT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleGTE applies the GTE predicate on the "point_of_contact_title" field.
func PointOfContactTitleGTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleLT applies the LT predicate on the "point_of_contact_title" field.
func PointOfContactTitleLT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleLTE applies the LTE predicate on the "point_of_contact_title" field.
func PointOfContactTitleLTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleContains applies the Contains predicate on the "point_of_contact_title" field.
func PointOfContactTitleContains(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleHasPrefix applies the HasPrefix predicate on the "point_of_contact_title" field.
func PointOfContactTitleHasPrefix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleHasSuffix applies the HasSuffix predicate on the "point_of_contact_title" field.
func PointOfContactTitleHasSuffix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleEqualFold applies the EqualFold predicate on the "point_of_contact_title" field.
func PointOfContactTitleEqualFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactTitleContainsFold applies the ContainsFold predicate on the "point_of_contact_title" field.
func PointOfContactTitleContainsFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointOfContactTitle), v))
	})
}

// PointOfContactOrganizationEQ applies the EQ predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationNEQ applies the NEQ predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationNEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationIn applies the In predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPointOfContactOrganization), v...))
	})
}

// PointOfContactOrganizationNotIn applies the NotIn predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationNotIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPointOfContactOrganization), v...))
	})
}

// PointOfContactOrganizationGT applies the GT predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationGT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationGTE applies the GTE predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationGTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationLT applies the LT predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationLT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationLTE applies the LTE predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationLTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationContains applies the Contains predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationContains(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationHasPrefix applies the HasPrefix predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationHasPrefix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationHasSuffix applies the HasSuffix predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationHasSuffix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationEqualFold applies the EqualFold predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationEqualFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactOrganizationContainsFold applies the ContainsFold predicate on the "point_of_contact_organization" field.
func PointOfContactOrganizationContainsFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointOfContactOrganization), v))
	})
}

// PointOfContactEmailEQ applies the EQ predicate on the "point_of_contact_email" field.
func PointOfContactEmailEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailNEQ applies the NEQ predicate on the "point_of_contact_email" field.
func PointOfContactEmailNEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailIn applies the In predicate on the "point_of_contact_email" field.
func PointOfContactEmailIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPointOfContactEmail), v...))
	})
}

// PointOfContactEmailNotIn applies the NotIn predicate on the "point_of_contact_email" field.
func PointOfContactEmailNotIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPointOfContactEmail), v...))
	})
}

// PointOfContactEmailGT applies the GT predicate on the "point_of_contact_email" field.
func PointOfContactEmailGT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailGTE applies the GTE predicate on the "point_of_contact_email" field.
func PointOfContactEmailGTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailLT applies the LT predicate on the "point_of_contact_email" field.
func PointOfContactEmailLT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailLTE applies the LTE predicate on the "point_of_contact_email" field.
func PointOfContactEmailLTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailContains applies the Contains predicate on the "point_of_contact_email" field.
func PointOfContactEmailContains(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailHasPrefix applies the HasPrefix predicate on the "point_of_contact_email" field.
func PointOfContactEmailHasPrefix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailHasSuffix applies the HasSuffix predicate on the "point_of_contact_email" field.
func PointOfContactEmailHasSuffix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailEqualFold applies the EqualFold predicate on the "point_of_contact_email" field.
func PointOfContactEmailEqualFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactEmailContainsFold applies the ContainsFold predicate on the "point_of_contact_email" field.
func PointOfContactEmailContainsFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointOfContactEmail), v))
	})
}

// PointOfContactPhoneEQ applies the EQ predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneNEQ applies the NEQ predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneNEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneIn applies the In predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPointOfContactPhone), v...))
	})
}

// PointOfContactPhoneNotIn applies the NotIn predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneNotIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPointOfContactPhone), v...))
	})
}

// PointOfContactPhoneGT applies the GT predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneGT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneGTE applies the GTE predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneGTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneLT applies the LT predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneLT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneLTE applies the LTE predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneLTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneContains applies the Contains predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneContains(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneHasPrefix applies the HasPrefix predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneHasPrefix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneHasSuffix applies the HasSuffix predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneHasSuffix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneEqualFold applies the EqualFold predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneEqualFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneContainsFold applies the ContainsFold predicate on the "point_of_contact_phone" field.
func PointOfContactPhoneContainsFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointOfContactPhone), v))
	})
}

// PointOfContactPhoneExtEQ applies the EQ predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtNEQ applies the NEQ predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtNEQ(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtIn applies the In predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPointOfContactPhoneExt), v...))
	})
}

// PointOfContactPhoneExtNotIn applies the NotIn predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtNotIn(vs ...string) predicate.PointOfContact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PointOfContact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPointOfContactPhoneExt), v...))
	})
}

// PointOfContactPhoneExtGT applies the GT predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtGT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtGTE applies the GTE predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtGTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtLT applies the LT predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtLT(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtLTE applies the LTE predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtLTE(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtContains applies the Contains predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtContains(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtHasPrefix applies the HasPrefix predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtHasPrefix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtHasSuffix applies the HasSuffix predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtHasSuffix(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtEqualFold applies the EqualFold predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtEqualFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// PointOfContactPhoneExtContainsFold applies the ContainsFold predicate on the "point_of_contact_phone_ext" field.
func PointOfContactPhoneExtContainsFold(v string) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointOfContactPhoneExt), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.MoreInfoModule) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PointOfContact) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PointOfContact) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
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
func Not(p predicate.PointOfContact) predicate.PointOfContact {
	return predicate.PointOfContact(func(s *sql.Selector) {
		p(s.Not())
	})
}
