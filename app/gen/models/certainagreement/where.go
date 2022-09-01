// Code generated (@generated) by entc, DO NOT EDIT.

package certainagreement

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AgreementPiSponsorEmployee applies equality check predicate on the "agreement_pi_sponsor_employee" field. It's identical to AgreementPiSponsorEmployeeEQ.
func AgreementPiSponsorEmployee(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementRestrictionType applies equality check predicate on the "agreement_restriction_type" field. It's identical to AgreementRestrictionTypeEQ.
func AgreementRestrictionType(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictiveAgreement applies equality check predicate on the "agreement_restrictive_agreement" field. It's identical to AgreementRestrictiveAgreementEQ.
func AgreementRestrictiveAgreement(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementOtherDetails applies equality check predicate on the "agreement_other_details" field. It's identical to AgreementOtherDetailsEQ.
func AgreementOtherDetails(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementPiSponsorEmployeeEQ applies the EQ predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeNEQ applies the NEQ predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeNEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeIn applies the In predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgreementPiSponsorEmployee), v...))
	})
}

// AgreementPiSponsorEmployeeNotIn applies the NotIn predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeNotIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgreementPiSponsorEmployee), v...))
	})
}

// AgreementPiSponsorEmployeeGT applies the GT predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeGT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeGTE applies the GTE predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeGTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeLT applies the LT predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeLT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeLTE applies the LTE predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeLTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeContains applies the Contains predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeContains(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeHasPrefix applies the HasPrefix predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeHasPrefix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeHasSuffix applies the HasSuffix predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeHasSuffix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeEqualFold applies the EqualFold predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeEqualFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementPiSponsorEmployeeContainsFold applies the ContainsFold predicate on the "agreement_pi_sponsor_employee" field.
func AgreementPiSponsorEmployeeContainsFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAgreementPiSponsorEmployee), v))
	})
}

// AgreementRestrictionTypeEQ applies the EQ predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeNEQ applies the NEQ predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeNEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeIn applies the In predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgreementRestrictionType), v...))
	})
}

// AgreementRestrictionTypeNotIn applies the NotIn predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeNotIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgreementRestrictionType), v...))
	})
}

// AgreementRestrictionTypeGT applies the GT predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeGT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeGTE applies the GTE predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeGTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeLT applies the LT predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeLT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeLTE applies the LTE predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeLTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeContains applies the Contains predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeContains(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeHasPrefix applies the HasPrefix predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeHasPrefix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeHasSuffix applies the HasSuffix predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeHasSuffix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeEqualFold applies the EqualFold predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeEqualFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictionTypeContainsFold applies the ContainsFold predicate on the "agreement_restriction_type" field.
func AgreementRestrictionTypeContainsFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAgreementRestrictionType), v))
	})
}

// AgreementRestrictiveAgreementEQ applies the EQ predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementNEQ applies the NEQ predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementNEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementIn applies the In predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgreementRestrictiveAgreement), v...))
	})
}

// AgreementRestrictiveAgreementNotIn applies the NotIn predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementNotIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgreementRestrictiveAgreement), v...))
	})
}

// AgreementRestrictiveAgreementGT applies the GT predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementGT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementGTE applies the GTE predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementGTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementLT applies the LT predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementLT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementLTE applies the LTE predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementLTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementContains applies the Contains predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementContains(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementHasPrefix applies the HasPrefix predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementHasPrefix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementHasSuffix applies the HasSuffix predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementHasSuffix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementEqualFold applies the EqualFold predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementEqualFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementRestrictiveAgreementContainsFold applies the ContainsFold predicate on the "agreement_restrictive_agreement" field.
func AgreementRestrictiveAgreementContainsFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAgreementRestrictiveAgreement), v))
	})
}

// AgreementOtherDetailsEQ applies the EQ predicate on the "agreement_other_details" field.
func AgreementOtherDetailsEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsNEQ applies the NEQ predicate on the "agreement_other_details" field.
func AgreementOtherDetailsNEQ(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsIn applies the In predicate on the "agreement_other_details" field.
func AgreementOtherDetailsIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAgreementOtherDetails), v...))
	})
}

// AgreementOtherDetailsNotIn applies the NotIn predicate on the "agreement_other_details" field.
func AgreementOtherDetailsNotIn(vs ...string) predicate.CertainAgreement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CertainAgreement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAgreementOtherDetails), v...))
	})
}

// AgreementOtherDetailsGT applies the GT predicate on the "agreement_other_details" field.
func AgreementOtherDetailsGT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsGTE applies the GTE predicate on the "agreement_other_details" field.
func AgreementOtherDetailsGTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsLT applies the LT predicate on the "agreement_other_details" field.
func AgreementOtherDetailsLT(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsLTE applies the LTE predicate on the "agreement_other_details" field.
func AgreementOtherDetailsLTE(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsContains applies the Contains predicate on the "agreement_other_details" field.
func AgreementOtherDetailsContains(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsHasPrefix applies the HasPrefix predicate on the "agreement_other_details" field.
func AgreementOtherDetailsHasPrefix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsHasSuffix applies the HasSuffix predicate on the "agreement_other_details" field.
func AgreementOtherDetailsHasSuffix(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsEqualFold applies the EqualFold predicate on the "agreement_other_details" field.
func AgreementOtherDetailsEqualFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAgreementOtherDetails), v))
	})
}

// AgreementOtherDetailsContainsFold applies the ContainsFold predicate on the "agreement_other_details" field.
func AgreementOtherDetailsContainsFold(v string) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAgreementOtherDetails), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.MoreInfoModule) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
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
func And(predicates ...predicate.CertainAgreement) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CertainAgreement) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
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
func Not(p predicate.CertainAgreement) predicate.CertainAgreement {
	return predicate.CertainAgreement(func(s *sql.Selector) {
		p(s.Not())
	})
}
