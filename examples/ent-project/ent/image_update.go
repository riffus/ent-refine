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
	"github.com/entkit/entkit/examples/ent-project/ent/company"
	"github.com/entkit/entkit/examples/ent-project/ent/image"
	"github.com/entkit/entkit/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetTitle sets the "title" field.
func (iu *ImageUpdate) SetTitle(s string) *ImageUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetOriginalURL sets the "original_url" field.
func (iu *ImageUpdate) SetOriginalURL(s string) *ImageUpdate {
	iu.mutation.SetOriginalURL(s)
	return iu
}

// SetGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID.
func (iu *ImageUpdate) SetGalleryCompanyID(id uuid.UUID) *ImageUpdate {
	iu.mutation.SetGalleryCompanyID(id)
	return iu
}

// SetNillableGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableGalleryCompanyID(id *uuid.UUID) *ImageUpdate {
	if id != nil {
		iu = iu.SetGalleryCompanyID(*id)
	}
	return iu
}

// SetGalleryCompany sets the "gallery_company" edge to the Company entity.
func (iu *ImageUpdate) SetGalleryCompany(c *Company) *ImageUpdate {
	return iu.SetGalleryCompanyID(c.ID)
}

// SetLogoCompanyID sets the "logo_company" edge to the Company entity by ID.
func (iu *ImageUpdate) SetLogoCompanyID(id uuid.UUID) *ImageUpdate {
	iu.mutation.SetLogoCompanyID(id)
	return iu
}

// SetNillableLogoCompanyID sets the "logo_company" edge to the Company entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableLogoCompanyID(id *uuid.UUID) *ImageUpdate {
	if id != nil {
		iu = iu.SetLogoCompanyID(*id)
	}
	return iu
}

// SetLogoCompany sets the "logo_company" edge to the Company entity.
func (iu *ImageUpdate) SetLogoCompany(c *Company) *ImageUpdate {
	return iu.SetLogoCompanyID(c.ID)
}

// SetCoverCompanyID sets the "cover_company" edge to the Company entity by ID.
func (iu *ImageUpdate) SetCoverCompanyID(id uuid.UUID) *ImageUpdate {
	iu.mutation.SetCoverCompanyID(id)
	return iu
}

// SetNillableCoverCompanyID sets the "cover_company" edge to the Company entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableCoverCompanyID(id *uuid.UUID) *ImageUpdate {
	if id != nil {
		iu = iu.SetCoverCompanyID(*id)
	}
	return iu
}

