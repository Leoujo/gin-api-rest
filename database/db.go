package database


import(
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/leoujo/api-go-gin/models"
)

var (
	DB  *gorm.DB
	err error
)

// Database configuration with GORM.
func DatabaseConnection() {
	connectionString := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		 log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Student{})
}