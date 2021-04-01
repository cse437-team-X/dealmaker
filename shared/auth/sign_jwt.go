package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"github.com/kataras/jwt"
	"net/http"
)

type SignTokenInterface interface {
	GetJwtAuth() *model.JwtAuth
	GetCredUser() *model.CredUser
}

func (w *WorkerInstance) SignTokenToScope(scope string) func (c *streamline.ConveyorBelt) int {
	return func(c *streamline.ConveyorBelt) int {
		jwtdata := c.DataDomain.(SignTokenInterface).GetJwtAuth()
		credUserData := c.DataDomain.(SignTokenInterface).GetCredUser()

		jwtdata.TokenClaim.Uid = credUserData.ID
		jwtdata.TokenClaim.Role = credUserData.Role
		jwtdata.TokenClaim.Scope = scope


		token, err := jwt.Sign(jwt.HS256, w.sharedKey, jwtdata.TokenClaim, jwt.MaxAge(w.TokenExpireTimes[scope]))
		if err != nil {
			panic(err)
		}

		jwtdata.Token = string(token)
		return http.StatusOK
	}
}