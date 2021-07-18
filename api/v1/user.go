package v1

import (
	"awesomeProject/middleware"
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	if _, isExist := model.IsExistUser(data.Username); isExist {
		code = utils.IS_EXIST
	} else {
		code = model.InsertUser(&data)
		SaveUserProfile(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func SaveUserProfile(user *model.User) {
	var userProfile model.UserProfile
	userProfile.ID = user.ID
	userProfile.Username = user.Username
	model.InsertUserProfile(&userProfile)
}

func ListUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	data, code := model.SelectUserList(pageSize, offset)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func GetUser(c *gin.Context) {

}

func RemoveUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func GetUserProfile(c *gin.Context) {
	data, code := model.SelectUserProfileById(sessions.Default(c).Get(middleware.SESSION_USER_ID_KEY).(int))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func UpdateUserProfile(c *gin.Context) {
	var data model.UserProfile
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	data, code = model.UpdateUserProfileById(sessions.Default(c).Get(middleware.SESSION_USER_ID_KEY).(int), &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}
