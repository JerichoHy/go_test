package model

import (
	"awesomeProject/utils"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title              string           `gorm:"type:varchar(50); not null" json:"title""`
	StartTime          *utils.LocalTime `gorm:"type:time; not null" json:"start_time"`
	EndTime            *utils.LocalTime `gorm:"type:time; not null" json:"end_time"`
	Content            string           `gorm:"type:varchar(50); " json:"content"`
	ActivityCategoryID int              `gorm:"type:int; not null" json:"activity_category_id""`
}

func InsertActivity(data *Activity) int {
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func SelectActivityList(pageSize int, offset int) ([]Activity, int) {
	var activities []Activity
	if err := db.Limit(pageSize).Offset(offset).Find(&activities).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR
	}
	return activities, utils.SUCCESS
}

func SelectActivityById(id int) (Activity, int) {
	var activity Activity
	err := db.Limit(1).Where("id = ?", id).Find(&activity).Error
	if err != nil {
		return activity, utils.ERROR
	}
	return activity, utils.SUCCESS
}

func DeleteActivityById(id int) (Activity, int) {
	var activity Activity
	err := db.Where("id = ? ", id).Delete(&activity).Error
	if err != nil {
		return activity, utils.ERROR
	}
	return activity, utils.SUCCESS
}

func UpdateActivityById(id int, data *Activity) (Activity, int) {
	var activity Activity
	err := db.Model(&activity).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return activity, utils.ERROR
	}
	return activity, utils.SUCCESS
}
