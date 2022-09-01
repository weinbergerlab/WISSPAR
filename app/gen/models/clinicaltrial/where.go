// Code generated (@generated) by entc, DO NOT EDIT.

package clinicaltrial

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StudyName applies equality check predicate on the "study_name" field. It's identical to StudyNameEQ.
func StudyName(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyName), v))
	})
}

// Sponsor applies equality check predicate on the "sponsor" field. It's identical to SponsorEQ.
func Sponsor(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSponsor), v))
	})
}

// ResponsibleParty applies equality check predicate on the "responsible_party" field. It's identical to ResponsiblePartyEQ.
func ResponsibleParty(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsibleParty), v))
	})
}

// StudyType applies equality check predicate on the "study_type" field. It's identical to StudyTypeEQ.
func StudyType(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyType), v))
	})
}

// Phase applies equality check predicate on the "phase" field. It's identical to PhaseEQ.
func Phase(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhase), v))
	})
}

// ActualEnrollment applies equality check predicate on the "actual_enrollment" field. It's identical to ActualEnrollmentEQ.
func ActualEnrollment(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualEnrollment), v))
	})
}

// Allocation applies equality check predicate on the "allocation" field. It's identical to AllocationEQ.
func Allocation(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocation), v))
	})
}

// InterventionModel applies equality check predicate on the "intervention_model" field. It's identical to InterventionModelEQ.
func InterventionModel(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterventionModel), v))
	})
}

// Masking applies equality check predicate on the "masking" field. It's identical to MaskingEQ.
func Masking(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMasking), v))
	})
}

// PrimaryPurpose applies equality check predicate on the "primary_purpose" field. It's identical to PrimaryPurposeEQ.
func PrimaryPurpose(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryPurpose), v))
	})
}

// OfficialTitle applies equality check predicate on the "official_title" field. It's identical to OfficialTitleEQ.
func OfficialTitle(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficialTitle), v))
	})
}

// ActualStudyStartDate applies equality check predicate on the "actual_study_start_date" field. It's identical to ActualStudyStartDateEQ.
func ActualStudyStartDate(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualPrimaryCompletionDate applies equality check predicate on the "actual_primary_completion_date" field. It's identical to ActualPrimaryCompletionDateEQ.
func ActualPrimaryCompletionDate(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualStudyCompletionDate applies equality check predicate on the "actual_study_completion_date" field. It's identical to ActualStudyCompletionDateEQ.
func ActualStudyCompletionDate(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualStudyCompletionDate), v))
	})
}

// StudyID applies equality check predicate on the "study_id" field. It's identical to StudyIDEQ.
func StudyID(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyID), v))
	})
}

// StudyNameEQ applies the EQ predicate on the "study_name" field.
func StudyNameEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyName), v))
	})
}

// StudyNameNEQ applies the NEQ predicate on the "study_name" field.
func StudyNameNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudyName), v))
	})
}

// StudyNameIn applies the In predicate on the "study_name" field.
func StudyNameIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudyName), v...))
	})
}

// StudyNameNotIn applies the NotIn predicate on the "study_name" field.
func StudyNameNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudyName), v...))
	})
}

// StudyNameGT applies the GT predicate on the "study_name" field.
func StudyNameGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStudyName), v))
	})
}

// StudyNameGTE applies the GTE predicate on the "study_name" field.
func StudyNameGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStudyName), v))
	})
}

// StudyNameLT applies the LT predicate on the "study_name" field.
func StudyNameLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStudyName), v))
	})
}

// StudyNameLTE applies the LTE predicate on the "study_name" field.
func StudyNameLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStudyName), v))
	})
}

// StudyNameContains applies the Contains predicate on the "study_name" field.
func StudyNameContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStudyName), v))
	})
}

// StudyNameHasPrefix applies the HasPrefix predicate on the "study_name" field.
func StudyNameHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStudyName), v))
	})
}

// StudyNameHasSuffix applies the HasSuffix predicate on the "study_name" field.
func StudyNameHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStudyName), v))
	})
}

// StudyNameEqualFold applies the EqualFold predicate on the "study_name" field.
func StudyNameEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStudyName), v))
	})
}

// StudyNameContainsFold applies the ContainsFold predicate on the "study_name" field.
func StudyNameContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStudyName), v))
	})
}

