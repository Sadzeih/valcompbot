package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TrackedEvent holds the schema definition for the TrackedEvent entity.
type TrackedEvent struct {
	ent.Schema
}

// Fields of the TrackedEvent.
func (TrackedEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Int("event_id").
			Unique(),
		field.String("name"),
	}
}

// Edges of the TrackedEvent.
func (TrackedEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("scheduledmatches", ScheduledMatch.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