// SetCoverCompany sets the "cover_company" edge to the Company entity.
func (iu *ImageUpdate) SetCoverCompany(c *Company) *ImageUpdate {
	return iu.SetCoverCompanyID(c.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearGalleryCompany clears the "gallery_company" edge to the Company entity.
func (iu *ImageUpdate) ClearGalleryCompany() *ImageUpdate {
	iu.mutation.ClearGalleryCompany()
	return iu
}

// ClearLogoCompany clears the "logo_company" edge to the Company entity.
func (iu *ImageUpdate) ClearLogoCompany() *ImageUpdate {
	iu.mutation.ClearLogoCompany()
	return iu
}

// ClearCoverCompany clears the "cover_company" edge to the Company entity.
func (iu *ImageUpdate) ClearCoverCompany() *ImageUpdate {
	iu.mutation.ClearCoverCompany()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ImageMutation](ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if v, ok := iu.mutation.Title(); ok {
		if err := image.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Image.title": %w`, err)}
		}
	}
	if v, ok := iu.mutation.OriginalURL(); ok {
		if err := image.OriginalURLValidator(v); err != nil {
			return &ValidationError{Name: "original_url", err: fmt.Errorf(`ent: validator failed for field "Image.original_url": %w`, err)}
		}
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.SetField(image.FieldTitle, field.TypeString, value)
	}
	if value, ok := iu.mutation.OriginalURL(); ok {
		_spec.SetField(image.FieldOriginalURL, field.TypeString, value)
	}
	if iu.mutation.GalleryCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.GalleryCompanyTable,
			Columns: []string{image.GalleryCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.GalleryCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.GalleryCompanyTable,
			Columns: []string{image.GalleryCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.LogoCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.LogoCompanyTable,
			Columns: []string{image.LogoCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.LogoCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.LogoCompanyTable,
			Columns: []string{image.LogoCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.CoverCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.CoverCompanyTable,
			Columns: []string{image.CoverCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CoverCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.CoverCompanyTable,
			Columns: []string{image.CoverCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetTitle sets the "title" field.
func (iuo *ImageUpdateOne) SetTitle(s string) *ImageUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetOriginalURL sets the "original_url" field.
func (iuo *ImageUpdateOne) SetOriginalURL(s string) *ImageUpdateOne {
	iuo.mutation.SetOriginalURL(s)
	return iuo
}

// SetGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID.
func (iuo *ImageUpdateOne) SetGalleryCompanyID(id uuid.UUID) *ImageUpdateOne {
	iuo.mutation.SetGalleryCompanyID(id)
	return iuo
}

// SetNillableGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableGalleryCompanyID(id *uuid.UUID) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetGalleryCompanyID(*id)
	}
	return iuo
}

// SetGalleryCompany sets the "gallery_company" edge to the Company entity.
func (iuo *ImageUpdateOne) SetGalleryCompany(c *Company) *ImageUpdateOne {
	return iuo.SetGalleryCompanyID(c.ID)
}

// SetLogoCompanyID sets the "logo_company" edge to the Company entity by ID.
func (iuo *ImageUpdateOne) SetLogoCompanyID(id uuid.UUID) *ImageUpdateOne {
	iuo.mutation.SetLogoCompanyID(id)
	return iuo
}

// SetNillableLogoCompanyID sets the "logo_company" edge to the Company entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableLogoCompanyID(id *uuid.UUID) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetLogoCompanyID(*id)
	}
	return iuo
}

// SetLogoCompany sets the "logo_company" edge to the Company entity.
func (iuo *ImageUpdateOne) SetLogoCompany(c *Company) *ImageUpdateOne {
	return iuo.SetLogoCompanyID(c.ID)
}

// SetCoverCompanyID sets the "cover_company" edge to the Company entity by ID.
func (iuo *ImageUpdateOne) SetCoverCompanyID(id uuid.UUID) *ImageUpdateOne {
	iuo.mutation.SetCoverCompanyID(id)
	return iuo
}

// SetNillableCoverCompanyID sets the "cover_company" edge to the Company entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableCoverCompanyID(id *uuid.UUID) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetCoverCompanyID(*id)
	}
	return iuo
}

// SetCoverCompany sets the "cover_company" edge to the Company entity.
func (iuo *ImageUpdateOne) SetCoverCompany(c *Company) *ImageUpdateOne {
	return iuo.SetCoverCompanyID(c.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearGalleryCompany clears the "gallery_company" edge to the Company entity.
func (iuo *ImageUpdateOne) ClearGalleryCompany() *ImageUpdateOne {
	iuo.mutation.ClearGalleryCompany()
	return iuo
}

// ClearLogoCompany clears the "logo_company" edge to the Company entity.
func (iuo *ImageUpdateOne) ClearLogoCompany() *ImageUpdateOne {
	iuo.mutation.ClearLogoCompany()
	return iuo
}

// ClearCoverCompany clears the "cover_company" edge to the Company entity.
func (iuo *ImageUpdateOne) ClearCoverCompany() *ImageUpdateOne {
	iuo.mutation.ClearCoverCompany()
	return iuo
}

// Where appends a list predicates to the ImageUpdate builder.
func (iuo *ImageUpdateOne) Where(ps ...predicate.Image) *ImageUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	return withHooks[*Image, ImageMutation](ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if v, ok := iuo.mutation.Title(); ok {
		if err := image.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Image.title": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.OriginalURL(); ok {
		if err := image.OriginalURLValidator(v); err != nil {
			return &ValidationError{Name: "original_url", err: fmt.Errorf(`ent: validator failed for field "Image.original_url": %w`, err)}
		}
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.SetField(image.FieldTitle, field.TypeString, value)
	}
	if value, ok := iuo.mutation.OriginalURL(); ok {
		_spec.SetField(image.FieldOriginalURL, field.TypeString, value)
	}
	if iuo.mutation.GalleryCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.GalleryCompanyTable,
			Columns: []string{image.GalleryCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.GalleryCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.GalleryCompanyTable,
			Columns: []string{image.GalleryCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.LogoCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.LogoCompanyTable,
			Columns: []string{image.LogoCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.LogoCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.LogoCompanyTable,
			Columns: []string{image.LogoCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.CoverCompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.CoverCompanyTable,
			Columns: []string{image.CoverCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CoverCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.CoverCompanyTable,
			Columns: []string{image.CoverCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
