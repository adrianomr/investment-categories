package integration

import (
	"adrianorodrigues.com.br/investment-categories/framework/data/sql"
	"log"
)

type PrepareForTests interface {
	Prepare()
}

type PrepareForTestsImpl struct{}

func NewPrepareForTests() PrepareForTests {
	return PrepareForTestsImpl{}
}

func (PrepareForTestsImpl) Prepare() {
	setUp()
}

func setUp() {

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
