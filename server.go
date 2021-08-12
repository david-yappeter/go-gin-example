package main

import (
	"io"
	"myapp/config"
	"myapp/controller"
	"myapp/docs"
	"myapp/middlewares"
	"myapp/migration"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const defaultPort = "8080"

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load()

	config.InitDB()

	db := config.DB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	migration.MigrateTable()

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Gin Sample Project"
	docs.SwaggerInfo.Description = "Basic Gin Project Example"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setup Logging
	f := setupLogOutput()
	defer f.Close()

	// Setup Router
	router := gin.Default()
	// router.Use(middlewares.BasicAuth())

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	apiRoutes := router.Group("/api")
	{
		apiRoutes.Use(middlewares.ContentJSON())

		apiRoutes.GET("/users/:uuID", controller.UserGetByUUID)

		apiRoutes.GET("/users", controller.UserGetAll)

		apiRoutes.POST("/users", controller.UserCreate)

		apiRoutes.POST("/login", controller.UserLogin)

		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.Use(middlewares.Auth())
			authRoutes.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "hello world")
			})
		}
	}

	// viewRoutes := router.Group("/view")
	// {
	// 	// viewRoutes.GET("/videos", ctrl.ShowAll)
	// }

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)
}

func setupLogOutput() *os.File {
	// f, _ := os.OpenFile("gin.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.FileMode(0777))
	f, _ := os.Create("gin.log")

	// 	f.WriteString(fmt.Sprintf(`
	// ###START_PROGRAM###
	// [%+v]

	// `, time.Now().UTC()))

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return f
}
