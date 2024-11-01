// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/pickemsevent"
	"github.com/google/uuid"
)

// PickemsEventCreate is the builder for creating a PickemsEvent entity.
type PickemsEventCreate struct {
	config
	mutation *PickemsEventMutation
	hooks    []Hook
}

// SetEventID sets the "event_id" field.
func (pec *PickemsEventCreate) SetEventID(i int) *PickemsEventCreate {
	pec.mutation.SetEventID(i)
	return pec
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (pec *PickemsEventCreate) SetNillableEventID(i *int) *PickemsEventCreate {
	if i != nil {
		pec.SetEventID(*i)
	}
	return pec
}

// SetTimestamp sets the "timestamp" field.
func (pec *PickemsEventCreate) SetTimestamp(t time.Time) *PickemsEventCreate {
	pec.mutation.SetTimestamp(t)
	return pec
}

// SetID sets the "id" field.
func (pec *PickemsEventCreate) SetID(u uuid.UUID) *PickemsEventCreate {
	pec.mutation.SetID(u)
	return pec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pec *PickemsEventCreate) SetNillableID(u *uuid.UUID) *PickemsEventCreate {
	if u != nil {
		pec.SetID(*u)
	}
	return pec
}

// Mutation returns the PickemsEventMutation object of the builder.
func (pec *PickemsEventCreate) Mutation() *PickemsEventMutation {
	return pec.mutation
}

// Save creates the PickemsEvent in the database.
func (pec *PickemsEventCreate) Save(ctx context.Context) (*PickemsEvent, error) {
	pec.defaults()
	return withHooks(ctx, pec.sqlSave, pec.mutation, pec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pec *PickemsEventCreate) SaveX(ctx context.Context) *PickemsEvent {
	v, err := pec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pec *PickemsEventCreate) Exec(ctx context.Context) error {
	_, err := pec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pec *PickemsEventCreate) ExecX(ctx context.Context) {
	if err := pec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pec *PickemsEventCreate) defaults() {
	if _, ok := pec.mutation.ID(); !ok {
		v := pickemsevent.DefaultID()
		pec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pec *PickemsEventCreate) check() error {
	if _, ok := pec.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "PickemsEvent.timestamp"`)}
	}
	return nil
}

func (pec *PickemsEventCreate) sqlSave(ctx context.Context) (*PickemsEvent, error) {
	if err := pec.check(); err != nil {
		return nil, err
	}
	_node, _spec := pec.createSpec()
	if err := sqlgraph.CreateNode(ctx, pec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pec.mutation.id = &_node.ID
	pec.mutation.done = true
	return _node, nil
}

func (pec *PickemsEventCreate) createSpec() (*PickemsEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &PickemsEvent{config: pec.config}
		_spec = sqlgraph.NewCreateSpec(pickemsevent.Table, sqlgraph.NewFieldSpec(pickemsevent.FieldID, field.TypeUUID))
	)
	if id, ok := pec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pec.mutation.EventID(); ok {
		_spec.SetField(pickemsevent.FieldEventID, field.TypeInt, value)
		_node.EventID = &value
	}
	if value, ok := pec.mutation.Timestamp(); ok {
		_spec.SetField(pickemsevent.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	return _node, _spec
}

// PickemsEventCreateBulk is the builder for creating many PickemsEvent entities in bulk.
type PickemsEventCreateBulk struct {
	config
	err      error
	builders []*PickemsEventCreate
}

// Save creates the PickemsEvent entities in the database.
func (pecb *PickemsEventCreateBulk) Save(ctx context.Context) ([]*PickemsEvent, error) {
	if pecb.err != nil {
		return nil, pecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pecb.builders))
	nodes := make([]*PickemsEvent, len(pecb.builders))
	mutators := make([]Mutator, len(pecb.builders))
	for i := range pecb.builders {
		func(i int, root context.Context) {
			builder := pecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PickemsEventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pecb *PickemsEventCreateBulk) SaveX(ctx context.Context) []*PickemsEvent {
	v, err := pecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pecb *PickemsEventCreateBulk) Exec(ctx context.Context) error {
	_, err := pecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pecb *PickemsEventCreateBulk) ExecX(ctx context.Context) {
	if err := pecb.Exec(ctx); err != nil {
		panic(err)
	}
}
