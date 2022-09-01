// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassDenomQuery is the builder for querying OutcomeClassDenom entities.
type OutcomeClassDenomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeClassDenom
	// eager-loading edges.
	withParent                     *OutcomeClassQuery
	withOutcomeClassDenomCountList *OutcomeClassDenomCountQuery
	withFKs                        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeClassDenomQuery builder.
func (ocdq *OutcomeClassDenomQuery) Where(ps ...predicate.OutcomeClassDenom) *OutcomeClassDenomQuery {
	ocdq.predicates = append(ocdq.predicates, ps...)
	return ocdq
}

// Limit adds a limit step to the query.
func (ocdq *OutcomeClassDenomQuery) Limit(limit int) *OutcomeClassDenomQuery {
	ocdq.limit = &limit
	return ocdq
}

// Offset adds an offset step to the query.
func (ocdq *OutcomeClassDenomQuery) Offset(offset int) *OutcomeClassDenomQuery {
	ocdq.offset = &offset
	return ocdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocdq *OutcomeClassDenomQuery) Unique(unique bool) *OutcomeClassDenomQuery {
	ocdq.unique = &unique
	return ocdq
}

// Order adds an order step to the query.
func (ocdq *OutcomeClassDenomQuery) Order(o ...OrderFunc) *OutcomeClassDenomQuery {
	ocdq.order = append(ocdq.order, o...)
	return ocdq
}

