package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alexesp/Go_FibePostgreSQL_Docker.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct{
	Db *gorm.DB

}

var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Sprintf("host=db user=%s password%s dbname=%s port=5432 sslmode=disable",
os.Getenv("DB_USER"),
os.Getenv("DB_PASSWORD"),
os.Getenv("DB_NAME"),
)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		log.Fatal("La coneccion a la base de datos fallo. \n", err)
	}

	log.Println("Conectado")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Ejecutamos migracion")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}