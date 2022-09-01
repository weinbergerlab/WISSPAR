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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassQuery is the builder for querying BaselineClass entities.
type BaselineClassQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineClass
	// eager-loading edges.
	withParent                 *BaselineMeasureQuery
	withBaselineClassDenomList *BaselineClassDenomQuery
	withBaselineCategoryList   *BaselineCategoryQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineClassQuery builder.
func (bcq *BaselineClassQuery) Where(ps ...predicate.BaselineClass) *BaselineClassQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit adds a limit step to the query.
func (bcq *BaselineClassQuery) Limit(limit int) *BaselineClassQuery {
	bcq.limit = &limit
	return bcq
}

// Offset adds an offset step to the query.
func (bcq *BaselineClassQuery) Offset(offset int) *BaselineClassQuery {
	bcq.offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BaselineClassQuery) Unique(unique bool) *BaselineClassQuery {
	bcq.unique = &unique
	return bcq
}

// Order adds an order step to the query.
func (bcq *BaselineClassQuery) Order(o ...OrderFunc) *BaselineClassQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryParent chains the current query on the "parent" edge.
func (bcq *BaselineClassQuery) QueryParent() *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, selector),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclass.ParentTable, baselineclass.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineClassDenomList chains the current query on the "baseline_class_denom_list" edge.
func (bcq *BaselineClassQuery) QueryBaselineClassDenomList() *BaselineClassDenomQuery {
	query := &BaselineClassDenomQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, selector),
			sqlgraph.To(baselineclassdenom.Table, baselineclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclass.BaselineClassDenomListTable, baselineclass.BaselineClassDenomListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineCategoryList chains the current query on the "baseline_category_list" edge.
func (bcq *BaselineClassQuery) QueryBaselineCategoryList() *BaselineCategoryQuery {
	query := &BaselineCategoryQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, selector),
			sqlgraph.To(baselinecategory.Table, baselinecategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclass.BaselineCategoryListTable, baselineclass.BaselineCategoryListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineClass entity from the query.
// Returns a *NotFoundError when no BaselineClass was found.
func (bcq *BaselineClassQuery) First(ctx context.Context) (*BaselineClass, error) {
	nodes, err := bcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselineclass.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BaselineClassQuery) FirstX(ctx context.Context) *BaselineClass {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineClass ID from the query.
// Returns a *NotFoundError when no BaselineClass ID was found.
func (bcq *BaselineClassQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselineclass.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BaselineClassQuery) FirstIDX(ctx context.Context) int {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineClass entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineClass entity is found.
// Returns a *NotFoundError when no BaselineClass entities are found.
func (bcq *BaselineClassQuery) Only(ctx context.Context) (*BaselineClass, error) {
	nodes, err := bcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselineclass.Label}
	default:
		return nil, &NotSingularError{baselineclass.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BaselineClassQuery) OnlyX(ctx context.Context) *BaselineClass {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineClass ID in the query.
// Returns a *NotSingularError when more than one BaselineClass ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BaselineClassQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselineclass.Label}
	default:
		err = &NotSingularError{baselineclass.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BaselineClassQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineClasses.
func (bcq *BaselineClassQuery) All(ctx context.Context) ([]*BaselineClass, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BaselineClassQuery) AllX(ctx context.Context) []*BaselineClass {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineClass IDs.
func (bcq *BaselineClassQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcq.Select(baselineclass.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BaselineClassQuery) IDsX(ctx context.Context) []int {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BaselineClassQuery) Count(ctx context.Context) (int, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BaselineClassQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BaselineClassQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BaselineClassQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BaselineClassQuery) Clone() *BaselineClassQuery {
	if bcq == nil {
		return nil
	}
	return &BaselineClassQuery{
		config:                     bcq.config,
		limit:                      bcq.limit,
		offset:                     bcq.offset,
		order:                      append([]OrderFunc{}, bcq.order...),
		predicates:                 append([]predicate.BaselineClass{}, bcq.predicates...),
		withParent:                 bcq.withParent.Clone(),
		withBaselineClassDenomList: bcq.withBaselineClassDenomList.Clone(),
		withBaselineCategoryList:   bcq.withBaselineCategoryList.Clone(),
		// clone intermediate query.
		sql:    bcq.sql.Clone(),
		path:   bcq.path,
		unique: bcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BaselineClassQuery) WithParent(opts ...func(*BaselineMeasureQuery)) *BaselineClassQuery {
	query := &BaselineMeasureQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withParent = query
	return bcq
}

// WithBaselineClassDenomList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_class_denom_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BaselineClassQuery) WithBaselineClassDenomList(opts ...func(*BaselineClassDenomQuery)) *BaselineClassQuery {
	query := &BaselineClassDenomQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBaselineClassDenomList = query
	return bcq
}

// WithBaselineCategoryList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_category_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BaselineClassQuery) WithBaselineCategoryList(opts ...func(*BaselineCategoryQuery)) *BaselineClassQuery {
	query := &BaselineCategoryQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBaselineCategoryList = query
	return bcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineClassTitle string `json:"baseline_class_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineClass.Query().
//		GroupBy(baselineclass.FieldBaselineClassTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bcq *BaselineClassQuery) GroupBy(field string, fields ...string) *BaselineClassGroupBy {
	grbuild := &BaselineClassGroupBy{config: bcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcq.sqlQuery(ctx), nil
	}
	grbuild.label = baselineclass.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineClassTitle string `json:"baseline_class_title,omitempty"`
//	}
//
//	client.BaselineClass.Query().
//		Select(baselineclass.FieldBaselineClassTitle).
//		Scan(ctx, &v)
//
func (bcq *BaselineClassQuery) Select(fields ...string) *BaselineClassSelect {
	bcq.fields = append(bcq.fields, fields...)
	selbuild := &BaselineClassSelect{BaselineClassQuery: bcq}
	selbuild.label = baselineclass.Label
	selbuild.flds, selbuild.scan = &bcq.fields, selbuild.Scan
	return selbuild
}

func (bcq *BaselineClassQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcq.fields {
		if !baselineclass.ValidColumn(f) {
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

func (bcq *BaselineClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineClass, error) {
	var (
		nodes       = []*BaselineClass{}
		withFKs     = bcq.withFKs
		_spec       = bcq.querySpec()
		loadedTypes = [3]bool{
			bcq.withParent != nil,
			bcq.withBaselineClassDenomList != nil,
			bcq.withBaselineCategoryList != nil,
		}
	)
	if bcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclass.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineClass).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineClass{config: bcq.config}
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
		nodeids := make(map[int][]*BaselineClass)
		for i := range nodes {
			if nodes[i].baseline_measure_baseline_class_list == nil {
				continue
			}
			fk := *nodes[i].baseline_measure_baseline_class_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinemeasure.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_baseline_class_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bcq.withBaselineClassDenomList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineClass)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineClassDenomList = []*BaselineClassDenom{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineClassDenom(func(s *sql.Selector) {
			s.Where(sql.InValues(baselineclass.BaselineClassDenomListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_class_baseline_class_denom_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_class_baseline_class_denom_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_baseline_class_denom_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineClassDenomList = append(node.Edges.BaselineClassDenomList, n)
		}
	}

	if query := bcq.withBaselineCategoryList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineClass)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineCategoryList = []*BaselineCategory{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineCategory(func(s *sql.Selector) {
			s.Where(sql.InValues(baselineclass.BaselineCategoryListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_class_baseline_category_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_class_baseline_category_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_baseline_category_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineCategoryList = append(node.Edges.BaselineCategoryList, n)
		}
	}

	return nodes, nil
}

func (bcq *BaselineClassQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	_spec.Node.Columns = bcq.fields
	if len(bcq.fields) > 0 {
		_spec.Unique = bcq.unique != nil && *bcq.unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BaselineClassQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcq *BaselineClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclass.Table,
			Columns: baselineclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclass.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclass.FieldID)
		for i := range fields {
			if fields[i] != baselineclass.FieldID {
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

func (bcq *BaselineClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(baselineclass.Table)
	columns := bcq.fields
	if len(columns) == 0 {
		columns = baselineclass.Columns
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

// BaselineClassGroupBy is the group-by builder for BaselineClass entities.
type BaselineClassGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BaselineClassGroupBy) Aggregate(fns ...AggregateFunc) *BaselineClassGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcgb *BaselineClassGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcgb.path(ctx)
	if err != nil {
		return err
	}
	bcgb.sql = query
	return bcgb.sqlScan(ctx, v)
}

func (bcgb *BaselineClassGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcgb.fields {
		if !baselineclass.ValidColumn(f) {
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

func (bcgb *BaselineClassGroupBy) sqlQuery() *sql.Selector {
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

// BaselineClassSelect is the builder for selecting fields of BaselineClass entities.
type BaselineClassSelect struct {
	*BaselineClassQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BaselineClassSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcs.sql = bcs.BaselineClassQuery.sqlQuery(ctx)
	return bcs.sqlScan(ctx, v)
}

func (bcs *BaselineClassSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcs.sql.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
