package ginRestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/common"
	appContext "github.com/kienmatu/restaurants-go/component/app-context"
	restaurantBiz "github.com/kienmatu/restaurants-go/module/restaurant/biz"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
	restaurantStorage "github.com/kienmatu/restaurants-go/module/restaurant/storage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantModel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.ID = id
		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
