// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ParticipantFlowModuleUpdate is the builder for updating ParticipantFlowModule entities.
type ParticipantFlowModuleUpdate struct {
	config
	hooks    []Hook
	mutation *ParticipantFlowModuleMutation
}

// Where appends a list predicates to the ParticipantFlowModuleUpdate builder.
func (pfmu *ParticipantFlowModuleUpdate) Where(ps ...predicate.ParticipantFlowModule) *ParticipantFlowModuleUpdate {
	pfmu.mutation.Where(ps...)
	return pfmu
}

// SetFlowPreAssignmentDetails sets the "flow_pre_assignment_details" field.
func (pfmu *ParticipantFlowModuleUpdate) SetFlowPreAssignmentDetails(s string) *ParticipantFlowModuleUpdate {
	pfmu.mutation.SetFlowPreAssignmentDetails(s)
	return pfmu
}

// SetFlowRecruitmentDetails sets the "flow_recruitment_details" field.
func (pfmu *ParticipantFlowModuleUpdate) SetFlowRecruitmentDetails(s string) *ParticipantFlowModuleUpdate {
	pfmu.mutation.SetFlowRecruitmentDetails(s)
	return pfmu
}

// SetFlowTypeUnitsAnalyzed sets the "flow_type_units_analyzed" field.
func (pfmu *ParticipantFlowModuleUpdate) SetFlowTypeUnitsAnalyzed(s string) *ParticipantFlowModuleUpdate {
	pfmu.mutation.SetFlowTypeUnitsAnalyzed(s)
	return pfmu
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (pfmu *ParticipantFlowModuleUpdate) SetParentID(id int) *ParticipantFlowModuleUpdate {
	pfmu.mutation.SetParentID(id)
	return pfmu
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (pfmu *ParticipantFlowModuleUpdate) SetNillableParentID(id *int) *ParticipantFlowModuleUpdate {
	if id != nil {
		pfmu = pfmu.SetParentID(*id)
	}
	return pfmu
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (pfmu *ParticipantFlowModuleUpdate) SetParent(r *ResultsDefinition) *ParticipantFlowModuleUpdate {
	return pfmu.SetParentID(r.ID)
}

// AddFlowGroupListIDs adds the "flow_group_list" edge to the FlowGroup entity by IDs.
func (pfmu *ParticipantFlowModuleUpdate) AddFlowGroupListIDs(ids ...int) *ParticipantFlowModuleUpdate {
	pfmu.mutation.AddFlowGroupListIDs(ids...)
	return pfmu
}

// AddFlowGroupList adds the "flow_group_list" edges to the FlowGroup entity.
func (pfmu *ParticipantFlowModuleUpdate) AddFlowGroupList(f ...*FlowGroup) *ParticipantFlowModuleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmu.AddFlowGroupListIDs(ids...)
}

// AddFlowPeriodListIDs adds the "flow_period_list" edge to the FlowPeriod entity by IDs.
func (pfmu *ParticipantFlowModuleUpdate) AddFlowPeriodListIDs(ids ...int) *ParticipantFlowModuleUpdate {
	pfmu.mutation.AddFlowPeriodListIDs(ids...)
	return pfmu
}

// AddFlowPeriodList adds the "flow_period_list" edges to the FlowPeriod entity.
func (pfmu *ParticipantFlowModuleUpdate) AddFlowPeriodList(f ...*FlowPeriod) *ParticipantFlowModuleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmu.AddFlowPeriodListIDs(ids...)
}

// Mutation returns the ParticipantFlowModuleMutation object of the builder.
func (pfmu *ParticipantFlowModuleUpdate) Mutation() *ParticipantFlowModuleMutation {
	return pfmu.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (pfmu *ParticipantFlowModuleUpdate) ClearParent() *ParticipantFlowModuleUpdate {
	pfmu.mutation.ClearParent()
	return pfmu
}

// ClearFlowGroupList clears all "flow_group_list" edges to the FlowGroup entity.
func (pfmu *ParticipantFlowModuleUpdate) ClearFlowGroupList() *ParticipantFlowModuleUpdate {
	pfmu.mutation.ClearFlowGroupList()
	return pfmu
}

// RemoveFlowGroupListIDs removes the "flow_group_list" edge to FlowGroup entities by IDs.
func (pfmu *ParticipantFlowModuleUpdate) RemoveFlowGroupListIDs(ids ...int) *ParticipantFlowModuleUpdate {
	pfmu.mutation.RemoveFlowGroupListIDs(ids...)
	return pfmu
}

// RemoveFlowGroupList removes "flow_group_list" edges to FlowGroup entities.
func (pfmu *ParticipantFlowModuleUpdate) RemoveFlowGroupList(f ...*FlowGroup) *ParticipantFlowModuleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmu.RemoveFlowGroupListIDs(ids...)
}

// ClearFlowPeriodList clears all "flow_period_list" edges to the FlowPeriod entity.
func (pfmu *ParticipantFlowModuleUpdate) ClearFlowPeriodList() *ParticipantFlowModuleUpdate {
	pfmu.mutation.ClearFlowPeriodList()
	return pfmu
}

// RemoveFlowPeriodListIDs removes the "flow_period_list" edge to FlowPeriod entities by IDs.
func (pfmu *ParticipantFlowModuleUpdate) RemoveFlowPeriodListIDs(ids ...int) *ParticipantFlowModuleUpdate {
	pfmu.mutation.RemoveFlowPeriodListIDs(ids...)
	return pfmu
}

// RemoveFlowPeriodList removes "flow_period_list" edges to FlowPeriod entities.
func (pfmu *ParticipantFlowModuleUpdate) RemoveFlowPeriodList(f ...*FlowPeriod) *ParticipantFlowModuleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmu.RemoveFlowPeriodListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pfmu *ParticipantFlowModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pfmu.hooks) == 0 {
		affected, err = pfmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantFlowModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pfmu.mutation = mutation
			affected, err = pfmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pfmu.hooks) - 1; i >= 0; i-- {
			if pfmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pfmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pfmu *ParticipantFlowModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := pfmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pfmu *ParticipantFlowModuleUpdate) Exec(ctx context.Context) error {
	_, err := pfmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmu *ParticipantFlowModuleUpdate) ExecX(ctx context.Context) {
	if err := pfmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pfmu *ParticipantFlowModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participantflowmodule.Table,
			Columns: participantflowmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participantflowmodule.FieldID,
			},
		},
	}
	if ps := pfmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pfmu.mutation.FlowPreAssignmentDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowPreAssignmentDetails,
		})
	}
	if value, ok := pfmu.mutation.FlowRecruitmentDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowRecruitmentDetails,
		})
	}
	if value, ok := pfmu.mutation.FlowTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowTypeUnitsAnalyzed,
		})
	}
	if pfmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   participantflowmodule.ParentTable,
			Columns: []string{participantflowmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   participantflowmodule.ParentTable,
			Columns: []string{participantflowmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pfmu.mutation.FlowGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmu.mutation.RemovedFlowGroupListIDs(); len(nodes) > 0 && !pfmu.mutation.FlowGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmu.mutation.FlowGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pfmu.mutation.FlowPeriodListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmu.mutation.RemovedFlowPeriodListIDs(); len(nodes) > 0 && !pfmu.mutation.FlowPeriodListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmu.mutation.FlowPeriodListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pfmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participantflowmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ParticipantFlowModuleUpdateOne is the builder for updating a single ParticipantFlowModule entity.
type ParticipantFlowModuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ParticipantFlowModuleMutation
}

// SetFlowPreAssignmentDetails sets the "flow_pre_assignment_details" field.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetFlowPreAssignmentDetails(s string) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.SetFlowPreAssignmentDetails(s)
	return pfmuo
}

