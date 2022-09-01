// Code generated (@generated) by entc, DO NOT EDIT.

package clinicaltrialmetadata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TagName applies equality check predicate on the "tag_name" field. It's identical to TagNameEQ.
func TagName(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagName), v))
	})
}

// TagValue applies equality check predicate on the "tag_value" field. It's identical to TagValueEQ.
func TagValue(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagValue), v))
	})
}

// UseCaseCode applies equality check predicate on the "use_case_code" field. It's identical to UseCaseCodeEQ.
func UseCaseCode(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseCode), v))
	})
}

// TagNameEQ applies the EQ predicate on the "tag_name" field.
func TagNameEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagName), v))
	})
}

// TagNameNEQ applies the NEQ predicate on the "tag_name" field.
func TagNameNEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagName), v))
	})
}

// TagNameIn applies the In predicate on the "tag_name" field.
func TagNameIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTagName), v...))
	})
}

// TagNameNotIn applies the NotIn predicate on the "tag_name" field.
func TagNameNotIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTagName), v...))
	})
}

// TagNameGT applies the GT predicate on the "tag_name" field.
func TagNameGT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTagName), v))
	})
}

// TagNameGTE applies the GTE predicate on the "tag_name" field.
func TagNameGTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTagName), v))
	})
}

// TagNameLT applies the LT predicate on the "tag_name" field.
func TagNameLT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTagName), v))
	})
}

// TagNameLTE applies the LTE predicate on the "tag_name" field.
func TagNameLTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTagName), v))
	})
}

// TagNameContains applies the Contains predicate on the "tag_name" field.
func TagNameContains(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTagName), v))
	})
}

// TagNameHasPrefix applies the HasPrefix predicate on the "tag_name" field.
func TagNameHasPrefix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTagName), v))
	})
}

// TagNameHasSuffix applies the HasSuffix predicate on the "tag_name" field.
func TagNameHasSuffix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTagName), v))
	})
}

// TagNameEqualFold applies the EqualFold predicate on the "tag_name" field.
func TagNameEqualFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTagName), v))
	})
}

// TagNameContainsFold applies the ContainsFold predicate on the "tag_name" field.
func TagNameContainsFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTagName), v))
	})
}

// TagValueEQ applies the EQ predicate on the "tag_value" field.
func TagValueEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagValue), v))
	})
}

// TagValueNEQ applies the NEQ predicate on the "tag_value" field.
func TagValueNEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagValue), v))
	})
}

// TagValueIn applies the In predicate on the "tag_value" field.
func TagValueIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTagValue), v...))
	})
}

// TagValueNotIn applies the NotIn predicate on the "tag_value" field.
func TagValueNotIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTagValue), v...))
	})
}

// TagValueGT applies the GT predicate on the "tag_value" field.
func TagValueGT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTagValue), v))
	})
}

// TagValueGTE applies the GTE predicate on the "tag_value" field.
func TagValueGTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTagValue), v))
	})
}

// TagValueLT applies the LT predicate on the "tag_value" field.
func TagValueLT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTagValue), v))
	})
}

// TagValueLTE applies the LTE predicate on the "tag_value" field.
func TagValueLTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTagValue), v))
	})
}

// TagValueContains applies the Contains predicate on the "tag_value" field.
func TagValueContains(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTagValue), v))
	})
}

// TagValueHasPrefix applies the HasPrefix predicate on the "tag_value" field.
func TagValueHasPrefix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTagValue), v))
	})
}

// TagValueHasSuffix applies the HasSuffix predicate on the "tag_value" field.
func TagValueHasSuffix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTagValue), v))
	})
}

// TagValueEqualFold applies the EqualFold predicate on the "tag_value" field.
func TagValueEqualFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTagValue), v))
	})
}

// TagValueContainsFold applies the ContainsFold predicate on the "tag_value" field.
func TagValueContainsFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTagValue), v))
	})
}

// UseCaseCodeEQ applies the EQ predicate on the "use_case_code" field.
func UseCaseCodeEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeNEQ applies the NEQ predicate on the "use_case_code" field.
func UseCaseCodeNEQ(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeIn applies the In predicate on the "use_case_code" field.
func UseCaseCodeIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func UseCaseCodeNotIn(vs ...string) predicate.ClinicalTrialMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func UseCaseCodeGT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeGTE applies the GTE predicate on the "use_case_code" field.
func UseCaseCodeGTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeLT applies the LT predicate on the "use_case_code" field.
func UseCaseCodeLT(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeLTE applies the LTE predicate on the "use_case_code" field.
func UseCaseCodeLTE(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeContains applies the Contains predicate on the "use_case_code" field.
func UseCaseCodeContains(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeHasPrefix applies the HasPrefix predicate on the "use_case_code" field.
func UseCaseCodeHasPrefix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeHasSuffix applies the HasSuffix predicate on the "use_case_code" field.
func UseCaseCodeHasSuffix(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeEqualFold applies the EqualFold predicate on the "use_case_code" field.
func UseCaseCodeEqualFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUseCaseCode), v))
	})
}

// UseCaseCodeContainsFold applies the ContainsFold predicate on the "use_case_code" field.
func UseCaseCodeContainsFold(v string) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUseCaseCode), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ClinicalTrial) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func And(predicates ...predicate.ClinicalTrialMetadata) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClinicalTrialMetadata) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
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
func Not(p predicate.ClinicalTrialMetadata) predicate.ClinicalTrialMetadata {
	return predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
		p(s.Not())
	})
}
