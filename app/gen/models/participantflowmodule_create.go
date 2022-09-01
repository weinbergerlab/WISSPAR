// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ParticipantFlowModuleCreate is the builder for creating a ParticipantFlowModule entity.
type ParticipantFlowModuleCreate struct {
	config
	mutation *ParticipantFlowModuleMutation
	hooks    []Hook
}

// SetFlowPreAssignmentDetails sets the "flow_pre_assignment_details" field.
func (pfmc *ParticipantFlowModuleCreate) SetFlowPreAssignmentDetails(s string) *ParticipantFlowModuleCreate {
	pfmc.mutation.SetFlowPreAssignmentDetails(s)
	return pfmc
}

// SetFlowRecruitmentDetails sets the "flow_recruitment_details" field.
func (pfmc *ParticipantFlowModuleCreate) SetFlowRecruitmentDetails(s string) *ParticipantFlowModuleCreate {
	pfmc.mutation.SetFlowRecruitmentDetails(s)
	return pfmc
}

// SetFlowTypeUnitsAnalyzed sets the "flow_type_units_analyzed" field.
func (pfmc *ParticipantFlowModuleCreate) SetFlowTypeUnitsAnalyzed(s string) *ParticipantFlowModuleCreate {
	pfmc.mutation.SetFlowTypeUnitsAnalyzed(s)
	return pfmc
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (pfmc *ParticipantFlowModuleCreate) SetParentID(id int) *ParticipantFlowModuleCreate {
	pfmc.mutation.SetParentID(id)
	return pfmc
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (pfmc *ParticipantFlowModuleCreate) SetNillableParentID(id *int) *ParticipantFlowModuleCreate {
	if id != nil {
		pfmc = pfmc.SetParentID(*id)
	}
	return pfmc
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (pfmc *ParticipantFlowModuleCreate) SetParent(r *ResultsDefinition) *ParticipantFlowModuleCreate {
	return pfmc.SetParentID(r.ID)
}

// AddFlowGroupListIDs adds the "flow_group_list" edge to the FlowGroup entity by IDs.
func (pfmc *ParticipantFlowModuleCreate) AddFlowGroupListIDs(ids ...int) *ParticipantFlowModuleCreate {
	pfmc.mutation.AddFlowGroupListIDs(ids...)
	return pfmc
}

// AddFlowGroupList adds the "flow_group_list" edges to the FlowGroup entity.
func (pfmc *ParticipantFlowModuleCreate) AddFlowGroupList(f ...*FlowGroup) *ParticipantFlowModuleCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmc.AddFlowGroupListIDs(ids...)
}

// AddFlowPeriodListIDs adds the "flow_period_list" edge to the FlowPeriod entity by IDs.
func (pfmc *ParticipantFlowModuleCreate) AddFlowPeriodListIDs(ids ...int) *ParticipantFlowModuleCreate {
	pfmc.mutation.AddFlowPeriodListIDs(ids...)
	return pfmc
}

// AddFlowPeriodList adds the "flow_period_list" edges to the FlowPeriod entity.
func (pfmc *ParticipantFlowModuleCreate) AddFlowPeriodList(f ...*FlowPeriod) *ParticipantFlowModuleCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pfmc.AddFlowPeriodListIDs(ids...)
}

// Mutation returns the ParticipantFlowModuleMutation object of the builder.
func (pfmc *ParticipantFlowModuleCreate) Mutation() *ParticipantFlowModuleMutation {
	return pfmc.mutation
}

// Save creates the ParticipantFlowModule in the database.
func (pfmc *ParticipantFlowModuleCreate) Save(ctx context.Context) (*ParticipantFlowModule, error) {
	var (
		err  error
		node *ParticipantFlowModule
	)
	if len(pfmc.hooks) == 0 {
		if err = pfmc.check(); err != nil {
			return nil, err
		}
		node, err = pfmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantFlowModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pfmc.check(); err != nil {
				return nil, err
			}
			pfmc.mutation = mutation
			if node, err = pfmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pfmc.hooks) - 1; i >= 0; i-- {
			if pfmc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pfmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pfmc *ParticipantFlowModuleCreate) SaveX(ctx context.Context) *ParticipantFlowModule {
	v, err := pfmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfmc *ParticipantFlowModuleCreate) Exec(ctx context.Context) error {
	_, err := pfmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmc *ParticipantFlowModuleCreate) ExecX(ctx context.Context) {
	if err := pfmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfmc *ParticipantFlowModuleCreate) check() error {
	if _, ok := pfmc.mutation.FlowPreAssignmentDetails(); !ok {
		return &ValidationError{Name: "flow_pre_assignment_details", err: errors.New(`models: missing required field "ParticipantFlowModule.flow_pre_assignment_details"`)}
	}
	if _, ok := pfmc.mutation.FlowRecruitmentDetails(); !ok {
		return &ValidationError{Name: "flow_recruitment_details", err: errors.New(`models: missing required field "ParticipantFlowModule.flow_recruitment_details"`)}
	}
	if _, ok := pfmc.mutation.FlowTypeUnitsAnalyzed(); !ok {
		return &ValidationError{Name: "flow_type_units_analyzed", err: errors.New(`models: missing required field "ParticipantFlowModule.flow_type_units_analyzed"`)}
	}
	return nil
}

func (pfmc *ParticipantFlowModuleCreate) sqlSave(ctx context.Context) (*ParticipantFlowModule, error) {
	_node, _spec := pfmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pfmc *ParticipantFlowModuleCreate) createSpec() (*ParticipantFlowModule, *sqlgraph.CreateSpec) {
	var (
		_node = &ParticipantFlowModule{config: pfmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: participantflowmodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participantflowmodule.FieldID,
			},
		}
	)
	if value, ok := pfmc.mutation.FlowPreAssignmentDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowPreAssignmentDetails,
		})
		_node.FlowPreAssignmentDetails = value
	}
	if value, ok := pfmc.mutation.FlowRecruitmentDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowRecruitmentDetails,
		})
		_node.FlowRecruitmentDetails = value
	}
	if value, ok := pfmc.mutation.FlowTypeUnitsAnalyzed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participantflowmodule.FieldFlowTypeUnitsAnalyzed,
		})
		_node.FlowTypeUnitsAnalyzed = value
	}
	if nodes := pfmc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.results_definition_participant_flow_module = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfmc.mutation.FlowGroupListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfmc.mutation.FlowPeriodListIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ParticipantFlowModuleCreateBulk is the builder for creating many ParticipantFlowModule entities in bulk.
type ParticipantFlowModuleCreateBulk struct {
	config
	builders []*ParticipantFlowModuleCreate
}

// Save creates the ParticipantFlowModule entities in the database.
func (pfmcb *ParticipantFlowModuleCreateBulk) Save(ctx context.Context) ([]*ParticipantFlowModule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pfmcb.builders))
	nodes := make([]*ParticipantFlowModule, len(pfmcb.builders))
	mutators := make([]Mutator, len(pfmcb.builders))
	for i := range pfmcb.builders {
		func(i int, root context.Context) {
			builder := pfmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParticipantFlowModuleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pfmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pfmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfmcb *ParticipantFlowModuleCreateBulk) SaveX(ctx context.Context) []*ParticipantFlowModule {
	v, err := pfmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfmcb *ParticipantFlowModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := pfmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfmcb *ParticipantFlowModuleCreateBulk) ExecX(ctx context.Context) {
	if err := pfmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
