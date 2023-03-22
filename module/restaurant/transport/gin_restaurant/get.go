package ginRestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/common"
	appContext "github.com/kienmatu/restaurants-go/component/app-context"
	restaurantBiz "github.com/kienmatu/restaurants-go/module/restaurant/biz"
	restaurantStorage "github.com/kienmatu/restaurants-go/module/restaurant/storage"
	"net/http"
)

func GetRestaurant(appCtx appContext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := common.UIDFromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewGetRestaurantBiz(store)
		res, err := biz.FindByCondition(c.Request.Context(), map[string]interface{}{"id": id.GetSequence()})
		if err != nil {
			panic(err)
		}
		res.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
