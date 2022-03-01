package main

import "adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest"

// função principal
func main() {
	rest.HttpServerSingleton().Init()

}