// SponsorEQ applies the EQ predicate on the "sponsor" field.
func SponsorEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSponsor), v))
	})
}

// SponsorNEQ applies the NEQ predicate on the "sponsor" field.
func SponsorNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSponsor), v))
	})
}

// SponsorIn applies the In predicate on the "sponsor" field.
func SponsorIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSponsor), v...))
	})
}

// SponsorNotIn applies the NotIn predicate on the "sponsor" field.
func SponsorNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSponsor), v...))
	})
}

// SponsorGT applies the GT predicate on the "sponsor" field.
func SponsorGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSponsor), v))
	})
}

// SponsorGTE applies the GTE predicate on the "sponsor" field.
func SponsorGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSponsor), v))
	})
}

// SponsorLT applies the LT predicate on the "sponsor" field.
func SponsorLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSponsor), v))
	})
}

// SponsorLTE applies the LTE predicate on the "sponsor" field.
func SponsorLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSponsor), v))
	})
}

// SponsorContains applies the Contains predicate on the "sponsor" field.
func SponsorContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSponsor), v))
	})
}

// SponsorHasPrefix applies the HasPrefix predicate on the "sponsor" field.
func SponsorHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSponsor), v))
	})
}

// SponsorHasSuffix applies the HasSuffix predicate on the "sponsor" field.
func SponsorHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSponsor), v))
	})
}

// SponsorEqualFold applies the EqualFold predicate on the "sponsor" field.
func SponsorEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSponsor), v))
	})
}

// SponsorContainsFold applies the ContainsFold predicate on the "sponsor" field.
func SponsorContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSponsor), v))
	})
}

// ResponsiblePartyEQ applies the EQ predicate on the "responsible_party" field.
func ResponsiblePartyEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyNEQ applies the NEQ predicate on the "responsible_party" field.
func ResponsiblePartyNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyIn applies the In predicate on the "responsible_party" field.
func ResponsiblePartyIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResponsibleParty), v...))
	})
}

// ResponsiblePartyNotIn applies the NotIn predicate on the "responsible_party" field.
func ResponsiblePartyNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResponsibleParty), v...))
	})
}

// ResponsiblePartyGT applies the GT predicate on the "responsible_party" field.
func ResponsiblePartyGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyGTE applies the GTE predicate on the "responsible_party" field.
func ResponsiblePartyGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyLT applies the LT predicate on the "responsible_party" field.
func ResponsiblePartyLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyLTE applies the LTE predicate on the "responsible_party" field.
func ResponsiblePartyLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyContains applies the Contains predicate on the "responsible_party" field.
func ResponsiblePartyContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyHasPrefix applies the HasPrefix predicate on the "responsible_party" field.
func ResponsiblePartyHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyHasSuffix applies the HasSuffix predicate on the "responsible_party" field.
func ResponsiblePartyHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyEqualFold applies the EqualFold predicate on the "responsible_party" field.
func ResponsiblePartyEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResponsibleParty), v))
	})
}

// ResponsiblePartyContainsFold applies the ContainsFold predicate on the "responsible_party" field.
func ResponsiblePartyContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResponsibleParty), v))
	})
}

// StudyTypeEQ applies the EQ predicate on the "study_type" field.
func StudyTypeEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyType), v))
	})
}

// StudyTypeNEQ applies the NEQ predicate on the "study_type" field.
func StudyTypeNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudyType), v))
	})
}

// StudyTypeIn applies the In predicate on the "study_type" field.
func StudyTypeIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudyType), v...))
	})
}

// StudyTypeNotIn applies the NotIn predicate on the "study_type" field.
func StudyTypeNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudyType), v...))
	})
}

// StudyTypeGT applies the GT predicate on the "study_type" field.
func StudyTypeGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStudyType), v))
	})
}

// StudyTypeGTE applies the GTE predicate on the "study_type" field.
func StudyTypeGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStudyType), v))
	})
}

// StudyTypeLT applies the LT predicate on the "study_type" field.
func StudyTypeLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStudyType), v))
	})
}

// StudyTypeLTE applies the LTE predicate on the "study_type" field.
func StudyTypeLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStudyType), v))
	})
}

// StudyTypeContains applies the Contains predicate on the "study_type" field.
func StudyTypeContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStudyType), v))
	})
}

// StudyTypeHasPrefix applies the HasPrefix predicate on the "study_type" field.
func StudyTypeHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStudyType), v))
	})
}

