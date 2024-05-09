package main

import (
	"customer-service/db"
	"customer-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	secret := db.GetSecretValue()
	db := db.GetDB(secret)

	a := service.GetApp(db)

	r := gin.Default()

	r.POST("/customers", a.PostHandler)
	r.GET("/customers/:customerId", func(c *gin.Context) {})
	r.PUT("/customers/:customerId", func(c *gin.Context) {})
	r.DELETE("/customers/:customerId", func(c *gin.Context) {})

	r.Run("localhost:8080")
}
