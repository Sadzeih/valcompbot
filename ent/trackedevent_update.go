// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/predicate"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
)

// TrackedEventUpdate is the builder for updating TrackedEvent entities.
type TrackedEventUpdate struct {
	config
	hooks    []Hook
	mutation *TrackedEventMutation
}

// Where appends a list predicates to the TrackedEventUpdate builder.
func (teu *TrackedEventUpdate) Where(ps ...predicate.TrackedEvent) *TrackedEventUpdate {
	teu.mutation.Where(ps...)
	return teu
}

// SetEventID sets the "event_id" field.
func (teu *TrackedEventUpdate) SetEventID(i int) *TrackedEventUpdate {
	teu.mutation.ResetEventID()
	teu.mutation.SetEventID(i)
	return teu
}

// AddEventID adds i to the "event_id" field.
func (teu *TrackedEventUpdate) AddEventID(i int) *TrackedEventUpdate {
	teu.mutation.AddEventID(i)
	return teu
}

// SetName sets the "name" field.
func (teu *TrackedEventUpdate) SetName(s string) *TrackedEventUpdate {
	teu.mutation.SetName(s)
	return teu
}

// Mutation returns the TrackedEventMutation object of the builder.
func (teu *TrackedEventUpdate) Mutation() *TrackedEventMutation {
	return teu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (teu *TrackedEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(teu.hooks) == 0 {
		affected, err = teu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackedEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			teu.mutation = mutation
			affected, err = teu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(teu.hooks) - 1; i >= 0; i-- {
			if teu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = teu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, teu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (teu *TrackedEventUpdate) SaveX(ctx context.Context) int {
	affected, err := teu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (teu *TrackedEventUpdate) Exec(ctx context.Context) error {
	_, err := teu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teu *TrackedEventUpdate) ExecX(ctx context.Context) {
	if err := teu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (teu *TrackedEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trackedevent.Table,
			Columns: trackedevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: trackedevent.FieldID,
			},
		},
	}
	if ps := teu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teu.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: trackedevent.FieldEventID,
		})
	}
	if value, ok := teu.mutation.AddedEventID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: trackedevent.FieldEventID,
		})
	}
	if value, ok := teu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: trackedevent.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, teu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trackedevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TrackedEventUpdateOne is the builder for updating a single TrackedEvent entity.
type TrackedEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrackedEventMutation
}

// SetEventID sets the "event_id" field.
func (teuo *TrackedEventUpdateOne) SetEventID(i int) *TrackedEventUpdateOne {
	teuo.mutation.ResetEventID()
	teuo.mutation.SetEventID(i)
	return teuo
}

// AddEventID adds i to the "event_id" field.
func (teuo *TrackedEventUpdateOne) AddEventID(i int) *TrackedEventUpdateOne {
	teuo.mutation.AddEventID(i)
	return teuo
}

// SetName sets the "name" field.
func (teuo *TrackedEventUpdateOne) SetName(s string) *TrackedEventUpdateOne {
	teuo.mutation.SetName(s)
	return teuo
}

// Mutation returns the TrackedEventMutation object of the builder.
func (teuo *TrackedEventUpdateOne) Mutation() *TrackedEventMutation {
	return teuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (teuo *TrackedEventUpdateOne) Select(field string, fields ...string) *TrackedEventUpdateOne {
	teuo.fields = append([]string{field}, fields...)
	return teuo
}

// Save executes the query and returns the updated TrackedEvent entity.
func (teuo *TrackedEventUpdateOne) Save(ctx context.Context) (*TrackedEvent, error) {
	var (
		err  error
		node *TrackedEvent
	)
	if len(teuo.hooks) == 0 {
		node, err = teuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackedEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			teuo.mutation = mutation
			node, err = teuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(teuo.hooks) - 1; i >= 0; i-- {
			if teuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = teuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, teuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TrackedEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TrackedEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (teuo *TrackedEventUpdateOne) SaveX(ctx context.Context) *TrackedEvent {
	node, err := teuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (teuo *TrackedEventUpdateOne) Exec(ctx context.Context) error {
	_, err := teuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teuo *TrackedEventUpdateOne) ExecX(ctx context.Context) {
	if err := teuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (teuo *TrackedEventUpdateOne) sqlSave(ctx context.Context) (_node *TrackedEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trackedevent.Table,
			Columns: trackedevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: trackedevent.FieldID,
			},
		},
	}
	id, ok := teuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TrackedEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := teuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trackedevent.FieldID)
		for _, f := range fields {
			if !trackedevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trackedevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := teuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teuo.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: trackedevent.FieldEventID,
		})
	}
	if value, ok := teuo.mutation.AddedEventID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: trackedevent.FieldEventID,
		})
	}
	if value, ok := teuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: trackedevent.FieldName,
		})
	}
	_node = &TrackedEvent{config: teuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, teuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trackedevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
