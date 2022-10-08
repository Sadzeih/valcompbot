package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PinnedComment holds the schema definition for the PinnedComment entity.
type PinnedComment struct {
	ent.Schema
}

// Fields of the PinnedComment.
func (PinnedComment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("comment_id").Unique(),
		field.String("parent_id"),
	}
}

// Edges of the PinnedComment.
func (PinnedComment) Edges() []ent.Edge {
	return nil
}
