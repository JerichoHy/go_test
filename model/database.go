package model

import (
	"awesomeProject/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDatabase() {
	databaseConfig := utils.C.Database
	dsn := fmt.Sprintf("%s:%s@(%s%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err database: %s \n", err)
	}

	db.AutoMigrate(&User{}, &Activity{}, &UserProfile{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("err database: %s \n", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
