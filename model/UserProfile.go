package model

import (
	"awesomeProject/utils"
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null" json:"username"`
	Email    string `gorm:"type:varchar(50); not null" json:"email"`
	Portrait string `gorm:"type:varchar(50); not null" json:"portrait"`
}

func SelectUserProfileById(data int) (UserProfile, int) {
	var userProfile UserProfile
	err := db.Limit(1).Where("id = ?", data).Find(&userProfile).Error
	if err != nil {
		return userProfile, utils.ERROR
	}
	return userProfile, utils.SUCCESS
}

func InsertUserProfile(data *UserProfile) int {
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func UpdateUserProfileById(id int, data *UserProfile) (UserProfile, int) {
	var userProfile UserProfile
	err := db.Model(&userProfile).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return userProfile, utils.ERROR
	}
	return userProfile, utils.SUCCESS
}
