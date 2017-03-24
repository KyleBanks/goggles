package gonamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// AttributeType defines a type of attribute used when representing data in DynamoDB.
type AttributeType string

const (
	// StringType is used for string fields.
	StringType AttributeType = "S"
	// NumberType is used for numeric fields.
	NumberType AttributeType = "N"
	// BoolType is used for boolean fields.
	BoolType AttributeType = "BOOL"
	// StringArrayType is used for string array types.
	StringArrayType AttributeType = "SS"
)

// AttributeMap represents a DynamoDB object in it's
// database representation. Each element of the map is
// identified by it's name, and contains the element
// type and value.
type AttributeMap map[string]*dynamodb.AttributeValue

// AttributeMaps is a slice of AttributeMap.
type AttributeMaps []map[string]*dynamodb.AttributeValue

// KeyType defines a type of Key (hash or range).
type KeyType string

const (
	// HashKey represents a DyanmoDB Hash key.
	HashKey KeyType = "HASH"
	// RangeKey represents a DynamoDB Range key.
	RangeKey KeyType = "RANGE"
)
