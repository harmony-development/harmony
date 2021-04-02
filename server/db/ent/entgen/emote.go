// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/emote"
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
)

// Emote is the model entity for the Emote schema.
type Emote struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmoteQuery when eager-loading is set.
	Edges            EmoteEdges `json:"edges"`
	emote_pack_emote *uint64
}

// EmoteEdges holds the relations/edges for other nodes in the graph.
type EmoteEdges struct {
	// Emotepack holds the value of the emotepack edge.
	Emotepack *EmotePack `json:"emotepack,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmotepackOrErr returns the Emotepack value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmoteEdges) EmotepackOrErr() (*EmotePack, error) {
	if e.loadedTypes[0] {
		if e.Emotepack == nil {
			// The edge emotepack was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: emotepack.Label}
		}
		return e.Emotepack, nil
	}
	return nil, &NotLoadedError{edge: "emotepack"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Emote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case emote.FieldID, emote.FieldName:
			values[i] = &sql.NullString{}
		case emote.ForeignKeys[0]: // emote_pack_emote
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Emote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Emote fields.
func (e *Emote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emote.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case emote.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case emote.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field emote_pack_emote", value)
			} else if value.Valid {
				e.emote_pack_emote = new(uint64)
				*e.emote_pack_emote = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryEmotepack queries the "emotepack" edge of the Emote entity.
func (e *Emote) QueryEmotepack() *EmotePackQuery {
	return (&EmoteClient{config: e.config}).QueryEmotepack(e)
}

// Update returns a builder for updating this Emote.
// Note that you need to call Emote.Unwrap() before calling this method if this Emote
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Emote) Update() *EmoteUpdateOne {
	return (&EmoteClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Emote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Emote) Unwrap() *Emote {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Emote is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Emote) String() string {
	var builder strings.Builder
	builder.WriteString("Emote(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Emotes is a parsable slice of Emote.
type Emotes []*Emote

func (e Emotes) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
