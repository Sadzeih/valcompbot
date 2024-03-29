// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/pinnedcomment"
	"github.com/Sadzeih/valcompbot/ent/predicate"
)

// PinnedCommentUpdate is the builder for updating PinnedComment entities.
type PinnedCommentUpdate struct {
	config
	hooks    []Hook
	mutation *PinnedCommentMutation
}

// Where appends a list predicates to the PinnedCommentUpdate builder.
func (pcu *PinnedCommentUpdate) Where(ps ...predicate.PinnedComment) *PinnedCommentUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetCommentID sets the "comment_id" field.
func (pcu *PinnedCommentUpdate) SetCommentID(s string) *PinnedCommentUpdate {
	pcu.mutation.SetCommentID(s)
	return pcu
}

// SetParentID sets the "parent_id" field.
func (pcu *PinnedCommentUpdate) SetParentID(s string) *PinnedCommentUpdate {
	pcu.mutation.SetParentID(s)
	return pcu
}

// Mutation returns the PinnedCommentMutation object of the builder.
func (pcu *PinnedCommentUpdate) Mutation() *PinnedCommentMutation {
	return pcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PinnedCommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcu.hooks) == 0 {
		affected, err = pcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PinnedCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcu.mutation = mutation
			affected, err = pcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcu.hooks) - 1; i >= 0; i-- {
			if pcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *PinnedCommentUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *PinnedCommentUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *PinnedCommentUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *PinnedCommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pinnedcomment.Table,
			Columns: pinnedcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pinnedcomment.FieldID,
			},
		},
	}
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.CommentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinnedcomment.FieldCommentID,
		})
	}
	if value, ok := pcu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinnedcomment.FieldParentID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pinnedcomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PinnedCommentUpdateOne is the builder for updating a single PinnedComment entity.
type PinnedCommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PinnedCommentMutation
}

// SetCommentID sets the "comment_id" field.
func (pcuo *PinnedCommentUpdateOne) SetCommentID(s string) *PinnedCommentUpdateOne {
	pcuo.mutation.SetCommentID(s)
	return pcuo
}

// SetParentID sets the "parent_id" field.
func (pcuo *PinnedCommentUpdateOne) SetParentID(s string) *PinnedCommentUpdateOne {
	pcuo.mutation.SetParentID(s)
	return pcuo
}

// Mutation returns the PinnedCommentMutation object of the builder.
func (pcuo *PinnedCommentUpdateOne) Mutation() *PinnedCommentMutation {
	return pcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *PinnedCommentUpdateOne) Select(field string, fields ...string) *PinnedCommentUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated PinnedComment entity.
func (pcuo *PinnedCommentUpdateOne) Save(ctx context.Context) (*PinnedComment, error) {
	var (
		err  error
		node *PinnedComment
	)
	if len(pcuo.hooks) == 0 {
		node, err = pcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PinnedCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcuo.mutation = mutation
			node, err = pcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pcuo.hooks) - 1; i >= 0; i-- {
			if pcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PinnedComment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PinnedCommentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *PinnedCommentUpdateOne) SaveX(ctx context.Context) *PinnedComment {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *PinnedCommentUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *PinnedCommentUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *PinnedCommentUpdateOne) sqlSave(ctx context.Context) (_node *PinnedComment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pinnedcomment.Table,
			Columns: pinnedcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pinnedcomment.FieldID,
			},
		},
	}
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PinnedComment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pinnedcomment.FieldID)
		for _, f := range fields {
			if !pinnedcomment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pinnedcomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.CommentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinnedcomment.FieldCommentID,
		})
	}
	if value, ok := pcuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinnedcomment.FieldParentID,
		})
	}
	_node = &PinnedComment{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pinnedcomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
