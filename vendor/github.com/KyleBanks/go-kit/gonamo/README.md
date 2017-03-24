# gonamo
--
    import "github.com/KyleBanks/go-kit/gonamo/"

Package gonamo provides a simple wrapper around the DynamoDB SDK.

The intention is to provide a minimal DynamoDB table representation that can be
created, written to, queried and scanned.

The main starting point of a gonamo implementation is to define a model that you
want to store in DynamoDB, and implement the "Persistable" interface:

    type Repository struct {
    	Owner string
    	Name  string

    	DateCreated time.Time
    }

    func (r Repository) HashKey() interface{} {
    	return r.Owner
    }

    func (r Repository) RangeKey() interface{} {
    	return r.Name
    }

    func (r Repository) Attributes() gonamo.AttributeMap {
    	return := gonamo.AttributeMap{
    		"owner":       gonamo.AttributeValue(gonamo.StringType, r.Owner),
    		"name":        gonamo.AttributeValue(gonamo.StringType, r.Name),
    		"dateCreated": gonamo.AttributeValue(gonamo.NumberType, r.DateCreated.Unix()),
    	}
    }

Next you will be able to create a Table using the NewTable method, by providing
a table name and the key structure:

    gonamo.HashRangeKeyDefinition{"owner", gonamo.StringType, "name", gonamo.StringType}
    tbl, err := gonamo.NewTable("tableName", key, nil)

## Usage

#### func  AttributeValue

```go
func AttributeValue(attributeType AttributeType, attributeValue interface{}) *dynamodb.AttributeValue
```
AttributeValue initializes and returns an AttributeValue based on the
attributeType provided.

#### func  AwsStringOrNil

```go
func AwsStringOrNil(s string) *string
```
AwsStringOrNil returns an aws.String string pointer, or nil if the provided
string is empty.

#### type AttributeMap

```go
type AttributeMap map[string]*dynamodb.AttributeValue
```

AttributeMap represents a DynamoDB object in it's database representation. Each
element of the map is identified by it's name, and contains the element type and
value.

#### type AttributeMaps

```go
type AttributeMaps []map[string]*dynamodb.AttributeValue
```

AttributeMaps is a slice of AttributeMap.

#### type AttributeType

```go
type AttributeType string
```

AttributeType defines a type of attribute used when representing data in
DynamoDB.

```go
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
```

#### type HashKeyDefinition

```go
type HashKeyDefinition struct {
	HashName string
	HashType AttributeType
}
```

HashKeyDefinition implements the KeyProvider interface for a table utilizing
only a Hash key element.

#### type HashRangeKeyDefinition

```go
type HashRangeKeyDefinition struct {
	HashName string
	HashType AttributeType

	RangeName string
	RangeType AttributeType
}
```

HashRangeKeyDefinition implements the KeyProvider interface for a table
utilizing Hash and Range key elements.

#### type KeyType

```go
type KeyType string
```

KeyType defines a type of Key (hash or range).

```go
const (
	// HashKey represents a DyanmoDB Hash key.
	HashKey KeyType = "HASH"
	// RangeKey represents a DynamoDB Range key.
	RangeKey KeyType = "RANGE"
)
```

#### type Options

```go
type Options struct {
	Endpoint            string
	Region              string
	DefaultProvisioning int64
}
```

Options represents optional configurations for DynamoDB.

#### type Persistable

```go
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
```

Persistable provides an interface for models that can be persisted.

#### type Table

```go
type Table struct {
}
```

Table provides a Dynamo API wrapper around a specific table.

#### func  NewTable

```go
func NewTable(tableName string, k keyProvider, opts *Options) (*Table, error)
```
NewTable intializes and returns a Table.

Providing the optional Options allows you to further configure the table, and
the client used to communicate with the DynamoDB service.

#### func (*Table) BatchDelete

```go
func (d *Table) BatchDelete(persistables []Persistable) error
```
BatchDelete performs a batch of delete requests, one for each of the
Persistables provided.

#### func (*Table) BatchWrite

```go
func (d *Table) BatchWrite(persistables []Persistable) error
```
BatchWrite handles paging and sending a batch write against a single table.

#### func (*Table) CreateTableIfNecessary

```go
func (d *Table) CreateTableIfNecessary() (bool, error)
```
CreateTableIfNecessary attempts to determine if a table already exists, and if
not, will create the table.

The boolean returned will be true if the table had to be created, and false
otherwise.

#### func (*Table) DeleteItem

```go
func (d *Table) DeleteItem(p Persistable) (*dynamodb.DeleteItemOutput, error)
```
DeleteItem deletes the item identified by the provided hash and (optional) range
key.

#### func (*Table) Describe

```go
func (d *Table) Describe() (*dynamodb.DescribeTableOutput, error)
```
Describe returns the DynamoDB Table description.

#### func (*Table) Find

```go
func (d *Table) Find(hashValue, rangeValue interface{}) (map[string]interface{}, error)
```
Find performs a GetItem on the underlying table with a particular hash and
(optional) range key.

If the table has no range key, this would be the same output as FindAllByHash.

#### func (*Table) FindAllByHash

```go
func (d *Table) FindAllByHash(hashValue interface{}) ([]map[string]interface{}, error)
```
FindAllByHash finds all items with the specified hash key value.

If the table has no range key, this would be the same output as Find.

#### func (*Table) PutItem

```go
func (d *Table) PutItem(item AttributeMap) (*dynamodb.PutItemOutput, error)
```
PutItem performs an insert into DynamoDB.

#### func (*Table) ScanAll

```go
func (d *Table) ScanAll(fn func(page []map[string]interface{}) bool) error
```
ScanAll performs a scan on the entire underlying DynamoDB table, and executes
the provided function for each page of results. If the function returns false,
the scan will complete.
