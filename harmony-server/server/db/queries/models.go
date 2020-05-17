// Code generated by sqlc. DO NOT EDIT.

package queries

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Attachment struct {
	MessageID     int64  `json:"message_id"`
	AttachmentUrl string `json:"attachment_url"`
}

type Channel struct {
	ChannelID   int64         `json:"channel_id"`
	GuildID     sql.NullInt64 `json:"guild_id"`
	ChannelName string        `json:"channel_name"`
}

type Guild struct {
	GuildID    int64  `json:"guild_id"`
	OwnerID    int64  `json:"owner_id"`
	GuildName  string `json:"guild_name"`
	PictureUrl string `json:"picture_url"`
}

type GuildMember struct {
	UserID  int64 `json:"user_id"`
	GuildID int64 `json:"guild_id"`
}

type Invite struct {
	InviteID     int64         `json:"invite_id"`
	Name         string        `json:"name"`
	Uses         int32         `json:"uses"`
	PossibleUses sql.NullInt32 `json:"possible_uses"`
	GuildID      int64         `json:"guild_id"`
}

type Message struct {
	MessageID int64             `json:"message_id"`
	GuildID   int64             `json:"guild_id"`
	ChannelID int64             `json:"channel_id"`
	UserID    int64             `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	EditedAt  sql.NullTime      `json:"edited_at"`
	Content   string            `json:"content"`
	Embeds    []json.RawMessage `json:"embeds"`
	Actions   []json.RawMessage `json:"actions"`
}

type Session struct {
	UserID  int64  `json:"user_id"`
	Session string `json:"session"`
}

type User struct {
	UserID   int64          `json:"user_id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
	Password []byte         `json:"password"`
}
