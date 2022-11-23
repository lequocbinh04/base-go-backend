package main

import (
	"cronbrowser/appCommon"
	"cronbrowser/plugin/tokenprovider"
	"cronbrowser/plugin/tokenprovider/paseto"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/logger"
	"time"
)

func main() {
	service, _ := goservice.New(
		goservice.WithInitRunnable(paseto.NewPasetoProvider(appCommon.PasetoProvider)),
	)
	service.Init()

	pro := service.MustGet(appCommon.PasetoProvider).(tokenprovider.Provider)
	token, err := pro.Generate(appCommon.TokenPayload{
		UId:   1,
		URole: "admin",
	}, 60*60*24*30*12*10)
	if err != nil {
		panic(err)
	}
	logger := logger.GetCurrent().GetLogger("token")
	logger.Infoln(token.GetToken())

	time.Sleep(time.Second * 2)

	logger.Infoln(pro.Validate(token.GetToken()))
	service.Start()
}
