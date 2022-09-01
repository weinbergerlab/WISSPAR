// Code generated (@generated) by entc, DO NOT EDIT.

package seriousevent

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
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
func IDGT(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SeriousEventTerm applies equality check predicate on the "serious_event_term" field. It's identical to SeriousEventTermEQ.
func SeriousEventTerm(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventOrganSystem applies equality check predicate on the "serious_event_organ_system" field. It's identical to SeriousEventOrganSystemEQ.
func SeriousEventOrganSystem(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventSourceVocabulary applies equality check predicate on the "serious_event_source_vocabulary" field. It's identical to SeriousEventSourceVocabularyEQ.
func SeriousEventSourceVocabulary(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventAssessmentType applies equality check predicate on the "serious_event_assessment_type" field. It's identical to SeriousEventAssessmentTypeEQ.
func SeriousEventAssessmentType(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventNotes applies equality check predicate on the "serious_event_notes" field. It's identical to SeriousEventNotesEQ.
func SeriousEventNotes(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventTermEQ applies the EQ predicate on the "serious_event_term" field.
func SeriousEventTermEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermNEQ applies the NEQ predicate on the "serious_event_term" field.
func SeriousEventTermNEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermIn applies the In predicate on the "serious_event_term" field.
func SeriousEventTermIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventTerm), v...))
	})
}

// SeriousEventTermNotIn applies the NotIn predicate on the "serious_event_term" field.
func SeriousEventTermNotIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventTerm), v...))
	})
}

// SeriousEventTermGT applies the GT predicate on the "serious_event_term" field.
func SeriousEventTermGT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermGTE applies the GTE predicate on the "serious_event_term" field.
func SeriousEventTermGTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermLT applies the LT predicate on the "serious_event_term" field.
func SeriousEventTermLT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermLTE applies the LTE predicate on the "serious_event_term" field.
func SeriousEventTermLTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermContains applies the Contains predicate on the "serious_event_term" field.
func SeriousEventTermContains(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermHasPrefix applies the HasPrefix predicate on the "serious_event_term" field.
func SeriousEventTermHasPrefix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermHasSuffix applies the HasSuffix predicate on the "serious_event_term" field.
func SeriousEventTermHasSuffix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermEqualFold applies the EqualFold predicate on the "serious_event_term" field.
func SeriousEventTermEqualFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventTermContainsFold applies the ContainsFold predicate on the "serious_event_term" field.
func SeriousEventTermContainsFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventTerm), v))
	})
}

// SeriousEventOrganSystemEQ applies the EQ predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemNEQ applies the NEQ predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemNEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemIn applies the In predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventOrganSystem), v...))
	})
}

// SeriousEventOrganSystemNotIn applies the NotIn predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemNotIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventOrganSystem), v...))
	})
}

// SeriousEventOrganSystemGT applies the GT predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemGT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemGTE applies the GTE predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemGTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemLT applies the LT predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemLT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemLTE applies the LTE predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemLTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemContains applies the Contains predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemContains(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemHasPrefix applies the HasPrefix predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemHasPrefix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemHasSuffix applies the HasSuffix predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemHasSuffix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemEqualFold applies the EqualFold predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemEqualFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventOrganSystemContainsFold applies the ContainsFold predicate on the "serious_event_organ_system" field.
func SeriousEventOrganSystemContainsFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventOrganSystem), v))
	})
}

// SeriousEventSourceVocabularyEQ applies the EQ predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyNEQ applies the NEQ predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyNEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyIn applies the In predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventSourceVocabulary), v...))
	})
}

// SeriousEventSourceVocabularyNotIn applies the NotIn predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyNotIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventSourceVocabulary), v...))
	})
}

// SeriousEventSourceVocabularyGT applies the GT predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyGT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyGTE applies the GTE predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyGTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyLT applies the LT predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyLT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyLTE applies the LTE predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyLTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyContains applies the Contains predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyContains(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyHasPrefix applies the HasPrefix predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyHasPrefix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyHasSuffix applies the HasSuffix predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyHasSuffix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyEqualFold applies the EqualFold predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyEqualFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventSourceVocabularyContainsFold applies the ContainsFold predicate on the "serious_event_source_vocabulary" field.
func SeriousEventSourceVocabularyContainsFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventSourceVocabulary), v))
	})
}

// SeriousEventAssessmentTypeEQ applies the EQ predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeNEQ applies the NEQ predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeNEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeIn applies the In predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventAssessmentType), v...))
	})
}

// SeriousEventAssessmentTypeNotIn applies the NotIn predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeNotIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventAssessmentType), v...))
	})
}

// SeriousEventAssessmentTypeGT applies the GT predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeGT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeGTE applies the GTE predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeGTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeLT applies the LT predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeLT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeLTE applies the LTE predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeLTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeContains applies the Contains predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeContains(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeHasPrefix applies the HasPrefix predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeHasPrefix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeHasSuffix applies the HasSuffix predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeHasSuffix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeEqualFold applies the EqualFold predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeEqualFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventAssessmentTypeContainsFold applies the ContainsFold predicate on the "serious_event_assessment_type" field.
func SeriousEventAssessmentTypeContainsFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventAssessmentType), v))
	})
}

// SeriousEventNotesEQ applies the EQ predicate on the "serious_event_notes" field.
func SeriousEventNotesEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesNEQ applies the NEQ predicate on the "serious_event_notes" field.
func SeriousEventNotesNEQ(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesIn applies the In predicate on the "serious_event_notes" field.
func SeriousEventNotesIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeriousEventNotes), v...))
	})
}

// SeriousEventNotesNotIn applies the NotIn predicate on the "serious_event_notes" field.
func SeriousEventNotesNotIn(vs ...string) predicate.SeriousEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SeriousEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeriousEventNotes), v...))
	})
}

// SeriousEventNotesGT applies the GT predicate on the "serious_event_notes" field.
func SeriousEventNotesGT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesGTE applies the GTE predicate on the "serious_event_notes" field.
func SeriousEventNotesGTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesLT applies the LT predicate on the "serious_event_notes" field.
func SeriousEventNotesLT(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesLTE applies the LTE predicate on the "serious_event_notes" field.
func SeriousEventNotesLTE(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesContains applies the Contains predicate on the "serious_event_notes" field.
func SeriousEventNotesContains(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesHasPrefix applies the HasPrefix predicate on the "serious_event_notes" field.
func SeriousEventNotesHasPrefix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesHasSuffix applies the HasSuffix predicate on the "serious_event_notes" field.
func SeriousEventNotesHasSuffix(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesEqualFold applies the EqualFold predicate on the "serious_event_notes" field.
func SeriousEventNotesEqualFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeriousEventNotes), v))
	})
}

// SeriousEventNotesContainsFold applies the ContainsFold predicate on the "serious_event_notes" field.
func SeriousEventNotesContainsFold(v string) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeriousEventNotes), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.AdverseEventsModule) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
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

// HasSeriousEventStatsList applies the HasEdge predicate on the "serious_event_stats_list" edge.
func HasSeriousEventStatsList() predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriousEventStatsListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeriousEventStatsListTable, SeriousEventStatsListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeriousEventStatsListWith applies the HasEdge predicate on the "serious_event_stats_list" edge with a given conditions (other predicates).
func HasSeriousEventStatsListWith(preds ...predicate.SeriousEventStats) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriousEventStatsListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeriousEventStatsListTable, SeriousEventStatsListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SeriousEvent) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SeriousEvent) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
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
func Not(p predicate.SeriousEvent) predicate.SeriousEvent {
	return predicate.SeriousEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}
