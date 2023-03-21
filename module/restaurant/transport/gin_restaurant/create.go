package ginRestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/common"
	appContext "github.com/kienmatu/restaurants-go/component/app-context"
	restaurantBiz "github.com/kienmatu/restaurants-go/module/restaurant/biz"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
	restaurantStorage "github.com/kienmatu/restaurants-go/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantModel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&data))
	}
}