// StudyTypeHasSuffix applies the HasSuffix predicate on the "study_type" field.
func StudyTypeHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStudyType), v))
	})
}

// StudyTypeEqualFold applies the EqualFold predicate on the "study_type" field.
func StudyTypeEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStudyType), v))
	})
}

// StudyTypeContainsFold applies the ContainsFold predicate on the "study_type" field.
func StudyTypeContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStudyType), v))
	})
}

// PhaseEQ applies the EQ predicate on the "phase" field.
func PhaseEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhase), v))
	})
}

// PhaseNEQ applies the NEQ predicate on the "phase" field.
func PhaseNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhase), v))
	})
}

// PhaseIn applies the In predicate on the "phase" field.
func PhaseIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhase), v...))
	})
}

// PhaseNotIn applies the NotIn predicate on the "phase" field.
func PhaseNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhase), v...))
	})
}

// PhaseGT applies the GT predicate on the "phase" field.
func PhaseGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhase), v))
	})
}

// PhaseGTE applies the GTE predicate on the "phase" field.
func PhaseGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhase), v))
	})
}

// PhaseLT applies the LT predicate on the "phase" field.
func PhaseLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhase), v))
	})
}

// PhaseLTE applies the LTE predicate on the "phase" field.
func PhaseLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhase), v))
	})
}

// PhaseContains applies the Contains predicate on the "phase" field.
func PhaseContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhase), v))
	})
}

// PhaseHasPrefix applies the HasPrefix predicate on the "phase" field.
func PhaseHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhase), v))
	})
}

// PhaseHasSuffix applies the HasSuffix predicate on the "phase" field.
func PhaseHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhase), v))
	})
}

// PhaseEqualFold applies the EqualFold predicate on the "phase" field.
func PhaseEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhase), v))
	})
}

// PhaseContainsFold applies the ContainsFold predicate on the "phase" field.
func PhaseContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhase), v))
	})
}

// ActualEnrollmentEQ applies the EQ predicate on the "actual_enrollment" field.
func ActualEnrollmentEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentNEQ applies the NEQ predicate on the "actual_enrollment" field.
func ActualEnrollmentNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentIn applies the In predicate on the "actual_enrollment" field.
func ActualEnrollmentIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActualEnrollment), v...))
	})
}

// ActualEnrollmentNotIn applies the NotIn predicate on the "actual_enrollment" field.
func ActualEnrollmentNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActualEnrollment), v...))
	})
}

// ActualEnrollmentGT applies the GT predicate on the "actual_enrollment" field.
func ActualEnrollmentGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentGTE applies the GTE predicate on the "actual_enrollment" field.
func ActualEnrollmentGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentLT applies the LT predicate on the "actual_enrollment" field.
func ActualEnrollmentLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentLTE applies the LTE predicate on the "actual_enrollment" field.
func ActualEnrollmentLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentContains applies the Contains predicate on the "actual_enrollment" field.
func ActualEnrollmentContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentHasPrefix applies the HasPrefix predicate on the "actual_enrollment" field.
func ActualEnrollmentHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentHasSuffix applies the HasSuffix predicate on the "actual_enrollment" field.
func ActualEnrollmentHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentEqualFold applies the EqualFold predicate on the "actual_enrollment" field.
func ActualEnrollmentEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActualEnrollment), v))
	})
}

// ActualEnrollmentContainsFold applies the ContainsFold predicate on the "actual_enrollment" field.
func ActualEnrollmentContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActualEnrollment), v))
	})
}

// AllocationEQ applies the EQ predicate on the "allocation" field.
func AllocationEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocation), v))
	})
}

// AllocationNEQ applies the NEQ predicate on the "allocation" field.
func AllocationNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllocation), v))
	})
}

// AllocationIn applies the In predicate on the "allocation" field.
func AllocationIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAllocation), v...))
	})
}

// AllocationNotIn applies the NotIn predicate on the "allocation" field.
func AllocationNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAllocation), v...))
	})
}

// AllocationGT applies the GT predicate on the "allocation" field.
func AllocationGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllocation), v))
	})
}

// AllocationGTE applies the GTE predicate on the "allocation" field.
func AllocationGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllocation), v))
	})
}

// AllocationLT applies the LT predicate on the "allocation" field.
func AllocationLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllocation), v))
	})
}

