package schema

import "entgo.io/ent"

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return nil
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return nil
}
