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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// CertainAgreementQuery is the builder for querying CertainAgreement entities.
type CertainAgreementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CertainAgreement
	// eager-loading edges.
	withParent *MoreInfoModuleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertainAgreementQuery builder.
func (caq *CertainAgreementQuery) Where(ps ...predicate.CertainAgreement) *CertainAgreementQuery {
	caq.predicates = append(caq.predicates, ps...)
	return caq
}

// Limit adds a limit step to the query.
func (caq *CertainAgreementQuery) Limit(limit int) *CertainAgreementQuery {
	caq.limit = &limit
	return caq
}

// Offset adds an offset step to the query.
func (caq *CertainAgreementQuery) Offset(offset int) *CertainAgreementQuery {
	caq.offset = &offset
	return caq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (caq *CertainAgreementQuery) Unique(unique bool) *CertainAgreementQuery {
	caq.unique = &unique
	return caq
}

// Order adds an order step to the query.
func (caq *CertainAgreementQuery) Order(o ...OrderFunc) *CertainAgreementQuery {
	caq.order = append(caq.order, o...)
	return caq
}

// QueryParent chains the current query on the "parent" edge.
func (caq *CertainAgreementQuery) QueryParent() *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: caq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := caq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certainagreement.Table, certainagreement.FieldID, selector),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certainagreement.ParentTable, certainagreement.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(caq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CertainAgreement entity from the query.
// Returns a *NotFoundError when no CertainAgreement was found.
func (caq *CertainAgreementQuery) First(ctx context.Context) (*CertainAgreement, error) {
	nodes, err := caq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certainagreement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (caq *CertainAgreementQuery) FirstX(ctx context.Context) *CertainAgreement {
	node, err := caq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CertainAgreement ID from the query.
// Returns a *NotFoundError when no CertainAgreement ID was found.
func (caq *CertainAgreementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = caq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certainagreement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (caq *CertainAgreementQuery) FirstIDX(ctx context.Context) int {
	id, err := caq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CertainAgreement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CertainAgreement entity is found.
// Returns a *NotFoundError when no CertainAgreement entities are found.
func (caq *CertainAgreementQuery) Only(ctx context.Context) (*CertainAgreement, error) {
	nodes, err := caq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certainagreement.Label}
	default:
		return nil, &NotSingularError{certainagreement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (caq *CertainAgreementQuery) OnlyX(ctx context.Context) *CertainAgreement {
	node, err := caq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CertainAgreement ID in the query.
// Returns a *NotSingularError when more than one CertainAgreement ID is found.
// Returns a *NotFoundError when no entities are found.
func (caq *CertainAgreementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = caq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certainagreement.Label}
	default:
		err = &NotSingularError{certainagreement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (caq *CertainAgreementQuery) OnlyIDX(ctx context.Context) int {
	id, err := caq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CertainAgreements.
func (caq *CertainAgreementQuery) All(ctx context.Context) ([]*CertainAgreement, error) {
	if err := caq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return caq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (caq *CertainAgreementQuery) AllX(ctx context.Context) []*CertainAgreement {
	nodes, err := caq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CertainAgreement IDs.
func (caq *CertainAgreementQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := caq.Select(certainagreement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (caq *CertainAgreementQuery) IDsX(ctx context.Context) []int {
	ids, err := caq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (caq *CertainAgreementQuery) Count(ctx context.Context) (int, error) {
	if err := caq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return caq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (caq *CertainAgreementQuery) CountX(ctx context.Context) int {
	count, err := caq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (caq *CertainAgreementQuery) Exist(ctx context.Context) (bool, error) {
	if err := caq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return caq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (caq *CertainAgreementQuery) ExistX(ctx context.Context) bool {
	exist, err := caq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertainAgreementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (caq *CertainAgreementQuery) Clone() *CertainAgreementQuery {
	if caq == nil {
		return nil
	}
	return &CertainAgreementQuery{
		config:     caq.config,
		limit:      caq.limit,
		offset:     caq.offset,
		order:      append([]OrderFunc{}, caq.order...),
		predicates: append([]predicate.CertainAgreement{}, caq.predicates...),
		withParent: caq.withParent.Clone(),
		// clone intermediate query.
		sql:    caq.sql.Clone(),
		path:   caq.path,
		unique: caq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (caq *CertainAgreementQuery) WithParent(opts ...func(*MoreInfoModuleQuery)) *CertainAgreementQuery {
	query := &MoreInfoModuleQuery{config: caq.config}
	for _, opt := range opts {
		opt(query)
	}
	caq.withParent = query
	return caq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AgreementPiSponsorEmployee string `json:"agreement_pi_sponsor_employee,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CertainAgreement.Query().
//		GroupBy(certainagreement.FieldAgreementPiSponsorEmployee).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (caq *CertainAgreementQuery) GroupBy(field string, fields ...string) *CertainAgreementGroupBy {
	grbuild := &CertainAgreementGroupBy{config: caq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return caq.sqlQuery(ctx), nil
	}
	grbuild.label = certainagreement.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AgreementPiSponsorEmployee string `json:"agreement_pi_sponsor_employee,omitempty"`
//	}
//
//	client.CertainAgreement.Query().
//		Select(certainagreement.FieldAgreementPiSponsorEmployee).
//		Scan(ctx, &v)
//
func (caq *CertainAgreementQuery) Select(fields ...string) *CertainAgreementSelect {
	caq.fields = append(caq.fields, fields...)
	selbuild := &CertainAgreementSelect{CertainAgreementQuery: caq}
	selbuild.label = certainagreement.Label
	selbuild.flds, selbuild.scan = &caq.fields, selbuild.Scan
	return selbuild
}

func (caq *CertainAgreementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range caq.fields {
		if !certainagreement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if caq.path != nil {
		prev, err := caq.path(ctx)
		if err != nil {
			return err
		}
		caq.sql = prev
	}
	return nil
}

func (caq *CertainAgreementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CertainAgreement, error) {
	var (
		nodes       = []*CertainAgreement{}
		withFKs     = caq.withFKs
		_spec       = caq.querySpec()
		loadedTypes = [1]bool{
			caq.withParent != nil,
		}
	)
	if caq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, certainagreement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CertainAgreement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CertainAgreement{config: caq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, caq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := caq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CertainAgreement)
		for i := range nodes {
			if nodes[i].certain_agreement_parent == nil {
				continue
			}
			fk := *nodes[i].certain_agreement_parent
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
				return nil, fmt.Errorf(`unexpected foreign-key "certain_agreement_parent" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (caq *CertainAgreementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := caq.querySpec()
	_spec.Node.Columns = caq.fields
	if len(caq.fields) > 0 {
		_spec.Unique = caq.unique != nil && *caq.unique
	}
	return sqlgraph.CountNodes(ctx, caq.driver, _spec)
}

func (caq *CertainAgreementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := caq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (caq *CertainAgreementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   certainagreement.Table,
			Columns: certainagreement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certainagreement.FieldID,
			},
		},
		From:   caq.sql,
		Unique: true,
	}
	if unique := caq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := caq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certainagreement.FieldID)
		for i := range fields {
			if fields[i] != certainagreement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := caq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := caq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := caq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := caq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (caq *CertainAgreementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(caq.driver.Dialect())
	t1 := builder.Table(certainagreement.Table)
	columns := caq.fields
	if len(columns) == 0 {
		columns = certainagreement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if caq.sql != nil {
		selector = caq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if caq.unique != nil && *caq.unique {
		selector.Distinct()
	}
	for _, p := range caq.predicates {
		p(selector)
	}
	for _, p := range caq.order {
		p(selector)
	}
	if offset := caq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := caq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CertainAgreementGroupBy is the group-by builder for CertainAgreement entities.
type CertainAgreementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cagb *CertainAgreementGroupBy) Aggregate(fns ...AggregateFunc) *CertainAgreementGroupBy {
	cagb.fns = append(cagb.fns, fns...)
	return cagb
}

// Scan applies the group-by query and scans the result into the given value.
func (cagb *CertainAgreementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cagb.path(ctx)
	if err != nil {
		return err
	}
	cagb.sql = query
	return cagb.sqlScan(ctx, v)
}

func (cagb *CertainAgreementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cagb.fields {
		if !certainagreement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cagb *CertainAgreementGroupBy) sqlQuery() *sql.Selector {
	selector := cagb.sql.Select()
	aggregation := make([]string, 0, len(cagb.fns))
	for _, fn := range cagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cagb.fields)+len(cagb.fns))
		for _, f := range cagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cagb.fields...)...)
}

// CertainAgreementSelect is the builder for selecting fields of CertainAgreement entities.
type CertainAgreementSelect struct {
	*CertainAgreementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cas *CertainAgreementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cas.prepareQuery(ctx); err != nil {
		return err
	}
	cas.sql = cas.CertainAgreementQuery.sqlQuery(ctx)
	return cas.sqlScan(ctx, v)
}

func (cas *CertainAgreementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cas.sql.Query()
	if err := cas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
