// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// PointOfContactQuery is the builder for querying PointOfContact entities.
type PointOfContactQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PointOfContact
	// eager-loading edges.
	withParent *MoreInfoModuleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PointOfContactQuery builder.
func (pocq *PointOfContactQuery) Where(ps ...predicate.PointOfContact) *PointOfContactQuery {
	pocq.predicates = append(pocq.predicates, ps...)
	return pocq
}

// Limit adds a limit step to the query.
func (pocq *PointOfContactQuery) Limit(limit int) *PointOfContactQuery {
	pocq.limit = &limit
	return pocq
}

// Offset adds an offset step to the query.
func (pocq *PointOfContactQuery) Offset(offset int) *PointOfContactQuery {
	pocq.offset = &offset
	return pocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pocq *PointOfContactQuery) Unique(unique bool) *PointOfContactQuery {
	pocq.unique = &unique
	return pocq
}

// Order adds an order step to the query.
func (pocq *PointOfContactQuery) Order(o ...OrderFunc) *PointOfContactQuery {
	pocq.order = append(pocq.order, o...)
	return pocq
}

// QueryParent chains the current query on the "parent" edge.
func (pocq *PointOfContactQuery) QueryParent() *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: pocq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pointofcontact.Table, pointofcontact.FieldID, selector),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pointofcontact.ParentTable, pointofcontact.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PointOfContact entity from the query.
// Returns a *NotFoundError when no PointOfContact was found.
func (pocq *PointOfContactQuery) First(ctx context.Context) (*PointOfContact, error) {
	nodes, err := pocq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pointofcontact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pocq *PointOfContactQuery) FirstX(ctx context.Context) *PointOfContact {
	node, err := pocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PointOfContact ID from the query.
// Returns a *NotFoundError when no PointOfContact ID was found.
func (pocq *PointOfContactQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pocq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pointofcontact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pocq *PointOfContactQuery) FirstIDX(ctx context.Context) int {
	id, err := pocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PointOfContact entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PointOfContact entity is found.
// Returns a *NotFoundError when no PointOfContact entities are found.
func (pocq *PointOfContactQuery) Only(ctx context.Context) (*PointOfContact, error) {
	nodes, err := pocq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pointofcontact.Label}
	default:
		return nil, &NotSingularError{pointofcontact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pocq *PointOfContactQuery) OnlyX(ctx context.Context) *PointOfContact {
	node, err := pocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PointOfContact ID in the query.
// Returns a *NotSingularError when more than one PointOfContact ID is found.
// Returns a *NotFoundError when no entities are found.
func (pocq *PointOfContactQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pocq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pointofcontact.Label}
	default:
		err = &NotSingularError{pointofcontact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pocq *PointOfContactQuery) OnlyIDX(ctx context.Context) int {
	id, err := pocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PointOfContacts.
func (pocq *PointOfContactQuery) All(ctx context.Context) ([]*PointOfContact, error) {
	if err := pocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pocq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pocq *PointOfContactQuery) AllX(ctx context.Context) []*PointOfContact {
	nodes, err := pocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PointOfContact IDs.
func (pocq *PointOfContactQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pocq.Select(pointofcontact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pocq *PointOfContactQuery) IDsX(ctx context.Context) []int {
	ids, err := pocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pocq *PointOfContactQuery) Count(ctx context.Context) (int, error) {
	if err := pocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pocq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pocq *PointOfContactQuery) CountX(ctx context.Context) int {
	count, err := pocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pocq *PointOfContactQuery) Exist(ctx context.Context) (bool, error) {
	if err := pocq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pocq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pocq *PointOfContactQuery) ExistX(ctx context.Context) bool {
	exist, err := pocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PointOfContactQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pocq *PointOfContactQuery) Clone() *PointOfContactQuery {
	if pocq == nil {
		return nil
	}
	return &PointOfContactQuery{
		config:     pocq.config,
		limit:      pocq.limit,
		offset:     pocq.offset,
		order:      append([]OrderFunc{}, pocq.order...),
		predicates: append([]predicate.PointOfContact{}, pocq.predicates...),
		withParent: pocq.withParent.Clone(),
		// clone intermediate query.
		sql:    pocq.sql.Clone(),
		path:   pocq.path,
		unique: pocq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (pocq *PointOfContactQuery) WithParent(opts ...func(*MoreInfoModuleQuery)) *PointOfContactQuery {
	query := &MoreInfoModuleQuery{config: pocq.config}
	for _, opt := range opts {
		opt(query)
	}
	pocq.withParent = query
	return pocq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PointOfContactTitle string `json:"point_of_contact_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PointOfContact.Query().
//		GroupBy(pointofcontact.FieldPointOfContactTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (pocq *PointOfContactQuery) GroupBy(field string, fields ...string) *PointOfContactGroupBy {
	grbuild := &PointOfContactGroupBy{config: pocq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pocq.sqlQuery(ctx), nil
	}
	grbuild.label = pointofcontact.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PointOfContactTitle string `json:"point_of_contact_title,omitempty"`
//	}
//
//	client.PointOfContact.Query().
//		Select(pointofcontact.FieldPointOfContactTitle).
//		Scan(ctx, &v)
//
func (pocq *PointOfContactQuery) Select(fields ...string) *PointOfContactSelect {
	pocq.fields = append(pocq.fields, fields...)
	selbuild := &PointOfContactSelect{PointOfContactQuery: pocq}
	selbuild.label = pointofcontact.Label
	selbuild.flds, selbuild.scan = &pocq.fields, selbuild.Scan
	return selbuild
}

func (pocq *PointOfContactQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pocq.fields {
		if !pointofcontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if pocq.path != nil {
		prev, err := pocq.path(ctx)
		if err != nil {
			return err
		}
		pocq.sql = prev
	}
	return nil
}

func (pocq *PointOfContactQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PointOfContact, error) {
	var (
		nodes       = []*PointOfContact{}
		withFKs     = pocq.withFKs
		_spec       = pocq.querySpec()
		loadedTypes = [1]bool{
			pocq.withParent != nil,
		}
	)
	if pocq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pointofcontact.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PointOfContact).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PointOfContact{config: pocq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pocq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PointOfContact)
		for i := range nodes {
			if nodes[i].point_of_contact_parent == nil {
				continue
			}
			fk := *nodes[i].point_of_contact_parent
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(moreinfomodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "point_of_contact_parent" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (pocq *PointOfContactQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pocq.querySpec()
	_spec.Node.Columns = pocq.fields
	if len(pocq.fields) > 0 {
		_spec.Unique = pocq.unique != nil && *pocq.unique
	}
	return sqlgraph.CountNodes(ctx, pocq.driver, _spec)
}

func (pocq *PointOfContactQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pocq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (pocq *PointOfContactQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pointofcontact.Table,
			Columns: pointofcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pointofcontact.FieldID,
			},
		},
		From:   pocq.sql,
		Unique: true,
	}
	if unique := pocq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pocq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pointofcontact.FieldID)
		for i := range fields {
			if fields[i] != pointofcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pocq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pocq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pocq *PointOfContactQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pocq.driver.Dialect())
	t1 := builder.Table(pointofcontact.Table)
	columns := pocq.fields
	if len(columns) == 0 {
		columns = pointofcontact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pocq.sql != nil {
		selector = pocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pocq.unique != nil && *pocq.unique {
		selector.Distinct()
	}
	for _, p := range pocq.predicates {
		p(selector)
	}
	for _, p := range pocq.order {
		p(selector)
	}
	if offset := pocq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pocq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PointOfContactGroupBy is the group-by builder for PointOfContact entities.
type PointOfContactGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pocgb *PointOfContactGroupBy) Aggregate(fns ...AggregateFunc) *PointOfContactGroupBy {
	pocgb.fns = append(pocgb.fns, fns...)
	return pocgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pocgb *PointOfContactGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pocgb.path(ctx)
	if err != nil {
		return err
	}
	pocgb.sql = query
	return pocgb.sqlScan(ctx, v)
}

func (pocgb *PointOfContactGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pocgb.fields {
		if !pointofcontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pocgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pocgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pocgb *PointOfContactGroupBy) sqlQuery() *sql.Selector {
	selector := pocgb.sql.Select()
	aggregation := make([]string, 0, len(pocgb.fns))
	for _, fn := range pocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pocgb.fields)+len(pocgb.fns))
		for _, f := range pocgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pocgb.fields...)...)
}

// PointOfContactSelect is the builder for selecting fields of PointOfContact entities.
type PointOfContactSelect struct {
	*PointOfContactQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pocs *PointOfContactSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pocs.prepareQuery(ctx); err != nil {
		return err
	}
	pocs.sql = pocs.PointOfContactQuery.sqlQuery(ctx)
	return pocs.sqlScan(ctx, v)
}

func (pocs *PointOfContactSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pocs.sql.Query()
	if err := pocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
