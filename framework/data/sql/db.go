package sql

import (
	"log"

	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database interface {
	connect() (*gorm.DB, error)
	GetDb() *gorm.DB
}
type DatabaseImpl struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

var database Database

func DatabaseSingleton() Database {
	if database == nil {
		database = newDbTest()
	}
	return database
}

func (d DatabaseImpl) GetDb() *gorm.DB {
	return d.Db
}

func newDb() Database {
	return &DatabaseImpl{}
}

func newDbTest() *DatabaseImpl {
	dbInstance := &DatabaseImpl{
		DsnTest:       ":memory:",
		DbTypeTest:    "sqlite3",
		Debug:         true,
		AutoMigrateDb: true,
		Env:           "test",
	}

	connection, err := dbInstance.connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	dbInstance.Db = connection

	return dbInstance
}

func (d *DatabaseImpl) connect() (*gorm.DB, error) {
	var err error
	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&dto.CategoryDto{}, &dto.Investment{})
		d.Db.Model(dto.CategoryDto{}).AddForeignKey("category_id", "category (id)", "NOTHING", "NOTHING")
		d.Db.Model(dto.Investment{}).AddForeignKey("category_id", "category (id)", "NOTHING", "NOTHING")
	}

	return d.Db, nil
}
