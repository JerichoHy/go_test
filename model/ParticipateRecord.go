package model

import (
	"awesomeProject/utils"
	"gorm.io/gorm"
)

type ParticipateRecord struct {
	gorm.Model
	UserId     int `gorm:"type:int; not null; uniqueIndex:user_id_activity_id_idx" json:"user_id"`
	ActivityId int `gorm:"type:int; not null; uniqueIndex:user_id_activity_id_idx" json:"activity_id"`
}

func InsertParticipateRecord(data *ParticipateRecord) int {
	if err := db.Create(&data).Error; err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeleteParticipateRecord(data *ParticipateRecord) (ParticipateRecord, int) {
	var participateRecord ParticipateRecord
	err := db.Where("user_id = ? AND activity_id = ?", data.UserId, data.ActivityId).Delete(&participateRecord).Error
	if err != nil {
		return participateRecord, utils.ERROR
	}
	return participateRecord, utils.SUCCESS
}
