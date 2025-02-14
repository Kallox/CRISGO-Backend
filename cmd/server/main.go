package main

import (
	"fmt"
	"net/http"

	db "github.com/Kallox/CRIS-Backend/internal/database"
	"github.com/Kallox/CRIS-Backend/internal/models"
	"github.com/Kallox/CRIS-Backend/internal/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// getConfig - Read the server configuration file and return the values
func getConfig() map[string]string {
	var config = make(map[string]string)
	cfg, err := ini.Load("config/server.cfg")
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		return nil
	}
	config["appName"] = cfg.Section("server").Key("app_name").String()
	config["appDescription"] = cfg.Section("server").Key("app_description").String()
	config["contact"] = cfg.Section("server").Key("contact").String()
	config["version"] = cfg.Section("server").Key("version").String()

	return config
}

// setupRouter - Setup the routes for the application
func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Main route to get the status of the application
	r.GET("/", func(c *gin.Context) {
		config := getConfig()
		c.JSON(http.StatusOK, config)
	})

	// Register all routes
	routes.RegisterRoutes(r)

	return r
}

func main() {
	db.DBConnection()                // Connect to the database
	db.DB.AutoMigrate(models.User{}) // Migrate the schema
	r := setupRouter()               // Setup the routes
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
