package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pedrolessa-dev/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=db_gin_api_rest port=5432 sslmode=disable", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
