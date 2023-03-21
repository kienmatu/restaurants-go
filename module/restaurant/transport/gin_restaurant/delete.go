package ginRestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/kienmatu/restaurants-go/common"
	appContext "github.com/kienmatu/restaurants-go/component/app-context"
	restaurantBiz "github.com/kienmatu/restaurants-go/module/restaurant/biz"
	restaurantStorage "github.com/kienmatu/restaurants-go/module/restaurant/storage"
)

func DeleteRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
