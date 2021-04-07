package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/lexorank"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *database) AddRoleToGuild(guildID, roleID uint64, name string, color int, hoist, pingable bool) (err error) {
	defer doRecovery(&err)
	d.Guild.
		UpdateOneID(guildID).
		AddRole(
			d.Role.
				Create().
				SetName(name).
				SetColor(color).
				SetHoist(hoist).
				SetPingable(pingable).
				SaveX(ctx),
		).
		ExecX(ctx)
	return
}

func (d *database) GetGuildRoles(guildID uint64) (roles []*entgen.Role, err error) {
	defer doRecovery(&err)
	roles = d.Guild.GetX(ctx, guildID).QueryRole().AllX(ctx)
	return
}

func (d *database) GetPermissions(roleID uint64) (permissions []types.PermissionsNode, err error) {
	defer doRecovery(&err)
	nodes := d.Role.GetX(ctx, roleID).QueryPermissionNode().AllX(ctx)
	for _, node := range nodes {
		permissions = append(permissions, types.PermissionsNode{
			Node:  node.Node,
			Allow: node.Allow,
		})
	}
	return
}

func (d *database) GetPermissionsData(guildID uint64) (data types.PermissionsData, err error) {
	defer doRecovery(&err)
	roles := d.Guild.GetX(ctx, guildID).QueryRole().WithPermissionNode().AllX(ctx)
	for _, role := range roles {
		if perms, err := d.GetPermissions(role.ID); err != nil {
			panic(err)
		} else {
			data.Roles[role.ID] = perms
		}
	}
	chans := d.Guild.Query().Where(guild.ID(guildID)).QueryChannel().Order(entgen.Asc(channel.FieldPosition)).WithRole().AllX(ctx)
	var category uint64 = 0
	data.Channels = make(map[uint64]map[uint64][]types.PermissionsNode)
	for _, c := range chans {
		if c.Kind == uint64(types.ChannelKindCategory) {
			category = c.ID
		} else if category != 0 {
			categoryData, ok := data.Categories[category]
			_ = ok
			data.Categories[category] = append(categoryData, c.ID)
		}
		for _, role := range roles {
			if perms, err := d.GetPermissions(role.ID); err != nil {
				panic(err)
			} else {
				data.Channels[c.ID][role.ID] = perms
			}
		}
	}
	return
}

func (d *database) MoveRole(guildID, roleID, previousRole, nextRole uint64) (err error) {
	defer doRecovery(&err)
	previousPos := d.Role.GetX(ctx, previousRole).Position
	nextPos := d.Role.GetX(ctx, nextRole).Position
	d.Role.
		UpdateOneID(roleID).
		SetPosition(
			lexorank.Rank(previousPos, nextPos),
		).
		ExecX(ctx)
	return
}
