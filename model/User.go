package model

import (
	"awesomeProject/utils"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null" json:"username"`
	Password string `gorm:"type:varchar(500); not null" json:"password"`
}

func InsertUser(data *User) int {
	data.Password = utils.GeneratePassword(data.Password)
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func SelectUserByUsername(data string) (User, int) {
	var user User
	err := db.Limit(1).Where("username = ?", data).Find(&user).Error
	if err != nil {
		return user, utils.ERROR
	}
	return user, utils.SUCCESS
}

func IsExistUser(username string) (int, bool) {
	var user User
	rowsAffected := db.Select("id").Where("username = ?", username).First(&user).RowsAffected
	if rowsAffected > 0 {
		return int(user.ID), true
	}
	return 0, false
}

func CheckUser(username string, password string) bool {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Printf("err check: %s \n", err)
		return false
	}
	return utils.ValidatePassword(user.Password, password)
}

func SelectUserList(pageSize int, offset int) ([]User, int) {
	var users []User
	if err := db.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR
	}
	return users, utils.SUCCESS
}
