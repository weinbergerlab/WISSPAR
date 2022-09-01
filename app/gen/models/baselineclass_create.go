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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
)

// BaselineClassCreate is the builder for creating a BaselineClass entity.
type BaselineClassCreate struct {
	config
	mutation *BaselineClassMutation
	hooks    []Hook
}

// SetBaselineClassTitle sets the "baseline_class_title" field.
func (bcc *BaselineClassCreate) SetBaselineClassTitle(s string) *BaselineClassCreate {
	bcc.mutation.SetBaselineClassTitle(s)
	return bcc
}

// SetParentID sets the "parent" edge to the BaselineMeasure entity by ID.
func (bcc *BaselineClassCreate) SetParentID(id int) *BaselineClassCreate {
	bcc.mutation.SetParentID(id)
	return bcc
}

// SetNillableParentID sets the "parent" edge to the BaselineMeasure entity by ID if the given value is not nil.
func (bcc *BaselineClassCreate) SetNillableParentID(id *int) *BaselineClassCreate {
	if id != nil {
		bcc = bcc.SetParentID(*id)
	}
	return bcc
}

// SetParent sets the "parent" edge to the BaselineMeasure entity.
func (bcc *BaselineClassCreate) SetParent(b *BaselineMeasure) *BaselineClassCreate {
	return bcc.SetParentID(b.ID)
}

// AddBaselineClassDenomListIDs adds the "baseline_class_denom_list" edge to the BaselineClassDenom entity by IDs.
func (bcc *BaselineClassCreate) AddBaselineClassDenomListIDs(ids ...int) *BaselineClassCreate {
	bcc.mutation.AddBaselineClassDenomListIDs(ids...)
	return bcc
}

// AddBaselineClassDenomList adds the "baseline_class_denom_list" edges to the BaselineClassDenom entity.
func (bcc *BaselineClassCreate) AddBaselineClassDenomList(b ...*BaselineClassDenom) *BaselineClassCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcc.AddBaselineClassDenomListIDs(ids...)
}

// AddBaselineCategoryListIDs adds the "baseline_category_list" edge to the BaselineCategory entity by IDs.
func (bcc *BaselineClassCreate) AddBaselineCategoryListIDs(ids ...int) *BaselineClassCreate {
	bcc.mutation.AddBaselineCategoryListIDs(ids...)
	return bcc
}

// AddBaselineCategoryList adds the "baseline_category_list" edges to the BaselineCategory entity.
func (bcc *BaselineClassCreate) AddBaselineCategoryList(b ...*BaselineCategory) *BaselineClassCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcc.AddBaselineCategoryListIDs(ids...)
}

// Mutation returns the BaselineClassMutation object of the builder.
func (bcc *BaselineClassCreate) Mutation() *BaselineClassMutation {
	return bcc.mutation
}

// Save creates the BaselineClass in the database.
func (bcc *BaselineClassCreate) Save(ctx context.Context) (*BaselineClass, error) {
	var (
		err  error
		node *BaselineClass
	)
	if len(bcc.hooks) == 0 {
		if err = bcc.check(); err != nil {
			return nil, err
		}
		node, err = bcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaselineClassMutation)
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
func (bcc *BaselineClassCreate) SaveX(ctx context.Context) *BaselineClass {
	v, err := bcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcc *BaselineClassCreate) Exec(ctx context.Context) error {
	_, err := bcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcc *BaselineClassCreate) ExecX(ctx context.Context) {
	if err := bcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcc *BaselineClassCreate) check() error {
	if _, ok := bcc.mutation.BaselineClassTitle(); !ok {
		return &ValidationError{Name: "baseline_class_title", err: errors.New(`models: missing required field "BaselineClass.baseline_class_title"`)}
	}
	return nil
}

func (bcc *BaselineClassCreate) sqlSave(ctx context.Context) (*BaselineClass, error) {
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

func (bcc *BaselineClassCreate) createSpec() (*BaselineClass, *sqlgraph.CreateSpec) {
	var (
		_node = &BaselineClass{config: bcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: baselineclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclass.FieldID,
			},
		}
	)
	if value, ok := bcc.mutation.BaselineClassTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: baselineclass.FieldBaselineClassTitle,
		})
		_node.BaselineClassTitle = value
	}
	if nodes := bcc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   baselineclass.ParentTable,
			Columns: []string{baselineclass.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.baseline_measure_baseline_class_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BaselineClassDenomListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineClassDenomListTable,
			Columns: []string{baselineclass.BaselineClassDenomListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselineclassdenom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BaselineCategoryListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   baselineclass.BaselineCategoryListTable,
			Columns: []string{baselineclass.BaselineCategoryListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecategory.FieldID,
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

// BaselineClassCreateBulk is the builder for creating many BaselineClass entities in bulk.
type BaselineClassCreateBulk struct {
	config
	builders []*BaselineClassCreate
}

// Save creates the BaselineClass entities in the database.
func (bccb *BaselineClassCreateBulk) Save(ctx context.Context) ([]*BaselineClass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bccb.builders))
	nodes := make([]*BaselineClass, len(bccb.builders))
	mutators := make([]Mutator, len(bccb.builders))
	for i := range bccb.builders {
		func(i int, root context.Context) {
			builder := bccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BaselineClassMutation)
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
func (bccb *BaselineClassCreateBulk) SaveX(ctx context.Context) []*BaselineClass {
	v, err := bccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bccb *BaselineClassCreateBulk) Exec(ctx context.Context) error {
	_, err := bccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bccb *BaselineClassCreateBulk) ExecX(ctx context.Context) {
	if err := bccb.Exec(ctx); err != nil {
		panic(err)
	}
}
