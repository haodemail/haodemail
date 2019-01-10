package main

import (
	"./controllers"
	"./middleware"
	"./models"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logFile, err := os.OpenFile("partner.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.OPTIONS("/*path", middleware.CORSMiddleware())
	router.NoRoute(controllers.Handle404)

	router.GET("/", controllers.HandleIndex)
	router.POST("/login", controllers.HandleLogin)
	router.GET("/logout", controllers.HandleLogout)

	// dashboard api
	api := router.Group("/api")
	api.Use(middleware.JWT())
	{
		// user

	}

	router.Run(models.Config.WebServer.Listen)
}
