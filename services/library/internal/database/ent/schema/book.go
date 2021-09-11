package schema

import "entgo.io/ent"

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return nil

	// return []ent.Field{
	// 	field.Uint64("caller_id").
	// 		Optional(),
	// 	field.Uint64("callee_id").
	// 		Optional(),
	// 	field.Enum("caller_audience").
	// 		Values("driver", "passenger", "third-party").
	// 		Optional(),
	// 	field.Enum("callee_audience").
	// 		Values("driver", "passenger", "third-party").
	// 		Optional(),
	// 	field.Text("caller_sdp_offer").
	// 		Optional(),
	// 	field.Text("callee_sdp_offer").
	// 		Optional(),
	// 	field.Float("duration").
	// 		Optional(),
	// 	field.Int("state").
	// 		Default(0).
	// 		Min(0).
	// 		Max(5),
	// 	field.String("audio_file").
	// 		Optional(),
	// 	field.String("media_server").
	// 		Optional(),
	// 	field.Time("created_at").
	// 		Default(func() time.Time {
	// 			return time.Now().UTC()
	// 		}),
	// }
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
