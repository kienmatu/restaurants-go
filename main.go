package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/component/app-context"
	ginRestaurant "github.com/kienmatu/restaurants-go/module/restaurant/transport/gin_restaurant"
	"github.com/kienmatu/restaurants-go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("FOOD DELIVERY RUNNING...")
	cfg := utils.NewConfig()
	db, err := gorm.Open(mysql.Open(cfg.DatabaseConnectionURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//if cfg.ENV == "DEV" {
	//	err := db.AutoMigrate(&restaurantModel.Restaurant{})
	//	if err != nil {
	//		panic("can not migrate db")
	//	}
	//}
	db.Debug()
	appCtx := appContext.NewAppContext(db, cfg)

	r := gin.Default()
	v1 := r.Group("/v1")

	restaurantApi := v1.Group("/restaurants")

	restaurantApi.GET("", ginRestaurant.ListRestaurant(appCtx))
	restaurantApi.POST("", ginRestaurant.CreateRestaurant(appCtx))
	restaurantApi.PATCH("/:id", ginRestaurant.UpdateRestaurant(appCtx))
	restaurantApi.GET("/:id", ginRestaurant.GetRestaurant(appCtx))
	restaurantApi.DELETE("/:id", ginRestaurant.DeleteRestaurant(appCtx))

	r.Run("localhost:" + cfg.Port)
}
