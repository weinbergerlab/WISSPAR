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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineCategoryQuery is the builder for querying BaselineCategory entities.
type BaselineCategoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineCategory
	// eager-loading edges.
	withParent                  *BaselineClassQuery
	withBaselineMeasurementList *BaselineMeasurementQuery
	withFKs                     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineCategoryQuery builder.
func (bcq *BaselineCategoryQuery) Where(ps ...predicate.BaselineCategory) *BaselineCategoryQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit adds a limit step to the query.
func (bcq *BaselineCategoryQuery) Limit(limit int) *BaselineCategoryQuery {
	bcq.limit = &limit
	return bcq
}

// Offset adds an offset step to the query.
func (bcq *BaselineCategoryQuery) Offset(offset int) *BaselineCategoryQuery {
	bcq.offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BaselineCategoryQuery) Unique(unique bool) *BaselineCategoryQuery {
	bcq.unique = &unique
	return bcq
}

// Order adds an order step to the query.
func (bcq *BaselineCategoryQuery) Order(o ...OrderFunc) *BaselineCategoryQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryParent chains the current query on the "parent" edge.
func (bcq *BaselineCategoryQuery) QueryParent() *BaselineClassQuery {
	query := &BaselineClassQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecategory.Table, baselinecategory.FieldID, selector),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinecategory.ParentTable, baselinecategory.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineMeasurementList chains the current query on the "baseline_measurement_list" edge.
