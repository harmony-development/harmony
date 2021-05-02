// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/foreignuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/profile"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges      UserEdges `json:"edges"`
	guild_bans *uint64
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// LocalUser holds the value of the local_user edge.
	LocalUser *LocalUser `json:"local_user,omitempty"`
	// ForeignUser holds the value of the foreign_user edge.
	ForeignUser *ForeignUser `json:"foreign_user,omitempty"`
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata []*UserMeta `json:"metadata,omitempty"`
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// Message holds the value of the message edge.
	Message []*Message `json:"message,omitempty"`
	// Guild holds the value of the guild edge.
	Guild []*Guild `json:"guild,omitempty"`
	// Emotepack holds the value of the emotepack edge.
	Emotepack []*EmotePack `json:"emotepack,omitempty"`
	// Createdpacks holds the value of the createdpacks edge.
	Createdpacks []*EmotePack `json:"createdpacks,omitempty"`
	// Listentry holds the value of the listentry edge.
	Listentry []*GuildListEntry `json:"listentry,omitempty"`
	// Role holds the value of the role edge.
	Role []*Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [11]bool
}

// LocalUserOrErr returns the LocalUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) LocalUserOrErr() (*LocalUser, error) {
	if e.loadedTypes[0] {
		if e.LocalUser == nil {
			// The edge local_user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: localuser.Label}
		}
		return e.LocalUser, nil
	}
	return nil, &NotLoadedError{edge: "local_user"}
}

// ForeignUserOrErr returns the ForeignUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ForeignUserOrErr() (*ForeignUser, error) {
	if e.loadedTypes[1] {
		if e.ForeignUser == nil {
			// The edge foreign_user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: foreignuser.Label}
		}
		return e.ForeignUser, nil
	}
	return nil, &NotLoadedError{edge: "foreign_user"}
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[2] {
		if e.Profile == nil {
			// The edge profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MetadataOrErr() ([]*UserMeta, error) {
	if e.loadedTypes[3] {
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[4] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// MessageOrErr returns the Message value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessageOrErr() ([]*Message, error) {
	if e.loadedTypes[5] {
		return e.Message, nil
	}
	return nil, &NotLoadedError{edge: "message"}
}

// GuildOrErr returns the Guild value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GuildOrErr() ([]*Guild, error) {
	if e.loadedTypes[6] {
		return e.Guild, nil
	}
	return nil, &NotLoadedError{edge: "guild"}
}

// EmotepackOrErr returns the Emotepack value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EmotepackOrErr() ([]*EmotePack, error) {
	if e.loadedTypes[7] {
		return e.Emotepack, nil
	}
	return nil, &NotLoadedError{edge: "emotepack"}
}

// CreatedpacksOrErr returns the Createdpacks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedpacksOrErr() ([]*EmotePack, error) {
	if e.loadedTypes[8] {
		return e.Createdpacks, nil
	}
	return nil, &NotLoadedError{edge: "createdpacks"}
}

// ListentryOrErr returns the Listentry value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ListentryOrErr() ([]*GuildListEntry, error) {
	if e.loadedTypes[9] {
		return e.Listentry, nil
	}
	return nil, &NotLoadedError{edge: "listentry"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoleOrErr() ([]*Role, error) {
	if e.loadedTypes[10] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[0]: // guild_bans
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint64(value.Int64)
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field guild_bans", value)
			} else if value.Valid {
				u.guild_bans = new(uint64)
				*u.guild_bans = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryLocalUser queries the "local_user" edge of the User entity.
func (u *User) QueryLocalUser() *LocalUserQuery {
	return (&UserClient{config: u.config}).QueryLocalUser(u)
}

// QueryForeignUser queries the "foreign_user" edge of the User entity.
func (u *User) QueryForeignUser() *ForeignUserQuery {
	return (&UserClient{config: u.config}).QueryForeignUser(u)
}

// QueryProfile queries the "profile" edge of the User entity.
func (u *User) QueryProfile() *ProfileQuery {
	return (&UserClient{config: u.config}).QueryProfile(u)
}

// QueryMetadata queries the "metadata" edge of the User entity.
func (u *User) QueryMetadata() *UserMetaQuery {
	return (&UserClient{config: u.config}).QueryMetadata(u)
}

// QuerySessions queries the "sessions" edge of the User entity.
func (u *User) QuerySessions() *SessionQuery {
	return (&UserClient{config: u.config}).QuerySessions(u)
}

// QueryMessage queries the "message" edge of the User entity.
func (u *User) QueryMessage() *MessageQuery {
	return (&UserClient{config: u.config}).QueryMessage(u)
}

// QueryGuild queries the "guild" edge of the User entity.
func (u *User) QueryGuild() *GuildQuery {
	return (&UserClient{config: u.config}).QueryGuild(u)
}

// QueryEmotepack queries the "emotepack" edge of the User entity.
func (u *User) QueryEmotepack() *EmotePackQuery {
	return (&UserClient{config: u.config}).QueryEmotepack(u)
}

// QueryCreatedpacks queries the "createdpacks" edge of the User entity.
func (u *User) QueryCreatedpacks() *EmotePackQuery {
	return (&UserClient{config: u.config}).QueryCreatedpacks(u)
}

// QueryListentry queries the "listentry" edge of the User entity.
func (u *User) QueryListentry() *GuildListEntryQuery {
	return (&UserClient{config: u.config}).QueryListentry(u)
}

// QueryRole queries the "role" edge of the User entity.
func (u *User) QueryRole() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRole(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entgen: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
