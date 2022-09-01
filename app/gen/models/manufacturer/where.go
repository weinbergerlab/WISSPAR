// Code generated (@generated) by entc, DO NOT EDIT.

package manufacturer

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ManufacturerName applies equality check predicate on the "manufacturer_name" field. It's identical to ManufacturerNameEQ.
func ManufacturerName(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameEQ applies the EQ predicate on the "manufacturer_name" field.
func ManufacturerNameEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameNEQ applies the NEQ predicate on the "manufacturer_name" field.
func ManufacturerNameNEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameIn applies the In predicate on the "manufacturer_name" field.
func ManufacturerNameIn(vs ...string) predicate.Manufacturer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManufacturerName), v...))
	})
}

// ManufacturerNameNotIn applies the NotIn predicate on the "manufacturer_name" field.
func ManufacturerNameNotIn(vs ...string) predicate.Manufacturer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManufacturerName), v...))
	})
}

// ManufacturerNameGT applies the GT predicate on the "manufacturer_name" field.
func ManufacturerNameGT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameGTE applies the GTE predicate on the "manufacturer_name" field.
func ManufacturerNameGTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameLT applies the LT predicate on the "manufacturer_name" field.
func ManufacturerNameLT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameLTE applies the LTE predicate on the "manufacturer_name" field.
func ManufacturerNameLTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameContains applies the Contains predicate on the "manufacturer_name" field.
func ManufacturerNameContains(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameHasPrefix applies the HasPrefix predicate on the "manufacturer_name" field.
func ManufacturerNameHasPrefix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameHasSuffix applies the HasSuffix predicate on the "manufacturer_name" field.
func ManufacturerNameHasSuffix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameEqualFold applies the EqualFold predicate on the "manufacturer_name" field.
func ManufacturerNameEqualFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManufacturerName), v))
	})
}

// ManufacturerNameContainsFold applies the ContainsFold predicate on the "manufacturer_name" field.
func ManufacturerNameContainsFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManufacturerName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
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
func Not(p predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		p(s.Not())
	})
}
