// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/examples/ent-project/ent/predicate"
	"github.com/entkit/entkit/examples/ent-project/ent/product"
	"github.com/entkit/entkit/examples/ent-project/ent/vendor"
	"github.com/entkit/entkit/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// VendorUpdate is the builder for updating Vendor entities.
type VendorUpdate struct {
	config
	hooks    []Hook
	mutation *VendorMutation
}

// Where appends a list predicates to the VendorUpdate builder.
func (vu *VendorUpdate) Where(ps ...predicate.Vendor) *VendorUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetName sets the "name" field.
func (vu *VendorUpdate) SetName(s string) *VendorUpdate {
	vu.mutation.SetName(s)
	return vu
}

// SetSchema sets the "schema" field.
func (vu *VendorUpdate) SetSchema(s string) *VendorUpdate {
	vu.mutation.SetSchema(s)
	return vu
}

// AddWarehouseIDs adds the "warehouses" edge to the Warehouse entity by IDs.
func (vu *VendorUpdate) AddWarehouseIDs(ids ...uuid.UUID) *VendorUpdate {
	vu.mutation.AddWarehouseIDs(ids...)
	return vu
}

// AddWarehouses adds the "warehouses" edges to the Warehouse entity.
func (vu *VendorUpdate) AddWarehouses(w ...*Warehouse) *VendorUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return vu.AddWarehouseIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (vu *VendorUpdate) AddProductIDs(ids ...uuid.UUID) *VendorUpdate {
	vu.mutation.AddProductIDs(ids...)
	return vu
}

// AddProducts adds the "products" edges to the Product entity.
func (vu *VendorUpdate) AddProducts(p ...*Product) *VendorUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vu.AddProductIDs(ids...)
}

// Mutation returns the VendorMutation object of the builder.
func (vu *VendorUpdate) Mutation() *VendorMutation {
	return vu.mutation
}

// ClearWarehouses clears all "warehouses" edges to the Warehouse entity.
func (vu *VendorUpdate) ClearWarehouses() *VendorUpdate {
	vu.mutation.ClearWarehouses()
	return vu
}

// RemoveWarehouseIDs removes the "warehouses" edge to Warehouse entities by IDs.
func (vu *VendorUpdate) RemoveWarehouseIDs(ids ...uuid.UUID) *VendorUpdate {
	vu.mutation.RemoveWarehouseIDs(ids...)
	return vu
}

// RemoveWarehouses removes "warehouses" edges to Warehouse entities.
func (vu *VendorUpdate) RemoveWarehouses(w ...*Warehouse) *VendorUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return vu.RemoveWarehouseIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (vu *VendorUpdate) ClearProducts() *VendorUpdate {
	vu.mutation.ClearProducts()
	return vu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (vu *VendorUpdate) RemoveProductIDs(ids ...uuid.UUID) *VendorUpdate {
	vu.mutation.RemoveProductIDs(ids...)
	return vu
}

// RemoveProducts removes "products" edges to Product entities.
func (vu *VendorUpdate) RemoveProducts(p ...*Product) *VendorUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VendorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, VendorMutation](ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VendorUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VendorUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VendorUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VendorUpdate) check() error {
	if v, ok := vu.mutation.Name(); ok {
		if err := vendor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vendor.name": %w`, err)}
		}
	}
	return nil
}

func (vu *VendorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(vendor.Table, vendor.Columns, sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.SetField(vendor.FieldName, field.TypeString, value)
	}
	if value, ok := vu.mutation.Schema(); ok {
		_spec.SetField(vendor.FieldSchema, field.TypeString, value)
	}
	if vu.mutation.WarehousesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedWarehousesIDs(); len(nodes) > 0 && !vu.mutation.WarehousesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.WarehousesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !vu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vendor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VendorUpdateOne is the builder for updating a single Vendor entity.
type VendorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VendorMutation
}

