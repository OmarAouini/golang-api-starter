package main

import (
	"github.com/OmarAouini/golang-api-starter/database"
	"github.com/OmarAouini/golang-api-starter/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()
	database.Migrate()
	//router and middleware config
	r := gin.Default()
	//routes register
	handlers.CompanyRoutes(r)
	//run
	r.Run("0.0.0.0:8080")
}
