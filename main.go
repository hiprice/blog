package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiprice/blog/controller"
	"github.com/hiprice/blog/env"
	db "github.com/hiprice/blog/helper/db"
	"github.com/hiprice/blog/middleware"
	"github.com/itsjamie/gin-cors"
)

const (
	APISERVER = ":3001"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(middleware.Connect)
	router.Use(middleware.ErrorHandler)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "yes")
	})

	router.Use(cors.Middleware(env.CorsConfig))
	router.Use(gin.Logger())

	user := router.Group("/api")

	uctr := controller.NewUser()
	user.POST("/signup", uctr.SignUp)
	user.POST("/signin", middleware.AuthMiddleware.LoginHandler)

	blog := router.Group("/api")
	blog.Use(middleware.AuthMiddleware.MiddlewareFunc("user"))

	bctr := controller.NewBlog()
	blog.GET("/blog", bctr.GetAll)
	blog.GET("/blog/:id", bctr.GetByID)
	blog.POST("/blog", bctr.Create)
	blog.PUT("/blog/:id", bctr.Update)
	blog.DELETE("/blog", bctr.Delete)

	userapi := router.Group("/api")
	userapi.Use(middleware.AuthMiddleware.MiddlewareFunc("user"))
	uictr := controller.NewUserInfo()
	userapi.POST("/user", uictr.SaveUserInfo)
	//userapi.GET("/user", uictr.GetUserInfoByID)

	router.Run(APISERVER)
}
