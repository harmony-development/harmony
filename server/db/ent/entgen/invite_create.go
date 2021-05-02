// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
)

// InviteCreate is the builder for creating a Invite entity.
type InviteCreate struct {
	config
	mutation *InviteMutation
	hooks    []Hook
}

// SetUses sets the "uses" field.
func (ic *InviteCreate) SetUses(i int64) *InviteCreate {
	ic.mutation.SetUses(i)
	return ic
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (ic *InviteCreate) SetNillableUses(i *int64) *InviteCreate {
	if i != nil {
		ic.SetUses(*i)
	}
	return ic
}

// SetPossibleUses sets the "possible_uses" field.
func (ic *InviteCreate) SetPossibleUses(i int64) *InviteCreate {
	ic.mutation.SetPossibleUses(i)
	return ic
}

// SetNillablePossibleUses sets the "possible_uses" field if the given value is not nil.
func (ic *InviteCreate) SetNillablePossibleUses(i *int64) *InviteCreate {
	if i != nil {
		ic.SetPossibleUses(*i)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InviteCreate) SetID(s string) *InviteCreate {
	ic.mutation.SetID(s)
	return ic
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (ic *InviteCreate) SetGuildID(id uint64) *InviteCreate {
	ic.mutation.SetGuildID(id)
	return ic
}

// SetNillableGuildID sets the "guild" edge to the Guild entity by ID if the given value is not nil.
func (ic *InviteCreate) SetNillableGuildID(id *uint64) *InviteCreate {
	if id != nil {
		ic = ic.SetGuildID(*id)
	}
	return ic
}

// SetGuild sets the "guild" edge to the Guild entity.
func (ic *InviteCreate) SetGuild(g *Guild) *InviteCreate {
	return ic.SetGuildID(g.ID)
}

// Mutation returns the InviteMutation object of the builder.
func (ic *InviteCreate) Mutation() *InviteMutation {
	return ic.mutation
}

// Save creates the Invite in the database.
func (ic *InviteCreate) Save(ctx context.Context) (*Invite, error) {
	var (
		err  error
		node *Invite
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			node, err = ic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InviteCreate) SaveX(ctx context.Context) *Invite {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ic *InviteCreate) defaults() {
	if _, ok := ic.mutation.Uses(); !ok {
		v := invite.DefaultUses
		ic.mutation.SetUses(v)
	}
	if _, ok := ic.mutation.PossibleUses(); !ok {
		v := invite.DefaultPossibleUses
		ic.mutation.SetPossibleUses(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InviteCreate) check() error {
	if _, ok := ic.mutation.Uses(); !ok {
		return &ValidationError{Name: "uses", err: errors.New("entgen: missing required field \"uses\"")}
	}
	if _, ok := ic.mutation.PossibleUses(); !ok {
		return &ValidationError{Name: "possible_uses", err: errors.New("entgen: missing required field \"possible_uses\"")}
	}
	return nil
}

func (ic *InviteCreate) sqlSave(ctx context.Context) (*Invite, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ic *InviteCreate) createSpec() (*Invite, *sqlgraph.CreateSpec) {
	var (
		_node = &Invite{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: invite.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: invite.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.Uses(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldUses,
		})
		_node.Uses = value
	}
	if value, ok := ic.mutation.PossibleUses(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldPossibleUses,
		})
		_node.PossibleUses = value
	}
	if nodes := ic.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.guild_invite = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InviteCreateBulk is the builder for creating many Invite entities in bulk.
type InviteCreateBulk struct {
	config
	builders []*InviteCreate
}

// Save creates the Invite entities in the database.
func (icb *InviteCreateBulk) Save(ctx context.Context) ([]*Invite, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invite, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InviteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InviteCreateBulk) SaveX(ctx context.Context) []*Invite {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
