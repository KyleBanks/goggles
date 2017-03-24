package gonamo

// Options represents optional configurations for DynamoDB.
type Options struct {
	Endpoint            string
	Region              string
	DefaultProvisioning int64
}

// defaultOptions returns a DynamoOptions struct containing
// sensible defaults for all values.
func defaultOptions() *Options {
	return &Options{
		DefaultProvisioning: 1,
	}
}
