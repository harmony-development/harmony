// Code generated by entc, DO NOT EDIT.

package usermeta

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
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
func IDGT(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Meta applies equality check predicate on the "meta" field. It's identical to MetaEQ.
func Meta(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMeta), v))
	})
}

// MetaEQ applies the EQ predicate on the "meta" field.
func MetaEQ(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMeta), v))
	})
}

// MetaNEQ applies the NEQ predicate on the "meta" field.
func MetaNEQ(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMeta), v))
	})
}

// MetaIn applies the In predicate on the "meta" field.
func MetaIn(vs ...string) predicate.UserMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMeta), v...))
	})
}

// MetaNotIn applies the NotIn predicate on the "meta" field.
func MetaNotIn(vs ...string) predicate.UserMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMeta), v...))
	})
}

// MetaGT applies the GT predicate on the "meta" field.
func MetaGT(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMeta), v))
	})
}

// MetaGTE applies the GTE predicate on the "meta" field.
func MetaGTE(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMeta), v))
	})
}

// MetaLT applies the LT predicate on the "meta" field.
func MetaLT(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMeta), v))
	})
}

// MetaLTE applies the LTE predicate on the "meta" field.
func MetaLTE(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMeta), v))
	})
}

// MetaContains applies the Contains predicate on the "meta" field.
func MetaContains(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMeta), v))
	})
}

// MetaHasPrefix applies the HasPrefix predicate on the "meta" field.
func MetaHasPrefix(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMeta), v))
	})
}

// MetaHasSuffix applies the HasSuffix predicate on the "meta" field.
func MetaHasSuffix(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMeta), v))
	})
}

// MetaEqualFold applies the EqualFold predicate on the "meta" field.
func MetaEqualFold(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMeta), v))
	})
}

// MetaContainsFold applies the ContainsFold predicate on the "meta" field.
func MetaContainsFold(v string) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMeta), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserMeta) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserMeta) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
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
func Not(p predicate.UserMeta) predicate.UserMeta {
	return predicate.UserMeta(func(s *sql.Selector) {
		p(s.Not())
	})
}
