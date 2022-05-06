package main

import (
	"api-rest-go-with-docker/models"
	"api-rest-go-with-docker/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Carlos", Historia: "Historia 1"},
		{Nome: "Maria", Historia: "Historia 2"},
	}
	fmt.Println("Servidor iniciado")
	routes.HandleRequest()
}
