package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PickemsEvent holds the schema definition for the TrackedEvent entity.
type PickemsEvent struct {
	ent.Schema
}

// Fields of the PickemsEvent.
func (PickemsEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Int("event_id").Nillable().Optional(),
		field.Time("timestamp"),
	}
}

// Edges of the PickemsEvent.
func (PickemsEvent) Edges() []ent.Edge {
	return nil
}
