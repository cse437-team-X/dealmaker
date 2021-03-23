package model

import (
	"github.com/kataras/jwt"
)

const (
	UserRoleUser = "user"
	UserRoleAdmin = "admin"
)

const (
	UserStatusInactive = 0
	UserStatusNormal = 1
	UserStatusOther = 2
)

type CredUser struct {
	//Uid string `gorm:"unique"`
	Role string
	HashedPassword string
	LoginName      string `gorm:"unique"`
	Status int
}
func (c *CredUser) GetCredUser() *CredUser {
	return c
}

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
}

func (t *TokenClaim) GetTokenClaim() *TokenClaim {
	return t
}


type JwtInterface interface {
	GetJwtAuth() *JwtAuth
}