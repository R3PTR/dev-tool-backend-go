package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", index)

	router.Run("localhost:8080")
}

func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{"hello", "world!"})
	test()
}
