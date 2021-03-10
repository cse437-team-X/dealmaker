package auth

import (
	"github.com/kataras/jwt"
)

type JwtAuth struct {
	TokenClaim
	Token string
	VToken *jwt.VerifiedToken
}
func (j *JwtAuth) GetClaims() TokenClaim {
	return j.TokenClaim
}
func (j *JwtAuth) SetClaims(t TokenClaim) {
	j.TokenClaim = t
}
func (j *JwtAuth) GetToken() string {
	return j.Token
}
func (j *JwtAuth) SetToken(t string) {
	j.Token = t
}
func (j *JwtAuth) SetVerifiedToken(vt *jwt.VerifiedToken) {
	j.VToken = vt
}
func (j *JwtAuth) GetVerifiedToken() *jwt.VerifiedToken {
	return j.VToken
}

type TokenClaim struct {
	Uid string
	Role string
}