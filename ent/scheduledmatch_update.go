// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/predicate"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/google/uuid"
)

// ScheduledMatchUpdate is the builder for updating ScheduledMatch entities.
type ScheduledMatchUpdate struct {
	config
	hooks    []Hook
	mutation *ScheduledMatchMutation
}

// Where appends a list predicates to the ScheduledMatchUpdate builder.
func (smu *ScheduledMatchUpdate) Where(ps ...predicate.ScheduledMatch) *ScheduledMatchUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetMatchID sets the "match_id" field.
func (smu *ScheduledMatchUpdate) SetMatchID(s string) *ScheduledMatchUpdate {
	smu.mutation.SetMatchID(s)
	return smu
}

// SetNillableMatchID sets the "match_id" field if the given value is not nil.
func (smu *ScheduledMatchUpdate) SetNillableMatchID(s *string) *ScheduledMatchUpdate {
	if s != nil {
		smu.SetMatchID(*s)
	}
	return smu
}

// SetDoneAt sets the "done_at" field.
func (smu *ScheduledMatchUpdate) SetDoneAt(t time.Time) *ScheduledMatchUpdate {
	smu.mutation.SetDoneAt(t)
	return smu
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (smu *ScheduledMatchUpdate) SetNillableDoneAt(t *time.Time) *ScheduledMatchUpdate {
	if t != nil {
		smu.SetDoneAt(*t)
	}
	return smu
}

// ClearDoneAt clears the value of the "done_at" field.
func (smu *ScheduledMatchUpdate) ClearDoneAt() *ScheduledMatchUpdate {
	smu.mutation.ClearDoneAt()
	return smu
}

// SetPostedAt sets the "posted_at" field.
func (smu *ScheduledMatchUpdate) SetPostedAt(t time.Time) *ScheduledMatchUpdate {
	smu.mutation.SetPostedAt(t)
	return smu
}

// SetNillablePostedAt sets the "posted_at" field if the given value is not nil.
func (smu *ScheduledMatchUpdate) SetNillablePostedAt(t *time.Time) *ScheduledMatchUpdate {
	if t != nil {
		smu.SetPostedAt(*t)
	}
	return smu
}

// ClearPostedAt clears the value of the "posted_at" field.
func (smu *ScheduledMatchUpdate) ClearPostedAt() *ScheduledMatchUpdate {
	smu.mutation.ClearPostedAt()
	return smu
}

// SetEventID sets the "event" edge to the TrackedEvent entity by ID.
func (smu *ScheduledMatchUpdate) SetEventID(id uuid.UUID) *ScheduledMatchUpdate {
	smu.mutation.SetEventID(id)
	return smu
}

// SetEvent sets the "event" edge to the TrackedEvent entity.
func (smu *ScheduledMatchUpdate) SetEvent(t *TrackedEvent) *ScheduledMatchUpdate {
	return smu.SetEventID(t.ID)
}

// Mutation returns the ScheduledMatchMutation object of the builder.
func (smu *ScheduledMatchUpdate) Mutation() *ScheduledMatchMutation {
	return smu.mutation
}

// ClearEvent clears the "event" edge to the TrackedEvent entity.
func (smu *ScheduledMatchUpdate) ClearEvent() *ScheduledMatchUpdate {
	smu.mutation.ClearEvent()
	return smu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *ScheduledMatchUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, smu.sqlSave, smu.mutation, smu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smu *ScheduledMatchUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *ScheduledMatchUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *ScheduledMatchUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smu *ScheduledMatchUpdate) check() error {
	if smu.mutation.EventCleared() && len(smu.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ScheduledMatch.event"`)
	}
	return nil
}

func (smu *ScheduledMatchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := smu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(scheduledmatch.Table, scheduledmatch.Columns, sqlgraph.NewFieldSpec(scheduledmatch.FieldID, field.TypeUUID))
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.MatchID(); ok {
		_spec.SetField(scheduledmatch.FieldMatchID, field.TypeString, value)
	}
	if value, ok := smu.mutation.DoneAt(); ok {
		_spec.SetField(scheduledmatch.FieldDoneAt, field.TypeTime, value)
	}
	if smu.mutation.DoneAtCleared() {
		_spec.ClearField(scheduledmatch.FieldDoneAt, field.TypeTime)
	}
	if value, ok := smu.mutation.PostedAt(); ok {
		_spec.SetField(scheduledmatch.FieldPostedAt, field.TypeTime, value)
	}
	if smu.mutation.PostedAtCleared() {
		_spec.ClearField(scheduledmatch.FieldPostedAt, field.TypeTime)
	}
	if smu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scheduledmatch.EventTable,
			Columns: []string{scheduledmatch.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trackedevent.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scheduledmatch.EventTable,
			Columns: []string{scheduledmatch.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trackedevent.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scheduledmatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	smu.mutation.done = true
	return n, nil
}

// ScheduledMatchUpdateOne is the builder for updating a single ScheduledMatch entity.
type ScheduledMatchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScheduledMatchMutation
}

// SetMatchID sets the "match_id" field.
func (smuo *ScheduledMatchUpdateOne) SetMatchID(s string) *ScheduledMatchUpdateOne {
	smuo.mutation.SetMatchID(s)
	return smuo
}

// SetNillableMatchID sets the "match_id" field if the given value is not nil.
func (smuo *ScheduledMatchUpdateOne) SetNillableMatchID(s *string) *ScheduledMatchUpdateOne {
	if s != nil {
		smuo.SetMatchID(*s)
	}
	return smuo
}

// SetDoneAt sets the "done_at" field.
func (smuo *ScheduledMatchUpdateOne) SetDoneAt(t time.Time) *ScheduledMatchUpdateOne {
	smuo.mutation.SetDoneAt(t)
	return smuo
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (smuo *ScheduledMatchUpdateOne) SetNillableDoneAt(t *time.Time) *ScheduledMatchUpdateOne {
	if t != nil {
		smuo.SetDoneAt(*t)
	}
	return smuo
}

// ClearDoneAt clears the value of the "done_at" field.
func (smuo *ScheduledMatchUpdateOne) ClearDoneAt() *ScheduledMatchUpdateOne {
	smuo.mutation.ClearDoneAt()
	return smuo
}

// SetPostedAt sets the "posted_at" field.
func (smuo *ScheduledMatchUpdateOne) SetPostedAt(t time.Time) *ScheduledMatchUpdateOne {
	smuo.mutation.SetPostedAt(t)
	return smuo
}

// SetNillablePostedAt sets the "posted_at" field if the given value is not nil.
func (smuo *ScheduledMatchUpdateOne) SetNillablePostedAt(t *time.Time) *ScheduledMatchUpdateOne {
	if t != nil {
		smuo.SetPostedAt(*t)
	}
	return smuo
}

// ClearPostedAt clears the value of the "posted_at" field.
func (smuo *ScheduledMatchUpdateOne) ClearPostedAt() *ScheduledMatchUpdateOne {
	smuo.mutation.ClearPostedAt()
	return smuo
}

// SetEventID sets the "event" edge to the TrackedEvent entity by ID.
func (smuo *ScheduledMatchUpdateOne) SetEventID(id uuid.UUID) *ScheduledMatchUpdateOne {
	smuo.mutation.SetEventID(id)
	return smuo
}

// SetEvent sets the "event" edge to the TrackedEvent entity.
func (smuo *ScheduledMatchUpdateOne) SetEvent(t *TrackedEvent) *ScheduledMatchUpdateOne {
	return smuo.SetEventID(t.ID)
}

// Mutation returns the ScheduledMatchMutation object of the builder.
func (smuo *ScheduledMatchUpdateOne) Mutation() *ScheduledMatchMutation {
	return smuo.mutation
}

// ClearEvent clears the "event" edge to the TrackedEvent entity.
func (smuo *ScheduledMatchUpdateOne) ClearEvent() *ScheduledMatchUpdateOne {
	smuo.mutation.ClearEvent()
	return smuo
}

// Where appends a list predicates to the ScheduledMatchUpdate builder.
func (smuo *ScheduledMatchUpdateOne) Where(ps ...predicate.ScheduledMatch) *ScheduledMatchUpdateOne {
	smuo.mutation.Where(ps...)
	return smuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *ScheduledMatchUpdateOne) Select(field string, fields ...string) *ScheduledMatchUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated ScheduledMatch entity.
func (smuo *ScheduledMatchUpdateOne) Save(ctx context.Context) (*ScheduledMatch, error) {
	return withHooks(ctx, smuo.sqlSave, smuo.mutation, smuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *ScheduledMatchUpdateOne) SaveX(ctx context.Context) *ScheduledMatch {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *ScheduledMatchUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *ScheduledMatchUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smuo *ScheduledMatchUpdateOne) check() error {
	if smuo.mutation.EventCleared() && len(smuo.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ScheduledMatch.event"`)
	}
	return nil
}

func (smuo *ScheduledMatchUpdateOne) sqlSave(ctx context.Context) (_node *ScheduledMatch, err error) {
	if err := smuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(scheduledmatch.Table, scheduledmatch.Columns, sqlgraph.NewFieldSpec(scheduledmatch.FieldID, field.TypeUUID))
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ScheduledMatch.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scheduledmatch.FieldID)
		for _, f := range fields {
			if !scheduledmatch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != scheduledmatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.MatchID(); ok {
		_spec.SetField(scheduledmatch.FieldMatchID, field.TypeString, value)
	}
	if value, ok := smuo.mutation.DoneAt(); ok {
		_spec.SetField(scheduledmatch.FieldDoneAt, field.TypeTime, value)
	}
	if smuo.mutation.DoneAtCleared() {
		_spec.ClearField(scheduledmatch.FieldDoneAt, field.TypeTime)
	}
	if value, ok := smuo.mutation.PostedAt(); ok {
		_spec.SetField(scheduledmatch.FieldPostedAt, field.TypeTime, value)
	}
	if smuo.mutation.PostedAtCleared() {
		_spec.ClearField(scheduledmatch.FieldPostedAt, field.TypeTime)
	}
	if smuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scheduledmatch.EventTable,
			Columns: []string{scheduledmatch.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trackedevent.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scheduledmatch.EventTable,
			Columns: []string{scheduledmatch.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trackedevent.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ScheduledMatch{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scheduledmatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	smuo.mutation.done = true
	return _node, nil
}