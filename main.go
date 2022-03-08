package main

import (
	"github.com/OmarAouini/golang-api-starter/database"
	docs "github.com/OmarAouini/golang-api-starter/docs"
	"github.com/OmarAouini/golang-api-starter/handlers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger Sample API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
func main() {
	database.ConnectDb()
	database.Migrate()
	//router and middleware config
	r := gin.Default()
	//routes register
	handlers.CompanyRoutes(r)
	//swagger endpoint
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	//run
	r.Run("0.0.0.0:8080")
}
