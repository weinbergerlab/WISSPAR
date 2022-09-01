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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ParticipantFlowModuleQuery is the builder for querying ParticipantFlowModule entities.
type ParticipantFlowModuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ParticipantFlowModule
	// eager-loading edges.
	withParent         *ResultsDefinitionQuery
	withFlowGroupList  *FlowGroupQuery
	withFlowPeriodList *FlowPeriodQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ParticipantFlowModuleQuery builder.
func (pfmq *ParticipantFlowModuleQuery) Where(ps ...predicate.ParticipantFlowModule) *ParticipantFlowModuleQuery {
	pfmq.predicates = append(pfmq.predicates, ps...)
	return pfmq
}

// Limit adds a limit step to the query.
func (pfmq *ParticipantFlowModuleQuery) Limit(limit int) *ParticipantFlowModuleQuery {
	pfmq.limit = &limit
	return pfmq
}

// Offset adds an offset step to the query.
func (pfmq *ParticipantFlowModuleQuery) Offset(offset int) *ParticipantFlowModuleQuery {
	pfmq.offset = &offset
	return pfmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pfmq *ParticipantFlowModuleQuery) Unique(unique bool) *ParticipantFlowModuleQuery {
	pfmq.unique = &unique
	return pfmq
}

// Order adds an order step to the query.
func (pfmq *ParticipantFlowModuleQuery) Order(o ...OrderFunc) *ParticipantFlowModuleQuery {
	pfmq.order = append(pfmq.order, o...)
	return pfmq
}

