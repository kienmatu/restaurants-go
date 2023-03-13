package utils

import (
	"fmt"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection(env string, connString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if env == "DEV" {
		err := db.AutoMigrate(&restaurantModel.Restaurant{})
		if err != nil {
			panic("Error when run migrations")
		}
	}
	return db
}
