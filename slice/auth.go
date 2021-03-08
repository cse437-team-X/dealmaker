package slice

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/base_model"
	"github.com/kataras/jwt"
	"net/http"
	"time"
)

var blockList *jwt.Blocklist

const InvalidTokenExpireTime = 15 * time.Minute
const TokenExpireTime = 10 * time.Minute
var sharedKey = []byte("p@ssw0rd")

func init() {
	blockList = jwt.NewBlocklist(InvalidTokenExpireTime)
}

type UserInfoInterface interface {
	GetHashedPassword() string
	GetUsername() string
}

type JwtInterface interface {
	GetToken() string
	SetToken(string)
	GetClaims() base_model.TokenClaim
	SetClaims(base_model.TokenClaim)
	SetVerifiedToken(*jwt.VerifiedToken)
	GetVerifiedToken() *jwt.VerifiedToken
}

//func SignUp(c *streamline.ConveyorBelt) int {
//
//}

func ValidateUsernamePassword(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(UserInfoInterface)
	d1 := data.GetHashedPassword()
	d2 := data.GetUsername()

	if queryUsernamePassword(d2, d1) != true {
		return http.StatusForbidden
	}

	c.Logger.Debugw("Login",
		"username",d1,
			"hashed_pw",d2)
	return http.StatusOK
}

func Logout(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface)
	token := data.GetVerifiedToken()
	err := blockList.InvalidateToken(token.Token, token.StandardClaims)
	if err != nil {
		panic(err)
	}
	return http.StatusOK
}


func queryUsernamePassword(username, hashpw string) bool {
	if username == "admin" && hashpw == "admin" {
		return true
	}
	return false
}

func SignToken(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface)

	token, err := jwt.Sign(jwt.HS256, sharedKey, base_model.TokenClaim{
		Uid: "0x00000000",
		Role: "admin",
	}, jwt.MaxAge(TokenExpireTime))
	if err != nil {
		panic(err)
	}

	data.SetToken(string(token))
	return http.StatusOK
}

func Validate(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface)
	token := data.GetToken()

	vtoken,err := jwt.Verify(jwt.HS256, sharedKey, []byte(token), blockList)
	if err != nil {
		return http.StatusForbidden
	}
	myclaims := base_model.TokenClaim{}
	err = vtoken.Claims(&myclaims)
	if err != nil {
		return http.StatusForbidden
	}

	data.SetClaims(myclaims)
	data.SetVerifiedToken(vtoken)

	c.Logger.Debugw("Authenticator",
		"token", token,
		"claims", c.DataDomain.(JwtInterface).GetClaims())

	return http.StatusOK
}
