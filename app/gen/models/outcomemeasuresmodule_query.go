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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// OutcomeMeasuresModuleQuery is the builder for querying OutcomeMeasuresModule entities.
type OutcomeMeasuresModuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeMeasuresModule
	// eager-loading edges.
	withParent             *ResultsDefinitionQuery
	withOutcomeMeasureList *OutcomeMeasureQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeMeasuresModuleQuery builder.
func (ommq *OutcomeMeasuresModuleQuery) Where(ps ...predicate.OutcomeMeasuresModule) *OutcomeMeasuresModuleQuery {
	ommq.predicates = append(ommq.predicates, ps...)
	return ommq
}

// Limit adds a limit step to the query.
func (ommq *OutcomeMeasuresModuleQuery) Limit(limit int) *OutcomeMeasuresModuleQuery {
	ommq.limit = &limit
	return ommq
}

// Offset adds an offset step to the query.
func (ommq *OutcomeMeasuresModuleQuery) Offset(offset int) *OutcomeMeasuresModuleQuery {
	ommq.offset = &offset
	return ommq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ommq *OutcomeMeasuresModuleQuery) Unique(unique bool) *OutcomeMeasuresModuleQuery {
	ommq.unique = &unique
	return ommq
}

// Order adds an order step to the query.
func (ommq *OutcomeMeasuresModuleQuery) Order(o ...OrderFunc) *OutcomeMeasuresModuleQuery {
	ommq.order = append(ommq.order, o...)
	return ommq
}

