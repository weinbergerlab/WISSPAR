// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// MoreInfoModuleQuery is the builder for querying MoreInfoModule entities.
type MoreInfoModuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MoreInfoModule
	// eager-loading edges.
	withParent           *ResultsDefinitionQuery
	withCertainAgreement *CertainAgreementQuery
	withPointOfContact   *PointOfContactQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MoreInfoModuleQuery builder.
func (mimq *MoreInfoModuleQuery) Where(ps ...predicate.MoreInfoModule) *MoreInfoModuleQuery {
	mimq.predicates = append(mimq.predicates, ps...)
	return mimq
}

// Limit adds a limit step to the query.
func (mimq *MoreInfoModuleQuery) Limit(limit int) *MoreInfoModuleQuery {
	mimq.limit = &limit
	return mimq
}

// Offset adds an offset step to the query.
func (mimq *MoreInfoModuleQuery) Offset(offset int) *MoreInfoModuleQuery {
	mimq.offset = &offset
	return mimq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mimq *MoreInfoModuleQuery) Unique(unique bool) *MoreInfoModuleQuery {
	mimq.unique = &unique
	return mimq
}

// Order adds an order step to the query.
func (mimq *MoreInfoModuleQuery) Order(o ...OrderFunc) *MoreInfoModuleQuery {
	mimq.order = append(mimq.order, o...)
	return mimq
}

