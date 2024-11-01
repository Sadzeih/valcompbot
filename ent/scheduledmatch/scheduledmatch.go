// Code generated by ent, DO NOT EDIT.

package scheduledmatch

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the scheduledmatch type in the database.
	Label = "scheduled_match"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldMatchID holds the string denoting the match_id field in the database.
	FieldMatchID = "match_id"
	// FieldDoneAt holds the string denoting the done_at field in the database.
	FieldDoneAt = "done_at"
	// FieldPostedAt holds the string denoting the posted_at field in the database.
	FieldPostedAt = "posted_at"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the scheduledmatch in the database.
	Table = "scheduled_matches"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "scheduled_matches"
	// EventInverseTable is the table name for the TrackedEvent entity.
	// It exists in this package in order to avoid circular dependency with the "trackedevent" package.
	EventInverseTable = "tracked_events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "tracked_event_scheduledmatches"
)

// Columns holds all SQL columns for scheduledmatch fields.
var Columns = []string{
	FieldID,
	FieldMatchID,
	FieldDoneAt,
	FieldPostedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "scheduled_matches"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tracked_event_scheduledmatches",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the ScheduledMatch queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMatchID orders the results by the match_id field.
func ByMatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMatchID, opts...).ToFunc()
}

// ByDoneAt orders the results by the done_at field.
func ByDoneAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDoneAt, opts...).ToFunc()
}

// ByPostedAt orders the results by the posted_at field.
func ByPostedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostedAt, opts...).ToFunc()
}

// ByEventField orders the results by event field.
func ByEventField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventStep(), sql.OrderByField(field, opts...))
	}
}
func newEventStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
	)
}