// QueryParent chains the current query on the "parent" edge.
func (pfmq *ParticipantFlowModuleQuery) QueryParent() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: pfmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pfmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, participantflowmodule.ParentTable, participantflowmodule.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pfmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowGroupList chains the current query on the "flow_group_list" edge.
func (pfmq *ParticipantFlowModuleQuery) QueryFlowGroupList() *FlowGroupQuery {
	query := &FlowGroupQuery{config: pfmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pfmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, selector),
			sqlgraph.To(flowgroup.Table, flowgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, participantflowmodule.FlowGroupListTable, participantflowmodule.FlowGroupListColumn),
		)
		fromU = sqlgraph.SetNeighbors(pfmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowPeriodList chains the current query on the "flow_period_list" edge.
func (pfmq *ParticipantFlowModuleQuery) QueryFlowPeriodList() *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: pfmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pfmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, selector),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, participantflowmodule.FlowPeriodListTable, participantflowmodule.FlowPeriodListColumn),
		)
		fromU = sqlgraph.SetNeighbors(pfmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ParticipantFlowModule entity from the query.
// Returns a *NotFoundError when no ParticipantFlowModule was found.
func (pfmq *ParticipantFlowModuleQuery) First(ctx context.Context) (*ParticipantFlowModule, error) {
	nodes, err := pfmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{participantflowmodule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) FirstX(ctx context.Context) *ParticipantFlowModule {
	node, err := pfmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ParticipantFlowModule ID from the query.
// Returns a *NotFoundError when no ParticipantFlowModule ID was found.
func (pfmq *ParticipantFlowModuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pfmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{participantflowmodule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) FirstIDX(ctx context.Context) int {
	id, err := pfmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ParticipantFlowModule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ParticipantFlowModule entity is found.
// Returns a *NotFoundError when no ParticipantFlowModule entities are found.
func (pfmq *ParticipantFlowModuleQuery) Only(ctx context.Context) (*ParticipantFlowModule, error) {
	nodes, err := pfmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{participantflowmodule.Label}
	default:
		return nil, &NotSingularError{participantflowmodule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) OnlyX(ctx context.Context) *ParticipantFlowModule {
	node, err := pfmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ParticipantFlowModule ID in the query.
// Returns a *NotSingularError when more than one ParticipantFlowModule ID is found.
// Returns a *NotFoundError when no entities are found.
func (pfmq *ParticipantFlowModuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pfmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{participantflowmodule.Label}
	default:
		err = &NotSingularError{participantflowmodule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := pfmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ParticipantFlowModules.
func (pfmq *ParticipantFlowModuleQuery) All(ctx context.Context) ([]*ParticipantFlowModule, error) {
	if err := pfmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pfmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) AllX(ctx context.Context) []*ParticipantFlowModule {
	nodes, err := pfmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ParticipantFlowModule IDs.
func (pfmq *ParticipantFlowModuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pfmq.Select(participantflowmodule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) IDsX(ctx context.Context) []int {
	ids, err := pfmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pfmq *ParticipantFlowModuleQuery) Count(ctx context.Context) (int, error) {
	if err := pfmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pfmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) CountX(ctx context.Context) int {
	count, err := pfmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pfmq *ParticipantFlowModuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := pfmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pfmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pfmq *ParticipantFlowModuleQuery) ExistX(ctx context.Context) bool {
	exist, err := pfmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ParticipantFlowModuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pfmq *ParticipantFlowModuleQuery) Clone() *ParticipantFlowModuleQuery {
	if pfmq == nil {
		return nil
	}
	return &ParticipantFlowModuleQuery{
		config:             pfmq.config,
		limit:              pfmq.limit,
		offset:             pfmq.offset,
		order:              append([]OrderFunc{}, pfmq.order...),
		predicates:         append([]predicate.ParticipantFlowModule{}, pfmq.predicates...),
		withParent:         pfmq.withParent.Clone(),
		withFlowGroupList:  pfmq.withFlowGroupList.Clone(),
		withFlowPeriodList: pfmq.withFlowPeriodList.Clone(),
		// clone intermediate query.
		sql:    pfmq.sql.Clone(),
		path:   pfmq.path,
		unique: pfmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (pfmq *ParticipantFlowModuleQuery) WithParent(opts ...func(*ResultsDefinitionQuery)) *ParticipantFlowModuleQuery {
	query := &ResultsDefinitionQuery{config: pfmq.config}
	for _, opt := range opts {
		opt(query)
	}
	pfmq.withParent = query
	return pfmq
}

// WithFlowGroupList tells the query-builder to eager-load the nodes that are connected to
// the "flow_group_list" edge. The optional arguments are used to configure the query builder of the edge.
func (pfmq *ParticipantFlowModuleQuery) WithFlowGroupList(opts ...func(*FlowGroupQuery)) *ParticipantFlowModuleQuery {
	query := &FlowGroupQuery{config: pfmq.config}
	for _, opt := range opts {
		opt(query)
	}
	pfmq.withFlowGroupList = query
	return pfmq
}

// WithFlowPeriodList tells the query-builder to eager-load the nodes that are connected to
// the "flow_period_list" edge. The optional arguments are used to configure the query builder of the edge.
func (pfmq *ParticipantFlowModuleQuery) WithFlowPeriodList(opts ...func(*FlowPeriodQuery)) *ParticipantFlowModuleQuery {
	query := &FlowPeriodQuery{config: pfmq.config}
	for _, opt := range opts {
		opt(query)
	}
	pfmq.withFlowPeriodList = query
	return pfmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowPreAssignmentDetails string `json:"flow_pre_assignment_details,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ParticipantFlowModule.Query().
//		GroupBy(participantflowmodule.FieldFlowPreAssignmentDetails).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (pfmq *ParticipantFlowModuleQuery) GroupBy(field string, fields ...string) *ParticipantFlowModuleGroupBy {
	grbuild := &ParticipantFlowModuleGroupBy{config: pfmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pfmq.sqlQuery(ctx), nil
	}
	grbuild.label = participantflowmodule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowPreAssignmentDetails string `json:"flow_pre_assignment_details,omitempty"`
//	}
//
//	client.ParticipantFlowModule.Query().
//		Select(participantflowmodule.FieldFlowPreAssignmentDetails).
//		Scan(ctx, &v)
//
func (pfmq *ParticipantFlowModuleQuery) Select(fields ...string) *ParticipantFlowModuleSelect {
	pfmq.fields = append(pfmq.fields, fields...)
	selbuild := &ParticipantFlowModuleSelect{ParticipantFlowModuleQuery: pfmq}
	selbuild.label = participantflowmodule.Label
	selbuild.flds, selbuild.scan = &pfmq.fields, selbuild.Scan
	return selbuild
}

func (pfmq *ParticipantFlowModuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pfmq.fields {
		if !participantflowmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if pfmq.path != nil {
		prev, err := pfmq.path(ctx)
		if err != nil {
			return err
		}
		pfmq.sql = prev
	}
	return nil
}

func (pfmq *ParticipantFlowModuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ParticipantFlowModule, error) {
	var (
		nodes       = []*ParticipantFlowModule{}
		withFKs     = pfmq.withFKs
		_spec       = pfmq.querySpec()
		loadedTypes = [3]bool{
			pfmq.withParent != nil,
			pfmq.withFlowGroupList != nil,
			pfmq.withFlowPeriodList != nil,
		}
	)
	if pfmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, participantflowmodule.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ParticipantFlowModule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ParticipantFlowModule{config: pfmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pfmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pfmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ParticipantFlowModule)
		for i := range nodes {
			if nodes[i].results_definition_participant_flow_module == nil {
				continue
			}
			fk := *nodes[i].results_definition_participant_flow_module
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
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_participant_flow_module" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := pfmq.withFlowGroupList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ParticipantFlowModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowGroupList = []*FlowGroup{}
		}
		query.withFKs = true
		query.Where(predicate.FlowGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(participantflowmodule.FlowGroupListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.participant_flow_module_flow_group_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "participant_flow_module_flow_group_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "participant_flow_module_flow_group_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowGroupList = append(node.Edges.FlowGroupList, n)
		}
	}

	if query := pfmq.withFlowPeriodList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ParticipantFlowModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowPeriodList = []*FlowPeriod{}
		}
		query.withFKs = true
		query.Where(predicate.FlowPeriod(func(s *sql.Selector) {
			s.Where(sql.InValues(participantflowmodule.FlowPeriodListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.participant_flow_module_flow_period_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "participant_flow_module_flow_period_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "participant_flow_module_flow_period_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowPeriodList = append(node.Edges.FlowPeriodList, n)
		}
	}

	return nodes, nil
}

func (pfmq *ParticipantFlowModuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pfmq.querySpec()
	_spec.Node.Columns = pfmq.fields
	if len(pfmq.fields) > 0 {
		_spec.Unique = pfmq.unique != nil && *pfmq.unique
	}
	return sqlgraph.CountNodes(ctx, pfmq.driver, _spec)
}

func (pfmq *ParticipantFlowModuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pfmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (pfmq *ParticipantFlowModuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participantflowmodule.Table,
			Columns: participantflowmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participantflowmodule.FieldID,
			},
		},
		From:   pfmq.sql,
		Unique: true,
	}
	if unique := pfmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pfmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, participantflowmodule.FieldID)
		for i := range fields {
			if fields[i] != participantflowmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pfmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pfmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pfmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pfmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pfmq *ParticipantFlowModuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pfmq.driver.Dialect())
	t1 := builder.Table(participantflowmodule.Table)
	columns := pfmq.fields
	if len(columns) == 0 {
		columns = participantflowmodule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pfmq.sql != nil {
		selector = pfmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pfmq.unique != nil && *pfmq.unique {
		selector.Distinct()
	}
	for _, p := range pfmq.predicates {
		p(selector)
	}
	for _, p := range pfmq.order {
		p(selector)
	}
	if offset := pfmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pfmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ParticipantFlowModuleGroupBy is the group-by builder for ParticipantFlowModule entities.
type ParticipantFlowModuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pfmgb *ParticipantFlowModuleGroupBy) Aggregate(fns ...AggregateFunc) *ParticipantFlowModuleGroupBy {
	pfmgb.fns = append(pfmgb.fns, fns...)
	return pfmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pfmgb *ParticipantFlowModuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pfmgb.path(ctx)
	if err != nil {
		return err
	}
	pfmgb.sql = query
	return pfmgb.sqlScan(ctx, v)
}

func (pfmgb *ParticipantFlowModuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pfmgb.fields {
		if !participantflowmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pfmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pfmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pfmgb *ParticipantFlowModuleGroupBy) sqlQuery() *sql.Selector {
	selector := pfmgb.sql.Select()
	aggregation := make([]string, 0, len(pfmgb.fns))
	for _, fn := range pfmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pfmgb.fields)+len(pfmgb.fns))
		for _, f := range pfmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pfmgb.fields...)...)
}

// ParticipantFlowModuleSelect is the builder for selecting fields of ParticipantFlowModule entities.
type ParticipantFlowModuleSelect struct {
	*ParticipantFlowModuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pfms *ParticipantFlowModuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pfms.prepareQuery(ctx); err != nil {
		return err
	}
	pfms.sql = pfms.ParticipantFlowModuleQuery.sqlQuery(ctx)
	return pfms.sqlScan(ctx, v)
}

func (pfms *ParticipantFlowModuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pfms.sql.Query()
	if err := pfms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
