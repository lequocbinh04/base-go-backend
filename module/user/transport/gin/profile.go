package ginuser

import (
	"cronbrowser/appCommon"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"net/http"
)

func Profile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(appCommon.CurrentUser)
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(u))
	}
}
