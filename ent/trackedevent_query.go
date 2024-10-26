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
	"github.com/Sadzeih/valcompbot/ent/predicate"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/google/uuid"
)

// TrackedEventQuery is the builder for querying TrackedEvent entities.
type TrackedEventQuery struct {
	config
	ctx        *QueryContext
	order      []trackedevent.OrderOption
	inters     []Interceptor
	predicates []predicate.TrackedEvent
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TrackedEventQuery builder.
func (teq *TrackedEventQuery) Where(ps ...predicate.TrackedEvent) *TrackedEventQuery {
	teq.predicates = append(teq.predicates, ps...)
	return teq
}

// Limit the number of records to be returned by this query.
func (teq *TrackedEventQuery) Limit(limit int) *TrackedEventQuery {
	teq.ctx.Limit = &limit
	return teq
}

// Offset to start from.
func (teq *TrackedEventQuery) Offset(offset int) *TrackedEventQuery {
	teq.ctx.Offset = &offset
	return teq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (teq *TrackedEventQuery) Unique(unique bool) *TrackedEventQuery {
	teq.ctx.Unique = &unique
	return teq
}

// Order specifies how the records should be ordered.
func (teq *TrackedEventQuery) Order(o ...trackedevent.OrderOption) *TrackedEventQuery {
	teq.order = append(teq.order, o...)
	return teq
}

// First returns the first TrackedEvent entity from the query.
// Returns a *NotFoundError when no TrackedEvent was found.
func (teq *TrackedEventQuery) First(ctx context.Context) (*TrackedEvent, error) {
	nodes, err := teq.Limit(1).All(setContextOp(ctx, teq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{trackedevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (teq *TrackedEventQuery) FirstX(ctx context.Context) *TrackedEvent {
	node, err := teq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TrackedEvent ID from the query.
// Returns a *NotFoundError when no TrackedEvent ID was found.
func (teq *TrackedEventQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = teq.Limit(1).IDs(setContextOp(ctx, teq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{trackedevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (teq *TrackedEventQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := teq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TrackedEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TrackedEvent entity is found.
// Returns a *NotFoundError when no TrackedEvent entities are found.
func (teq *TrackedEventQuery) Only(ctx context.Context) (*TrackedEvent, error) {
	nodes, err := teq.Limit(2).All(setContextOp(ctx, teq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{trackedevent.Label}
	default:
		return nil, &NotSingularError{trackedevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (teq *TrackedEventQuery) OnlyX(ctx context.Context) *TrackedEvent {
	node, err := teq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TrackedEvent ID in the query.
// Returns a *NotSingularError when more than one TrackedEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (teq *TrackedEventQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = teq.Limit(2).IDs(setContextOp(ctx, teq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{trackedevent.Label}
	default:
		err = &NotSingularError{trackedevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (teq *TrackedEventQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := teq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TrackedEvents.
func (teq *TrackedEventQuery) All(ctx context.Context) ([]*TrackedEvent, error) {
	ctx = setContextOp(ctx, teq.ctx, ent.OpQueryAll)
	if err := teq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TrackedEvent, *TrackedEventQuery]()
	return withInterceptors[[]*TrackedEvent](ctx, teq, qr, teq.inters)
}

// AllX is like All, but panics if an error occurs.
func (teq *TrackedEventQuery) AllX(ctx context.Context) []*TrackedEvent {
	nodes, err := teq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TrackedEvent IDs.
func (teq *TrackedEventQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if teq.ctx.Unique == nil && teq.path != nil {
		teq.Unique(true)
	}
	ctx = setContextOp(ctx, teq.ctx, ent.OpQueryIDs)
	if err = teq.Select(trackedevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (teq *TrackedEventQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := teq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (teq *TrackedEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, teq.ctx, ent.OpQueryCount)
	if err := teq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, teq, querierCount[*TrackedEventQuery](), teq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (teq *TrackedEventQuery) CountX(ctx context.Context) int {
	count, err := teq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (teq *TrackedEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, teq.ctx, ent.OpQueryExist)
	switch _, err := teq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (teq *TrackedEventQuery) ExistX(ctx context.Context) bool {
	exist, err := teq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TrackedEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (teq *TrackedEventQuery) Clone() *TrackedEventQuery {
	if teq == nil {
		return nil
	}
	return &TrackedEventQuery{
		config:     teq.config,
		ctx:        teq.ctx.Clone(),
		order:      append([]trackedevent.OrderOption{}, teq.order...),
		inters:     append([]Interceptor{}, teq.inters...),
		predicates: append([]predicate.TrackedEvent{}, teq.predicates...),
		// clone intermediate query.
		sql:  teq.sql.Clone(),
		path: teq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TrackedEvent.Query().
//		GroupBy(trackedevent.FieldEventID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (teq *TrackedEventQuery) GroupBy(field string, fields ...string) *TrackedEventGroupBy {
	teq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TrackedEventGroupBy{build: teq}
	grbuild.flds = &teq.ctx.Fields
	grbuild.label = trackedevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id,omitempty"`
//	}
//
//	client.TrackedEvent.Query().
//		Select(trackedevent.FieldEventID).
//		Scan(ctx, &v)
func (teq *TrackedEventQuery) Select(fields ...string) *TrackedEventSelect {
	teq.ctx.Fields = append(teq.ctx.Fields, fields...)
	sbuild := &TrackedEventSelect{TrackedEventQuery: teq}
	sbuild.label = trackedevent.Label
	sbuild.flds, sbuild.scan = &teq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TrackedEventSelect configured with the given aggregations.
func (teq *TrackedEventQuery) Aggregate(fns ...AggregateFunc) *TrackedEventSelect {
	return teq.Select().Aggregate(fns...)
}

func (teq *TrackedEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range teq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, teq); err != nil {
				return err
			}
		}
	}
	for _, f := range teq.ctx.Fields {
		if !trackedevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if teq.path != nil {
		prev, err := teq.path(ctx)
		if err != nil {
			return err
		}
		teq.sql = prev
	}
	return nil
}

func (teq *TrackedEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TrackedEvent, error) {
	var (
		nodes = []*TrackedEvent{}
		_spec = teq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TrackedEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TrackedEvent{config: teq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, teq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (teq *TrackedEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := teq.querySpec()
	_spec.Node.Columns = teq.ctx.Fields
	if len(teq.ctx.Fields) > 0 {
		_spec.Unique = teq.ctx.Unique != nil && *teq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, teq.driver, _spec)
}

func (teq *TrackedEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(trackedevent.Table, trackedevent.Columns, sqlgraph.NewFieldSpec(trackedevent.FieldID, field.TypeUUID))
	_spec.From = teq.sql
	if unique := teq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if teq.path != nil {
		_spec.Unique = true
	}
	if fields := teq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trackedevent.FieldID)
		for i := range fields {
			if fields[i] != trackedevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := teq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := teq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := teq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := teq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (teq *TrackedEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(teq.driver.Dialect())
	t1 := builder.Table(trackedevent.Table)
	columns := teq.ctx.Fields
	if len(columns) == 0 {
		columns = trackedevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if teq.sql != nil {
		selector = teq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if teq.ctx.Unique != nil && *teq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range teq.predicates {
		p(selector)
	}
	for _, p := range teq.order {
		p(selector)
	}
	if offset := teq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := teq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TrackedEventGroupBy is the group-by builder for TrackedEvent entities.
type TrackedEventGroupBy struct {
	selector
	build *TrackedEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tegb *TrackedEventGroupBy) Aggregate(fns ...AggregateFunc) *TrackedEventGroupBy {
	tegb.fns = append(tegb.fns, fns...)
	return tegb
}

// Scan applies the selector query and scans the result into the given value.
func (tegb *TrackedEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tegb.build.ctx, ent.OpQueryGroupBy)
	if err := tegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TrackedEventQuery, *TrackedEventGroupBy](ctx, tegb.build, tegb, tegb.build.inters, v)
}

func (tegb *TrackedEventGroupBy) sqlScan(ctx context.Context, root *TrackedEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tegb.fns))
	for _, fn := range tegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tegb.flds)+len(tegb.fns))
		for _, f := range *tegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TrackedEventSelect is the builder for selecting fields of TrackedEvent entities.
type TrackedEventSelect struct {
	*TrackedEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tes *TrackedEventSelect) Aggregate(fns ...AggregateFunc) *TrackedEventSelect {
	tes.fns = append(tes.fns, fns...)
	return tes
}

// Scan applies the selector query and scans the result into the given value.
func (tes *TrackedEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tes.ctx, ent.OpQuerySelect)
	if err := tes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TrackedEventQuery, *TrackedEventSelect](ctx, tes.TrackedEventQuery, tes, tes.inters, v)
}

func (tes *TrackedEventSelect) sqlScan(ctx context.Context, root *TrackedEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tes.fns))
	for _, fn := range tes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
