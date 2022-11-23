package middleware

import (
	"context"
	"cronbrowser/appCommon"
	"cronbrowser/module/user/model"
	"cronbrowser/plugin/tokenprovider"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *appCommon.AppError {
	return appCommon.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(sc goservice.ServiceContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := sc.MustGet(appCommon.PasetoProvider).(tokenprovider.Provider)

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			panic(err)
		}

		if user.Status != usermodel.StatusActive {
			panic(appCommon.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(appCommon.DbTypeUser)

		c.Set(appCommon.CurrentUser, user)
		c.Next()
	}
}
