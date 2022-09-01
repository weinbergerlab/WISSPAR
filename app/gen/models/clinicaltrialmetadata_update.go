// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ClinicalTrialMetadataUpdate is the builder for updating ClinicalTrialMetadata entities.
type ClinicalTrialMetadataUpdate struct {
	config
	hooks    []Hook
	mutation *ClinicalTrialMetadataMutation
}

// Where appends a list predicates to the ClinicalTrialMetadataUpdate builder.
func (ctmu *ClinicalTrialMetadataUpdate) Where(ps ...predicate.ClinicalTrialMetadata) *ClinicalTrialMetadataUpdate {
	ctmu.mutation.Where(ps...)
	return ctmu
}

// SetTagName sets the "tag_name" field.
func (ctmu *ClinicalTrialMetadataUpdate) SetTagName(s string) *ClinicalTrialMetadataUpdate {
	ctmu.mutation.SetTagName(s)
	return ctmu
}

// SetTagValue sets the "tag_value" field.
func (ctmu *ClinicalTrialMetadataUpdate) SetTagValue(s string) *ClinicalTrialMetadataUpdate {
	ctmu.mutation.SetTagValue(s)
	return ctmu
}

// SetUseCaseCode sets the "use_case_code" field.
func (ctmu *ClinicalTrialMetadataUpdate) SetUseCaseCode(s string) *ClinicalTrialMetadataUpdate {
	ctmu.mutation.SetUseCaseCode(s)
	return ctmu
}

// SetNillableUseCaseCode sets the "use_case_code" field if the given value is not nil.
func (ctmu *ClinicalTrialMetadataUpdate) SetNillableUseCaseCode(s *string) *ClinicalTrialMetadataUpdate {
	if s != nil {
		ctmu.SetUseCaseCode(*s)
	}
	return ctmu
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (ctmu *ClinicalTrialMetadataUpdate) SetParentID(id int) *ClinicalTrialMetadataUpdate {
	ctmu.mutation.SetParentID(id)
	return ctmu
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (ctmu *ClinicalTrialMetadataUpdate) SetNillableParentID(id *int) *ClinicalTrialMetadataUpdate {
	if id != nil {
		ctmu = ctmu.SetParentID(*id)
	}
	return ctmu
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (ctmu *ClinicalTrialMetadataUpdate) SetParent(c *ClinicalTrial) *ClinicalTrialMetadataUpdate {
	return ctmu.SetParentID(c.ID)
}

// Mutation returns the ClinicalTrialMetadataMutation object of the builder.
func (ctmu *ClinicalTrialMetadataUpdate) Mutation() *ClinicalTrialMetadataMutation {
	return ctmu.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (ctmu *ClinicalTrialMetadataUpdate) ClearParent() *ClinicalTrialMetadataUpdate {
	ctmu.mutation.ClearParent()
	return ctmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctmu *ClinicalTrialMetadataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctmu.hooks) == 0 {
		affected, err = ctmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctmu.mutation = mutation
			affected, err = ctmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctmu.hooks) - 1; i >= 0; i-- {
			if ctmu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctmu *ClinicalTrialMetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := ctmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctmu *ClinicalTrialMetadataUpdate) Exec(ctx context.Context) error {
	_, err := ctmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmu *ClinicalTrialMetadataUpdate) ExecX(ctx context.Context) {
	if err := ctmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctmu *ClinicalTrialMetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrialmetadata.Table,
			Columns: clinicaltrialmetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrialmetadata.FieldID,
			},
		},
	}
	if ps := ctmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctmu.mutation.TagName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagName,
		})
	}
	if value, ok := ctmu.mutation.TagValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagValue,
		})
	}
	if value, ok := ctmu.mutation.UseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldUseCaseCode,
		})
	}
	if ctmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicaltrialmetadata.ParentTable,
			Columns: []string{clinicaltrialmetadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicaltrialmetadata.ParentTable,
			Columns: []string{clinicaltrialmetadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicaltrialmetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClinicalTrialMetadataUpdateOne is the builder for updating a single ClinicalTrialMetadata entity.
type ClinicalTrialMetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClinicalTrialMetadataMutation
}

// SetTagName sets the "tag_name" field.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetTagName(s string) *ClinicalTrialMetadataUpdateOne {
	ctmuo.mutation.SetTagName(s)
	return ctmuo
}

