// Code generated (@generated) by entc, DO NOT EDIT.

package baselinecharacteristicsmodule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BaselinePopulationDescription applies equality check predicate on the "baseline_population_description" field. It's identical to BaselinePopulationDescriptionEQ.
func BaselinePopulationDescription(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselineTypeUnitsAnalyzed applies equality check predicate on the "baseline_type_units_analyzed" field. It's identical to BaselineTypeUnitsAnalyzedEQ.
func BaselineTypeUnitsAnalyzed(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselinePopulationDescriptionEQ applies the EQ predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionEQ(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionNEQ applies the NEQ predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionNEQ(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionIn applies the In predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionIn(vs ...string) predicate.BaselineCharacteristicsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselinePopulationDescription), v...))
	})
}

// BaselinePopulationDescriptionNotIn applies the NotIn predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionNotIn(vs ...string) predicate.BaselineCharacteristicsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselinePopulationDescription), v...))
	})
}

// BaselinePopulationDescriptionGT applies the GT predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionGT(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionGTE applies the GTE predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionGTE(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionLT applies the LT predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionLT(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionLTE applies the LTE predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionLTE(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionContains applies the Contains predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionContains(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionHasPrefix applies the HasPrefix predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionHasPrefix(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionHasSuffix applies the HasSuffix predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionHasSuffix(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionEqualFold applies the EqualFold predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionEqualFold(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselinePopulationDescriptionContainsFold applies the ContainsFold predicate on the "baseline_population_description" field.
func BaselinePopulationDescriptionContainsFold(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselinePopulationDescription), v))
	})
}

// BaselineTypeUnitsAnalyzedEQ applies the EQ predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedEQ(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedNEQ applies the NEQ predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedNEQ(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedIn applies the In predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedIn(vs ...string) predicate.BaselineCharacteristicsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBaselineTypeUnitsAnalyzed), v...))
	})
}

// BaselineTypeUnitsAnalyzedNotIn applies the NotIn predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedNotIn(vs ...string) predicate.BaselineCharacteristicsModule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBaselineTypeUnitsAnalyzed), v...))
	})
}

// BaselineTypeUnitsAnalyzedGT applies the GT predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedGT(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedGTE applies the GTE predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedGTE(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedLT applies the LT predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedLT(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedLTE applies the LTE predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedLTE(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedContains applies the Contains predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedContains(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedHasPrefix applies the HasPrefix predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedHasPrefix(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedHasSuffix applies the HasSuffix predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedHasSuffix(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedEqualFold applies the EqualFold predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedEqualFold(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// BaselineTypeUnitsAnalyzedContainsFold applies the ContainsFold predicate on the "baseline_type_units_analyzed" field.
func BaselineTypeUnitsAnalyzedContainsFold(v string) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaselineTypeUnitsAnalyzed), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.ResultsDefinition) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineGroupList applies the HasEdge predicate on the "baseline_group_list" edge.
func HasBaselineGroupList() predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineGroupListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineGroupListTable, BaselineGroupListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineGroupListWith applies the HasEdge predicate on the "baseline_group_list" edge with a given conditions (other predicates).
func HasBaselineGroupListWith(preds ...predicate.BaselineGroup) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineGroupListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineGroupListTable, BaselineGroupListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineDenomList applies the HasEdge predicate on the "baseline_denom_list" edge.
func HasBaselineDenomList() predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineDenomListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineDenomListTable, BaselineDenomListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineDenomListWith applies the HasEdge predicate on the "baseline_denom_list" edge with a given conditions (other predicates).
func HasBaselineDenomListWith(preds ...predicate.BaselineDenom) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineDenomListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineDenomListTable, BaselineDenomListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBaselineMeasureList applies the HasEdge predicate on the "baseline_measure_list" edge.
func HasBaselineMeasureList() predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureListTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureListTable, BaselineMeasureListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaselineMeasureListWith applies the HasEdge predicate on the "baseline_measure_list" edge with a given conditions (other predicates).
func HasBaselineMeasureListWith(preds ...predicate.BaselineMeasure) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaselineMeasureListInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BaselineMeasureListTable, BaselineMeasureListColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaselineCharacteristicsModule) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaselineCharacteristicsModule) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
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
func Not(p predicate.BaselineCharacteristicsModule) predicate.BaselineCharacteristicsModule {
	return predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
		p(s.Not())
	})
}
