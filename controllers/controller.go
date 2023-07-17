package controllers

import (
	"github.com/leoujo/api-go-gin/database"
	"github.com/leoujo/api-go-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Retrieve all students.
func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// Create a new student
func CreateNewStudent(c *gin.Context) {
	var student models.Student

	// should bind json -> Get the request body and an package this based on the struct.
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// Get student by id
func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	// If no student is found
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student was not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// Delete student by id
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted successfully!"})
}

// Edit student by id
func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Find(&student, id)

	// If some error happens
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

// Get student by cpf
func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	// If no student is found
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student was not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}