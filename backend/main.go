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
	// router.Use(cors.Default())

	router.GET("/about.json", servicesApi.AboutJson)
	router.POST("/mobile/token", githubApi.StoreMobileToken)

	apiRoutes := router.Group("/api")
	{

		user := apiRoutes.Group("/user", middlewares.Authorization())
		{
			user.GET("services", userApi.GetServices)
			user.GET("workflows", userApi.GetWorkflows)
			user.PUT("service/logout", userApi.LogoutService)
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
			workflow.DELETE("", workflowApi.DeleteWorkflow)
			workflow.GET("/reaction", workflowApi.GetMostRecentReaction)
		}

		action := apiRoutes.Group("/action", middlewares.Authorization())
		{
			action.POST("", actionApi.CreateAction)
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
	// Services
	jwtService                  services.JWTService                  = services.NewJWTService()
	serviceToken                services.TokenService                = services.NewTokenService(tokenRepository, userService)
	userService                 services.UserService                 = services.NewUserService(userRepository, jwtService)
	reactionResponseDataService services.ReactionResponseDataService = services.NewReactionResponseDataService(reactionResponseDataRepository)
	githubService               services.GithubService               = services.NewGithubService(githubRepository, tokenRepository, workflowsRepository, reactionRepository, reactionResponseDataService, userService)
	servicesService             services.ServicesService             = services.NewServicesService(servicesRepository, githubService)
	actionService               services.ActionService               = services.NewActionService(actionRepository, servicesService, userService)
	reactionService             services.ReactionService             = services.NewReactionService(reactionRepository, servicesService)
	workflowsService            services.WorkflowService             = services.NewWorkflowService(workflowsRepository, userService, actionService, reactionService, servicesService, serviceToken, reactionResponseDataService)

	// Controllers
	userController     controllers.UserController     = controllers.NewUserController(userService, jwtService, servicesService, reactionService, actionService, serviceToken)
	githubController   controllers.GithubController   = controllers.NewGithubController(githubService, userService, serviceToken, servicesService)
	actionController   controllers.ActionController   = controllers.NewActionController(actionService)
	servicesController controllers.ServicesController = controllers.NewServiceController(servicesService, actionService, reactionService)
	workflowController controllers.WorkflowController = controllers.NewWorkflowController(workflowsService, reactionService, actionService)
)

var (
	userApi     *api.UserApi     = api.NewUserApi(userController)
	githubApi   *api.GithubApi   = api.NewGithubApi(githubController)
	servicesApi *api.ServicesApi = api.NewServicesApi(servicesController, workflowController)
	workflowApi *api.WorkflowApi = api.NewWorkflowApi(workflowController)
	actionApi   *api.ActionApi   = api.NewActionApi(actionController)
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
