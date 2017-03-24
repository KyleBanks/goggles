package gonamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	errTableDoesntExist string = "ResourceNotFoundException"

	maxBatchWriteLen int = 25
)

// Table provides a Dynamo API wrapper around a specific table.
type Table struct {
	client *dynamodb.DynamoDB

	tableName           string
	defaultProvisioning int64

	key keyProvider
}

// NewTable intializes and returns a Table.
//
// Providing the optional Options allows you to further configure the table,
// and the client used to communicate with the DynamoDB service.
func NewTable(tableName string, k keyProvider, opts *Options) (*Table, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	if opts == nil {
		opts = defaultOptions()
	}

	t := Table{
		client: dynamodb.New(sess, &aws.Config{
			Region:   AwsStringOrNil(opts.Region),
			Endpoint: AwsStringOrNil(opts.Endpoint),
		}),

		tableName:           tableName,
		defaultProvisioning: opts.DefaultProvisioning,

		key: k,
	}

	return &t, nil
}

// CreateTableIfNecessary attempts to determine if a table already exists, and if not,
// will create the table.
//
// The boolean returned will be true if the table had to be created, and false otherwise.
func (d *Table) CreateTableIfNecessary() (bool, error) {
	if _, err := d.Describe(); err == nil { // Table exists
		return false, nil
	} else if err.(awserr.Error).Code() != errTableDoesntExist { // Different error case
		return false, err
	}

	// Table doesn't exist, create it.
	return true, d.createTable()
}

// Describe returns the DynamoDB Table description.
func (d *Table) Describe() (*dynamodb.DescribeTableOutput, error) {
	return d.client.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(d.tableName),
	})
}

// Find performs a GetItem on the underlying table with a particular hash and (optional) range key.
//
// If the table has no range key, this would be the same output as FindAllByHash.
func (d *Table) Find(hashValue, rangeValue interface{}) (map[string]interface{}, error) {
	req := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			d.key.hashName(): AttributeValue(d.key.hashType(), hashValue),
		},
		TableName: aws.String(d.tableName),
	}

	if d.key.hasRange() {
		req.Key[d.key.rangeName()] = AttributeValue(d.key.rangeType(), rangeValue)
	}

	res, err := d.client.GetItem(&req)
	if err != nil {
		return nil, err
	}

	return flattenAttributeMap(res.Item), nil
}

// FindAllByHash finds all items with the specified hash key value.
//
// If the table has no range key, this would be the same output as Find.
func (d *Table) FindAllByHash(hashValue interface{}) ([]map[string]interface{}, error) {
	req := dynamodb.QueryInput{
		KeyConditions: map[string]*dynamodb.Condition{
			d.key.hashName(): {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					AttributeValue(d.key.hashType(), hashValue),
				},
			},
		},
		TableName: aws.String(d.tableName),
	}

	res, err := d.client.Query(&req)
	if err != nil {
		return nil, err
	}

	return flattenAttributeMaps(res.Items), nil
}

// PutItem performs an insert into DynamoDB.
func (d *Table) PutItem(item AttributeMap) (*dynamodb.PutItemOutput, error) {
	req := dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.tableName),
	}

	return d.client.PutItem(&req)
}

// DeleteItem deletes the item identified by the provided hash and (optional) range key.
func (d *Table) DeleteItem(p Persistable) (*dynamodb.DeleteItemOutput, error) {
	req := dynamodb.DeleteItemInput{
		Key:       d.keyAttributeMap(p),
		TableName: aws.String(d.tableName),
	}

	return d.client.DeleteItem(&req)
}

// BatchDelete performs a batch of delete requests, one for each of the Persistables provided.
func (d *Table) BatchDelete(persistables []Persistable) error {
	deletes := make([]*dynamodb.WriteRequest, len(persistables))
	for i, p := range persistables {
		deletes[i] = &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: d.keyAttributeMap(p),
			},
		}
	}

	return d.batch(deletes)
}

// BatchWrite handles paging and sending a batch write against a single table.
func (d *Table) BatchWrite(persistables []Persistable) error {
	writes := make([]*dynamodb.WriteRequest, len(persistables))
	for i, p := range persistables {
		writes[i] = &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: p.Attributes(),
			},
		}
	}

	return d.batch(writes)
}

// ScanAll performs a scan on the entire underlying DynamoDB table, and executes the provided function
// for each page of results. If the function returns false, the scan will complete.
func (d *Table) ScanAll(fn func(page []map[string]interface{}) bool) error {
	req := dynamodb.ScanInput{
		TableName: aws.String(d.tableName),
	}

	return d.client.ScanPages(&req, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		return fn(flattenAttributeMaps(page.Items))
	})
}

// createTable creates a Dynamo table using the provided table definition.
func (d *Table) createTable() error {
	req := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{attributeDefinition(d.key.hashName(), d.key.hashType())},
		KeySchema:            []*dynamodb.KeySchemaElement{keySchemaElement(d.key.hashName(), HashKey)},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(d.defaultProvisioning),
			WriteCapacityUnits: aws.Int64(d.defaultProvisioning),
		},
		TableName: aws.String(d.tableName),
	}

	if d.key.hasRange() {
		req.AttributeDefinitions = append(req.AttributeDefinitions, attributeDefinition(d.key.rangeName(), d.key.rangeType()))
		req.KeySchema = append(req.KeySchema, keySchemaElement(d.key.rangeName(), RangeKey))
	}

	_, err := d.client.CreateTable(req)
	if err != nil {
		return err
	}

	return nil
}

// batch executes a batch of writes/deletes on the underlying DynamoDB
// Table.
//
// The batch is executed in pages if necessary, ensuring the DynamoDB batch
// size limit is respected.
func (d *Table) batch(writes []*dynamodb.WriteRequest) error {
	if len(writes) == 0 {
		return nil
	}

	var page []*dynamodb.WriteRequest
	for i, w := range writes {
		page = append(page, w)

		// Each time we've got 25 writes, or if it's the end of the writes slice,
		// send the batch.
		if i == len(writes)-1 || (i > 0 && i%maxBatchWriteLen == 0) {
			batch := &dynamodb.BatchWriteItemInput{
				RequestItems: map[string][]*dynamodb.WriteRequest{
					d.tableName: page,
				},
			}

			if _, err := d.client.BatchWriteItem(batch); err != nil {
				return err
			}

			page = make([]*dynamodb.WriteRequest, 0)
		}
	}

	return nil
}

// KeyAttributeMap returns an AttributeMap containing the Persistable's keys.
func (d *Table) keyAttributeMap(p Persistable) AttributeMap {
	m := map[string]*dynamodb.AttributeValue{
		d.key.hashName(): AttributeValue(d.key.hashType(), p.HashKey()),
	}

	if d.key.hasRange() {
		m[d.key.rangeName()] = AttributeValue(d.key.rangeType(), p.RangeKey())
	}

	return m
}
