package gonamo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

var (
	tableCount int

	testOpts = &Options{
		Endpoint:            "http://localhost:8999",
		Region:              "us-east-1",
		DefaultProvisioning: 10,
	}
)

func TestNewTable(t *testing.T) {
	tests := []struct {
		opts *Options

		expectedRegion       string
		expectedEndpoint     string
		expectedProvisioning int64

		keys keyProvider
	}{
		{
			nil,
			"",
			"",
			1,
			HashKeyDefinition{"hashName", StringType},
		},
		{
			testOpts,
			testOpts.Region,
			testOpts.Endpoint,
			testOpts.DefaultProvisioning,
			HashRangeKeyDefinition{"hashName", StringType, "rangeName", NumberType},
		},
	}

	for _, tt := range tests {
		// Clear the AWS_REGION environment variable to ensure we're only
		// testing the custom value.
		os.Setenv("AWS_REGION", "")

		tblName := newTableName()
		tbl, err := NewTable(tblName, tt.keys, tt.opts)
		if err != nil {
			t.Fatal(err)
		}

		if tbl.tableName != tblName {
			t.Fatalf("Unexpected table name, expected=%v, got=%v", tblName, tbl.tableName)
		} else if tbl.client == nil {
			t.Fatal("Expected client not to be null")
		}

		if *tbl.client.Config.Region != tt.expectedRegion {
			t.Fatalf("Unexpected region, expected=%v, got=%v", tt.expectedRegion, *tbl.client.Config.Region)
		} else if tbl.client.Config.Endpoint == nil && len(tt.expectedEndpoint) > 0 {
			t.Fatalf("Unexpected nil endpoint, expected=%v", tt.expectedEndpoint)
		} else if len(tt.expectedEndpoint) > 0 && *tbl.client.Config.Endpoint != tt.expectedEndpoint {
			t.Fatalf("Unexpected endpoint, expected=%v, got=%v", tt.expectedEndpoint, *tbl.client.Config.Endpoint)
		} else if tbl.defaultProvisioning != tt.expectedProvisioning {
			t.Fatalf("Unexpected defaultProvisioning, expected=%v, got=%v", tt.expectedProvisioning, tbl.defaultProvisioning)
		}

		testHashKey(t, tbl, tt.keys)
		testRangeKey(t, tbl, tt.keys)
	}
}

func TestCreateTableIfNecessary(t *testing.T) {
	tests := []struct {
		key keyProvider
	}{
		{HashKeyDefinition{"hashName", StringType}},
		{HashRangeKeyDefinition{"hashName", NumberType, "rangeName", NumberType}},
	}

	for _, tt := range tests {
		tbl, err := NewTable(newTableName(), tt.key, testOpts)
		if err != nil {
			t.Fatal(err)
		}

		// Should create the table the first time...
		if created, err := tbl.CreateTableIfNecessary(); err != nil {
			t.Fatal(err)
		} else if !created {
			t.Fatal("Expected table to be created on first call")
		}

		// ... and not the second.
		if created, err := tbl.CreateTableIfNecessary(); err != nil {
			t.Fatal(err)
		} else if created {
			t.Fatal("Expected table not to be created on second call")
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		key keyProvider
	}{
		{HashKeyDefinition{"hashName", StringType}},
		{HashRangeKeyDefinition{"hashName", NumberType, "rangeName", NumberType}},
	}

	for _, tt := range tests {
		tbl, err := NewTable(newTableName(), tt.key, testOpts)
		if err != nil {
			t.Fatal(err)
		}

		// Should return the "TableNotFound" error...
		if _, err := tbl.Describe(); err == nil {
			t.Fatal("Expected error, got nil")
		} else if err.(awserr.Error).Code() != errTableDoesntExist {
			t.Fatalf("Unexpected error, expected=%v, got=%v", errTableDoesntExist, err)
		}

		if _, err := tbl.CreateTableIfNecessary(); err != nil {
			t.Fatal(err)
		}

		// ... but not after the table is created.
		res, err := tbl.Describe()
		if err != nil {
			t.Fatal(err)
		}

		if *res.Table.AttributeDefinitions[0].AttributeName != tt.key.hashName() {
			t.Fatalf("Unexepected Hash AttributeName, expected=%v, got=%v", tt.key.hashName(), *res.Table.AttributeDefinitions[0].AttributeName)
		} else if *res.Table.AttributeDefinitions[0].AttributeType != string(tt.key.hashType()) {
			t.Fatalf("Unexepected Hash AttributeType, expected=%v, got=%v", tt.key.hashType(), *res.Table.AttributeDefinitions[0].AttributeType)
		}

		if !tt.key.hasRange() {
			if len(res.Table.AttributeDefinitions) > 1 {
				t.Fatalf("Unexpected Range key when Range was not provided. got=%v", res.Table.AttributeDefinitions[1])
			}
		} else {
			if *res.Table.AttributeDefinitions[1].AttributeName != tt.key.rangeName() {
				t.Fatalf("Unexepected Range AttributeName, expected=%v, got=%v", tt.key.rangeName(), *res.Table.AttributeDefinitions[1].AttributeName)
			} else if *res.Table.AttributeDefinitions[1].AttributeType != string(tt.key.rangeType()) {
				t.Fatalf("Unexepected Range AttributeType, expected=%v, got=%v", tt.key.rangeType(), *res.Table.AttributeDefinitions[1].AttributeType)
			}
		}
	}
}

func newTableName() string {
	tableCount++
	return fmt.Sprintf("test-%v-%v", time.Now().UnixNano(), tableCount)
}

func testHashKey(t *testing.T, tbl *Table, k keyProvider) {
	if tbl.key.hashName() != k.hashName() {
		t.Fatalf("Unexpected hashKey, expected=%v, got=%v", k.hashName(), tbl.key.hashName())
	} else if tbl.key.hashType() != k.hashType() {
		t.Fatalf("Unexpected hashType, expected=%v, got=%v", k.hashType(), tbl.key.hashType())
	}
}

func testRangeKey(t *testing.T, tbl *Table, k keyProvider) {
	if tbl.key.rangeName() != k.rangeName() {
		t.Fatalf("Unexpected rangeKey, expected=%v, got=%v", k.rangeName(), tbl.key.rangeName())
	} else if tbl.key.rangeType() != k.rangeType() {
		t.Fatalf("Unexpected rangeType, expected=%v, got=%v", k.rangeType(), tbl.key.rangeType())
	} else if tbl.key.hasRange() != k.hasRange() {
		t.Fatalf("Unexpected hasRange, expected=%v, got=%v", k.hasRange(), tbl.key.hasRange())
	}
}
