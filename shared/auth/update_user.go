package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UpdateUserInterface interface {
	GetCredUser() *model.CredUser
	GetJwtAuth() *model.JwtAuth
}

func (w *WorkerInstance) UpdateUser(c *streamline.ConveyorBelt) int {
	myInterface := c.DataDomain.(UpdateUserInterface)
	newUser := myInterface.GetCredUser()
	a := myInterface.GetJwtAuth()
	c.Debugw("old token claim", a.TokenClaim)
	newUser.ID = a.TokenClaim.Uid

	err := w.FuncUpdateCredUser(newUser)
	if err != nil {
		c.Errorw("update cred user failed", err)
		return http.StatusInternalServerError
	}
	c.Infow("cred user updated", newUser)
	return http.StatusOK
}