func (bcq *BaselineCategoryQuery) QueryBaselineMeasurementList() *BaselineMeasurementQuery {
	query := &BaselineMeasurementQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecategory.Table, baselinecategory.FieldID, selector),
			sqlgraph.To(baselinemeasurement.Table, baselinemeasurement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecategory.BaselineMeasurementListTable, baselinecategory.BaselineMeasurementListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineCategory entity from the query.
// Returns a *NotFoundError when no BaselineCategory was found.
func (bcq *BaselineCategoryQuery) First(ctx context.Context) (*BaselineCategory, error) {
	nodes, err := bcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) FirstX(ctx context.Context) *BaselineCategory {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineCategory ID from the query.
// Returns a *NotFoundError when no BaselineCategory ID was found.
func (bcq *BaselineCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineCategory entity is found.
// Returns a *NotFoundError when no BaselineCategory entities are found.
func (bcq *BaselineCategoryQuery) Only(ctx context.Context) (*BaselineCategory, error) {
	nodes, err := bcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinecategory.Label}
	default:
		return nil, &NotSingularError{baselinecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) OnlyX(ctx context.Context) *BaselineCategory {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineCategory ID in the query.
// Returns a *NotSingularError when more than one BaselineCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BaselineCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinecategory.Label}
	default:
		err = &NotSingularError{baselinecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineCategories.
func (bcq *BaselineCategoryQuery) All(ctx context.Context) ([]*BaselineCategory, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) AllX(ctx context.Context) []*BaselineCategory {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineCategory IDs.
func (bcq *BaselineCategoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcq.Select(baselinecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BaselineCategoryQuery) Count(ctx context.Context) (int, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BaselineCategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BaselineCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BaselineCategoryQuery) Clone() *BaselineCategoryQuery {
	if bcq == nil {
		return nil
	}
	return &BaselineCategoryQuery{
		config:                      bcq.config,
		limit:                       bcq.limit,
		offset:                      bcq.offset,
		order:                       append([]OrderFunc{}, bcq.order...),
		predicates:                  append([]predicate.BaselineCategory{}, bcq.predicates...),
		withParent:                  bcq.withParent.Clone(),
		withBaselineMeasurementList: bcq.withBaselineMeasurementList.Clone(),
		// clone intermediate query.
		sql:    bcq.sql.Clone(),
		path:   bcq.path,
		unique: bcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BaselineCategoryQuery) WithParent(opts ...func(*BaselineClassQuery)) *BaselineCategoryQuery {
	query := &BaselineClassQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withParent = query
	return bcq
}

// WithBaselineMeasurementList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_measurement_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BaselineCategoryQuery) WithBaselineMeasurementList(opts ...func(*BaselineMeasurementQuery)) *BaselineCategoryQuery {
	query := &BaselineMeasurementQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBaselineMeasurementList = query
	return bcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineCategoryTitle string `json:"baseline_category_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineCategory.Query().
//		GroupBy(baselinecategory.FieldBaselineCategoryTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bcq *BaselineCategoryQuery) GroupBy(field string, fields ...string) *BaselineCategoryGroupBy {
	grbuild := &BaselineCategoryGroupBy{config: bcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinecategory.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineCategoryTitle string `json:"baseline_category_title,omitempty"`
//	}
//
//	client.BaselineCategory.Query().
//		Select(baselinecategory.FieldBaselineCategoryTitle).
//		Scan(ctx, &v)
//
func (bcq *BaselineCategoryQuery) Select(fields ...string) *BaselineCategorySelect {
	bcq.fields = append(bcq.fields, fields...)
	selbuild := &BaselineCategorySelect{BaselineCategoryQuery: bcq}
	selbuild.label = baselinecategory.Label
	selbuild.flds, selbuild.scan = &bcq.fields, selbuild.Scan
	return selbuild
}

func (bcq *BaselineCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcq.fields {
		if !baselinecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BaselineCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineCategory, error) {
	var (
		nodes       = []*BaselineCategory{}
		withFKs     = bcq.withFKs
		_spec       = bcq.querySpec()
		loadedTypes = [2]bool{
			bcq.withParent != nil,
			bcq.withBaselineMeasurementList != nil,
		}
	)
	if bcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecategory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineCategory{config: bcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineCategory)
		for i := range nodes {
			if nodes[i].baseline_class_baseline_category_list == nil {
				continue
			}
			fk := *nodes[i].baseline_class_baseline_category_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselineclass.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_baseline_category_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bcq.withBaselineMeasurementList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineCategory)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineMeasurementList = []*BaselineMeasurement{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineMeasurement(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinecategory.BaselineMeasurementListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_category_baseline_measurement_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_category_baseline_measurement_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_category_baseline_measurement_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineMeasurementList = append(node.Edges.BaselineMeasurementList, n)
		}
	}

	return nodes, nil
}

func (bcq *BaselineCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	_spec.Node.Columns = bcq.fields
	if len(bcq.fields) > 0 {
		_spec.Unique = bcq.unique != nil && *bcq.unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BaselineCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcq *BaselineCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecategory.Table,
			Columns: baselinecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecategory.FieldID,
			},
		},
		From:   bcq.sql,
		Unique: true,
	}
	if unique := bcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecategory.FieldID)
		for i := range fields {
			if fields[i] != baselinecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BaselineCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(baselinecategory.Table)
	columns := bcq.fields
	if len(columns) == 0 {
		columns = baselinecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcq.unique != nil && *bcq.unique {
		selector.Distinct()
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineCategoryGroupBy is the group-by builder for BaselineCategory entities.
type BaselineCategoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BaselineCategoryGroupBy) Aggregate(fns ...AggregateFunc) *BaselineCategoryGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcgb *BaselineCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcgb.path(ctx)
	if err != nil {
		return err
	}
	bcgb.sql = query
	return bcgb.sqlScan(ctx, v)
}

func (bcgb *BaselineCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcgb.fields {
		if !baselinecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcgb *BaselineCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := bcgb.sql.Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcgb.fields)+len(bcgb.fns))
		for _, f := range bcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcgb.fields...)...)
}

// BaselineCategorySelect is the builder for selecting fields of BaselineCategory entities.
type BaselineCategorySelect struct {
	*BaselineCategoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BaselineCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcs.sql = bcs.BaselineCategoryQuery.sqlQuery(ctx)
	return bcs.sqlScan(ctx, v)
}

func (bcs *BaselineCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcs.sql.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
