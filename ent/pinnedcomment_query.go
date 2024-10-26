// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
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
	ctx        *QueryContext
	order      []pinnedcomment.OrderOption
	inters     []Interceptor
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

// Limit the number of records to be returned by this query.
func (pcq *PinnedCommentQuery) Limit(limit int) *PinnedCommentQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PinnedCommentQuery) Offset(offset int) *PinnedCommentQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PinnedCommentQuery) Unique(unique bool) *PinnedCommentQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PinnedCommentQuery) Order(o ...pinnedcomment.OrderOption) *PinnedCommentQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// First returns the first PinnedComment entity from the query.
// Returns a *NotFoundError when no PinnedComment was found.
func (pcq *PinnedCommentQuery) First(ctx context.Context) (*PinnedComment, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, ent.OpQueryFirst))
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
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, ent.OpQueryFirstID)); err != nil {
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
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, ent.OpQueryOnly))
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
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, ent.OpQueryOnlyID)); err != nil {
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
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryAll)
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PinnedComment, *PinnedCommentQuery]()
	return withInterceptors[[]*PinnedComment](ctx, pcq, qr, pcq.inters)
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
func (pcq *PinnedCommentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryIDs)
	if err = pcq.Select(pinnedcomment.FieldID).Scan(ctx, &ids); err != nil {
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
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryCount)
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PinnedCommentQuery](), pcq.inters)
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
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryExist)
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
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
		ctx:        pcq.ctx.Clone(),
		order:      append([]pinnedcomment.OrderOption{}, pcq.order...),
		inters:     append([]Interceptor{}, pcq.inters...),
		predicates: append([]predicate.PinnedComment{}, pcq.predicates...),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
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
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PinnedCommentGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = pinnedcomment.Label
	grbuild.scan = grbuild.Scan
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
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PinnedCommentSelect{PinnedCommentQuery: pcq}
	sbuild.label = pinnedcomment.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PinnedCommentSelect configured with the given aggregations.
func (pcq *PinnedCommentQuery) Aggregate(fns ...AggregateFunc) *PinnedCommentSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PinnedCommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
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
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PinnedComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
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
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PinnedCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pinnedcomment.Table, pinnedcomment.Columns, sqlgraph.NewFieldSpec(pinnedcomment.FieldID, field.TypeUUID))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
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
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
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
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = pinnedcomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PinnedCommentGroupBy is the group-by builder for PinnedComment entities.
type PinnedCommentGroupBy struct {
	selector
	build *PinnedCommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PinnedCommentGroupBy) Aggregate(fns ...AggregateFunc) *PinnedCommentGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PinnedCommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, ent.OpQueryGroupBy)
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PinnedCommentQuery, *PinnedCommentGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PinnedCommentGroupBy) sqlScan(ctx context.Context, root *PinnedCommentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PinnedCommentSelect is the builder for selecting fields of PinnedComment entities.
type PinnedCommentSelect struct {
	*PinnedCommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PinnedCommentSelect) Aggregate(fns ...AggregateFunc) *PinnedCommentSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PinnedCommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, ent.OpQuerySelect)
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PinnedCommentQuery, *PinnedCommentSelect](ctx, pcs.PinnedCommentQuery, pcs, pcs.inters, v)
}

func (pcs *PinnedCommentSelect) sqlScan(ctx context.Context, root *PinnedCommentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
