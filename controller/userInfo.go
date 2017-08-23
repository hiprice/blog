package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hiprice/blog/httputil"
	"github.com/hiprice/blog/model"
	//"strconv"
	"gopkg.in/mgo.v2"
)

type UserInfo struct {
}

func NewUserInfo() UserInfo {
	return UserInfo{}
}
func (u *UserInfo) SaveUserInfo(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	userInfo := model.UserInfo{}

	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(400, httputil.NewErrorResponse(err))
		return
	}
	err := db.C(model.CollectionUserInfo).Insert(userInfo)

	if err != nil {
		c.JSON(400, httputil.NewErrorResponse(err))
		return
	}
	c.JSON(200, httputil.NewSuccessResponse("ok"))
	//	c.Redirect(http.StatusMovedPermanently, "/articles")
}

/*
func (u *UserInfo) GetUserInfoByID(c *gin.Context) {
	userId := c.Param("userId")

	if len(userId) <= 0 {
		c.JSON(400, gin.H{"status": "userId should be bigger than 0"})
		return
	}
	userInfo := model.NewUserInfo()
	userInfo.UserId = userId
	result, _ := userInfo.ReadByID()
	c.JSON(200, httputil.NewSuccessResponse(result))
}
*/

// Edit an userInfo
func Edit(c *gin.Context) {
	userId := c.Param("userId")

	if len(userId) <= 0 {
		c.JSON(400, gin.H{"status": "userId should be bigger than 0"})
		return
	}
	db := c.MustGet("db").(*mgo.Database)
	userInfo := model.UserInfo{}
	userInfo.UserId = userId
	//oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(model.CollectionUserInfo).Find(userInfo).One(&userInfo)
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, httputil.NewSuccessResponse("ok"))
}
