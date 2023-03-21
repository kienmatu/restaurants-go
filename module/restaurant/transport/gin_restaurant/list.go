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

func ListRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		var filter restaurantModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewListRestaurantBiz(store)
		var data []*restaurantModel.Restaurant

		data, err := biz.ListByCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SuccessResponse(data, &paging, &filter))
	}
}
