package main

import (
	"adrianorodrigues.com.br/investment-categories/application"
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest"
)

// função principal
func main() {
	application.NewApplication().Start()
	rest.HttpServerSingleton().Init()

}
