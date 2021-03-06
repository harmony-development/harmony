// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// EmotePack is the model entity for the EmotePack schema.
type EmotePack struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmotePackQuery when eager-loading is set.
	Edges             EmotePackEdges `json:"edges"`
	user_emotepack    *uint64
	user_createdpacks *uint64
}

// EmotePackEdges holds the relations/edges for other nodes in the graph.
type EmotePackEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Emote holds the value of the emote edge.
	Emote []*Emote `json:"emote,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmotePackEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmotePackEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// EmoteOrErr returns the Emote value or an error if the edge
// was not loaded in eager-loading.
func (e EmotePackEdges) EmoteOrErr() ([]*Emote, error) {
	if e.loadedTypes[2] {
		return e.Emote, nil
	}
	return nil, &NotLoadedError{edge: "emote"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmotePack) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case emotepack.FieldID:
			values[i] = new(sql.NullInt64)
		case emotepack.FieldName:
			values[i] = new(sql.NullString)
		case emotepack.ForeignKeys[0]: // user_emotepack
			values[i] = new(sql.NullInt64)
		case emotepack.ForeignKeys[1]: // user_createdpacks
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EmotePack", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmotePack fields.
func (ep *EmotePack) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emotepack.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ep.ID = uint64(value.Int64)
		case emotepack.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ep.Name = value.String
			}
		case emotepack.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_emotepack", value)
			} else if value.Valid {
				ep.user_emotepack = new(uint64)
				*ep.user_emotepack = uint64(value.Int64)
			}
		case emotepack.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_createdpacks", value)
			} else if value.Valid {
				ep.user_createdpacks = new(uint64)
				*ep.user_createdpacks = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the EmotePack entity.
func (ep *EmotePack) QueryUser() *UserQuery {
	return (&EmotePackClient{config: ep.config}).QueryUser(ep)
}

// QueryOwner queries the "owner" edge of the EmotePack entity.
func (ep *EmotePack) QueryOwner() *UserQuery {
	return (&EmotePackClient{config: ep.config}).QueryOwner(ep)
}

// QueryEmote queries the "emote" edge of the EmotePack entity.
func (ep *EmotePack) QueryEmote() *EmoteQuery {
	return (&EmotePackClient{config: ep.config}).QueryEmote(ep)
}

// Update returns a builder for updating this EmotePack.
// Note that you need to call EmotePack.Unwrap() before calling this method if this EmotePack
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EmotePack) Update() *EmotePackUpdateOne {
	return (&EmotePackClient{config: ep.config}).UpdateOne(ep)
}

// Unwrap unwraps the EmotePack entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *EmotePack) Unwrap() *EmotePack {
	tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("entgen: EmotePack is not a transactional entity")
	}
	ep.config.driver = tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EmotePack) String() string {
	var builder strings.Builder
	builder.WriteString("EmotePack(")
	builder.WriteString(fmt.Sprintf("id=%v", ep.ID))
	builder.WriteString(", name=")
	builder.WriteString(ep.Name)
	builder.WriteByte(')')
	return builder.String()
}

// EmotePacks is a parsable slice of EmotePack.
type EmotePacks []*EmotePack

func (ep EmotePacks) config(cfg config) {
	for _i := range ep {
		ep[_i].config = cfg
	}
}
