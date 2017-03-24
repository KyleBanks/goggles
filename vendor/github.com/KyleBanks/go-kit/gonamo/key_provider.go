package gonamo

// keyProvider makes up the definition of Hash and Range keys used to
// construct and interact with a DynamoDB table.
type keyProvider interface {
	// hashName returns the name of the Hash portion of the table's Key.
	hashName() string
	// hashType returns the type of the Hash portion of the table's Key.
	hashType() AttributeType

	// hasRange returns a boolean indicating if the table has a Range key.
	hasRange() bool
	// rangeName returns the name of the Range portion of the table's Key,
	// if applicable.
	rangeName() string
	// rangeType returns the type of the Range portion of the table's Key,
	// if applicable.
	rangeType() AttributeType
}

// HashKeyDefinition implements the KeyProvider interface for a table
// utilizing only a Hash key element.
type HashKeyDefinition struct {
	HashName string
	HashType AttributeType
}

func (h HashKeyDefinition) hashName() string {
	return h.HashName
}

func (h HashKeyDefinition) hashType() AttributeType {
	return h.HashType
}

func (h HashKeyDefinition) hasRange() bool {
	return false
}

func (h HashKeyDefinition) rangeName() string {
	return ""
}

func (h HashKeyDefinition) rangeType() AttributeType {
	return ""
}

// HashRangeKeyDefinition implements the KeyProvider interface for a table
// utilizing Hash and Range key elements.
type HashRangeKeyDefinition struct {
	HashName string
	HashType AttributeType

	RangeName string
	RangeType AttributeType
}

func (h HashRangeKeyDefinition) hashName() string {
	return h.HashName
}

func (h HashRangeKeyDefinition) hashType() AttributeType {
	return h.HashType
}

func (h HashRangeKeyDefinition) hasRange() bool {
	return true
}

func (h HashRangeKeyDefinition) rangeName() string {
	return h.RangeName
}

func (h HashRangeKeyDefinition) rangeType() AttributeType {
	return h.RangeType
}
