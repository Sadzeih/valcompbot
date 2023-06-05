// Code generated by ent, DO NOT EDIT.

package pickemsevent

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the pickemsevent type in the database.
	Label = "pickems_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// Table holds the table name of the pickemsevent in the database.
	Table = "pickems_events"
)

// Columns holds all SQL columns for pickemsevent fields.
var Columns = []string{
	FieldID,
	FieldEventID,
	FieldTimestamp,
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
