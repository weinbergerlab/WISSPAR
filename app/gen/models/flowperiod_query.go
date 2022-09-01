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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowPeriodQuery is the builder for querying FlowPeriod entities.
type FlowPeriodQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowPeriod
	// eager-loading edges.
	withParent               *ParticipantFlowModuleQuery
	withFlowMilestoneList    *FlowMilestoneQuery
	withFlowDropWithdrawList *FlowDropWithdrawQuery
	withFKs                  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowPeriodQuery builder.
func (fpq *FlowPeriodQuery) Where(ps ...predicate.FlowPeriod) *FlowPeriodQuery {
	fpq.predicates = append(fpq.predicates, ps...)
	return fpq
}

// Limit adds a limit step to the query.
func (fpq *FlowPeriodQuery) Limit(limit int) *FlowPeriodQuery {
	fpq.limit = &limit
	return fpq
}

// Offset adds an offset step to the query.
func (fpq *FlowPeriodQuery) Offset(offset int) *FlowPeriodQuery {
	fpq.offset = &offset
	return fpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fpq *FlowPeriodQuery) Unique(unique bool) *FlowPeriodQuery {
	fpq.unique = &unique
	return fpq
}

// Order adds an order step to the query.
func (fpq *FlowPeriodQuery) Order(o ...OrderFunc) *FlowPeriodQuery {
	fpq.order = append(fpq.order, o...)
	return fpq
}

