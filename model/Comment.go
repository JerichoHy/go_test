package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId     string `gorm:"type:varchar(50); not null" json:"user_id"`
	ActivityId string `gorm:"type:varchar(50); not null" json:"activity_id"`
	Content    string `gorm:"type:varchar(500); not null" json:"content"`
}