// AllocationLTE applies the LTE predicate on the "allocation" field.
func AllocationLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllocation), v))
	})
}

// AllocationContains applies the Contains predicate on the "allocation" field.
func AllocationContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAllocation), v))
	})
}

// AllocationHasPrefix applies the HasPrefix predicate on the "allocation" field.
func AllocationHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAllocation), v))
	})
}

// AllocationHasSuffix applies the HasSuffix predicate on the "allocation" field.
func AllocationHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAllocation), v))
	})
}

// AllocationEqualFold applies the EqualFold predicate on the "allocation" field.
func AllocationEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAllocation), v))
	})
}

// AllocationContainsFold applies the ContainsFold predicate on the "allocation" field.
func AllocationContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAllocation), v))
	})
}

// InterventionModelEQ applies the EQ predicate on the "intervention_model" field.
func InterventionModelEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelNEQ applies the NEQ predicate on the "intervention_model" field.
func InterventionModelNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelIn applies the In predicate on the "intervention_model" field.
func InterventionModelIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInterventionModel), v...))
	})
}

// InterventionModelNotIn applies the NotIn predicate on the "intervention_model" field.
func InterventionModelNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInterventionModel), v...))
	})
}

// InterventionModelGT applies the GT predicate on the "intervention_model" field.
func InterventionModelGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelGTE applies the GTE predicate on the "intervention_model" field.
func InterventionModelGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelLT applies the LT predicate on the "intervention_model" field.
func InterventionModelLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelLTE applies the LTE predicate on the "intervention_model" field.
func InterventionModelLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelContains applies the Contains predicate on the "intervention_model" field.
func InterventionModelContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelHasPrefix applies the HasPrefix predicate on the "intervention_model" field.
func InterventionModelHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelHasSuffix applies the HasSuffix predicate on the "intervention_model" field.
func InterventionModelHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelEqualFold applies the EqualFold predicate on the "intervention_model" field.
func InterventionModelEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInterventionModel), v))
	})
}

// InterventionModelContainsFold applies the ContainsFold predicate on the "intervention_model" field.
func InterventionModelContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInterventionModel), v))
	})
}

// MaskingEQ applies the EQ predicate on the "masking" field.
func MaskingEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMasking), v))
	})
}

// MaskingNEQ applies the NEQ predicate on the "masking" field.
func MaskingNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMasking), v))
	})
}

// MaskingIn applies the In predicate on the "masking" field.
func MaskingIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMasking), v...))
	})
}

// MaskingNotIn applies the NotIn predicate on the "masking" field.
func MaskingNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMasking), v...))
	})
}

// MaskingGT applies the GT predicate on the "masking" field.
func MaskingGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMasking), v))
	})
}

// MaskingGTE applies the GTE predicate on the "masking" field.
func MaskingGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMasking), v))
	})
}

// MaskingLT applies the LT predicate on the "masking" field.
func MaskingLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMasking), v))
	})
}

// MaskingLTE applies the LTE predicate on the "masking" field.
func MaskingLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMasking), v))
	})
}

// MaskingContains applies the Contains predicate on the "masking" field.
func MaskingContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMasking), v))
	})
}

// MaskingHasPrefix applies the HasPrefix predicate on the "masking" field.
func MaskingHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMasking), v))
	})
}

// MaskingHasSuffix applies the HasSuffix predicate on the "masking" field.
func MaskingHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMasking), v))
	})
}

// MaskingEqualFold applies the EqualFold predicate on the "masking" field.
func MaskingEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMasking), v))
	})
}

// MaskingContainsFold applies the ContainsFold predicate on the "masking" field.
func MaskingContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMasking), v))
	})
}

// PrimaryPurposeEQ applies the EQ predicate on the "primary_purpose" field.
func PrimaryPurposeEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeNEQ applies the NEQ predicate on the "primary_purpose" field.
func PrimaryPurposeNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeIn applies the In predicate on the "primary_purpose" field.
func PrimaryPurposeIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrimaryPurpose), v...))
	})
}

// PrimaryPurposeNotIn applies the NotIn predicate on the "primary_purpose" field.
func PrimaryPurposeNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrimaryPurpose), v...))
	})
}

