package v1

import (
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAdministrator(c *gin.Context) {
	var data model.Administrator
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("err bind: %s \n", err)
		return
	}
	var code int
	if _, isExist := model.IsExistAdministrator(data.Username); isExist {
		code = utils.IS_EXIST
	} else {
		code = model.InsertAdministrator(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    utils.GetErrMsg(code),
	})
}
