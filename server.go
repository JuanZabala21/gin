package main

import (
	"gin-project/controllers"
	"gin-project/middlewares"
	"gin-project/repository"
	"gin-project/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func main() {
	defer VideoRepository.CloseDB()

	var routeVideos string = "/videos"

	setupLogOutPut()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "The Application is Up!",
		})
	})

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		/* Endpoints */
		apiRoutes.GET(routeVideos, func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST(routeVideos, func(ctx *gin.Context) {
			if err := videoController.Save(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video is Saved!"})
			}
		})
		apiRoutes.PUT(routeVideos+"/:id", func(ctx *gin.Context) {
			if err := videoController.Update(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video is Updated!"})
			}
		})
		apiRoutes.DELETE(routeVideos+"/:id", func(ctx *gin.Context) {
			if err := videoController.Delete(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video is Deleted!"})
			}
		})
		/************/
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET(routeVideos, videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
