package v1

import (
	"awesomeProject/middleware"
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	if userId, isExist := model.IsExistUser(data.Username); !isExist {
		code = utils.IS_NOT_EXIST
	} else if !model.CheckUser(data.Username, data.Password) {
		code = utils.WRONG_PASSWORD
	} else {
		code = utils.SUCCESS
		session := sessions.Default(c)
		session.Set(middleware.SESSION_USER_ID_KEY, userId)
		session.Set(middleware.SESSION_USER_ROLE_KEY, middleware.USER_ROLE_USER)
		session.Save()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func LoginAdminHandler(c *gin.Context) {
	var data model.Administrator
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	if userId, isExist := model.IsExistAdministrator(data.Username); !isExist {
		code = utils.IS_NOT_EXIST
	} else if !model.CheckAdministrator(data.Username, data.Password) {
		code = utils.WRONG_PASSWORD
	} else {
		code = utils.SUCCESS
		session := sessions.Default(c)
		session.Set(middleware.SESSION_USER_ID_KEY, userId)
		session.Set(middleware.SESSION_USER_ROLE_KEY, middleware.USER_ROLE_ADMINSTRATOR)
		session.Save()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}