// SetName sets the "name" field.
func (vuo *VendorUpdateOne) SetName(s string) *VendorUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// SetSchema sets the "schema" field.
func (vuo *VendorUpdateOne) SetSchema(s string) *VendorUpdateOne {
	vuo.mutation.SetSchema(s)
	return vuo
}

// AddWarehouseIDs adds the "warehouses" edge to the Warehouse entity by IDs.
func (vuo *VendorUpdateOne) AddWarehouseIDs(ids ...uuid.UUID) *VendorUpdateOne {
	vuo.mutation.AddWarehouseIDs(ids...)
	return vuo
}

// AddWarehouses adds the "warehouses" edges to the Warehouse entity.
func (vuo *VendorUpdateOne) AddWarehouses(w ...*Warehouse) *VendorUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return vuo.AddWarehouseIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (vuo *VendorUpdateOne) AddProductIDs(ids ...uuid.UUID) *VendorUpdateOne {
	vuo.mutation.AddProductIDs(ids...)
	return vuo
}

// AddProducts adds the "products" edges to the Product entity.
func (vuo *VendorUpdateOne) AddProducts(p ...*Product) *VendorUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vuo.AddProductIDs(ids...)
}

// Mutation returns the VendorMutation object of the builder.
func (vuo *VendorUpdateOne) Mutation() *VendorMutation {
	return vuo.mutation
}

// ClearWarehouses clears all "warehouses" edges to the Warehouse entity.
func (vuo *VendorUpdateOne) ClearWarehouses() *VendorUpdateOne {
	vuo.mutation.ClearWarehouses()
	return vuo
}

// RemoveWarehouseIDs removes the "warehouses" edge to Warehouse entities by IDs.
func (vuo *VendorUpdateOne) RemoveWarehouseIDs(ids ...uuid.UUID) *VendorUpdateOne {
	vuo.mutation.RemoveWarehouseIDs(ids...)
	return vuo
}

// RemoveWarehouses removes "warehouses" edges to Warehouse entities.
func (vuo *VendorUpdateOne) RemoveWarehouses(w ...*Warehouse) *VendorUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return vuo.RemoveWarehouseIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (vuo *VendorUpdateOne) ClearProducts() *VendorUpdateOne {
	vuo.mutation.ClearProducts()
	return vuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (vuo *VendorUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *VendorUpdateOne {
	vuo.mutation.RemoveProductIDs(ids...)
	return vuo
}

// RemoveProducts removes "products" edges to Product entities.
func (vuo *VendorUpdateOne) RemoveProducts(p ...*Product) *VendorUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vuo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the VendorUpdate builder.
func (vuo *VendorUpdateOne) Where(ps ...predicate.Vendor) *VendorUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VendorUpdateOne) Select(field string, fields ...string) *VendorUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vendor entity.
func (vuo *VendorUpdateOne) Save(ctx context.Context) (*Vendor, error) {
	return withHooks[*Vendor, VendorMutation](ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VendorUpdateOne) SaveX(ctx context.Context) *Vendor {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VendorUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VendorUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VendorUpdateOne) check() error {
	if v, ok := vuo.mutation.Name(); ok {
		if err := vendor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vendor.name": %w`, err)}
		}
	}
	return nil
}

func (vuo *VendorUpdateOne) sqlSave(ctx context.Context) (_node *Vendor, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(vendor.Table, vendor.Columns, sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vendor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vendor.FieldID)
		for _, f := range fields {
			if !vendor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vendor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Name(); ok {
		_spec.SetField(vendor.FieldName, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Schema(); ok {
		_spec.SetField(vendor.FieldSchema, field.TypeString, value)
	}
	if vuo.mutation.WarehousesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedWarehousesIDs(); len(nodes) > 0 && !vuo.mutation.WarehousesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.WarehousesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.WarehousesTable,
			Columns: []string{vendor.WarehousesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !vuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ProductsTable,
			Columns: []string{vendor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Vendor{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vendor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
