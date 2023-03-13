package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/component/app-context"
	"github.com/kienmatu/restaurants-go/module/restaurant/transport"
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
	appCtx := appContext.NewAppContext(db, cfg)

	r := gin.Default()
	v1 := r.Group("/v1")

	restaurantApi := v1.Group("/restaurants")

	restaurantApi.GET("", transport.GetRestaurant(appCtx))
}
