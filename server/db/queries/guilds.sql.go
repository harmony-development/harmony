// Code generated by sqlc. DO NOT EDIT.
// source: guilds.sql

package queries

import (
	"context"
	"database/sql"
)

const addUserToGuild = `-- name: AddUserToGuild :exec
INSERT INTO Guild_Members (
    User_ID, Guild_ID
) VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING
`

type AddUserToGuildParams struct {
	UserID  uint64 `json:"user_id"`
	GuildID uint64 `json:"guild_id"`
}

func (q *Queries) AddUserToGuild(ctx context.Context, arg AddUserToGuildParams) error {
	_, err := q.exec(ctx, q.addUserToGuildStmt, addUserToGuild, arg.UserID, arg.GuildID)
	return err
}

const createChannel = `-- name: CreateChannel :one
INSERT INTO Channels (
    Guild_ID, Channel_Name
) VALUES (
    $1, $2
)
RETURNING channel_id, guild_id, channel_name
`

type CreateChannelParams struct {
	GuildID     sql.NullInt64 `json:"guild_id"`
	ChannelName string        `json:"channel_name"`
}

func (q *Queries) CreateChannel(ctx context.Context, arg CreateChannelParams) (Channel, error) {
	row := q.queryRow(ctx, q.createChannelStmt, createChannel, arg.GuildID, arg.ChannelName)
	var i Channel
	err := row.Scan(&i.ChannelID, &i.GuildID, &i.ChannelName)
	return i, err
}

const createGuild = `-- name: CreateGuild :one
INSERT INTO Guilds (
    Guild_ID, Owner_ID, Guild_Name, Picture_URL
) VALUES (
    $1, $2, $3, $4
)
RETURNING guild_id, owner_id, guild_name, picture_url
`

type CreateGuildParams struct {
	GuildID    uint64 `json:"guild_id"`
	OwnerID    uint64 `json:"owner_id"`
	GuildName  string `json:"guild_name"`
	PictureUrl string `json:"picture_url"`
}

func (q *Queries) CreateGuild(ctx context.Context, arg CreateGuildParams) (Guild, error) {
	row := q.queryRow(ctx, q.createGuildStmt, createGuild,
		arg.GuildID,
		arg.OwnerID,
		arg.GuildName,
		arg.PictureUrl,
	)
	var i Guild
	err := row.Scan(
		&i.GuildID,
		&i.OwnerID,
		&i.GuildName,
		&i.PictureUrl,
	)
	return i, err
}

const deleteChannel = `-- name: DeleteChannel :exec
DELETE FROM Channels
    WHERE Guild_ID = $1
    AND Channel_ID = $2
`

type DeleteChannelParams struct {
	GuildID   sql.NullInt64 `json:"guild_id"`
	ChannelID uint64        `json:"channel_id"`
}

func (q *Queries) DeleteChannel(ctx context.Context, arg DeleteChannelParams) error {
	_, err := q.exec(ctx, q.deleteChannelStmt, deleteChannel, arg.GuildID, arg.ChannelID)
	return err
}

const deleteGuild = `-- name: DeleteGuild :exec
DELETE FROM Guilds
    WHERE Guild_ID = $1
`

func (q *Queries) DeleteGuild(ctx context.Context, guildID uint64) error {
	_, err := q.exec(ctx, q.deleteGuildStmt, deleteGuild, guildID)
	return err
}

const getChannels = `-- name: GetChannels :many
SELECT channel_id, guild_id, channel_name FROM Channels
    WHERE Guild_ID = $1
`