// QueryParent chains the current query on the "parent" edge.
func (mimq *MoreInfoModuleQuery) QueryParent() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: mimq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mimq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mimq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, moreinfomodule.ParentTable, moreinfomodule.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mimq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCertainAgreement chains the current query on the "certain_agreement" edge.
func (mimq *MoreInfoModuleQuery) QueryCertainAgreement() *CertainAgreementQuery {
	query := &CertainAgreementQuery{config: mimq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mimq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mimq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, selector),
			sqlgraph.To(certainagreement.Table, certainagreement.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, moreinfomodule.CertainAgreementTable, moreinfomodule.CertainAgreementColumn),
		)
		fromU = sqlgraph.SetNeighbors(mimq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPointOfContact chains the current query on the "point_of_contact" edge.
func (mimq *MoreInfoModuleQuery) QueryPointOfContact() *PointOfContactQuery {
	query := &PointOfContactQuery{config: mimq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mimq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mimq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, selector),
			sqlgraph.To(pointofcontact.Table, pointofcontact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, moreinfomodule.PointOfContactTable, moreinfomodule.PointOfContactColumn),
		)
		fromU = sqlgraph.SetNeighbors(mimq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MoreInfoModule entity from the query.
// Returns a *NotFoundError when no MoreInfoModule was found.
func (mimq *MoreInfoModuleQuery) First(ctx context.Context) (*MoreInfoModule, error) {
	nodes, err := mimq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{moreinfomodule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) FirstX(ctx context.Context) *MoreInfoModule {
	node, err := mimq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MoreInfoModule ID from the query.
// Returns a *NotFoundError when no MoreInfoModule ID was found.
func (mimq *MoreInfoModuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mimq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{moreinfomodule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) FirstIDX(ctx context.Context) int {
	id, err := mimq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MoreInfoModule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MoreInfoModule entity is found.
// Returns a *NotFoundError when no MoreInfoModule entities are found.
func (mimq *MoreInfoModuleQuery) Only(ctx context.Context) (*MoreInfoModule, error) {
	nodes, err := mimq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{moreinfomodule.Label}
	default:
		return nil, &NotSingularError{moreinfomodule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) OnlyX(ctx context.Context) *MoreInfoModule {
	node, err := mimq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MoreInfoModule ID in the query.
// Returns a *NotSingularError when more than one MoreInfoModule ID is found.
// Returns a *NotFoundError when no entities are found.
func (mimq *MoreInfoModuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mimq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{moreinfomodule.Label}
	default:
		err = &NotSingularError{moreinfomodule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := mimq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MoreInfoModules.
func (mimq *MoreInfoModuleQuery) All(ctx context.Context) ([]*MoreInfoModule, error) {
	if err := mimq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mimq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) AllX(ctx context.Context) []*MoreInfoModule {
	nodes, err := mimq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MoreInfoModule IDs.
func (mimq *MoreInfoModuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mimq.Select(moreinfomodule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) IDsX(ctx context.Context) []int {
	ids, err := mimq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mimq *MoreInfoModuleQuery) Count(ctx context.Context) (int, error) {
	if err := mimq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mimq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) CountX(ctx context.Context) int {
	count, err := mimq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mimq *MoreInfoModuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := mimq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mimq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mimq *MoreInfoModuleQuery) ExistX(ctx context.Context) bool {
	exist, err := mimq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MoreInfoModuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mimq *MoreInfoModuleQuery) Clone() *MoreInfoModuleQuery {
	if mimq == nil {
		return nil
	}
	return &MoreInfoModuleQuery{
		config:               mimq.config,
		limit:                mimq.limit,
		offset:               mimq.offset,
		order:                append([]OrderFunc{}, mimq.order...),
		predicates:           append([]predicate.MoreInfoModule{}, mimq.predicates...),
		withParent:           mimq.withParent.Clone(),
		withCertainAgreement: mimq.withCertainAgreement.Clone(),
		withPointOfContact:   mimq.withPointOfContact.Clone(),
		// clone intermediate query.
		sql:    mimq.sql.Clone(),
		path:   mimq.path,
		unique: mimq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (mimq *MoreInfoModuleQuery) WithParent(opts ...func(*ResultsDefinitionQuery)) *MoreInfoModuleQuery {
	query := &ResultsDefinitionQuery{config: mimq.config}
	for _, opt := range opts {
		opt(query)
	}
	mimq.withParent = query
	return mimq
}

// WithCertainAgreement tells the query-builder to eager-load the nodes that are connected to
// the "certain_agreement" edge. The optional arguments are used to configure the query builder of the edge.
func (mimq *MoreInfoModuleQuery) WithCertainAgreement(opts ...func(*CertainAgreementQuery)) *MoreInfoModuleQuery {
	query := &CertainAgreementQuery{config: mimq.config}
	for _, opt := range opts {
		opt(query)
	}
	mimq.withCertainAgreement = query
	return mimq
}

// WithPointOfContact tells the query-builder to eager-load the nodes that are connected to
// the "point_of_contact" edge. The optional arguments are used to configure the query builder of the edge.
func (mimq *MoreInfoModuleQuery) WithPointOfContact(opts ...func(*PointOfContactQuery)) *MoreInfoModuleQuery {
	query := &PointOfContactQuery{config: mimq.config}
	for _, opt := range opts {
		opt(query)
	}
	mimq.withPointOfContact = query
	return mimq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LimitationsAndCaveatsDescription string `json:"limitations_and_caveats_description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MoreInfoModule.Query().
//		GroupBy(moreinfomodule.FieldLimitationsAndCaveatsDescription).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (mimq *MoreInfoModuleQuery) GroupBy(field string, fields ...string) *MoreInfoModuleGroupBy {
	grbuild := &MoreInfoModuleGroupBy{config: mimq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mimq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mimq.sqlQuery(ctx), nil
	}
	grbuild.label = moreinfomodule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LimitationsAndCaveatsDescription string `json:"limitations_and_caveats_description,omitempty"`
//	}
//
//	client.MoreInfoModule.Query().
//		Select(moreinfomodule.FieldLimitationsAndCaveatsDescription).
//		Scan(ctx, &v)
//
func (mimq *MoreInfoModuleQuery) Select(fields ...string) *MoreInfoModuleSelect {
	mimq.fields = append(mimq.fields, fields...)
	selbuild := &MoreInfoModuleSelect{MoreInfoModuleQuery: mimq}
	selbuild.label = moreinfomodule.Label
	selbuild.flds, selbuild.scan = &mimq.fields, selbuild.Scan
	return selbuild
}

func (mimq *MoreInfoModuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mimq.fields {
		if !moreinfomodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if mimq.path != nil {
		prev, err := mimq.path(ctx)
		if err != nil {
			return err
		}
		mimq.sql = prev
	}
	return nil
}

func (mimq *MoreInfoModuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MoreInfoModule, error) {
	var (
		nodes       = []*MoreInfoModule{}
		withFKs     = mimq.withFKs
		_spec       = mimq.querySpec()
		loadedTypes = [3]bool{
			mimq.withParent != nil,
			mimq.withCertainAgreement != nil,
			mimq.withPointOfContact != nil,
		}
	)
	if mimq.withParent != nil || mimq.withCertainAgreement != nil || mimq.withPointOfContact != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, moreinfomodule.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MoreInfoModule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MoreInfoModule{config: mimq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mimq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mimq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MoreInfoModule)
		for i := range nodes {
			if nodes[i].results_definition_more_info_module == nil {
				continue
			}
			fk := *nodes[i].results_definition_more_info_module
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
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_more_info_module" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := mimq.withCertainAgreement; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MoreInfoModule)
		for i := range nodes {
			if nodes[i].more_info_module_certain_agreement == nil {
				continue
			}
			fk := *nodes[i].more_info_module_certain_agreement
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(certainagreement.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "more_info_module_certain_agreement" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CertainAgreement = n
			}
		}
	}

	if query := mimq.withPointOfContact; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MoreInfoModule)
		for i := range nodes {
			if nodes[i].more_info_module_point_of_contact == nil {
				continue
			}
			fk := *nodes[i].more_info_module_point_of_contact
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(pointofcontact.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "more_info_module_point_of_contact" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.PointOfContact = n
			}
		}
	}

	return nodes, nil
}

func (mimq *MoreInfoModuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mimq.querySpec()
	_spec.Node.Columns = mimq.fields
	if len(mimq.fields) > 0 {
		_spec.Unique = mimq.unique != nil && *mimq.unique
	}
	return sqlgraph.CountNodes(ctx, mimq.driver, _spec)
}

func (mimq *MoreInfoModuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mimq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (mimq *MoreInfoModuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moreinfomodule.Table,
			Columns: moreinfomodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moreinfomodule.FieldID,
			},
		},
		From:   mimq.sql,
		Unique: true,
	}
	if unique := mimq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mimq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moreinfomodule.FieldID)
		for i := range fields {
			if fields[i] != moreinfomodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mimq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mimq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mimq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mimq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mimq *MoreInfoModuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mimq.driver.Dialect())
	t1 := builder.Table(moreinfomodule.Table)
	columns := mimq.fields
	if len(columns) == 0 {
		columns = moreinfomodule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mimq.sql != nil {
		selector = mimq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mimq.unique != nil && *mimq.unique {
		selector.Distinct()
	}
	for _, p := range mimq.predicates {
		p(selector)
	}
	for _, p := range mimq.order {
		p(selector)
	}
	if offset := mimq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mimq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MoreInfoModuleGroupBy is the group-by builder for MoreInfoModule entities.
type MoreInfoModuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mimgb *MoreInfoModuleGroupBy) Aggregate(fns ...AggregateFunc) *MoreInfoModuleGroupBy {
	mimgb.fns = append(mimgb.fns, fns...)
	return mimgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mimgb *MoreInfoModuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mimgb.path(ctx)
	if err != nil {
		return err
	}
	mimgb.sql = query
	return mimgb.sqlScan(ctx, v)
}

func (mimgb *MoreInfoModuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mimgb.fields {
		if !moreinfomodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mimgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mimgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mimgb *MoreInfoModuleGroupBy) sqlQuery() *sql.Selector {
	selector := mimgb.sql.Select()
	aggregation := make([]string, 0, len(mimgb.fns))
	for _, fn := range mimgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mimgb.fields)+len(mimgb.fns))
		for _, f := range mimgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mimgb.fields...)...)
}

// MoreInfoModuleSelect is the builder for selecting fields of MoreInfoModule entities.
type MoreInfoModuleSelect struct {
	*MoreInfoModuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mims *MoreInfoModuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mims.prepareQuery(ctx); err != nil {
		return err
	}
	mims.sql = mims.MoreInfoModuleQuery.sqlQuery(ctx)
	return mims.sqlScan(ctx, v)
}

func (mims *MoreInfoModuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mims.sql.Query()
	if err := mims.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
