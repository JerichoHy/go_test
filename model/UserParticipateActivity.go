package model

import "awesomeProject/utils"

func SelectParticipateActivityListByUserId(userId int) ([]Activity, int) {
	var activities []Activity
	err := db.Table("participate_records").Select("activities.*").Joins("left join activities on activities.id = participate_records.activity_id where participate_records.user_id = ?", userId).Scan(&activities).Error
	if err != nil {
		return nil, utils.ERROR
	}
	return activities, utils.SUCCESS
}
