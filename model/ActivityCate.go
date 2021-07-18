package model

import (
	"awesomeProject/utils"
	"gorm.io/gorm"
)

type ActivityCate struct {
	gorm.Model
	Title string `gorm:"type:varchar(50); not null" json:"title""`
}

func InsertActivityCate(data *ActivityCate) int {
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func SelectActivityCateList(pageSize int, offset int) ([]ActivityCate, int) {
	var activityCate []ActivityCate
	if err := db.Limit(pageSize).Offset(offset).Find(&activityCate).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR
	}
	return activityCate, utils.SUCCESS
}

func SelectActivityCateById(id int) (ActivityCate, int) {
	var activityCate ActivityCate
	err := db.Limit(1).Where("id = ?", id).Find(&activityCate).Error
	if err != nil {
		return activityCate, utils.ERROR
	}
	return activityCate, utils.SUCCESS
}

func DeleteActivityCateById(id int) (ActivityCate, int) {
	var activityCate ActivityCate
	err := db.Where("id = ? ", id).Delete(&activityCate).Error
	if err != nil {
		return activityCate, utils.ERROR
	}
	return activityCate, utils.SUCCESS
}

func UpdateActivityCateById(id int, data *ActivityCate) (ActivityCate, int) {
	var activityCate ActivityCate
	err := db.Model(&activityCate).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return activityCate, utils.ERROR
	}
	return activityCate, utils.SUCCESS
}
