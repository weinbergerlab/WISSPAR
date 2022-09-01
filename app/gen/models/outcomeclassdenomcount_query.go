// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomCountQuery is the builder for querying OutcomeClassDenomCount entities.
type OutcomeClassDenomCountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeClassDenomCount
	// eager-loading edges.
	withParent *OutcomeClassDenomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeClassDenomCountQuery builder.
func (ocdcq *OutcomeClassDenomCountQuery) Where(ps ...predicate.OutcomeClassDenomCount) *OutcomeClassDenomCountQuery {
	ocdcq.predicates = append(ocdcq.predicates, ps...)
	return ocdcq
}

// Limit adds a limit step to the query.
func (ocdcq *OutcomeClassDenomCountQuery) Limit(limit int) *OutcomeClassDenomCountQuery {
	ocdcq.limit = &limit
	return ocdcq
}

// Offset adds an offset step to the query.
func (ocdcq *OutcomeClassDenomCountQuery) Offset(offset int) *OutcomeClassDenomCountQuery {
	ocdcq.offset = &offset
	return ocdcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocdcq *OutcomeClassDenomCountQuery) Unique(unique bool) *OutcomeClassDenomCountQuery {
	ocdcq.unique = &unique
	return ocdcq
}

// Order adds an order step to the query.
func (ocdcq *OutcomeClassDenomCountQuery) Order(o ...OrderFunc) *OutcomeClassDenomCountQuery {
	ocdcq.order = append(ocdcq.order, o...)
	return ocdcq
}

