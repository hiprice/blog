package controller

//推荐用户
import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/blog/httputil"
	"github.com/makki0205/blog/model"
)

type RecommendUser struct {
	recommendUserRep model.RecommendUserRepository
	userRep          model.UserRepository
}

func NewRecommendUser() RecommendUser {
	return RecommendUser{}
}

func (r *RecommendUser) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"hoge": "hoge",
	})
}

func (r *RecommendUser) Create(c *gin.Context) {
	//请求处理
	var recommendUser model.RecommendUser
	err := c.BindJSON(&recommendUser)
	if err != nil {
		panic(err)
	}
	// Insert
	ok, err := r.recommendUserRep.Create(recommendUser)
	if err != nil {
		panic(err)
	}
	c.JSON(200, httputil.NewSuccessResponse(ok))
}
