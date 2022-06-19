// Code generated by entc, DO NOT EDIT.

package trackedevent

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the trackedevent type in the database.
	Label = "tracked_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the trackedevent in the database.
	Table = "tracked_events"
)

// Columns holds all SQL columns for trackedevent fields.
var Columns = []string{
	FieldID,
	FieldEventID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
