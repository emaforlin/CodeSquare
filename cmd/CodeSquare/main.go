package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// setup routes and handlers
	r.GET("/ping", ping)

	return r
}

func main() {
	router := setupRouter()

	router.Run()
}
