package fingerprintgin

import (
	"cronbrowser/appCommon"
	fingerprintbiz "cronbrowser/module/fingerprint/biz"
	fingerprintmodel "cronbrowser/module/fingerprint/model"
	fingerprintstorage "cronbrowser/module/fingerprint/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"gorm.io/gorm"
	"net/http"
)

func GetRandomFingerprint(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter fingerprintmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*gorm.DB)
		store := fingerprintstorage.NewSQLStore(db)
		biz := fingerprintbiz.NewFindRandomFingerprintBiz(store)
		data, err := biz.FindRandomFingerprint(c, &filter)
		if err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data))
	}
}
