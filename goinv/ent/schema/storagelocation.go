package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StorageLocation holds the schema definition for the StorageLocation entity.
type StorageLocation struct {
	ent.Schema
}

func (StorageLocation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the StorageLocation.
func (StorageLocation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("size").Values("half", "full", "other"),
		field.Enum("color").Values("white", "gray", "stealth", "orange", "black", "other"),
	}
}

// Edges of the StorageLocation.
func (StorageLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("items", Item.Type).
			Ref("storage_location"),
	}
}
