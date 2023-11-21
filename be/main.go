package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// initilaize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	r := gin.Default()

	// connect to database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))
}
