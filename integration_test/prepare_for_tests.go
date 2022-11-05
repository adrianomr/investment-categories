package integration

import (
	"adrianorodrigues.com.br/investment-categories/application"
	"adrianorodrigues.com.br/investment-categories/framework/data/sql"
	"log"
	"os"
)

type PrepareForTests interface {
	Prepare()
}

type PrepareForTestsImpl struct{}

func NewPrepareForTests() PrepareForTests {
	return PrepareForTestsImpl{}
}

func (PrepareForTestsImpl) Prepare() {
	os.Setenv("profile", "test")
	application.NewApplication().Start()
	setUp()
}

func setUp() {

	(*sql.DatabaseSingleton()).FlushInMemoryDb()

	repository := sql.NewCategoryRepository()

	_, err := repository.Save(CategoryDatabase)
	if err != nil {
		log.Fatal(err)
	}
	_, err = repository.Save(CategoryDatabase2)
	if err != nil {
		log.Fatal(err)
	}
}
