package jwtc

import (
	"github.com/golang-jwt/jwt/v5"
	"sunflower-blog-svc/internal/pkg/ctxdata"
)

type Payload struct {
	Uid string
	Iat int64 // issued at，unix timestamp，token颁发时间
	Exp int64
}

func GenJwtToken(secretKey string, payload *Payload) (string, error) {
	iat := payload.Iat
	exp := payload.Exp
	userID := payload.Uid
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims[ctxdata.CtxKeyUid] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