// SetFlowRecruitmentDetails sets the "flow_recruitment_details" field.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetFlowRecruitmentDetails(s string) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.SetFlowRecruitmentDetails(s)
	return pfmuo
}

// SetFlowTypeUnitsAnalyzed sets the "flow_type_units_analyzed" field.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetFlowTypeUnitsAnalyzed(s string) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.SetFlowTypeUnitsAnalyzed(s)
	return pfmuo
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetParentID(id int) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.SetParentID(id)
	return pfmuo
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetNillableParentID(id *int) *ParticipantFlowModuleUpdateOne {
	if id != nil {
		pfmuo = pfmuo.SetParentID(*id)
	}
	return pfmuo
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) SetParent(r *ResultsDefinition) *ParticipantFlowModuleUpdateOne {
	return pfmuo.SetParentID(r.ID)
}

// AddFlowGroupListIDs adds the "flow_group_list" edge to the FlowGroup entity by IDs.
func (pfmuo *ParticipantFlowModuleUpdateOne) AddFlowGroupListIDs(ids ...int) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.AddFlowGroupListIDs(ids...)
	return pfmuo
}

// AddFlowGroupList adds the "flow_group_list" edges to the FlowGroup entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) AddFlowGroupList(f ...*FlowGroup) *ParticipantFlowModuleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmuo.AddFlowGroupListIDs(ids...)
}

// AddFlowPeriodListIDs adds the "flow_period_list" edge to the FlowPeriod entity by IDs.
func (pfmuo *ParticipantFlowModuleUpdateOne) AddFlowPeriodListIDs(ids ...int) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.AddFlowPeriodListIDs(ids...)
	return pfmuo
}

