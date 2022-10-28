package main

import (
	"github.com/danielsdev/api-go-gin/database"
	"github.com/danielsdev/api-go-gin/routes"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Erro ao carregar env")
	}

	database.ConectaComBancoDeDados()
	//fixtures.GerarAlunos()
	//fmt.Printf("%+v\n", models.Alunos)
	routes.HandleRequests()
}
