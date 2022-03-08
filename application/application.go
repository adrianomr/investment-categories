package application

import (
	"adrianorodrigues.com.br/investment-categories/framework/data/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Application struct {
}

var app = Application{}

func NewApplication() Application {
	return app
}

func (app Application) Start() {
	loadEnvVariables()
	sql.DatabaseSingleton()
}

func loadEnvVariables() {
	_, b, _, _ := runtime.Caller(0)
	projectRootPath := filepath.Join(filepath.Dir(b), "../")

	profile := os.Getenv("profile")
	if profile == "" {
		log.Default().Printf("Loading default profile")
	} else {
		profile = "-" + profile
		log.Default().Printf("Loading profile: %v", profile)
	}

	err := godotenv.Load(projectRootPath + "/.env" + profile)
	if err != nil {
		panic(err)
	}
}
