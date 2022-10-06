// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/pinnedcomment"
	"github.com/Sadzeih/valcompbot/ent/predicate"
	"github.com/google/uuid"
)

// PinnedCommentQuery is the builder for querying PinnedComment entities.
type PinnedCommentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PinnedComment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PinnedCommentQuery builder.
func (pcq *PinnedCommentQuery) Where(ps ...predicate.PinnedComment) *PinnedCommentQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit adds a limit step to the query.
func (pcq *PinnedCommentQuery) Limit(limit int) *PinnedCommentQuery {
	pcq.limit = &limit
	return pcq
}

// Offset adds an offset step to the query.
func (pcq *PinnedCommentQuery) Offset(offset int) *PinnedCommentQuery {
	pcq.offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PinnedCommentQuery) Unique(unique bool) *PinnedCommentQuery {
	pcq.unique = &unique
	return pcq
}

// Order adds an order step to the query.
func (pcq *PinnedCommentQuery) Order(o ...OrderFunc) *PinnedCommentQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// First returns the first PinnedComment entity from the query.
// Returns a *NotFoundError when no PinnedComment was found.
func (pcq *PinnedCommentQuery) First(ctx context.Context) (*PinnedComment, error) {
	nodes, err := pcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pinnedcomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PinnedCommentQuery) FirstX(ctx context.Context) *PinnedComment {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PinnedComment ID from the query.
// Returns a *NotFoundError when no PinnedComment ID was found.
func (pcq *PinnedCommentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pinnedcomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PinnedCommentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PinnedComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PinnedComment entity is found.
// Returns a *NotFoundError when no PinnedComment entities are found.
func (pcq *PinnedCommentQuery) Only(ctx context.Context) (*PinnedComment, error) {
	nodes, err := pcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pinnedcomment.Label}
	default:
		return nil, &NotSingularError{pinnedcomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PinnedCommentQuery) OnlyX(ctx context.Context) *PinnedComment {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PinnedComment ID in the query.
// Returns a *NotSingularError when more than one PinnedComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PinnedCommentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pinnedcomment.Label}
	default:
		err = &NotSingularError{pinnedcomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PinnedCommentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PinnedComments.
func (pcq *PinnedCommentQuery) All(ctx context.Context) ([]*PinnedComment, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PinnedCommentQuery) AllX(ctx context.Context) []*PinnedComment {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PinnedComment IDs.
func (pcq *PinnedCommentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pcq.Select(pinnedcomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PinnedCommentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PinnedCommentQuery) Count(ctx context.Context) (int, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PinnedCommentQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PinnedCommentQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PinnedCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PinnedCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PinnedCommentQuery) Clone() *PinnedCommentQuery {
	if pcq == nil {
		return nil
	}
	return &PinnedCommentQuery{
		config:     pcq.config,
		limit:      pcq.limit,
		offset:     pcq.offset,
		order:      append([]OrderFunc{}, pcq.order...),
		predicates: append([]predicate.PinnedComment{}, pcq.predicates...),
		// clone intermediate query.
		sql:    pcq.sql.Clone(),
		path:   pcq.path,
		unique: pcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CommentID string `json:"comment_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PinnedComment.Query().
//		GroupBy(pinnedcomment.FieldCommentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PinnedCommentQuery) GroupBy(field string, fields ...string) *PinnedCommentGroupBy {
	grbuild := &PinnedCommentGroupBy{config: pcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcq.sqlQuery(ctx), nil
	}
	grbuild.label = pinnedcomment.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CommentID string `json:"comment_id,omitempty"`
//	}
//
//	client.PinnedComment.Query().
//		Select(pinnedcomment.FieldCommentID).
//		Scan(ctx, &v)
func (pcq *PinnedCommentQuery) Select(fields ...string) *PinnedCommentSelect {
	pcq.fields = append(pcq.fields, fields...)
	selbuild := &PinnedCommentSelect{PinnedCommentQuery: pcq}
	selbuild.label = pinnedcomment.Label
	selbuild.flds, selbuild.scan = &pcq.fields, selbuild.Scan
	return selbuild
}

func (pcq *PinnedCommentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcq.fields {
		if !pinnedcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PinnedCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PinnedComment, error) {
	var (
		nodes = []*PinnedComment{}
		_spec = pcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PinnedComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PinnedComment{config: pcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pcq *PinnedCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.fields
	if len(pcq.fields) > 0 {
		_spec.Unique = pcq.unique != nil && *pcq.unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PinnedCommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcq *PinnedCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pinnedcomment.Table,
			Columns: pinnedcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pinnedcomment.FieldID,
			},
		},
		From:   pcq.sql,
		Unique: true,
	}
	if unique := pcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pinnedcomment.FieldID)
		for i := range fields {
			if fields[i] != pinnedcomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PinnedCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(pinnedcomment.Table)
	columns := pcq.fields
	if len(columns) == 0 {
		columns = pinnedcomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.unique != nil && *pcq.unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PinnedCommentGroupBy is the group-by builder for PinnedComment entities.
type PinnedCommentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PinnedCommentGroupBy) Aggregate(fns ...AggregateFunc) *PinnedCommentGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcgb *PinnedCommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcgb.path(ctx)
	if err != nil {
		return err
	}
	pcgb.sql = query
	return pcgb.sqlScan(ctx, v)
}

func (pcgb *PinnedCommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcgb.fields {
		if !pinnedcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcgb *PinnedCommentGroupBy) sqlQuery() *sql.Selector {
	selector := pcgb.sql.Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcgb.fields)+len(pcgb.fns))
		for _, f := range pcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcgb.fields...)...)
}

// PinnedCommentSelect is the builder for selecting fields of PinnedComment entities.
type PinnedCommentSelect struct {
	*PinnedCommentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PinnedCommentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	pcs.sql = pcs.PinnedCommentQuery.sqlQuery(ctx)
	return pcs.sqlScan(ctx, v)
}

func (pcs *PinnedCommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcs.sql.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}