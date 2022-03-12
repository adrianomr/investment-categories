package sql

import (
	"errors"
	"log"
	"os"
	"strconv"

	"adrianorodrigues.com.br/investment-categories/framework/data/sql/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database interface {
	GetDb() (*gorm.DB, error)
	FlushInMemoryDb()
}
type DatabaseImpl struct {
	db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

var database Database

func DatabaseSingleton() *Database {
	if database == nil {
		configDb()
	}
	return &database
}

func configDb() {
	debugMode, err := strconv.ParseBool(os.Getenv("database.debug"))
	dns := os.Getenv("database.dsn")
	if dns == "" {
		panic(errors.New("No database found"))
	}
	if err != nil {
		log.Printf("Using default debug mode")
		debugMode = false
	}
	autoMigrate, err := strconv.ParseBool(os.Getenv("database.autoMigrate"))
	if err != nil {
		log.Printf("Using default autoMigrate mode")
		autoMigrate = false
	}
	database = &DatabaseImpl{
		Dsn:           dns,
		DbType:        os.Getenv("database.type"),
		Debug:         debugMode,
		AutoMigrateDb: autoMigrate,
		Env:           os.Getenv("profile"),
	}
}

func (d *DatabaseImpl) GetDb() (*gorm.DB, error) {
	if d.Dsn == ":memory:" {
		if d.db == nil {
			db, err := d.connect()
			if err != nil {
				return nil, err
			}
			d.db = db
		}
		return d.db, nil
	}
	db, err := d.connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *DatabaseImpl) FlushInMemoryDb() {
	if d.db != nil {
		d.db.Close()
		db, err := d.connect()
		if err != nil {
			log.Fatal(err)
		}
		d.db = db
	}
}

func (d *DatabaseImpl) connect() (*gorm.DB, error) {
	db, err := gorm.Open(d.DbType, d.Dsn)
	if err != nil {
		return nil, err
	}

	if d.Debug {
		db.LogMode(true)
	}

	if d.AutoMigrateDb {
		db.AutoMigrate(&dto.CategoryDto{}, &dto.InvestmentDto{})
		db.Model(dto.CategoryDto{}).AddForeignKey("category_id", "category (id)", "NOTHING", "NOTHING")
		db.Model(dto.InvestmentDto{}).AddForeignKey("category_id", "category (id)", "NOTHING", "NOTHING")
	}

	return db, nil
}
