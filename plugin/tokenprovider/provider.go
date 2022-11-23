package tokenprovider

import (
	"cronbrowser/appCommon"
	"errors"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	//SecretKey() string
}

type TokenPayload interface {
	UserId() int64
	Role() string
}

type Token interface {
	GetToken() string
}

var (
	ErrNotFound = appCommon.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = appCommon.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = appCommon.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
