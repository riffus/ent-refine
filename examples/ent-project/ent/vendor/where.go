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

package vendor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/entkit/entkit/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldName, v))
}

// Schema applies equality check predicate on the "schema" field. It's identical to SchemaEQ.
func Schema(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldSchema, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContainsFold(FieldName, v))
}

// SchemaEQ applies the EQ predicate on the "schema" field.
func SchemaEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldSchema, v))
}

// SchemaNEQ applies the NEQ predicate on the "schema" field.
func SchemaNEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldSchema, v))
}

// SchemaIn applies the In predicate on the "schema" field.
func SchemaIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldSchema, vs...))
}

// SchemaNotIn applies the NotIn predicate on the "schema" field.
func SchemaNotIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldSchema, vs...))
}

// SchemaGT applies the GT predicate on the "schema" field.
func SchemaGT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldSchema, v))
}

// SchemaGTE applies the GTE predicate on the "schema" field.
func SchemaGTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldSchema, v))
}

// SchemaLT applies the LT predicate on the "schema" field.
func SchemaLT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldSchema, v))
}

// SchemaLTE applies the LTE predicate on the "schema" field.
func SchemaLTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldSchema, v))
}

// SchemaContains applies the Contains predicate on the "schema" field.
func SchemaContains(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContains(FieldSchema, v))
}

// SchemaHasPrefix applies the HasPrefix predicate on the "schema" field.
func SchemaHasPrefix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasPrefix(FieldSchema, v))
}

// SchemaHasSuffix applies the HasSuffix predicate on the "schema" field.
func SchemaHasSuffix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasSuffix(FieldSchema, v))
}

// SchemaEqualFold applies the EqualFold predicate on the "schema" field.
func SchemaEqualFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEqualFold(FieldSchema, v))
}

// SchemaContainsFold applies the ContainsFold predicate on the "schema" field.
func SchemaContainsFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContainsFold(FieldSchema, v))
}

// HasWarehouses applies the HasEdge predicate on the "warehouses" edge.
func HasWarehouses() predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WarehousesTable, WarehousesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWarehousesWith applies the HasEdge predicate on the "warehouses" edge with a given conditions (other predicates).
func HasWarehousesWith(preds ...predicate.Warehouse) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WarehousesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WarehousesTable, WarehousesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		p(s.Not())
	})
}
