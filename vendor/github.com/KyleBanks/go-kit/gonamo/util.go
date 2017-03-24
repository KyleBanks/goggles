package gonamo

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// AwsStringOrNil returns an aws.String string pointer, or nil if the provided
// string is empty.
func AwsStringOrNil(s string) *string {
	if len(s) == 0 {
		return nil
	}

	return aws.String(s)
}

// AttributeValue initializes and returns an AttributeValue based on the attributeType provided.
func AttributeValue(attributeType AttributeType, attributeValue interface{}) *dynamodb.AttributeValue {
	if attributeValue == nil {
		return nil
	}

	var attr dynamodb.AttributeValue

	switch attributeType {
	case StringType:
		attr.S = aws.String(attributeValue.(string))
	case StringArrayType:
		slice := attributeValue.([]string)
		ss := make([]*string, len(slice))
		for i, s := range slice {
			ss[i] = AwsStringOrNil(s)
		}

		attr.SS = ss
	case NumberType:
		attr.N = aws.String(fmt.Sprintf("%v", attributeValue))
	case BoolType:
		attr.BOOL = aws.Bool(attributeValue.(bool))

	default:
		// TODO: Return an error?
		log.Printf("WARNING: Unknown attributeType provided to attributeValue(%v, %v)\n", attributeType, attributeValue)
	}

	return &attr
}

// attributeDefinition initializes and returns an AttributeDefition.
func attributeDefinition(name string, attributeType AttributeType) *dynamodb.AttributeDefinition {
	return &dynamodb.AttributeDefinition{
		AttributeName: aws.String(name),
		AttributeType: aws.String(string(attributeType)),
	}
}

// keySchemaElement initializes and returns a KeySchemaElement.
func keySchemaElement(name string, keyType KeyType) *dynamodb.KeySchemaElement {
	return &dynamodb.KeySchemaElement{
		AttributeName: aws.String(name),
		KeyType:       aws.String(string(keyType)),
	}
}

// flattenAttributeMap takes a DynamoDB attribute map and flattens it into a map[string]interface{}.
func flattenAttributeMap(m AttributeMap) map[string]interface{} {
	if m == nil {
		return nil
	}

	out := make(map[string]interface{})
	for k, v := range m {
		if v.S != nil {
			out[k] = *v.S
		} else if v.SS != nil {
			ss := make([]string, len(v.SS))
			for i, s := range v.SS {
				ss[i] = *s
			}
			out[k] = ss
		} else if v.BOOL != nil {
			out[k] = *v.BOOL
		} else if v.N != nil {
			out[k], _ = strconv.Atoi(*v.N)
		} else {
			out[k] = nil
		}
	}

	return out
}

// flattenAttributeMaps takes a slice of DynamoDB attribute maps and flattens them
// into a slice of map[string]interface{}.
func flattenAttributeMaps(maps AttributeMaps) []map[string]interface{} {
	if maps == nil {
		return nil
	}

	res := make([]map[string]interface{}, len(maps))
	for i, m := range maps {
		res[i] = flattenAttributeMap(m)
	}

	return res
}
