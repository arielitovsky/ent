// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/privacy/ent/planet"
	"github.com/facebook/ent/entc/integration/privacy/ent/predicate"
	"github.com/facebook/ent/schema/field"
)

// PlanetUpdate is the builder for updating Planet entities.
type PlanetUpdate struct {
	config
	hooks      []Hook
	mutation   *PlanetMutation
	predicates []predicate.Planet
}

// Where adds a new predicate for the builder.
func (pu *PlanetUpdate) Where(ps ...predicate.Planet) *PlanetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAge sets the age field.
func (pu *PlanetUpdate) SetAge(u uint) *PlanetUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(u)
	return pu
}

// SetNillableAge sets the age field if the given value is not nil.
func (pu *PlanetUpdate) SetNillableAge(u *uint) *PlanetUpdate {
	if u != nil {
		pu.SetAge(*u)
	}
	return pu
}

// AddAge adds u to age.
func (pu *PlanetUpdate) AddAge(u uint) *PlanetUpdate {
	pu.mutation.AddAge(u)
	return pu
}

// ClearAge clears the value of age.
func (pu *PlanetUpdate) ClearAge() *PlanetUpdate {
	pu.mutation.ClearAge()
	return pu
}

// AddNeighborIDs adds the neighbors edge to Planet by ids.
func (pu *PlanetUpdate) AddNeighborIDs(ids ...int) *PlanetUpdate {
	pu.mutation.AddNeighborIDs(ids...)
	return pu
}

// AddNeighbors adds the neighbors edges to Planet.
func (pu *PlanetUpdate) AddNeighbors(p ...*Planet) *PlanetUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddNeighborIDs(ids...)
}

// Mutation returns the PlanetMutation object of the builder.
func (pu *PlanetUpdate) Mutation() *PlanetMutation {
	return pu.mutation
}

// RemoveNeighborIDs removes the neighbors edge to Planet by ids.
func (pu *PlanetUpdate) RemoveNeighborIDs(ids ...int) *PlanetUpdate {
	pu.mutation.RemoveNeighborIDs(ids...)
	return pu
}

// RemoveNeighbors removes neighbors edges to Planet.
func (pu *PlanetUpdate) RemoveNeighbors(p ...*Planet) *PlanetUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveNeighborIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PlanetUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlanetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlanetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlanetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlanetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   planet.Table,
			Columns: planet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: planet.FieldAge,
		})
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: planet.FieldAge,
		})
	}
	if pu.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Column: planet.FieldAge,
		})
	}
	if nodes := pu.mutation.RemovedNeighborsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   planet.NeighborsTable,
			Columns: planet.NeighborsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NeighborsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   planet.NeighborsTable,
			Columns: planet.NeighborsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{planet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PlanetUpdateOne is the builder for updating a single Planet entity.
type PlanetUpdateOne struct {
	config
	hooks    []Hook
	mutation *PlanetMutation
}

// SetAge sets the age field.
func (puo *PlanetUpdateOne) SetAge(u uint) *PlanetUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(u)
	return puo
}

// SetNillableAge sets the age field if the given value is not nil.
func (puo *PlanetUpdateOne) SetNillableAge(u *uint) *PlanetUpdateOne {
	if u != nil {
		puo.SetAge(*u)
	}
	return puo
}

// AddAge adds u to age.
func (puo *PlanetUpdateOne) AddAge(u uint) *PlanetUpdateOne {
	puo.mutation.AddAge(u)
	return puo
}

// ClearAge clears the value of age.
func (puo *PlanetUpdateOne) ClearAge() *PlanetUpdateOne {
	puo.mutation.ClearAge()
	return puo
}

// AddNeighborIDs adds the neighbors edge to Planet by ids.
func (puo *PlanetUpdateOne) AddNeighborIDs(ids ...int) *PlanetUpdateOne {
	puo.mutation.AddNeighborIDs(ids...)
	return puo
}

// AddNeighbors adds the neighbors edges to Planet.
func (puo *PlanetUpdateOne) AddNeighbors(p ...*Planet) *PlanetUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddNeighborIDs(ids...)
}

// Mutation returns the PlanetMutation object of the builder.
func (puo *PlanetUpdateOne) Mutation() *PlanetMutation {
	return puo.mutation
}

// RemoveNeighborIDs removes the neighbors edge to Planet by ids.
func (puo *PlanetUpdateOne) RemoveNeighborIDs(ids ...int) *PlanetUpdateOne {
	puo.mutation.RemoveNeighborIDs(ids...)
	return puo
}

// RemoveNeighbors removes neighbors edges to Planet.
func (puo *PlanetUpdateOne) RemoveNeighbors(p ...*Planet) *PlanetUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveNeighborIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PlanetUpdateOne) Save(ctx context.Context) (*Planet, error) {

	var (
		err  error
		node *Planet
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlanetUpdateOne) SaveX(ctx context.Context) *Planet {
	pl, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pl
}

// Exec executes the query on the entity.
func (puo *PlanetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlanetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlanetUpdateOne) sqlSave(ctx context.Context) (pl *Planet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   planet.Table,
			Columns: planet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Planet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: planet.FieldAge,
		})
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: planet.FieldAge,
		})
	}
	if puo.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Column: planet.FieldAge,
		})
	}
	if nodes := puo.mutation.RemovedNeighborsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   planet.NeighborsTable,
			Columns: planet.NeighborsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NeighborsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   planet.NeighborsTable,
			Columns: planet.NeighborsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pl = &Planet{config: puo.config}
	_spec.Assign = pl.assignValues
	_spec.ScanValues = pl.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{planet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pl, nil
}
