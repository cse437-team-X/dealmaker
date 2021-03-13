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

type CredUserInterface interface {
	GetHashedPassword() string
	SetHashedPassword(v string)

	GetLoginName() string
	SetLoginName(v string)

	GetStatus() int
	SetStatus(v int)

	GetUid() string
	SetUid(v string)

	GetRole() string
	SetRole(string)
}


type JwtInterface interface {
	//GetHashedPassword() string
	//GetLoginName() string
	//GetStatus() int
	//SetUid(string)
	//GetUid() string
	GetToken() string
	SetToken(string)
	GetClaims() model.TokenClaim
	SetClaims(model.TokenClaim)
	SetVerifiedToken(*jwt.VerifiedToken)
	GetVerifiedToken() *jwt.VerifiedToken
}

func SignUp(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(CredUserInterface)
	cuser := model.CredUser{
		LoginName:      data.GetLoginName(),
		HashedPassword: data.GetHashedPassword(),
		Status:         data.GetStatus(),
		Role:           model.RoleUser,
	}

	err := dal.AddCredUser(cuser)
	if err != nil {
		c.Infow("Error", err.Error(), "req user", cuser)
		return http.StatusForbidden
	}
	c.Debugw("filled CredUser", cuser)
	return http.StatusOK
}

func ValidateCredUser(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(CredUserInterface)
	hpw := data.GetHashedPassword()
	loginName := data.GetLoginName()

	userInstance := dal.GetCredUser(loginName)
	if userInstance.GetHashedPassword() != hpw ||
		userInstance.GetStatus() == model.UserStatusInvalid {
		return http.StatusForbidden
	}

	data.SetUid(userInstance.GetUid())
	data.SetRole(userInstance.GetRole())
	data.SetStatus(userInstance.GetStatus())

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
	credUserData := c.DataDomain.(CredUserInterface)

	token, err := jwt.Sign(jwt.HS256, sharedKey, model.TokenClaim{
		Uid: credUserData.GetUid(),
		Role: credUserData.GetRole(),
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