// AddFlowPeriodList adds the "flow_period_list" edges to the FlowPeriod entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) AddFlowPeriodList(f ...*FlowPeriod) *ParticipantFlowModuleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmuo.AddFlowPeriodListIDs(ids...)
}

// Mutation returns the ParticipantFlowModuleMutation object of the builder.
func (pfmuo *ParticipantFlowModuleUpdateOne) Mutation() *ParticipantFlowModuleMutation {
	return pfmuo.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) ClearParent() *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.ClearParent()
	return pfmuo
}

// ClearFlowGroupList clears all "flow_group_list" edges to the FlowGroup entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) ClearFlowGroupList() *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.ClearFlowGroupList()
	return pfmuo
}

// RemoveFlowGroupListIDs removes the "flow_group_list" edge to FlowGroup entities by IDs.
func (pfmuo *ParticipantFlowModuleUpdateOne) RemoveFlowGroupListIDs(ids ...int) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.RemoveFlowGroupListIDs(ids...)
	return pfmuo
}

// RemoveFlowGroupList removes "flow_group_list" edges to FlowGroup entities.
func (pfmuo *ParticipantFlowModuleUpdateOne) RemoveFlowGroupList(f ...*FlowGroup) *ParticipantFlowModuleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmuo.RemoveFlowGroupListIDs(ids...)
}

// ClearFlowPeriodList clears all "flow_period_list" edges to the FlowPeriod entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) ClearFlowPeriodList() *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.ClearFlowPeriodList()
	return pfmuo
}

// RemoveFlowPeriodListIDs removes the "flow_period_list" edge to FlowPeriod entities by IDs.
func (pfmuo *ParticipantFlowModuleUpdateOne) RemoveFlowPeriodListIDs(ids ...int) *ParticipantFlowModuleUpdateOne {
	pfmuo.mutation.RemoveFlowPeriodListIDs(ids...)
	return pfmuo
}

// RemoveFlowPeriodList removes "flow_period_list" edges to FlowPeriod entities.
func (pfmuo *ParticipantFlowModuleUpdateOne) RemoveFlowPeriodList(f ...*FlowPeriod) *ParticipantFlowModuleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmuo.RemoveFlowPeriodListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pfmuo *ParticipantFlowModuleUpdateOne) Select(field string, fields ...string) *ParticipantFlowModuleUpdateOne {
	pfmuo.fields = append([]string{field}, fields...)
	return pfmuo
}

// Save executes the query and returns the updated ParticipantFlowModule entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) Save(ctx context.Context) (*ParticipantFlowModule, error) {
	var (
		err  error
		node *ParticipantFlowModule
	)
	if len(pfmuo.hooks) == 0 {
		node, err = pfmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantFlowModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pfmuo.mutation = mutation
			node, err = pfmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pfmuo.hooks) - 1; i >= 0; i-- {
			if pfmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pfmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pfmuo *ParticipantFlowModuleUpdateOne) SaveX(ctx context.Context) *ParticipantFlowModule {
	node, err := pfmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pfmuo *ParticipantFlowModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := pfmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmuo *ParticipantFlowModuleUpdateOne) ExecX(ctx context.Context) {
	if err := pfmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pfmuo *ParticipantFlowModuleUpdateOne) sqlSave(ctx context.Context) (_node *ParticipantFlowModule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participantflowmodule.Table,
			Columns: participantflowmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participantflowmodule.FieldID,
			},
		},
	}
	id, ok := pfmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "ParticipantFlowModule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pfmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, participantflowmodule.FieldID)
		for _, f := range fields {
			if !participantflowmodule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != participantflowmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pfmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pfmuo.mutation.FlowPreAssignmentDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowPreAssignmentDetails,
		})
	}
	if value, ok := pfmuo.mutation.FlowRecruitmentDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowRecruitmentDetails,
		})
	}
	if value, ok := pfmuo.mutation.FlowTypeUnitsAnalyzed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowTypeUnitsAnalyzed,
		})
	}
	if pfmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   participantflowmodule.ParentTable,
			Columns: []string{participantflowmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   participantflowmodule.ParentTable,
			Columns: []string{participantflowmodule.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pfmuo.mutation.FlowGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmuo.mutation.RemovedFlowGroupListIDs(); len(nodes) > 0 && !pfmuo.mutation.FlowGroupListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmuo.mutation.FlowGroupListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowGroupListTable,
			Columns: []string{participantflowmodule.FlowGroupListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pfmuo.mutation.FlowPeriodListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmuo.mutation.RemovedFlowPeriodListIDs(); len(nodes) > 0 && !pfmuo.mutation.FlowPeriodListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pfmuo.mutation.FlowPeriodListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participantflowmodule.FlowPeriodListTable,
			Columns: []string{participantflowmodule.FlowPeriodListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flowperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ParticipantFlowModule{config: pfmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pfmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participantflowmodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