// PrimaryPurposeGT applies the GT predicate on the "primary_purpose" field.
func PrimaryPurposeGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeGTE applies the GTE predicate on the "primary_purpose" field.
func PrimaryPurposeGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeLT applies the LT predicate on the "primary_purpose" field.
func PrimaryPurposeLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeLTE applies the LTE predicate on the "primary_purpose" field.
func PrimaryPurposeLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeContains applies the Contains predicate on the "primary_purpose" field.
func PrimaryPurposeContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeHasPrefix applies the HasPrefix predicate on the "primary_purpose" field.
func PrimaryPurposeHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeHasSuffix applies the HasSuffix predicate on the "primary_purpose" field.
func PrimaryPurposeHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeEqualFold applies the EqualFold predicate on the "primary_purpose" field.
func PrimaryPurposeEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrimaryPurpose), v))
	})
}

// PrimaryPurposeContainsFold applies the ContainsFold predicate on the "primary_purpose" field.
func PrimaryPurposeContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrimaryPurpose), v))
	})
}

// OfficialTitleEQ applies the EQ predicate on the "official_title" field.
func OfficialTitleEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleNEQ applies the NEQ predicate on the "official_title" field.
func OfficialTitleNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleIn applies the In predicate on the "official_title" field.
func OfficialTitleIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOfficialTitle), v...))
	})
}

// OfficialTitleNotIn applies the NotIn predicate on the "official_title" field.
func OfficialTitleNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOfficialTitle), v...))
	})
}

// OfficialTitleGT applies the GT predicate on the "official_title" field.
func OfficialTitleGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleGTE applies the GTE predicate on the "official_title" field.
func OfficialTitleGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleLT applies the LT predicate on the "official_title" field.
func OfficialTitleLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleLTE applies the LTE predicate on the "official_title" field.
func OfficialTitleLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleContains applies the Contains predicate on the "official_title" field.
func OfficialTitleContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleHasPrefix applies the HasPrefix predicate on the "official_title" field.
func OfficialTitleHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleHasSuffix applies the HasSuffix predicate on the "official_title" field.
func OfficialTitleHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleEqualFold applies the EqualFold predicate on the "official_title" field.
func OfficialTitleEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOfficialTitle), v))
	})
}

// OfficialTitleContainsFold applies the ContainsFold predicate on the "official_title" field.
func OfficialTitleContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOfficialTitle), v))
	})
}

// ActualStudyStartDateEQ applies the EQ predicate on the "actual_study_start_date" field.
func ActualStudyStartDateEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualStudyStartDateNEQ applies the NEQ predicate on the "actual_study_start_date" field.
func ActualStudyStartDateNEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualStudyStartDateIn applies the In predicate on the "actual_study_start_date" field.
func ActualStudyStartDateIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActualStudyStartDate), v...))
	})
}

// ActualStudyStartDateNotIn applies the NotIn predicate on the "actual_study_start_date" field.
func ActualStudyStartDateNotIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActualStudyStartDate), v...))
	})
}

// ActualStudyStartDateGT applies the GT predicate on the "actual_study_start_date" field.
func ActualStudyStartDateGT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualStudyStartDateGTE applies the GTE predicate on the "actual_study_start_date" field.
func ActualStudyStartDateGTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualStudyStartDateLT applies the LT predicate on the "actual_study_start_date" field.
func ActualStudyStartDateLT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualStudyStartDateLTE applies the LTE predicate on the "actual_study_start_date" field.
func ActualStudyStartDateLTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActualStudyStartDate), v))
	})
}

// ActualPrimaryCompletionDateEQ applies the EQ predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualPrimaryCompletionDateNEQ applies the NEQ predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateNEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualPrimaryCompletionDateIn applies the In predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActualPrimaryCompletionDate), v...))
	})
}

// ActualPrimaryCompletionDateNotIn applies the NotIn predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateNotIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActualPrimaryCompletionDate), v...))
	})
}

// ActualPrimaryCompletionDateGT applies the GT predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateGT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualPrimaryCompletionDateGTE applies the GTE predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateGTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualPrimaryCompletionDateLT applies the LT predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateLT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualPrimaryCompletionDateLTE applies the LTE predicate on the "actual_primary_completion_date" field.
func ActualPrimaryCompletionDateLTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActualPrimaryCompletionDate), v))
	})
}

// ActualStudyCompletionDateEQ applies the EQ predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActualStudyCompletionDate), v))
	})
}

