package gonamo

// Persistable provides an interface for models that can be persisted.
type Persistable interface {
	// HashKey returns the value of the Hash key used to represent this individual entity.
	HashKey() interface{}

	// RangeKey returns the value of the Range key used to represent this individual entity.
	// If the table has no Range key, this may simply return nil.
	RangeKey() interface{}

	// Attributes returns all attributes (including Hash and Range keys)
	// and their type for the persistable model.
	Attributes() AttributeMap
}
