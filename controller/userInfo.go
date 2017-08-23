package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hiprice/blog/httputil"
	"github.com/hiprice/blog/model"
	//"strconv"
)

type UserInfo struct {
}

func NewUserInfo() UserInfo {
	return UserInfo{}
}
func (u *UserInfo) SaveUserInfo(c *gin.Context) {
	var userInfo model.UserInfo
	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(400, httputil.NewErrorResponse(err))
		return
	}
	err := userInfo.SaveUserInfo()
	if err != nil {
		c.JSON(400, httputil.NewErrorResponse(err))
		return
	}
	c.JSON(200, httputil.NewSuccessResponse("ok"))
}

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
