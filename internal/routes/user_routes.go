package routes

import (
	"net/http"

	"github.com/Kallox/CRIS-Backend/internal/models"
	"github.com/Kallox/CRIS-Backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var newUser models.User
	// Validate the input
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// RegisterUserRoutes - Register the routes for the user model
func RegisterUserRoutes(router *gin.Engine) {
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", AddUser)
}
