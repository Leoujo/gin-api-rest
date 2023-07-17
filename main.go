package main

import (
	"github.com/leoujo/api-go-gin/database"
	"github.com/leoujo/api-go-gin/routes"
)

// Application entry point
// Call the database and routes.
func main() {
	database.DatabaseConnection()
	routes.HandleRequest()
}