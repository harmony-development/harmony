// Code generated by entc, DO NOT EDIT.

package guild

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
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
func IDGT(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Picture applies equality check predicate on the "picture" field. It's identical to PictureEQ.
func Picture(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPicture), v))
	})
}

// Metadata applies equality check predicate on the "metadata" field. It's identical to MetadataEQ.
func Metadata(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwner), v))
	})
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...uint64) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwner), v...))
	})
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...uint64) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwner), v...))
	})
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwner), v))
	})
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwner), v))
	})
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwner), v))
	})
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v uint64) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwner), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PictureEQ applies the EQ predicate on the "picture" field.
func PictureEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPicture), v))
	})
}

// PictureNEQ applies the NEQ predicate on the "picture" field.
func PictureNEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPicture), v))
	})
}

// PictureIn applies the In predicate on the "picture" field.
func PictureIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPicture), v...))
	})
}

// PictureNotIn applies the NotIn predicate on the "picture" field.
func PictureNotIn(vs ...string) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPicture), v...))
	})
}

// PictureGT applies the GT predicate on the "picture" field.
func PictureGT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPicture), v))
	})
}

// PictureGTE applies the GTE predicate on the "picture" field.
func PictureGTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPicture), v))
	})
}

// PictureLT applies the LT predicate on the "picture" field.
func PictureLT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPicture), v))
	})
}

// PictureLTE applies the LTE predicate on the "picture" field.
func PictureLTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPicture), v))
	})
}

// PictureContains applies the Contains predicate on the "picture" field.
func PictureContains(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPicture), v))
	})
}

// PictureHasPrefix applies the HasPrefix predicate on the "picture" field.
func PictureHasPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPicture), v))
	})
}

// PictureHasSuffix applies the HasSuffix predicate on the "picture" field.
func PictureHasSuffix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPicture), v))
	})
}

// PictureEqualFold applies the EqualFold predicate on the "picture" field.
func PictureEqualFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPicture), v))
	})
}

// PictureContainsFold applies the ContainsFold predicate on the "picture" field.
func PictureContainsFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPicture), v))
	})
}

// MetadataEQ applies the EQ predicate on the "metadata" field.
func MetadataEQ(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// MetadataNEQ applies the NEQ predicate on the "metadata" field.
func MetadataNEQ(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMetadata), v))
	})
}

// MetadataIn applies the In predicate on the "metadata" field.
func MetadataIn(vs ...[]byte) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMetadata), v...))
	})
}

// MetadataNotIn applies the NotIn predicate on the "metadata" field.
func MetadataNotIn(vs ...[]byte) predicate.Guild {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMetadata), v...))
	})
}

// MetadataGT applies the GT predicate on the "metadata" field.
func MetadataGT(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMetadata), v))
	})
}

// MetadataGTE applies the GTE predicate on the "metadata" field.
func MetadataGTE(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMetadata), v))
	})
}

// MetadataLT applies the LT predicate on the "metadata" field.
func MetadataLT(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMetadata), v))
	})
}

// MetadataLTE applies the LTE predicate on the "metadata" field.
func MetadataLTE(v []byte) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMetadata), v))
	})
}

// HasInvite applies the HasEdge predicate on the "invite" edge.
func HasInvite() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InviteTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InviteTable, InviteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInviteWith applies the HasEdge predicate on the "invite" edge with a given conditions (other predicates).
func HasInviteWith(preds ...predicate.Invite) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InviteInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InviteTable, InviteColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBans applies the HasEdge predicate on the "bans" edge.
func HasBans() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BansTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BansTable, BansColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBansWith applies the HasEdge predicate on the "bans" edge with a given conditions (other predicates).
func HasBansWith(preds ...predicate.User) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BansInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BansTable, BansColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RoleTable, RolePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RoleTable, RolePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
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
func Not(p predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		p(s.Not())
	})
}
