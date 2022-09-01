// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// CertainAgreementUpdate is the builder for updating CertainAgreement entities.
type CertainAgreementUpdate struct {
	config
	hooks    []Hook
	mutation *CertainAgreementMutation
}

// Where appends a list predicates to the CertainAgreementUpdate builder.
func (cau *CertainAgreementUpdate) Where(ps ...predicate.CertainAgreement) *CertainAgreementUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetAgreementPiSponsorEmployee sets the "agreement_pi_sponsor_employee" field.
func (cau *CertainAgreementUpdate) SetAgreementPiSponsorEmployee(s string) *CertainAgreementUpdate {
	cau.mutation.SetAgreementPiSponsorEmployee(s)
	return cau
}

// SetAgreementRestrictionType sets the "agreement_restriction_type" field.
func (cau *CertainAgreementUpdate) SetAgreementRestrictionType(s string) *CertainAgreementUpdate {
	cau.mutation.SetAgreementRestrictionType(s)
	return cau
}

// SetAgreementRestrictiveAgreement sets the "agreement_restrictive_agreement" field.
func (cau *CertainAgreementUpdate) SetAgreementRestrictiveAgreement(s string) *CertainAgreementUpdate {
	cau.mutation.SetAgreementRestrictiveAgreement(s)
	return cau
}

// SetAgreementOtherDetails sets the "agreement_other_details" field.
func (cau *CertainAgreementUpdate) SetAgreementOtherDetails(s string) *CertainAgreementUpdate {
	cau.mutation.SetAgreementOtherDetails(s)
	return cau
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (cau *CertainAgreementUpdate) SetParentID(id int) *CertainAgreementUpdate {
	cau.mutation.SetParentID(id)
	return cau
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (cau *CertainAgreementUpdate) SetParent(m *MoreInfoModule) *CertainAgreementUpdate {
	return cau.SetParentID(m.ID)
}

// Mutation returns the CertainAgreementMutation object of the builder.
func (cau *CertainAgreementUpdate) Mutation() *CertainAgreementMutation {
	return cau.mutation
}

// ClearParent clears the "parent" edge to the MoreInfoModule entity.
func (cau *CertainAgreementUpdate) ClearParent() *CertainAgreementUpdate {
	cau.mutation.ClearParent()
	return cau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CertainAgreementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cau.hooks) == 0 {
		if err = cau.check(); err != nil {
			return 0, err
		}
		affected, err = cau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertainAgreementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cau.check(); err != nil {
				return 0, err
			}
			cau.mutation = mutation
			affected, err = cau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cau.hooks) - 1; i >= 0; i-- {
			if cau.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = cau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cau *CertainAgreementUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CertainAgreementUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CertainAgreementUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cau *CertainAgreementUpdate) check() error {
	if _, ok := cau.mutation.ParentID(); cau.mutation.ParentCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "CertainAgreement.parent"`)
	}
	return nil
}

func (cau *CertainAgreementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   certainagreement.Table,
			Columns: certainagreement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certainagreement.FieldID,
			},
		},
	}
	if ps := cau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cau.mutation.AgreementPiSponsorEmployee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementPiSponsorEmployee,
		})
	}
	if value, ok := cau.mutation.AgreementRestrictionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictionType,
		})
	}
	if value, ok := cau.mutation.AgreementRestrictiveAgreement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictiveAgreement,
		})
	}
	if value, ok := cau.mutation.AgreementOtherDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementOtherDetails,
		})
	}
	if cau.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certainagreement.ParentTable,
			Columns: []string{certainagreement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cau.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certainagreement.ParentTable,
			Columns: []string{certainagreement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certainagreement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CertainAgreementUpdateOne is the builder for updating a single CertainAgreement entity.
type CertainAgreementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CertainAgreementMutation
}

// SetAgreementPiSponsorEmployee sets the "agreement_pi_sponsor_employee" field.
func (cauo *CertainAgreementUpdateOne) SetAgreementPiSponsorEmployee(s string) *CertainAgreementUpdateOne {
	cauo.mutation.SetAgreementPiSponsorEmployee(s)
	return cauo
}

// SetAgreementRestrictionType sets the "agreement_restriction_type" field.
func (cauo *CertainAgreementUpdateOne) SetAgreementRestrictionType(s string) *CertainAgreementUpdateOne {
	cauo.mutation.SetAgreementRestrictionType(s)
	return cauo
}

// SetAgreementRestrictiveAgreement sets the "agreement_restrictive_agreement" field.
func (cauo *CertainAgreementUpdateOne) SetAgreementRestrictiveAgreement(s string) *CertainAgreementUpdateOne {
	cauo.mutation.SetAgreementRestrictiveAgreement(s)
	return cauo
}

// SetAgreementOtherDetails sets the "agreement_other_details" field.
func (cauo *CertainAgreementUpdateOne) SetAgreementOtherDetails(s string) *CertainAgreementUpdateOne {
	cauo.mutation.SetAgreementOtherDetails(s)
	return cauo
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (cauo *CertainAgreementUpdateOne) SetParentID(id int) *CertainAgreementUpdateOne {
	cauo.mutation.SetParentID(id)
	return cauo
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (cauo *CertainAgreementUpdateOne) SetParent(m *MoreInfoModule) *CertainAgreementUpdateOne {
	return cauo.SetParentID(m.ID)
}

// Mutation returns the CertainAgreementMutation object of the builder.
func (cauo *CertainAgreementUpdateOne) Mutation() *CertainAgreementMutation {
	return cauo.mutation
}

// ClearParent clears the "parent" edge to the MoreInfoModule entity.
func (cauo *CertainAgreementUpdateOne) ClearParent() *CertainAgreementUpdateOne {
	cauo.mutation.ClearParent()
	return cauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CertainAgreementUpdateOne) Select(field string, fields ...string) *CertainAgreementUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CertainAgreement entity.
func (cauo *CertainAgreementUpdateOne) Save(ctx context.Context) (*CertainAgreement, error) {
	var (
		err  error
		node *CertainAgreement
	)
	if len(cauo.hooks) == 0 {
		if err = cauo.check(); err != nil {
			return nil, err
		}
		node, err = cauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertainAgreementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cauo.check(); err != nil {
				return nil, err
			}
			cauo.mutation = mutation
			node, err = cauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cauo.hooks) - 1; i >= 0; i-- {
			if cauo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = cauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CertainAgreementUpdateOne) SaveX(ctx context.Context) *CertainAgreement {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CertainAgreementUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CertainAgreementUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cauo *CertainAgreementUpdateOne) check() error {
	if _, ok := cauo.mutation.ParentID(); cauo.mutation.ParentCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "CertainAgreement.parent"`)
	}
	return nil
}

func (cauo *CertainAgreementUpdateOne) sqlSave(ctx context.Context) (_node *CertainAgreement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   certainagreement.Table,
			Columns: certainagreement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certainagreement.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "CertainAgreement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certainagreement.FieldID)
		for _, f := range fields {
			if !certainagreement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != certainagreement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cauo.mutation.AgreementPiSponsorEmployee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementPiSponsorEmployee,
		})
	}
	if value, ok := cauo.mutation.AgreementRestrictionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictionType,
		})
	}
	if value, ok := cauo.mutation.AgreementRestrictiveAgreement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementRestrictiveAgreement,
		})
	}
	if value, ok := cauo.mutation.AgreementOtherDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certainagreement.FieldAgreementOtherDetails,
		})
	}
	if cauo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certainagreement.ParentTable,
			Columns: []string{certainagreement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cauo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certainagreement.ParentTable,
			Columns: []string{certainagreement.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CertainAgreement{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certainagreement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
