package main

import (
	"github.com/danielsdev/api-go-gin/database"
	"github.com/danielsdev/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	//fixtures.GerarAlunos()
	//fmt.Printf("%+v\n", models.Alunos)
	routes.HandleRequests()
}
