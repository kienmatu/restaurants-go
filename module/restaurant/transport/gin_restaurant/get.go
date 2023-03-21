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

func GetRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewGetRestaurantBiz(store)
		res, err := biz.FindByCondition(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
