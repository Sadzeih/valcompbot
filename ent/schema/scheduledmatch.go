package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ScheduledMatch holds the schema definition for the ScheduledMatch entity.
type ScheduledMatch struct {
	ent.Schema
}

// Fields of the ScheduledMatch.
func (ScheduledMatch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("match_id").
			Unique(),
		field.Time("done_at").Optional().Nillable(),
		field.Time("posted_at").Optional().Nillable(),
	}
}

// Edges of the ScheduledMatch.
func (ScheduledMatch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", TrackedEvent.Type).
			Ref("scheduledmatches").
			Unique().
			Required(),
	}
}

func (ScheduledMatch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("match_id"),
	}
}
