package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kienmatu/restaurants-go/common"
	appContext "github.com/kienmatu/restaurants-go/component/app-context"
)

func Recover(ctx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}
				if err, ok := err.(error); ok {
					appErr := common.ErrInternal(err.(error))
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}
				appErr := common.ErrInternal(fmt.Errorf("panic: %s", err))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return
			}
		}()
		c.Next()
	}
}
