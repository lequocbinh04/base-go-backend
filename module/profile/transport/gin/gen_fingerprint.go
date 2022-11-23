package ginprofile

import (
	"cronbrowser/appCommon"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"net/http"
)

func GenFingerprint(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(appCommon.GenProfileFingerprint()))
	}
}
