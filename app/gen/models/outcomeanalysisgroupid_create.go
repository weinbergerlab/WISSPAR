// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
)

// OutcomeAnalysisGroupIDCreate is the builder for creating a OutcomeAnalysisGroupID entity.
type OutcomeAnalysisGroupIDCreate struct {
	config
	mutation *OutcomeAnalysisGroupIDMutation
	hooks    []Hook
}

// SetOutcomeAnalysisGroupID sets the "outcome_analysis_group_id" field.
func (oagic *OutcomeAnalysisGroupIDCreate) SetOutcomeAnalysisGroupID(s string) *OutcomeAnalysisGroupIDCreate {
	oagic.mutation.SetOutcomeAnalysisGroupID(s)
	return oagic
}

// SetParentID sets the "parent" edge to the OutcomeAnalysis entity by ID.
func (oagic *OutcomeAnalysisGroupIDCreate) SetParentID(id int) *OutcomeAnalysisGroupIDCreate {
	oagic.mutation.SetParentID(id)
	return oagic
}

// SetNillableParentID sets the "parent" edge to the OutcomeAnalysis entity by ID if the given value is not nil.
func (oagic *OutcomeAnalysisGroupIDCreate) SetNillableParentID(id *int) *OutcomeAnalysisGroupIDCreate {
	if id != nil {
		oagic = oagic.SetParentID(*id)
	}
	return oagic
}

// SetParent sets the "parent" edge to the OutcomeAnalysis entity.
func (oagic *OutcomeAnalysisGroupIDCreate) SetParent(o *OutcomeAnalysis) *OutcomeAnalysisGroupIDCreate {
	return oagic.SetParentID(o.ID)
}

// Mutation returns the OutcomeAnalysisGroupIDMutation object of the builder.
func (oagic *OutcomeAnalysisGroupIDCreate) Mutation() *OutcomeAnalysisGroupIDMutation {
	return oagic.mutation
}

// Save creates the OutcomeAnalysisGroupID in the database.
func (oagic *OutcomeAnalysisGroupIDCreate) Save(ctx context.Context) (*OutcomeAnalysisGroupID, error) {
	var (
		err  error
		node *OutcomeAnalysisGroupID
	)
	if len(oagic.hooks) == 0 {
		if err = oagic.check(); err != nil {
			return nil, err
		}
		node, err = oagic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeAnalysisGroupIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oagic.check(); err != nil {
				return nil, err
			}
			oagic.mutation = mutation
			if node, err = oagic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oagic.hooks) - 1; i >= 0; i-- {
			if oagic.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oagic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oagic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oagic *OutcomeAnalysisGroupIDCreate) SaveX(ctx context.Context) *OutcomeAnalysisGroupID {
	v, err := oagic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oagic *OutcomeAnalysisGroupIDCreate) Exec(ctx context.Context) error {
	_, err := oagic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oagic *OutcomeAnalysisGroupIDCreate) ExecX(ctx context.Context) {
	if err := oagic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oagic *OutcomeAnalysisGroupIDCreate) check() error {
	if _, ok := oagic.mutation.OutcomeAnalysisGroupID(); !ok {
		return &ValidationError{Name: "outcome_analysis_group_id", err: errors.New(`models: missing required field "OutcomeAnalysisGroupID.outcome_analysis_group_id"`)}
	}
	return nil
}

func (oagic *OutcomeAnalysisGroupIDCreate) sqlSave(ctx context.Context) (*OutcomeAnalysisGroupID, error) {
	_node, _spec := oagic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oagic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oagic *OutcomeAnalysisGroupIDCreate) createSpec() (*OutcomeAnalysisGroupID, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeAnalysisGroupID{config: oagic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeanalysisgroupid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysisgroupid.FieldID,
			},
		}
	)
	if value, ok := oagic.mutation.OutcomeAnalysisGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID,
		})
		_node.OutcomeAnalysisGroupID = value
	}
	if nodes := oagic.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeanalysisgroupid.ParentTable,
			Columns: []string{outcomeanalysisgroupid.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomeanalysis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.outcome_analysis_outcome_analysis_group_id_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeAnalysisGroupIDCreateBulk is the builder for creating many OutcomeAnalysisGroupID entities in bulk.
type OutcomeAnalysisGroupIDCreateBulk struct {
	config
	builders []*OutcomeAnalysisGroupIDCreate
}

// Save creates the OutcomeAnalysisGroupID entities in the database.
func (oagicb *OutcomeAnalysisGroupIDCreateBulk) Save(ctx context.Context) ([]*OutcomeAnalysisGroupID, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oagicb.builders))
	nodes := make([]*OutcomeAnalysisGroupID, len(oagicb.builders))
	mutators := make([]Mutator, len(oagicb.builders))
	for i := range oagicb.builders {
		func(i int, root context.Context) {
			builder := oagicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeAnalysisGroupIDMutation)
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
					_, err = mutators[i+1].Mutate(root, oagicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oagicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oagicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oagicb *OutcomeAnalysisGroupIDCreateBulk) SaveX(ctx context.Context) []*OutcomeAnalysisGroupID {
	v, err := oagicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oagicb *OutcomeAnalysisGroupIDCreateBulk) Exec(ctx context.Context) error {
	_, err := oagicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oagicb *OutcomeAnalysisGroupIDCreateBulk) ExecX(ctx context.Context) {
	if err := oagicb.Exec(ctx); err != nil {
		panic(err)
	}
}