// ActualStudyCompletionDateNEQ applies the NEQ predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateNEQ(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActualStudyCompletionDate), v))
	})
}

// ActualStudyCompletionDateIn applies the In predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActualStudyCompletionDate), v...))
	})
}

// ActualStudyCompletionDateNotIn applies the NotIn predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateNotIn(vs ...time.Time) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActualStudyCompletionDate), v...))
	})
}

// ActualStudyCompletionDateGT applies the GT predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateGT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActualStudyCompletionDate), v))
	})
}

// ActualStudyCompletionDateGTE applies the GTE predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateGTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActualStudyCompletionDate), v))
	})
}

// ActualStudyCompletionDateLT applies the LT predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateLT(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActualStudyCompletionDate), v))
	})
}

// ActualStudyCompletionDateLTE applies the LTE predicate on the "actual_study_completion_date" field.
func ActualStudyCompletionDateLTE(v time.Time) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActualStudyCompletionDate), v))
	})
}

// StudyIDEQ applies the EQ predicate on the "study_id" field.
func StudyIDEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudyID), v))
	})
}

// StudyIDNEQ applies the NEQ predicate on the "study_id" field.
func StudyIDNEQ(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudyID), v))
	})
}

// StudyIDIn applies the In predicate on the "study_id" field.
func StudyIDIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudyID), v...))
	})
}

// StudyIDNotIn applies the NotIn predicate on the "study_id" field.
func StudyIDNotIn(vs ...string) predicate.ClinicalTrial {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudyID), v...))
	})
}

// StudyIDGT applies the GT predicate on the "study_id" field.
func StudyIDGT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStudyID), v))
	})
}

// StudyIDGTE applies the GTE predicate on the "study_id" field.
func StudyIDGTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStudyID), v))
	})
}

// StudyIDLT applies the LT predicate on the "study_id" field.
func StudyIDLT(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStudyID), v))
	})
}

// StudyIDLTE applies the LTE predicate on the "study_id" field.
func StudyIDLTE(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStudyID), v))
	})
}

// StudyIDContains applies the Contains predicate on the "study_id" field.
func StudyIDContains(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStudyID), v))
	})
}

// StudyIDHasPrefix applies the HasPrefix predicate on the "study_id" field.
func StudyIDHasPrefix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStudyID), v))
	})
}

// StudyIDHasSuffix applies the HasSuffix predicate on the "study_id" field.
func StudyIDHasSuffix(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStudyID), v))
	})
}

// StudyIDEqualFold applies the EqualFold predicate on the "study_id" field.
func StudyIDEqualFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStudyID), v))
	})
}

// StudyIDContainsFold applies the ContainsFold predicate on the "study_id" field.
func StudyIDContainsFold(v string) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStudyID), v))
	})
}

// HasResultsDefinition applies the HasEdge predicate on the "results_definition" edge.
func HasResultsDefinition() predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResultsDefinitionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ResultsDefinitionTable, ResultsDefinitionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResultsDefinitionWith applies the HasEdge predicate on the "results_definition" edge with a given conditions (other predicates).
func HasResultsDefinitionWith(preds ...predicate.ResultsDefinition) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResultsDefinitionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ResultsDefinitionTable, ResultsDefinitionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudyLocations applies the HasEdge predicate on the "study_locations" edge.
func HasStudyLocations() predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyLocationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudyLocationsTable, StudyLocationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudyLocationsWith applies the HasEdge predicate on the "study_locations" edge with a given conditions (other predicates).
func HasStudyLocationsWith(preds ...predicate.StudyLocation) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyLocationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudyLocationsTable, StudyLocationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudyEligibility applies the HasEdge predicate on the "study_eligibility" edge.
func HasStudyEligibility() predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyEligibilityTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudyEligibilityTable, StudyEligibilityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudyEligibilityWith applies the HasEdge predicate on the "study_eligibility" edge with a given conditions (other predicates).
func HasStudyEligibilityWith(preds ...predicate.StudyEligibility) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyEligibilityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudyEligibilityTable, StudyEligibilityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.ClinicalTrialMetadata) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClinicalTrial) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClinicalTrial) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
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
func Not(p predicate.ClinicalTrial) predicate.ClinicalTrial {
	return predicate.ClinicalTrial(func(s *sql.Selector) {
		p(s.Not())
	})
}
