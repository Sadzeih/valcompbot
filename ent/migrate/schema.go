// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TrackedEventsColumns holds the columns for the "tracked_events" table.
	TrackedEventsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "event_id", Type: field.TypeInt, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// TrackedEventsTable holds the schema information for the "tracked_events" table.
	TrackedEventsTable = &schema.Table{
		Name:       "tracked_events",
		Columns:    TrackedEventsColumns,
		PrimaryKey: []*schema.Column{TrackedEventsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TrackedEventsTable,
	}
)

func init() {
}
