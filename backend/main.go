package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"area51/api"
	"area51/controllers"
	"area51/database"
	"area51/repository"
	"area51/services"
)

func setupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/about.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "simple about.json route",
		})
	})

	apiRoutes := router.Group("/api")
	{
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", userApi.Login)
			auth.POST("/register", userApi.Register)
		}

		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", )
		}
	}

	return router
}

var (
	// Database connection
	databaseConnection *gorm.DB = database.Connection()
	// Repositories
	userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
	// Services
	jwtService 		services.JWTService						= services.NewJWTService()
	userService        services.UserService        = services.NewUserService(userRepository, jwtService)
	// Controllers
	userController        controllers.UserController        = controllers.NewUserController(userService, jwtService)
)

var (
userApi       *api.UserApi        = api.NewUserApi(userController)
githubApi     *api.GithubApi      = api.NewGithubApi(githubController)
)


// func initDependencies(dependencies *api.UserDependencies) {
// 	// Database connection
// 	databaseConnection := database.Connection()
// 	// Repositories
// 	userRepository := repository.NewUserRepository(databaseConnection)
// 	// Services
// 	jwtService := services.NewJWTService()
// 	userService := services.NewUserService(userRepository, jwtService)
// 	// Controllers
// 	userController := controllers.NewUserController(userService, jwtService)

// 	dependencies.UserApi = api.NewUserApi(userController)
// }

func main() {

	// schemas.Dependencies
	// pass the reference of the dependencies struct to the initDependencies function
	// initDependencies(&api.UserDependencies{})


	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}