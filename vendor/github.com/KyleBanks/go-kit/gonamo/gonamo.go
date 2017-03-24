// Package gonamo provides a simple wrapper around the DynamoDB SDK.
//
// The intention is to provide a minimal DynamoDB table representation
// that can be created, written to, queried and scanned.
//
// The main starting point of a gonamo implementation is to define a model that
// you want to store in DynamoDB, and implement the "Persistable" interface:
//
// 	type Repository struct {
// 		Owner string
// 		Name  string
//
//		DateCreated time.Time
// 	}
//
// 	func (r Repository) HashKey() interface{} {
// 		return r.Owner
// 	}
//
// 	func (r Repository) RangeKey() interface{} {
// 		return r.Name
// 	}
//
// 	func (r Repository) Attributes() gonamo.AttributeMap {
// 		return := gonamo.AttributeMap{
// 			"owner":       gonamo.AttributeValue(gonamo.StringType, r.Owner),
// 			"name":        gonamo.AttributeValue(gonamo.StringType, r.Name),
// 			"dateCreated": gonamo.AttributeValue(gonamo.NumberType, r.DateCreated.Unix()),
// 		}
// 	}
//
// Next you will be able to create a Table using the NewTable method, by providing a
// table name and the key structure:
//
// 	gonamo.HashRangeKeyDefinition{"owner", gonamo.StringType, "name", gonamo.StringType}
// 	tbl, err := gonamo.NewTable("tableName", key, nil)
package gonamo
