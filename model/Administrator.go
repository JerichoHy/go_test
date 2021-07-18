package model

import (
	"awesomeProject/utils"
	"fmt"
	"gorm.io/gorm"
)

type Administrator struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null" json:"username"`
	Password string `gorm:"type:varchar(500); not null" json:"password"`
}

func InsertAdministrator(data *Administrator) int {
	data.Password = utils.GeneratePassword(data.Password)
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func IsExistAdministrator(username string) (int, bool) {
	var administrator Administrator
	rowsAffected := db.Select("id").Where("username = ?", username).First(&administrator).RowsAffected
	if rowsAffected > 0 {
		return int(administrator.ID), true
	}
	return 0, false
}

func CheckAdministrator(username string, password string) bool {
	var administrator Administrator
	err := db.Where("username = ?", username).First(&administrator).Error
	if err != nil {
		fmt.Printf("err check: %s \n", err)
		return false
	}
	return utils.ValidatePassword(administrator.Password, password)
}
