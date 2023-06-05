// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sadzeih/valcompbot/ent/pickemsevent"
	"github.com/Sadzeih/valcompbot/ent/predicate"
	"github.com/google/uuid"
)

// PickemsEventQuery is the builder for querying PickemsEvent entities.
type PickemsEventQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PickemsEvent
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PickemsEventQuery builder.
func (peq *PickemsEventQuery) Where(ps ...predicate.PickemsEvent) *PickemsEventQuery {
	peq.predicates = append(peq.predicates, ps...)
	return peq
}

// Limit adds a limit step to the query.
func (peq *PickemsEventQuery) Limit(limit int) *PickemsEventQuery {
	peq.limit = &limit
	return peq
}

// Offset adds an offset step to the query.
func (peq *PickemsEventQuery) Offset(offset int) *PickemsEventQuery {
	peq.offset = &offset
	return peq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (peq *PickemsEventQuery) Unique(unique bool) *PickemsEventQuery {
	peq.unique = &unique
	return peq
}

// Order adds an order step to the query.
func (peq *PickemsEventQuery) Order(o ...OrderFunc) *PickemsEventQuery {
	peq.order = append(peq.order, o...)
	return peq
}

// First returns the first PickemsEvent entity from the query.
// Returns a *NotFoundError when no PickemsEvent was found.
func (peq *PickemsEventQuery) First(ctx context.Context) (*PickemsEvent, error) {
	nodes, err := peq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pickemsevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (peq *PickemsEventQuery) FirstX(ctx context.Context) *PickemsEvent {
	node, err := peq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PickemsEvent ID from the query.
// Returns a *NotFoundError when no PickemsEvent ID was found.
func (peq *PickemsEventQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = peq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pickemsevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (peq *PickemsEventQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := peq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PickemsEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PickemsEvent entity is found.
// Returns a *NotFoundError when no PickemsEvent entities are found.
func (peq *PickemsEventQuery) Only(ctx context.Context) (*PickemsEvent, error) {
	nodes, err := peq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pickemsevent.Label}
	default:
		return nil, &NotSingularError{pickemsevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (peq *PickemsEventQuery) OnlyX(ctx context.Context) *PickemsEvent {
	node, err := peq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PickemsEvent ID in the query.
// Returns a *NotSingularError when more than one PickemsEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (peq *PickemsEventQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = peq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pickemsevent.Label}
	default:
		err = &NotSingularError{pickemsevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (peq *PickemsEventQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := peq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PickemsEvents.
func (peq *PickemsEventQuery) All(ctx context.Context) ([]*PickemsEvent, error) {
	if err := peq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return peq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (peq *PickemsEventQuery) AllX(ctx context.Context) []*PickemsEvent {
	nodes, err := peq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PickemsEvent IDs.
func (peq *PickemsEventQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := peq.Select(pickemsevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (peq *PickemsEventQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := peq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (peq *PickemsEventQuery) Count(ctx context.Context) (int, error) {
	if err := peq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return peq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (peq *PickemsEventQuery) CountX(ctx context.Context) int {
	count, err := peq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (peq *PickemsEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := peq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return peq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (peq *PickemsEventQuery) ExistX(ctx context.Context) bool {
	exist, err := peq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PickemsEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (peq *PickemsEventQuery) Clone() *PickemsEventQuery {
	if peq == nil {
		return nil
	}
	return &PickemsEventQuery{
		config:     peq.config,
		limit:      peq.limit,
		offset:     peq.offset,
		order:      append([]OrderFunc{}, peq.order...),
		predicates: append([]predicate.PickemsEvent{}, peq.predicates...),
		// clone intermediate query.
		sql:    peq.sql.Clone(),
		path:   peq.path,
		unique: peq.unique,
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
//	client.PickemsEvent.Query().
//		GroupBy(pickemsevent.FieldEventID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (peq *PickemsEventQuery) GroupBy(field string, fields ...string) *PickemsEventGroupBy {
	grbuild := &PickemsEventGroupBy{config: peq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := peq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return peq.sqlQuery(ctx), nil
	}
	grbuild.label = pickemsevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.PickemsEvent.Query().
//		Select(pickemsevent.FieldEventID).
//		Scan(ctx, &v)
func (peq *PickemsEventQuery) Select(fields ...string) *PickemsEventSelect {
	peq.fields = append(peq.fields, fields...)
	selbuild := &PickemsEventSelect{PickemsEventQuery: peq}
	selbuild.label = pickemsevent.Label
	selbuild.flds, selbuild.scan = &peq.fields, selbuild.Scan
	return selbuild
}

func (peq *PickemsEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range peq.fields {
		if !pickemsevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if peq.path != nil {
		prev, err := peq.path(ctx)
		if err != nil {
			return err
		}
		peq.sql = prev
	}
	return nil
}

func (peq *PickemsEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PickemsEvent, error) {
	var (
		nodes = []*PickemsEvent{}
		_spec = peq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PickemsEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PickemsEvent{config: peq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, peq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (peq *PickemsEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := peq.querySpec()
	_spec.Node.Columns = peq.fields
	if len(peq.fields) > 0 {
		_spec.Unique = peq.unique != nil && *peq.unique
	}
	return sqlgraph.CountNodes(ctx, peq.driver, _spec)
}

func (peq *PickemsEventQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := peq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (peq *PickemsEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pickemsevent.Table,
			Columns: pickemsevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pickemsevent.FieldID,
			},
		},
		From:   peq.sql,
		Unique: true,
	}
	if unique := peq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := peq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pickemsevent.FieldID)
		for i := range fields {
			if fields[i] != pickemsevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := peq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := peq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := peq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := peq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (peq *PickemsEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(peq.driver.Dialect())
	t1 := builder.Table(pickemsevent.Table)
	columns := peq.fields
	if len(columns) == 0 {
		columns = pickemsevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if peq.sql != nil {
		selector = peq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if peq.unique != nil && *peq.unique {
		selector.Distinct()
	}
	for _, p := range peq.predicates {
		p(selector)
	}
	for _, p := range peq.order {
		p(selector)
	}
	if offset := peq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := peq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PickemsEventGroupBy is the group-by builder for PickemsEvent entities.
type PickemsEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pegb *PickemsEventGroupBy) Aggregate(fns ...AggregateFunc) *PickemsEventGroupBy {
	pegb.fns = append(pegb.fns, fns...)
	return pegb
}

// Scan applies the group-by query and scans the result into the given value.
func (pegb *PickemsEventGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pegb.path(ctx)
	if err != nil {
		return err
	}
	pegb.sql = query
	return pegb.sqlScan(ctx, v)
}

func (pegb *PickemsEventGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pegb.fields {
		if !pickemsevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pegb *PickemsEventGroupBy) sqlQuery() *sql.Selector {
	selector := pegb.sql.Select()
	aggregation := make([]string, 0, len(pegb.fns))
	for _, fn := range pegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pegb.fields)+len(pegb.fns))
		for _, f := range pegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pegb.fields...)...)
}

// PickemsEventSelect is the builder for selecting fields of PickemsEvent entities.
type PickemsEventSelect struct {
	*PickemsEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pes *PickemsEventSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pes.prepareQuery(ctx); err != nil {
		return err
	}
	pes.sql = pes.PickemsEventQuery.sqlQuery(ctx)
	return pes.sqlScan(ctx, v)
}

func (pes *PickemsEventSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pes.sql.Query()
	if err := pes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}