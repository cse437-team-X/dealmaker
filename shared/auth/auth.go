package auth

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/auth/model"
	"github.com/kataras/jwt"
	"net/http"
	"strconv"
	"time"
)

var blockList *jwt.Blocklist

const InvalidTokenExpireTime = 15 * time.Minute
const TokenExpireTime = 10 * time.Minute
var sharedKey = []byte("p@ssw0rd")

func init() {
	blockList = jwt.NewBlocklist(InvalidTokenExpireTime)
}

type CredUserInterface interface {
	GetCredUser() *model.CredUser
}


type JwtInterface interface {
	GetJwtAuth() *model.JwtAuth
}

func ValidateSignUp(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(CredUserInterface).GetCredUser()
	data.Role = model.UserRoleUser
	data.Status = model.UserStatusInactive

	c.Debugw("filled CredUser", data)
	return http.StatusOK
}

func ValidateCredUser(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(CredUserInterface).GetCredUser()
	hpw := data.HashedPassword
	loginName := data.LoginName
	c.Debugw("loginname",loginName,
		"hashedpw",hpw)
	return http.StatusOK
}

func Logout(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface).GetJwtAuth()
	token := data.VToken
	err := blockList.InvalidateToken(token.Token, token.StandardClaims)
	if err != nil {
		panic(err)
	}
	return http.StatusOK
}

func SignToken(c *streamline.ConveyorBelt) int {
	jwtdata := c.DataDomain.(JwtInterface).GetJwtAuth()
	credUserData := c.DataDomain.(auth_db.AuthDBInterface).GetUserCredModel()

	token, err := jwt.Sign(jwt.HS256, sharedKey, model.TokenClaim{
		Uid: strconv.FormatUint(uint64(credUserData.ID),10),
		Role: credUserData.Role,
	}, jwt.MaxAge(TokenExpireTime))
	if err != nil {
		panic(err)
	}

	jwtdata.Token = string(token)
	return http.StatusOK
}

func Validate(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface).GetJwtAuth()
	token := data.Token

	vtoken,err := jwt.Verify(jwt.HS256, sharedKey, []byte(token), blockList)
	if err != nil {
		return http.StatusForbidden
	}
	myclaims := model.TokenClaim{}
	err = vtoken.Claims(&myclaims)
	if err != nil {
		return http.StatusForbidden
	}

	data.TokenClaim = myclaims
	data.VToken = vtoken

	c.Debugw(
		"token", token,
		"claims", c.DataDomain.(JwtInterface).GetJwtAuth().TokenClaim)

	return http.StatusOK
}
