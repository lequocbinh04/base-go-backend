package cmd

import (
	"cronbrowser/appCommon"
	"cronbrowser/cmd/handlers"
	"cronbrowser/middleware"
	"cronbrowser/plugin/appredis"
	"cronbrowser/plugin/locker"
	appgrpc "cronbrowser/plugin/remotecall/grpc"
	"cronbrowser/plugin/tokenprovider/paseto"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/storage/sdkgorm"
	"github.com/spf13/cobra"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("cron-profile"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", appCommon.DBMain)),
		goservice.WithInitRunnable(paseto.NewPasetoProvider(appCommon.PasetoProvider)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", appCommon.PluginRedis)),
		goservice.WithInitRunnable(appgrpc.NewGRPCServer(appCommon.PluginGrpcServer)),
	)
	if err := service.Init(); err != nil {
		panic(err)
	}
	service.Add(
		goservice.WithInitRunnable(locker.NewLocker(appCommon.PluginLocker, service)),
	)

	if err := service.InitPrefix(appCommon.PluginLocker); err != nil {
		panic(err)
	}

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start a CronBrowser service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())
			engine.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})

			handlers.MainRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
