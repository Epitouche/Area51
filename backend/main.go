package main

import (
	"area51/api"
	"area51/controllers"
	"area51/database"
	"area51/repository"
	"area51/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/about.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "simple about.json route",
		})
	})

	var databaseConnection *gorm.DB = database.Connection()

	userRepository := repository.NewUserRepository(databaseConnection)
	jwtService := services.NewJWTService()
	userService := services.NewUserService(userRepository, jwtService)
	userController := controllers.NewUserController(userService, jwtService)
	userApi := api.NewUserAPI(userController)

	githubTokenRepository := repository.NewGithubTokenRepository(databaseConnection)
	githubTokenService := services.NewGithubTokenService(githubTokenRepository)
	githubTokenController := controllers.NewGithubTokenController(githubTokenService, userService)
	githubApi := api.NewGithubAPI(githubTokenController)

	apiRoutes := router.Group("/")
	{
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/register", userApi.Register)
		}

		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func (ctx *gin.Context) {
				githubApi.RedirectToGithub(ctx, github.BasePath()+"/auth/callback")
			})
		}
	}

	return router
}

func main() {
	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}