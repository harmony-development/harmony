// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/permissionnode"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// PermissionNode is the model entity for the PermissionNode schema.
type PermissionNode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Node holds the value of the "node" field.
	Node string `json:"node,omitempty"`
	// Allow holds the value of the "allow" field.
	Allow bool `json:"allow,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionNodeQuery when eager-loading is set.
	Edges                   PermissionNodeEdges `json:"edges"`
	channel_permission_node *uint64
	guild_permission_node   *uint64
	role_permission_node    *uint64
}

// PermissionNodeEdges holds the relations/edges for other nodes in the graph.
type PermissionNodeEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// Guild holds the value of the guild edge.
	Guild *Guild `json:"guild,omitempty"`
	// Channel holds the value of the channel edge.
	Channel *Channel `json:"channel,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionNodeEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[0] {
		if e.Role == nil {
			// The edge role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// GuildOrErr returns the Guild value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionNodeEdges) GuildOrErr() (*Guild, error) {
	if e.loadedTypes[1] {
		if e.Guild == nil {
			// The edge guild was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: guild.Label}
		}
		return e.Guild, nil
	}
	return nil, &NotLoadedError{edge: "guild"}
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionNodeEdges) ChannelOrErr() (*Channel, error) {
	if e.loadedTypes[2] {
		if e.Channel == nil {
			// The edge channel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: channel.Label}
		}
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionNode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissionnode.FieldAllow:
			values[i] = new(sql.NullBool)
		case permissionnode.FieldID:
			values[i] = new(sql.NullInt64)
		case permissionnode.FieldNode:
			values[i] = new(sql.NullString)
		case permissionnode.ForeignKeys[0]: // channel_permission_node
			values[i] = new(sql.NullInt64)
		case permissionnode.ForeignKeys[1]: // guild_permission_node
			values[i] = new(sql.NullInt64)
		case permissionnode.ForeignKeys[2]: // role_permission_node
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PermissionNode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PermissionNode fields.
func (pn *PermissionNode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permissionnode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pn.ID = int(value.Int64)
		case permissionnode.FieldNode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node", values[i])
			} else if value.Valid {
				pn.Node = value.String
			}
		case permissionnode.FieldAllow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field allow", values[i])
			} else if value.Valid {
				pn.Allow = value.Bool
			}
		case permissionnode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field channel_permission_node", value)
			} else if value.Valid {
				pn.channel_permission_node = new(uint64)
				*pn.channel_permission_node = uint64(value.Int64)
			}
		case permissionnode.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field guild_permission_node", value)
			} else if value.Valid {
				pn.guild_permission_node = new(uint64)
				*pn.guild_permission_node = uint64(value.Int64)
			}
		case permissionnode.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_permission_node", value)
			} else if value.Valid {
				pn.role_permission_node = new(uint64)
				*pn.role_permission_node = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryRole queries the "role" edge of the PermissionNode entity.
func (pn *PermissionNode) QueryRole() *RoleQuery {
	return (&PermissionNodeClient{config: pn.config}).QueryRole(pn)
}

// QueryGuild queries the "guild" edge of the PermissionNode entity.
func (pn *PermissionNode) QueryGuild() *GuildQuery {
	return (&PermissionNodeClient{config: pn.config}).QueryGuild(pn)
}

// QueryChannel queries the "channel" edge of the PermissionNode entity.
func (pn *PermissionNode) QueryChannel() *ChannelQuery {
	return (&PermissionNodeClient{config: pn.config}).QueryChannel(pn)
}

// Update returns a builder for updating this PermissionNode.
// Note that you need to call PermissionNode.Unwrap() before calling this method if this PermissionNode
// was returned from a transaction, and the transaction was committed or rolled back.
func (pn *PermissionNode) Update() *PermissionNodeUpdateOne {
	return (&PermissionNodeClient{config: pn.config}).UpdateOne(pn)
}

// Unwrap unwraps the PermissionNode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pn *PermissionNode) Unwrap() *PermissionNode {
	tx, ok := pn.config.driver.(*txDriver)
	if !ok {
		panic("entgen: PermissionNode is not a transactional entity")
	}
	pn.config.driver = tx.drv
	return pn
}

// String implements the fmt.Stringer.
func (pn *PermissionNode) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionNode(")
	builder.WriteString(fmt.Sprintf("id=%v", pn.ID))
	builder.WriteString(", node=")
	builder.WriteString(pn.Node)
	builder.WriteString(", allow=")
	builder.WriteString(fmt.Sprintf("%v", pn.Allow))
	builder.WriteByte(')')
	return builder.String()
}

// PermissionNodes is a parsable slice of PermissionNode.
type PermissionNodes []*PermissionNode

func (pn PermissionNodes) config(cfg config) {
	for _i := range pn {
		pn[_i].config = cfg
	}
}
