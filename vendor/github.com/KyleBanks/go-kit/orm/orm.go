// Package orm manages access to a database, including ORM-like functionality.
//
// The package wraps the GORM library, which can then be potentially swapped out with
// minimal changes.
package orm

import (
	"github.com/KyleBanks/go-kit/log"
	// Required to initialize the mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// Required to initialize the sqlite driver.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// ErrRecordNotFound is the error returned when trying to load a record
	// that cannot be found.
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// ORM is a container for the underlying database connection.
type ORM struct {
	conn *gorm.DB
}

// Model is a type added to domain models that will provide ORM functionality.
type Model struct {
	gorm.Model
}

// Open creates a database connection, or returns an existing one if present.
func (orm *ORM) Open(dialect, connectionString string) (*gorm.DB, error) {
	if orm.conn != nil {
		return orm.conn, nil
	}

	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		return nil, err
	}

	// Configure
	// TODO: Accept options as a param to Open
	db.SetLogger(log.Logger)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(0) // Unlimited

	orm.conn = db
	return orm.conn, nil
}

// AutoMigrate performs database migration for all Model types provided.
func (orm ORM) AutoMigrate(models []interface{}) error {
	for _, model := range models {
		if err := orm.conn.AutoMigrate(model).Error; err != nil {
			return err
		}
	}

	return nil
}

// Exec performs a raw SQL query against the underlying database.
func (orm ORM) Exec(query string, output interface{}) *gorm.DB {
	return orm.conn.Raw(query).Scan(output)
}

// Begin starts a new database transaction.
func (orm ORM) Begin() *gorm.DB {
	return orm.conn.Begin()
}

// Where performs a query with "Where" parameters.
func (orm ORM) Where(query interface{}, args ...interface{}) *gorm.DB {
	return orm.conn.Where(query, args...)
}

// Create inserts a new model instance into the database.
func (orm ORM) Create(model interface{}) *gorm.DB {
	return orm.conn.Create(model)
}

// Save updates a model with the given attributes.
func (orm ORM) Save(value interface{}) *gorm.DB {
	return orm.conn.Save(value)
}

// Model specifies the domain model that subsequent queries will be run against.
func (orm ORM) Model(model interface{}) *gorm.DB {
	return orm.conn.Model(model)
}

// First returns the first model (ordered by ID) that matches the specified query.
func (orm ORM) First(model interface{}, where ...interface{}) *gorm.DB {
	return orm.conn.First(model, where...)
}

// Last returns the last model (ordered by ID) that matches the specified query.
func (orm ORM) Last(model interface{}, where ...interface{}) *gorm.DB {
	return orm.conn.Last(model, where...)
}

// ModelWithID returns an instance of the specified model with the given ID.
func (orm ORM) ModelWithID(model interface{}, id uint) error {
	// First check if the Model exists.
	// We do this so that we can avoid an error returned by the ORM
	// when a query returns no results.
	if exists, err := orm.ModelExistsWithID(model, id); err != nil {
		return err
	} else if !exists {
		return nil
	}

	// It exists, so let's load it
	if err := orm.First(model, id).Error; err != nil {
		return err
	}

	return nil
}

// ModelExistsWithID returns a boolean indicating if an instance of the
// specified model exists with a given ID.
func (orm ORM) ModelExistsWithID(model interface{}, id uint) (bool, error) {
	var count int64

	err := orm.Model(model).Where(id).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