// QueryParent chains the current query on the "parent" edge.
func (ocdq *OutcomeClassDenomQuery) QueryParent() *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: ocdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenom.Table, outcomeclassdenom.FieldID, selector),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclassdenom.ParentTable, outcomeclassdenom.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeClassDenomCountList chains the current query on the "outcome_class_denom_count_list" edge.
func (ocdq *OutcomeClassDenomQuery) QueryOutcomeClassDenomCountList() *OutcomeClassDenomCountQuery {
	query := &OutcomeClassDenomCountQuery{config: ocdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenom.Table, outcomeclassdenom.FieldID, selector),
			sqlgraph.To(outcomeclassdenomcount.Table, outcomeclassdenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclassdenom.OutcomeClassDenomCountListTable, outcomeclassdenom.OutcomeClassDenomCountListColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeClassDenom entity from the query.
// Returns a *NotFoundError when no OutcomeClassDenom was found.
func (ocdq *OutcomeClassDenomQuery) First(ctx context.Context) (*OutcomeClassDenom, error) {
	nodes, err := ocdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeclassdenom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) FirstX(ctx context.Context) *OutcomeClassDenom {
	node, err := ocdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeClassDenom ID from the query.
// Returns a *NotFoundError when no OutcomeClassDenom ID was found.
func (ocdq *OutcomeClassDenomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeclassdenom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) FirstIDX(ctx context.Context) int {
	id, err := ocdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeClassDenom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeClassDenom entity is found.
// Returns a *NotFoundError when no OutcomeClassDenom entities are found.
func (ocdq *OutcomeClassDenomQuery) Only(ctx context.Context) (*OutcomeClassDenom, error) {
	nodes, err := ocdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeclassdenom.Label}
	default:
		return nil, &NotSingularError{outcomeclassdenom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) OnlyX(ctx context.Context) *OutcomeClassDenom {
	node, err := ocdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeClassDenom ID in the query.
// Returns a *NotSingularError when more than one OutcomeClassDenom ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocdq *OutcomeClassDenomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeclassdenom.Label}
	default:
		err = &NotSingularError{outcomeclassdenom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeClassDenoms.
func (ocdq *OutcomeClassDenomQuery) All(ctx context.Context) ([]*OutcomeClassDenom, error) {
	if err := ocdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ocdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) AllX(ctx context.Context) []*OutcomeClassDenom {
	nodes, err := ocdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeClassDenom IDs.
func (ocdq *OutcomeClassDenomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ocdq.Select(outcomeclassdenom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) IDsX(ctx context.Context) []int {
	ids, err := ocdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocdq *OutcomeClassDenomQuery) Count(ctx context.Context) (int, error) {
	if err := ocdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ocdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) CountX(ctx context.Context) int {
	count, err := ocdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocdq *OutcomeClassDenomQuery) Exist(ctx context.Context) (bool, error) {
	if err := ocdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ocdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ocdq *OutcomeClassDenomQuery) ExistX(ctx context.Context) bool {
	exist, err := ocdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeClassDenomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocdq *OutcomeClassDenomQuery) Clone() *OutcomeClassDenomQuery {
	if ocdq == nil {
		return nil
	}
	return &OutcomeClassDenomQuery{
		config:                         ocdq.config,
		limit:                          ocdq.limit,
		offset:                         ocdq.offset,
		order:                          append([]OrderFunc{}, ocdq.order...),
		predicates:                     append([]predicate.OutcomeClassDenom{}, ocdq.predicates...),
		withParent:                     ocdq.withParent.Clone(),
		withOutcomeClassDenomCountList: ocdq.withOutcomeClassDenomCountList.Clone(),
		// clone intermediate query.
		sql:    ocdq.sql.Clone(),
		path:   ocdq.path,
		unique: ocdq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ocdq *OutcomeClassDenomQuery) WithParent(opts ...func(*OutcomeClassQuery)) *OutcomeClassDenomQuery {
	query := &OutcomeClassQuery{config: ocdq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocdq.withParent = query
	return ocdq
}

// WithOutcomeClassDenomCountList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_class_denom_count_list" edge. The optional arguments are used to configure the query builder of the edge.
func (ocdq *OutcomeClassDenomQuery) WithOutcomeClassDenomCountList(opts ...func(*OutcomeClassDenomCountQuery)) *OutcomeClassDenomQuery {
	query := &OutcomeClassDenomCountQuery{config: ocdq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocdq.withOutcomeClassDenomCountList = query
	return ocdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeClassDenomUnits string `json:"outcome_class_denom_units,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeClassDenom.Query().
//		GroupBy(outcomeclassdenom.FieldOutcomeClassDenomUnits).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ocdq *OutcomeClassDenomQuery) GroupBy(field string, fields ...string) *OutcomeClassDenomGroupBy {
	grbuild := &OutcomeClassDenomGroupBy{config: ocdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ocdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ocdq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeclassdenom.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeClassDenomUnits string `json:"outcome_class_denom_units,omitempty"`
//	}
//
//	client.OutcomeClassDenom.Query().
//		Select(outcomeclassdenom.FieldOutcomeClassDenomUnits).
//		Scan(ctx, &v)
//
func (ocdq *OutcomeClassDenomQuery) Select(fields ...string) *OutcomeClassDenomSelect {
	ocdq.fields = append(ocdq.fields, fields...)
	selbuild := &OutcomeClassDenomSelect{OutcomeClassDenomQuery: ocdq}
	selbuild.label = outcomeclassdenom.Label
	selbuild.flds, selbuild.scan = &ocdq.fields, selbuild.Scan
	return selbuild
}

func (ocdq *OutcomeClassDenomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ocdq.fields {
		if !outcomeclassdenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ocdq.path != nil {
		prev, err := ocdq.path(ctx)
		if err != nil {
			return err
		}
		ocdq.sql = prev
	}
	return nil
}

func (ocdq *OutcomeClassDenomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeClassDenom, error) {
	var (
		nodes       = []*OutcomeClassDenom{}
		withFKs     = ocdq.withFKs
		_spec       = ocdq.querySpec()
		loadedTypes = [2]bool{
			ocdq.withParent != nil,
			ocdq.withOutcomeClassDenomCountList != nil,
		}
	)
	if ocdq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeClassDenom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeClassDenom{config: ocdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ocdq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeClassDenom)
		for i := range nodes {
			if nodes[i].outcome_class_outcome_class_denom_list == nil {
				continue
			}
			fk := *nodes[i].outcome_class_outcome_class_denom_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomeclass.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_outcome_class_denom_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ocdq.withOutcomeClassDenomCountList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeClassDenom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeClassDenomCountList = []*OutcomeClassDenomCount{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeClassDenomCount(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomeclassdenom.OutcomeClassDenomCountListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_class_denom_outcome_class_denom_count_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_class_denom_outcome_class_denom_count_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_denom_outcome_class_denom_count_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeClassDenomCountList = append(node.Edges.OutcomeClassDenomCountList, n)
		}
	}

	return nodes, nil
}

func (ocdq *OutcomeClassDenomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocdq.querySpec()
	_spec.Node.Columns = ocdq.fields
	if len(ocdq.fields) > 0 {
		_spec.Unique = ocdq.unique != nil && *ocdq.unique
	}
	return sqlgraph.CountNodes(ctx, ocdq.driver, _spec)
}

func (ocdq *OutcomeClassDenomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ocdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ocdq *OutcomeClassDenomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclassdenom.Table,
			Columns: outcomeclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclassdenom.FieldID,
			},
		},
		From:   ocdq.sql,
		Unique: true,
	}
	if unique := ocdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ocdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclassdenom.FieldID)
		for i := range fields {
			if fields[i] != outcomeclassdenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocdq *OutcomeClassDenomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocdq.driver.Dialect())
	t1 := builder.Table(outcomeclassdenom.Table)
	columns := ocdq.fields
	if len(columns) == 0 {
		columns = outcomeclassdenom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocdq.sql != nil {
		selector = ocdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocdq.unique != nil && *ocdq.unique {
		selector.Distinct()
	}
	for _, p := range ocdq.predicates {
		p(selector)
	}
	for _, p := range ocdq.order {
		p(selector)
	}
	if offset := ocdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeClassDenomGroupBy is the group-by builder for OutcomeClassDenom entities.
type OutcomeClassDenomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocdgb *OutcomeClassDenomGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeClassDenomGroupBy {
	ocdgb.fns = append(ocdgb.fns, fns...)
	return ocdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ocdgb *OutcomeClassDenomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ocdgb.path(ctx)
	if err != nil {
		return err
	}
	ocdgb.sql = query
	return ocdgb.sqlScan(ctx, v)
}

func (ocdgb *OutcomeClassDenomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ocdgb.fields {
		if !outcomeclassdenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ocdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ocdgb *OutcomeClassDenomGroupBy) sqlQuery() *sql.Selector {
	selector := ocdgb.sql.Select()
	aggregation := make([]string, 0, len(ocdgb.fns))
	for _, fn := range ocdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ocdgb.fields)+len(ocdgb.fns))
		for _, f := range ocdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ocdgb.fields...)...)
}

// OutcomeClassDenomSelect is the builder for selecting fields of OutcomeClassDenom entities.
type OutcomeClassDenomSelect struct {
	*OutcomeClassDenomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ocds *OutcomeClassDenomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ocds.prepareQuery(ctx); err != nil {
		return err
	}
	ocds.sql = ocds.OutcomeClassDenomQuery.sqlQuery(ctx)
	return ocds.sqlScan(ctx, v)
}

func (ocds *OutcomeClassDenomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ocds.sql.Query()
	if err := ocds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
