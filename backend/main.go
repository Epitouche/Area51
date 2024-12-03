package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/about.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "simple about.json route",
		})
	})

	// apiRoutes := router.Group("/api")
	// {
	// 	auth := apiRoutes.Group("/auth")
	// 	{
	// 		auth.POST("/login", userApi.Login)
	// 	}
	// }

	return router
}

// var (
// 	// Database connection
// 	databaseConnection *gorm.DB = database.Connection()
// 	// Repositories
// 	userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
// 	// Services
// 	userService        service.UserService        = service.NewUserService(userRepository, jwtService)
// 	// Controllers
// 	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
// )

// var (
// userApi       *api.UserApi        = api.NewUserAPI(userController)
// )


// func initDependencies(dependencies *api.Dependencies) {
// 	// Database connection
// 	databaseConnection := database.Connection()
// 	// Repositories
// 	userRepository := repository.NewUserRepository(databaseConnection)
// 	// Services
// 	userService := service.NewUserService(userRepository, jwtService)
// 	// Controllers
// 	userController := controllers.NewUserController(userService, jwtService)

// 	return &api.Dependencies{
// 		UserApi: api.NewUserApi(userController),
// 	}
// }

func main() {

	// schemas.Dependencies
	// pass the reference of the dependencies struct to the initDependencies function
	// initDependencies(&api.Dependencies{})


	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}