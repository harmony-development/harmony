// Code generated by entc, DO NOT EDIT.

package guildlistentry

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
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
func IDGT(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHost), v))
	})
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.GuildListEntry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildListEntry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHost), v...))
	})
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.GuildListEntry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildListEntry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHost), v...))
	})
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHost), v))
	})
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHost), v))
	})
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHost), v))
	})
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHost), v))
	})
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHost), v))
	})
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHost), v))
	})
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHost), v))
	})
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHost), v))
	})
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHost), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.GuildListEntry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildListEntry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.GuildListEntry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildListEntry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
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
func And(predicates ...predicate.GuildListEntry) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GuildListEntry) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
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
func Not(p predicate.GuildListEntry) predicate.GuildListEntry {
	return predicate.GuildListEntry(func(s *sql.Selector) {
		p(s.Not())
	})
}
