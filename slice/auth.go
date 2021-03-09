package slice

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/base_model"
	"github.com/dealmaker/base_model/obj"
	"github.com/dealmaker/dal"
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
	GetEmail() string
	GetStatus() int
}

type JwtInterface interface {
	GetToken() string
	SetToken(string)
	GetClaims() base_model.TokenClaim
	SetClaims(base_model.TokenClaim)
	SetVerifiedToken(*jwt.VerifiedToken)
	GetVerifiedToken() *jwt.VerifiedToken
}

func SignUp(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(UserInfoInterface)
	dal.AddUser(obj.User{
		Username:       data.GetUsername(),
		Email:          data.GetEmail(),
		HashedPassword: data.GetHashedPassword(),
		Status:         data.GetStatus(),
	})
	c.Logger.Debugw("signup", "filled user", obj.User{
		Username:       data.GetUsername(),
		Email:          data.GetEmail(),
		HashedPassword: data.GetHashedPassword(),
		Status:         data.GetStatus(),
	})
	return http.StatusOK
}

func ValidateUsernamePassword(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(UserInfoInterface)
	hpw := data.GetHashedPassword()
	email := data.GetEmail()

	if queryUsernamePassword(email, hpw) != true {
		return http.StatusForbidden
	}

	c.Logger.Debugw("Login",
		"email",email,
			"hashed_pw",hpw)
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


func queryUsernamePassword(email, hpw string) bool {
	u := dal.GetUser(email)
	return u.GetHashedPassword() == hpw
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
