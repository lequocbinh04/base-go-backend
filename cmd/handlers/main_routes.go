package handlers

import (
	"cronbrowser/appCommon"
	"cronbrowser/middleware"
	fingerprintgin "cronbrowser/module/fingerprint/transport/gin"
	"cronbrowser/module/profile/transport/gin"
	"cronbrowser/module/user/storage"
	"cronbrowser/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"gorm.io/gorm"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	dbConn := sc.MustGet(appCommon.DBMain).(*gorm.DB)
	userStore := userstorage.NewSQLStore(dbConn)

	v1 := router.Group("/v1")
	{
		authRoutes := v1.Group("/", middleware.RequiredAuth(sc, userStore))
		{
			authRoutes.GET("/profile", ginuser.Profile(sc))

			profileRoutes := authRoutes.Group("/profiles")
			{
				profileRoutes.GET("/fingerprint", ginprofile.GenFingerprint(sc))
			}

			fingerprintRoutes := authRoutes.Group("/fingerprint")
			{
				fingerprintRoutes.GET("/random", fingerprintgin.GetRandomFingerprint(sc))
			}
		}

	}

}
