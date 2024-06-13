package main

import (
	"gin-project/api"
	"gin-project/controllers"
	"gin-project/docs"
	"gin-project/repository"
	"gin-project/services"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	VideoRepository repository.VideoRepository = repository.NewVideoRepository()
	VideoService    services.VideoService      = services.New(VideoRepository)
	loginService    services.LoginService      = services.NewLoginService()
	jwtService      services.JWTServices       = services.NewJWTService()

	videoController controllers.VideoController = controllers.New(VideoService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func setupLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	/* Swagger */
	docs.SwaggerInfo.Title = "My First API with Golang"
	docs.SwaggerInfo.Description = "My Description API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	/**********/

	defer VideoRepository.CloseDB()

	/* Server */
	server := gin.Default()
	/*server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middlewares.Logger())*/
	setupLogOutPut()
	/**********/

	/* API */
	appAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		/* Health */
		server.GET("/health", appAPI.Authenticate)

		/* Login */
		login := apiRoutes.Group("/login")
		{
			login.POST("/token", appAPI.Authenticate)
		}

		/* Video */
		videos := apiRoutes.Group("/videos")
		{
			videos.GET("", appAPI.GetVideos)
			videos.POST("", appAPI.CreateVideo)
			videos.PUT(":id", appAPI.UpdateVideo)
			videos.DELETE(":id", appAPI.DeleteVideo)
		}

		/* View */
		viewRoutes := server.Group("/view")
		{
			viewRoutes.GET("/videos", videoController.ShowAll)
		}
	}
	/********/

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
