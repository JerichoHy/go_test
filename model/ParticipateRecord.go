package model

import (
	"gorm.io/gorm"
)

type ParticipateRecord struct {
	gorm.Model
	UserId     string `gorm:"type:varchar(50); not null" json:"user_id"`
	ActivityId string `gorm:"type:varchar(50); not null" json:"activity_id"`
}
