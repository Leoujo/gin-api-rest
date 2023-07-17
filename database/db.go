package database


import(
	"log"
	"os"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/leoujo/api-go-gin/models"
	"github.com/joho/godotenv"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	// Retrieving env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded.")
	}
	
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	// Database configuration with GORM.
	connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=5432 sslmode=disable", user, password)
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		 log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Student{})
}