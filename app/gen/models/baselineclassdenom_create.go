// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
)

// BaselineClassDenomCreate is the builder for creating a BaselineClassDenom entity.
type BaselineClassDenomCreate struct {
	config
	mutation *BaselineClassDenomMutation
	hooks    []Hook
}

// SetBaselineClassDenomUnits sets the "baseline_class_denom_units" field.
func (bcdc *BaselineClassDenomCreate) SetBaselineClassDenomUnits(s string) *BaselineClassDenomCreate {
	bcdc.mutation.SetBaselineClassDenomUnits(s)
	return bcdc
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcdc *BaselineClassDenomCreate) SetParentID(id int) *BaselineClassDenomCreate {
	bcdc.mutation.SetParentID(id)
	return bcdc
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcdc *BaselineClassDenomCreate) SetNillableParentID(id *int) *BaselineClassDenomCreate {
	if id != nil {
		bcdc = bcdc.SetParentID(*id)
	}
	return bcdc
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcdc *BaselineClassDenomCreate) SetParent(b *BaselineClass) *BaselineClassDenomCreate {
	return bcdc.SetParentID(b.ID)
}

// AddBaselineClassDenomCountListIDs adds the "baseline_class_denom_count_list" edge to the BaselineClassDenomCount entity by IDs.
func (bcdc *BaselineClassDenomCreate) AddBaselineClassDenomCountListIDs(ids ...int) *BaselineClassDenomCreate {
	bcdc.mutation.AddBaselineClassDenomCountListIDs(ids...)
	return bcdc
}

// AddBaselineClassDenomCountList adds the "baseline_class_denom_count_list" edges to the BaselineClassDenomCount entity.
func (bcdc *BaselineClassDenomCreate) AddBaselineClassDenomCountList(b ...*BaselineClassDenomCount) *BaselineClassDenomCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcdc.AddBaselineClassDenomCountListIDs(ids...)
}

// Mutation returns the BaselineClassDenomMutation object of the builder.
func (bcdc *BaselineClassDenomCreate) Mutation() *BaselineClassDenomMutation {
	return bcdc.mutation
}

// Save creates the BaselineClassDenom in the database.
func (bcdc *BaselineClassDenomCreate) Save(ctx context.Context) (*BaselineClassDenom, error) {
	var (
		err  error
		node *BaselineClassDenom
	)
	if len(bcdc.hooks) == 0 {
		if err = bcdc.check(); err != nil {
			return nil, err
		}
		node, err = bcdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassDenomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcdc.check(); err != nil {
				return nil, err
			}
			bcdc.mutation = mutation
			if node, err = bcdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcdc.hooks) - 1; i >= 0; i-- {
			if bcdc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcdc *BaselineClassDenomCreate) SaveX(ctx context.Context) *BaselineClassDenom {
	v, err := bcdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcdc *BaselineClassDenomCreate) Exec(ctx context.Context) error {
	_, err := bcdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdc *BaselineClassDenomCreate) ExecX(ctx context.Context) {
	if err := bcdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcdc *BaselineClassDenomCreate) check() error {
	if _, ok := bcdc.mutation.BaselineClassDenomUnits(); !ok {
		return &ValidationError{Name: "baseline_class_denom_units", err: errors.New(`models: missing required field "BaselineClassDenom.baseline_class_denom_units"`)}
	}
	return nil
}

func (bcdc *BaselineClassDenomCreate) sqlSave(ctx context.Context) (*BaselineClassDenom, error) {
	_node, _spec := bcdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bcdc *BaselineClassDenomCreate) createSpec() (*BaselineClassDenom, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineClassDenom{config: bcdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselineclassdenom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenom.FieldID,
			},
		}
	)
	if value, ok := bcdc.mutation.BaselineClassDenomUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclassdenom.FieldBaselineClassDenomUnits,
		})
		_node.BaselineClassDenomUnits = value
	}
	if nodes := bcdc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclassdenom.ParentTable,
			Columns: []string{baselineclassdenom.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_class_baseline_class_denom_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcdc.mutation.BaselineClassDenomCountListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclassdenom.BaselineClassDenomCountListTable,
			Columns: []string{baselineclassdenom.BaselineClassDenomCountListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenomcount.FieldID,
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

// BaselineClassDenomCreateBulk is the builder for creating many BaselineClassDenom entities in bulk.
type BaselineClassDenomCreateBulk struct {
	config
	builders []*BaselineClassDenomCreate
}

// Save creates the BaselineClassDenom entities in the database.
func (bcdcb *BaselineClassDenomCreateBulk) Save(ctx context.Context) ([]*BaselineClassDenom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcdcb.builders))
	nodes := make([]*BaselineClassDenom, len(bcdcb.builders))
	mutators := make([]Mutator, len(bcdcb.builders))
	for i := range bcdcb.builders {
		func(i int, root context.Context) {
			builder := bcdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineClassDenomMutation)
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
					_, err = mutators[i+1].Mutate(root, bcdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcdcb *BaselineClassDenomCreateBulk) SaveX(ctx context.Context) []*BaselineClassDenom {
	v, err := bcdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcdcb *BaselineClassDenomCreateBulk) Exec(ctx context.Context) error {
	_, err := bcdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdcb *BaselineClassDenomCreateBulk) ExecX(ctx context.Context) {
	if err := bcdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
