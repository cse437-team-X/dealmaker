package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/itzmeerkat/streamline"
	"github.com/kataras/jwt"
	"net/http"
)

type ValidateJwtInterface interface {
	GetJwtAuth() *model.JwtAuth
	GetBase() *base.Base
}

func (w *WorkerInstance) ValidateJwt(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(ValidateJwtInterface).GetJwtAuth()
	b := c.DataDomain.(ValidateJwtInterface).GetBase()
	token := data.Token

	vtoken, err := jwt.Verify(jwt.HS256, w.sharedKey, []byte(token), w.blockList)
	if err != nil {
		b.BaseMessage = err.Error()
		c.Infow("jwt", b.BaseMessage)
		return http.StatusForbidden
	}
	myclaims := model.TokenClaim{}
	err = vtoken.Claims(&myclaims)
	if err != nil {
		b.BaseMessage = err.Error()
		return http.StatusForbidden
	}

	data.TokenClaim = myclaims
	data.VToken = vtoken

	c.Debugw(
		"token", token,
		"claims", vtoken.StandardClaims)

	return http.StatusOK
}