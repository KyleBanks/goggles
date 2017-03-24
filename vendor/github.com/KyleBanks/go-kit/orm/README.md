# orm
--
    import "github.com/KyleBanks/go-kit/orm/"

Package orm manages access to a database, including ORM-like functionality.

The package wraps the GORM library, which can then be potentially swapped out
with minimal changes.

## Usage

```go
var (
	// ErrRecordNotFound is the error returned when trying to load a record
	// that cannot be found.
	ErrRecordNotFound = gorm.ErrRecordNotFound
)
```

#### type Model

```go
type Model struct {
	gorm.Model
}
```

Model is a type added to domain models that will provide ORM functionality.

#### type ORM

```go
type ORM struct {
}
```

ORM is a container for the underlying database connection.

#### func (ORM) AutoMigrate

```go
func (orm ORM) AutoMigrate(models []interface{}) error
```
AutoMigrate performs database migration for all Model types provided.

#### func (ORM) Begin

```go
func (orm ORM) Begin() *gorm.DB
```
Begin starts a new database transaction.

#### func (ORM) Create

```go
func (orm ORM) Create(model interface{}) *gorm.DB
```
Create inserts a new model instance into the database.

#### func (ORM) Exec

```go
func (orm ORM) Exec(query string, output interface{}) *gorm.DB
```
Exec performs a raw SQL query against the underlying database.

#### func (ORM) First

```go
func (orm ORM) First(model interface{}, where ...interface{}) *gorm.DB
```
First returns the first model (ordered by ID) that matches the specified query.

#### func (ORM) Last

```go
func (orm ORM) Last(model interface{}, where ...interface{}) *gorm.DB
```
Last returns the last model (ordered by ID) that matches the specified query.

#### func (ORM) Model

```go
func (orm ORM) Model(model interface{}) *gorm.DB
```
Model specifies the domain model that subsequent queries will be run against.

#### func (ORM) ModelExistsWithID

```go
func (orm ORM) ModelExistsWithID(model interface{}, id uint) (bool, error)
```
ModelExistsWithID returns a boolean indicating if an instance of the specified
model exists with a given ID.

#### func (ORM) ModelWithID

```go
func (orm ORM) ModelWithID(model interface{}, id uint) error
```
ModelWithID returns an instance of the specified model with the given ID.

#### func (*ORM) Open

```go
func (orm *ORM) Open(dialect, connectionString string) (*gorm.DB, error)
```
Open creates a database connection, or returns an existing one if present.

#### func (ORM) Save

```go
func (orm ORM) Save(value interface{}) *gorm.DB
```
Save updates a model with the given attributes.

#### func (ORM) Where

```go
func (orm ORM) Where(query interface{}, args ...interface{}) *gorm.DB
```
Where performs a query with "Where" parameters.
