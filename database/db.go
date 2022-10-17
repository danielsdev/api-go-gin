package database

import (
	"log"

	"github.com/danielsdev/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	//Cria a tabela com base na struct de Aluno
	DB.AutoMigrate(&models.Aluno{})
}
