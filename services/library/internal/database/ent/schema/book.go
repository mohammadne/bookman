package schema

import "entgo.io/ent"

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return nil
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
