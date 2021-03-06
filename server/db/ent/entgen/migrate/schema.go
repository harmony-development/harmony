// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "kind", Type: field.TypeUint64},
		{Name: "position", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "guild_channel", Type: field.TypeUint64, Nullable: true},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "channels_guilds_channel",
				Columns:    []*schema.Column{ChannelsColumns[5]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmotesColumns holds the columns for the "emotes" table.
	EmotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "emote_pack_emote", Type: field.TypeUint64, Nullable: true},
	}
	// EmotesTable holds the schema information for the "emotes" table.
	EmotesTable = &schema.Table{
		Name:       "emotes",
		Columns:    EmotesColumns,
		PrimaryKey: []*schema.Column{EmotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emotes_emote_packs_emote",
				Columns:    []*schema.Column{EmotesColumns[2]},
				RefColumns: []*schema.Column{EmotePacksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmotePacksColumns holds the columns for the "emote_packs" table.
	EmotePacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_emotepack", Type: field.TypeUint64, Nullable: true},
		{Name: "user_createdpacks", Type: field.TypeUint64, Nullable: true},
	}
	// EmotePacksTable holds the schema information for the "emote_packs" table.
	EmotePacksTable = &schema.Table{
		Name:       "emote_packs",
		Columns:    EmotePacksColumns,
		PrimaryKey: []*schema.Column{EmotePacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emote_packs_users_emotepack",
				Columns:    []*schema.Column{EmotePacksColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "emote_packs_users_createdpacks",
				Columns:    []*schema.Column{EmotePacksColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "contenttype", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:        "files",
		Columns:     FilesColumns,
		PrimaryKey:  []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FileHashesColumns holds the columns for the "file_hashes" table.
	FileHashesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeBytes},
		{Name: "fileid", Type: field.TypeString},
	}
	// FileHashesTable holds the schema information for the "file_hashes" table.
	FileHashesTable = &schema.Table{
		Name:        "file_hashes",
		Columns:     FileHashesColumns,
		PrimaryKey:  []*schema.Column{FileHashesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "filehash_hash",
				Unique:  true,
				Columns: []*schema.Column{FileHashesColumns[1]},
			},
		},
	}
	// ForeignUsersColumns holds the columns for the "foreign_users" table.
	ForeignUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "foreignid", Type: field.TypeUint64},
		{Name: "host", Type: field.TypeString},
		{Name: "user_foreign_user", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// ForeignUsersTable holds the schema information for the "foreign_users" table.
	ForeignUsersTable = &schema.Table{
		Name:       "foreign_users",
		Columns:    ForeignUsersColumns,
		PrimaryKey: []*schema.Column{ForeignUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "foreign_users_users_foreign_user",
				Columns:    []*schema.Column{ForeignUsersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GuildsColumns holds the columns for the "guilds" table.
	GuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "picture", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeBytes},
		{Name: "guild_owner", Type: field.TypeUint64, Nullable: true},
	}
	// GuildsTable holds the schema information for the "guilds" table.
	GuildsTable = &schema.Table{
		Name:       "guilds",
		Columns:    GuildsColumns,
		PrimaryKey: []*schema.Column{GuildsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "guilds_users_owner",
				Columns:    []*schema.Column{GuildsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GuildListEntriesColumns holds the columns for the "guild_list_entries" table.
	GuildListEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "guild_id", Type: field.TypeUint64},
		{Name: "host", Type: field.TypeString},
		{Name: "position", Type: field.TypeString},
		{Name: "user_listentry", Type: field.TypeUint64, Nullable: true},
	}
	// GuildListEntriesTable holds the schema information for the "guild_list_entries" table.
	GuildListEntriesTable = &schema.Table{
		Name:       "guild_list_entries",
		Columns:    GuildListEntriesColumns,
		PrimaryKey: []*schema.Column{GuildListEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "guild_list_entries_users_listentry",
				Columns:    []*schema.Column{GuildListEntriesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HostsColumns holds the columns for the "hosts" table.
	HostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "host", Type: field.TypeString, Unique: true},
		{Name: "eventqueue", Type: field.TypeBytes},
	}
	// HostsTable holds the schema information for the "hosts" table.
	HostsTable = &schema.Table{
		Name:        "hosts",
		Columns:     HostsColumns,
		PrimaryKey:  []*schema.Column{HostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// InvitesColumns holds the columns for the "invites" table.
	InvitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "uses", Type: field.TypeInt64, Default: 0},
		{Name: "possible_uses", Type: field.TypeInt64, Default: -1},
		{Name: "guild_invite", Type: field.TypeUint64, Nullable: true},
	}
	// InvitesTable holds the schema information for the "invites" table.
	InvitesTable = &schema.Table{
		Name:       "invites",
		Columns:    InvitesColumns,
		PrimaryKey: []*schema.Column{InvitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invites_guilds_invite",
				Columns:    []*schema.Column{InvitesColumns[3]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocalUsersColumns holds the columns for the "local_users" table.
	LocalUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
		{Name: "user_local_user", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// LocalUsersTable holds the schema information for the "local_users" table.
	LocalUsersTable = &schema.Table{
		Name:       "local_users",
		Columns:    LocalUsersColumns,
		PrimaryKey: []*schema.Column{LocalUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "local_users_users_local_user",
				Columns:    []*schema.Column{LocalUsersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "createdat", Type: field.TypeTime},
		{Name: "editedat", Type: field.TypeTime, Nullable: true},
		{Name: "metadata", Type: field.TypeBytes, Nullable: true},
		{Name: "override", Type: field.TypeBytes, Nullable: true},
		{Name: "content", Type: field.TypeBytes},
		{Name: "channel_message", Type: field.TypeUint64, Nullable: true},
		{Name: "message_replies", Type: field.TypeUint64, Nullable: true},
		{Name: "user_message", Type: field.TypeUint64, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_channels_message",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_messages_replies",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_users_message",
				Columns:    []*schema.Column{MessagesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PermissionNodesColumns holds the columns for the "permission_nodes" table.
	PermissionNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "node", Type: field.TypeString},
		{Name: "allow", Type: field.TypeBool},
		{Name: "channel_permission_node", Type: field.TypeUint64, Nullable: true},
		{Name: "guild_permission_node", Type: field.TypeUint64, Nullable: true},
		{Name: "role_permission_node", Type: field.TypeUint64, Nullable: true},
	}
	// PermissionNodesTable holds the schema information for the "permission_nodes" table.
	PermissionNodesTable = &schema.Table{
		Name:       "permission_nodes",
		Columns:    PermissionNodesColumns,
		PrimaryKey: []*schema.Column{PermissionNodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permission_nodes_channels_permission_node",
				Columns:    []*schema.Column{PermissionNodesColumns[3]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "permission_nodes_guilds_permission_node",
				Columns:    []*schema.Column{PermissionNodesColumns[4]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "permission_nodes_roles_permission_node",
				Columns:    []*schema.Column{PermissionNodesColumns[5]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "status", Type: field.TypeInt16, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "is_bot", Type: field.TypeBool, Default: false},
		{Name: "user_profile", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profiles_users_profile",
				Columns:    []*schema.Column{ProfilesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "color", Type: field.TypeInt},
		{Name: "hoist", Type: field.TypeBool},
		{Name: "pingable", Type: field.TypeBool},
		{Name: "position", Type: field.TypeString},
		{Name: "channel_role", Type: field.TypeUint64, Nullable: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "roles_channels_role",
				Columns:    []*schema.Column{RolesColumns[6]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "expires", Type: field.TypeTime},
		{Name: "local_user_sessions", Type: field.TypeInt, Nullable: true},
		{Name: "user_sessions", Type: field.TypeUint64, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_local_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[2]},
				RefColumns: []*schema.Column{LocalUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "guild_bans", Type: field.TypeUint64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_guilds_bans",
				Columns:    []*schema.Column{UsersColumns[1]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserMetaColumns holds the columns for the "user_meta" table.
	UserMetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "meta", Type: field.TypeString},
		{Name: "user_metadata", Type: field.TypeUint64, Nullable: true},
	}
	// UserMetaTable holds the schema information for the "user_meta" table.
	UserMetaTable = &schema.Table{
		Name:       "user_meta",
		Columns:    UserMetaColumns,
		PrimaryKey: []*schema.Column{UserMetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_meta_users_metadata",
				Columns:    []*schema.Column{UserMetaColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GuildRoleColumns holds the columns for the "guild_role" table.
	GuildRoleColumns = []*schema.Column{
		{Name: "guild_id", Type: field.TypeUint64},
		{Name: "role_id", Type: field.TypeUint64},
	}
	// GuildRoleTable holds the schema information for the "guild_role" table.
	GuildRoleTable = &schema.Table{
		Name:       "guild_role",
		Columns:    GuildRoleColumns,
		PrimaryKey: []*schema.Column{GuildRoleColumns[0], GuildRoleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "guild_role_guild_id",
				Columns:    []*schema.Column{GuildRoleColumns[0]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "guild_role_role_id",
				Columns:    []*schema.Column{GuildRoleColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleMembersColumns holds the columns for the "role_members" table.
	RoleMembersColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// RoleMembersTable holds the schema information for the "role_members" table.
	RoleMembersTable = &schema.Table{
		Name:       "role_members",
		Columns:    RoleMembersColumns,
		PrimaryKey: []*schema.Column{RoleMembersColumns[0], RoleMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_members_role_id",
				Columns:    []*schema.Column{RoleMembersColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_members_user_id",
				Columns:    []*schema.Column{RoleMembersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGuildColumns holds the columns for the "user_guild" table.
	UserGuildColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "guild_id", Type: field.TypeUint64},
	}
	// UserGuildTable holds the schema information for the "user_guild" table.
	UserGuildTable = &schema.Table{
		Name:       "user_guild",
		Columns:    UserGuildColumns,
		PrimaryKey: []*schema.Column{UserGuildColumns[0], UserGuildColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_guild_user_id",
				Columns:    []*schema.Column{UserGuildColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_guild_guild_id",
				Columns:    []*schema.Column{UserGuildColumns[1]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		EmotesTable,
		EmotePacksTable,
		FilesTable,
		FileHashesTable,
		ForeignUsersTable,
		GuildsTable,
		GuildListEntriesTable,
		HostsTable,
		InvitesTable,
		LocalUsersTable,
		MessagesTable,
		PermissionNodesTable,
		ProfilesTable,
		RolesTable,
		SessionsTable,
		UsersTable,
		UserMetaTable,
		GuildRoleTable,
		RoleMembersTable,
		UserGuildTable,
	}
)

func init() {
	ChannelsTable.ForeignKeys[0].RefTable = GuildsTable
	EmotesTable.ForeignKeys[0].RefTable = EmotePacksTable
	EmotePacksTable.ForeignKeys[0].RefTable = UsersTable
	EmotePacksTable.ForeignKeys[1].RefTable = UsersTable
	ForeignUsersTable.ForeignKeys[0].RefTable = UsersTable
	GuildsTable.ForeignKeys[0].RefTable = UsersTable
	GuildListEntriesTable.ForeignKeys[0].RefTable = UsersTable
	InvitesTable.ForeignKeys[0].RefTable = GuildsTable
	LocalUsersTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = ChannelsTable
	MessagesTable.ForeignKeys[1].RefTable = MessagesTable
	MessagesTable.ForeignKeys[2].RefTable = UsersTable
	PermissionNodesTable.ForeignKeys[0].RefTable = ChannelsTable
	PermissionNodesTable.ForeignKeys[1].RefTable = GuildsTable
	PermissionNodesTable.ForeignKeys[2].RefTable = RolesTable
	ProfilesTable.ForeignKeys[0].RefTable = UsersTable
	RolesTable.ForeignKeys[0].RefTable = ChannelsTable
	SessionsTable.ForeignKeys[0].RefTable = LocalUsersTable
	SessionsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = GuildsTable
	UserMetaTable.ForeignKeys[0].RefTable = UsersTable
	GuildRoleTable.ForeignKeys[0].RefTable = GuildsTable
	GuildRoleTable.ForeignKeys[1].RefTable = RolesTable
	RoleMembersTable.ForeignKeys[0].RefTable = RolesTable
	RoleMembersTable.ForeignKeys[1].RefTable = UsersTable
	UserGuildTable.ForeignKeys[0].RefTable = UsersTable
	UserGuildTable.ForeignKeys[1].RefTable = GuildsTable
}
