package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"area51/api"
	"area51/controllers"
	"area51/database"
	"area51/middlewares"
	"area51/repository"
	"area51/services"
)

func setupRouter() *gin.Engine {

	router := gin.Default()
	fullCors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	router.Use(fullCors)

	router.GET("/about.json", servicesApi.AboutJson)

	apiRoutes := router.Group("/api")
	{
		mobile := apiRoutes.Group("/mobile")
		{
			mobile.POST("/token", mobileApi.StoreMobileToken)
		}

		user := apiRoutes.Group("/user", middlewares.Authorization())
		{
			user.GET("services", userApi.GetServices)
			user.GET("workflows", userApi.GetWorkflows)
			user.PUT("service/logout", userApi.LogoutService)
			user.DELETE("account", userApi.DeleteAccount)
		}

		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", userApi.Login)
			auth.POST("/register", userApi.Register)
		}

		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func(ctx *gin.Context) {
				githubApi.RedirectToGithub(ctx, github.BasePath()+"/callback")
			})
			github.POST("/callback", func(ctx *gin.Context) {
				githubApi.HandleGithubTokenCallback(ctx, github.BasePath()+"/callback")
			})
		}
		workflow := apiRoutes.Group("/workflow", middlewares.Authorization())
		{
			workflow.POST("", workflowApi.CreateWorkflow)
			workflow.PUT("/activation", workflowApi.ActivateWorkflow)
			workflow.PUT("", workflowApi.UpdateWorkflow)
			workflow.DELETE("", workflowApi.DeleteWorkflow)
			workflow.GET("/reaction/latest/", workflowApi.GetMostRecentReaction)
			workflow.GET("/reactions", workflowApi.GetAllReactionsForAWorkflow)
		}

		spotify := apiRoutes.Group("/spotify")
		{
			spotify.GET("/auth", func(ctx *gin.Context) {
				spotifyApi.RedirectToSpotify(ctx, spotify.BasePath()+"/callback")
			})
			spotify.POST("/callback", func(ctx *gin.Context) {
				spotifyApi.HandleSpotifyTokenCallback(ctx, spotify.BasePath()+"/callback")
			})
		}

		microsoft := apiRoutes.Group("/microsoft")
		{
			microsoft.GET("/auth", func(ctx *gin.Context) {
				microsoftApi.RedirectToMicrosoft(ctx, "/callback")
			})
			microsoft.POST("/callback", func(ctx *gin.Context) {
				microsoftApi.HandleMicrosoftTokenCallback(ctx, "/callback")
			})
		}

		google := apiRoutes.Group("/google")
		{
			google.GET("/auth", func(ctx *gin.Context) {
				googleApi.RedirectToGoogle(ctx, "/callback")
			})
			google.POST("/callback", func(ctx *gin.Context) {
				googleApi.HandleGoogleTokenCallback(ctx, "/callback")
			})
		}
	}

	return router
}

var (
	// Database connection
	databaseConnection *gorm.DB = database.Connection()

	// Repositories
	userRepository                 repository.UserRepository                 = repository.NewUserRepository(databaseConnection)
	githubRepository               repository.GithubRepository               = repository.NewGithubRepository(databaseConnection)
	tokenRepository                repository.TokenRepository                = repository.NewTokenRepository(databaseConnection)
	servicesRepository             repository.ServiceRepository              = repository.NewServiceRepository(databaseConnection)
	actionRepository               repository.ActionRepository               = repository.NewActionRepository(databaseConnection)
	reactionRepository             repository.ReactionRepository             = repository.NewReactionRepository(databaseConnection)
	workflowsRepository            repository.WorkflowRepository             = repository.NewWorkflowRepository(databaseConnection)
	reactionResponseDataRepository repository.ReactionResponseDataRepository = repository.NewReactionResponseDataRepository(databaseConnection)
	spotifyRepository              repository.SpotifyRepository              = repository.NewSpotifyRepository(databaseConnection)
	googleRepository               repository.GoogleRepository               = repository.NewGoogleRepository(databaseConnection)

	// Services
	jwtService                  services.JWTService                  = services.NewJWTService()
	serviceToken                services.TokenService                = services.NewTokenService(tokenRepository, userService)
	userService                 services.UserService                 = services.NewUserService(userRepository, jwtService)
	reactionResponseDataService services.ReactionResponseDataService = services.NewReactionResponseDataService(reactionResponseDataRepository)
	githubService               services.GithubService               = services.NewGithubService(githubRepository, tokenRepository, workflowsRepository, reactionRepository, reactionResponseDataService, userService, servicesRepository)
	weatherService              services.WeatherService              = services.NewWeatherService(workflowsRepository, userService, reactionResponseDataService)
	servicesService             services.ServicesService             = services.NewServicesService(servicesRepository, githubService, spotifyService, googleService, microsoftService, weatherService, interpolService)
	actionService               services.ActionService               = services.NewActionService(actionRepository, servicesService, userService)
	reactionService             services.ReactionService             = services.NewReactionService(reactionRepository, servicesService)
	interpolService             services.InterpolService             = services.NewInterpolService(workflowsRepository, reactionRepository, userService, reactionResponseDataRepository)
	workflowsService            services.WorkflowService             = services.NewWorkflowService(workflowsRepository, userService, actionService, reactionService, servicesService, serviceToken, reactionResponseDataService, googleRepository, githubRepository)
	spotifyService              services.SpotifyService              = services.NewSpotifyService(userService, spotifyRepository, workflowsRepository, actionRepository, reactionRepository, tokenRepository, servicesRepository)
	googleService               services.GoogleService               = services.NewGoogleService(serviceToken, userService, workflowsRepository, servicesRepository, googleRepository)
	microsoftService            services.MicrosoftService            = services.NewMicrosoftService(serviceToken, userService, workflowsRepository, servicesRepository)

	// Controllers
	userController      controllers.UserController      = controllers.NewUserController(userService, jwtService, servicesService, reactionService, actionService, serviceToken, workflowsService, googleService, githubService)
	githubController    controllers.GithubController    = controllers.NewGithubController(githubService, userService, serviceToken, servicesService)
	servicesController  controllers.ServicesController  = controllers.NewServiceController(servicesService, actionService, reactionService)
	workflowController  controllers.WorkflowController  = controllers.NewWorkflowController(workflowsService, reactionService, actionService)
	spotifyController   controllers.SpotifyController   = controllers.NewSpotifyController(spotifyService, servicesService, userService, serviceToken)
	microsoftController controllers.MicrosoftController = controllers.NewMicrosoftController(microsoftService, userService, servicesService, serviceToken)
	googleController    controllers.GoogleController    = controllers.NewGoogleController(googleService, userService, servicesService, serviceToken)
	mobileController    controllers.MobileController    = controllers.NewMobileController(userService, serviceToken, servicesService)
)

var (
	userApi      *api.UserApi      = api.NewUserApi(userController)
	githubApi    *api.GithubApi    = api.NewGithubApi(githubController)
	servicesApi  *api.ServicesApi  = api.NewServicesApi(servicesController, workflowController)
	workflowApi  *api.WorkflowApi  = api.NewWorkflowApi(workflowController)
	spotifyApi   *api.SpotifyApi   = api.NewSpotifyApi(spotifyController)
	mobileApi    *api.MobileApi    = api.NewMobileApi(mobileController)
	microsoftApi *api.MicrosoftApi = api.NewMicrosoftApi(microsoftController)
	googleApi    *api.GoogleApi    = api.NewGoogleApi(googleController)
)

func main() {
	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
