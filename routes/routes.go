package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leoujo/api-go-gin/controllers"
)

func HandleRequest() {

	// Using Gin default routes
	r := gin.Default()

	// Get all students, get student by id or by cpf.
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)


	// Create a student.
	r.POST("/students", controllers.CreateNewStudent)

	// Delete or updates a student, by it's ID.
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)

	// Run the routes.
	r.Run()
}
