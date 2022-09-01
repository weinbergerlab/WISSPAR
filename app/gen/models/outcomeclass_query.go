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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeClassQuery is the builder for querying OutcomeClass entities.
type OutcomeClassQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeClass
	// eager-loading edges.
	withParent                *OutcomeMeasureQuery
	withOutcomeClassDenomList *OutcomeClassDenomQuery
	withOutcomeCategoryList   *OutcomeCategoryQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeClassQuery builder.
func (ocq *OutcomeClassQuery) Where(ps ...predicate.OutcomeClass) *OutcomeClassQuery {
	ocq.predicates = append(ocq.predicates, ps...)
	return ocq
}

// Limit adds a limit step to the query.
func (ocq *OutcomeClassQuery) Limit(limit int) *OutcomeClassQuery {
	ocq.limit = &limit
	return ocq
}

// Offset adds an offset step to the query.
func (ocq *OutcomeClassQuery) Offset(offset int) *OutcomeClassQuery {
	ocq.offset = &offset
	return ocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocq *OutcomeClassQuery) Unique(unique bool) *OutcomeClassQuery {
	ocq.unique = &unique
	return ocq
}

// Order adds an order step to the query.
func (ocq *OutcomeClassQuery) Order(o ...OrderFunc) *OutcomeClassQuery {
	ocq.order = append(ocq.order, o...)
	return ocq
}

