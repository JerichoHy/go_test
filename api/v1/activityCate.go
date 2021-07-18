package v1

import (
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveActivityCate(c *gin.Context) {
	var data model.ActivityCate
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	code = model.InsertActivityCate(&data)
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func ListActivityCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	data, code := model.SelectActivityCateList(pageSize, offset)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func GetActivityCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.SelectActivityCateById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func RemoveActivityCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.DeleteActivityCateById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}

func UpdateActivityCate(c *gin.Context) {
	var data model.ActivityCate
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	data, code = model.UpdateActivityCateById(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}
