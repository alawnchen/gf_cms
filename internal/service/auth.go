// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"

	jwt "github.com/gogf/gf-jwt/v2"
)

type IAuth interface {
	JWTAuth() *jwt.GfJWTMiddleware
	PayloadFunc(data interface{}) jwt.MapClaims
	IdentityHandler(ctx context.Context) interface{}
	Unauthorized(ctx context.Context, code int, message string)
	Authenticator(ctx context.Context) (interface{}, error)
}

var localAuth IAuth

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
