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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeCategoryQuery is the builder for querying OutcomeCategory entities.
type OutcomeCategoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeCategory
	// eager-loading edges.
	withParent                 *OutcomeClassQuery
	withOutcomeMeasurementList *OutcomeMeasurementQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeCategoryQuery builder.
func (ocq *OutcomeCategoryQuery) Where(ps ...predicate.OutcomeCategory) *OutcomeCategoryQuery {
	ocq.predicates = append(ocq.predicates, ps...)
	return ocq
}

// Limit adds a limit step to the query.
func (ocq *OutcomeCategoryQuery) Limit(limit int) *OutcomeCategoryQuery {
	ocq.limit = &limit
	return ocq
}

// Offset adds an offset step to the query.
func (ocq *OutcomeCategoryQuery) Offset(offset int) *OutcomeCategoryQuery {
	ocq.offset = &offset
	return ocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocq *OutcomeCategoryQuery) Unique(unique bool) *OutcomeCategoryQuery {
	ocq.unique = &unique
	return ocq
}

// Order adds an order step to the query.
func (ocq *OutcomeCategoryQuery) Order(o ...OrderFunc) *OutcomeCategoryQuery {
	ocq.order = append(ocq.order, o...)
	return ocq
}

// QueryParent chains the current query on the "parent" edge.
func (ocq *OutcomeCategoryQuery) QueryParent() *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: ocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomecategory.Table, outcomecategory.FieldID, selector),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomecategory.ParentTable, outcomecategory.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeMeasurementList chains the current query on the "outcome_measurement_list" edge.
func (ocq *OutcomeCategoryQuery) QueryOutcomeMeasurementList() *OutcomeMeasurementQuery {
	query := &OutcomeMeasurementQuery{config: ocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomecategory.Table, outcomecategory.FieldID, selector),
			sqlgraph.To(outcomemeasurement.Table, outcomemeasurement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomecategory.OutcomeMeasurementListTable, outcomecategory.OutcomeMeasurementListColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeCategory entity from the query.
// Returns a *NotFoundError when no OutcomeCategory was found.
func (ocq *OutcomeCategoryQuery) First(ctx context.Context) (*OutcomeCategory, error) {
	nodes, err := ocq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) FirstX(ctx context.Context) *OutcomeCategory {
	node, err := ocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeCategory ID from the query.
// Returns a *NotFoundError when no OutcomeCategory ID was found.
func (ocq *OutcomeCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := ocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeCategory entity is found.
// Returns a *NotFoundError when no OutcomeCategory entities are found.
func (ocq *OutcomeCategoryQuery) Only(ctx context.Context) (*OutcomeCategory, error) {
	nodes, err := ocq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomecategory.Label}
	default:
		return nil, &NotSingularError{outcomecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) OnlyX(ctx context.Context) *OutcomeCategory {
	node, err := ocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeCategory ID in the query.
// Returns a *NotSingularError when more than one OutcomeCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocq *OutcomeCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomecategory.Label}
	default:
		err = &NotSingularError{outcomecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeCategories.
func (ocq *OutcomeCategoryQuery) All(ctx context.Context) ([]*OutcomeCategory, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ocq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) AllX(ctx context.Context) []*OutcomeCategory {
	nodes, err := ocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeCategory IDs.
func (ocq *OutcomeCategoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ocq.Select(outcomecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := ocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocq *OutcomeCategoryQuery) Count(ctx context.Context) (int, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ocq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) CountX(ctx context.Context) int {
	count, err := ocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocq *OutcomeCategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ocq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ocq *OutcomeCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocq *OutcomeCategoryQuery) Clone() *OutcomeCategoryQuery {
	if ocq == nil {
		return nil
	}
	return &OutcomeCategoryQuery{
		config:                     ocq.config,
		limit:                      ocq.limit,
		offset:                     ocq.offset,
		order:                      append([]OrderFunc{}, ocq.order...),
		predicates:                 append([]predicate.OutcomeCategory{}, ocq.predicates...),
		withParent:                 ocq.withParent.Clone(),
		withOutcomeMeasurementList: ocq.withOutcomeMeasurementList.Clone(),
		// clone intermediate query.
		sql:    ocq.sql.Clone(),
		path:   ocq.path,
		unique: ocq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OutcomeCategoryQuery) WithParent(opts ...func(*OutcomeClassQuery)) *OutcomeCategoryQuery {
	query := &OutcomeClassQuery{config: ocq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocq.withParent = query
	return ocq
}

// WithOutcomeMeasurementList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_measurement_list" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OutcomeCategoryQuery) WithOutcomeMeasurementList(opts ...func(*OutcomeMeasurementQuery)) *OutcomeCategoryQuery {
	query := &OutcomeMeasurementQuery{config: ocq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocq.withOutcomeMeasurementList = query
	return ocq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeCategoryTitle string `json:"outcome_category_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeCategory.Query().
//		GroupBy(outcomecategory.FieldOutcomeCategoryTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ocq *OutcomeCategoryQuery) GroupBy(field string, fields ...string) *OutcomeCategoryGroupBy {
	grbuild := &OutcomeCategoryGroupBy{config: ocq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ocq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomecategory.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeCategoryTitle string `json:"outcome_category_title,omitempty"`
//	}
//
//	client.OutcomeCategory.Query().
//		Select(outcomecategory.FieldOutcomeCategoryTitle).
//		Scan(ctx, &v)
//
func (ocq *OutcomeCategoryQuery) Select(fields ...string) *OutcomeCategorySelect {
	ocq.fields = append(ocq.fields, fields...)
	selbuild := &OutcomeCategorySelect{OutcomeCategoryQuery: ocq}
	selbuild.label = outcomecategory.Label
	selbuild.flds, selbuild.scan = &ocq.fields, selbuild.Scan
	return selbuild
}

func (ocq *OutcomeCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ocq.fields {
		if !outcomecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ocq.path != nil {
		prev, err := ocq.path(ctx)
		if err != nil {
			return err
		}
		ocq.sql = prev
	}
	return nil
}

func (ocq *OutcomeCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeCategory, error) {
	var (
		nodes       = []*OutcomeCategory{}
		withFKs     = ocq.withFKs
		_spec       = ocq.querySpec()
		loadedTypes = [2]bool{
			ocq.withParent != nil,
			ocq.withOutcomeMeasurementList != nil,
		}
	)
	if ocq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomecategory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeCategory{config: ocq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ocq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeCategory)
		for i := range nodes {
			if nodes[i].outcome_class_outcome_category_list == nil {
				continue
			}
			fk := *nodes[i].outcome_class_outcome_category_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_outcome_category_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ocq.withOutcomeMeasurementList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeCategory)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeMeasurementList = []*OutcomeMeasurement{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeMeasurement(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomecategory.OutcomeMeasurementListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_category_outcome_measurement_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_category_outcome_measurement_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_category_outcome_measurement_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeMeasurementList = append(node.Edges.OutcomeMeasurementList, n)
		}
	}

	return nodes, nil
}

func (ocq *OutcomeCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocq.querySpec()
	_spec.Node.Columns = ocq.fields
	if len(ocq.fields) > 0 {
		_spec.Unique = ocq.unique != nil && *ocq.unique
	}
	return sqlgraph.CountNodes(ctx, ocq.driver, _spec)
}

func (ocq *OutcomeCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ocq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ocq *OutcomeCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomecategory.Table,
			Columns: outcomecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomecategory.FieldID,
			},
		},
		From:   ocq.sql,
		Unique: true,
	}
	if unique := ocq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ocq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomecategory.FieldID)
		for i := range fields {
			if fields[i] != outcomecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocq *OutcomeCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocq.driver.Dialect())
	t1 := builder.Table(outcomecategory.Table)
	columns := ocq.fields
	if len(columns) == 0 {
		columns = outcomecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocq.sql != nil {
		selector = ocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocq.unique != nil && *ocq.unique {
		selector.Distinct()
	}
	for _, p := range ocq.predicates {
		p(selector)
	}
	for _, p := range ocq.order {
		p(selector)
	}
	if offset := ocq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeCategoryGroupBy is the group-by builder for OutcomeCategory entities.
type OutcomeCategoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocgb *OutcomeCategoryGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeCategoryGroupBy {
	ocgb.fns = append(ocgb.fns, fns...)
	return ocgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ocgb *OutcomeCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ocgb.path(ctx)
	if err != nil {
		return err
	}
	ocgb.sql = query
	return ocgb.sqlScan(ctx, v)
}

func (ocgb *OutcomeCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ocgb.fields {
		if !outcomecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ocgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ocgb *OutcomeCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := ocgb.sql.Select()
	aggregation := make([]string, 0, len(ocgb.fns))
	for _, fn := range ocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ocgb.fields)+len(ocgb.fns))
		for _, f := range ocgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ocgb.fields...)...)
}

// OutcomeCategorySelect is the builder for selecting fields of OutcomeCategory entities.
type OutcomeCategorySelect struct {
	*OutcomeCategoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ocs *OutcomeCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ocs.prepareQuery(ctx); err != nil {
		return err
	}
	ocs.sql = ocs.OutcomeCategoryQuery.sqlQuery(ctx)
	return ocs.sqlScan(ctx, v)
}

func (ocs *OutcomeCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ocs.sql.Query()
	if err := ocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
