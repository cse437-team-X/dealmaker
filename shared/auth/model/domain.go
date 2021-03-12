package model

import (
	"github.com/kataras/jwt"
)

type CredUserInterface interface {
	GetHashedPassword() string
	SetHashedPassword(v string)
	GetLoginName() string
	SetLoginName(v string)
	GetStatus() int
	SetStatus(v int)
	GetUid() string
	SetUid(v string)
}

type CredUser struct {
	Uid string
	HashedPassword string
	LoginName      string
	Status int
}

func (u *CredUser) GetHashedPassword() string {
	return u.HashedPassword
}
func (u *CredUser) GetLoginName() string {
	return u.LoginName
}
func (u *CredUser) SetHashedPassword(v string) {
	u.HashedPassword = v
}
func (u *CredUser) SetLoginName(v string) {
	u.LoginName = v
}
func (u *CredUser) GetStatus() int {
	return u.Status
}
func (u *CredUser) SetStatus(v int) {
	u.Status = v
}
func (u *CredUser) GetUid() string {
	return u.Uid
}
func (u *CredUser) SetUid(v string) {
	u.Uid = v
}


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

func (t *TokenClaim) GetUid() string {
	return t.Uid
}
func (t *TokenClaim) SetUid(v string) {
	t.Uid = v
}
func (t *TokenClaim) GetRole() string {
	return t.Role
}
func (t *TokenClaim) SetRole(v string) {
	t.Role = v
}