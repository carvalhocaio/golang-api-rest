package main

import (
	"fmt"

	"github.com/carvalhocaio/golang-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
