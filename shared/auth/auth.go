package auth

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/shared/auth/model"
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

type JwtInterface interface {
	GetHashedPassword() string
	GetLoginName() string
	GetStatus() int
	SetUid(string)
	GetUid() string
	GetToken() string
	SetToken(string)
	GetClaims() model.TokenClaim
	SetClaims(model.TokenClaim)
	SetVerifiedToken(*jwt.VerifiedToken)
	GetVerifiedToken() *jwt.VerifiedToken
}

func SignUp(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(model.CredUserInterface)
	cuser := model.CredUser{
		LoginName:      data.GetLoginName(),
		HashedPassword: data.GetHashedPassword(),
		Status:         data.GetStatus(),
	}

	dal.AddCredUser(cuser)
	c.Debugw("filled CredUser", cuser)
	return http.StatusOK
}

func ValidateUsernamePassword(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface)
	hpw := data.GetHashedPassword()
	loginName := data.GetLoginName()

	userInstance := dal.GetCredUser(loginName)
	if userInstance.GetHashedPassword() != hpw {
		return http.StatusForbidden
	}

	data.SetUid(userInstance.GetUid())

	c.Debugw("db user",userInstance)
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


func SignToken(c *streamline.ConveyorBelt) int {
	jwtdata := c.DataDomain.(JwtInterface)

	token, err := jwt.Sign(jwt.HS256, sharedKey, model.TokenClaim{
		Uid: jwtdata.GetUid(),
		Role: "admin",
	}, jwt.MaxAge(TokenExpireTime))
	if err != nil {
		panic(err)
	}

	jwtdata.SetToken(string(token))
	return http.StatusOK
}

func Validate(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(JwtInterface)
	token := data.GetToken()

	vtoken,err := jwt.Verify(jwt.HS256, sharedKey, []byte(token), blockList)
	if err != nil {
		return http.StatusForbidden
	}
	myclaims := model.TokenClaim{}
	err = vtoken.Claims(&myclaims)
	if err != nil {
		return http.StatusForbidden
	}

	data.SetClaims(myclaims)
	data.SetVerifiedToken(vtoken)

	c.Debugw(
		"token", token,
		"claims", c.DataDomain.(JwtInterface).GetClaims())

	return http.StatusOK
}
