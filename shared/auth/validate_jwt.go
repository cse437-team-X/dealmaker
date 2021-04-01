package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"github.com/kataras/jwt"
	"net/http"
)

type ValidateJwtInterface interface {
	GetJwtAuth() *model.JwtAuth
}

func (w *WorkerInstance) ValidateJwt(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(ValidateJwtInterface).GetJwtAuth()
	token := data.Token

	vtoken, err := jwt.Verify(jwt.HS256, w.sharedKey, []byte(token), w.blockList)
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
		"claims", vtoken.StandardClaims)

	return http.StatusOK
}