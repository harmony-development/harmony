package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
		field.String("contenttype"),
		field.Int("size"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
