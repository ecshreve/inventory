// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (i *Item) StorageLocation(ctx context.Context) (*StorageLocation, error) {
	result, err := i.Edges.StorageLocationOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryStorageLocation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (sl *StorageLocation) Items(ctx context.Context) (result []*Item, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = sl.NamedItems(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = sl.Edges.ItemsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = sl.QueryItems().All(ctx)
	}
	return result, err
}
