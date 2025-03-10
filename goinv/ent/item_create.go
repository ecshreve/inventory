// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"goinv/ent/item"
	"goinv/ent/storagelocation"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *ItemCreate) SetCreatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ItemCreate) SetUpdatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *ItemCreate) SetName(s string) *ItemCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetQuantity sets the "quantity" field.
func (ic *ItemCreate) SetQuantity(i int) *ItemCreate {
	ic.mutation.SetQuantity(i)
	return ic
}

// SetCategory sets the "category" field.
func (ic *ItemCreate) SetCategory(i item.Category) *ItemCreate {
	ic.mutation.SetCategory(i)
	return ic
}

// SetStorageLocationID sets the "storage_location" edge to the StorageLocation entity by ID.
func (ic *ItemCreate) SetStorageLocationID(id int) *ItemCreate {
	ic.mutation.SetStorageLocationID(id)
	return ic
}

// SetNillableStorageLocationID sets the "storage_location" edge to the StorageLocation entity by ID if the given value is not nil.
func (ic *ItemCreate) SetNillableStorageLocationID(id *int) *ItemCreate {
	if id != nil {
		ic = ic.SetStorageLocationID(*id)
	}
	return ic
}

// SetStorageLocation sets the "storage_location" edge to the StorageLocation entity.
func (ic *ItemCreate) SetStorageLocation(s *StorageLocation) *ItemCreate {
	return ic.SetStorageLocationID(s.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := item.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := item.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Item.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Item.updated_at"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Item.name"`)}
	}
	if _, ok := ic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Item.quantity"`)}
	}
	if _, ok := ic.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Item.category"`)}
	}
	if v, ok := ic.mutation.Category(); ok {
		if err := item.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Item.category": %w`, err)}
		}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(item.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := ic.mutation.Category(); ok {
		_spec.SetField(item.FieldCategory, field.TypeEnum, value)
		_node.Category = value
	}
	if nodes := ic.mutation.StorageLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   item.StorageLocationTable,
			Columns: []string{item.StorageLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagelocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.item_storage_location = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	err      error
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
