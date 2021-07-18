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

func SaveActivity(c *gin.Context) {
	var data model.Activity
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	code = model.InsertActivity(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func ListActivity(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	data, code := model.SelectActivityList(pageSize, offset)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func GetActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.SelectActivityById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func RemoveActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.DeleteActivityById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func UpdateActivity(c *gin.Context) {
	var data model.Activity
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	data, code = model.UpdateActivityById(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func ParticipateActivity(c *gin.Context) {
	var data model.ParticipateRecord
	data.UserId = sessions.Default(c).Get(middleware.SESSION_USER_ID_KEY).(int)
	data.ActivityId, _ = strconv.Atoi(c.Param("id"))
	code := model.InsertParticipateRecord(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func UnParticipateActivity(c *gin.Context) {
	var data model.ParticipateRecord
	data.UserId = sessions.Default(c).Get(middleware.SESSION_USER_ID_KEY).(int)
	data.ActivityId, _ = strconv.Atoi(c.Param("id"))
	participateRecord, code := model.DeleteParticipateRecord(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   participateRecord,
		"msg":    utils.GetErrMsg(code),
	})
}

func GetParticipateActivityListByUserId(c *gin.Context) {
	userId := sessions.Default(c).Get(middleware.SESSION_USER_ID_KEY).(int)
	data, code := model.SelectParticipateActivityListByUserId(userId)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}
