package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HighlightedComment holds the schema definition for the HighlightedComment entity.
type HighlightedComment struct {
	ent.Schema
}

// Fields of the HighlightedComment.
func (HighlightedComment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("comment_id").Unique(),
		field.String("body"),
		field.String("author"),
		field.String("flair"),
		field.String("parent_id"),
		field.String("link"),
		field.String("author_type"),
		field.Time("timestamp"),
	}
}

// Edges of the HighlightedComment.
func (HighlightedComment) Edges() []ent.Edge {
	return nil
}