// QueryParent chains the current query on the "parent" edge.
func (ocq *OutcomeClassQuery) QueryParent() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: ocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclass.ParentTable, outcomeclass.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeClassDenomList chains the current query on the "outcome_class_denom_list" edge.
func (ocq *OutcomeClassQuery) QueryOutcomeClassDenomList() *OutcomeClassDenomQuery {
	query := &OutcomeClassDenomQuery{config: ocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, selector),
			sqlgraph.To(outcomeclassdenom.Table, outcomeclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclass.OutcomeClassDenomListTable, outcomeclass.OutcomeClassDenomListColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeCategoryList chains the current query on the "outcome_category_list" edge.
func (ocq *OutcomeClassQuery) QueryOutcomeCategoryList() *OutcomeCategoryQuery {
	query := &OutcomeCategoryQuery{config: ocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, selector),
			sqlgraph.To(outcomecategory.Table, outcomecategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclass.OutcomeCategoryListTable, outcomeclass.OutcomeCategoryListColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeClass entity from the query.
// Returns a *NotFoundError when no OutcomeClass was found.
func (ocq *OutcomeClassQuery) First(ctx context.Context) (*OutcomeClass, error) {
	nodes, err := ocq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeclass.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocq *OutcomeClassQuery) FirstX(ctx context.Context) *OutcomeClass {
	node, err := ocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeClass ID from the query.
// Returns a *NotFoundError when no OutcomeClass ID was found.
func (ocq *OutcomeClassQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeclass.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocq *OutcomeClassQuery) FirstIDX(ctx context.Context) int {
	id, err := ocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeClass entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeClass entity is found.
// Returns a *NotFoundError when no OutcomeClass entities are found.
func (ocq *OutcomeClassQuery) Only(ctx context.Context) (*OutcomeClass, error) {
	nodes, err := ocq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeclass.Label}
	default:
		return nil, &NotSingularError{outcomeclass.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocq *OutcomeClassQuery) OnlyX(ctx context.Context) *OutcomeClass {
	node, err := ocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeClass ID in the query.
// Returns a *NotSingularError when more than one OutcomeClass ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocq *OutcomeClassQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeclass.Label}
	default:
		err = &NotSingularError{outcomeclass.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocq *OutcomeClassQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeClasses.
func (ocq *OutcomeClassQuery) All(ctx context.Context) ([]*OutcomeClass, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ocq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ocq *OutcomeClassQuery) AllX(ctx context.Context) []*OutcomeClass {
	nodes, err := ocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeClass IDs.
func (ocq *OutcomeClassQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ocq.Select(outcomeclass.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocq *OutcomeClassQuery) IDsX(ctx context.Context) []int {
	ids, err := ocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocq *OutcomeClassQuery) Count(ctx context.Context) (int, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ocq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ocq *OutcomeClassQuery) CountX(ctx context.Context) int {
	count, err := ocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocq *OutcomeClassQuery) Exist(ctx context.Context) (bool, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ocq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ocq *OutcomeClassQuery) ExistX(ctx context.Context) bool {
	exist, err := ocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocq *OutcomeClassQuery) Clone() *OutcomeClassQuery {
	if ocq == nil {
		return nil
	}
	return &OutcomeClassQuery{
		config:                    ocq.config,
		limit:                     ocq.limit,
		offset:                    ocq.offset,
		order:                     append([]OrderFunc{}, ocq.order...),
		predicates:                append([]predicate.OutcomeClass{}, ocq.predicates...),
		withParent:                ocq.withParent.Clone(),
		withOutcomeClassDenomList: ocq.withOutcomeClassDenomList.Clone(),
		withOutcomeCategoryList:   ocq.withOutcomeCategoryList.Clone(),
		// clone intermediate query.
		sql:    ocq.sql.Clone(),
		path:   ocq.path,
		unique: ocq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OutcomeClassQuery) WithParent(opts ...func(*OutcomeMeasureQuery)) *OutcomeClassQuery {
	query := &OutcomeMeasureQuery{config: ocq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocq.withParent = query
	return ocq
}

// WithOutcomeClassDenomList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_class_denom_list" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OutcomeClassQuery) WithOutcomeClassDenomList(opts ...func(*OutcomeClassDenomQuery)) *OutcomeClassQuery {
	query := &OutcomeClassDenomQuery{config: ocq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocq.withOutcomeClassDenomList = query
	return ocq
}

// WithOutcomeCategoryList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_category_list" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OutcomeClassQuery) WithOutcomeCategoryList(opts ...func(*OutcomeCategoryQuery)) *OutcomeClassQuery {
	query := &OutcomeCategoryQuery{config: ocq.config}
	for _, opt := range opts {
		opt(query)
	}
	ocq.withOutcomeCategoryList = query
	return ocq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeClassTitle string `json:"outcome_class_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeClass.Query().
//		GroupBy(outcomeclass.FieldOutcomeClassTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ocq *OutcomeClassQuery) GroupBy(field string, fields ...string) *OutcomeClassGroupBy {
	grbuild := &OutcomeClassGroupBy{config: ocq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ocq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeclass.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeClassTitle string `json:"outcome_class_title,omitempty"`
//	}
//
//	client.OutcomeClass.Query().
//		Select(outcomeclass.FieldOutcomeClassTitle).
//		Scan(ctx, &v)
//
func (ocq *OutcomeClassQuery) Select(fields ...string) *OutcomeClassSelect {
	ocq.fields = append(ocq.fields, fields...)
	selbuild := &OutcomeClassSelect{OutcomeClassQuery: ocq}
	selbuild.label = outcomeclass.Label
	selbuild.flds, selbuild.scan = &ocq.fields, selbuild.Scan
	return selbuild
}

func (ocq *OutcomeClassQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ocq.fields {
		if !outcomeclass.ValidColumn(f) {
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

func (ocq *OutcomeClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeClass, error) {
	var (
		nodes       = []*OutcomeClass{}
		withFKs     = ocq.withFKs
		_spec       = ocq.querySpec()
		loadedTypes = [3]bool{
			ocq.withParent != nil,
			ocq.withOutcomeClassDenomList != nil,
			ocq.withOutcomeCategoryList != nil,
		}
	)
	if ocq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclass.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeClass).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeClass{config: ocq.config}
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
		nodeids := make(map[int][]*OutcomeClass)
		for i := range nodes {
			if nodes[i].outcome_measure_outcome_class_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measure_outcome_class_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomemeasure.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_class_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ocq.withOutcomeClassDenomList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeClass)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeClassDenomList = []*OutcomeClassDenom{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeClassDenom(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomeclass.OutcomeClassDenomListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_class_outcome_class_denom_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_class_outcome_class_denom_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_outcome_class_denom_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeClassDenomList = append(node.Edges.OutcomeClassDenomList, n)
		}
	}

	if query := ocq.withOutcomeCategoryList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeClass)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeCategoryList = []*OutcomeCategory{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeCategory(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomeclass.OutcomeCategoryListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_class_outcome_category_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_class_outcome_category_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_class_outcome_category_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeCategoryList = append(node.Edges.OutcomeCategoryList, n)
		}
	}

	return nodes, nil
}

func (ocq *OutcomeClassQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocq.querySpec()
	_spec.Node.Columns = ocq.fields
	if len(ocq.fields) > 0 {
		_spec.Unique = ocq.unique != nil && *ocq.unique
	}
	return sqlgraph.CountNodes(ctx, ocq.driver, _spec)
}

func (ocq *OutcomeClassQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ocq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ocq *OutcomeClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeclass.Table,
			Columns: outcomeclass.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeclass.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeclass.FieldID)
		for i := range fields {
			if fields[i] != outcomeclass.FieldID {
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

func (ocq *OutcomeClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocq.driver.Dialect())
	t1 := builder.Table(outcomeclass.Table)
	columns := ocq.fields
	if len(columns) == 0 {
		columns = outcomeclass.Columns
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

// OutcomeClassGroupBy is the group-by builder for OutcomeClass entities.
type OutcomeClassGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocgb *OutcomeClassGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeClassGroupBy {
	ocgb.fns = append(ocgb.fns, fns...)
	return ocgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ocgb *OutcomeClassGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ocgb.path(ctx)
	if err != nil {
		return err
	}
	ocgb.sql = query
	return ocgb.sqlScan(ctx, v)
}

func (ocgb *OutcomeClassGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ocgb.fields {
		if !outcomeclass.ValidColumn(f) {
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

func (ocgb *OutcomeClassGroupBy) sqlQuery() *sql.Selector {
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

// OutcomeClassSelect is the builder for selecting fields of OutcomeClass entities.
type OutcomeClassSelect struct {
	*OutcomeClassQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ocs *OutcomeClassSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ocs.prepareQuery(ctx); err != nil {
		return err
	}
	ocs.sql = ocs.OutcomeClassQuery.sqlQuery(ctx)
	return ocs.sqlScan(ctx, v)
}

func (ocs *OutcomeClassSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ocs.sql.Query()
	if err := ocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
