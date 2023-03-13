package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/common"
	restaurantBiz "github.com/kienmatu/restaurants-go/module/restaurant/biz"
	restaurantStorage "github.com/kienmatu/restaurants-go/module/restaurant/storage"
	"net/http"
	"strconv"

	appContext "github.com/kienmatu/restaurants-go/component/app-context"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func GetRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewGetRestaurantBiz(store)
		res, err := biz.FindByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleSuccessResponse(res))
			return
		}
		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(res))
		return
	}
}

func CreateRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantModel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
	}
}
func UpdateRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantModel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}
}