// QueryParent chains the current query on the "parent" edge.
func (fpq *FlowPeriodQuery) QueryParent() *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: fpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, selector),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowperiod.ParentTable, flowperiod.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowMilestoneList chains the current query on the "flow_milestone_list" edge.
func (fpq *FlowPeriodQuery) QueryFlowMilestoneList() *FlowMilestoneQuery {
	query := &FlowMilestoneQuery{config: fpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, selector),
			sqlgraph.To(flowmilestone.Table, flowmilestone.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowperiod.FlowMilestoneListTable, flowperiod.FlowMilestoneListColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowDropWithdrawList chains the current query on the "flow_drop_withdraw_list" edge.
func (fpq *FlowPeriodQuery) QueryFlowDropWithdrawList() *FlowDropWithdrawQuery {
	query := &FlowDropWithdrawQuery{config: fpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, selector),
			sqlgraph.To(flowdropwithdraw.Table, flowdropwithdraw.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowperiod.FlowDropWithdrawListTable, flowperiod.FlowDropWithdrawListColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowPeriod entity from the query.
// Returns a *NotFoundError when no FlowPeriod was found.
func (fpq *FlowPeriodQuery) First(ctx context.Context) (*FlowPeriod, error) {
	nodes, err := fpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowperiod.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fpq *FlowPeriodQuery) FirstX(ctx context.Context) *FlowPeriod {
	node, err := fpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowPeriod ID from the query.
// Returns a *NotFoundError when no FlowPeriod ID was found.
func (fpq *FlowPeriodQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowperiod.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fpq *FlowPeriodQuery) FirstIDX(ctx context.Context) int {
	id, err := fpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowPeriod entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowPeriod entity is found.
// Returns a *NotFoundError when no FlowPeriod entities are found.
func (fpq *FlowPeriodQuery) Only(ctx context.Context) (*FlowPeriod, error) {
	nodes, err := fpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowperiod.Label}
	default:
		return nil, &NotSingularError{flowperiod.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fpq *FlowPeriodQuery) OnlyX(ctx context.Context) *FlowPeriod {
	node, err := fpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowPeriod ID in the query.
// Returns a *NotSingularError when more than one FlowPeriod ID is found.
// Returns a *NotFoundError when no entities are found.
func (fpq *FlowPeriodQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowperiod.Label}
	default:
		err = &NotSingularError{flowperiod.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fpq *FlowPeriodQuery) OnlyIDX(ctx context.Context) int {
	id, err := fpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowPeriods.
func (fpq *FlowPeriodQuery) All(ctx context.Context) ([]*FlowPeriod, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fpq *FlowPeriodQuery) AllX(ctx context.Context) []*FlowPeriod {
	nodes, err := fpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowPeriod IDs.
func (fpq *FlowPeriodQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fpq.Select(flowperiod.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fpq *FlowPeriodQuery) IDsX(ctx context.Context) []int {
	ids, err := fpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fpq *FlowPeriodQuery) Count(ctx context.Context) (int, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fpq *FlowPeriodQuery) CountX(ctx context.Context) int {
	count, err := fpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fpq *FlowPeriodQuery) Exist(ctx context.Context) (bool, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fpq *FlowPeriodQuery) ExistX(ctx context.Context) bool {
	exist, err := fpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowPeriodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fpq *FlowPeriodQuery) Clone() *FlowPeriodQuery {
	if fpq == nil {
		return nil
	}
	return &FlowPeriodQuery{
		config:                   fpq.config,
		limit:                    fpq.limit,
		offset:                   fpq.offset,
		order:                    append([]OrderFunc{}, fpq.order...),
		predicates:               append([]predicate.FlowPeriod{}, fpq.predicates...),
		withParent:               fpq.withParent.Clone(),
		withFlowMilestoneList:    fpq.withFlowMilestoneList.Clone(),
		withFlowDropWithdrawList: fpq.withFlowDropWithdrawList.Clone(),
		// clone intermediate query.
		sql:    fpq.sql.Clone(),
		path:   fpq.path,
		unique: fpq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *FlowPeriodQuery) WithParent(opts ...func(*ParticipantFlowModuleQuery)) *FlowPeriodQuery {
	query := &ParticipantFlowModuleQuery{config: fpq.config}
	for _, opt := range opts {
		opt(query)
	}
	fpq.withParent = query
	return fpq
}

// WithFlowMilestoneList tells the query-builder to eager-load the nodes that are connected to
// the "flow_milestone_list" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *FlowPeriodQuery) WithFlowMilestoneList(opts ...func(*FlowMilestoneQuery)) *FlowPeriodQuery {
	query := &FlowMilestoneQuery{config: fpq.config}
	for _, opt := range opts {
		opt(query)
	}
	fpq.withFlowMilestoneList = query
	return fpq
}

// WithFlowDropWithdrawList tells the query-builder to eager-load the nodes that are connected to
// the "flow_drop_withdraw_list" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *FlowPeriodQuery) WithFlowDropWithdrawList(opts ...func(*FlowDropWithdrawQuery)) *FlowPeriodQuery {
	query := &FlowDropWithdrawQuery{config: fpq.config}
	for _, opt := range opts {
		opt(query)
	}
	fpq.withFlowDropWithdrawList = query
	return fpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowPeriodTitle string `json:"flow_period_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowPeriod.Query().
//		GroupBy(flowperiod.FieldFlowPeriodTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (fpq *FlowPeriodQuery) GroupBy(field string, fields ...string) *FlowPeriodGroupBy {
	grbuild := &FlowPeriodGroupBy{config: fpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fpq.sqlQuery(ctx), nil
	}
	grbuild.label = flowperiod.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowPeriodTitle string `json:"flow_period_title,omitempty"`
//	}
//
//	client.FlowPeriod.Query().
//		Select(flowperiod.FieldFlowPeriodTitle).
//		Scan(ctx, &v)
//
func (fpq *FlowPeriodQuery) Select(fields ...string) *FlowPeriodSelect {
	fpq.fields = append(fpq.fields, fields...)
	selbuild := &FlowPeriodSelect{FlowPeriodQuery: fpq}
	selbuild.label = flowperiod.Label
	selbuild.flds, selbuild.scan = &fpq.fields, selbuild.Scan
	return selbuild
}

func (fpq *FlowPeriodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fpq.fields {
		if !flowperiod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if fpq.path != nil {
		prev, err := fpq.path(ctx)
		if err != nil {
			return err
		}
		fpq.sql = prev
	}
	return nil
}

func (fpq *FlowPeriodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowPeriod, error) {
	var (
		nodes       = []*FlowPeriod{}
		withFKs     = fpq.withFKs
		_spec       = fpq.querySpec()
		loadedTypes = [3]bool{
			fpq.withParent != nil,
			fpq.withFlowMilestoneList != nil,
			fpq.withFlowDropWithdrawList != nil,
		}
	)
	if fpq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowperiod.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowPeriod).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowPeriod{config: fpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fpq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowPeriod)
		for i := range nodes {
			if nodes[i].participant_flow_module_flow_period_list == nil {
				continue
			}
			fk := *nodes[i].participant_flow_module_flow_period_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(participantflowmodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "participant_flow_module_flow_period_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := fpq.withFlowMilestoneList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FlowPeriod)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowMilestoneList = []*FlowMilestone{}
		}
		query.withFKs = true
		query.Where(predicate.FlowMilestone(func(s *sql.Selector) {
			s.Where(sql.InValues(flowperiod.FlowMilestoneListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flow_period_flow_milestone_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flow_period_flow_milestone_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_period_flow_milestone_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowMilestoneList = append(node.Edges.FlowMilestoneList, n)
		}
	}

	if query := fpq.withFlowDropWithdrawList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FlowPeriod)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowDropWithdrawList = []*FlowDropWithdraw{}
		}
		query.withFKs = true
		query.Where(predicate.FlowDropWithdraw(func(s *sql.Selector) {
			s.Where(sql.InValues(flowperiod.FlowDropWithdrawListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flow_period_flow_drop_withdraw_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flow_period_flow_drop_withdraw_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_period_flow_drop_withdraw_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowDropWithdrawList = append(node.Edges.FlowDropWithdrawList, n)
		}
	}

	return nodes, nil
}

func (fpq *FlowPeriodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fpq.querySpec()
	_spec.Node.Columns = fpq.fields
	if len(fpq.fields) > 0 {
		_spec.Unique = fpq.unique != nil && *fpq.unique
	}
	return sqlgraph.CountNodes(ctx, fpq.driver, _spec)
}

func (fpq *FlowPeriodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (fpq *FlowPeriodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowperiod.Table,
			Columns: flowperiod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowperiod.FieldID,
			},
		},
		From:   fpq.sql,
		Unique: true,
	}
	if unique := fpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowperiod.FieldID)
		for i := range fields {
			if fields[i] != flowperiod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fpq *FlowPeriodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fpq.driver.Dialect())
	t1 := builder.Table(flowperiod.Table)
	columns := fpq.fields
	if len(columns) == 0 {
		columns = flowperiod.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fpq.sql != nil {
		selector = fpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fpq.unique != nil && *fpq.unique {
		selector.Distinct()
	}
	for _, p := range fpq.predicates {
		p(selector)
	}
	for _, p := range fpq.order {
		p(selector)
	}
	if offset := fpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowPeriodGroupBy is the group-by builder for FlowPeriod entities.
type FlowPeriodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fpgb *FlowPeriodGroupBy) Aggregate(fns ...AggregateFunc) *FlowPeriodGroupBy {
	fpgb.fns = append(fpgb.fns, fns...)
	return fpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fpgb *FlowPeriodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fpgb.path(ctx)
	if err != nil {
		return err
	}
	fpgb.sql = query
	return fpgb.sqlScan(ctx, v)
}

func (fpgb *FlowPeriodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fpgb.fields {
		if !flowperiod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fpgb *FlowPeriodGroupBy) sqlQuery() *sql.Selector {
	selector := fpgb.sql.Select()
	aggregation := make([]string, 0, len(fpgb.fns))
	for _, fn := range fpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fpgb.fields)+len(fpgb.fns))
		for _, f := range fpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fpgb.fields...)...)
}

// FlowPeriodSelect is the builder for selecting fields of FlowPeriod entities.
type FlowPeriodSelect struct {
	*FlowPeriodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fps *FlowPeriodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fps.prepareQuery(ctx); err != nil {
		return err
	}
	fps.sql = fps.FlowPeriodQuery.sqlQuery(ctx)
	return fps.sqlScan(ctx, v)
}

func (fps *FlowPeriodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fps.sql.Query()
	if err := fps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