func (q *Queries) GetChannels(ctx context.Context, guildID sql.NullInt64) ([]Channel, error) {
	rows, err := q.query(ctx, q.getChannelsStmt, getChannels, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Channel
	for rows.Next() {
		var i Channel
		if err := rows.Scan(&i.ChannelID, &i.GuildID, &i.ChannelName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGuildData = `-- name: GetGuildData :one
SELECT guild_id, owner_id, guild_name, picture_url FROM Guilds
    WHERE Guild_ID = $1
`

func (q *Queries) GetGuildData(ctx context.Context, guildID uint64) (Guild, error) {
	row := q.queryRow(ctx, q.getGuildDataStmt, getGuildData, guildID)
	var i Guild
	err := row.Scan(
		&i.GuildID,
		&i.OwnerID,
		&i.GuildName,
		&i.PictureUrl,
	)
	return i, err
}

const getGuildMembers = `-- name: GetGuildMembers :many
SELECT User_ID FROM Guild_Members
    WHERE Guild_ID = $1
`

func (q *Queries) GetGuildMembers(ctx context.Context, guildID uint64) ([]uint64, error) {
	rows, err := q.query(ctx, q.getGuildMembersStmt, getGuildMembers, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uint64
	for rows.Next() {
		var user_id uint64
		if err := rows.Scan(&user_id); err != nil {
			return nil, err
		}
		items = append(items, user_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGuildOwner = `-- name: GetGuildOwner :one
SELECT Owner_ID FROM GUILDS
    WHERE Guild_ID = $1
`

func (q *Queries) GetGuildOwner(ctx context.Context, guildID uint64) (uint64, error) {
	row := q.queryRow(ctx, q.getGuildOwnerStmt, getGuildOwner, guildID)
	var owner_id uint64
	err := row.Scan(&owner_id)
	return owner_id, err
}

const getGuildPicture = `-- name: GetGuildPicture :one
SELECT Picture_URL FROM Guilds
    WHERE Guild_ID = $1
`

func (q *Queries) GetGuildPicture(ctx context.Context, guildID uint64) (string, error) {
	row := q.queryRow(ctx, q.getGuildPictureStmt, getGuildPicture, guildID)
	var picture_url string
	err := row.Scan(&picture_url)
	return picture_url, err
}

const guildWithIDExists = `-- name: GuildWithIDExists :one
SELECT EXISTS (
    SELECT 1 FROM Guilds
             WHERE Guild_ID = $1
)
`

func (q *Queries) GuildWithIDExists(ctx context.Context, guildID uint64) (bool, error) {
	row := q.queryRow(ctx, q.guildWithIDExistsStmt, guildWithIDExists, guildID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const guildsForUser = `-- name: GuildsForUser :many
SELECT Guilds.Guild_ID FROM Guild_Members
    INNER JOIN guilds
    ON Guild_Members.Guild_ID = Guilds.Guild_ID
    WHERE User_ID = $1
`

func (q *Queries) GuildsForUser(ctx context.Context, userID uint64) ([]uint64, error) {
	rows, err := q.query(ctx, q.guildsForUserStmt, guildsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uint64
	for rows.Next() {
		var guild_id uint64
		if err := rows.Scan(&guild_id); err != nil {
			return nil, err
		}
		items = append(items, guild_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const guildsForUserWithData = `-- name: GuildsForUserWithData :many
SELECT user_id, guild_members.guild_id, guilds.guild_id, owner_id, guild_name, picture_url FROM Guild_Members
    INNER JOIN guilds
    ON Guild_Members.Guild_ID = Guilds.Guild_ID
    WHERE User_ID = $1
`

type GuildsForUserWithDataRow struct {
	UserID     uint64 `json:"user_id"`
	GuildID    uint64 `json:"guild_id"`
	GuildID_2  uint64 `json:"guild_id_2"`
	OwnerID    uint64 `json:"owner_id"`
	GuildName  string `json:"guild_name"`
	PictureUrl string `json:"picture_url"`
}

func (q *Queries) GuildsForUserWithData(ctx context.Context, userID uint64) ([]GuildsForUserWithDataRow, error) {
	rows, err := q.query(ctx, q.guildsForUserWithDataStmt, guildsForUserWithData, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GuildsForUserWithDataRow
	for rows.Next() {
		var i GuildsForUserWithDataRow
		if err := rows.Scan(
			&i.UserID,
			&i.GuildID,
			&i.GuildID_2,
			&i.OwnerID,
			&i.GuildName,
			&i.PictureUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const numChannelsWithID = `-- name: NumChannelsWithID :one
SELECT COUNT(*) FROM Channels
    WHERE Guild_ID = $1
      AND Channel_ID = $2
`

type NumChannelsWithIDParams struct {
	GuildID   sql.NullInt64 `json:"guild_id"`
	ChannelID uint64        `json:"channel_id"`
}

func (q *Queries) NumChannelsWithID(ctx context.Context, arg NumChannelsWithIDParams) (int64, error) {
	row := q.queryRow(ctx, q.numChannelsWithIDStmt, numChannelsWithID, arg.GuildID, arg.ChannelID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const removeUserFromGuild = `-- name: RemoveUserFromGuild :exec
DELETE FROM Guild_Members
    WHERE Guild_ID = $1
      AND User_ID = $2
`

type RemoveUserFromGuildParams struct {
	GuildID uint64 `json:"guild_id"`
	UserID  uint64 `json:"user_id"`
}

func (q *Queries) RemoveUserFromGuild(ctx context.Context, arg RemoveUserFromGuildParams) error {
	_, err := q.exec(ctx, q.removeUserFromGuildStmt, removeUserFromGuild, arg.GuildID, arg.UserID)
	return err
}

const setGuildName = `-- name: SetGuildName :exec
UPDATE Guilds
    SET Guild_Name = $1
    WHERE Guild_ID = $2
`

type SetGuildNameParams struct {
	GuildName string `json:"guild_name"`
	GuildID   uint64 `json:"guild_id"`
}

func (q *Queries) SetGuildName(ctx context.Context, arg SetGuildNameParams) error {
	_, err := q.exec(ctx, q.setGuildNameStmt, setGuildName, arg.GuildName, arg.GuildID)
	return err
}

const setGuildPicture = `-- name: SetGuildPicture :exec
UPDATE Guilds
    SET Picture_URL = $1
    WHERE Guild_ID = $2
`

type SetGuildPictureParams struct {
	PictureUrl string `json:"picture_url"`
	GuildID    uint64 `json:"guild_id"`
}

func (q *Queries) SetGuildPicture(ctx context.Context, arg SetGuildPictureParams) error {
	_, err := q.exec(ctx, q.setGuildPictureStmt, setGuildPicture, arg.PictureUrl, arg.GuildID)
	return err
}

const updateChannelName = `-- name: UpdateChannelName :exec
UPDATE Channels
      SET Channel_Name = $1
    WHERE Guild_ID = $2
      AND Channel_ID = $3
`

type UpdateChannelNameParams struct {
	ChannelName string        `json:"channel_name"`
	GuildID     sql.NullInt64 `json:"guild_id"`
	ChannelID   uint64        `json:"channel_id"`
}

func (q *Queries) UpdateChannelName(ctx context.Context, arg UpdateChannelNameParams) error {
	_, err := q.exec(ctx, q.updateChannelNameStmt, updateChannelName, arg.ChannelName, arg.GuildID, arg.ChannelID)
	return err
}

const userInGuild = `-- name: UserInGuild :one
SELECT User_ID FROM Guild_Members
    WHERE User_ID = $1
    AND Guild_ID = $2
`

type UserInGuildParams struct {
	UserID  uint64 `json:"user_id"`
	GuildID uint64 `json:"guild_id"`
}

func (q *Queries) UserInGuild(ctx context.Context, arg UserInGuildParams) (uint64, error) {
	row := q.queryRow(ctx, q.userInGuildStmt, userInGuild, arg.UserID, arg.GuildID)
	var user_id uint64
	err := row.Scan(&user_id)
	return user_id, err
}
