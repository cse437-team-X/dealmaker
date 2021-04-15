package model

import (
	"github.com/kataras/jwt"
)

const (
	JwtScopeNormal = "normal"
	JwtScopeActivate = "activate"
	JwtScopeRecover = "recover"
)

type JwtAuth struct {
	TokenClaim
	Token string
	VToken *jwt.VerifiedToken
}
func (j *JwtAuth) GetJwtAuth() *JwtAuth {
	return j
}


type TokenClaim struct {
	Uid uint
	Role string
	Scope string
	LoginName string
}

func (t *TokenClaim) GetTokenClaim() *TokenClaim {
	return t
}


type JwtInterface interface {
	GetJwtAuth() *JwtAuth
}