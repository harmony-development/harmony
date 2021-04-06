// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guildlistentry"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// GuildListEntryCreate is the builder for creating a GuildListEntry entity.
type GuildListEntryCreate struct {
	config
	mutation *GuildListEntryMutation
	hooks    []Hook
}

// SetHost sets the "host" field.
func (glec *GuildListEntryCreate) SetHost(s string) *GuildListEntryCreate {
	glec.mutation.SetHost(s)
	return glec
}

// SetPosition sets the "position" field.
func (glec *GuildListEntryCreate) SetPosition(s string) *GuildListEntryCreate {
	glec.mutation.SetPosition(s)
	return glec
}

// SetID sets the "id" field.
func (glec *GuildListEntryCreate) SetID(u uint64) *GuildListEntryCreate {
	glec.mutation.SetID(u)
	return glec
}

// SetUserID sets the "user" edge to the User entity by ID.
func (glec *GuildListEntryCreate) SetUserID(id uint64) *GuildListEntryCreate {
	glec.mutation.SetUserID(id)
	return glec
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (glec *GuildListEntryCreate) SetNillableUserID(id *uint64) *GuildListEntryCreate {
	if id != nil {
		glec = glec.SetUserID(*id)
	}
	return glec
}

// SetUser sets the "user" edge to the User entity.
func (glec *GuildListEntryCreate) SetUser(u *User) *GuildListEntryCreate {
	return glec.SetUserID(u.ID)
}

// Mutation returns the GuildListEntryMutation object of the builder.
func (glec *GuildListEntryCreate) Mutation() *GuildListEntryMutation {
	return glec.mutation
}

// Save creates the GuildListEntry in the database.
func (glec *GuildListEntryCreate) Save(ctx context.Context) (*GuildListEntry, error) {
	var (
		err  error
		node *GuildListEntry
	)
	if len(glec.hooks) == 0 {
		if err = glec.check(); err != nil {
			return nil, err
		}
		node, err = glec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildListEntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = glec.check(); err != nil {
				return nil, err
			}
			glec.mutation = mutation
			node, err = glec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(glec.hooks) - 1; i >= 0; i-- {
			mut = glec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, glec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (glec *GuildListEntryCreate) SaveX(ctx context.Context) *GuildListEntry {
	v, err := glec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (glec *GuildListEntryCreate) check() error {
	if _, ok := glec.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New("entgen: missing required field \"host\"")}
	}
	if _, ok := glec.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New("entgen: missing required field \"position\"")}
	}
	return nil
}

func (glec *GuildListEntryCreate) sqlSave(ctx context.Context) (*GuildListEntry, error) {
	_node, _spec := glec.createSpec()
	if err := sqlgraph.CreateNode(ctx, glec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (glec *GuildListEntryCreate) createSpec() (*GuildListEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &GuildListEntry{config: glec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: guildlistentry.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guildlistentry.FieldID,
			},
		}
	)
	if id, ok := glec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := glec.mutation.Host(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildlistentry.FieldHost,
		})
		_node.Host = value
	}
	if value, ok := glec.mutation.Position(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildlistentry.FieldPosition,
		})
		_node.Position = value
	}
	if nodes := glec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   guildlistentry.UserTable,
			Columns: []string{guildlistentry.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_listentry = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GuildListEntryCreateBulk is the builder for creating many GuildListEntry entities in bulk.
type GuildListEntryCreateBulk struct {
	config
	builders []*GuildListEntryCreate
}

// Save creates the GuildListEntry entities in the database.
func (glecb *GuildListEntryCreateBulk) Save(ctx context.Context) ([]*GuildListEntry, error) {
	specs := make([]*sqlgraph.CreateSpec, len(glecb.builders))
	nodes := make([]*GuildListEntry, len(glecb.builders))
	mutators := make([]Mutator, len(glecb.builders))
	for i := range glecb.builders {
		func(i int, root context.Context) {
			builder := glecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GuildListEntryMutation)
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
					_, err = mutators[i+1].Mutate(root, glecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, glecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, glecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (glecb *GuildListEntryCreateBulk) SaveX(ctx context.Context) []*GuildListEntry {
	v, err := glecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}