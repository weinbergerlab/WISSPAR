// Code generated (@generated) by entc, DO NOT EDIT.

package otherevent

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OtherEventTerm applies equality check predicate on the "other_event_term" field. It's identical to OtherEventTermEQ.
func OtherEventTerm(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventOrganSystem applies equality check predicate on the "other_event_organ_system" field. It's identical to OtherEventOrganSystemEQ.
func OtherEventOrganSystem(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventSourceVocabulary applies equality check predicate on the "other_event_source_vocabulary" field. It's identical to OtherEventSourceVocabularyEQ.
func OtherEventSourceVocabulary(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventAssessmentType applies equality check predicate on the "other_event_assessment_type" field. It's identical to OtherEventAssessmentTypeEQ.
func OtherEventAssessmentType(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventNotes applies equality check predicate on the "other_event_notes" field. It's identical to OtherEventNotesEQ.
func OtherEventNotes(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventTermEQ applies the EQ predicate on the "other_event_term" field.
func OtherEventTermEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermNEQ applies the NEQ predicate on the "other_event_term" field.
func OtherEventTermNEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermIn applies the In predicate on the "other_event_term" field.
func OtherEventTermIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventTerm), v...))
	})
}

// OtherEventTermNotIn applies the NotIn predicate on the "other_event_term" field.
func OtherEventTermNotIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventTerm), v...))
	})
}

// OtherEventTermGT applies the GT predicate on the "other_event_term" field.
func OtherEventTermGT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermGTE applies the GTE predicate on the "other_event_term" field.
func OtherEventTermGTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermLT applies the LT predicate on the "other_event_term" field.
func OtherEventTermLT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermLTE applies the LTE predicate on the "other_event_term" field.
func OtherEventTermLTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermContains applies the Contains predicate on the "other_event_term" field.
func OtherEventTermContains(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermHasPrefix applies the HasPrefix predicate on the "other_event_term" field.
func OtherEventTermHasPrefix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermHasSuffix applies the HasSuffix predicate on the "other_event_term" field.
func OtherEventTermHasSuffix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermEqualFold applies the EqualFold predicate on the "other_event_term" field.
func OtherEventTermEqualFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventTermContainsFold applies the ContainsFold predicate on the "other_event_term" field.
func OtherEventTermContainsFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventTerm), v))
	})
}

// OtherEventOrganSystemEQ applies the EQ predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemNEQ applies the NEQ predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemNEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemIn applies the In predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventOrganSystem), v...))
	})
}

// OtherEventOrganSystemNotIn applies the NotIn predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemNotIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventOrganSystem), v...))
	})
}

// OtherEventOrganSystemGT applies the GT predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemGT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemGTE applies the GTE predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemGTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemLT applies the LT predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemLT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemLTE applies the LTE predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemLTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemContains applies the Contains predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemContains(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemHasPrefix applies the HasPrefix predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemHasPrefix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemHasSuffix applies the HasSuffix predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemHasSuffix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemEqualFold applies the EqualFold predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemEqualFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventOrganSystemContainsFold applies the ContainsFold predicate on the "other_event_organ_system" field.
func OtherEventOrganSystemContainsFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventOrganSystem), v))
	})
}

// OtherEventSourceVocabularyEQ applies the EQ predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyNEQ applies the NEQ predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyNEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyIn applies the In predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventSourceVocabulary), v...))
	})
}

// OtherEventSourceVocabularyNotIn applies the NotIn predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyNotIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventSourceVocabulary), v...))
	})
}

// OtherEventSourceVocabularyGT applies the GT predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyGT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyGTE applies the GTE predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyGTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyLT applies the LT predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyLT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyLTE applies the LTE predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyLTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyContains applies the Contains predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyContains(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyHasPrefix applies the HasPrefix predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyHasPrefix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyHasSuffix applies the HasSuffix predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyHasSuffix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyEqualFold applies the EqualFold predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyEqualFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventSourceVocabularyContainsFold applies the ContainsFold predicate on the "other_event_source_vocabulary" field.
func OtherEventSourceVocabularyContainsFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventSourceVocabulary), v))
	})
}

// OtherEventAssessmentTypeEQ applies the EQ predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeNEQ applies the NEQ predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeNEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeIn applies the In predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventAssessmentType), v...))
	})
}

// OtherEventAssessmentTypeNotIn applies the NotIn predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeNotIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventAssessmentType), v...))
	})
}

// OtherEventAssessmentTypeGT applies the GT predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeGT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeGTE applies the GTE predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeGTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeLT applies the LT predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeLT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeLTE applies the LTE predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeLTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeContains applies the Contains predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeContains(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeHasPrefix applies the HasPrefix predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeHasPrefix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeHasSuffix applies the HasSuffix predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeHasSuffix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeEqualFold applies the EqualFold predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeEqualFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventAssessmentTypeContainsFold applies the ContainsFold predicate on the "other_event_assessment_type" field.
func OtherEventAssessmentTypeContainsFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventAssessmentType), v))
	})
}

// OtherEventNotesEQ applies the EQ predicate on the "other_event_notes" field.
func OtherEventNotesEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesNEQ applies the NEQ predicate on the "other_event_notes" field.
func OtherEventNotesNEQ(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesIn applies the In predicate on the "other_event_notes" field.
func OtherEventNotesIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherEventNotes), v...))
	})
}

// OtherEventNotesNotIn applies the NotIn predicate on the "other_event_notes" field.
func OtherEventNotesNotIn(vs ...string) predicate.OtherEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherEvent(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherEventNotes), v...))
	})
}

// OtherEventNotesGT applies the GT predicate on the "other_event_notes" field.
func OtherEventNotesGT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesGTE applies the GTE predicate on the "other_event_notes" field.
func OtherEventNotesGTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesLT applies the LT predicate on the "other_event_notes" field.
func OtherEventNotesLT(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesLTE applies the LTE predicate on the "other_event_notes" field.
func OtherEventNotesLTE(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesContains applies the Contains predicate on the "other_event_notes" field.
func OtherEventNotesContains(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesHasPrefix applies the HasPrefix predicate on the "other_event_notes" field.
func OtherEventNotesHasPrefix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesHasSuffix applies the HasSuffix predicate on the "other_event_notes" field.
func OtherEventNotesHasSuffix(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesEqualFold applies the EqualFold predicate on the "other_event_notes" field.
func OtherEventNotesEqualFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherEventNotes), v))
	})
}

// OtherEventNotesContainsFold applies the ContainsFold predicate on the "other_event_notes" field.
func OtherEventNotesContainsFold(v string) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherEventNotes), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.AdverseEventsModule) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
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

// HasOtherEventStatsList applies the HasEdge predicate on the "other_event_stats_list" edge.
func HasOtherEventStatsList() predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OtherEventStatsListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OtherEventStatsListTable, OtherEventStatsListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOtherEventStatsListWith applies the HasEdge predicate on the "other_event_stats_list" edge with a given conditions (other predicates).
func HasOtherEventStatsListWith(preds ...predicate.OtherEventStats) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OtherEventStatsListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OtherEventStatsListTable, OtherEventStatsListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OtherEvent) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OtherEvent) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
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
func Not(p predicate.OtherEvent) predicate.OtherEvent {
	return predicate.OtherEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}
