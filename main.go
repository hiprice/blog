package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiprice/blog/controller"
	"github.com/hiprice/blog/env"
	db "github.com/hiprice/blog/helper/db"
	"github.com/hiprice/blog/middleware"
	"github.com/itsjamie/gin-cors"
	"gopkg.in/mgo.v2"
)

const (
	APISERVER          = ":3001"
	DATABASESERVER     = "localhost:27017"
	DATABASENAME       = ""
	DATABASECOLLECTION = ""
)

func init() {
	mongoSession, err := mgo.Dial(DATABASESERVER)
	if err != nil {
		panic(err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
	db.MongoSession = mongoSession
	db.UserInfoCollection = db.MongoSession.DB(DATABASENAME).C(DATABASECOLLECTION)
}

func main() {
	defer db.MongoSession.Close()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "yes")
	})

	r.Use(cors.Middleware(env.CorsConfig))
	r.Use(gin.Logger())

	user := r.Group("/api")

	uctr := controller.NewUser()
	user.POST("/signup", uctr.SignUp)
	user.POST("/signin", middleware.AuthMiddleware.LoginHandler)

	blog := r.Group("/api")
	blog.Use(middleware.AuthMiddleware.MiddlewareFunc("user"))

	bctr := controller.NewBlog()
	blog.GET("/blog", bctr.GetAll)
	blog.GET("/blog/:id", bctr.GetByID)
	blog.POST("/blog", bctr.Create)
	blog.PUT("/blog/:id", bctr.Update)
	blog.DELETE("/blog", bctr.Delete)

	r.Run(APISERVER)
}
