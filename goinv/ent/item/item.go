// Code generated by ent, DO NOT EDIT.

package item

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// EdgeStorageLocation holds the string denoting the storage_location edge name in mutations.
	EdgeStorageLocation = "storage_location"
	// Table holds the table name of the item in the database.
	Table = "items"
	// StorageLocationTable is the table that holds the storage_location relation/edge.
	StorageLocationTable = "items"
	// StorageLocationInverseTable is the table name for the StorageLocation entity.
	// It exists in this package in order to avoid circular dependency with the "storagelocation" package.
	StorageLocationInverseTable = "storage_locations"
	// StorageLocationColumn is the table column denoting the storage_location relation/edge.
	StorageLocationColumn = "item_storage_location"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldQuantity,
	FieldCategory,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"item_storage_location",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryAdapter Category = "adapter"
	CategoryCable   Category = "cable"
	CategoryDevice  Category = "device"
	CategoryMisc    Category = "misc"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryAdapter, CategoryCable, CategoryDevice, CategoryMisc:
		return nil
	default:
		return fmt.Errorf("item: invalid enum value for category field: %q", c)
	}
}

// OrderOption defines the ordering options for the Item queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByStorageLocationField orders the results by storage_location field.
func ByStorageLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStorageLocationStep(), sql.OrderByField(field, opts...))
	}
}
func newStorageLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StorageLocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StorageLocationTable, StorageLocationColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Category) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Category) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Category(str)
	if err := CategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}
