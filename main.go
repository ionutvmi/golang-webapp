package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ApplicationConfig struct {
	mode string
	port string
	host string
}

var AppConfig ApplicationConfig

func init() {
	AppConfig.mode = os.Getenv("APP_MODE")
	AppConfig.port = os.Getenv("PORT")

	if AppConfig.port == "" {
		AppConfig.port = "8002"
	}
	if AppConfig.mode == "" {
		AppConfig.mode = "development"
	}

	AppConfig.host = "127.0.0.1"

	if AppConfig.mode == "production" {
		AppConfig.host = "0.0.0.0"
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		var number = rand.Intn(100)
		var msg = fmt.Sprintf("Hello there from golang %d", number)
		c.String(http.StatusOK, msg)
	})

	// start the server
	router.Run(AppConfig.host + ":" + AppConfig.port)
}
