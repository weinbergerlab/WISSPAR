// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
)

// BaselineCategoryCreate is the builder for creating a BaselineCategory entity.
type BaselineCategoryCreate struct {
	config
	mutation *BaselineCategoryMutation
	hooks    []Hook
}

// SetBaselineCategoryTitle sets the "baseline_category_title" field.
func (bcc *BaselineCategoryCreate) SetBaselineCategoryTitle(s string) *BaselineCategoryCreate {
	bcc.mutation.SetBaselineCategoryTitle(s)
	return bcc
}

// SetParentID sets the "parent" edge to the BaselineClass entity by ID.
func (bcc *BaselineCategoryCreate) SetParentID(id int) *BaselineCategoryCreate {
	bcc.mutation.SetParentID(id)
	return bcc
}

// SetNillableParentID sets the "parent" edge to the BaselineClass entity by ID if the given value is not nil.
func (bcc *BaselineCategoryCreate) SetNillableParentID(id *int) *BaselineCategoryCreate {
	if id != nil {
		bcc = bcc.SetParentID(*id)
	}
	return bcc
}

// SetParent sets the "parent" edge to the BaselineClass entity.
func (bcc *BaselineCategoryCreate) SetParent(b *BaselineClass) *BaselineCategoryCreate {
	return bcc.SetParentID(b.ID)
}

// AddBaselineMeasurementListIDs adds the "baseline_measurement_list" edge to the BaselineMeasurement entity by IDs.
func (bcc *BaselineCategoryCreate) AddBaselineMeasurementListIDs(ids ...int) *BaselineCategoryCreate {
	bcc.mutation.AddBaselineMeasurementListIDs(ids...)
	return bcc
}

// AddBaselineMeasurementList adds the "baseline_measurement_list" edges to the BaselineMeasurement entity.
func (bcc *BaselineCategoryCreate) AddBaselineMeasurementList(b ...*BaselineMeasurement) *BaselineCategoryCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcc.AddBaselineMeasurementListIDs(ids...)
}

// Mutation returns the BaselineCategoryMutation object of the builder.
func (bcc *BaselineCategoryCreate) Mutation() *BaselineCategoryMutation {
	return bcc.mutation
}

// Save creates the BaselineCategory in the database.
func (bcc *BaselineCategoryCreate) Save(ctx context.Context) (*BaselineCategory, error) {
	var (
		err  error
		node *BaselineCategory
	)
	if len(bcc.hooks) == 0 {
		if err = bcc.check(); err != nil {
			return nil, err
		}
		node, err = bcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcc.check(); err != nil {
				return nil, err
			}
			bcc.mutation = mutation
			if node, err = bcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcc.hooks) - 1; i >= 0; i-- {
			if bcc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcc *BaselineCategoryCreate) SaveX(ctx context.Context) *BaselineCategory {
	v, err := bcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcc *BaselineCategoryCreate) Exec(ctx context.Context) error {
	_, err := bcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcc *BaselineCategoryCreate) ExecX(ctx context.Context) {
	if err := bcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcc *BaselineCategoryCreate) check() error {
	if _, ok := bcc.mutation.BaselineCategoryTitle(); !ok {
		return &ValidationError{Name: "baseline_category_title", err: errors.New(`models: missing required field "BaselineCategory.baseline_category_title"`)}
	}
	return nil
}

func (bcc *BaselineCategoryCreate) sqlSave(ctx context.Context) (*BaselineCategory, error) {
	_node, _spec := bcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bcc *BaselineCategoryCreate) createSpec() (*BaselineCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineCategory{config: bcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselinecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecategory.FieldID,
			},
		}
	)
	if value, ok := bcc.mutation.BaselineCategoryTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselinecategory.FieldBaselineCategoryTitle,
		})
		_node.BaselineCategoryTitle = value
	}
	if nodes := bcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselinecategory.ParentTable,
			Columns: []string{baselinecategory.ParentColumn},
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
		_node.baseline_class_baseline_category_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BaselineMeasurementListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselinecategory.BaselineMeasurementListTable,
			Columns: []string{baselinecategory.BaselineMeasurementListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasurement.FieldID,
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

// BaselineCategoryCreateBulk is the builder for creating many BaselineCategory entities in bulk.
type BaselineCategoryCreateBulk struct {
	config
	builders []*BaselineCategoryCreate
}

// Save creates the BaselineCategory entities in the database.
func (bccb *BaselineCategoryCreateBulk) Save(ctx context.Context) ([]*BaselineCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bccb.builders))
	nodes := make([]*BaselineCategory, len(bccb.builders))
	mutators := make([]Mutator, len(bccb.builders))
	for i := range bccb.builders {
		func(i int, root context.Context) {
			builder := bccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, bccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bccb *BaselineCategoryCreateBulk) SaveX(ctx context.Context) []*BaselineCategory {
	v, err := bccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bccb *BaselineCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := bccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bccb *BaselineCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := bccb.Exec(ctx); err != nil {
		panic(err)
	}
}