// QueryParent chains the current query on the "parent" edge.
func (ommq *OutcomeMeasuresModuleQuery) QueryParent() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: ommq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ommq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ommq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, outcomemeasuresmodule.ParentTable, outcomemeasuresmodule.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ommq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeMeasureList chains the current query on the "outcome_measure_list" edge.
func (ommq *OutcomeMeasuresModuleQuery) QueryOutcomeMeasureList() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: ommq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ommq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ommq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasuresmodule.OutcomeMeasureListTable, outcomemeasuresmodule.OutcomeMeasureListColumn),
		)
		fromU = sqlgraph.SetNeighbors(ommq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeMeasuresModule entity from the query.
// Returns a *NotFoundError when no OutcomeMeasuresModule was found.
func (ommq *OutcomeMeasuresModuleQuery) First(ctx context.Context) (*OutcomeMeasuresModule, error) {
	nodes, err := ommq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomemeasuresmodule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) FirstX(ctx context.Context) *OutcomeMeasuresModule {
	node, err := ommq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeMeasuresModule ID from the query.
// Returns a *NotFoundError when no OutcomeMeasuresModule ID was found.
func (ommq *OutcomeMeasuresModuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ommq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomemeasuresmodule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) FirstIDX(ctx context.Context) int {
	id, err := ommq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeMeasuresModule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeMeasuresModule entity is found.
// Returns a *NotFoundError when no OutcomeMeasuresModule entities are found.
func (ommq *OutcomeMeasuresModuleQuery) Only(ctx context.Context) (*OutcomeMeasuresModule, error) {
	nodes, err := ommq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomemeasuresmodule.Label}
	default:
		return nil, &NotSingularError{outcomemeasuresmodule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) OnlyX(ctx context.Context) *OutcomeMeasuresModule {
	node, err := ommq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeMeasuresModule ID in the query.
// Returns a *NotSingularError when more than one OutcomeMeasuresModule ID is found.
// Returns a *NotFoundError when no entities are found.
func (ommq *OutcomeMeasuresModuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ommq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomemeasuresmodule.Label}
	default:
		err = &NotSingularError{outcomemeasuresmodule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := ommq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeMeasuresModules.
func (ommq *OutcomeMeasuresModuleQuery) All(ctx context.Context) ([]*OutcomeMeasuresModule, error) {
	if err := ommq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ommq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) AllX(ctx context.Context) []*OutcomeMeasuresModule {
	nodes, err := ommq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeMeasuresModule IDs.
func (ommq *OutcomeMeasuresModuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ommq.Select(outcomemeasuresmodule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) IDsX(ctx context.Context) []int {
	ids, err := ommq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ommq *OutcomeMeasuresModuleQuery) Count(ctx context.Context) (int, error) {
	if err := ommq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ommq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) CountX(ctx context.Context) int {
	count, err := ommq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ommq *OutcomeMeasuresModuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := ommq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ommq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ommq *OutcomeMeasuresModuleQuery) ExistX(ctx context.Context) bool {
	exist, err := ommq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeMeasuresModuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ommq *OutcomeMeasuresModuleQuery) Clone() *OutcomeMeasuresModuleQuery {
	if ommq == nil {
		return nil
	}
	return &OutcomeMeasuresModuleQuery{
		config:                 ommq.config,
		limit:                  ommq.limit,
		offset:                 ommq.offset,
		order:                  append([]OrderFunc{}, ommq.order...),
		predicates:             append([]predicate.OutcomeMeasuresModule{}, ommq.predicates...),
		withParent:             ommq.withParent.Clone(),
		withOutcomeMeasureList: ommq.withOutcomeMeasureList.Clone(),
		// clone intermediate query.
		sql:    ommq.sql.Clone(),
		path:   ommq.path,
		unique: ommq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ommq *OutcomeMeasuresModuleQuery) WithParent(opts ...func(*ResultsDefinitionQuery)) *OutcomeMeasuresModuleQuery {
	query := &ResultsDefinitionQuery{config: ommq.config}
	for _, opt := range opts {
		opt(query)
	}
	ommq.withParent = query
	return ommq
}

// WithOutcomeMeasureList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_measure_list" edge. The optional arguments are used to configure the query builder of the edge.
func (ommq *OutcomeMeasuresModuleQuery) WithOutcomeMeasureList(opts ...func(*OutcomeMeasureQuery)) *OutcomeMeasuresModuleQuery {
	query := &OutcomeMeasureQuery{config: ommq.config}
	for _, opt := range opts {
		opt(query)
	}
	ommq.withOutcomeMeasureList = query
	return ommq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ommq *OutcomeMeasuresModuleQuery) GroupBy(field string, fields ...string) *OutcomeMeasuresModuleGroupBy {
	grbuild := &OutcomeMeasuresModuleGroupBy{config: ommq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ommq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ommq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomemeasuresmodule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ommq *OutcomeMeasuresModuleQuery) Select(fields ...string) *OutcomeMeasuresModuleSelect {
	ommq.fields = append(ommq.fields, fields...)
	selbuild := &OutcomeMeasuresModuleSelect{OutcomeMeasuresModuleQuery: ommq}
	selbuild.label = outcomemeasuresmodule.Label
	selbuild.flds, selbuild.scan = &ommq.fields, selbuild.Scan
	return selbuild
}

func (ommq *OutcomeMeasuresModuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ommq.fields {
		if !outcomemeasuresmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ommq.path != nil {
		prev, err := ommq.path(ctx)
		if err != nil {
			return err
		}
		ommq.sql = prev
	}
	return nil
}

func (ommq *OutcomeMeasuresModuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeMeasuresModule, error) {
	var (
		nodes       = []*OutcomeMeasuresModule{}
		withFKs     = ommq.withFKs
		_spec       = ommq.querySpec()
		loadedTypes = [2]bool{
			ommq.withParent != nil,
			ommq.withOutcomeMeasureList != nil,
		}
	)
	if ommq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasuresmodule.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeMeasuresModule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeMeasuresModule{config: ommq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ommq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ommq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeMeasuresModule)
		for i := range nodes {
			if nodes[i].results_definition_outcome_measures_module == nil {
				continue
			}
			fk := *nodes[i].results_definition_outcome_measures_module
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(resultsdefinition.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_outcome_measures_module" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ommq.withOutcomeMeasureList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasuresModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeMeasureList = []*OutcomeMeasure{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeMeasure(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasuresmodule.OutcomeMeasureListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measures_module_outcome_measure_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measures_module_outcome_measure_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measures_module_outcome_measure_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeMeasureList = append(node.Edges.OutcomeMeasureList, n)
		}
	}

	return nodes, nil
}

func (ommq *OutcomeMeasuresModuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ommq.querySpec()
	_spec.Node.Columns = ommq.fields
	if len(ommq.fields) > 0 {
		_spec.Unique = ommq.unique != nil && *ommq.unique
	}
	return sqlgraph.CountNodes(ctx, ommq.driver, _spec)
}

func (ommq *OutcomeMeasuresModuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ommq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ommq *OutcomeMeasuresModuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasuresmodule.Table,
			Columns: outcomemeasuresmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasuresmodule.FieldID,
			},
		},
		From:   ommq.sql,
		Unique: true,
	}
	if unique := ommq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ommq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasuresmodule.FieldID)
		for i := range fields {
			if fields[i] != outcomemeasuresmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ommq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ommq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ommq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ommq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ommq *OutcomeMeasuresModuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ommq.driver.Dialect())
	t1 := builder.Table(outcomemeasuresmodule.Table)
	columns := ommq.fields
	if len(columns) == 0 {
		columns = outcomemeasuresmodule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ommq.sql != nil {
		selector = ommq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ommq.unique != nil && *ommq.unique {
		selector.Distinct()
	}
	for _, p := range ommq.predicates {
		p(selector)
	}
	for _, p := range ommq.order {
		p(selector)
	}
	if offset := ommq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ommq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeMeasuresModuleGroupBy is the group-by builder for OutcomeMeasuresModule entities.
type OutcomeMeasuresModuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ommgb *OutcomeMeasuresModuleGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeMeasuresModuleGroupBy {
	ommgb.fns = append(ommgb.fns, fns...)
	return ommgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ommgb *OutcomeMeasuresModuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ommgb.path(ctx)
	if err != nil {
		return err
	}
	ommgb.sql = query
	return ommgb.sqlScan(ctx, v)
}

func (ommgb *OutcomeMeasuresModuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ommgb.fields {
		if !outcomemeasuresmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ommgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ommgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ommgb *OutcomeMeasuresModuleGroupBy) sqlQuery() *sql.Selector {
	selector := ommgb.sql.Select()
	aggregation := make([]string, 0, len(ommgb.fns))
	for _, fn := range ommgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ommgb.fields)+len(ommgb.fns))
		for _, f := range ommgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ommgb.fields...)...)
}

// OutcomeMeasuresModuleSelect is the builder for selecting fields of OutcomeMeasuresModule entities.
type OutcomeMeasuresModuleSelect struct {
	*OutcomeMeasuresModuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (omms *OutcomeMeasuresModuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := omms.prepareQuery(ctx); err != nil {
		return err
	}
	omms.sql = omms.OutcomeMeasuresModuleQuery.sqlQuery(ctx)
	return omms.sqlScan(ctx, v)
}

func (omms *OutcomeMeasuresModuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := omms.sql.Query()
	if err := omms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
