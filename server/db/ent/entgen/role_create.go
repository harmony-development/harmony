// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/permissionnode"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetColor sets the "color" field.
func (rc *RoleCreate) SetColor(i int) *RoleCreate {
	rc.mutation.SetColor(i)
	return rc
}

// SetHoist sets the "hoist" field.
func (rc *RoleCreate) SetHoist(b bool) *RoleCreate {
	rc.mutation.SetHoist(b)
	return rc
}

// SetPingable sets the "pingable" field.
func (rc *RoleCreate) SetPingable(b bool) *RoleCreate {
	rc.mutation.SetPingable(b)
	return rc
}

// SetPosition sets the "position" field.
func (rc *RoleCreate) SetPosition(s string) *RoleCreate {
	rc.mutation.SetPosition(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(u uint64) *RoleCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddGuildIDs adds the "guild" edge to the Guild entity by IDs.
func (rc *RoleCreate) AddGuildIDs(ids ...uint64) *RoleCreate {
	rc.mutation.AddGuildIDs(ids...)
	return rc
}

// AddGuild adds the "guild" edges to the Guild entity.
func (rc *RoleCreate) AddGuild(g ...*Guild) *RoleCreate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return rc.AddGuildIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (rc *RoleCreate) AddMemberIDs(ids ...uint64) *RoleCreate {
	rc.mutation.AddMemberIDs(ids...)
	return rc
}

// AddMembers adds the "members" edges to the User entity.
func (rc *RoleCreate) AddMembers(u ...*User) *RoleCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddMemberIDs(ids...)
}

// AddPermissionNodeIDs adds the "permission_node" edge to the PermissionNode entity by IDs.
func (rc *RoleCreate) AddPermissionNodeIDs(ids ...int) *RoleCreate {
	rc.mutation.AddPermissionNodeIDs(ids...)
	return rc
}

// AddPermissionNode adds the "permission_node" edges to the PermissionNode entity.
func (rc *RoleCreate) AddPermissionNode(p ...*PermissionNode) *RoleCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionNodeIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	var (
		err  error
		node *Role
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("entgen: missing required field \"name\"")}
	}
	if _, ok := rc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New("entgen: missing required field \"color\"")}
	}
	if _, ok := rc.mutation.Hoist(); !ok {
		return &ValidationError{Name: "hoist", err: errors.New("entgen: missing required field \"hoist\"")}
	}
	if _, ok := rc.mutation.Pingable(); !ok {
		return &ValidationError{Name: "pingable", err: errors.New("entgen: missing required field \"pingable\"")}
	}
	if _, ok := rc.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New("entgen: missing required field \"position\"")}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: role.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: role.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: role.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := rc.mutation.Hoist(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: role.FieldHoist,
		})
		_node.Hoist = value
	}
	if value, ok := rc.mutation.Pingable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: role.FieldPingable,
		})
		_node.Pingable = value
	}
	if value, ok := rc.mutation.Position(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldPosition,
		})
		_node.Position = value
	}
	if nodes := rc.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.GuildTable,
			Columns: role.GuildPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MembersTable,
			Columns: role.MembersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.PermissionNodeTable,
			Columns: []string{role.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	builders []*RoleCreate
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