// SetTagValue sets the "tag_value" field.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetTagValue(s string) *ClinicalTrialMetadataUpdateOne {
	ctmuo.mutation.SetTagValue(s)
	return ctmuo
}

// SetUseCaseCode sets the "use_case_code" field.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetUseCaseCode(s string) *ClinicalTrialMetadataUpdateOne {
	ctmuo.mutation.SetUseCaseCode(s)
	return ctmuo
}

// SetNillableUseCaseCode sets the "use_case_code" field if the given value is not nil.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetNillableUseCaseCode(s *string) *ClinicalTrialMetadataUpdateOne {
	if s != nil {
		ctmuo.SetUseCaseCode(*s)
	}
	return ctmuo
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetParentID(id int) *ClinicalTrialMetadataUpdateOne {
	ctmuo.mutation.SetParentID(id)
	return ctmuo
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetNillableParentID(id *int) *ClinicalTrialMetadataUpdateOne {
	if id != nil {
		ctmuo = ctmuo.SetParentID(*id)
	}
	return ctmuo
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SetParent(c *ClinicalTrial) *ClinicalTrialMetadataUpdateOne {
	return ctmuo.SetParentID(c.ID)
}

// Mutation returns the ClinicalTrialMetadataMutation object of the builder.
func (ctmuo *ClinicalTrialMetadataUpdateOne) Mutation() *ClinicalTrialMetadataMutation {
	return ctmuo.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (ctmuo *ClinicalTrialMetadataUpdateOne) ClearParent() *ClinicalTrialMetadataUpdateOne {
	ctmuo.mutation.ClearParent()
	return ctmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctmuo *ClinicalTrialMetadataUpdateOne) Select(field string, fields ...string) *ClinicalTrialMetadataUpdateOne {
	ctmuo.fields = append([]string{field}, fields...)
	return ctmuo
}

// Save executes the query and returns the updated ClinicalTrialMetadata entity.
func (ctmuo *ClinicalTrialMetadataUpdateOne) Save(ctx context.Context) (*ClinicalTrialMetadata, error) {
	var (
		err  error
		node *ClinicalTrialMetadata
	)
	if len(ctmuo.hooks) == 0 {
		node, err = ctmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctmuo.mutation = mutation
			node, err = ctmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctmuo.hooks) - 1; i >= 0; i-- {
			if ctmuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctmuo *ClinicalTrialMetadataUpdateOne) SaveX(ctx context.Context) *ClinicalTrialMetadata {
	node, err := ctmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctmuo *ClinicalTrialMetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := ctmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctmuo *ClinicalTrialMetadataUpdateOne) ExecX(ctx context.Context) {
	if err := ctmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctmuo *ClinicalTrialMetadataUpdateOne) sqlSave(ctx context.Context) (_node *ClinicalTrialMetadata, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrialmetadata.Table,
			Columns: clinicaltrialmetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrialmetadata.FieldID,
			},
		},
	}
	id, ok := ctmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "ClinicalTrialMetadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clinicaltrialmetadata.FieldID)
		for _, f := range fields {
			if !clinicaltrialmetadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != clinicaltrialmetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctmuo.mutation.TagName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagName,
		})
	}
	if value, ok := ctmuo.mutation.TagValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldTagValue,
		})
	}
	if value, ok := ctmuo.mutation.UseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrialmetadata.FieldUseCaseCode,
		})
	}
	if ctmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicaltrialmetadata.ParentTable,
			Columns: []string{clinicaltrialmetadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicaltrialmetadata.ParentTable,
			Columns: []string{clinicaltrialmetadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClinicalTrialMetadata{config: ctmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicaltrialmetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
