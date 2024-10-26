// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/Sadzeih/valcompbot/ent/highlightedcomment"
	"github.com/Sadzeih/valcompbot/ent/pickemsevent"
	"github.com/Sadzeih/valcompbot/ent/pinnedcomment"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
	"github.com/Sadzeih/valcompbot/ent/schema"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	highlightedcommentFields := schema.HighlightedComment{}.Fields()
	_ = highlightedcommentFields
	// highlightedcommentDescID is the schema descriptor for id field.
	highlightedcommentDescID := highlightedcommentFields[0].Descriptor()
	// highlightedcomment.DefaultID holds the default value on creation for the id field.
	highlightedcomment.DefaultID = highlightedcommentDescID.Default.(func() uuid.UUID)
	pickemseventFields := schema.PickemsEvent{}.Fields()
	_ = pickemseventFields
	// pickemseventDescID is the schema descriptor for id field.
	pickemseventDescID := pickemseventFields[0].Descriptor()
	// pickemsevent.DefaultID holds the default value on creation for the id field.
	pickemsevent.DefaultID = pickemseventDescID.Default.(func() uuid.UUID)
	pinnedcommentFields := schema.PinnedComment{}.Fields()
	_ = pinnedcommentFields
	// pinnedcommentDescID is the schema descriptor for id field.
	pinnedcommentDescID := pinnedcommentFields[0].Descriptor()
	// pinnedcomment.DefaultID holds the default value on creation for the id field.
	pinnedcomment.DefaultID = pinnedcommentDescID.Default.(func() uuid.UUID)
	scheduledmatchFields := schema.ScheduledMatch{}.Fields()
	_ = scheduledmatchFields
	// scheduledmatchDescID is the schema descriptor for id field.
	scheduledmatchDescID := scheduledmatchFields[0].Descriptor()
	// scheduledmatch.DefaultID holds the default value on creation for the id field.
	scheduledmatch.DefaultID = scheduledmatchDescID.Default.(func() uuid.UUID)
	trackedeventFields := schema.TrackedEvent{}.Fields()
	_ = trackedeventFields
	// trackedeventDescID is the schema descriptor for id field.
	trackedeventDescID := trackedeventFields[0].Descriptor()
	// trackedevent.DefaultID holds the default value on creation for the id field.
	trackedevent.DefaultID = trackedeventDescID.Default.(func() uuid.UUID)
}