// QueryParent chains the current query on the "parent" edge.
func (ocdcq *OutcomeClassDenomCountQuery) QueryParent() *OutcomeClassDenomQuery {
	query := &OutcomeClassDenomQuery{config: ocdcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocdcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenomcount.Table, outcomeclassdenomcount.FieldID, selector),
			sqlgraph.To(outcomeclassdenom.Table, outcomeclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclassdenomcount.ParentTable, outcomeclassdenomcount.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocdcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeClassDenomCount entity from the query.
// Returns a *NotFoundError when no OutcomeClassDenomCount was found.
func (ocdcq *OutcomeClassDenomCountQuery) First(ctx context.Context) (*OutcomeClassDenomCount, error) {
	nodes, err := ocdcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeclassdenomcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) FirstX(ctx context.Context) *OutcomeClassDenomCount {
	node, err := ocdcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeClassDenomCount ID from the query.
// Returns a *NotFoundError when no OutcomeClassDenomCount ID was found.
func (ocdcq *OutcomeClassDenomCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocdcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeclassdenomcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) FirstIDX(ctx context.Context) int {
	id, err := ocdcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeClassDenomCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeClassDenomCount entity is found.
// Returns a *NotFoundError when no OutcomeClassDenomCount entities are found.
func (ocdcq *OutcomeClassDenomCountQuery) Only(ctx context.Context) (*OutcomeClassDenomCount, error) {
	nodes, err := ocdcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeclassdenomcount.Label}
	default:
		return nil, &NotSingularError{outcomeclassdenomcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) OnlyX(ctx context.Context) *OutcomeClassDenomCount {
	node, err := ocdcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeClassDenomCount ID in the query.
// Returns a *NotSingularError when more than one OutcomeClassDenomCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocdcq *OutcomeClassDenomCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocdcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeclassdenomcount.Label}
	default:
		err = &NotSingularError{outcomeclassdenomcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocdcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeClassDenomCounts.
func (ocdcq *OutcomeClassDenomCountQuery) All(ctx context.Context) ([]*OutcomeClassDenomCount, error) {
	if err := ocdcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ocdcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) AllX(ctx context.Context) []*OutcomeClassDenomCount {
	nodes, err := ocdcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeClassDenomCount IDs.
func (ocdcq *OutcomeClassDenomCountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ocdcq.Select(outcomeclassdenomcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) IDsX(ctx context.Context) []int {
	ids, err := ocdcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocdcq *OutcomeClassDenomCountQuery) Count(ctx context.Context) (int, error) {
	if err := ocdcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ocdcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) CountX(ctx context.Context) int {
	count, err := ocdcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocdcq *OutcomeClassDenomCountQuery) Exist(ctx context.Context) (bool, error) {
	if err := ocdcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ocdcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ocdcq *OutcomeClassDenomCountQuery) ExistX(ctx context.Context) bool {
	exist, err := ocdcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeClassDenomCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocdcq *OutcomeClassDenomCountQuery) Clone() *OutcomeClassDenomCountQuery {
	if ocdcq == nil {
		return nil
	}
	return &OutcomeClassDenomCountQuery{
		config:     ocdcq.config,
		limit:      ocdcq.limit,
		offset:     ocdcq.offset,
		order:      append([]OrderFunc{}, ocdcq.order...),
		predicates: append([]predicate.OutcomeClassDenomCount{}, ocdcq.predicates...),
		withParent: ocdcq.withParent.Clone(),
		// clone intermediate query.
		sql:    ocdcq.sql.Clone(),
		path:   ocdcq.path,
		unique: ocdcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ocdcq *OutcomeClassDenomCountQuery) WithParent(opts ...func(*OutcomeClassDenomQuery)) *OutcomeClassDenomCountQuery {
	query := &OutcomeClassDenomQuery{config: ocdcq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocdcq.withParent = query
	return ocdcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeClassDenomCountGroupID string `json:"outcome_class_denom_count_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeClassDenomCount.Query().
//		GroupBy(outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ocdcq *OutcomeClassDenomCountQuery) GroupBy(field string, fields ...string) *OutcomeClassDenomCountGroupBy {
	grbuild := &OutcomeClassDenomCountGroupBy{config: ocdcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ocdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ocdcq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeclassdenomcount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeClassDenomCountGroupID string `json:"outcome_class_denom_count_group_id,omitempty"`
//	}
//
//	client.OutcomeClassDenomCount.Query().
//		Select(outcomeclassdenomcount.FieldOutcomeClassDenomCountGroupID).
//		Scan(ctx, &v)
//
func (ocdcq *OutcomeClassDenomCountQuery) Select(fields ...string) *OutcomeClassDenomCountSelect {
	ocdcq.fields = append(ocdcq.fields, fields...)
	selbuild := &OutcomeClassDenomCountSelect{OutcomeClassDenomCountQuery: ocdcq}
	selbuild.label = outcomeclassdenomcount.Label
	selbuild.flds, selbuild.scan = &ocdcq.fields, selbuild.Scan
	return selbuild
}

func (ocdcq *OutcomeClassDenomCountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ocdcq.fields {
		if !outcomeclassdenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ocdcq.path != nil {
		prev, err := ocdcq.path(ctx)
		if err != nil {
			return err
		}
		ocdcq.sql = prev
	}
	return nil
}

func (ocdcq *OutcomeClassDenomCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeClassDenomCount, error) {
	var (
		nodes       = []*OutcomeClassDenomCount{}
		withFKs     = ocdcq.withFKs
		_spec       = ocdcq.querySpec()
		loadedTypes = [1]bool{
			ocdcq.withParent != nil,
		}
	)
	if ocdcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenomcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeClassDenomCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeClassDenomCount{config: ocdcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocdcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ocdcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeClassDenomCount)
		for i := range nodes {
			if nodes[i].outcome_class_denom_outcome_class_denom_count_list == nil {
				continue
			}
			fk := *nodes[i].outcome_class_denom_outcome_class_denom_count_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomeclassdenom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_denom_outcome_class_denom_count_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (ocdcq *OutcomeClassDenomCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocdcq.querySpec()
	_spec.Node.Columns = ocdcq.fields
	if len(ocdcq.fields) > 0 {
		_spec.Unique = ocdcq.unique != nil && *ocdcq.unique
	}
	return sqlgraph.CountNodes(ctx, ocdcq.driver, _spec)
}

func (ocdcq *OutcomeClassDenomCountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ocdcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ocdcq *OutcomeClassDenomCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenomcount.Table,
			Columns: outcomeclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenomcount.FieldID,
			},
		},
		From:   ocdcq.sql,
		Unique: true,
	}
	if unique := ocdcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ocdcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenomcount.FieldID)
		for i := range fields {
			if fields[i] != outcomeclassdenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocdcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocdcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocdcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocdcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocdcq *OutcomeClassDenomCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocdcq.driver.Dialect())
	t1 := builder.Table(outcomeclassdenomcount.Table)
	columns := ocdcq.fields
	if len(columns) == 0 {
		columns = outcomeclassdenomcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocdcq.sql != nil {
		selector = ocdcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocdcq.unique != nil && *ocdcq.unique {
		selector.Distinct()
	}
	for _, p := range ocdcq.predicates {
		p(selector)
	}
	for _, p := range ocdcq.order {
		p(selector)
	}
	if offset := ocdcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocdcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeClassDenomCountGroupBy is the group-by builder for OutcomeClassDenomCount entities.
type OutcomeClassDenomCountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocdcgb *OutcomeClassDenomCountGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeClassDenomCountGroupBy {
	ocdcgb.fns = append(ocdcgb.fns, fns...)
	return ocdcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ocdcgb *OutcomeClassDenomCountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ocdcgb.path(ctx)
	if err != nil {
		return err
	}
	ocdcgb.sql = query
	return ocdcgb.sqlScan(ctx, v)
}

func (ocdcgb *OutcomeClassDenomCountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ocdcgb.fields {
		if !outcomeclassdenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ocdcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocdcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ocdcgb *OutcomeClassDenomCountGroupBy) sqlQuery() *sql.Selector {
	selector := ocdcgb.sql.Select()
	aggregation := make([]string, 0, len(ocdcgb.fns))
	for _, fn := range ocdcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ocdcgb.fields)+len(ocdcgb.fns))
		for _, f := range ocdcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ocdcgb.fields...)...)
}

// OutcomeClassDenomCountSelect is the builder for selecting fields of OutcomeClassDenomCount entities.
type OutcomeClassDenomCountSelect struct {
	*OutcomeClassDenomCountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ocdcs *OutcomeClassDenomCountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ocdcs.prepareQuery(ctx); err != nil {
		return err
	}
	ocdcs.sql = ocdcs.OutcomeClassDenomCountQuery.sqlQuery(ctx)
	return ocdcs.sqlScan(ctx, v)
}

func (ocdcs *OutcomeClassDenomCountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ocdcs.sql.Query()
	if err := ocdcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
