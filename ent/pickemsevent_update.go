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
	"github.com/Sadzeih/valcompbot/ent/pickemsevent"
	"github.com/Sadzeih/valcompbot/ent/predicate"
)

// PickemsEventUpdate is the builder for updating PickemsEvent entities.
type PickemsEventUpdate struct {
	config
	hooks    []Hook
	mutation *PickemsEventMutation
}

// Where appends a list predicates to the PickemsEventUpdate builder.
func (peu *PickemsEventUpdate) Where(ps ...predicate.PickemsEvent) *PickemsEventUpdate {
	peu.mutation.Where(ps...)
	return peu
}

// SetEventID sets the "event_id" field.
func (peu *PickemsEventUpdate) SetEventID(i int) *PickemsEventUpdate {
	peu.mutation.ResetEventID()
	peu.mutation.SetEventID(i)
	return peu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (peu *PickemsEventUpdate) SetNillableEventID(i *int) *PickemsEventUpdate {
	if i != nil {
		peu.SetEventID(*i)
	}
	return peu
}

// AddEventID adds i to the "event_id" field.
func (peu *PickemsEventUpdate) AddEventID(i int) *PickemsEventUpdate {
	peu.mutation.AddEventID(i)
	return peu
}

// ClearEventID clears the value of the "event_id" field.
func (peu *PickemsEventUpdate) ClearEventID() *PickemsEventUpdate {
	peu.mutation.ClearEventID()
	return peu
}

// SetTimestamp sets the "timestamp" field.
func (peu *PickemsEventUpdate) SetTimestamp(t time.Time) *PickemsEventUpdate {
	peu.mutation.SetTimestamp(t)
	return peu
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (peu *PickemsEventUpdate) SetNillableTimestamp(t *time.Time) *PickemsEventUpdate {
	if t != nil {
		peu.SetTimestamp(*t)
	}
	return peu
}

// Mutation returns the PickemsEventMutation object of the builder.
func (peu *PickemsEventUpdate) Mutation() *PickemsEventMutation {
	return peu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (peu *PickemsEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, peu.sqlSave, peu.mutation, peu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (peu *PickemsEventUpdate) SaveX(ctx context.Context) int {
	affected, err := peu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (peu *PickemsEventUpdate) Exec(ctx context.Context) error {
	_, err := peu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (peu *PickemsEventUpdate) ExecX(ctx context.Context) {
	if err := peu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (peu *PickemsEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pickemsevent.Table, pickemsevent.Columns, sqlgraph.NewFieldSpec(pickemsevent.FieldID, field.TypeUUID))
	if ps := peu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := peu.mutation.EventID(); ok {
		_spec.SetField(pickemsevent.FieldEventID, field.TypeInt, value)
	}
	if value, ok := peu.mutation.AddedEventID(); ok {
		_spec.AddField(pickemsevent.FieldEventID, field.TypeInt, value)
	}
	if peu.mutation.EventIDCleared() {
		_spec.ClearField(pickemsevent.FieldEventID, field.TypeInt)
	}
	if value, ok := peu.mutation.Timestamp(); ok {
		_spec.SetField(pickemsevent.FieldTimestamp, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, peu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pickemsevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	peu.mutation.done = true
	return n, nil
}

// PickemsEventUpdateOne is the builder for updating a single PickemsEvent entity.
type PickemsEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PickemsEventMutation
}

// SetEventID sets the "event_id" field.
func (peuo *PickemsEventUpdateOne) SetEventID(i int) *PickemsEventUpdateOne {
	peuo.mutation.ResetEventID()
	peuo.mutation.SetEventID(i)
	return peuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (peuo *PickemsEventUpdateOne) SetNillableEventID(i *int) *PickemsEventUpdateOne {
	if i != nil {
		peuo.SetEventID(*i)
	}
	return peuo
}

// AddEventID adds i to the "event_id" field.
func (peuo *PickemsEventUpdateOne) AddEventID(i int) *PickemsEventUpdateOne {
	peuo.mutation.AddEventID(i)
	return peuo
}

// ClearEventID clears the value of the "event_id" field.
func (peuo *PickemsEventUpdateOne) ClearEventID() *PickemsEventUpdateOne {
	peuo.mutation.ClearEventID()
	return peuo
}

// SetTimestamp sets the "timestamp" field.
func (peuo *PickemsEventUpdateOne) SetTimestamp(t time.Time) *PickemsEventUpdateOne {
	peuo.mutation.SetTimestamp(t)
	return peuo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (peuo *PickemsEventUpdateOne) SetNillableTimestamp(t *time.Time) *PickemsEventUpdateOne {
	if t != nil {
		peuo.SetTimestamp(*t)
	}
	return peuo
}

// Mutation returns the PickemsEventMutation object of the builder.
func (peuo *PickemsEventUpdateOne) Mutation() *PickemsEventMutation {
	return peuo.mutation
}

// Where appends a list predicates to the PickemsEventUpdate builder.
func (peuo *PickemsEventUpdateOne) Where(ps ...predicate.PickemsEvent) *PickemsEventUpdateOne {
	peuo.mutation.Where(ps...)
	return peuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (peuo *PickemsEventUpdateOne) Select(field string, fields ...string) *PickemsEventUpdateOne {
	peuo.fields = append([]string{field}, fields...)
	return peuo
}

// Save executes the query and returns the updated PickemsEvent entity.
func (peuo *PickemsEventUpdateOne) Save(ctx context.Context) (*PickemsEvent, error) {
	return withHooks(ctx, peuo.sqlSave, peuo.mutation, peuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (peuo *PickemsEventUpdateOne) SaveX(ctx context.Context) *PickemsEvent {
	node, err := peuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (peuo *PickemsEventUpdateOne) Exec(ctx context.Context) error {
	_, err := peuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (peuo *PickemsEventUpdateOne) ExecX(ctx context.Context) {
	if err := peuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (peuo *PickemsEventUpdateOne) sqlSave(ctx context.Context) (_node *PickemsEvent, err error) {
	_spec := sqlgraph.NewUpdateSpec(pickemsevent.Table, pickemsevent.Columns, sqlgraph.NewFieldSpec(pickemsevent.FieldID, field.TypeUUID))
	id, ok := peuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PickemsEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := peuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pickemsevent.FieldID)
		for _, f := range fields {
			if !pickemsevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pickemsevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := peuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := peuo.mutation.EventID(); ok {
		_spec.SetField(pickemsevent.FieldEventID, field.TypeInt, value)
	}
	if value, ok := peuo.mutation.AddedEventID(); ok {
		_spec.AddField(pickemsevent.FieldEventID, field.TypeInt, value)
	}
	if peuo.mutation.EventIDCleared() {
		_spec.ClearField(pickemsevent.FieldEventID, field.TypeInt)
	}
	if value, ok := peuo.mutation.Timestamp(); ok {
		_spec.SetField(pickemsevent.FieldTimestamp, field.TypeTime, value)
	}
	_node = &PickemsEvent{config: peuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, peuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pickemsevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	peuo.mutation.done = true
	return _node, nil
}
