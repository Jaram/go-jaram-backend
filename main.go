package main

import (
	together "go-jaram-backend/jaram-together"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Main Function
func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	urlTogether := r.Group("/together")
	{
		urlTogether.POST("/msgRequest", together.MsgRequest)
	}

	log.Fatal(http.ListenAndServe(":8080", r))
}
